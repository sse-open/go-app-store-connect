package inapppurchase

import (
	"context"

	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/response"
)

//go:generate mockery --name IInAppPurchaseService
type IInAppPurchaseService interface {
	ListAppInAppPurchases(ctx context.Context, appID string, queryParams *ListAppInAppPurchasesQuery) (*InAppPurchasesV2Response, *response.ClientResponse, error)
	ListInAppPurchaseManualPrices(ctx context.Context, inAppPurchaseID string, queryParams *ListInAppPurchaseManualPricesQuery) (*InAppPurchasePricesResponse, *response.ClientResponse, error)
	ListInAppPurchaseAutomaticPrices(ctx context.Context, inAppPurchaseID string, queryParams *ListInAppPurchaseAutomaticPricesQuery) (*InAppPurchasePricesResponse, *response.ClientResponse, error)
}

type InAppPurchaseService struct {
	client *client.Client
}

func NewInAppPurchaseService(client *client.Client) *InAppPurchaseService {
	return &InAppPurchaseService{
		client: client,
	}
}
