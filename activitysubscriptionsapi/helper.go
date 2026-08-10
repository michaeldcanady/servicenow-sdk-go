package actsubapi

import (
	"reflect"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

// isNilOrEmpty reports whether value is nil or empty.
//
// A value is considered empty when it is a zero-length string, slice, map, array
// or channel, or when it is the zero value for its type (0, false, an empty
// struct, ...). Pointers and interfaces are followed, so a pointer to an empty
// string is empty, while a pointer to a non-empty string is not.
func isNilOrEmpty(value interface{}) bool {
	if conversion.IsNil(value) {
		return true
	}

	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return true
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		return v.Len() == 0
	default:
		return v.IsZero()
	}
}
