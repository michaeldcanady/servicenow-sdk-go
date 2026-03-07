package policyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputStatus_String(t *testing.T) {
	tests := []struct {
		name string
		val  InputStatus
		want string
	}{
		{"Unknown", InputStatusUnknown, "unknown"},
		{"Invalid", InputStatusInvalid, "invalid"},
		{"Valid", InputStatusValid, "valid"},
		{"OutOfBounds", InputStatus(999), "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.val.String())
		})
	}
}

func TestParseInputStatus(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    interface{}
		wantErr bool
	}{
		{"unknown", "unknown", InputStatusUnknown, false},
		{"invalid", "invalid", InputStatusInvalid, false},
		{"valid", "valid", InputStatusValid, false},
		{"error", "something-else", InputStatusUnknown, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInputStatus(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
