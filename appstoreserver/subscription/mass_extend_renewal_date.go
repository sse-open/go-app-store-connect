package subscription

import (
	"context"

	resourcesubscription "github.com/sse-open/go-app-store-connect/appstoreserver/resource/subscription"
	"github.com/sse-open/go-app-store-connect/client/response"
)

// https://developer.apple.com/documentation/appstoreserverapi/extend-subscription-renewal-dates-for-all-active-subscribers
func (ss *SubscriptionService) MassExtendSubscriptionRenewalDates(ctx context.Context, req *resourcesubscription.MassExtendRenewalDateRequest) (*resourcesubscription.MassExtendRenewalDateResponse, *response.ClientResponse, error) {
	url := "inApps/v1/subscriptions/extend/mass"
	respPayload := &resourcesubscription.MassExtendRenewalDateResponse{}
	resp, err := ss.client.Post(ctx, url, req, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}
