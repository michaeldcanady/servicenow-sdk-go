package tableapi

import (
	"reflect"
	"testing"
)

func TestConvertType(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
		err      bool
	}{
		{"Ok", 42, 42, false},
		{"Bad", "s", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := convertType[int](tt.input)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	var s *string
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"Nil", nil, true},
		{"NilPtr", s, true},
		{"NotNil", "v", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isNil(tt.input) != tt.expected {
				t.Errorf("failed for %v", tt.input)
			}
		})
	}
}

func TestConvertToPage(t *testing.T) {
	resp := &TableCollectionResponse{
		Result: []*TableEntry{{}},
	}
	tests := []struct {
		name  string
		input interface{}
		err   error
	}{
		{"Ok", resp, nil},
		{"OkPtr", &resp, nil}, // Pointer to pointer to collection response
		{"Nil", nil, ErrNilResponse},
		{"WrongType", 123, ErrWrongResponseType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := convertToPage(tt.input)
			if tt.err != nil {
				if err != tt.err {
					t.Errorf("got err %v, expected %v", err, tt.err)
				}
			} else if err != nil {
				t.Errorf("unexpected err %v", err)
			}
		})
	}
}

func TestConvertFromTableEntry(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]string
		err      bool
	}{
		{"Map", map[string]string{"a": "b"}, map[string]string{"a": "b"}, false},
		{"Entry", TableEntry{"a": 1, "b": "s"}, map[string]string{"a": "1", "b": "s"}, false},
		{"PtrEntry", &TableEntry{"a": 1}, map[string]string{"a": "1"}, false},
		{"BadType", 123, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := convertFromTableEntry(tt.input)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if !tt.err && !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}
