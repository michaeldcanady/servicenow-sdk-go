// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateServiceRequest represents the request body for creating an application service.
type CreateServiceRequest struct {
	core.BaseModel
}

// NewCreateServiceRequest creates a new instance of CreateServiceRequest.
func NewCreateServiceRequest() *CreateServiceRequest {
	return &CreateServiceRequest{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *CreateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey, m.GetComments),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CreateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:     internalSerialization.DeserializeStringFunc(m.setName),
		commentsKey: internalSerialization.DeserializeStringFunc(m.setComments),
	}
}

// GetName returns the name value.
func (m *CreateServiceRequest) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceRequest, *string](m, nameKey)
}

func (m *CreateServiceRequest) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// SetName sets the name value.
func (m *CreateServiceRequest) SetName(val *string) error {
	return m.setName(val)
}

// GetComments returns the comments value.
func (m *CreateServiceRequest) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceRequest, *string](m, commentsKey)
}

func (m *CreateServiceRequest) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, commentsKey, val)
}

// SetComments sets the comments value.
func (m *CreateServiceRequest) SetComments(val *string) error {
	return m.setComments(val)
}

// CreateCreateServiceRequestFromDiscriminatorValue creates a new CreateServiceRequest from a ParseNode.
func CreateCreateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateServiceRequest(), nil
}

// CreateServiceResult represents the result details of a created application service.
type CreateServiceResult struct {
	core.BaseModel
}

// NewCreateServiceResult creates a new instance of CreateServiceResult.
func NewCreateServiceResult() *CreateServiceResult {
	return &CreateServiceResult{BaseModel: *core.NewBaseModel()}
}

const (
	urlKey           = "url"
	getContentUrlKey = "getContentUrl"
	infoKey          = "info"
)

// Serialize writes the objects properties to the current writer.
func (m *CreateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey, m.GetComments),
		internalSerialization.SerializeStringFunc(urlKey, m.GetURL),
		internalSerialization.SerializeStringFunc(getContentUrlKey, m.GetGetContentURL),
		internalSerialization.SerializeStringFunc(infoKey, m.GetInfo),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CreateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:         internalSerialization.DeserializeStringFunc(m.setSysID),
		nameKey:          internalSerialization.DeserializeStringFunc(m.setName),
		commentsKey:      internalSerialization.DeserializeStringFunc(m.setComments),
		urlKey:           internalSerialization.DeserializeStringFunc(m.setURL),
		getContentUrlKey: internalSerialization.DeserializeStringFunc(m.setGetContentURL),
		infoKey:          internalSerialization.DeserializeStringFunc(m.setInfo),
	}
}

// GetSysID returns the sys id value.
func (m *CreateServiceResult) GetSysID() (*string, error) {
	val, err := store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, sysIDKey)
	if err == nil && val != nil && *val != "" {
		return val, nil
	}
	// Fallback: extract from URL
	urlVal, err := m.GetURL()
	if err == nil && urlVal != nil && *urlVal != "" {
		parts := strings.Split(*urlVal, "/")
		if len(parts) > 0 {
			sysID := parts[len(parts)-1]
			if len(sysID) == 32 { // standard sys_id length
				return &sysID, nil
			}
		}
	}
	// Fallback 2: extract from getContentUrl
	getContentURLVal, err := m.GetGetContentURL()
	if err == nil && getContentURLVal != nil && *getContentURLVal != "" {
		parts := strings.Split(*getContentURLVal, "/")
		if len(parts) >= 2 {
			sysID := parts[len(parts)-2]
			if len(sysID) == 32 {
				return &sysID, nil
			}
		}
	}
	return nil, nil
}

func (m *CreateServiceResult) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// SetSysID sets the sys id value.
func (m *CreateServiceResult) SetSysID(val *string) error {
	return m.setSysID(val)
}

// GetName returns the name value.
func (m *CreateServiceResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, nameKey)
}

func (m *CreateServiceResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// SetName sets the name value.
func (m *CreateServiceResult) SetName(val *string) error {
	return m.setName(val)
}

// GetComments returns the comments value.
func (m *CreateServiceResult) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, commentsKey)
}

func (m *CreateServiceResult) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, commentsKey, val)
}

// SetComments sets the comments value.
func (m *CreateServiceResult) SetComments(val *string) error {
	return m.setComments(val)
}

// GetURL returns the url value.
func (m *CreateServiceResult) GetURL() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, urlKey)
}

func (m *CreateServiceResult) setURL(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, urlKey, val)
}

// SetURL sets the url value.
func (m *CreateServiceResult) SetURL(val *string) error {
	return m.setURL(val)
}

// GetGetContentURL returns the getContentUrl value.
func (m *CreateServiceResult) GetGetContentURL() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, getContentUrlKey)
}

func (m *CreateServiceResult) setGetContentURL(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, getContentUrlKey, val)
}

// SetGetContentURL sets the getContentUrl value.
func (m *CreateServiceResult) SetGetContentURL(val *string) error {
	return m.setGetContentURL(val)
}

// GetInfo returns the info value.
func (m *CreateServiceResult) GetInfo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, infoKey)
}

func (m *CreateServiceResult) setInfo(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, infoKey, val)
}

// SetInfo sets the info value.
func (m *CreateServiceResult) SetInfo(val *string) error {
	return m.setInfo(val)
}

// CreateCreateServiceResultFromDiscriminatorValue creates a new CreateServiceResult from a ParseNode.
func CreateCreateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateServiceResult(), nil
}
