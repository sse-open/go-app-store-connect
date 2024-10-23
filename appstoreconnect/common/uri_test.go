package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseURICursor(t *testing.T) {
	var uri URI
	err := uri.parseURI("https://api.appstoreconnect.apple.com/v1/apps?cursor=123")
	assert.NoError(t, err)

	uri.Scheme = "https"
	uri.Host = "api.appstoreconnect.apple.com"
	uri.Path = "/v1/apps"
	uri.RawQuery = "cursor=123"

	assert.Equal(t, "123", uri.Cursor())
}

func TestParseURIEmptyCursor(t *testing.T) {
	var uri URI
	err := uri.parseURI("https://api.appstoreconnect.apple.com/v1/apps")
	assert.NoError(t, err)

	uri.Scheme = "https"
	uri.Host = "api.appstoreconnect.apple.com"
	uri.Path = "/v1/apps"
	uri.RawQuery = ""

	assert.Equal(t, "", uri.Cursor())
}
