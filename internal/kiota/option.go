package kiota

import "github.com/michaeldcanady/servicenow-sdk-go/internal/utils"

// Option[T] represents an optional parameter.
type Option[T any] func(T) error

// ApplyOptions applies options to the applicable type.
func ApplyOptions[T any](config T, opts ...Option[T]) error {
	if utils.IsNil(config) {
		return nil
	}
	if len(opts) == 0 {
		return nil
	}
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return err
		}
	}
	return nil
}
