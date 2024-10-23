package territories

import "github.com/sse-open/go-app-store-connect/appstoreconnect/common"

// Attributes that describe a Territories resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/territory/attributes-data.dictionary
type TerritoryAttributes struct {
	Currency string `json:"currency"`
}

// The data structure that represents a Territories resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/territory
type Territory struct {
	Attributes *TerritoryAttributes `json:"attributes,omitempty"`
	ID         string               `json:"id"`
	Links      common.ResourceLinks `json:"links,omitempty"`
	Type       string               `json:"type"`
}
