package client

import (
	"fmt"
	"net/http"
)

// The error details that an API returns in the response body whenever the API request isn’t successful.
//
// https://developer.apple.com/documentation/appstoreconnectapi/errorresponse
type ErrorResponse struct {
	Response *http.Response       `json:"-"`
	Errors   []ErrorResponseError `json:"errors,omitempty"`
}

// The details about an error that are returned when an API request isn’t successful.
//
// https://developer.apple.com/documentation/appstoreconnectapi/errorresponse/errors-data.dictionary
type ErrorResponseError struct {
	// A machine-readable code indicating the type of error. The code is a hierarchical value with levels of
	// specificity separated by the ‘.’ character. This value is parseable for programmatic error handling in code.
	Code string `json:"code"`
	// The HTTP status code of the error. This status code usually matches the response’s status code; however, if
	// the request produces multiple errors, these two codes may differ.
	Status string `json:"status"`
	// The unique ID of a specific instance of an error, request, and response. Use this ID when providing feedback
	// to or debugging issues with Apple.
	ID *string `json:"id,omitempty"`
	// A summary of the error. Do not use this field for programmatic error handling.
	Title string `json:"title"`
	// A detailed explanation of the error. Do not use this field for programmatic error handling.
	Detail string `json:"detail"`
	// One of two possible types of values: source.Parameter, provided when a query parameter produced the error, or
	// source.JsonPointer, provided when a problem with the entity produced the error.
	//
	// https://developer.apple.com/documentation/appstoreconnectapi/jsonpointer
	// https://developer.apple.com/documentation/appstoreconnectapi/parameter
	Source *ErrorSource `json:"source,omitempty"`
	// An object that contains the error itself or associated errors. Ignored for now.
	Meta *ErrorMeta `json:"-"`
}

type ErrorSource struct {
	// A JSON pointer that indicates the location in the request entity where the error originates.
	Pointer string `json:"pointer,omitempty"`
	// The query parameter that produced the error.
	Parameter string `json:"parameter,omitempty"`
}

type ErrorMeta struct {
	Errors map[any]any `json:"-"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf(
		"%v %v: %d",
		e.Response.Request.Method,
		e.Response.Request.URL,
		e.Response.StatusCode,
	)
}
