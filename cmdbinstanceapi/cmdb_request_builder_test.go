package cmdbinstanceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCmdbRequestBuilder(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}

	builder := NewCmdbRequestBuilder("https://example.com/api/now/v1/cmdb", adapter)

	require.NotNil(t, builder)
	assert.Equal(t, cmdbURLTemplate, builder.GetURLTemplate())
}

func TestNewCmdbRequestBuilderInternal(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}

	builder := NewCmdbRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	require.NotNil(t, builder)
	assert.Equal(t, cmdbURLTemplate, builder.GetURLTemplate())
}

func TestCmdbRequestBuilder_Instance(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCmdbRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	instanceBuilder := builder.Instance()

	require.NotNil(t, instanceBuilder)
	assert.Equal(t, cmdbInstanceURLTemplate, instanceBuilder.GetURLTemplate())
}

func TestCmdbRequestBuilder_AppService(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCmdbRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	appServiceBuilder := builder.AppService()

	require.NotNil(t, appServiceBuilder)
}
