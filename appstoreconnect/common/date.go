package common

import (
	"encoding/json"
	"fmt"
	"time"
)

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format(time.DateOnly)
	return json.Marshal(formatted)
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("failed to unmarshal to string: %w", err)
	}
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}
	*d = Date(t)
	return nil
}
