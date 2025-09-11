package inapps

import "github.com/sse-open/go-app-store-connect/appstoreserver/resource/common"

// A response that contains signed transaction information for a single transaction.
//
// https://developer.apple.com/documentation/appstoreserverapi/transactioninforesponse
type TransactionInfoResponse struct {
	SignedTransactionInfo common.JWSTransaction `json:"signedTransactionInfo"`
}
