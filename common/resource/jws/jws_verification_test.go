package jws

import (
	"context"
	"crypto/x509"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/sse-open/go-app-store-connect/common/resource/certificate"
	"github.com/sse-open/go-app-store-connect/testutils"
	"github.com/stretchr/testify/assert"
)

type testClaimsPayload struct {
	jwt.RegisteredClaims
	ProductId string `json:"productId"`
}

func TestVerifyJWS(t *testing.T) {
	chain := testutils.NewTestChain(t)

	t.Run("verifies a server notification payload", func(t *testing.T) {
		signed := chain.Sign(t, &testClaimsPayload{
			ProductId: "product-id",
		}, chain.X5C(), jwt.SigningMethodES256)

		payload, err := VerifyJWS[testClaimsPayload](t.Context(), chain.RootCertificateProvider(), signed)

		assert.NoError(t, err)
		assert.Equal(t, "product-id", payload.ProductId)
	})

	t.Run("rejects a chain that does not lead to the trusted root", func(t *testing.T) {
		signed := chain.Sign(t, &testClaimsPayload{}, chain.X5C(), jwt.SigningMethodES256)

		verifier := certificate.NewRootCertificateProviderWithRootCertificateFetcher(
			func(_ context.Context) (*x509.Certificate, error) {
				return chain.UnknownRoot(), nil
			},
		)

		_, err := VerifyJWS[testClaimsPayload](t.Context(), verifier, signed)

		assert.ErrorContains(t, err, "certificate chain is not trusted")
	})

	t.Run("rejects a signing method other than ES256", func(t *testing.T) {
		signed := chain.Sign(t, &testClaimsPayload{}, chain.X5C(), jwt.SigningMethodHS256)

		_, err := VerifyJWS[testClaimsPayload](t.Context(), chain.RootCertificateProvider(), signed)

		assert.ErrorContains(t, err, "could not verify signed payload")
	})

	t.Run("rejects an x5c chain with too few certificates", func(t *testing.T) {
		signed := chain.Sign(t, &testClaimsPayload{}, chain.LeafOnlyX5C(), jwt.SigningMethodES256)

		_, err := VerifyJWS[testClaimsPayload](t.Context(), chain.RootCertificateProvider(), signed)

		assert.ErrorContains(t, err, "missing or insufficient x5c certificate chain")
	})

	t.Run("returns an error when the root certificate cannot be fetched", func(t *testing.T) {
		verifier := certificate.NewRootCertificateProviderWithRootCertificateFetcher(
			func(_ context.Context) (*x509.Certificate, error) {
				return nil, errors.New("fetch failed")
			},
		)

		_, err := VerifyJWS[testClaimsPayload](t.Context(), verifier, "does-not-matter")

		assert.ErrorContains(t, err, "could not obtain trusted root certificate")
	})
}

func TestRootCertificateCaching(t *testing.T) {
	chain := testutils.NewTestChain(t)

	t.Run("fetches the root certificate only once", func(t *testing.T) {
		calls := 0
		verifier := certificate.NewRootCertificateProviderWithRootCertificateFetcher(
			func(_ context.Context) (*x509.Certificate, error) {
				calls++
				return chain.RootCertificateProvider().RootCertificate(t.Context())
			},
		)

		signed := chain.Sign(t, &testClaimsPayload{}, chain.X5C(), jwt.SigningMethodES256)

		_, err := VerifyJWS[testClaimsPayload](t.Context(), verifier, signed)
		assert.NoError(t, err)
		_, err = VerifyJWS[testClaimsPayload](t.Context(), verifier, signed)
		assert.NoError(t, err)

		assert.Equal(t, 1, calls)
	})
}
