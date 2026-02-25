package attachmentapi

import (
	"testing"
)

func TestInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		err      bool
	}{
		{"Ok", `"1"`, 1, false},
		{"Empty", `""`, 0, false},
		{"Invalid", `"s"`, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i Int
			err := i.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if int(i) != tt.expected {
				t.Errorf("got %v, expected %v", i, tt.expected)
			}
		})
	}
}
