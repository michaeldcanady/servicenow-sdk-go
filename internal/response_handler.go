package internal

import "net/http"

type ResponseHandler struct {
	BaseHandler[*http.Response]
	next Handler[*http.Response]
}

func NewResponseHandler() *ResponseHandler {
	return &ResponseHandler{
		*NewBaseHandler[*http.Response](),
		nil,
	}
}

func (b *ResponseHandler) Handle(request *http.Response) error {
	return b.BaseHandler.Handle(request)
}
