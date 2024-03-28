package internal

import "net/http"

type ErrorFactory struct {
	mapping map[string]interface{}
}

func NewErrorFactory() *ErrorFactory {
	return &ErrorFactory{}
}

func (f *ErrorFactory) CreateError(response *http.Response) error {
	return nil
}

func (f *ErrorFactory) RegisterSerializer(statusCode string, serializer interface{}) {
	f.mapping[statusCode] = serializer
}
