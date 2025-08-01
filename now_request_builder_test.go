package servicenowsdkgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNowRequestBuilder(t *testing.T) {
	client := &ServiceNowClient{} // Replace with your client implementation
	url := "https://example.service-now.com/api"
	builder := NewNowRequestBuilder(url, client)

	expected := map[string]string{
		"baseurl": url,
	}

	assert.NotNil(t, builder)
	assert.Equal(t, expected, builder.PathParameters)
}

func TestNowRequestBuilder_Table(t *testing.T) {
	client := &ServiceNowClient{} // Replace with your client implementation
	url := "https://example.service-now.com/api"
	builder := NewNowRequestBuilder(url, client)
	tableName := "incident"

	expected := map[string]string{
		"baseurl": url,
		"table":   tableName,
	}

	tableBuilder := builder.Table(tableName)

	assert.NotNil(t, tableBuilder)
	assert.Equal(t, expected, tableBuilder.PathParameters)
}

func TestNowRequestBuilder_Attachment(t *testing.T) {
	client := &ServiceNowClient{} // Replace with your client implementation
	url := "https://example.service-now.com/api"
	builder := NewNowRequestBuilder(url, client)

	expected := map[string]string{
		"baseurl": url,
	}

	attachmentBuilder := builder.Attachment()

	assert.NotNil(t, attachmentBuilder)
	assert.Equal(t, expected, attachmentBuilder.PathParameters)
}

func TestNowRequestBuilder_Attachment2(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				client := &ServiceNowClient{}
				url := "https://example.service-now.com/api"
				builder := NewNowRequestBuilder(url, client)

				expected := map[string]string{
					"baseurl": url,
				}

				attachmentBuilder := builder.Attachment2()

				assert.NotNil(t, attachmentBuilder)
				assert.Equal(t, expected, attachmentBuilder.GetPathParameters())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNowRequestBuilder_Batch(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				client := &ServiceNowClient{} // Replace with your client implementation
				url := "https://example.service-now.com/api"
				builder := NewNowRequestBuilder(url, client)

				expected := map[string]string{
					"baseurl": url,
				}

				attachmentBuilder := builder.Batch()

				assert.NotNil(t, attachmentBuilder)
				assert.Equal(t, expected, attachmentBuilder.GetPathParameters())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
