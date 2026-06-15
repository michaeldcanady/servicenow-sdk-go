package caseapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestCaseRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	t.Run("Search Cases", func(t *testing.T) {
		requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, requestInfo)
		assert.Equal(t, caseURLTemplate, requestInfo.UrlTemplate)
	})

	t.Run("Get Case By ID", func(t *testing.T) {
		itemBuilder := builder.ByID("test-id")
		requestInfo, err := itemBuilder.ToGetRequestInformation(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, requestInfo)
		assert.Equal(t, caseItemURLTemplate, requestInfo.UrlTemplate)
		assert.Equal(t, "test-id", requestInfo.PathParameters["id"])
	})
}
