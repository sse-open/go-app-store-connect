package certificate

import (
	"context"
	"crypto/x509"
	"io"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const appleRootCertificateURL = "https://www.apple.com/certificateauthority/AppleRootCA-G3.cer"

type RootCertificateFetcher func(ctx context.Context) (*x509.Certificate, error)

type RootCertificateProvider struct {
	fetchRootCertificate RootCertificateFetcher

	mu         sync.Mutex
	cachedRoot *x509.Certificate
}

func NewRootCertificateProvider() *RootCertificateProvider {
	return &RootCertificateProvider{
		fetchRootCertificate: fetchAppleRootCertificate,
	}
}

func NewRootCertificateProviderWithRootCertificateFetcher(
	fetcher RootCertificateFetcher,
) *RootCertificateProvider {
	return &RootCertificateProvider{
		fetchRootCertificate: fetcher,
	}
}

func fetchAppleRootCertificate(ctx context.Context) (*x509.Certificate, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, appleRootCertificateURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request for Apple root certificate")
	}

	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch Apple root certificate")
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log := slog.Default()
			log.Error("could not close Apple root certificate response body", "error", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected status code %d fetching Apple root certificate", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read Apple root certificate response")
	}

	certificate, err := x509.ParseCertificate(body)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse Apple root certificate")
	}

	return certificate, nil
}

func (v *RootCertificateProvider) RootCertificate(ctx context.Context) (*x509.Certificate, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if v.cachedRoot != nil {
		return v.cachedRoot, nil
	}

	certificate, err := v.fetchRootCertificate(ctx)
	if err != nil {
		return nil, err
	}

	v.cachedRoot = certificate
	return certificate, nil
}
