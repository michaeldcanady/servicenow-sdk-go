package requesthandler

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

type requiredPathParameterHandler struct {
	requestHandler
	keys []string
}

func NewRequiredPathParameterHandler(keys []string) RequestHandler {
	return &requiredPathParameterHandler{
		keys: keys,
	}
}

func (rH *requiredPathParameterHandler) Handle(req core.RequestInformation) error {
	return rH.requestHandler.Handle(req)
}
func (rH *requiredPathParameterHandler) GetNext() RequestHandler {
	return rH.next
}
func (rH *requiredPathParameterHandler) SetNext(next RequestHandler) {
	rH.next = next
}
