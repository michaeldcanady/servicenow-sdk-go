package core

type RequestConfiguration struct {
	Header          interface{}
	QueryParameters interface{}
	Data            interface{}
	ErrorMapping    ErrorMapping
	Response        Response
}
