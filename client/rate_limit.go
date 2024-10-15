package client

import (
	"net/http"
	"strconv"
	"strings"
)

type RateLimitInfo struct {
	Limit     *int
	Remaining *int
}

func ParseRateLimitInfo(response *http.Response) *RateLimitInfo {
	var rate RateLimitInfo

	header := response.Header.Get("X-Rate-Limit")

	if header == "" {
		return nil
	}

	for _, rateLimitPart := range strings.Split(header, ";") {
		if rateLimitPart == "" {
			continue
		}

		rateLimitPartInfo := strings.Split(rateLimitPart, ":")
		if len(rateLimitPartInfo) != 2 {
			continue
		}

		rateLimitKey := rateLimitPartInfo[0]
		rateLimitValue, err := strconv.Atoi(rateLimitPartInfo[1])

		if err != nil {
			continue
		}

		switch rateLimitKey {
		case "user-hour-lim":
			rate.Limit = &rateLimitValue
		case "user-hour-rem":
			rate.Remaining = &rateLimitValue
		}
	}

	if rate.Limit == nil && rate.Remaining == nil {
		return nil
	}

	return &rate
}
