package subscription

import (
	"context"
	"fmt"

	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/subscription"
	"github.com/sse-open/go-app-store-connect/client/response"
)

// https://developer.apple.com/documentation/appstoreserverapi/get-all-subscription-statuses
type GetSubscriptionStatusQuery struct {
	Status []subscription.AutoRenewableStatus `url:"status,omitempty"`
}

// https://developer.apple.com/documentation/appstoreserverapi/get-all-subscription-statuses
func (ss *SubscriptionService) GetSubscriptionStatus(ctx context.Context, transactionId string, queryParams *GetSubscriptionStatusQuery) (*subscription.StatusResponse, *response.ClientResponse, error) {
	url := fmt.Sprintf("inApps/v1/subscriptions/%s", transactionId)
	respPayload := &subscription.StatusResponse{}
	resp, err := ss.client.Get(ctx, url, queryParams, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}
