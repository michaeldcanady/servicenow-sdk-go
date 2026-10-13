// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// NodesResponse represents the response for a list of nodes.
type NodesResponse interface {
	core.ServiceNowCollectionResponse[*NodeResultModel]
}

// CreateNodesResponseFromDiscriminatorValue creates a new NodesResponse from a ParseNode.
func CreateNodesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil
}

// NodeResponse represents the response for a single node.
type NodeResponse interface {
	core.ServiceNowItemResponse[*NodeResultModel]
}

// CreateNodeResponseFromDiscriminatorValue creates a new NodeResponse from a ParseNode.
func CreateNodeResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil
}

// ValidationResponse represents the response for configuration validation.
type ValidationResponse interface {
	core.ServiceNowItemResponse[*ValidationResultModel]
}

// CreateValidationResponseFromDiscriminatorValue creates a new ValidationResponse from a ParseNode.
func CreateValidationResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ValidationResultModel](CreateValidationResultFromDiscriminatorValue), nil
}
