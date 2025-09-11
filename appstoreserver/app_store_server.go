package appstoreserver

import (
	"net/http"

	errorsPkg "github.com/pkg/errors"
	"github.com/sse-open/go-app-store-connect/appstoreserver/inapps"
	"github.com/sse-open/go-app-store-connect/appstoreserver/subscription"
	"github.com/sse-open/go-app-store-connect/client"
)

type IAppStoreServer interface {
	InAppsService() inapps.IInAppsService
	SubscriptionService() subscription.ISubscriptionService
}

type AppStoreServer struct {
	client *client.Client
}

func NewAppStoreServer(httpClient *http.Client, jwtProvider client.IJWTProvider, sandbox bool) (*AppStoreServer, error) {
	client, err := client.NewServerClient(httpClient, jwtProvider, sandbox)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to create client")
	}

	return &AppStoreServer{
		client: client,
	}, nil
}

func (asc *AppStoreServer) InAppsService() inapps.IInAppsService {
	return inapps.NewInAppsService(asc.client)
}

func (asc *AppStoreServer) SubscriptionService() subscription.ISubscriptionService {
	return subscription.NewSubscriptionService(asc.client)
}
