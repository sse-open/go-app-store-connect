package response

import (
	"net/http"

	"github.com/sse-open/go-app-store-connect/client/ratelimit"
)

type ClientResponse struct {
	*http.Response
	RateLimitInfo *ratelimit.RateLimitInfo
}

func NewResponse(response *http.Response) *ClientResponse {
	return &ClientResponse{
		Response:      response,
		RateLimitInfo: ratelimit.ParseRateLimitInfo(response),
	}
}
