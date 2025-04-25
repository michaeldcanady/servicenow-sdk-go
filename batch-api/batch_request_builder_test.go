package batchapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestNewBatchRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				var params = map[string]string{"key": "value"}
				var mockAdapter = mocking.NewMockRequestAdapter()

				builder := NewBatchRequestBuilderInternal(params, mockAdapter)

				assert.NotNil(t, builder)
				assert.IsType(t, &BatchRequestBuilder{}, builder)
				assert.Equal(t, mockAdapter, builder.RequestBuilder.GetRequestAdapter())
				assert.Equal(t, params, builder.RequestBuilder.GetPathParameters())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNewBatchRequestBuilder2(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				const rawURL = ""
				var mockAdapter = mocking.NewMockRequestAdapter()

				builder := NewBatchRequestBuilder(rawURL, mockAdapter)

				assert.NotNil(t, builder)
				assert.IsType(t, &BatchRequestBuilder{}, builder)
				assert.Equal(t, mockAdapter, builder.RequestBuilder.GetRequestAdapter())
				assert.Equal(t, map[string]string{newInternal.RawURLKey: rawURL}, builder.RequestBuilder.GetPathParameters())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestBatchRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestBatchRequestBuilder_toPostRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
