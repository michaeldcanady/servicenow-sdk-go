package internal

type RequestConfiguration interface {
	Header() RequestHeader
	Query() map[string]string
	Data() string
	ErrorMapping() ErrorMapping
	SetResponse(interface{})
}
