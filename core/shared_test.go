package core

import (
	"context"
	"io"
	"net/http"
	"strings"
)

type mockPage struct{}

func (m *mockPage) GetResult() ([]*any, error)       { return nil, nil } // Fixed signature
func (m *mockPage) GetNextPageLink() string          { return "" }       // Fixed signature
func (m *mockPage) GetPreviousPageLink() string      { return "" }       // Fixed signature
func (m *mockPage) GetFirstPageLink() string         { return "" }       // Fixed signature
func (m *mockPage) GetLastPageLink() string          { return "" }       // Fixed signature
func (m *mockPage) ToPage() PageResult[any]          { return PageResult[any]{} }
func (m *mockPage) ParseHeaders(headers http.Header) {}

type mockCoreClient struct {
	SendFunc func(ri IRequestInformation, em ErrorMapping) (*http.Response, error)
}

func (m *mockCoreClient) Send(ri IRequestInformation, em ErrorMapping) (*http.Response, error) {
	if m.SendFunc != nil {
		return m.SendFunc(ri, em)
	}
	return &http.Response{Body: io.NopCloser(strings.NewReader(`{}`))}, nil
}

type mockCoreClient2 struct {
	mockCoreClient
	SendWithContextFunc func(ctx context.Context, ri IRequestInformation, em ErrorMapping) (*http.Response, error)
}

func (m *mockCoreClient2) SendWithContext(ctx context.Context, ri IRequestInformation, em ErrorMapping) (*http.Response, error) {
	if m.SendWithContextFunc != nil {
		return m.SendWithContextFunc(ctx, ri, em)
	}
	return &http.Response{Body: io.NopCloser(strings.NewReader(`{}`))}, nil
}
