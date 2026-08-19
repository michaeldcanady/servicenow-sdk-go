package internal

type ErrorThrower interface {
	Throw(typeName string, statusCode int64, contentType string, content []byte) error
}
