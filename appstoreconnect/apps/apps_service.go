package apps

import (
	"context"

	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/response"
)

//go:generate mockery --name IAppsService
type IAppsService interface {
	ListApps(ctx context.Context, queryParams *ListAppsQuery) (*AppsResponse, *response.ClientResponse, error)
}

type AppsService struct {
	client *client.Client
}

func NewAppsService(client *client.Client) *AppsService {
	return &AppsService{
		client: client,
	}
}
