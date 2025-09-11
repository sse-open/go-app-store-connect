package inapps

import (
	"context"

	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/inapps"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/response"
)

//go:generate mockery --name IInAppsService
type IInAppsService interface {
	LookUpOrderID(ctx context.Context, orderID string) (*inapps.OrderLookupResponse, *response.ClientResponse, error)
	GetTransactionInfo(ctx context.Context, transactionId string) (*inapps.TransactionInfoResponse, *response.ClientResponse, error)
}

type InAppsService struct {
	client *client.Client
}

func NewInAppsService(client *client.Client) *InAppsService {
	return &InAppsService{
		client: client,
	}
}
