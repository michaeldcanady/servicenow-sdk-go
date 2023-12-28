package core

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response interface {
	ParseHeaders(headers http.Header)
}

// ParseResponse parses the HTTP Response to the provided type
func ParseResponse[T Response](response *http.Response, value T) error {
	err := FromJSON(response, value)
	if err != nil {
		return err
	}

	value.ParseHeaders(response.Header)

	return nil
}

func FromJSON[T any](response *http.Response, v T) error {
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
