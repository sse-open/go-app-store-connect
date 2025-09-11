package common

import (
	"encoding/json"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t.Time).UnixMilli(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}
	*t = Timestamp{time.UnixMilli(timestamp).UTC()}
	return nil
}
