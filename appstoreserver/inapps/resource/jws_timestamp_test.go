package resource

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWSTimestampMarshal(t *testing.T) {
	timestamp := JWSTimestamp{time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)}
	expected := `1696161600`

	data, err := json.Marshal(timestamp)
	require.NoError(t, err)

	assert.Equal(t, expected, string(data))
}

func TestJWSTimestampUnmarshal(t *testing.T) {
	timestamp := JWSTimestamp{time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)}
	expected := `1696161600`
	data := []byte(expected)

	var unmarshalled JWSTimestamp
	err := json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)

	assert.Equal(t, timestamp.Unix(), unmarshalled.Unix())
}
