package attachmentapi

import (
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		err      bool
	}{
		{"Ok", `"2006-01-02 15:04:05"`, time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC), false},
		{"BadFormat", `"bad"`, time.Time{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ti Time
			err := ti.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if !tt.err && !time.Time(ti).Equal(tt.expected) {
				t.Errorf("got %v, expected %v", ti, tt.expected)
			}
		})
	}
}
