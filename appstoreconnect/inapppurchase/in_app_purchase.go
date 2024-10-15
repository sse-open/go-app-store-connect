package inapppurchase

import (
	"context"
	"fmt"

	"github.com/sse-open/go-app-store-connect/appstoreconnect/common"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/included"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/resource/inapppurchase"
	"github.com/sse-open/go-app-store-connect/client"
)

// A response that contains a list of InAppPurchasesV2 resources for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasesv2response
type InAppPurchasesV2Response struct {
	Data     []inapppurchase.InAppPurchasesV2 `json:"data"`
	Included []included.Included              `json:"included,omitempty"` // Skipping declaration of included type payloads for now
	Links    common.PagedDocumentLinks        `json:"links"`
	Meta     *common.PagingInformation        `json:"meta,omitempty"`
}

// Available query parameters for List All In-App Purchases for an App
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-apps-_id_-inapppurchasesv2
type ListAppInAppPurchasesQuery struct {
	FieldsInAppPurchaseAppStoreReviewScreenshots string `url:"fields[inAppPurchaseAppStoreReviewScreenshots],omitempty"`
	FieldsInAppPurchaseContents                  string `url:"fields[inAppPurchaseContents],omitempty"`
	FieldsInAppPurchaseLocalizations             string `url:"fields[inAppPurchaseLocalizations],omitempty"`
	FieldsInAppPurchases                         string `url:"fields[inAppPurchases],omitempty"`
	FieldsPromotedPurchases                      string `url:"fields[promotedPurchases],omitempty"`
	FilterInAppPurchaseType                      string `url:"filter[inAppPurchaseType],omitempty"`
	FilterName                                   string `url:"filter[name],omitempty"`
	FilterProductID                              string `url:"filter[productId],omitempty"`
	FilterState                                  string `url:"filter[state],omitempty"`
	Include                                      string `url:"include,omitempty"`
	Limit                                        int    `url:"limit,omitempty"`
	LimitInAppPurchaseLocalizations              int    `url:"limit[inAppPurchaseLocalizations],omitempty"`
	Sort                                         string `url:"sort,omitempty"`
	FieldsInAppPurchasePriceSchedules            string `url:"fields[inAppPurchasePriceSchedules],omitempty"`
	FieldsInAppPurchaseAvailabilities            string `url:"fields[inAppPurchaseAvailabilities],omitempty"`
	FieldsInAppPurchaseImages                    string `url:"fields[inAppPurchaseImages],omitempty"`
	LimitImages                                  int    `url:"limit[images],omitempty"`
	Cursor                                       string `url:"cursor,omitempty"`
}

func (iaps *InAppPurchaseService) ListAppInAppPurchases(ctx context.Context, appID string, queryParams *ListAppInAppPurchasesQuery) (*InAppPurchasesV2Response, *client.ClientResponse, error) {
	url := fmt.Sprintf("v1/apps/%s/inAppPurchasesV2", appID)
	respPayload := &InAppPurchasesV2Response{}
	resp, err := iaps.client.Get(ctx, url, queryParams, respPayload)
	if err != nil {
		return nil, nil, err
	}

	return respPayload, resp, nil
}
