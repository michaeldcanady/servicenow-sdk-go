package cdmeditorapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestNodesRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	sysId := "123"
	config := &NodesRequestBuilderGetRequestConfiguration{
		QueryParameters: &NodesRequestBuilderGetQueryParameters{
			SysId: &sysId,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/editor/v1/nodes?sys_id=123", uri.String())
}

func TestNodeItemRequestBuilder_Put(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewNodeItemRequestBuilderInternal(map[string]string{
		"baseurl":     "https://example.service-now.com",
		"node_sys_id": "node123",
	}, adapter)

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, builder.GetURLTemplate(), builder.GetPathParameters())

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/editor/v1/nodes/node123", uri.String())
}

func TestValidationRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewValidationRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	cdmId := "cdm456"
	config := &ValidationRequestBuilderGetRequestConfiguration{
		QueryParameters: &ValidationRequestBuilderGetQueryParameters{
			CdmId: &cdmId,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, _ := requestInfo.GetUri()
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/editor/v1/validation?cdm_id=cdm456", uri.String())
}
