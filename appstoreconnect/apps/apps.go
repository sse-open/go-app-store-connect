package apps

import (
	"context"

	"github.com/sse-open/go-app-store-connect/appstoreconnect/common"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/included"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/resource/apps"
	"github.com/sse-open/go-app-store-connect/client"
)

// A response that contains a list of Apps resources.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appsresponse
type AppsResponse struct {
	Data     []apps.App                `json:"data"`
	Included []included.Included       `json:"included,omitempty"` // Skipping declaration of included type payloads for now
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

// Available query parameters for List Apps
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-apps
type ListAppsQuery struct {
	FieldsApps                            string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements           string `url:"fields[betaLicenseAgreements],omitempty"`
	FilterBundleID                        string `url:"filter[bundleId],omitempty"`
	FilterID                              string `url:"filter[id],omitempty"`
	FilterName                            string `url:"filter[name],omitempty"`
	FilterSKU                             string `url:"filter[sku],omitempty"`
	Include                               string `url:"include,omitempty"`
	Limit                                 int    `url:"limit,omitempty"`
	Sort                                  string `url:"sort,omitempty"`
	FieldsPreReleaseVersions              string `url:"fields[preReleaseVersions],omitempty"`
	LimitPreReleaseVersions               int    `url:"limit[preReleaseVersions],omitempty"`
	FieldsBetaAppReviewDetails            string `url:"fields[betaAppReviewDetails],omitempty"`
	FieldsBetaAppLocalizations            string `url:"fields[betaAppLocalizations],omitempty"`
	FieldsBuilds                          string `url:"fields[builds],omitempty"`
	FieldsBetaGroups                      string `url:"fields[betaGroups],omitempty"`
	LimitBuilds                           int    `url:"limit[builds],omitempty"`
	LimitBetaGroups                       int    `url:"limit[betaGroups],omitempty"`
	LimitBetaAppLocalizations             int    `url:"limit[betaAppLocalizations],omitempty"`
	LimitAppStoreVersions                 int    `url:"limit[appStoreVersions],omitempty"`
	LimitAppInfos                         int    `url:"limit[appInfos],omitempty"`
	FieldsEndUserLicenseAgreements        string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsAppStoreVersions                string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppInfos                        string `url:"fields[appInfos],omitempty"`
	FilterAppStoreVersions                string `url:"filter[appStoreVersions],omitempty"`
	FilterAppStoreVersionsPlatform        string `url:"filter[appStoreVersions.platform],omitempty"`
	FilterAppStoreVersionsAppStoreState   string `url:"filter[appStoreVersions.appStoreState],omitempty"`
	FieldsInAppPurchases                  string `url:"fields[inAppPurchases],omitempty"`
	FieldsCIProducts                      string `url:"fields[ciProducts],omitempty"`
	LimitAppClips                         int    `url:"limit[appClips],omitempty"`
	FieldsAppClips                        string `url:"fields[appClips],omitempty"`
	FieldsReviewSubmissions               string `url:"fields[reviewSubmissions],omitempty"`
	FieldsAppCustomProductPages           string `url:"fields[appCustomProductPages],omitempty"`
	FieldsAppEvents                       string `url:"fields[appEvents],omitempty"`
	LimitAppCustomProductPages            int    `url:"limit[appCustomProductPages],omitempty"`
	LimitAppEvents                        int    `url:"limit[appEvents],omitempty"`
	LimitReviewSubmissions                int    `url:"limit[reviewSubmissions],omitempty"`
	FieldsSubscriptionGracePeriods        string `url:"fields[subscriptionGracePeriods],omitempty"`
	FieldsPromotedPurchases               string `url:"fields[promotedPurchases],omitempty"`
	FieldsSubscriptionGroups              string `url:"fields[subscriptionGroups],omitempty"`
	LimitInAppPurchasesV2                 int    `url:"limit[inAppPurchasesV2],omitempty"`
	LimitPromotedPurchases                int    `url:"limit[promotedPurchases],omitempty"`
	LimitSubscriptionGroups               int    `url:"limit[subscriptionGroups],omitempty"`
	FieldsAppStoreVersionExperiments      string `url:"fields[appStoreVersionExperiments],omitempty"`
	LimitAppStoreVersionExperimentsV2     int    `url:"limit[appStoreVersionExperimentsV2],omitempty"`
	FieldsAppEncryptionDeclarations       string `url:"fields[appEncryptionDeclarations],omitempty"`
	LimitAppEncryptionDeclarations        int    `url:"limit[appEncryptionDeclarations],omitempty"`
	FieldsGameCenterDetails               string `url:"fields[gameCenterDetails],omitempty"`
	FieldsAppStoreVersionsAppVersionState string `url:"filter[appStoreVersions.appVersionState],omitempty"`
	FilterReviewSubmissionsPlatform       string `url:"filter[reviewSubmissions.platform],omitempty"`
	FilterReviewSubmissionsState          string `url:"filter[reviewSubmissions.state],omitempty"`
	Cursor                                string `url:"cursor,omitempty"`
}

// List available apps
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-apps
func (as *AppsService) ListApps(ctx context.Context, queryParams *ListAppsQuery) (*AppsResponse, *client.ClientResponse, error) {
	res := new(AppsResponse)
	resp, err := as.client.Get(ctx, "v1/apps", queryParams, res)
	if err != nil {
		return nil, nil, err
	}

	return res, resp, nil
}
