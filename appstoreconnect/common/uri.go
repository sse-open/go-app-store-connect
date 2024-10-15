package common

import (
	"encoding/json"
	"net/url"
)

type URI struct {
	url.URL
}

func (r URI) Cursor() string {
	return r.Query().Get("cursor")
}

func (r URI) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *URI) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	if u != nil {
		r.URL = *u
	}

	return nil
}
