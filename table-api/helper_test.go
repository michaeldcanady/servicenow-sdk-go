package tableapi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertType(t *testing.T) {
	tests := []test[int]{
		{
			Title:    "Valid",
			Input:    42,
			Expected: 42,
		},
		{
			Title:    "NotAnInt",
			Input:    "not_an_int",
			Expected: 0,
			Error:    fmt.Errorf("value (%v) cannot be converted to %T", "not_an_int", 42),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			page, err := convertType[int](tt.Input)

			assert.Equal(t, err, tt.Error)
			assert.Equal(t, tt.Expected, page)
		})
	}
}

func TestCovertToPage(t *testing.T) {
	tests := []test[PageResult]{
		{
			Title:    "NilResponse",
			Input:    nil,
			Expected: PageResult{},
			Error:    ErrNilResponse,
		},
		{
			Title:    "WrongResponseType",
			Input:    TableItemResponse{},
			Expected: PageResult{},
			Error:    ErrWrongResponseType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			page, err := convertToPage(tt.Input)

			assert.Equal(t, err, tt.Error)
			assert.Equal(t, tt.Expected, page)
		})
	}
}

func TestConvertFromTableEntry(t *testing.T) {
	tests := []test[map[string]string]{
		{
			Title:    "Test with map[string]string",
			Input:    map[string]string{"key": "value"},
			Expected: map[string]string{"key": "value"},
		},
		{
			Title:    "Test with TableEntry",
			Input:    TableEntry{"key": 123},
			Expected: map[string]string{"key": "123"},
		},
		{
			Title:     "Test with unsupported type",
			Input:     123,
			expectErr: true,
		},
		{
			Title:    "Test with pointer to map[string]string",
			Input:    &map[string]string{"key": "value"},
			Expected: map[string]string{"key": "value"},
		},
		{
			Title:    "Test with pointer to TableEntry int",
			Input:    &TableEntry{"key": 123},
			Expected: map[string]string{"key": "123"},
		},
		{
			Title:    "Test with pointer to TableEntry bool",
			Input:    &TableEntry{"key": true},
			Expected: map[string]string{"key": "true"},
		},
		{
			Title:    "Test with pointer to TableEntry string",
			Input:    &TableEntry{"key": "value"},
			Expected: map[string]string{"key": "value"},
		},
		{
			Title:    "Test with pointer to TableEntry float",
			Input:    &TableEntry{"key": 1.2},
			Expected: map[string]string{"key": "1.2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			got, err := convertFromTableEntry(tt.Input)
			if (err != nil) != tt.expectErr {
				t.Errorf("convertFromTableEntry() error = %v, wantErr %v", err, tt.expectErr)
				return
			}
			assert.Equal(t, tt.Expected, got)
		})
	}
}
