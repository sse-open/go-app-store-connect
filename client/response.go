package client

import "net/http"

type ClientResponse struct {
	*http.Response
	RateLimitInfo *RateLimitInfo
}

func NewResponse(response *http.Response) *ClientResponse {
	return &ClientResponse{
		Response:      response,
		RateLimitInfo: ParseRateLimitInfo(response),
	}
}
