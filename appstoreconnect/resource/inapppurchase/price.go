package inapppurchase

import (
	"time"

	"github.com/sse-open/go-app-store-connect/appstoreconnect/common"
)

// Attributes that describe an InAppPurchasePrice resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchaseprice/attributes-data.dictionary
type InAppPurchasePriceAttributes struct {
	EndDate   *time.Time `json:"endDate,omitempty"`
	Manual    *bool      `json:"manual,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
}

// The relationships you included in the request and those on which you can operate.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasev2/relationships-data.dictionary
type InAppPurchasePriceRelationships struct {
	InAppPurchasePricePoint *common.RelationshipDataOnly `json:"inAppPurchasePricePoint,omitempty"`
	Territory               *common.RelationshipDataOnly `json:"territory,omitempty"`
}

// The data structure that represents an InAppPurchaseV2 resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchaseprice
type InAppPurchasePrice struct {
	Attributes    *InAppPurchasePriceAttributes    `json:"attributes,omitempty"`
	ID            string                           `json:"id"`
	Links         common.ResourceLinks             `json:"links"`
	Relationships *InAppPurchasePriceRelationships `json:"relationships,omitempty"`
	Type          string                           `json:"type"`
}
