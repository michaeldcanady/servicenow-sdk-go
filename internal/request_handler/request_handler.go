package requesthandler

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

type RequestHandler interface {
	Handle(core.RequestInformation) error
	GetNext() RequestHandler
	SetNext(RequestHandler)
}

type requestHandler struct {
	next RequestHandler
}

func NewRequestHandler() RequestHandler {
	return &requestHandler{}
}

func (rH *requestHandler) Handle(req core.RequestInformation) error {
	if nextHandler := rH.GetNext(); nextHandler != nil {
		return nextHandler.Handle(req)
	}
	return nil
}
func (rH *requestHandler) GetNext() RequestHandler {
	return rH.next
}
func (rH *requestHandler) SetNext(next RequestHandler) {
	rH.next = next
}
