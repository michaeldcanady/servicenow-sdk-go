package core

//TODO: Convert to interface
type RequestConfiguration struct {
	Header          interface{}
	QueryParameters interface{}
	Data            interface{}
	ErrorMapping    ErrorMapping
	Response        Response
}
