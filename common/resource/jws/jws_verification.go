package jws

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/sse-open/go-app-store-connect/common/resource/certificate"
)

type Claims[P any] interface {
	*P
	jwt.Claims
}

func VerifyJWS[P any, C Claims[P]](
	ctx context.Context,
	v *certificate.RootCertificateProvider,
	signedPayload string,
) (P, error) {
	var empty P

	root, err := v.RootCertificate(ctx)
	if err != nil {
		return empty, errors.Wrap(err, "could not obtain trusted root certificate")
	}

	claims := C(new(P))
	_, err = jwt.ParseWithClaims(
		signedPayload,
		claims,
		func(token *jwt.Token) (any, error) {
			return leafPublicKeyFromChain(token, root)
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodES256.Alg()}),
		jwt.WithPaddingAllowed(),
	)
	if err != nil {
		return empty, errors.Wrap(err, "could not verify signed payload")
	}

	return *claims, nil
}

func leafPublicKeyFromChain(token *jwt.Token, root *x509.Certificate) (*ecdsa.PublicKey, error) {
	rawChain, ok := token.Header["x5c"].([]any)
	if !ok || len(rawChain) < 2 {
		return nil, errors.New("missing or insufficient x5c certificate chain")
	}

	certificates := make([]*x509.Certificate, 0, len(rawChain))
	for _, rawCert := range rawChain {
		certString, ok := rawCert.(string)
		if !ok {
			return nil, errors.New("x5c entry was not a string")
		}
		der, err := base64.StdEncoding.DecodeString(certString)
		if err != nil {
			return nil, errors.Wrap(err, "could not decode x5c certificate")
		}
		certificate, err := x509.ParseCertificate(der)
		if err != nil {
			return nil, errors.Wrap(err, "could not parse x5c certificate")
		}
		certificates = append(certificates, certificate)
	}

	leaf := certificates[0]
	intermediates := x509.NewCertPool()
	for _, intermediate := range certificates[1:] {
		intermediates.AddCert(intermediate)
	}

	roots := x509.NewCertPool()
	roots.AddCert(root)

	if _, err := leaf.Verify(x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermediates,
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	}); err != nil {
		return nil, errors.Wrap(err, "certificate chain is not trusted")
	}

	publicKey, ok := leaf.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("leaf certificate does not have an ECDSA public key")
	}

	return publicKey, nil
}
