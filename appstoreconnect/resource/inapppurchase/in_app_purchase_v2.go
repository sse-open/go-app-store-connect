package inapppurchase

import "github.com/sse-open/go-app-store-connect/appstoreconnect/common"

// Attributes that describe an InAppPurchaseV2 resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasev2/attributes-data.dictionary
type InAppPurchasesV2Attributes struct {
	ContentHosting    *bool   `json:"contentHosting,omitempty"`
	FamilySharable    *bool   `json:"familySharable,omitempty"`
	InAppPurchaseType *string `json:"inAppPurchaseType,omitempty"`
	Name              *string `json:"name,omitempty"`
	ProductID         *string `json:"productId,omitempty"`
	ReviewNote        *string `json:"reviewNote,omitempty"`
	State             *string `json:"state,omitempty"`
}

// The relationships you included in the request and those on which you can operate.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasev2/relationships-data.dictionary
type InAppPurchasesV2Relationships struct {
	AppStoreReviewScreenshot   *common.Relationship      `json:"appStoreReviewScreenshot,omitempty"`
	Content                    *common.Relationship      `json:"content,omitempty"`
	IapPriceSchedule           *common.Relationship      `json:"iapPriceSchedule,omitempty"`
	Images                     *common.PagedRelationship `json:"images,omitempty"`
	InAppPurchaseAvailability  *common.Relationship      `json:"inAppPurchaseAvailability,omitempty"`
	InAppPurchaseLocalizations *common.PagedRelationship `json:"inAppPurchaseLocalizations,omitempty"`
	PricePoints                *common.PagedRelationship `json:"pricePoints,omitempty"`
	PromotedPurchase           *common.Relationship      `json:"promotedPurchase,omitempty"`
}

// The data structure that represents an InAppPurchaseV2 resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasev2
type InAppPurchasesV2 struct {
	Attributes    *InAppPurchasesV2Attributes    `json:"attributes,omitempty"`
	ID            string                         `json:"id"`
	Links         common.ResourceLinks           `json:"links"`
	Relationships *InAppPurchasesV2Relationships `json:"relationships,omitempty"`
	Type          string                         `json:"type"`
}
