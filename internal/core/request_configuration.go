package core

type RequestConfiguration struct {
	Header          interface{}
	QueryParameters interface{}
	Data            interface{}
	ErrorMapping    ErrorMapping
	Response        Response
}

type RequestConfiguration2 interface {
	Header() interface{}
	Query() interface{}
	Data() interface{}
	Mapping() ErrorMapping
	Response() Response
}
