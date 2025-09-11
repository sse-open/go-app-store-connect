package common

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/sse-open/go-app-store-connect/appstoreserver/common"
)

// A string that describes whether the transaction was purchased by the customer, or is available to them through Family Sharing.
//
// https://developer.apple.com/documentation/appstoreserverapi/inappownershiptype
type InAppOwnershipType string

var (
	InAppOwnershipTypeFamilyShared InAppOwnershipType = "FAMILY_SHARED"
	InAppOwnershipTypePurchased    InAppOwnershipType = "PURCHASED"
)

// The payment mode for subscription offers on an auto-renewable subscription.
//
// https://developer.apple.com/documentation/appstoreserverapi/offerdiscounttype
type OfferDiscountType string

var (
	OfferDiscountTypeFreeTrial  OfferDiscountType = "FREE_TRIAL"
	OfferDiscountTypePayAsYouGo OfferDiscountType = "PAY_AS_YOU_GO"
	OfferDiscountTypePayUpFront OfferDiscountType = "PAY_UP_FRONT"
)

// The duration of the offer.
//
// https://developer.apple.com/documentation/appstoreserverapi/offerperiod
type OfferPeriod string

var (
	OfferPeriodOneMonth  OfferPeriod = "P1M"
	OfferPeriodTwoMonth  OfferPeriod = "P2M"
	OfferPeriodThreeDays OfferPeriod = "P3D"
)

// The type of subscription offer.
//
// https://developer.apple.com/documentation/appstoreserverapi/offertype
type OfferType int

var (
	OfferTypeIntroductory OfferType = 1
	OfferTypePromotional  OfferType = 2
	OfferTypeOfferCode    OfferType = 3
	OfferTypeWinBack      OfferType = 4
)

// The reason for a refunded transaction.
//
// https://developer.apple.com/documentation/appstoreserverapi/revocationreason
type RevocationReason int

var (
	RevocationReasonOtherReason          RevocationReason = 0
	RevocationReasonActualPerceivedIssue RevocationReason = 1
)

// The cause of a purchase transaction, which indicates whether it’s a customer’s
// purchase or a renewal for an auto-renewable subscription that the system initiates.
//
// https://developer.apple.com/documentation/appstoreserverapi/transactionreason
type TransactionReason string

var (
	TransactionReasonPurchase TransactionReason = "PURCHASE"
	TransactionReasonRenewal  TransactionReason = "RENEWAL"
)

// The type of In-App Purchase products you can offer in your app.
//
// https://developer.apple.com/documentation/appstoreserverapi/type
type TransactionType string

var (
	TransactionTypeAutoRenewableSubscription TransactionType = "Auto-Renewable Subscription"
	TransactionTypeNonConsumable             TransactionType = "Non-Consumable"
	TransactionTypeConsumable                TransactionType = "Consumable"
	TransactionTypeNonRenewingSubscription   TransactionType = "Non-Renewing Subscription"
)

// The price, in milliunits, of the In-App Purchase that the system records in the transaction.
//
// https://developer.apple.com/documentation/appstoreserverapi/price
type Price uint64

func (s Price) ToDecimal() decimal.Decimal {
	return decimal.NewFromUint64(uint64(s)).Div(decimal.NewFromInt(1000))
}

// The renewal status for an auto-renewable subscription.
//
// https://developer.apple.com/documentation/appstoreserverapi/autorenewstatus
type AutoRenewStatus int

var (
	AutoRenewStatusOff AutoRenewStatus = 0
	AutoRenewStatusOn  AutoRenewStatus = 1
)

// The reason an auto-renewable subscription expired.
//
// https://developer.apple.com/documentation/appstoreserverapi/expirationintent
type ExpirationIntent int

var (
	ExpirationIntentCustomerCanceled            ExpirationIntent = 1
	ExpirationIntentBillingError                ExpirationIntent = 2
	ExpirationIntentMissingPriceIncreaseConsent ExpirationIntent = 3
	ExpirationIntentProductUnavailable          ExpirationIntent = 4
	ExpirationIntentOtherReason                 ExpirationIntent = 5
)

// The status that indicates whether an auto-renewable subscription is subject to a price increase.
//
// https://developer.apple.com/documentation/appstoreserverapi/priceincreasestatus
type PriceIncreaseStatus int

var (
	PriceIncreaseStatusNoConsentResponse PriceIncreaseStatus = 0
	PriceIncreaseStatusConsent           PriceIncreaseStatus = 1
)

// The renewal price, in milliunits, of the auto-renewable subscription that renews at the next billing period.
//
// https://developer.apple.com/documentation/appstoreserverapi/renewalprice
type RenewalPrice uint64

func (s RenewalPrice) ToDecimal() decimal.Decimal {
	return decimal.NewFromUint64(uint64(s)).Div(decimal.NewFromInt(1000))
}

type JWSDecodedHeader struct{}

type JWSTransaction string

func (jt JWSTransaction) Decode() (*JWSTransactionDecodedPayload, error) {
	payload := &JWSTransactionDecodedPayload{}
	_, _, err := jwt.NewParser(jwt.WithPaddingAllowed()).ParseUnverified(string(jt), payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// a decoded payload that contains transaction information.
//
// https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
type JWSTransactionDecodedPayload struct {
	jwt.RegisteredClaims
	AppAccountToken             *uuid.UUID         `json:"appAccountToken,omitempty"`
	AppTransactionId            string             `json:"appTransactionId"`
	BundleId                    string             `json:"bundleId"`
	Environment                 Environment        `json:"environment"`
	ExpiresDate                 *common.Timestamp  `json:"expiresDate,omitempty"`
	InAppOwnershipType          InAppOwnershipType `json:"inAppOwnershipType"`
	IsUpgraded                  bool               `json:"isUpgraded"`
	OfferDiscountType           *OfferDiscountType `json:"offerDiscountType,omitempty"`
	OfferIdentifier             *string            `json:"offerIdentifier,omitempty"`
	OfferPeriod                 *OfferPeriod       `json:"offerPeriod,omitempty"`
	OfferType                   *OfferType         `json:"offerType,omitempty"`
	OriginalPurchaseDate        *common.Timestamp  `json:"originalPurchaseDate"`
	OriginalTransactionId       string             `json:"originalTransactionId"`
	ProductId                   string             `json:"productId"`
	PurchaseDate                *common.Timestamp  `json:"purchaseDate"`
	Quantity                    int                `json:"quantity"`
	RevocationDate              *common.Timestamp  `json:"revocationDate,omitempty"`
	RevocationReason            *RevocationReason  `json:"revocationReason,omitempty"`
	SignedDate                  *common.Timestamp  `json:"signedDate"`
	Storefront                  string             `json:"storefront"`
	StorefrontId                string             `json:"storefrontId"`
	SubscriptionGroupIdentifier *string            `json:"subscriptionGroupIdentifier,omitempty"`
	TransactionId               string             `json:"transactionId"`
	TransactionReason           TransactionReason  `json:"transactionReason"`
	Type                        TransactionType    `json:"type"`
	WebOrderLineItemId          *string            `json:"webOrderLineItemId,omitempty"`

	// An integer value that represents the price multiplied by 1000 of the
	// in-app purchase or subscription offer you configured in App Store Connect
	// and that the system records at the time of the purchase
	Price    Price  `json:"price"`
	Currency string `json:"currency"`
}

type JWSRenewalInfo string

func (jt JWSRenewalInfo) Decode() (*JWSRenewalInfoDecodedPayload, error) {
	payload := &JWSRenewalInfoDecodedPayload{}
	_, _, err := jwt.NewParser(jwt.WithPaddingAllowed()).ParseUnverified(string(jt), payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// a decoded payload that contains renewal information.
//
// https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfodecodedpayload
type JWSRenewalInfoDecodedPayload struct {
	jwt.RegisteredClaims
	AppAccountToken             *uuid.UUID           `json:"appAccountToken,omitempty"`
	AppTransactionId            string               `json:"appTransactionId"`
	AutoRenewProductId          string               `json:"autoRenewProductId"`
	AutoRenewStatus             AutoRenewStatus      `json:"autoRenewStatus"`
	Currency                    string               `json:"currency"`
	EligibleWinBackOfferIds     []string             `json:"eligibleWinBackOfferIds"`
	Environment                 Environment          `json:"environment"`
	ExpirationIntent            *ExpirationIntent    `json:"expirationIntent,omitempty"`
	GracePeriodExpiresDate      *common.Timestamp    `json:"gracePeriodExpiresDate"`
	IsInBillingRetryPeriod      bool                 `json:"isInBillingRetryPeriod"`
	OfferDiscountType           *OfferDiscountType   `json:"offerDiscountType,omitempty"`
	OfferIdentifier             *string              `json:"offerIdentifier,omitempty"`
	OfferPeriod                 *OfferPeriod         `json:"offerPeriod,omitempty"`
	OfferType                   *OfferType           `json:"offerType,omitempty"`
	OriginalTransactionId       string               `json:"originalTransactionId"`
	PriceIncreaseStatus         *PriceIncreaseStatus `json:"priceIncreaseStatus,omitempty"`
	ProductId                   string               `json:"productId"`
	RecentSubscriptionStartDate *common.Timestamp    `json:"recentSubscriptionStartDate"`
	RenewalDate                 *common.Timestamp    `json:"renewalDate"`
	RenewalPrice                RenewalPrice         `json:"renewalPrice"`
	SignedDate                  *common.Timestamp    `json:"signedDate"`
}
