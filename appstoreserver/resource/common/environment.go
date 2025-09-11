package common

// The server environment, either sandbox or production.
//
// https://developer.apple.com/documentation/appstoreserverapi/environment
type Environment string

var (
	EnvironmentSandbox    Environment = "Sandbox"
	EnvironmentProduction Environment = "Production"
)
