package subscription

import (
	"context"
	"fmt"

	resourcesubscription "github.com/sse-open/go-app-store-connect/appstoreserver/resource/subscription"
	"github.com/sse-open/go-app-store-connect/client/response"
)

// https://developer.apple.com/documentation/appstoreserverapi/extend-a-subscription-renewal-date
func (ss *SubscriptionService) ExtendSubscriptionRenewalDate(ctx context.Context, originalTransactionId string, req *resourcesubscription.ExtendRenewalDateRequest) (*resourcesubscription.ExtendRenewalDateResponse, *response.ClientResponse, error) {
	url := fmt.Sprintf("inApps/v1/subscriptions/extend/%s", originalTransactionId)
	respPayload := &resourcesubscription.ExtendRenewalDateResponse{}
	resp, err := ss.client.Put(ctx, url, req, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}
