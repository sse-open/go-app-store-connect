package appstoreconnect

import (
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

func NewAppStoreConnect(jwtProvider client.IJWTProvider) (*AppStoreConnect, error) {
	client, err := client.NewClient(nil, jwtProvider)
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
