package core

type RequestConfiguration interface {
	GetErrorMapping() ErrorMapping
	SetResponse(Response)
	GetResponse() Response
}

type RequestConfigurationImpl struct {
	Header          interface{}
	QueryParameters interface{}
	Data            interface{}
	ErrorMapping    ErrorMapping
	Response        Response
}

func (rC *RequestConfigurationImpl) GetHeaders() interface{} {
	return rC.Header
}
func (rC *RequestConfigurationImpl) GetQueryParams() interface{} {
	return rC.QueryParameters
}
func (rC *RequestConfigurationImpl) GetData() interface{} {
	return rC.Data
}

func (rC *RequestConfigurationImpl) GetErrorMapping() ErrorMapping {
	return rC.ErrorMapping
}
func (rC *RequestConfigurationImpl) SetResponse(resp Response) {
	rC.Response = resp
}
func (rC *RequestConfigurationImpl) GetResponse() Response {
	return rC.Response
}

type SupportsHeaders[T any] interface {
	GetHeaders() T
}

type SupportsQueryParams[T any] interface {
	GetQueryParams() interface{}
}

type SupportsData[T any] interface {
	GetData() T
}
