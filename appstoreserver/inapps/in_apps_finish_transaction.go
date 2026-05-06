package inapps

import (
	"context"
	"fmt"

	"github.com/sse-open/go-app-store-connect/client/response"
)

// https://developer.apple.com/documentation/appstoreserverapi/finish-transaction
func (iaps *InAppsService) FinishTransaction(ctx context.Context, transactionId string) (*response.ClientResponse, error) {
	url := fmt.Sprintf("inApps/v1/transactions/%s/finish", transactionId)
	resp, err := iaps.client.Post(ctx, url, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
