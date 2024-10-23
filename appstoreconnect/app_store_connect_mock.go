package appstoreconnect

import (
	"testing"

	appsMocks "github.com/sse-open/go-app-store-connect/appstoreconnect/apps/mocks"
	inAppPurchaseMocks "github.com/sse-open/go-app-store-connect/appstoreconnect/inapppurchase/mocks"
)

type MockAppStoreConnect struct {
	MockAppsService          *appsMocks.IAppsService
	MockInAppPurchaseService *inAppPurchaseMocks.IInAppPurchaseService
}

func NewMockSDK(t *testing.T) (*MockAppStoreConnect, error) {
	return &MockAppStoreConnect{
		MockAppsService:          appsMocks.NewIAppsService(t),
		MockInAppPurchaseService: inAppPurchaseMocks.NewIInAppPurchaseService(t),
	}, nil
}

func (c *MockAppStoreConnect) AppsService() *appsMocks.IAppsService {
	return c.MockAppsService
}

func (c *MockAppStoreConnect) InAppPurchaseService() *inAppPurchaseMocks.IInAppPurchaseService {
	return c.MockInAppPurchaseService
}
