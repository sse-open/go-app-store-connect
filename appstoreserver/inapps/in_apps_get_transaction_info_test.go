package inapps

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/common"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionInfo(t *testing.T) {
	t.Run("get success", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		transactionID := "transaction123"

		gock.New("https://api.appstoreconnect.apple.com").
			Get(fmt.Sprintf("/inApps/v1/transactions/%s", transactionID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(200).
			JSON(map[string]any{
				"signedTransactionInfo": "xxx.yyy.zzz",
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewConnectClient(nil, mockedJWTProvider)
		assert.NoError(t, err)

		appsService := NewInAppsService(c)

		response, clientResponse, err := appsService.GetTransactionInfo(ctx, transactionID)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, clientResponse)

		assert.Equal(t, common.JWSTransaction("xxx.yyy.zzz"), response.SignedTransactionInfo)
	})

	t.Run("invalid transaction id", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		transactionID := "invalidTransactionID"

		gock.New("https://api.appstoreconnect.apple.com").
			Get(fmt.Sprintf("/inApps/v1/transactions/%s", transactionID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(400).
			JSON(map[string]any{
				"errorCode":    4000006,
				"errorMessage": "Invalid transaction id.",
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewConnectClient(nil, mockedJWTProvider)
		assert.NoError(t, err)

		appsService := NewInAppsService(c)

		response, clientResponse, err := appsService.GetTransactionInfo(ctx, transactionID)
		responseError := client.ErrorResponse{}
		if assert.ErrorAs(t, err, &responseError) {
			assert.NotNil(t, responseError.Response)
			assert.Equal(t, responseError.Response.StatusCode, http.StatusBadRequest)
		}
		assert.Nil(t, response)
		assert.Nil(t, clientResponse)
	})
}
