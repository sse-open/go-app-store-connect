package subscription

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/common"
	"github.com/sse-open/go-app-store-connect/appstoreserver/resource/subscription"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetSubscriptionStatus(t *testing.T) {
	t.Run("get success", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		transactionID := "transaction123"

		gock.New("https://api.appstoreconnect.apple.com").
			Get(fmt.Sprintf("/inApps/v1/subscriptions/%s", transactionID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(200).
			JSON(map[string]any{
				"data": []map[string]any{
					{
						"subscriptionGroupIdentifier": "group123",
						"lastTransactions": []map[string]any{
							{
								"originalTransactionId": "originalTransaction123",
								"signedTransactionInfo": "xxx.yyy.zzz",
								"signedRenewalInfo":     "aaa.bbb.ccc",
								"status":                1,
							},
						},
					},
				},
				"environment": "Sandbox",
				"appAppleId":  123456789,
				"bundleId":    "com.example.app",
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewConnectClient(nil, mockedJWTProvider)
		assert.NoError(t, err)

		subscriptionService := NewSubscriptionService(c)

		response, clientResponse, err := subscriptionService.GetSubscriptionStatus(ctx, transactionID, nil)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, clientResponse)

		assert.Equal(t, common.EnvironmentSandbox, response.Environment)
		assert.Equal(t, int64(123456789), response.AppAppleId)
		assert.Equal(t, "com.example.app", response.BundleId)
		assert.Len(t, response.Data, 1)

		group := response.Data[0]
		assert.Equal(t, "group123", group.SubscriptionGroupIdentifier)
		assert.Len(t, group.LastTransactions, 1)

		lastTransaction := group.LastTransactions[0]
		assert.Equal(t, "originalTransaction123", lastTransaction.OriginalTransactionId)
		assert.Equal(t, subscription.AutoRenewableStatusActive, lastTransaction.Status)
		assert.Equal(t, common.JWSTransaction("xxx.yyy.zzz"), lastTransaction.SignedTransactionInfo)
		assert.Equal(t, common.JWSRenewalInfo("aaa.bbb.ccc"), lastTransaction.SignedRenewalInfo)
	})

	t.Run("get success with status param", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		transactionID := "transaction123"

		gock.New("https://api.appstoreconnect.apple.com").
			Get(fmt.Sprintf("/inApps/v1/subscriptions/%s", transactionID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			MatchParam("status", "1").
			Reply(200).
			JSON(map[string]any{
				"data": []map[string]any{
					{
						"subscriptionGroupIdentifier": "group123",
						"lastTransactions": []map[string]any{
							{
								"originalTransactionId": "originalTransaction123",
								"signedTransactionInfo": "xxx.yyy.zzz",
								"signedRenewalInfo":     "aaa.bbb.ccc",
								"status":                1,
							},
						},
					},
				},
				"environment": "Sandbox",
				"appAppleId":  123456789,
				"bundleId":    "com.example.app",
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewConnectClient(nil, mockedJWTProvider)
		assert.NoError(t, err)

		subscriptionService := NewSubscriptionService(c)

		response, clientResponse, err := subscriptionService.GetSubscriptionStatus(ctx, transactionID, &GetSubscriptionStatusQuery{
			Status: []subscription.AutoRenewableStatus{
				subscription.AutoRenewableStatusActive,
			},
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, clientResponse)

		assert.Equal(t, common.EnvironmentSandbox, response.Environment)
		assert.Equal(t, int64(123456789), response.AppAppleId)
		assert.Equal(t, "com.example.app", response.BundleId)
		assert.Len(t, response.Data, 1)

		group := response.Data[0]
		assert.Equal(t, "group123", group.SubscriptionGroupIdentifier)
		assert.Len(t, group.LastTransactions, 1)

		lastTransaction := group.LastTransactions[0]
		assert.Equal(t, "originalTransaction123", lastTransaction.OriginalTransactionId)
		assert.Equal(t, subscription.AutoRenewableStatusActive, lastTransaction.Status)
		assert.Equal(t, common.JWSTransaction("xxx.yyy.zzz"), lastTransaction.SignedTransactionInfo)
		assert.Equal(t, common.JWSRenewalInfo("aaa.bbb.ccc"), lastTransaction.SignedRenewalInfo)
	})

	t.Run("invalid transaction id", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()

		transactionID := "invalidTransactionID"

		gock.New("https://api.appstoreconnect.apple.com").
			Get(fmt.Sprintf("/inApps/v1/subscriptions/%s", transactionID)).
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

		subscriptionService := NewSubscriptionService(c)

		response, clientResponse, err := subscriptionService.GetSubscriptionStatus(ctx, transactionID, nil)
		responseError := client.ErrorResponse{}
		if assert.ErrorAs(t, err, &responseError) {
			assert.NotNil(t, responseError.Response)
			assert.Equal(t, responseError.Response.StatusCode, http.StatusBadRequest)
		}
		assert.Nil(t, response)
		assert.Nil(t, clientResponse)
	})
}
