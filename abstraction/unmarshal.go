package abstraction

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
)

func IsPointer(value interface{}) bool {

	if value == nil {
		return false
	}

	valueKind := reflect.ValueOf(value).Kind()

	return valueKind == reflect.Ptr
}

func FromJson[T interface{}](response *http.Response) (*T, error) {

	if !IsPointer(response) {
		return nil, errors.New("response is nil or not a pointer")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var value T
	if err := json.Unmarshal(body, &value); err != nil {
		return nil, err
	}

	return &value, nil
}
