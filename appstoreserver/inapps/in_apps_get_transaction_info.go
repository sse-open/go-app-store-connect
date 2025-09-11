package inapps

import (
	"context"
	"fmt"

	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/inapps"
	"github.com/sse-open/go-app-store-connect/client/response"
)

// https://developer.apple.com/documentation/appstoreserverapi/get-transaction-info
func (iaps *InAppsService) GetTransactionInfo(ctx context.Context, transactionId string) (*inapps.TransactionInfoResponse, *response.ClientResponse, error) {
	url := fmt.Sprintf("inApps/v1/transactions/%s", transactionId)
	respPayload := &inapps.TransactionInfoResponse{}
	resp, err := iaps.client.Get(ctx, url, nil, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}
