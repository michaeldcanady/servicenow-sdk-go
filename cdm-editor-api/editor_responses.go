package cdmeditorapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// NodesResponse represents the response for a list of nodes.
type NodesResponse interface {
	newInternal.ServiceNowCollectionResponse[*NodeResultModel]
}

func CreateNodesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil
}

// NodeResponse represents the response for a single node.
type NodeResponse interface {
	newInternal.ServiceNowItemResponse[*NodeResultModel]
}

func CreateNodeResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil
}

// ValidationResponse represents the response for configuration validation.
type ValidationResponse interface {
	newInternal.ServiceNowItemResponse[*ValidationResultModel]
}

func CreateValidationResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*ValidationResultModel](CreateValidationResultFromDiscriminatorValue), nil
}

// NodeDeleteResponse represents the response for deleting a node.
type NodeDeleteResponse interface {
	newInternal.ServiceNowItemResponse[*MessageResult]
}

func CreateNodeDeleteResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*MessageResult](CreateMessageResultFromDiscriminatorValue), nil
}
