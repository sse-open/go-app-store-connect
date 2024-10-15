package apps

import "github.com/sse-open/go-app-store-connect/appstoreconnect/common"

// Attributes that describe an Apps resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/app/attributes-data.dictionary
type AppAttributes struct {
	BundleID                               *string `json:"bundleId,omitempty"`
	Name                                   *string `json:"name,omitempty"`
	PrimaryLocale                          *string `json:"primaryLocale,omitempty"`
	Sku                                    *string `json:"sku,omitempty"`
	ContentRightsDeclaration               *string `json:"contentRightsDeclaration,omitempty"`
	IsOrEverWasMadeForKids                 *bool   `json:"isOrEverWasMadeForKids,omitempty"`
	SubscriptionStatusUrl                  *string `json:"subscriptionStatusUrl,omitempty"`
	SubscriptionStatusUrlForSandbox        *string `json:"subscriptionStatusUrlForSandbox,omitempty"`
	SubscriptionStatusUrlVersion           *string `json:"subscriptionStatusUrlVersion,omitempty"`
	SubscriptionStatusUrlVersionForSandbox *string `json:"subscriptionStatusUrlVersionForSandbox,omitempty"`
	StreamlinedPurchasingEnabled           *bool   `json:"streamlinedPurchasingEnabled,omitempty"`
}

// The relationships you included in the request and those on which you can operate.
//
// https://developer.apple.com/documentation/appstoreconnectapi/app/relationships-data.dictionary
type AppRelationships struct {
	BetaLicenseAgreement         *common.Relationship          `json:"betaLicenseAgreement,omitempty"`
	PreReleaseVersions           *common.PagedRelationship     `json:"preReleaseVersions,omitempty"`
	BetaAppLocalizations         *common.PagedRelationship     `json:"betaAppLocalizations,omitempty"`
	BetaGroups                   *common.PagedRelationship     `json:"betaGroups,omitempty"`
	Builds                       *common.PagedRelationship     `json:"builds,omitempty"`
	BetaAppReviewDetail          *common.Relationship          `json:"betaAppReviewDetail,omitempty"`
	AppInfos                     *common.PagedRelationship     `json:"appInfos,omitempty"`
	AppStoreVersions             *common.PagedRelationship     `json:"appStoreVersions,omitempty"`
	EndUserLicenseAgreement      *common.Relationship          `json:"endUserLicenseAgreement,omitempty"`
	PreOrder                     *common.Relationship          `json:"preOrder,omitempty"`
	GameCenterEnabledVersions    *common.PagedRelationship     `json:"gameCenterEnabledVersions,omitempty"`
	CIProduct                    *common.Relationship          `json:"ciProduct,omitempty"`
	AppClips                     *common.PagedRelationship     `json:"appClips,omitempty"`
	AppCustomProductPages        *common.PagedRelationship     `json:"appCustomProductPages,omitempty"`
	AppEvents                    *common.PagedRelationship     `json:"appEvents,omitempty"`
	ReviewSubmissions            *common.PagedRelationship     `json:"reviewSubmissions,omitempty"`
	SubscriptionGracePeriods     *common.Relationship          `json:"subscriptionGracePeriods,omitempty"`
	InAppPurchasesV2             *common.PagedRelationship     `json:"inAppPurchasesV2,omitempty"`
	PromotedPurchases            *common.PagedRelationship     `json:"promotedPurchases,omitempty"`
	SubscriptionGroups           *common.PagedRelationship     `json:"subscriptionGroups,omitempty"`
	AppStoreVersionExperimentsV2 *common.PagedRelationship     `json:"appStoreVersionExperimentsV2,omitempty"`
	AppEncryptionDeclarations    *common.PagedRelationship     `json:"appEncryptionDeclarations,omitempty"`
	GameCenterDetails            *common.Relationship          `json:"gameCenterDetails,omitempty"`
	AlternativeDistributionKeys  *common.RelationshipLinksOnly `json:"alternativeDistributionKeys,omitempty"`
	AnalyticsReportRequests      *common.RelationshipLinksOnly `json:"analyticsReportRequests,omitempty"`
	AppAvailability              *common.RelationshipLinksOnly `json:"appAvailability,omitempty"`
	AppAvailabilityV2            *common.RelationshipLinksOnly `json:"appAvailabilityV2,omitempty"`
	AppPricePoints               *common.RelationshipLinksOnly `json:"appPricePoints,omitempty"`
	AppPriceSchedule             *common.RelationshipLinksOnly `json:"appPriceSchedule,omitempty"`
	BetaTesters                  *common.RelationshipLinksOnly `json:"betaTesters,omitempty"`
	CustomerReviews              *common.RelationshipLinksOnly `json:"customerReviews,omitempty"`
	MarketplaceSearchDetail      *common.RelationshipLinksOnly `json:"marketplaceSearchDetail,omitempty"`
	PerfPowerMetrics             *common.RelationshipLinksOnly `json:"perfPowerMetrics,omitempty"`
}

// The data structure that represents an Apps resource.
//
// https://developer.apple.com/documentation/appstoreconnectapi/app
type App struct {
	Attributes    *AppAttributes       `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Relationships *AppRelationships    `json:"relationships,omitempty"`
	Type          string               `json:"type"`
	Links         common.ResourceLinks `json:"links"`
}
