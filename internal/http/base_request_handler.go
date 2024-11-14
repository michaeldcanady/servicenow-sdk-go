package http

import "github.com/RecoLabs/servicenow-sdk-go/internal/core"

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
	if next := b.Next(); !core.IsNil(next) {
		return next.Handle(request)
	}
	return nil
}
