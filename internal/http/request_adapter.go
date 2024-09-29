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
	requestHandler  RequestHandler
	client          *http.Client
	responseHandler any
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

func (ra *requestAdapter) AddHandler(handler RequestHandler) error {
	if ra.requestHandler == nil {
		ra.requestHandler = handler
		return nil
	}

	current := ra.requestHandler
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

func (ra *requestAdapter) GetBaseURL() string {
	return ""
}
