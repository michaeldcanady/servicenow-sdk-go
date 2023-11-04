package tableapi

import (
	"testing"
)

func TestConvertType(t *testing.T) {
	// Test case with valid conversion
	var validValue interface{} = 42
	expectedInt := 42
	validInt, validErr := convertType[int](validValue)
	if validErr != nil {
		t.Errorf("Expected no error, got: %v", validErr)
	}
	if validInt != expectedInt {
		t.Errorf("Expected %d, got: %d", expectedInt, validInt)
	}

	// Test case with invalid conversion
	invalidValue := "not_an_int"
	var invalidInt int
	invalidInt, invalidErr := convertType[int](invalidValue)
	if invalidErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if invalidInt != 0 {
		t.Errorf("Expected default value 0, got: %d", invalidInt)
	}
}
