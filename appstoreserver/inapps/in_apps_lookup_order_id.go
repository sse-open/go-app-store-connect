package inapps

import (
	"context"
	"fmt"

	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/inapps"
	"github.com/sse-open/go-app-store-connect/client/response"
)

// https://developer.apple.com/documentation/appstoreserverapi/look-up-order-id
func (iaps *InAppsService) LookUpOrderID(ctx context.Context, orderID string) (*inapps.OrderLookupResponse, *response.ClientResponse, error) {
	url := fmt.Sprintf("inApps/v1/lookup/%s", orderID)
	respPayload := &inapps.OrderLookupResponse{}
	resp, err := iaps.client.Get(ctx, url, nil, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}
