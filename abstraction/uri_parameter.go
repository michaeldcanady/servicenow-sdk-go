package abstraction

import (
	"reflect"
	"strconv"
	"strings"
)

const tagValue = "uriparametername"

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "uriParameter: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "uriParameter: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "uriParameter: Unmarshal(nil " + e.Type.String() + ")"
}

type QueryParameter struct {
	Name  string
	Value string
}

// uriParamValue generates a map of query parameters from a struct using reflection.
// It takes an input 'source' (a pointer to a struct), and returns a map of string key-value pairs
// representing query parameters and an error if the 'source' is invalid or unmarshalable.
func uriParamValue(source interface{}) (map[string]string, error) {
	valOfP := reflect.ValueOf(source)

	// Validate the input source
	if valOfP.Kind() != reflect.Ptr || valOfP.IsNil() {
		return nil, &InvalidUnmarshalError{reflect.TypeOf(source)}
	}

	queryParameters := make(map[string]string)

	structType := valOfP.Elem().Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldName := getFieldName(field)

		fieldValue := valOfP.Elem().Field(i).Interface()
		if fieldValue == nil {
			continue
		}

		queryParameters[fieldName] = asString(fieldValue)
	}

	return queryParameters, nil
}

// getFieldName extracts the field name from a struct field, considering struct tags.
// It takes a 'field' of type reflect.StructField and returns the field name as a string.
func getFieldName(field reflect.StructField) string {
	fieldName := field.Name
	tagValue := field.Tag.Get("uriparametername")
	if tagValue != "" {
		fieldName = tagValue
	}
	return fieldName
}

// asString converts a value of various types to a string representation.
// It handles types like *string, *bool, int32, and []string and returns their string representation.
// If the value is not a valid type, it returns an empty string.
func asString(value interface{}) string {

	returnValue := ""

	switch v := value.(type) {
	case *string:
		if v != nil && *v != "" {
			returnValue = *v
		}
	case *bool:
		if v != nil {
			returnValue = strconv.FormatBool(*v)
		}
	case int32:
		if v != 0 {
			returnValue = strconv.FormatInt(int64(v), 10)
		}
	case []string:
		if len(v) > 0 {
			returnValue = strings.Join(v, ",")
		}
	}
	return returnValue
}
