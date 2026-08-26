package subscription

import (
	"github.com/sse-open/go-app-store-connect/common"
)

// The code that represents the reason for the subscription-renewal-date extension.
//
// https://developer.apple.com/documentation/appstoreserverapi/extendreasoncode
type ExtendReasonCode int

var (
	ExtendReasonCodeUndeclared           ExtendReasonCode = 0
	ExtendReasonCodeCustomerSatisfaction ExtendReasonCode = 1
	ExtendReasonCodeOther                ExtendReasonCode = 2
	ExtendReasonCodeServiceIssue         ExtendReasonCode = 3
)

// The request body that contains subscription-renewal-extension data for an individual subscription.
//
// https://developer.apple.com/documentation/appstoreserverapi/extendrenewaldaterequest
type ExtendRenewalDateRequest struct {
	ExtendByDays      int              `json:"extendByDays"`
	ExtendReasonCode  ExtendReasonCode `json:"extendReasonCode"`
	RequestIdentifier string           `json:"requestIdentifier"`
}

// A response that indicates whether an individual renewal-date extension succeeded, and related details.
//
// https://developer.apple.com/documentation/appstoreserverapi/extendrenewaldateresponse
type ExtendRenewalDateResponse struct {
	EffectiveDate         *common.Timestamp `json:"effectiveDate,omitempty"`
	OriginalTransactionId string            `json:"originalTransactionId"`
	Success               bool              `json:"success"`
	WebOrderLineItemId    string            `json:"webOrderLineItemId"`
}

// The request body that contains subscription-renewal-extension data to apply for all eligible active subscribers.
//
// https://developer.apple.com/documentation/appstoreserverapi/massextendrenewaldaterequest
type MassExtendRenewalDateRequest struct {
	RequestIdentifier      string           `json:"requestIdentifier"`
	ExtendByDays           int              `json:"extendByDays"`
	ExtendReasonCode       ExtendReasonCode `json:"extendReasonCode"`
	ProductId              string           `json:"productId"`
	StorefrontCountryCodes []string         `json:"storefrontCountryCodes,omitempty"`
}

// A response that indicates the server successfully received the subscription-renewal-date extension request.
//
// https://developer.apple.com/documentation/appstoreserverapi/massextendrenewaldateresponse
type MassExtendRenewalDateResponse struct {
	RequestIdentifier string `json:"requestIdentifier"`
}
