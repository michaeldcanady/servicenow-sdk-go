package internal

type RequestHandler interface {
	Handle(RequestInformation) error
	SetNext(RequestHandler)
	Next() RequestHandler
}

type BaseHandler struct {
	next RequestHandler
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// SetNext method for BaseHandler
func (b *BaseHandler) SetNext(handler RequestHandler) {
	b.next = handler
}

func (b *BaseHandler) Next() RequestHandler {
	return b.next
}

func (b *BaseHandler) Handle(request RequestInformation) error {
	if !IsNil(b.Next()) {
		return b.Next().Handle(request)
	}
	return nil
}
