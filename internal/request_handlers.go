package internal

type RequestHandler struct {
	BaseHandler[RequestInformation]
	next Handler[RequestInformation]
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		*NewBaseHandler[RequestInformation](),
		nil,
	}
}

func (b *RequestHandler) Handle(request RequestInformation) error {
	return b.BaseHandler.Handle(request)
}
