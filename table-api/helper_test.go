package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestConvertType(t *testing.T) {
	// Test case with valid conversion
	var validValue interface{} = 42
	expectedInt := 42
	validInt, validErr := internal.ConvertType[int](validValue)
	assert.Nil(t, validErr)
	assert.Equal(t, expectedInt, validInt)

	// Test case with invalid conversion
	invalidValue := "not_an_int"
	var invalidInt int
	invalidInt, invalidErr := internal.ConvertType[int](invalidValue)
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
