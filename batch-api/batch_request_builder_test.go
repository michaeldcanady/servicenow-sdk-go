package batchapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBatchRequestBuilder(t *testing.T) {
	client := MockClient{}
	pathParameters := map[string]string{}

	builder := NewBatchRequestBuilder(&client, pathParameters)

	assert.Equal(t, "{+baseurl}/v1/batch", builder.UrlTemplate)
	assert.Equal(t, pathParameters, builder.PathParameters)
	assert.Equal(t, &client, builder.Client)
}
