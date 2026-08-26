package common

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/sse-open/go-app-store-connect/common"
	commonresource "github.com/sse-open/go-app-store-connect/common/resource"
	"github.com/sse-open/go-app-store-connect/common/resource/certificate"
	"github.com/sse-open/go-app-store-connect/common/resource/jws"
)

// https://developer.apple.com/documentation/appstoreservernotifications/billingplantype
type BillingPlanType string

var (
	BillingPlanTypeBilledUpFront BillingPlanType = "BILLED_UPFRONT"
	BillingPlanTypeMonthly       BillingPlanType = "MONTHLY"
)

// A string that describes whether the transaction was purchased by the customer, or is available to them through Family Sharing.
//
// https://developer.apple.com/documentation/appstoreservernotifications/inappownershiptype
type InAppOwnershipType string

var (
	InAppOwnershipTypeFamilyShared InAppOwnershipType = "FAMILY_SHARED"
	InAppOwnershipTypePurchased    InAppOwnershipType = "PURCHASED"
)

// The payment mode for a discount offer on an In-App Purchase.
//
// https://developer.apple.com/documentation/appstoreservernotifications/offerdiscounttype
type OfferDiscountType string

var (
	OfferDiscountTypeFreeTrial  OfferDiscountType = "FREE_TRIAL"
	OfferDiscountTypePayAsYouGo OfferDiscountType = "PAY_AS_YOU_GO"
	OfferDiscountTypePayUpFront OfferDiscountType = "PAY_UP_FRONT"
	OfferDiscountTypeOneTime    OfferDiscountType = "ONE_TIME"
)

// The duration of the offer.
//
// https://developer.apple.com/documentation/appstoreservernotifications/offerperiod
type OfferPeriod string

var (
	OfferPeriodOneMonth  OfferPeriod = "P1M"
	OfferPeriodTwoMonth  OfferPeriod = "P2M"
	OfferPeriodThreeDays OfferPeriod = "P3D"
)

// The type of offer.
//
// https://developer.apple.com/documentation/appstoreservernotifications/offertype
type OfferType int

var (
	OfferTypeIntroductory OfferType = 1
	OfferTypePromotional  OfferType = 2
	OfferTypeOfferCode    OfferType = 3
	OfferTypeWinBack      OfferType = 4
)

// The reason for a revoked or refunded transaction.
//
// https://developer.apple.com/documentation/appstoreservernotifications/revocationreason
type RevocationReason int

var (
	RevocationReasonOtherReason          RevocationReason = 0
	RevocationReasonActualPerceivedIssue RevocationReason = 1
)

// The type of the refund or revocation that applies to the transaction.
//
// https://developer.apple.com/documentation/appstoreservernotifications/revocationtype
type RevocationType string

var (
	RevocationTypeRefundFull     RevocationType = "REFUND_FULL"
	RevocationTypeRefundProrated RevocationType = "REFUND_PRORATED"
	RevocationTypeFamilyRevoke   RevocationType = "FAMILY_REVOKE"
)

// The cause of a purchase transaction, which indicates whether it’s a customer’s purchase or a
// renewal for an auto-renewable subscription that the system initiates.
//
// https://developer.apple.com/documentation/appstoreservernotifications/transactionreason
type TransactionReason string

var (
	TransactionReasonPurchase TransactionReason = "PURCHASE"
	TransactionReasonRenewal  TransactionReason = "RENEWAL"
)

// The product type of the In-App Purchase.
//
// https://developer.apple.com/documentation/appstoreservernotifications/type
type TransactionType string

var (
	TransactionTypeAutoRenewableSubscription TransactionType = "Auto-Renewable Subscription"
	TransactionTypeNonConsumable             TransactionType = "Non-Consumable"
	TransactionTypeConsumable                TransactionType = "Consumable"
	TransactionTypeNonRenewingSubscription   TransactionType = "Non-Renewing Subscription"
)

// https://developer.apple.com/documentation/appstoreservernotifications/transactioncommitmentinfo
type TransactionCommitmentInfo struct {
	BillingPeriodNumber   int               `json:"billingPeriodNumber,omitempty"`
	TotalBillingPeriods   int               `json:"totalBillingPeriods,omitempty"`
	CommitmentExpiresDate *common.Timestamp `json:"commitmentExpiresDate,omitempty"`
	CommitmentPrice       int64             `json:"commitmentPrice,omitempty"`
}

// The price, in milliunits, of the In-App Purchase that the system records in the transaction.
//
// https://developer.apple.com/documentation/appstoreservernotifications/price
type Price uint64

func (s Price) ToDecimal() decimal.Decimal {
	return decimal.NewFromUint64(uint64(s)).Div(decimal.NewFromInt(1000))
}

type JWSTransaction string

func (jt JWSTransaction) Decode() (*JWSTransactionDecodedPayload, error) {
	payload := &JWSTransactionDecodedPayload{}
	_, _, err := jwt.NewParser(jwt.WithPaddingAllowed()).ParseUnverified(string(jt), payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (jt JWSTransaction) VerifyClaims(ctx context.Context, rootCertificateProvider *certificate.RootCertificateProvider) (*JWSTransactionDecodedPayload, error) {
	claims, err := jws.VerifyJWS[JWSTransactionDecodedPayload](ctx, rootCertificateProvider, string(jt))
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

// A decoded payload that contains transaction information.
//
// https://developer.apple.com/documentation/appstoreservernotifications/jwstransactiondecodedpayload
type JWSTransactionDecodedPayload struct {
	jwt.RegisteredClaims
	AppAccountToken               *uuid.UUID                 `json:"appAccountToken,omitempty"`
	AppTransactionId              string                     `json:"appTransactionId"`
	BundleId                      string                     `json:"bundleId"`
	BillingPlanType               *BillingPlanType           `json:"billingPlanType,omitempty"`
	CommitmentInfo                *TransactionCommitmentInfo `json:"commitmentInfo,omitempty"`
	Environment                   commonresource.Environment `json:"environment"`
	ExpiresDate                   *common.Timestamp          `json:"expiresDate,omitempty"`
	InAppOwnershipType            InAppOwnershipType         `json:"inAppOwnershipType"`
	IsUpgraded                    bool                       `json:"isUpgraded"`
	OfferDiscountType             *OfferDiscountType         `json:"offerDiscountType,omitempty"`
	OfferIdentifier               *string                    `json:"offerIdentifier,omitempty"`
	OfferPeriod                   *OfferPeriod               `json:"offerPeriod,omitempty"`
	OfferType                     *OfferType                 `json:"offerType,omitempty"`
	OriginalPurchaseDate          *common.Timestamp          `json:"originalPurchaseDate"`
	OriginalTransactionId         string                     `json:"originalTransactionId"`
	PreviousOriginalTransactionId *string                    `json:"previousOriginalTransactionId,omitempty"`
	ProductId                     string                     `json:"productId"`
	PurchaseDate                  *common.Timestamp          `json:"purchaseDate"`
	Quantity                      int                        `json:"quantity"`
	RevocationDate                *common.Timestamp          `json:"revocationDate,omitempty"`
	RevocationPercentage          *int                       `json:"revocationPercentage,omitempty"`
	RevocationReason              *RevocationReason          `json:"revocationReason,omitempty"`
	RevocationType                *RevocationType            `json:"revocationType,omitempty"`
	SignedDate                    *common.Timestamp          `json:"signedDate"`
	Storefront                    string                     `json:"storefront"`
	StorefrontId                  string                     `json:"storefrontId"`
	SubscriptionGroupIdentifier   *string                    `json:"subscriptionGroupIdentifier,omitempty"`
	TransactionId                 string                     `json:"transactionId"`
	TransactionReason             TransactionReason          `json:"transactionReason"`
	Type                          TransactionType            `json:"type"`
	WebOrderLineItemId            *string                    `json:"webOrderLineItemId,omitempty"`

	// An integer value that represents the price multiplied by 1000 of the
	// in-app purchase or subscription offer you configured in App Store Connect
	// and that the system records at the time of the purchase
	Price    *Price  `json:"price,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// The renewal status for an auto-renewable subscription.
//
// https://developer.apple.com/documentation/appstoreservernotifications/autorenewstatus
type AutoRenewStatus int

var (
	AutoRenewStatusOff AutoRenewStatus = 0
	AutoRenewStatusOn  AutoRenewStatus = 1
)

// The reason an auto-renewable subscription expired.
//
// https://developer.apple.com/documentation/appstoreservernotifications/expirationintent
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
// https://developer.apple.com/documentation/appstoreservernotifications/priceincreasestatus
type PriceIncreaseStatus int

var (
	PriceIncreaseStatusNoConsentResponse PriceIncreaseStatus = 0
	PriceIncreaseStatusConsent           PriceIncreaseStatus = 1
)

// https://developer.apple.com/documentation/appstoreservernotifications/renewalbillingplantype
type RenewalBillingPlanType string

var (
	RenewalBillingPlanTypeBilledUpFront RenewalBillingPlanType = "BILLED_UPFRONT"
	RenewalBillingPlanTypeMonthly       RenewalBillingPlanType = "MONTHLY"
)

// The renewal price, in milliunits, of the auto-renewable subscription that renews at the next billing period.
//
// https://developer.apple.com/documentation/appstoreservernotifications/renewalprice
type RenewalPrice uint64

func (s RenewalPrice) ToDecimal() decimal.Decimal {
	return decimal.NewFromUint64(uint64(s)).Div(decimal.NewFromInt(1000))
}

type CommitmentRenewalBillingPlanType string

var (
	CommitmentRenewalBillingPlanTypeBilledUpFront CommitmentRenewalBillingPlanType = "BILLED_UPFRONT"
	CommitmentRenewalBillingPlanTypeMonthly       CommitmentRenewalBillingPlanType = "MONTHLY"
)

// https://developer.apple.com/documentation/appstoreservernotifications/renewalcommitmentinfo
type RenewalCommitmentInfo struct {
	CommitmentAutoRenewProductId     string                            `json:"commitmentAutoRenewProductId,omitempty"`
	CommitmentAutoRenewStatus        int                               `json:"commitmentAutoRenewStatus,omitempty"`
	CommitmentRenewalBillingPlanType *CommitmentRenewalBillingPlanType `json:"commitmentRenewalBillingPlanType,omitempty"`
	CommitmentRenewalDate            *common.Timestamp                 `json:"commitmentRenewalDate,omitempty"`
	CommitmentRenewalPrice           int64                             `json:"commitmentRenewalPrice,omitempty"`
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

func (jt JWSRenewalInfo) VerifyClaims(ctx context.Context, rootCertificateProvider *certificate.RootCertificateProvider) (*JWSRenewalInfoDecodedPayload, error) {
	claims, err := jws.VerifyJWS[JWSRenewalInfoDecodedPayload](ctx, rootCertificateProvider, string(jt))
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

// A decoded payload containing subscription renewal information for an auto-renewable subscription.
//
// https://developer.apple.com/documentation/appstoreservernotifications/jwsrenewalinfodecodedpayload
type JWSRenewalInfoDecodedPayload struct {
	jwt.RegisteredClaims
	AppAccountToken             *uuid.UUID                 `json:"appAccountToken,omitempty"`
	AppTransactionId            string                     `json:"appTransactionId"`
	AutoRenewProductId          string                     `json:"autoRenewProductId"`
	AutoRenewStatus             AutoRenewStatus            `json:"autoRenewStatus"`
	Currency                    string                     `json:"currency"`
	CommitmentInfo              *RenewalCommitmentInfo     `json:"commitmentInfo,omitempty"`
	EligibleWinBackOfferIds     []string                   `json:"eligibleWinBackOfferIds"`
	Environment                 commonresource.Environment `json:"environment"`
	ExpirationIntent            *ExpirationIntent          `json:"expirationIntent,omitempty"`
	GracePeriodExpiresDate      *common.Timestamp          `json:"gracePeriodExpiresDate"`
	IsInBillingRetryPeriod      bool                       `json:"isInBillingRetryPeriod"`
	OfferDiscountType           *OfferDiscountType         `json:"offerDiscountType,omitempty"`
	OfferIdentifier             *string                    `json:"offerIdentifier,omitempty"`
	OfferPeriod                 *OfferPeriod               `json:"offerPeriod,omitempty"`
	OfferType                   *OfferType                 `json:"offerType,omitempty"`
	OriginalTransactionId       string                     `json:"originalTransactionId"`
	PriceIncreaseStatus         *PriceIncreaseStatus       `json:"priceIncreaseStatus,omitempty"`
	ProductId                   string                     `json:"productId"`
	RecentSubscriptionStartDate *common.Timestamp          `json:"recentSubscriptionStartDate,omitempty"`
	RenewalBillingPlanType      *RenewalBillingPlanType    `json:"renewalBillingPlanType,omitempty"`
	RenewalDate                 *common.Timestamp          `json:"renewalDate"`
	RenewalPrice                RenewalPrice               `json:"renewalPrice"`
	SignedDate                  *common.Timestamp          `json:"signedDate"`
}
