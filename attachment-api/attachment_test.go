package attachmentapi

import (
	"encoding/json"
	"testing"
)

func TestAttachment_UnmarshalJSON(t *testing.T) {
	rawJSON := []byte(`{
		"table_sys_id": "sid",
		"size_bytes": "462",
		"compressed": "true",
		"sys_updated_on": "2009-05-21 04:12:21"
	}`)

	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"Basic", rawJSON, "sid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a Attachment
			if err := json.Unmarshal(tt.input, &a); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if a.TableSysId != tt.expected {
				t.Errorf("got %s, expected %s", a.TableSysId, tt.expected)
			}
		})
	}
}
