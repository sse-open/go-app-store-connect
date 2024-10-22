# go-app-store-connect

go-app-store-connect is a Go library used for communicating with the Apple [App Store Connect API](https://developer.apple.com/documentation/appstoreconnectapi).

## Version requirement and support

The library is currently written using go 1.22, and implemented to work against the, at the time of writing, latest App Store Connect API version 3.6.
Parts of the API that at this time has been marked for deprecation has been omitted from the library.

Currently the library only supports a subset of the API's available functionality, namely the following parts:

| Endpoint (App Store Connect documentation link)                           | Further Details |
|:--------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------|
| [List Apps](https://developer.apple.com/documentation/appstoreconnectapi) | Currently supports all query parameters, but only all parts of the response payload except the included field resource payloads. |
| [List All In-App Purchases for an App](https://developer.apple.com/documentation/appstoreconnectapi/get-v1-apps-_id_-inapppurchasesv2)    | All aspects should be supported |
| [Read price information for an in-app purchase price schedule](https://developer.apple.com/documentation/appstoreconnectapi/get-v1-inapppurchasepriceschedules-_id_-manualprices) | All aspects should be supported |
| [List automatically generated prices for an in-app purchase price](https://developer.apple.com/documentation/appstoreconnectapi/get-v1-inapppurchasepriceschedules-_id_-automaticprices) | All aspects should be supported |

## Usage

The library is instantiated via creating a App Store Connect object, with a given optional `http.Client` to use for requests and a JWT token provider capable of generation bearer tokens, and then creating the relevant service exposing endpoints from that object. If a client isn't provided one is created automatically internally.

An example of this:
```go
    import (
        "errors"

        "github.com/sse-open/go-app-store-connect/appstoreconnect"
        "github.com/sse-open/go-app-store-connect/appstoreconnect/apps"
        "github.com/sse-open/go-app-store-connect/client"
    )

    ...

    jwtProvider, err := client.NewJWTProvider(keyId, issuerId, 20*time.Minute, privateKeyData)
    if err != nil {
        panic(err)
    }

    appStoreConnect, err := appstoreconnect.NewAppStoreConnect(nil, jwtProvider)
    if err != nil {
        panic(err)
    }

    appsService := appStoreConnect.AppsService()

    ctx := context.Background()
    params := apps.ListAppsQuery{
        FilterBundleID: "com.example.app",
    }
    appsResponse, clientResponse, err := appsService.ListApps(ctx, &params)
```

The JWT provider will be used to generate a token as needed by the client, using the provided Key ID, Issuer ID and Private Key all of which is given out upon creation of a new API key in App Store Connect. For more information on creating a key, and using it in communication with the api, please see [Creating API Keys for App Store Connect API
](https://developer.apple.com/documentation/appstoreconnectapi/creating-api-keys-for-app-store-connect-api) and [Generating Tokens for API Requests](https://developer.apple.com/documentation/appstoreconnectapi/generating-tokens-for-api-requests).

The returned data from the endpoint functions generally follow the same structure, in that a parsed response payload comes first followed by a client response containing information regarding the pure HTTP response, and finally a potential error.

If an error is returned it can be inspected further by converting it to a client.ErrorResponse type, making the App Store Connect API specific error information accessible. For example:
```go
    import (
        "errors"

        "github.com/sse-open/go-app-store-connect/appstoreconnect"
        "github.com/sse-open/go-app-store-connect/appstoreconnect/apps"
        "github.com/sse-open/go-app-store-connect/client"
    )

    ...

    _, _, err := appsService.ListApps(ctx, &params)

    if err != nil {
        var errorReponse client.ErrorResponse
        if errors.As(err, &errorResponse) {
            // App Store Connect API error handling
        } else if errors.Is(err, ErrRateLimitExceeded) {
            // Do something rate limit related
        } else {
            // Regular go error handling
        }
    }
```

## Code Structure

The code is separated into two main parts, the client (`client/`) containing the api client, a token provider implementation for authenticating and request/response/error structs, and the library itself (`appstoreconnect/`) implementing the different endpoints and response payloads as well as resource payloads.
The library part is separated into directories and sub directories with the intent of structuring the code similar to the hierarchy in which the api documentation is written. Each sub directory implements a service including endpoints and query parameter payloads as well response payloads.
The resource payloads have been separated out into a directory structure under `appstoreconnect/resource/`, since these model the central resources of the api and are re-used in various endpoints response payloads.

To ease in the implementation of tests, mocks of the various parts of the code is shipped with the library, and typically reside in a `mocks/` directory beside the production code equivalent.
These have been generated using [mockery v2](https://github.com/vektra/mockery), the exact version of which is specified in the *Makefile*.

## License notices
MIT License. Copyright (c) 2024 Star Stable Entertainment AB