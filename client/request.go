package client

type AppStoreConnectRequestPayload struct {
	Data     interface{} `json:"data"`
	Included interface{} `json:"included,omitempty"`
}
