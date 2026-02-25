package attachmentapi

import (
	"testing"
)

func TestBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		err      bool
	}{
		{"True", `"true"`, true, false},
		{"False", `"false"`, false, false},
		{"Invalid", `"bad"`, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Bool
			err := b.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if bool(b) != tt.expected {
				t.Errorf("got %v, expected %v", b, tt.expected)
			}
		})
	}
}
