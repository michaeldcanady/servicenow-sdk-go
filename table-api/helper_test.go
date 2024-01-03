package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertType(t *testing.T) {
	// Test case with valid conversion
	var validValue interface{} = 42
	expectedInt := 42
	validInt, validErr := convertType[int](validValue)
	assert.Nil(t, validErr)
	assert.Equal(t, expectedInt, validInt)

	// Test case with invalid conversion
	invalidValue := "not_an_int"
	var invalidInt int
	invalidInt, invalidErr := convertType[int](invalidValue)
	assert.Error(t, invalidErr)
	assert.Equal(t, 0, invalidInt)
}

func TestConvertToPageNilResponse(t *testing.T) {
	page, err := convertToPage(nil)

	assert.Equal(t, PageResult{}, page)
	assert.ErrorIs(t, err, ErrNilResponse)
}

func TestConvertToPageWrongResponseType(t *testing.T) {
	response := TableItemResponse{}

	page, err := convertToPage(response)

	assert.Equal(t, PageResult{}, page)
	assert.ErrorIs(t, err, ErrWrongResponseType)
}

func TestConvertFromTableEntry(t *testing.T) {
	tests := []test[map[string]string]{
		{
			title:    "Test with map[string]string",
			value:    map[string]string{"key": "value"},
			expected: map[string]string{"key": "value"},
		},
		{
			title:    "Test with TableEntry",
			value:    TableEntry{"key": 123},
			expected: map[string]string{"key": "123"},
		},
		{
			title:     "Test with unsupported type",
			value:     123,
			expectErr: true,
		},
		{
			title:    "Test with pointer to map[string]string",
			value:    &map[string]string{"key": "value"},
			expected: map[string]string{"key": "value"},
		},
		{
			title:    "Test with pointer to TableEntry int",
			value:    &TableEntry{"key": 123},
			expected: map[string]string{"key": "123"},
		},
		{
			title:    "Test with pointer to TableEntry bool",
			value:    &TableEntry{"key": true},
			expected: map[string]string{"key": "true"},
		},
		{
			title:    "Test with pointer to TableEntry string",
			value:    &TableEntry{"key": "value"},
			expected: map[string]string{"key": "value"},
		},
		{
			title:    "Test with pointer to TableEntry float",
			value:    &TableEntry{"key": 1.2},
			expected: map[string]string{"key": "1.2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got, err := convertFromTableEntry(tt.value)
			if (err != nil) != tt.expectErr {
				t.Errorf("convertFromTableEntry() error = %v, wantErr %v", err, tt.expectErr)
				return
			}
			assert.Equal(t, tt.expected, got)
		})
	}
}
