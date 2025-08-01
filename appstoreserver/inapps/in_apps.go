package inapps

import (
	"context"
	"fmt"

	"github.com/sse-open/go-app-store-connect/appstoreserver/inapps/resource"
	"github.com/sse-open/go-app-store-connect/client/response"
)

// The status that indicates whether the order ID is valid.
//
// https://developer.apple.com/documentation/appstoreserverapi/orderlookupstatus
type OrderLookupStatus int

func (s OrderLookupStatus) IsSuccess() bool {
	return s == OrderLookupStatusSuccess
}

var (
	OrderLookupStatusSuccess OrderLookupStatus = 0
	OrderLookupStatusFailure OrderLookupStatus = 1
)

// Contains information about an in-app purchase order.
//
// https://developer.apple.com/documentation/appstoreserverapi/orderlookupresponse
type OrderLookupResponse struct {
	Status             OrderLookupStatus         `json:"status"`
	SignedTransactions []resource.JWSTransaction `json:"signedTransactions"`
}

func (iaps *InAppsService) LookUpOrderID(ctx context.Context, orderID string) (*OrderLookupResponse, *response.ClientResponse, error) {
	url := fmt.Sprintf("inApps/v1/lookup/%s", orderID)
	respPayload := &OrderLookupResponse{}
	resp, err := iaps.client.Get(ctx, url, nil, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}
