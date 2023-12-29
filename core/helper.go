package core

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

// Deprecated: deprecated since v{version}. Will be removed from public API
//
// ToQueryMap converts a struct to query parameter map
func ToQueryMap(source interface{}) (map[string]string, error) {
	return core.ToQueryMap(source)
}

// Deprecated: deprecated since v{version}. Will be removed from public API
func IsPointer(value interface{}) bool {
	return core.IsPointer(value)
}

// Deprecated: deprecated since v{version}. Will be removed from public API
func FromJson[T any](response *http.Response, v *T) error { //nolint:stylecheck
	if response == nil {
		return ErrNilResponse
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close() //nolint:errcheck

	if len(body) == 0 {
		return ErrNilResponseBody
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

// Deprecated: deprecated since v{version}. Will be removed from public API
//
// ParseResponse parses the HTTP Response to the provided type
func ParseResponse[T core.Response](response *http.Response, value *T) error {
	return core.ParseResponse[T](response, value)
}

// Deprecated: deprecated since v{version}. Will be removed from public API
//
// Deprecated: deprecated in version {version}. Please use `SendGet3`.
func SendGet2(requestBuilder *core.RequestBuilder, config *core.RequestConfiguration) error {
	return core.SendGet2(requestBuilder, config)
}

// Deprecated: deprecated since v{version}. Will be removed from public API
//
// SendGet3 Sends a GET request utilizing the provided RequestBuilder and RequestConfigurations
func SendGet3(requestBuilder *core.RequestBuilder, config core.RequestConfiguration2) error {
	return core.SendGet3(requestBuilder, config)
}

// Deprecated: deprecated since v{version}. Will be removed from public API
func SendPost2(requestBuilder *core.RequestBuilder, config *core.RequestConfiguration) error {
	return core.SendPost2(requestBuilder, config)
}
