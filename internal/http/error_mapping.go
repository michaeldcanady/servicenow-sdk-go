package http

type ErrorMapping interface {
	Set(code, err string)
	Len() int
	Get(code int) (string, bool)
}
