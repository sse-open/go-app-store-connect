package appstoreconnect

import (
	"net/http"

	errorsPkg "github.com/pkg/errors"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/apps"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/inapppurchase"
	"github.com/sse-open/go-app-store-connect/client"
)

type IAppStoreConnect interface {
	AppsService() apps.IAppsService
	InAppPurchaseService() inapppurchase.IInAppPurchaseService
}

type AppStoreConnect struct {
	client *client.Client
}

func NewAppStoreConnect(httpClient *http.Client, jwtProvider client.IJWTProvider) (*AppStoreConnect, error) {
	client, err := client.NewConnectClient(httpClient, jwtProvider)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to create client")
	}

	return &AppStoreConnect{
		client: client,
	}, nil
}

func (asc *AppStoreConnect) AppsService() apps.IAppsService {
	return apps.NewAppsService(asc.client)
}

func (asc *AppStoreConnect) InAppPurchaseService() inapppurchase.IInAppPurchaseService {
	return inapppurchase.NewInAppPurchaseService(asc.client)
}
