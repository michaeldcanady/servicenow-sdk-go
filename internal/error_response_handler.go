package internal

import (
	"net/http"
)

type ErrorResponseHandler struct {
	BaseHandler[*http.Response]
	ErrorFactory *ErrorFactory
}

func NewErrorResponseHandler(errorFactory *ErrorFactory) *ErrorResponseHandler {
	return &ErrorResponseHandler{
		*NewBaseHandler[*http.Response](),
		errorFactory,
	}
}

func (e *ErrorResponseHandler) Handle(response *http.Response) error {
	if response.StatusCode < 400 {
		return nil
	}

	return e.ErrorFactory.CreateError(response)
}
