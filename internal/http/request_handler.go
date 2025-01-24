package http

type RequestHandler interface {
	Handle(RequestInformation) error
	SetNext(RequestHandler)
	Next() RequestHandler
}
