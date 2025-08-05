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

func TestListInAppPurchaseManualPrices(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	inAppPurchaseID := "123456"
	territory := "USA"

	gock.New("https://api.appstoreconnect.apple.com").
		Get(fmt.Sprintf("/v1/inAppPurchasePriceSchedules/%s/manualPrices", inAppPurchaseID)).
		MatchParam("include", "inAppPurchasePricePoint,territory").
		MatchParam("fields[inAppPurchasePricePoints]", "customerPrice,territory").
		MatchParam("fields[territories]", "currency").
		MatchParam("filter[territory]", territory).
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(200).
		JSON(map[string]any{
			"data": []any{
				map[string]any{
					"id": "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIiLCJzZCI6MC4wLCJlZCI6MC4wfQ",
					"attributes": map[string]any{
						"manual": true,
					},
					"type": "inAppPurchasePrices",
					"relationships": map[string]any{
						"inAppPurchasePricePoint": map[string]any{
							"data": map[string]string{
								"id":   "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ",
								"type": "inAppPurchasePricePoints",
							},
						},
						"territory": map[string]any{
							"data": map[string]string{
								"id":   territory,
								"type": "territories",
							},
						},
					},
					"links": map[string]string{
						"self": "https://api.appstoreconnect.apple.com/v1/inAppPurchasePrices/eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIiLCJzZCI6MC4wLCJlZCI6MC4wfQ",
					},
				},
			},
			"included": []map[string]any{
				{
					"attributes": map[string]string{
						"customerPrice": "4.99",
						"proceeds":      "3.5",
					},
					"id": "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ",
					"links": map[string]string{
						"self": "https://api.appstoreconnect.apple.com/v1/inAppPurchasePricePoints/eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ",
					},
					"relationships": map[string]any{
						"territory": map[string]any{},
					},
					"type": "inAppPurchasePricePoints",
				},
				{
					"attributes": map[string]string{
						"currency": "USD",
					},
					"id": "USA",
					"links": map[string]string{
						"self": "https://api.appstoreconnect.apple.com/v1/territories/USA",
					},
					"type": "territories",
				},
			},
			"links": map[string]string{
				"self": "https://api.appstoreconnect.apple.com/v1/inAppPurchasePriceSchedules/123456/manualPrices?include=inAppPurchasePricePoint%2Cterritory\u0026filter%5Bterritory%5D=USA\u0026fields%5BinAppPurchasePricePoints%5D=customerPrice%2Cterritory\u0026fields%5Bterritories%5D=currency",
			},
			"meta": map[string]any{
				"paging": map[string]int{
					"limit": 50,
					"total": 1,
				},
			},
		})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := client.NewConnectClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	appsService := NewInAppPurchaseService(c)

	queryParams := ListInAppPurchaseManualPricesQuery{
		Include:                        "inAppPurchasePricePoint,territory",
		FieldsInAppPurchasePricePoints: "customerPrice,territory",
		FieldsTerritories:              "currency",
		FilterTerritory:                territory,
	}
	responsePayload, resp, err := appsService.ListInAppPurchaseManualPrices(ctx, inAppPurchaseID, &queryParams)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, responsePayload)

	assert.Equal(t, 1, len(responsePayload.Data))

	assert.Equal(t, "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIiLCJzZCI6MC4wLCJlZCI6MC4wfQ", responsePayload.Data[0].ID)
	assert.True(t, *responsePayload.Data[0].Attributes.Manual)
	assert.NotNil(t, *responsePayload.Data[0].Relationships.InAppPurchasePricePoint)
	assert.Equal(t, "inAppPurchasePricePoints", responsePayload.Data[0].Relationships.InAppPurchasePricePoint.Data.Type)
	assert.Equal(t, "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ", responsePayload.Data[0].Relationships.InAppPurchasePricePoint.Data.ID)
	assert.Equal(t, "territories", responsePayload.Data[0].Relationships.Territory.Data.Type)
	assert.Equal(t, territory, responsePayload.Data[0].Relationships.Territory.Data.ID)

	assert.Equal(t, 2, len(responsePayload.Included))
	includedTerritories := responsePayload.GetIncludedTerritories()
	assert.Equal(t, 1, len(includedTerritories))
	// assert.Equal(t, "USA", includedTerritories[0].ID)
	assert.Equal(t, "USD", includedTerritories[0].Attributes.Currency)
	includedInAppPurchasePricePoints := responsePayload.GetIncludedInAppPurchasePricePoints()
	assert.Equal(t, 1, len(includedInAppPurchasePricePoints))
	assert.Equal(t, "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ", includedInAppPurchasePricePoints[0].ID)
	assert.Equal(t, "4.99", *includedInAppPurchasePricePoints[0].Attributes.CustomerPrice)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestListInAppPurchaseManualPricesBadRequest(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	inAppPurchaseID := "123456"
	territory := "UNKNOWN"

	gock.New("https://api.appstoreconnect.apple.com").
		Get(fmt.Sprintf("/v1/inAppPurchasePriceSchedules/%s/manualPrices", inAppPurchaseID)).
		MatchParam("include", "inAppPurchasePricePoint,territory").
		MatchParam("fields[inAppPurchasePricePoints]", "customerPrice,territory").
		MatchParam("fields[territories]", "currency").
		MatchParam("filter[territory]", territory).
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(400).
		JSON(map[string]any{
			"errors": []any{
				map[string]any{
					"id":     "123456789",
					"status": "400",
					"code":   "BAD_REQUEST",
					"title":  "Invalid value for filter[territory].",
					"detail": "The filter[territory] contains an invalid parameter value.",
					"meta": map[string]any{
						"associatedErrors": map[string]any{
							"/v1/inAppPurchasePriceSchedules/123456/manualPrices": []any{
								map[string]any{
									"id":     "987654321",
									"status": "400",
									"code":   "BAD_REQUEST",
									"title":  "Bad Request",
									"detail": "Invalid input for field filter[territory]",
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

	queryParams := ListInAppPurchaseManualPricesQuery{
		Include:                        "inAppPurchasePricePoint,territory",
		FieldsInAppPurchasePricePoints: "customerPrice,territory",
		FieldsTerritories:              "currency",
		FilterTerritory:                territory,
	}
	responsePayload, resp, err := appsService.ListInAppPurchaseManualPrices(ctx, inAppPurchaseID, &queryParams)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responsePayload)

	var errorResponse client.ErrorResponse
	assert.True(t, errors.As(err, &errorResponse))
	assert.Equal(t, 400, errorResponse.Response.StatusCode)
	assert.Equal(t, "123456789", *errorResponse.Errors[0].ID)
	assert.Equal(t, "400", errorResponse.Errors[0].Status)
	assert.Equal(t, "BAD_REQUEST", errorResponse.Errors[0].Code)
	assert.Equal(t, "Invalid value for filter[territory].", errorResponse.Errors[0].Title)
	assert.Equal(t, "The filter[territory] contains an invalid parameter value.", errorResponse.Errors[0].Detail)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestListInAppPurchaseAutomaticPrices(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	inAppPurchaseID := "123456"
	territory := "USA"

	gock.New("https://api.appstoreconnect.apple.com").
		Get(fmt.Sprintf("/v1/inAppPurchasePriceSchedules/%s/automaticPrices", inAppPurchaseID)).
		MatchParam("include", "inAppPurchasePricePoint,territory").
		MatchParam("fields[inAppPurchasePricePoints]", "customerPrice,territory").
		MatchParam("fields[territories]", "currency").
		MatchParam("filter[territory]", territory).
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(200).
		JSON(map[string]any{
			"data": []any{
				map[string]any{
					"id": "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIiLCJzZCI6MC4wLCJlZCI6MC4wfQ",
					"attributes": map[string]any{
						"manual": false,
					},
					"type": "inAppPurchasePrices",
					"relationships": map[string]any{
						"inAppPurchasePricePoint": map[string]any{
							"data": map[string]string{
								"id":   "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ",
								"type": "inAppPurchasePricePoints",
							},
						},
						"territory": map[string]any{
							"data": map[string]string{
								"id":   territory,
								"type": "territories",
							},
						},
					},
					"links": map[string]string{
						"self": "https://api.appstoreconnect.apple.com/v1/inAppPurchasePrices/eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIiLCJzZCI6MC4wLCJlZCI6MC4wfQ",
					},
				},
			},
			"included": []map[string]any{
				{
					"attributes": map[string]string{
						"customerPrice": "4.99",
						"proceeds":      "3.5",
					},
					"id": "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ",
					"links": map[string]string{
						"self": "https://api.appstoreconnect.apple.com/v1/inAppPurchasePricePoints/eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ",
					},
					"relationships": map[string]any{
						"territory": map[string]any{},
					},
					"type": "inAppPurchasePricePoints",
				},
				{
					"attributes": map[string]string{
						"currency": "USD",
					},
					"id": "USA",
					"links": map[string]string{
						"self": "https://api.appstoreconnect.apple.com/v1/territories/USA",
					},
					"type": "territories",
				},
			},
			"links": map[string]string{
				"self": "https://api.appstoreconnect.apple.com/v1/inAppPurchasePriceSchedules/123456/automaticPrices?include=inAppPurchasePricePoint%2Cterritory\u0026filter%5Bterritory%5D=USA\u0026fields%5BinAppPurchasePricePoints%5D=customerPrice%2Cterritory\u0026fields%5Bterritories%5D=currency",
			},
			"meta": map[string]any{
				"paging": map[string]int{
					"limit": 50,
					"total": 1,
				},
			},
		})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := client.NewConnectClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	appsService := NewInAppPurchaseService(c)

	queryParams := ListInAppPurchaseAutomaticPricesQuery{
		Include:                        "inAppPurchasePricePoint,territory",
		FieldsInAppPurchasePricePoints: "customerPrice,territory",
		FieldsTerritories:              "currency",
		FilterTerritory:                territory,
	}
	responsePayload, resp, err := appsService.ListInAppPurchaseAutomaticPrices(ctx, inAppPurchaseID, &queryParams)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, responsePayload)

	assert.Equal(t, 1, len(responsePayload.Data))

	assert.Equal(t, "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIiLCJzZCI6MC4wLCJlZCI6MC4wfQ", responsePayload.Data[0].ID)
	assert.False(t, *responsePayload.Data[0].Attributes.Manual)
	assert.NotNil(t, *responsePayload.Data[0].Relationships.InAppPurchasePricePoint)
	assert.Equal(t, "inAppPurchasePricePoints", responsePayload.Data[0].Relationships.InAppPurchasePricePoint.Data.Type)
	assert.Equal(t, "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ", responsePayload.Data[0].Relationships.InAppPurchasePricePoint.Data.ID)
	assert.Equal(t, "territories", responsePayload.Data[0].Relationships.Territory.Data.Type)
	assert.Equal(t, territory, responsePayload.Data[0].Relationships.Territory.Data.ID)

	assert.Equal(t, 2, len(responsePayload.Included))
	includedTerritories := responsePayload.GetIncludedTerritories()
	assert.Equal(t, 1, len(includedTerritories))
	// assert.Equal(t, "USA", includedTerritories[0].ID)
	assert.Equal(t, "USD", includedTerritories[0].Attributes.Currency)
	includedInAppPurchasePricePoints := responsePayload.GetIncludedInAppPurchasePricePoints()
	assert.Equal(t, 1, len(includedInAppPurchasePricePoints))
	assert.Equal(t, "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ", includedInAppPurchasePricePoints[0].ID)
	assert.Equal(t, "4.99", *includedInAppPurchasePricePoints[0].Attributes.CustomerPrice)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestListInAppPurchaseAutomaticPricesBadRequest(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	inAppPurchaseID := "123456"
	territory := "UNKNOWN"

	gock.New("https://api.appstoreconnect.apple.com").
		Get(fmt.Sprintf("/v1/inAppPurchasePriceSchedules/%s/automaticPrices", inAppPurchaseID)).
		MatchParam("include", "inAppPurchasePricePoint,territory").
		MatchParam("fields[inAppPurchasePricePoints]", "customerPrice,territory").
		MatchParam("fields[territories]", "currency").
		MatchParam("filter[territory]", territory).
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(400).
		JSON(map[string]any{
			"errors": []any{
				map[string]any{
					"id":     "123456789",
					"status": "400",
					"code":   "BAD_REQUEST",
					"title":  "Invalid value for filter[territory].",
					"detail": "The filter[territory] contains an invalid parameter value.",
					"meta": map[string]any{
						"associatedErrors": map[string]any{
							"/v1/inAppPurchasePriceSchedules/123456/automaticPrices": []any{
								map[string]any{
									"id":     "987654321",
									"status": "400",
									"code":   "BAD_REQUEST",
									"title":  "Bad Request",
									"detail": "Invalid input for field filter[territory]",
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

	queryParams := ListInAppPurchaseAutomaticPricesQuery{
		Include:                        "inAppPurchasePricePoint,territory",
		FieldsInAppPurchasePricePoints: "customerPrice,territory",
		FieldsTerritories:              "currency",
		FilterTerritory:                territory,
	}
	responsePayload, resp, err := appsService.ListInAppPurchaseAutomaticPrices(ctx, inAppPurchaseID, &queryParams)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responsePayload)

	var errorResponse client.ErrorResponse
	assert.True(t, errors.As(err, &errorResponse))
	assert.Equal(t, 400, errorResponse.Response.StatusCode)
	assert.Equal(t, "123456789", *errorResponse.Errors[0].ID)
	assert.Equal(t, "400", errorResponse.Errors[0].Status)
	assert.Equal(t, "BAD_REQUEST", errorResponse.Errors[0].Code)
	assert.Equal(t, "Invalid value for filter[territory].", errorResponse.Errors[0].Title)
	assert.Equal(t, "The filter[territory] contains an invalid parameter value.", errorResponse.Errors[0].Detail)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}
