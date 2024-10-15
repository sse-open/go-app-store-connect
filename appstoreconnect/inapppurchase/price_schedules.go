package inapppurchase

import (
	"context"
	"fmt"

	"github.com/sse-open/go-app-store-connect/appstoreconnect/common"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/included"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/resource/inapppurchase"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/resource/territories"
	"github.com/sse-open/go-app-store-connect/client"
)

// A response that contains a list of InAppPurchasePricesResponse resource for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasepricesresponse
type InAppPurchasePricesResponse struct {
	Data     []inapppurchase.InAppPurchasePrice `json:"data"`
	Included []included.Included                `json:"included,omitempty"` // Skipping declaration of included type payloads for now
	Links    common.PagedDocumentLinks          `json:"links"`
	Meta     *common.PagingInformation          `json:"meta,omitempty"`
}

// Get information about a set price or prices for an in-app purchase price schedule.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-inapppurchasepriceschedules-_id_-manualprices
type ListInAppPurchaseManualPricesQuery struct {
	FieldsInAppPurchasePricePoints string `url:"fields[inAppPurchasePricePoints],omitempty"`
	FieldsInAppPurchasePrices      string `url:"inAppPurchasePrices,omitempty"`
	Include                        string `url:"include,omitempty"`
	Limit                          int    `url:"limit,omitempty"`
	FilterTerritory                string `url:"filter[territory],omitempty"`
	FieldsTerritories              string `url:"fields[territories],omitempty"`
}

// Get information about a set price or prices for an in-app purchase price schedule.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-inapppurchasepriceschedules-_id_-manualprices
func (iaps *InAppPurchaseService) ListInAppPurchaseManualPrices(ctx context.Context, inAppPurchaseID string, queryParams *ListInAppPurchaseManualPricesQuery) (*InAppPurchasePricesResponse, *client.ClientResponse, error) {
	url := fmt.Sprintf("v1/inAppPurchasePriceSchedules/%s/manualPrices", inAppPurchaseID)
	respPayload := &InAppPurchasePricesResponse{}
	resp, err := iaps.client.Get(ctx, url, queryParams, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}

// Get information about a price or prices automatically set based on a base territory for an in-app purchase price schedule.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-inapppurchasepriceschedules-_id_-automaticprices
type ListInAppPurchaseAutomaticPricesQuery struct {
	FieldsInAppPurchasePricePoints string `url:"fields[inAppPurchasePricePoints],omitempty"`
	FieldsInAppPurchasePrices      string `url:"inAppPurchasePrices,omitempty"`
	Include                        string `url:"include,omitempty"`
	Limit                          int    `url:"limit,omitempty"`
	FilterTerritory                string `url:"filter[territory],omitempty"`
	FieldsTerritories              string `url:"fields[territories],omitempty"`
}

// Get information about a price or prices automatically set based on a base territory for an in-app purchase price schedule.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-inapppurchasepriceschedules-_id_-automaticprices
func (iaps *InAppPurchaseService) ListInAppPurchaseAutomaticPrices(ctx context.Context, inAppPurchaseID string, queryParams *ListInAppPurchaseAutomaticPricesQuery) (*InAppPurchasePricesResponse, *client.ClientResponse, error) {
	url := fmt.Sprintf("v1/inAppPurchasePriceSchedules/%s/automaticPrices", inAppPurchaseID)
	respPayload := &InAppPurchasePricesResponse{}
	resp, err := iaps.client.Get(ctx, url, queryParams, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}

func (i *InAppPurchasePricesResponse) GetIncludedTerritories() []territories.Territory {
	var incTerritories []territories.Territory
	for _, i := range i.Included {
		if v, ok := i.TypeData.(territories.Territory); ok {
			incTerritories = append(incTerritories, v)
		}
	}

	return incTerritories
}

func (i *InAppPurchasePricesResponse) GetIncludedInAppPurchasePricePoints() []inapppurchase.InAppPurchasePricePoint {
	var incInAppPurchasePricePoints []inapppurchase.InAppPurchasePricePoint
	for _, i := range i.Included {
		if v, ok := i.TypeData.(inapppurchase.InAppPurchasePricePoint); ok {
			incInAppPurchasePricePoints = append(incInAppPurchasePricePoints, v)
		}
	}

	return incInAppPurchasePricePoints
}
