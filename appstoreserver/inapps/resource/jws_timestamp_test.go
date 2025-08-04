package resource

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJWSTimestampMarshal(t *testing.T) {
	timestamp := JWSTimestamp(time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC))
	expected := `1696161600`
	data, err := json.Marshal(timestamp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(data) != expected {
		t.Errorf("expected %s, got %s", expected, data)
	}
}

func TestJWSTimestampUnmarshal(t *testing.T) {
	timestamp := JWSTimestamp(time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC))
	expected := `1696161600`
	data := []byte(expected)

	var unmarshalled JWSTimestamp
	err := json.Unmarshal(data, &unmarshalled)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if time.Time(unmarshalled).Unix() != time.Time(timestamp).Unix() {
		t.Errorf("expected %d, got %d", time.Time(timestamp).Unix(), time.Time(unmarshalled).Unix())
	}
}
