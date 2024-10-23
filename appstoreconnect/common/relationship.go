package common

// Paging details such as the total number of resources and the per-page limit.
//
// https://developer.apple.com/documentation/appstoreconnectapi/paginginformation/paging-data.dictionary
type Paging struct {
	Limit int `json:"limit"`
	Total int `json:"total"`
}

// Paging information for data responses.
//
// https://developer.apple.com/documentation/appstoreconnectapi/paginginformation
type PagingInformation struct {
	Paging Paging `json:"paging"`
}

// Links related to the response document, including self links.
//
// https://developer.apple.com/documentation/appstoreconnectapi/relationshiplinks
type RelationshipLinks struct {
	Related *URI `json:"related,omitempty"`
	Self    *URI `json:"self,omitempty"`
}

// Common type to describe a relationship between resources including only links.
type RelationshipLinksOnly struct {
	Links *RelationshipLinks `json:"links,omitempty"`
}

// Common type to describe a relationship between resources including only data.
type RelationshipDataOnly struct {
	Data *RelationshipData `json:"data,omitempty"`
}

// Common type to describe a relationship between resources including data and links.
type Relationship struct {
	Data  *RelationshipData  `json:"data,omitempty"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

// Common type to describe a relationship between resources which includes paging information.
type PagedRelationship struct {
	Data  []RelationshipData `json:"data,omitempty"`
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// The type and ID of a related resource.
type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Self-links to requested resources.
//
// https://developer.apple.com/documentation/appstoreconnectapi/resourcelinks
type ResourceLinks struct {
	Self URI `json:"self"`
}
