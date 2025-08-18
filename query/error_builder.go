//go:build preview.query

package query

type ErrorBuilder interface {
	addErrors(errs ...error)
}
