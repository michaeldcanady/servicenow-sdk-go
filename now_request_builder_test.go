package servicenowsdkgo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNowRequestBuilder(t *testing.T) {
	client := &ServiceNowClient{} // Replace with your client implementation
	url := "https://example.service-now.com/api"
	builder := NewNowRequestBuilder(context.Background(), url, client)

	expected := map[string]string{
		"baseurl": url,
	}

	assert.NotNil(t, builder)
	assert.Equal(t, expected, builder.RequestBuilder.PathParameters)
}

func TestNowRequestBuilderTable(t *testing.T) {
	client := &ServiceNowClient{} // Replace with your client implementation
	url := "https://example.service-now.com/api"
	builder := NewNowRequestBuilder(context.Background(), url, client)
	tableName := "incident"

	expected := map[string]string{
		"baseurl": url,
		"table":   tableName,
	}

	tableBuilder := builder.Table(tableName)

	assert.NotNil(t, tableBuilder)
	assert.Equal(t, expected, tableBuilder.RequestBuilder.PathParameters)
}

func TestNowRequestBuilderAttachment(t *testing.T) {
	client := &ServiceNowClient{} // Replace with your client implementation
	url := "https://example.service-now.com/api"
	builder := NewNowRequestBuilder(context.Background(), url, client)

	expected := map[string]string{
		"baseurl": url,
	}

	attachmentBuilder := builder.Attachment()

	assert.NotNil(t, attachmentBuilder)
	assert.Equal(t, expected, attachmentBuilder.RequestBuilder.PathParameters)
}
