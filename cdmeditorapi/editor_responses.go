package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// NodesResponse represents the response for a list of nodes.
type NodesResponse interface {
	core.ServiceNowCollectionResponse[*NodeResultModel]
}

func CreateNodesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil
}

// NodeResponse represents the response for a single node.
type NodeResponse interface {
	core.ServiceNowItemResponse[*NodeResultModel]
}

func CreateNodeResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil
}

// ValidationResponse represents the response for configuration validation.
type ValidationResponse interface {
	core.ServiceNowItemResponse[*ValidationResultModel]
}

func CreateValidationResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ValidationResultModel](CreateValidationResultFromDiscriminatorValue), nil
}

// NodeDeleteResponse represents the response for deleting a node.
type NodeDeleteResponse interface {
	core.ServiceNowItemResponse[*MessageResult]
}

func CreateNodeDeleteResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*MessageResult](CreateMessageResultFromDiscriminatorValue), nil
}
