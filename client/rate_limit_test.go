package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRateLimitInfo(t *testing.T) {
	t.Parallel()

	inputResponse := http.Response{
		Header: map[string][]string{
			"X-Rate-Limit": {"user-hour-lim:3500;user-hour-rem:500;"},
		},
	}
	rateLimitInfo := ParseRateLimitInfo(&inputResponse)
	assert.NotNil(t, rateLimitInfo)

	assert.Equal(t, 3500, *rateLimitInfo.Limit)
	assert.Equal(t, 500, *rateLimitInfo.Remaining)
}

func TestParseRateLimitMissingHeader(t *testing.T) {
	t.Parallel()

	inputResponse := http.Response{}
	rateLimitInfo := ParseRateLimitInfo(&inputResponse)
	assert.Nil(t, rateLimitInfo)
}

func TestParseRateLimitInfoExtra(t *testing.T) {
	t.Parallel()

	inputResponse := http.Response{
		Header: map[string][]string{
			"X-Rate-Limit": {"user-hour-lim:3500;user-hour-rem:500;extra_to:drop;"},
		},
	}
	rateLimitInfo := ParseRateLimitInfo(&inputResponse)
	assert.NotNil(t, rateLimitInfo)

	assert.Equal(t, 3500, *rateLimitInfo.Limit)
	assert.Equal(t, 500, *rateLimitInfo.Remaining)
}

func TestParseRateLimitInvalidData(t *testing.T) {
	t.Parallel()

	inputResponse := http.Response{
		Header: map[string][]string{
			"X-Rate-Limit": {"user-hour-lim:abc;user-hour-rem:def;"},
		},
	}
	rateLimitInfo := ParseRateLimitInfo(&inputResponse)
	assert.Nil(t, rateLimitInfo)
}
