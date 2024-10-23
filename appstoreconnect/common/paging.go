package common

// Links related to the response document, including paging links.
//
// https://developer.apple.com/documentation/appstoreconnectapi/pageddocumentlinks
type PagedDocumentLinks struct {
	First *URI `json:"first,omitempty"`
	Next  *URI `json:"next,omitempty"`
	Self  URI  `json:"self"`
}
