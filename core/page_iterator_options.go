package core

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type PageIteratorOption[T serialization.Parsable] = internal.Option[*PageIterator[T]]

// WithHeaders sets the headers for the next page request.
func WithHeaders[T serialization.Parsable](headers *abstractions.RequestHeaders) PageIteratorOption[T] {
	return func(i *PageIterator[T]) error {
		i.headers = headers
		return nil
	}
}

// WithRequestOptions adds the request options for the next page request.
func WithRequestOptions[T serialization.Parsable](options ...abstractions.RequestOption) PageIteratorOption[T] {
	return func(i *PageIterator[T]) error {
		i.reqOptions = append(i.reqOptions, options...)
		return nil
	}
}
