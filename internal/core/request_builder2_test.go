package core

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockClient2 struct {
	mock.Mock
}

func (c *mockClient2) Send(requestInfo RequestInformation, errorMapping ErrorMapping) (*http.Response, error) {
	args := c.Called(requestInfo, errorMapping)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (c *mockClient2) SendWithContext(ctx context.Context, requestInfo RequestInformation, errorMapping ErrorMapping) (*http.Response, error) {
	args := c.Called(ctx, requestInfo, errorMapping)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (c *mockClient2) GetBaseURL() string {
	args := c.Called()
	return args.String(0)
}

func TestNewRequestBuilder2(t *testing.T) {
	client := &mockClient2{}
	urlTemplate := "http://example.com/{param}"
	pathParameters := map[string]string{"param": "test"}

	builder := NewRequestBuilder2(client, urlTemplate, pathParameters).(*requestBuilder2)

	assert.Equal(t, client, builder.client)
	assert.Equal(t, urlTemplate, builder.urlTemplate)
	assert.Equal(t, pathParameters, builder.pathParameters)
}

func TestGetPathParameters(t *testing.T) {
	pathParameters := map[string]string{"param": "test"}
	builder := &requestBuilder2{pathParameters: pathParameters}

	got := builder.GetPathParameters()

	assert.Equal(t, pathParameters, got)
}

func TestGetClient(t *testing.T) {
	client := &mockClient2{}
	builder := &requestBuilder2{client: client}

	got := builder.GetClient()

	assert.Equal(t, client, got)
}

func TestGetURLTemplate(t *testing.T) {
	urlTemplate := "http://example.com/{param}"
	builder := &requestBuilder2{urlTemplate: urlTemplate}

	got := builder.GetURLTemplate()

	assert.Equal(t, urlTemplate, got)
}

func TestSend(t *testing.T) {
	client := &mockClient2{}
	builder := NewRequestBuilder2(client, "/test/{param}", map[string]string{"param": "value"}).(*requestBuilder2)

	tests := []Test[any]{
		{
			Title: "Success",
			Setup: func() {
				client.On("SendWithContext", context.Background(), GET, []RequestConfigurationOption{WithData("test")}).Return("test", nil)
			},
			ExpectedErr: nil,
			Expected:    "test",
			Cleanup: func() {
				ResetCalls(client.ExpectedCalls...)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			resp, err := builder.Send(context.Background(), GET, []RequestConfigurationOption{WithData("test")}...)

			assert.Equal(t, test.Expected, resp)
			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
