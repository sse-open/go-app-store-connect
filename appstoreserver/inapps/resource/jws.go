package resource

import "github.com/golang-jwt/jwt/v5"

type JWSDecodedHeader struct{}

type JWSTransaction string

func (jt JWSTransaction) Decode() (*JWSTransactionDecodedPayload, error) {
	payload := &JWSTransactionDecodedPayload{}
	_, _, err := jwt.NewParser().ParseUnverified(string(jt), payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

type JWSTransactionDecodedPayload struct {
	jwt.RegisteredClaims
	AppAccountToken             string `json:"appAccountToken"`
	AppTransactionId            string `json:"appTransactionId"`
	BundleId                    string `json:"bundleId"`
	Environment                 string `json:"environment"`
	ExpiresDate                 int    `json:"expiresDate"`
	InAppOwnershipType          string `json:"inAppOwnershipType"`
	IsUpgraded                  bool   `json:"isUpgraded"`
	OfferDiscountType           string `json:"offerDiscountType"`
	OfferIdentifier             string `json:"offerIdentifier"`
	OfferPeriod                 string `json:"offerPeriod"`
	OfferType                   string `json:"offerType"`
	OriginalPurchaseDate        int    `json:"originalPurchaseDate"`
	OriginalTransactionId       string `json:"originalTransactionId"`
	ProductId                   string `json:"productId"`
	PurchaseDate                int    `json:"purchaseDate"`
	Quantity                    int    `json:"quantity"`
	RevocationDate              int    `json:"revocationDate"`
	RevocationReason            string `json:"revocationReason"`
	SignedDate                  int    `json:"signedDate"`
	Storefront                  string `json:"storefront"`
	StorefrontId                string `json:"storefrontId"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	TransactionId               string `json:"transactionId"`
	TransactionReason           string `json:"transactionReason"`
	Type                        string `json:"type"`
	WebOrderLineItemId          string `json:"webOrderLineItemId"`

	// An integer value that represents the price multiplied by 1000 of the
	// in-app purchase or subscription offer you configured in App Store Connect
	// and that the system records at the time of the purchase
	Price    int64  `json:"price"`
	Currency string `json:"currency"`
}

type JWSRenewalInfo string

func (jt JWSRenewalInfo) Decode() (*JWSRenewalInfoDecodedPayload, error) {
	payload := &JWSRenewalInfoDecodedPayload{}
	_, _, err := jwt.NewParser().ParseUnverified(string(jt), payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

type JWSRenewalInfoDecodedPayload struct {
	jwt.RegisteredClaims
	AppAccountToken             string   `json:"appAccountToken"`
	AppTransactionId            string   `json:"appTransactionId"`
	AutoRenewProductId          string   `json:"autoRenewProductId"`
	AutoRenewStatus             string   `json:"autoRenewStatus"`
	Currency                    string   `json:"currency"`
	EligibleWinBackOfferIds     []string `json:"eligibleWinBackOfferIds"`
	Environment                 string   `json:"environment"`
	ExpirationIntent            string   `json:"expirationIntent"`
	GracePeriodExpiresDate      int64    `json:"gracePeriodExpiresDate"`
	IsInBillingRetryPeriod      bool     `json:"isInBillingRetryPeriod"`
	OfferDiscountType           string   `json:"offerDiscountType"`
	OfferIdentifier             string   `json:"offerIdentifier"`
	OfferPeriod                 string   `json:"offerPeriod"`
	OfferType                   string   `json:"offerType"`
	OriginalTransactionId       string   `json:"originalTransactionId"`
	PriceIncreaseStatus         string   `json:"priceIncreaseStatus"`
	ProductId                   string   `json:"productId"`
	RecentSubscriptionStartDate int64    `json:"recentSubscriptionStartDate"`
	RenewalDate                 int64    `json:"renewalDate"`
	RenewalPrice                int64    `json:"renewalPrice"`
	SignedDate                  int64    `json:"signedDate"`
}
