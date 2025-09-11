package inapps

import "github.com/sse-open/go-app-store-connect/appstoreserver/resource/common"

// The status that indicates whether the order ID is valid.
//
// https://developer.apple.com/documentation/appstoreserverapi/orderlookupstatus
type OrderLookupStatus int

func (s OrderLookupStatus) IsSuccess() bool {
	return s == OrderLookupStatusSuccess
}

var (
	OrderLookupStatusSuccess OrderLookupStatus = 0
	OrderLookupStatusFailure OrderLookupStatus = 1
)

// Contains information about an in-app purchase order.
//
// https://developer.apple.com/documentation/appstoreserverapi/orderlookupresponse
type OrderLookupResponse struct {
	Status             OrderLookupStatus       `json:"status"`
	SignedTransactions []common.JWSTransaction `json:"signedTransactions"`
}
