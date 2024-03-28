package internal

import (
	"context"
	"fmt"
	"net/http"
)

type Client struct {
	requestHandler  Handler[RequestInformation]
	responseHandler Handler[*http.Response]
	session         http.Client
}

func NewClient(authProvider AuthorizationProvider) *Client {
	return NewClientWithHandlers(NewAuthorizationHandler(authProvider), nil)
}

func NewClientWithHandlers(requestHandler Handler[RequestInformation], responseHandler Handler[*http.Response]) *Client {
	return &Client{
		requestHandler:  requestHandler,
		responseHandler: responseHandler,
		session:         http.Client{},
	}
}

func NewClientWithHandlersWithSerialization(requestHandler Handler[RequestInformation], responseHandler Handler[*http.Response], serializationFactory) *Client {
	return &Client{
		requestHandler:  requestHandler,
		responseHandler: responseHandler,
		session:         http.Client{},
	}
}

func (c *Client) toNativeRequest(ctx context.Context, requestInfo RequestInformation) (*http.Request, error) {
	if requestInfo == nil {
		return nil, ErrNilRequestInfo
	}

	if ctx == nil {
		return nil, ErrNilContext
	}

	err := c.requestHandler.Handle(requestInfo)
	if err != nil {
		return nil, err
	}

	request, err := ConvertToNativeRequest(ctx, requestInfo)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	return request, err
}

func (c *Client) setErrorResponseHandler(errorMapping ErrorMapping) {
	factory := NewErrorFactory()

	for statusCode, serializer := range errorMapping.mapping {
		factory.RegisterSerializer(statusCode, serializer)
	}

	errorHandler := NewErrorResponseHandler(factory)

	if IsNil(c.responseHandler) {
		c.responseHandler = errorHandler
	} else {
		c.responseHandler.SetNext(errorHandler)
	}
}

func (c *Client) SendWithContext(ctx context.Context, requestInfo RequestInformation, errorMapping ErrorMapping) (*http.Response, error) {
	request, err := c.toNativeRequest(ctx, requestInfo)
	if err != nil {
		return nil, err
	}

	response, err := c.session.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request: %s", err)
	}

	c.setErrorResponseHandler(errorMapping)

	err = c.responseHandler.Handle(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
