//go:build preview.query

package query

type ErrorAdder interface {
	addErrors(errs ...error)
}
