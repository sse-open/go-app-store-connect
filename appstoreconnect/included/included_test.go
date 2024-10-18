package included

import (
	"encoding/json"
	"testing"

	"github.com/sse-open/go-app-store-connect/appstoreconnect/resource/inapppurchase"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/resource/territories"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalTerritory(t *testing.T) {
	t.Parallel()

	payload := []byte(`{
		"type": "territories",
		"attributes": {
			"currency": "USD"
		},
		"id": "USA",
		"links": {
			"self": "https://api.appstoreconnect.apple.com/v1/territories/USA"
		}
	}`)

	var included Included
	err := json.Unmarshal(payload, &included)
	assert.NoError(t, err)

	territory := included.TypeData.(territories.Territory)
	assert.Equal(t, "USA", territory.ID)
	assert.Equal(t, "USD", territory.Attributes.Currency)
}

func TestUnmarshalInAppPurchasePricePoint(t *testing.T) {
	t.Parallel()

	payload := []byte(`{
		"attributes": {
			"customerPrice": "4.99",
			"proceeds": "3.5"
		},
		"id": "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ",
		"links": {
			"self": "https://api.appstoreconnect.apple.com/v1/inAppPurchasePricePoints/eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ"
		},
		"relationships": {
			"territory": {}
		},
		"type": "inAppPurchasePricePoints"
	}`)

	var included Included
	err := json.Unmarshal(payload, &included)
	assert.NoError(t, err)

	iapPricePoint := included.TypeData.(inapppurchase.InAppPurchasePricePoint)
	assert.Equal(t, "eyJzIjoiNjUwMjg2ODgwMCIsInQiOiJVU0EiLCJwIjoiMTAwNjIifQ", iapPricePoint.ID)
	assert.Equal(t, "4.99", *iapPricePoint.Attributes.CustomerPrice)
	assert.Equal(t, "3.5", *iapPricePoint.Attributes.Proceeds)
}

func TestUnmarshalUnkownIncludedType(t *testing.T) {
	t.Parallel()

	payload := []byte(`{
		"type": "unknown"
	}`)

	var included Included
	err := json.Unmarshal(payload, &included)
	assert.Error(t, err)
	assert.IsType(t, ErrUnsupportedIncludedType{}, err)
}
