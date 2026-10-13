// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserWindowRequest represents the request body for the userwindow endpoint. The spec defines
// no schema for this body.
type UserWindowRequest struct {
	core.BaseModel
}

// NewUserWindowRequest creates a new instance of UserWindowRequest.
func NewUserWindowRequest() *UserWindowRequest {
	return &UserWindowRequest{BaseModel: *core.NewBaseModel()}
}

// CreateUserWindowRequestFromDiscriminatorValue creates a new UserWindowRequest from a ParseNode.
func CreateUserWindowRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserWindowRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *UserWindowRequest) Serialize(_ serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *UserWindowRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}
