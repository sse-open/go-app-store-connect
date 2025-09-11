package subscription

import "github.com/sse-open/go-app-store-connect/appstoreserver/resource/common"

// The status of an auto-renewable subscription.
//
// https://developer.apple.com/documentation/appstoreserverapi/status
type AutoRenewableStatus int

var (
	AutoRenewableStatusActive             AutoRenewableStatus = 1
	AutoRenewableStatusExpired            AutoRenewableStatus = 2
	AutoRenewableStatusBillingRetryPeriod AutoRenewableStatus = 3
	AutoRenewableStatusBillingGracePeriod AutoRenewableStatus = 4
	AutoRenewableStatusRevoked            AutoRenewableStatus = 5
)

type LastTransactionsItem struct {
	OriginalTransactionId string                `json:"originalTransactionId"`
	SignedTransactionInfo common.JWSTransaction `json:"signedTransactionInfo"`
	SignedRenewalInfo     common.JWSRenewalInfo `json:"signedRenewalInfo"`
	Status                AutoRenewableStatus   `json:"status"`
}

// Information for auto-renewable subscriptions, including signed transaction information and signed renewal information,
// for one subscription group.
//
// https://developer.apple.com/documentation/appstoreserverapi/subscriptiongroupidentifieritem
type SubscriptionGroupIdentifierItem struct {
	SubscriptionGroupIdentifier string                 `json:"subscriptionGroupIdentifier"`
	LastTransactions            []LastTransactionsItem `json:"lastTransactions"`
}

// A response that contains status information for all of a customer’s auto-renewable subscriptions in your app.
//
// https://developer.apple.com/documentation/appstoreserverapi/statusresponse
type StatusResponse struct {
	Data        []SubscriptionGroupIdentifierItem `json:"data"`
	Environment common.Environment                `json:"environment"`
	AppAppleId  int64                             `json:"appAppleId"`
	BundleId    string                            `json:"bundleId"`
}
