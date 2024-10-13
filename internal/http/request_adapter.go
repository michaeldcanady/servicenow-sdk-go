package http

import (
	"context"
	"errors"
	"net/http"
	"reflect"
)

type RequestConverter interface {
	Convert(ctx context.Context, info RequestInformation) (*http.Request, error)
}

type ResponseConverter interface {
	Convert(resp *http.Response) (Response, error)
}

type RequestAdapter interface {
	Send(ctx context.Context, info RequestInformation) (*http.Response, error)
	AddHandler(handler RequestHandler) error
	GetBaseURL() string
}

type Client[T any, R any] interface {
	Send(T) (R, error)
}

type requestAdapter struct {
	requestHandler RequestHandler
	client         *http.Client
	reqConverter   RequestConverter
	respConverter  ResponseConverter
}

func NewClient2CompatibleRequestAdapter() Client2 {
	return &client2Adapter{adaptee: NewRequestAdapter()}
}

func NewRequestAdapter() RequestAdapter {
	return &requestAdapter{
		client: http.DefaultClient,
	}
}

func (rA *requestAdapter) Send(ctx context.Context, info RequestInformation, parsable Parsable, mapping ErrorMapping) (interface{}, error) {
	if err := rA.requestHandler.Handle(info); err != nil {
		return nil, err
	}

	req, err := rA.reqConverter.Convert(ctx, info)
	if err != nil {
		return nil, err
	}

	resp, err := rA.client.Do(req)
	if err != nil {
		return nil, err
	}

	return rA.respConverter.Convert(resp)
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
