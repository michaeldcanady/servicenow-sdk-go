package tableapi

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/RecoLabs/servicenow-sdk-go/internal"
)

func convertType[T any](val interface{}) (T, error) {
	v, ok := val.(T)
	if !ok {
		return v, fmt.Errorf("value (%v) cannot be converted to %T", val, v)
	}
	return v, nil
}

// convertToPage converts a response into a PageResult.
func convertToPage(response interface{}) (PageResult, error) {
	var page PageResult

	if internal.IsNil(response) {
		return page, ErrNilResponse
	}

	valType := reflect.ValueOf(response)

	if valType.Kind() == reflect.Pointer {
		response = reflect.Indirect(valType.Elem()).Interface()
	}

	collectionRep, ok := response.(TableCollectionResponse)
	if !ok {
		return page, ErrWrongResponseType
	}

	page.Result = collectionRep.Result
	page.FirstPageLink = collectionRep.FirstPageLink
	page.LastPageLink = collectionRep.LastPageLink
	page.NextPageLink = collectionRep.NextPageLink
	page.PreviousPageLink = collectionRep.PreviousPageLink

	return page, nil
}

func convertFromTableEntry(entry interface{}) (map[string]string, error) {
	retVal := map[string]string{}

	if internal.IsNil(entry) {
		return nil, errors.New("entry is nil")
	}

	// Check if entry is a pointer
	val := reflect.ValueOf(entry)
	if val.Kind() == reflect.Ptr {
		// Dereference the pointer and call the function again
		return convertFromTableEntry(val.Elem().Interface())
	}

	switch v := entry.(type) {
	case map[string]string:
		retVal = v
	case TableEntry:
		for key, value := range v {
			switch value := value.(type) {
			case int:
				retVal[key] = fmt.Sprintf("%d", value)
			default:
				retVal[key] = fmt.Sprintf("%v", value)
			}
		}
	default:
		return nil, fmt.Errorf("expected (%T) or (%T), not (%T)", map[string]string{}, TableEntry{}, entry)
	}

	return retVal, nil
}
