package subscription

import (
	"context"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	resourcesubscription "github.com/sse-open/go-app-store-connect/appstoreserver/resource/subscription"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMassExtendSubscriptionRenewalDates(t *testing.T) {
	t.Run("mass extend success", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()
		requestIdentifier := "fdf6a8c8-2bef-4a8c-9e0d-3c0a8b1d6f33"
		productId := "com.example.subscription.monthly"

		gock.New("https://api.storekit-sandbox.itunes.apple.com").
			Post("/inApps/v1/subscriptions/extend/mass").
			MatchHeader("Authorization", "Bearer fakeToken").
			JSON(map[string]any{
				"requestIdentifier":      requestIdentifier,
				"extendByDays":           7,
				"extendReasonCode":       3,
				"productId":              productId,
				"storefrontCountryCodes": []string{"USA", "GBR"},
			}).
			Reply(200).
			JSON(map[string]any{
				"requestIdentifier": requestIdentifier,
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewServerClient(nil, mockedJWTProvider, true)
		assert.NoError(t, err)

		subscriptionService := NewSubscriptionService(c)

		response, clientResponse, err := subscriptionService.MassExtendSubscriptionRenewalDates(ctx, &resourcesubscription.MassExtendRenewalDateRequest{
			RequestIdentifier:      requestIdentifier,
			ExtendByDays:           7,
			ExtendReasonCode:       resourcesubscription.ExtendReasonCodeServiceIssue,
			ProductId:              productId,
			StorefrontCountryCodes: []string{"USA", "GBR"},
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, clientResponse)

		assert.Equal(t, requestIdentifier, response.RequestIdentifier)
	})

	t.Run("invalid product id", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()
		requestIdentifier := "fdf6a8c8-2bef-4a8c-9e0d-3c0a8b1d6f33"

		gock.New("https://api.storekit-sandbox.itunes.apple.com").
			Post("/inApps/v1/subscriptions/extend/mass").
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(400).
			JSON(map[string]any{
				"errors": []map[string]any{
					{
						"status": "400",
						"code":   "4000023",
						"title":  "Invalid product id",
					},
				},
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewServerClient(nil, mockedJWTProvider, true)
		assert.NoError(t, err)

		subscriptionService := NewSubscriptionService(c)

		response, clientResponse, err := subscriptionService.MassExtendSubscriptionRenewalDates(ctx, &resourcesubscription.MassExtendRenewalDateRequest{
			RequestIdentifier: requestIdentifier,
			ExtendByDays:      7,
			ExtendReasonCode:  resourcesubscription.ExtendReasonCodeServiceIssue,
			ProductId:         "invalid.product.id",
		})
		responseError := client.ErrorResponse{}
		if assert.ErrorAs(t, err, &responseError) {
			assert.NotNil(t, responseError.Response)
			assert.Equal(t, http.StatusBadRequest, responseError.Response.StatusCode)
		}
		assert.Nil(t, response)
		assert.Nil(t, clientResponse)
	})
}
