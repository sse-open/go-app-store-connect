package subscription

import (
	"context"

	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/subscription"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/response"
)

//go:generate mockery --name ISubscriptionService
type ISubscriptionService interface {
	GetSubscriptionStatus(ctx context.Context, transactionId string, queryParams *GetSubscriptionStatusQuery) (*subscription.StatusResponse, *response.ClientResponse, error)
	ExtendSubscriptionRenewalDate(ctx context.Context, originalTransactionId string, req *subscription.ExtendRenewalDateRequest) (*subscription.ExtendRenewalDateResponse, *response.ClientResponse, error)
	MassExtendSubscriptionRenewalDates(ctx context.Context, req *subscription.MassExtendRenewalDateRequest) (*subscription.MassExtendRenewalDateResponse, *response.ClientResponse, error)
}

type SubscriptionService struct {
	client *client.Client
}

func NewSubscriptionService(client *client.Client) *SubscriptionService {
	return &SubscriptionService{
		client: client,
	}
}
