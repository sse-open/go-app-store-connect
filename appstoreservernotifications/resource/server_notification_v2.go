package resource

import "github.com/google/uuid"

// The type that describes the In-App Purchase or external purchase event for which the App Store sends the version 2 notification.
//
// https://developer.apple.com/documentation/appstoreservernotifications/notificationtype
type NotificationType string

var (
	NotificationTypeConsumptionRequest     NotificationType = "CONSUMPTION_REQUEST"
	NotificationTypeDidChangeRenewalPref   NotificationType = "DID_CHANGE_RENEWAL_PREF"
	NotificationTypeDidChangeRenewalStatus NotificationType = "DID_CHANGE_RENEWAL_STATUS"
	NotificationTypeDidFailToRenew         NotificationType = "DID_FAIL_TO_RENEW"
	NotificationTypeDidRenew               NotificationType = "DID_RENEW"
	NotificationTypeExpired                NotificationType = "EXPIRED"
	NotificationTypeExternalPurchaseToken  NotificationType = "EXTERNAL_PURCHASE_TOKEN"
	NotificationTypeGracePeriodExpired     NotificationType = "GRACE_PERIOD_EXPIRED"
	NotificationTypeMetadataUpdate         NotificationType = "METADATA_UPDATE"
	NotificationTypeMigration              NotificationType = "MIGRATION"
	NotificationTypeOfferRedeemed          NotificationType = "OFFER_REDEEMED"
	NotificationTypeOneTimeCharge          NotificationType = "ONE_TIME_CHARGE"
	NotificationTypePriceChange            NotificationType = "PRICE_CHANGE"
	NotificationTypePriceIncrease          NotificationType = "PRICE_INCREASE"
	NotificationTypeRefund                 NotificationType = "REFUND"
	NotificationTypeRefundDeclined         NotificationType = "REFUND_DECLINED"
	NotificationTypeRefundReversed         NotificationType = "REFUND_REVERSED"
	NotificationTypeRenewalExtended        NotificationType = "RENEWAL_EXTENDED"
	NotificationTypeRenewalExtension       NotificationType = "RENEWAL_EXTENSION"
	NotificationTypeRescindConsent         NotificationType = "RESCIND_CONSENT"
	NotificationTypeRevoke                 NotificationType = "REVOKE"
	NotificationTypeSubscribed             NotificationType = "SUBSCRIBED"
	NotificationTypeTest                   NotificationType = "TEST"
)

// The sub type that describes the In-App Purchase or external purchase event for which the App Store sends the version 2 notification.
//
// https://developer.apple.com/documentation/appstoreservernotifications/subtype
type NotificationSubType string

var (
	NotificationSubTypeUpgrade             NotificationSubType = "UPGRADE"
	NotificationSubTypeDowngrade           NotificationSubType = "DOWNGRADE"
	NotificationSubTypeAutoRenewEnabled    NotificationSubType = "AUTO_RENEW_ENABLED"
	NotificationSubTypeAutoRenewDisabled   NotificationSubType = "AUTO_RENEW_DISABLED"
	NotificationSubTypeGracePeriod         NotificationSubType = "GRACE_PERIOD"
	NotificationSubTypeBillingRecovery     NotificationSubType = "BILLING_RECOVERY"
	NotificationSubTypeVoluntary           NotificationSubType = "VOLUNTARY"
	NotificationSubTypeBillingRetry        NotificationSubType = "BILLING_RETRY"
	NotificationSubTypePriceIncrease       NotificationSubType = "PRICE_INCREASE"
	NotificationSubTypeProductNotForSale   NotificationSubType = "PRODUCT_NOT_FOR_SALE"
	NotificationSubTypeCreated             NotificationSubType = "CREATED"
	NotificationSubTypeActiveTokenReminder NotificationSubType = "ACTIVE_TOKEN_REMINDER"
	NotificationSubTypeUnreported          NotificationSubType = "UNREPORTED"
	NotificationSubTypePending             NotificationSubType = "PENDING"
	NotificationSubTypeAccepted            NotificationSubType = "ACCEPTED"
	NotificationSubTypeSummary             NotificationSubType = "SUMMARY"
	NotificationSubTypeFailure             NotificationSubType = "FAILURE"
	NotificationSubTypeResubscribe         NotificationSubType = "RESUBSCRIBE"
	NotificationSubTypeInitialBuy          NotificationSubType = "INITIAL_BUY"
)

// The response body the App Store sends in a version 2 server notification.
//
// https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2
type ServerNotificationResponseBodyV2 struct {
	SignedPayload string `json:"signedPayload"`
}

// The customer-provided reason for a refund request.
//
// https://developer.apple.com/documentation/appstoreservernotifications/consumptionrequestreason
type ConsumptionRequestReason string

var (
	ConsumptionRequestReasonUnintendedPurchase      ConsumptionRequestReason = "UNINTENDED_PURCHASE"
	ConsumptionRequestReasonFulfillmentIssue        ConsumptionRequestReason = "FULFILLMENT_ISSUE"
	ConsumptionRequestReasonUnsatisfiedWithPurchase ConsumptionRequestReason = "UNSATISFIED_WITH_PURCHASE"
	ConsumptionRequestReasonLegal                   ConsumptionRequestReason = "LEGAL"
	ConsumptionRequestReasonOther                   ConsumptionRequestReason = "OTHER"
)

// The status of an auto-renewable subscription at the time the App Store signs the notification.
//
// https://developer.apple.com/documentation/appstoreservernotifications/status
type NotificationStatus int

var (
	NotificationStatusActive             NotificationStatus = 1
	NotificationStatusExpired            NotificationStatus = 2
	NotificationStatusBillingRetryPeriod NotificationStatus = 3
	NotificationStatusBillingGracePeriod NotificationStatus = 4
	NotificationStatusRevoked            NotificationStatus = 5
)

// The payload data that contains app metadata and the signed renewal and transaction information.
//
// https://developer.apple.com/documentation/appstoreservernotifications/data
type NotificationData struct {
	AppAppleID                      *int64                    `json:"appAppleId,omitempty"`
	BundleID                        string                    `json:"bundleId"`
	BundleVersion                   *string                   `json:"bundleVersion,omitempty"`
	ConsumptionRequestReason        *ConsumptionRequestReason `json:"consumptionRequestReason,omitempty"`
	Environment                     Environment               `json:"environment"`
	SignedRenewalInfo               *string                   `json:"signedRenewalInfo,omitempty"`
	SignedTransactionInfo           string                    `json:"signedTransactionInfo"`
	AutoRenewableSubscriptionStatus *NotificationStatus       `json:"status,omitempty"`
}

// The payload data for a subscription-renewal-date extension notification.
//
// https://developer.apple.com/documentation/appstoreservernotifications/summary
type NotificationSummary struct {
	RequestIdentifier      uuid.UUID   `json:"requestIdentifier"`
	Environment            Environment `json:"environment"`
	AppAppleID             *int64      `json:"appAppleId,omitempty"`
	BundleID               string      `json:"bundleId"`
	ProductID              string      `json:"productId"`
	StorefrontCountryCodes []string    `json:"storefrontCountryCodes,omitempty"`
	FailedCount            *int64      `json:"failedCount,omitempty"`
	SucceededCount         *int64      `json:"succeededCount,omitempty"`
}

// The type of an external purchase custom link token.
//
// https://developer.apple.com/documentation/appstoreservernotifications/tokentype
type ExternalPurchaseTokenType string

var (
	ExternalPurchaseTokenTypeAcquisition ExternalPurchaseTokenType = "ACQUISITION"
	ExternalPurchaseTokenTypeService     ExternalPurchaseTokenType = "SERVICE"
)

// The payload data that contains an external purchase token.
//
// https://developer.apple.com/documentation/appstoreservernotifications/externalpurchasetoken
type NotificationExternalPurchaseToken struct {
	ExternalPurchaseID  string                     `json:"externalPurchaseId"`
	TokenCreationDate   Timestamp                  `json:"tokenCreationDate"`
	AppAppleID          int64                      `json:"appAppleId"`
	BundleID            string                     `json:"bundleId"`
	TokenExpirationDate *Timestamp                 `json:"tokenExpirationDate,omitempty"`
	TokenType           *ExternalPurchaseTokenType `json:"tokenType,omitempty"`
}

// The object that contains the app metadata and signed app transaction information.
//
// https://developer.apple.com/documentation/appstoreservernotifications/appdata
type NotificationAppData struct {
	AppAppleID               *int64      `json:"appAppleId,omitempty"`
	BundleID                 string      `json:"bundleId"`
	Environment              Environment `json:"environment"`
	SignedAppTransactionInfo string      `json:"signedAppTransactionInfo"`
}

// A decoded payload that contains the version 2 notification data.
//
// https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2decodedpayload
type ServerNotificationResponseBodyV2DecodedPayload struct {
	NotificationType      NotificationType                   `json:"notificationType"`
	Subtype               NotificationSubType                `json:"subtype,omitempty"`
	Data                  *NotificationData                  `json:"data,omitempty"`
	Summary               *NotificationSummary               `json:"summary,omitempty"`
	ExternalPurchaseToken *NotificationExternalPurchaseToken `json:"externalPurchaseToken,omitempty"`
	AppData               *NotificationAppData               `json:"appData,omitempty"`
	Version               string                             `json:"version"`
	SignedDate            Timestamp                          `json:"signedDate"`
	NotificationUUID      string                             `json:"notificationUUID"`
}
