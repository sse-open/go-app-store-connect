package appstoreserver

import (
	"testing"

	"github.com/sse-open/go-app-store-connect/appstoreserver/inapps"
	inAppsMocks "github.com/sse-open/go-app-store-connect/appstoreserver/inapps/mocks"
)

type MockAppStoreServer struct {
	MockInAppsService *inAppsMocks.IInAppsService
}

func NewMockSDK(t *testing.T) (*MockAppStoreServer, error) {
	return &MockAppStoreServer{
		MockInAppsService: inAppsMocks.NewIInAppsService(t),
	}, nil
}

func (c *MockAppStoreServer) InAppsService() inapps.IInAppsService {
	return c.MockInAppsService
}
