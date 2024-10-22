package apps

import (
	"context"
	"errors"
	"testing"

	"github.com/h2non/gock"
	"github.com/nbio/st"
	"github.com/sse-open/go-app-store-connect/client"
	"github.com/sse-open/go-app-store-connect/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListApps(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	gock.New("https://api.appstoreconnect.apple.com").
		Get("/v1/apps").
		MatchParam("filter[bundleId]", "com.example.app").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(200).
		JSON(map[string]any{
			"data": []any{
				map[string]any{
					"id": "123456",
					"attributes": map[string]any{
						"bundleId": "com.example.app",
						"name":     "An Example App",
					},
					"type": "apps",
				},
			},
		})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := client.NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	appsService := NewAppsService(c)

	queryParams := ListAppsQuery{
		FilterBundleID: "com.example.app",
	}
	responsePayload, resp, err := appsService.ListApps(ctx, &queryParams)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, responsePayload)

	assert.Equal(t, 1, len(responsePayload.Data))
	assert.Equal(t, "123456", responsePayload.Data[0].ID)
	assert.Equal(t, "com.example.app", *responsePayload.Data[0].Attributes.BundleID)
	assert.Equal(t, "An Example App", *responsePayload.Data[0].Attributes.Name)
	assert.Equal(t, "apps", responsePayload.Data[0].Type)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}

func TestListAppsBadRequest(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	ctx := context.Background()

	gock.New("https://api.appstoreconnect.apple.com").
		Get("/v1/apps").
		MatchParam("filter[bundleId]", "com.example.app").
		MatchHeader("Authorization", "Bearer fakeToken").
		Reply(400).
		JSON(map[string]any{
			"errors": []any{
				map[string]any{
					"id":     "123456789",
					"status": "400",
					"code":   "BAD_REQUEST",
					"title":  "Invalid bundleId.",
					"detail": "The provided bundleId is invalid.",
					"meta": map[string]any{
						"associatedErrors": map[string]any{
							"/v1/apps": []any{
								map[string]any{
									"id":     "987654321",
									"status": "400",
									"code":   "BAD_REQUEST",
									"title":  "Bad Request",
									"detail": "Invalid input for field filter[bundleId]",
								},
							},
						},
					},
				},
			},
		})

	mockedJWTProvider := mocks.NewIJWTProvider(t)
	mockedJWTProvider.EXPECT().GetJWTToken().Return("fakeToken", nil)

	c, err := client.NewClient(nil, mockedJWTProvider)
	assert.NoError(t, err)

	appsService := NewAppsService(c)

	queryParams := ListAppsQuery{
		FilterBundleID: "com.example.app",
	}
	responsePayload, resp, err := appsService.ListApps(ctx, &queryParams)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, responsePayload)

	var errorResponse client.ErrorResponse
	assert.True(t, errors.As(err, &errorResponse))
	assert.Equal(t, 400, errorResponse.Response.StatusCode)
	assert.Equal(t, "123456789", *errorResponse.Errors[0].ID)
	assert.Equal(t, "400", errorResponse.Errors[0].Status)
	assert.Equal(t, "BAD_REQUEST", errorResponse.Errors[0].Code)
	assert.Equal(t, "Invalid bundleId.", errorResponse.Errors[0].Title)
	assert.Equal(t, "The provided bundleId is invalid.", errorResponse.Errors[0].Detail)

	mockedJWTProvider.AssertExpectations(t)
	st.Expect(t, gock.IsDone(), true)
}
