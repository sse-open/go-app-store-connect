package inapppurchase

import "github.com/sse-open/go-app-store-connect/appstoreconnect/common"

// Attributes that describe a InAppPurchasePricePoint resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasepricepoint/attributes-data.dictionary
type InAppPurchasePricePointAttributes struct {
	CustomerPrice *string `json:"customerPrice,omitempty"`
	Proceeds      *string `json:"proceeds,omitempty"`
}

// The relationships you included in the request and those on which you can operate.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasepricepoint/relationships-data.dictionary
type InAppPurchasePricePointRelationships struct {
	Territory *common.RelationshipDataOnly `json:"territory,omitempty"`
}

// The data structure that represents an InAppPurchasePricePoint resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasepricepoint
type InAppPurchasePricePoint struct {
	Attributes    *InAppPurchasePricePointAttributes    `json:"attributes,omitempty"`
	ID            string                                `json:"id"`
	Links         common.ResourceLinks                  `json:"links,omitempty"`
	Relationships *InAppPurchasePricePointRelationships `json:"relationships,omitempty"`
	Type          string                                `json:"type"`
}
