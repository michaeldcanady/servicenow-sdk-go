package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// NodesResponse represents the response for a list of nodes.
type NodesResponse interface {
	internal.ServiceNowCollectionResponse[*NodeResultModel]
}

func CreateNodesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil
}

// NodeResponse represents the response for a single node.
type NodeResponse interface {
	internal.ServiceNowItemResponse[*NodeResultModel]
}

func CreateNodeResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil
}

// ValidationResponse represents the response for configuration validation.
type ValidationResponse interface {
	internal.ServiceNowItemResponse[*ValidationResultModel]
}

func CreateValidationResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*ValidationResultModel](CreateValidationResultFromDiscriminatorValue), nil
}

// NodeDeleteResponse represents the response for deleting a node.
type NodeDeleteResponse interface {
	internal.ServiceNowItemResponse[*MessageResult]
}

func CreateNodeDeleteResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*MessageResult](CreateMessageResultFromDiscriminatorValue), nil
}
