package testutils

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"math/big"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sse-open/go-app-store-connect/common/resource/certificate"
	"github.com/stretchr/testify/assert"
)

type TestChain struct {
	root        *x509.Certificate
	leafKey     *ecdsa.PrivateKey
	x5c         []string
	leafOnlyX5C []string
	unknownRoot *x509.Certificate
}

func newCertificate(
	t *testing.T,
	template *x509.Certificate,
	parent *x509.Certificate,
	parentKey *ecdsa.PrivateKey,
) (*x509.Certificate, *ecdsa.PrivateKey) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)

	signingParent := parent
	signingKey := parentKey
	if signingParent == nil {
		signingParent = template
		signingKey = key
	}

	der, err := x509.CreateCertificate(rand.Reader, template, signingParent, &key.PublicKey, signingKey)
	assert.NoError(t, err)

	certificate, err := x509.ParseCertificate(der)
	assert.NoError(t, err)

	return certificate, key
}

func caTemplate(serial int64, name string) *x509.Certificate {
	return &x509.Certificate{
		SerialNumber:          big.NewInt(serial),
		Subject:               pkix.Name{CommonName: name},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
}

func leafTemplate(serial int64, name string) *x509.Certificate {
	return &x509.Certificate{
		SerialNumber:          big.NewInt(serial),
		Subject:               pkix.Name{CommonName: name},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}
}

func NewTestChain(t *testing.T) TestChain {
	t.Helper()

	root, rootKey := newCertificate(t, caTemplate(1, "Test Root"), nil, nil)
	intermediate, intermediateKey := newCertificate(t, caTemplate(2, "Test Intermediate"), root, rootKey)
	leaf, leafKey := newCertificate(t, leafTemplate(3, "Test Leaf"), intermediate, intermediateKey)

	otherRoot, _ := newCertificate(t, caTemplate(4, "Other Root"), nil, nil)

	encode := func(certificates ...*x509.Certificate) []string {
		encoded := make([]string, 0, len(certificates))
		for _, certificate := range certificates {
			encoded = append(encoded, base64.StdEncoding.EncodeToString(certificate.Raw))
		}
		return encoded
	}

	return TestChain{
		root:        root,
		leafKey:     leafKey,
		x5c:         encode(leaf, intermediate, root),
		leafOnlyX5C: encode(leaf),
		unknownRoot: otherRoot,
	}
}

func (c TestChain) Sign(t *testing.T, claims jwt.Claims, x5c []string, method jwt.SigningMethod) string {
	t.Helper()

	token := jwt.NewWithClaims(method, claims)
	token.Header["x5c"] = x5c

	var (
		signed string
		err    error
	)
	if method == jwt.SigningMethodES256 {
		signed, err = token.SignedString(c.leafKey)
	} else {
		signed, err = token.SignedString([]byte("secret"))
	}
	assert.NoError(t, err)

	return signed
}

func (c TestChain) RootCertificateProvider() *certificate.RootCertificateProvider {
	return certificate.NewRootCertificateProviderWithRootCertificateFetcher(
		func(_ context.Context) (*x509.Certificate, error) {
			return c.root, nil
		},
	)
}

func (c TestChain) X5C() []string {
	return c.x5c
}

func (c TestChain) LeafOnlyX5C() []string {
	return c.leafOnlyX5C
}

func (c TestChain) UnknownRoot() *x509.Certificate {
	return c.unknownRoot
}
