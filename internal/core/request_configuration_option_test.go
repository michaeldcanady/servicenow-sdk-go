package core

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockResponse struct {
	headers http.Header
}

func (m *mockResponse) ParseHeaders(headers http.Header) {
	m.headers = headers
}

type mockErrorMapping struct {
	*mock.Mock
}

func (m *mockErrorMapping) Set(code, err string) {
	_ = m.Called(code, err)
}

func (m *mockErrorMapping) Len() int {
	args := m.Called()
	return args.Int(0)
}

func (m *mockErrorMapping) Get(code int) (string, bool) {
	args := m.Called(code)
	return args.String(0), args.Bool(1)
}

func TestWithHeader(t *testing.T) {
	header := "TestHeader"
	opt := WithHeader(header)
	config := &RequestConfigurationImpl{}
	opt(config)

	assert.Equal(t, header, config.Header)
}

func TestWithQueryParameters(t *testing.T) {
	queryParams := "TestQueryParameters"
	opt := WithQueryParameters(queryParams)
	config := &RequestConfigurationImpl{}
	opt(config)

	assert.Equal(t, queryParams, config.QueryParameters)
}

func TestWithData(t *testing.T) {
	data := "TestData"
	opt := WithData(data)
	config := &RequestConfigurationImpl{}
	opt(config)

	assert.Equal(t, data, config.Data)
}

func TestWithErrorMapping(t *testing.T) {
	errorMapping := &mockErrorMapping{}
	opt := WithErrorMapping(errorMapping)
	config := &RequestConfigurationImpl{}
	opt(config)

	assert.Equal(t, errorMapping, config.ErrorMapping)
}

func TestWithResponse(t *testing.T) {
	response := &mockResponse{}
	opt := WithResponse(response)
	config := &RequestConfigurationImpl{}
	opt(config)

	if !reflect.DeepEqual(config.Response, response) {
		t.Errorf("Expected response %v, got %v", response, config.Response)
	}
}

func TestApplyOptions(t *testing.T) {
	header := "TestHeader"
	queryParams := "TestQueryParameters"
	data := "TestData"
	errorMapping := &mockErrorMapping{}
	response := &mockResponse{}

	config := ApplyOptions(
		WithHeader(header),
		WithQueryParameters(queryParams),
		WithData(data),
		WithErrorMapping(errorMapping),
		WithResponse(response),
	)

	assert.Equal(t, header, config.Header)
	assert.Equal(t, queryParams, config.QueryParameters)
	assert.Equal(t, data, config.Data)
	assert.Equal(t, errorMapping, config.ErrorMapping)
}
