package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/nbio/st"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/sse-open/go-app-store-connect/client/request"
	"github.com/stretchr/testify/assert"
)

type TestQuery struct {
	TestParam string `url:"testParam,omitempty"`
}

type TestPayload struct {
	TestField string `json:"testField,omitempty"`
}

type TestResponse struct {
	TestResponseField string `json:"testResponseField,omitempty"`
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	c, err := NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	assert.Equal(t, defaultBaseURL, c.baseURL.String())
}

func TestClientSetBaseURL(t *testing.T) {
	t.Parallel()

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	c, err := NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	c.SetBaseURL("https://example.com")

	assert.Equal(t, "https://example.com", c.baseURL.String())
}

func TestClientGet(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	gock.New("https://api.appstoreconnect.apple.com").
		Get("/test").
		MatchParam("testParam", "abc123").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(200).
		JSON(map[string]string{"testResponseField": "testResponseValue"})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	queryParams := TestQuery{
		TestParam: "abc123",
	}
	var testRes TestResponse
	resp, err := c.Get(ctx, "test", queryParams, &testRes)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "testResponseValue", testRes.TestResponseField)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestClientGetNotFound(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	gock.New("https://api.appstoreconnect.apple.com").
		Get("/test").
		MatchParam("testParam", "abc123").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(404).JSON("{}")

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	queryParams := TestQuery{
		TestParam: "abc123",
	}
	var testRes TestResponse
	resp, err := c.Get(ctx, "test", queryParams, &testRes)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Empty(t, testRes.TestResponseField)

	var errorResponse ErrorResponse
	assert.True(t, errors.As(err, &errorResponse))
	assert.Equal(t, 404, errorResponse.Response.StatusCode)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestClientPost(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	gock.New("https://api.appstoreconnect.apple.com").
		Post("/test").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(200).
		JSON(map[string]string{"testResponseField": "testResponseValue"})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	testPayload := TestPayload{
		TestField: "testFieldValue",
	}
	appStoreConnectRequestPayload := request.AppStoreConnectRequestPayload{
		Data: testPayload,
	}
	var testRes TestResponse
	resp, err := c.Post(ctx, "test", &appStoreConnectRequestPayload, &testRes)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "testResponseValue", testRes.TestResponseField)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestClientPatch(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	gock.New("https://api.appstoreconnect.apple.com").
		Patch("/test").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(200).
		JSON(map[string]string{"testResponseField": "testResponseValue"})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	testPayload := TestPayload{
		TestField: "testFieldValue",
	}
	appStoreConnectRequestPayload := request.AppStoreConnectRequestPayload{
		Data: testPayload,
	}
	var testRes TestResponse
	resp, err := c.Patch(ctx, "test", &appStoreConnectRequestPayload, &testRes)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "testResponseValue", testRes.TestResponseField)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestClientDelete(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	gock.New("https://api.appstoreconnect.apple.com").
		Delete("/test").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(204)

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	resp, err := c.Delete(ctx, "test")
	assert.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestHandleErrorResponse(t *testing.T) {
	response := http.Response{
		StatusCode: 404,
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{},
		},
		Body: io.NopCloser(strings.NewReader(`{
			"errors": [
				  {
					"id": "123456789",
					"status": "404",
					"code": "NOT_FOUND",
					"title": "Resource not found.",
					"detail": "The requests resource was not found.",
					"meta": {
						  "associatedErrors": {
							"/v1/route/": [
								  {
									"id": "987654321",
									"status": "400",
									"code": "BAD_REQUEST",
									"title": "Bad Request",
									"detail": "Invalid input for field test"
								  }
							]
						  }
					}
				  }
			]
		}`)),
	}
	err := handleErrorResponse(&response)

	var errorResponse ErrorResponse
	assert.True(t, errors.As(err, &errorResponse))
	assert.Equal(t, 404, errorResponse.Response.StatusCode)
	assert.Equal(t, "123456789", *errorResponse.Errors[0].ID)
	assert.Equal(t, "404", errorResponse.Errors[0].Status)
	assert.Equal(t, "Resource not found.", errorResponse.Errors[0].Title)
}

func TestHandleErrorResponseStatusOK(t *testing.T) {
	response := http.Response{
		StatusCode: 200,
	}
	err := handleErrorResponse(&response)

	assert.NoError(t, err)
}

func TestHandleErrorResponseRateLimitExceeded(t *testing.T) {
	response := http.Response{
		StatusCode: 429,
		Header: http.Header{
			"X-Rate-Limit": []string{"user-hour-lim:100;user-hour-rem:0"},
		},
	}
	err := handleErrorResponse(&response)

	expectedErr := ErrRateLimitExceeded
	expectedErr.rateLimitLimit = 100
	expectedErr.rateLimitRemaining = 0
	assert.True(t, errors.Is(err, expectedErr))
}

func TestHandleErrorResponseRateLimitMissing(t *testing.T) {
	response := http.Response{
		StatusCode: 429,
	}
	err := handleErrorResponse(&response)

	expectedErr := ErrRateLimitExceeded
	assert.True(t, errors.Is(err, expectedErr))
}
