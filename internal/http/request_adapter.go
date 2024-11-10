package http

import (
	"context"
	"errors"
	"net/http"
	"reflect"
)

type RequestAdapter interface {
	Send(ctx context.Context, info RequestInformation) (*http.Response, error)
	AddHandler(handler RequestHandler) error
	GetBaseURL() string
}

type requestAdapter struct {
	requestHandler RequestHandler
	client         *http.Client
}

func NewClient2CompatibleRequestAdapter() Client2 {
	return &client2Adapter{adaptee: NewRequestAdapter()}
}

func NewRequestAdapter() RequestAdapter {
	return &requestAdapter{
		client: http.DefaultClient,
	}
}

func (rA *requestAdapter) Send(ctx context.Context, info RequestInformation) (*http.Response, error) {
	if err := rA.requestHandler.Handle(info); err != nil {
		return nil, err
	}

	req, err := info.ToRequestWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return rA.client.Do(req)
}

func (rA *requestAdapter) AddHandler(handler RequestHandler) error {
	if rA.requestHandler == nil {
		rA.requestHandler = handler
		return nil
	}

	current := rA.requestHandler
	for current != nil {
		if reflect.TypeOf(current) == reflect.TypeOf(handler) {
			return errors.New("handler of this type already exists")
		}
		if next := current.Next(); next == nil {
			current.SetNext(handler)
			return nil
		} else {
			current = next
		}
	}
	return nil
}

func (rA *requestAdapter) GetBaseURL() string {
	return ""
}
