package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	errorsPkg "github.com/pkg/errors"
	"github.com/sse-open/go-app-store-connect/client/ratelimit"
	"github.com/sse-open/go-app-store-connect/client/request"
	"github.com/sse-open/go-app-store-connect/client/response"
)

const (
	defaultBaseURL = "https://api.appstoreconnect.apple.com/"
)

type ErrorWithRateLimit struct {
	message            string
	rateLimitLimit     int
	rateLimitRemaining int
}

func (e ErrorWithRateLimit) Error() string {
	if e.rateLimitLimit > 0 {
		return fmt.Sprintf("%s. Rate limit: %d/%d", e.message, e.rateLimitRemaining, e.rateLimitLimit)
	}
	return e.message
}

var ErrRateLimitExceeded = ErrorWithRateLimit{"hourly rate limit exceeded", 0, 0}

//go:generate mockery --name IClient
type IClient interface {
	SetBaseURL(baseURL string)
	Get(ctx context.Context, path string, query interface{}, respPayload interface{}) (*response.ClientResponse, error)
	Post(ctx context.Context, path string, body *request.AppStoreConnectRequestPayload, respPayload interface{}) (*response.ClientResponse, error)
	Patch(ctx context.Context, path string, body *request.AppStoreConnectRequestPayload, respPayload interface{}) (*response.ClientResponse, error)
	Delete(ctx context.Context, path string) (*response.ClientResponse, error)
}

type Client struct {
	client      *http.Client
	baseURL     *url.URL
	jwtProvider IJWTProvider
}

func NewClient(httpClient *http.Client, jwtProvider IJWTProvider) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to parse default base URL")
	}

	c := &Client{
		client:      httpClient,
		baseURL:     baseURL,
		jwtProvider: jwtProvider,
	}

	return c, nil
}

func (c *Client) SetBaseURL(baseURL string) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return
	}
	c.baseURL = u
}

func (c *Client) Get(ctx context.Context, path string, query interface{}, respPayload interface{}) (*response.ClientResponse, error) {
	resp, err := c.createAndExecuteRequest(ctx, "GET", path, query, nil, respPayload)

	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to perform GET request")
	}

	return resp, nil
}

func (c *Client) Post(ctx context.Context, path string, body *request.AppStoreConnectRequestPayload, respPayload interface{}) (*response.ClientResponse, error) {
	resp, err := c.createAndExecuteRequest(ctx, "POST", path, nil, body, respPayload)

	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to perform POST request")
	}

	return resp, nil
}

func (c *Client) Patch(ctx context.Context, path string, body *request.AppStoreConnectRequestPayload, respPayload interface{}) (*response.ClientResponse, error) {
	resp, err := c.createAndExecuteRequest(ctx, "PATCH", path, nil, body, respPayload)

	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to perform PATCH request")
	}

	return resp, nil
}

func (c *Client) Delete(ctx context.Context, path string) (*response.ClientResponse, error) {
	resp, err := c.createAndExecuteRequest(ctx, "DELETE", path, nil, nil, nil)

	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to perform DELETE request")
	}

	return resp, nil
}

func (c *Client) createAndExecuteRequest(ctx context.Context, method string, path string, query interface{}, body *request.AppStoreConnectRequestPayload, respPayload interface{}) (*response.ClientResponse, error) {
	req, err := c.newHTTPRequest(ctx, method, path, query, body)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to create new HTTP request")
	}

	resp, err := c.executeHTTPRequest(req, respPayload)
	if err != nil {
		return resp, errorsPkg.Wrap(err, "failed to execute HTTP request")
	}

	return resp, err
}

func (c *Client) newHTTPRequest(ctx context.Context, method string, path string, queryParameters interface{}, body *request.AppStoreConnectRequestPayload) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to parse path")
	}

	var u *url.URL
	if rel.IsAbs() {
		u = rel
	} else {
		u = c.baseURL.ResolveReference(rel)
	}

	if queryParameters != nil {
		qs, err := query.Values(queryParameters)
		if err != nil {
			return nil, errorsPkg.Wrap(err, "failed to serialize query parameters")
		}

		u.RawQuery = qs.Encode()
	}

	buf := new(bytes.Buffer)

	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	token, err := c.jwtProvider.GetJWTToken()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) executeHTTPRequest(request *http.Request, respPayload interface{}) (*response.ClientResponse, error) {
	resp, err := c.client.Do(request)

	if err != nil {
		if resp != nil {
			return response.NewResponse(resp), err
		} else {
			return nil, err
		}
	}

	if err := handleErrorResponse(resp); err != nil {
		return nil, err
	}

	if respPayload != nil && resp.StatusCode != http.StatusNoContent {
		err = json.NewDecoder(resp.Body).Decode(respPayload)
		if err != nil {
			return nil, err
		}
	}
	return response.NewResponse(resp), nil
}

func handleErrorResponse(response *http.Response) error {
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		return nil
	}

	if response.StatusCode == 429 {
		rateLimitInfo := ratelimit.ParseRateLimitInfo(response)
		errRateLimitExceeded := ErrRateLimitExceeded
		if rateLimitInfo != nil {
			errRateLimitExceeded.rateLimitLimit = *rateLimitInfo.Limit
			errRateLimitExceeded.rateLimitRemaining = *rateLimitInfo.Remaining
		}
		return errRateLimitExceeded
	}

	var errorResponse ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&errorResponse)
	if err != nil {
		return err
	}
	errorResponse.Response = response

	return errorResponse
}
