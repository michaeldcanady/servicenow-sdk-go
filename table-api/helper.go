package tableapi

import (
	"reflect"
)

// isNil checks if a value is nil or a nil interface
func isNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// convertToPage converts a response into a PageResult.
func convertToPage(response interface{}) (PageResult, error) {
	var page PageResult

	if isNil(response) {
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
