package internal

type RequestConfiguration interface {
	Header() interface{}
	QueryParameters() interface{}
	Data() interface{}
	ErrorMapping() ErrorMapping
	Response() Response
}
