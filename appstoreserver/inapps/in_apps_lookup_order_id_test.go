package inapps

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/common"
	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/inapps"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLookupOrderId(t *testing.T) {
	t.Run("lookup success", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		orderID := "order123"

		gock.New("https://api.appstoreconnect.apple.com").
			Get(fmt.Sprintf("/inApps/v1/lookup/%s", orderID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(200).
			JSON(map[string]any{
				"status": 0,
				"signedTransactions": []string{
					"xxx.yyy.zzz",
				},
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewConnectClient(nil, mockedJWTProvider)
		assert.NoError(t, err)

		appsService := NewInAppsService(c)

		response, clientResponse, err := appsService.LookUpOrderID(ctx, orderID)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, clientResponse)

		assert.Equal(t, inapps.OrderLookupStatusSuccess, response.Status)
		if assert.Len(t, response.SignedTransactions, 1) {
			assert.Equal(t, common.JWSTransaction("xxx.yyy.zzz"), response.SignedTransactions[0])
		}
	})

	t.Run("lookup failure", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		orderID := "orderNotFound"

		gock.New("https://api.appstoreconnect.apple.com").
			Get(fmt.Sprintf("/inApps/v1/lookup/%s", orderID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(200).
			JSON(map[string]any{
				"status": 1,
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewConnectClient(nil, mockedJWTProvider)
		assert.NoError(t, err)

		appsService := NewInAppsService(c)

		response, clientResponse, err := appsService.LookUpOrderID(ctx, orderID)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, clientResponse)
		assert.Equal(t, inapps.OrderLookupStatusFailure, response.Status)
	})

	t.Run("wrong app id", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		orderID := "orderWrongApp"

		gock.New("https://api.appstoreconnect.apple.com").
			Get(fmt.Sprintf("/inApps/v1/lookup/%s", orderID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(401).
			Body(strings.NewReader("Unauthenticated\n\nRequest ID: EOZKUGINZDWMNLY4EDTKN6RWWI.0.0"))

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewConnectClient(nil, mockedJWTProvider)
		assert.NoError(t, err)

		appsService := NewInAppsService(c)

		response, clientResponse, err := appsService.LookUpOrderID(ctx, orderID)
		responseError := client.ErrorResponse{}
		if assert.ErrorAs(t, err, &responseError) {
			assert.NotNil(t, responseError.Response)
			assert.Equal(t, []client.ErrorResponseError{
				{
					Status: "401 Unauthorized",
				},
			}, responseError.Errors)
		}
		assert.Nil(t, response)
		assert.Nil(t, clientResponse)
	})
}
