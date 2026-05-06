package inapps

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFinishTransaction(t *testing.T) {
	t.Run("finish success", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		transactionID := "transaction123"

		gock.New("https://api.storekit-sandbox.itunes.apple.com").
			Post(fmt.Sprintf("/inApps/v1/transactions/%s/finish", transactionID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(200)

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewServerClient(nil, mockedJWTProvider, true)
		assert.NoError(t, err)

		appsService := NewInAppsService(c)

		clientResponse, err := appsService.FinishTransaction(ctx, transactionID)
		assert.NoError(t, err)
		assert.NotNil(t, clientResponse)
	})

	t.Run("transaction id not found", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		transactionID := "missingTransactionID"

		gock.New("https://api.storekit-sandbox.itunes.apple.com").
			Post(fmt.Sprintf("/inApps/v1/transactions/%s/finish", transactionID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(404).
			JSON(map[string]any{
				"errorCode":    4040010,
				"errorMessage": "Transaction id not found.",
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewServerClient(nil, mockedJWTProvider, true)
		assert.NoError(t, err)

		appsService := NewInAppsService(c)

		clientResponse, err := appsService.FinishTransaction(ctx, transactionID)
		responseError := client.ErrorResponse{}
		if assert.ErrorAs(t, err, &responseError) {
			assert.NotNil(t, responseError.Response)
			assert.Equal(t, http.StatusNotFound, responseError.Response.StatusCode)
		}
		assert.Nil(t, clientResponse)
	})
}
