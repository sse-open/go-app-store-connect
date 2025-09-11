package appstoreserver

import (
	"testing"

	"github.com/sse-open/go-app-store-connect/appstoreserver/inapps"
	inAppsMocks "github.com/sse-open/go-app-store-connect/appstoreserver/inapps/mocks"
	"github.com/sse-open/go-app-store-connect/appstoreserver/subscription"
	subscriptionMocks "github.com/sse-open/go-app-store-connect/appstoreserver/subscription/mocks"
)

type MockAppStoreServer struct {
	MockInAppsService       *inAppsMocks.IInAppsService
	MockSubscriptionService *subscriptionMocks.ISubscriptionService
}

func NewMockSDK(t *testing.T) (*MockAppStoreServer, error) {
	return &MockAppStoreServer{
		MockInAppsService:       inAppsMocks.NewIInAppsService(t),
		MockSubscriptionService: subscriptionMocks.NewISubscriptionService(t),
	}, nil
}

func (c *MockAppStoreServer) InAppsService() inapps.IInAppsService {
	return c.MockInAppsService
}

func (c *MockAppStoreServer) SubscriptionService() subscription.ISubscriptionService {
	return c.MockSubscriptionService
}
