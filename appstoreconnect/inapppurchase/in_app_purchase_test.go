package inapppurchase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/h2non/gock"
	"github.com/nbio/st"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListAppInAppPurchases(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	appID := "123456"

	gock.New("https://api.appstoreconnect.apple.com").
		Get(fmt.Sprintf("/v1/apps/%s/inAppPurchasesV2", appID)).
		MatchParam("filter[productId]", "com.example.app.product1,com.example.app.product2,com.example.app.product3").
		MatchParam("filter[state]", "APPROVED").
		MatchParam("limit", "2").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(200).
		JSON(map[string]any{
			"data": []any{
				map[string]any{
					"id": "7890",
					"attributes": map[string]any{
						"productId":         "com.example.app.product1",
						"name":              "My example iap",
						"familySharable":    false,
						"state":             "APPROVED",
						"inAppPurchaseType": "CONSUMABLE",
					},
					"type": "inAppPurchases",
					"relationships": map[string]any{
						"iapPriceSchedule": map[string]any{
							"links": map[string]string{
								"related": "https://api.appstoreconnect.apple.com/v2/inAppPurchases/7890/iapPriceSchedule",
								"self":    "https://api.appstoreconnect.apple.com/v2/inAppPurchases/7890/relationships/iapPriceSchedule",
							},
						},
					},
				},
				map[string]any{
					"id": "78901234",
					"attributes": map[string]any{
						"productId":         "com.example.app.product2",
						"name":              "My example iap 2",
						"familySharable":    false,
						"state":             "APPROVED",
						"inAppPurchaseType": "CONSUMABLE",
					},
					"type": "inAppPurchases",
					"relationships": map[string]any{
						"iapPriceSchedule": map[string]any{
							"links": map[string]string{
								"related": "https://api.appstoreconnect.apple.com/v2/inAppPurchases/78901234/iapPriceSchedule",
								"self":    "https://api.appstoreconnect.apple.com/v2/inAppPurchases/78901234/relationships/iapPriceSchedule",
							},
						},
					},
				},
			},
			"links": map[string]string{
				"next": "https://api.appstoreconnect.apple.com/v1/apps/123456/inAppPurchasesV2?cursor=Ag.ANwgDAk\u0026filter%5BproductId%5D=com.example.app.product1%2Ccom.example.app.product2%2Ccom.example.app.product3\u0026filter%5Bstate%5D=APPROVED\u0026limit=2",
				"self": "https://api.appstoreconnect.apple.com/v1/apps/123456/inAppPurchasesV2?filter%5BproductId%5D=com.example.app.product1%2Ccom.example.app.product2%2Ccom.example.app.product3\u0026filter%5Bstate%5D=APPROVED\u0026limit=2",
			},
			"meta": map[string]any{
				"paging": map[string]int{
					"limit": 2,
					"total": 3,
				},
			},
		})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := client.NewConnectClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	appsService := NewInAppPurchaseService(c)

	queryParams := ListAppInAppPurchasesQuery{
		FilterProductID: "com.example.app.product1,com.example.app.product2,com.example.app.product3",
		FilterState:     "APPROVED",
		Limit:           2,
	}
	responsePayload, resp, err := appsService.ListAppInAppPurchases(ctx, appID, &queryParams)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, responsePayload)

	assert.Equal(t, 2, len(responsePayload.Data))

	assert.Equal(t, "7890", responsePayload.Data[0].ID)
	assert.Equal(t, "com.example.app.product1", *responsePayload.Data[0].Attributes.ProductID)
	assert.Equal(t, "My example iap", *responsePayload.Data[0].Attributes.Name)
	assert.Equal(t, "APPROVED", *responsePayload.Data[0].Attributes.State)
	assert.Equal(t, "CONSUMABLE", *responsePayload.Data[0].Attributes.InAppPurchaseType)
	assert.Equal(t, "inAppPurchases", responsePayload.Data[0].Type)

	assert.Equal(t, "78901234", responsePayload.Data[1].ID)
	assert.Equal(t, "com.example.app.product2", *responsePayload.Data[1].Attributes.ProductID)
	assert.Equal(t, "My example iap 2", *responsePayload.Data[1].Attributes.Name)
	assert.Equal(t, "APPROVED", *responsePayload.Data[1].Attributes.State)
	assert.Equal(t, "CONSUMABLE", *responsePayload.Data[1].Attributes.InAppPurchaseType)
	assert.Equal(t, "inAppPurchases", responsePayload.Data[1].Type)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestListAppInAppPurchasesBadRequest(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	appID := "123456"

	gock.New("https://api.appstoreconnect.apple.com").
		Get(fmt.Sprintf("/v1/apps/%s/inAppPurchasesV2", appID)).
		MatchParam("filter[productId]", "com.example.app.product1,com.example.app.product2,com.example.app.product3").
		MatchParam("filter[state]", "INVALID_STATE").
		MatchParam("limit", "2").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(400).
		JSON(map[string]any{
			"errors": []any{
				map[string]any{
					"id":     "123456789",
					"status": "400",
					"code":   "BAD_REQUEST",
					"title":  "Invalid value for filter[state].",
					"detail": "The filter[state] contains an invalid parameter value.",
					"meta": map[string]any{
						"associatedErrors": map[string]any{
							"/v1/apps/123456/inAppPurchasesV2": []any{
								map[string]any{
									"id":     "987654321",
									"status": "400",
									"code":   "BAD_REQUEST",
									"title":  "Bad Request",
									"detail": "Invalid input for field filter[state]",
								},
							},
						},
					},
				},
			},
		})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := client.NewConnectClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	appsService := NewInAppPurchaseService(c)

	queryParams := ListAppInAppPurchasesQuery{
		FilterProductID: "com.example.app.product1,com.example.app.product2,com.example.app.product3",
		FilterState:     "INVALID_STATE",
		Limit:           2,
	}
	responsePayload, resp, err := appsService.ListAppInAppPurchases(ctx, appID, &queryParams)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responsePayload)

	var errorResponse client.ErrorResponse
	assert.True(t, errors.As(err, &errorResponse))
	assert.Equal(t, 400, errorResponse.Response.StatusCode)
	assert.Equal(t, "123456789", *errorResponse.Errors[0].ID)
	assert.Equal(t, "400", errorResponse.Errors[0].Status)
	assert.Equal(t, "BAD_REQUEST", errorResponse.Errors[0].Code)
	assert.Equal(t, "Invalid value for filter[state].", errorResponse.Errors[0].Title)
	assert.Equal(t, "The filter[state] contains an invalid parameter value.", errorResponse.Errors[0].Detail)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}
