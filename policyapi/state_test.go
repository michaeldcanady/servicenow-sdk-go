package policyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_String(t *testing.T) {
	tests := []struct {
		name string
		val  State
		want string
	}{
		{"Unknown", StateUnknown, "unknown"},
		{"Active", StateActive, "active"},
		{"Inactive", StateInactive, "inactive"},
		{"Invalid", State(999), "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.val.String())
		})
	}
}

func TestParseState(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    interface{}
		wantErr bool
	}{
		{"unknown", "unknown", StateUnknown, false},
		{"active", "active", StateActive, false},
		{"inactive", "inactive", StateInactive, false},
		{"invalid", "something-else", StateUnknown, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseState(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
