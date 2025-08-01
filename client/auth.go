package client

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	errorsPkg "github.com/pkg/errors"
)

var ErrInvalidPEM = errors.New("invalid PEM format")

var ErrInvalidPrivateKey = errors.New("key could not be parsed")

var ErrInvalidKeyID = errors.New("key ID is invalid")

var ErrInvalidIssuerID = errors.New("issuer ID is invalid")

type jwtClaims struct {
	jwt.RegisteredClaims
	BundleID string `json:"bid,omitempty"` // Optional, can be used to associate the token with a specific app
}

//go:generate mockery --name IJWTProvider
type IJWTProvider interface {
	GetJWTToken() (string, error)
}

type JWTProvider struct {
	keyID          string
	issuerID       string
	bundleID       string // Optional
	expireDuration time.Duration
	privateKey     *ecdsa.PrivateKey

	token            string
	tokenGeneratedAt *time.Time
}

func (p *JWTProvider) SetBundleID(bundleID string) {
	p.bundleID = bundleID
}

func NewJWTProvider(keyID string, issuerID string, expireDuration time.Duration, privateKey []byte) (*JWTProvider, error) {
	key, err := parsePrivateKey(privateKey)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "failed to parse private key")
	}

	if keyID == "" {
		return nil, ErrInvalidKeyID
	}

	if issuerID == "" {
		return nil, ErrInvalidIssuerID
	}

	return &JWTProvider{
		keyID:          keyID,
		issuerID:       issuerID,
		privateKey:     key,
		expireDuration: expireDuration,
	}, nil
}

func parsePrivateKey(blob []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(blob)
	if block == nil {
		return nil, ErrInvalidPEM
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "not a valid PKCS8 key")
	}

	if key, ok := parsedKey.(*ecdsa.PrivateKey); ok {
		return key, nil
	}

	return nil, ErrInvalidPrivateKey
}

func (p *JWTProvider) GetJWTToken() (string, error) {
	if p.isValid() {
		return p.token, nil
	}

	issuedAt := time.Now()

	expiry := issuedAt.Add(p.expireDuration)

	claims := &jwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"appstoreconnect-v1"},
			ExpiresAt: jwt.NewNumericDate(expiry),
			Issuer:    p.issuerID,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
		},
		BundleID: p.bundleID,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t.Header["kid"] = p.keyID

	token, err := t.SignedString(p.privateKey)
	if err != nil {
		p.token = ""
		p.tokenGeneratedAt = nil
		return "", errorsPkg.Wrap(err, "failed to sign token")
	}

	p.token = token
	p.tokenGeneratedAt = &issuedAt

	return token, nil
}

func (p *JWTProvider) isValid() bool {
	if p.token != "" && p.tokenGeneratedAt != nil && time.Since(p.tokenGeneratedAt.Add(-10*time.Second)) < p.expireDuration {
		return true
	}
	return false
}
