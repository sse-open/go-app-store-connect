package common

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimestampMarshal(t *testing.T) {
	timestamp := Timestamp{time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)}
	expected := `1696161600`

	data, err := json.Marshal(timestamp)
	require.NoError(t, err)

	assert.Equal(t, expected, string(data))
}

func TestTimestampUnmarshal(t *testing.T) {
	timestamp := Timestamp{time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)}
	expected := `1696161600`
	data := []byte(expected)

	var unmarshalled Timestamp
	err := json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)

	assert.Equal(t, timestamp.Unix(), unmarshalled.Unix())
}
