package inapppurchase

import (
	"context"

	"github.com/sse-open/go-app-store-connect/client"
)

type IInAppPurchaseService interface {
	ListAppInAppPurchases(ctx context.Context, appID string, queryParams *ListAppInAppPurchasesQuery) (*InAppPurchasesV2Response, *client.ClientResponse, error)
	ListInAppPurchaseManualPrices(ctx context.Context, inAppPurchaseID string, queryParams *ListInAppPurchaseManualPricesQuery) (*InAppPurchasePricesResponse, *client.ClientResponse, error)
	ListInAppPurchaseAutomaticPrices(ctx context.Context, inAppPurchaseID string, queryParams *ListInAppPurchaseAutomaticPricesQuery) (*InAppPurchasePricesResponse, *client.ClientResponse, error)
}

type InAppPurchaseService struct {
	client *client.Client
}

func NewInAppPurchaseService(client *client.Client) *InAppPurchaseService {
	return &InAppPurchaseService{
		client: client,
	}
}
