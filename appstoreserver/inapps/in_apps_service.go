package inapps

import (
	"context"

	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/response"
)

//go:generate mockery --name IInAppsService
type IInAppsService interface {
	LookUpOrderID(ctx context.Context, orderID string) (*OrderLookupResponse, *response.ClientResponse, error)
}

type InAppsService struct {
	client *client.Client
}

func NewInAppsService(client *client.Client) *InAppsService {
	return &InAppsService{
		client: client,
	}
}
