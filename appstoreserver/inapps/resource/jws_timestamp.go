package resource

import (
	"encoding/json"
	"strconv"
	"time"
)

type JWSTimestamp struct {
	time.Time
}

func (t JWSTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t.Time).Unix(), 10)), nil
}

func (t *JWSTimestamp) UnmarshalJSON(data []byte) error {
	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}
	*t = JWSTimestamp{time.Unix(timestamp, 0)}
	return nil
}
