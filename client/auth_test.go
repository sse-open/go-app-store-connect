package client

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewJWTProvider(t *testing.T) {
	t.Parallel()

	// Manually generated key for testing purposes
	var privPEMData = []byte(`
-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgOn2DXPVis6YGyaw6
6eA3FypZ2yVEwaqY1bBl2xeG0o6hRANCAATeLhoTgI7CFp21U2pgC7f0k1Yf8hBO
jsI9lUtwZOpeuYJga3p442cnJbNrQaOsHEAI1cdESnOJO1/m4ZTh3aGS
-----END PRIVATE KEY-----
`)

	jwtProvider, err := NewJWTProvider("fakeKeyID", "fakeIssuerID", 5*time.Minute, privPEMData)
	assert.NoError(t, err)

	token, err := jwtProvider.GetJWTToken()
	assert.NoError(t, err)

	parsed, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			if token.Header["kid"] != "fakeKeyID" {
				return nil, errors.New("invalid key id header")
			}
			if token.Header["alg"] != jwt.SigningMethodES256.Alg() {
				return nil, errors.New("invalid signing method")
			}
			return &jwtProvider.privateKey.PublicKey, nil
		},
		jwt.WithAudience("appstoreconnect-v1"),
		jwt.WithIssuer("fakeIssuerID"),
		jwt.WithValidMethods([]string{jwt.SigningMethodES256.Alg()}),
	)
	assert.NoError(t, err)
	assert.True(t, parsed.Valid)

	tokenCached, err := jwtProvider.GetJWTToken()
	assert.NoError(t, err)
	assert.Equal(t, token, tokenCached)
}

func TestNewJWTProviderInvalidPEM(t *testing.T) {
	t.Parallel()

	_, err := NewJWTProvider("fakeKeyID", "fakeIssuerID", 20*time.Minute, []byte("an invalid pem"))
	assert.Error(t, err, "invalid PEM format")
}

func TestNewJWTProviderInvalidKeyID(t *testing.T) {
	t.Parallel()

	// Manually generated key for testing purposes
	var privPEMData = []byte(`
-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgOn2DXPVis6YGyaw6
6eA3FypZ2yVEwaqY1bBl2xeG0o6hRANCAATeLhoTgI7CFp21U2pgC7f0k1Yf8hBO
jsI9lUtwZOpeuYJga3p442cnJbNrQaOsHEAI1cdESnOJO1/m4ZTh3aGS
-----END PRIVATE KEY-----
`)
	_, err := NewJWTProvider("", "fakeIssuerID", 20*time.Minute, privPEMData)
	assert.Error(t, err, "key ID is invalid")
}

func TestNewJWTProviderInvalidIssuerID(t *testing.T) {
	t.Parallel()

	// Manually generated key for testing purposes
	var privPEMData = []byte(`
-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgOn2DXPVis6YGyaw6
6eA3FypZ2yVEwaqY1bBl2xeG0o6hRANCAATeLhoTgI7CFp21U2pgC7f0k1Yf8hBO
jsI9lUtwZOpeuYJga3p442cnJbNrQaOsHEAI1cdESnOJO1/m4ZTh3aGS
-----END PRIVATE KEY-----
`)
	_, err := NewJWTProvider("fakeKeyID", "", 20*time.Minute, privPEMData)
	assert.Error(t, err, "issuer ID is invalid")
}
