package appstoreserver

import (
	"net/http"

	errorsPkg "github.com/pkg/errors"
	"github.com/sse-open/go-app-store-connect/client"
)

type IAppStoreServer interface {
}

type AppStoreServer struct {
	client *client.Client
}

func NewAppStoreServer(httpClient *http.Client, jwtProvider client.IJWTProvider) (*AppStoreServer, error) {
	client, err := client.NewServerClient(httpClient, jwtProvider)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to create client")
	}

	return &AppStoreServer{
		client: client,
	}, nil
}

// func (asc *AppStoreServer) InAppsService() apps.IAppsService {
// 	return apps.NewAppsService(asc.client)
// }
