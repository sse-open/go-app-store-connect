package resource

import (
	"context"
	"crypto/x509"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sse-open/go-app-store-connect/common/resource/certificate"
	"github.com/sse-open/go-app-store-connect/testutils"
	"github.com/stretchr/testify/assert"
)

func TestVerifyJWS(t *testing.T) {
	chain := testutils.NewTestChain(t)

	t.Run("verifies a server notification payload", func(t *testing.T) {
		signed := chain.Sign(t, &ServerNotificationV2PayloadWithClaims{
			ServerNotificationResponseBodyV2DecodedPayload: ServerNotificationResponseBodyV2DecodedPayload{
				NotificationType: NotificationTypeSubscribed,
				NotificationUUID: "notification-uuid",
				Version:          "2.0",
			},
		}, chain.X5C(), jwt.SigningMethodES256)

		serverNotification := ServerNotificationResponseBodyV2{
			SignedPayload: signed,
		}

		payload, err := serverNotification.VerifySignedPayloadClaims(t.Context(), chain.RootCertificateProvider())

		assert.NoError(t, err)
		assert.Equal(t, NotificationTypeSubscribed, payload.NotificationType)
		assert.Equal(t, "notification-uuid", payload.NotificationUUID)
		assert.Equal(t, "2.0", payload.Version)
	})

	t.Run("propagates verification errors", func(t *testing.T) {
		signed := chain.Sign(t, &ServerNotificationV2PayloadWithClaims{
			ServerNotificationResponseBodyV2DecodedPayload: ServerNotificationResponseBodyV2DecodedPayload{
				NotificationType: NotificationTypeSubscribed,
				NotificationUUID: "notification-uuid",
				Version:          "2.0",
			},
		}, chain.X5C(), jwt.SigningMethodES256)

		verifier := certificate.NewRootCertificateProviderWithRootCertificateFetcher(
			func(_ context.Context) (*x509.Certificate, error) {
				return chain.UnknownRoot(), nil
			},
		)

		serverNotification := ServerNotificationResponseBodyV2{
			SignedPayload: signed,
		}

		payload, err := serverNotification.VerifySignedPayloadClaims(t.Context(), verifier)

		assert.ErrorContains(t, err, "certificate chain is not trusted")
		assert.Nil(t, payload)
	})
}
