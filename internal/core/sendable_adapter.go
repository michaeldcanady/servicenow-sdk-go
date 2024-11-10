package core

import (
	"context"
	"net/http"
)

type ClientSendableAdapterFunc[T any] func(T, context.Context, RequestInformation, ErrorMapping) (*http.Response, error)

type clientSendableAdapter[T any] struct {
	adapter ClientSendableAdapterFunc[T]
	adaptee T
}

func NewClietSendableAdapter[T any](adapter ClientSendableAdapterFunc[T], adaptee T) ClientSendable {
	return &clientSendableAdapter[T]{
		adapter: adapter,
		adaptee: adaptee,
	}
}

func (s *clientSendableAdapter[T]) SendWithContext(ctx context.Context, info RequestInformation, mapping ErrorMapping) (*http.Response, error) {
	return s.adapter(s.adaptee, ctx, info, mapping)
}
