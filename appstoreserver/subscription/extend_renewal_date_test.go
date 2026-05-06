package subscription

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/h2non/gock"
	appstoreservercommon "github.com/sse-open/go-app-store-connect/appstoreserver/common"
	resourcesubscription "github.com/sse-open/go-app-store-connect/appstoreserver/resource/subscription"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestExtendSubscriptionRenewalDate(t *testing.T) {
	t.Run("extend success", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()
		originalTransactionID := "1000000123456789"
		effectiveDate := time.Date(2026, 5, 6, 12, 0, 0, 0, time.UTC)

		gock.New("https://api.storekit-sandbox.itunes.apple.com").
			Put(fmt.Sprintf("/inApps/v1/subscriptions/extend/%s", originalTransactionID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			JSON(map[string]any{
				"extendByDays":      7,
				"extendReasonCode":  3,
				"requestIdentifier": "request-123",
			}).
			Reply(200).
			JSON(map[string]any{
				"effectiveDate":         effectiveDate.UnixMilli(),
				"originalTransactionId": originalTransactionID,
				"success":               true,
				"webOrderLineItemId":    "230000123456789",
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewServerClient(nil, mockedJWTProvider, true)
		assert.NoError(t, err)

		subscriptionService := NewSubscriptionService(c)

		response, clientResponse, err := subscriptionService.ExtendSubscriptionRenewalDate(ctx, originalTransactionID, &resourcesubscription.ExtendRenewalDateRequest{
			ExtendByDays:      7,
			ExtendReasonCode:  resourcesubscription.ExtendReasonCodeServiceIssue,
			RequestIdentifier: "request-123",
		})
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, clientResponse)

		assert.Equal(t, originalTransactionID, response.OriginalTransactionId)
		assert.True(t, response.Success)
		assert.Equal(t, "230000123456789", response.WebOrderLineItemId)
		if assert.NotNil(t, response.EffectiveDate) {
			assert.Equal(t, appstoreservercommon.Timestamp{Time: effectiveDate}, *response.EffectiveDate)
		}
	})

	t.Run("invalid original transaction id", func(t *testing.T) {
		defer gock.Off() // Flush pending mocks after test execution

		ctx := context.Background()
		originalTransactionID := "invalidTransactionID"

		gock.New("https://api.storekit-sandbox.itunes.apple.com").
			Put(fmt.Sprintf("/inApps/v1/subscriptions/extend/%s", originalTransactionID)).
			MatchHeader("Authorization", "Bearer fakeToken").
			Reply(400).
			JSON(map[string]any{
				"errors": []map[string]any{
					{
						"status": "400",
						"code":   "4000008",
						"title":  "Invalid original transaction id",
					},
				},
			})

		mockedJWTProvider := mocks.NewIJWTProvider(t)
		mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

		c, err := client.NewServerClient(nil, mockedJWTProvider, true)
		assert.NoError(t, err)

		subscriptionService := NewSubscriptionService(c)

		response, clientResponse, err := subscriptionService.ExtendSubscriptionRenewalDate(ctx, originalTransactionID, &resourcesubscription.ExtendRenewalDateRequest{
			ExtendByDays:      7,
			ExtendReasonCode:  resourcesubscription.ExtendReasonCodeServiceIssue,
			RequestIdentifier: "request-123",
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
