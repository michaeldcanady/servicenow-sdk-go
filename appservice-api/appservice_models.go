package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// CreateOrUpdateServiceRequest represents the request body for creating or updating an application service.
type CreateOrUpdateServiceRequest struct {
	newInternal.BaseModel
}

func NewCreateOrUpdateServiceRequest() *CreateOrUpdateServiceRequest {
	return &CreateOrUpdateServiceRequest{BaseModel: *newInternal.NewBaseModel()}
}

func (m *CreateOrUpdateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
	)
}

func (m *CreateOrUpdateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey: internalSerialization.DeserializeStringFunc()(m.setName),
		typeKey: internalSerialization.DeserializeStringFunc()(m.setType),
	}
}

func (m *CreateOrUpdateServiceRequest) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}

func (m *CreateOrUpdateServiceRequest) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}

func (m *CreateOrUpdateServiceRequest) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}

func (m *CreateOrUpdateServiceRequest) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}

func CreateCreateOrUpdateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateOrUpdateServiceRequest(), nil
}

// CreateOrUpdateServiceResult represents the result details of a created or updated application service.
type CreateOrUpdateServiceResult struct {
	newInternal.BaseModel
}

func NewCreateOrUpdateServiceResult() *CreateOrUpdateServiceResult {
	return &CreateOrUpdateServiceResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *CreateOrUpdateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(classNameKey)(m.GetClassName),
	)
}

func (m *CreateOrUpdateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:     internalSerialization.DeserializeStringFunc()(m.setSysId),
		nameKey:      internalSerialization.DeserializeStringFunc()(m.setName),
		classNameKey: internalSerialization.DeserializeStringFunc()(m.setClassName),
	}
}

func (m *CreateOrUpdateServiceResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}

func (m *CreateOrUpdateServiceResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}

func (m *CreateOrUpdateServiceResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}

func (m *CreateOrUpdateServiceResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}

func (m *CreateOrUpdateServiceResult) GetClassName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), classNameKey)
}

func (m *CreateOrUpdateServiceResult) setClassName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), classNameKey, val)
}

func CreateCreateOrUpdateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateOrUpdateServiceResult(), nil
}

// CreateOrUpdateServiceResponse represents the response containing the created or updated application service details.
type CreateOrUpdateServiceResponse interface {
	newInternal.ServiceNowItemResponse[*CreateOrUpdateServiceResult]
}

func CreateCreateOrUpdateServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*CreateOrUpdateServiceResult](CreateCreateOrUpdateServiceResultFromDiscriminatorValue), nil
}

// CIInfo represents a Configuration Item in the application service content.
type CIInfo struct {
	newInternal.BaseModel
}

func NewCIInfo() *CIInfo {
	return &CIInfo{BaseModel: *newInternal.NewBaseModel()}
}

func (m *CIInfo) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(classNameKey)(m.GetClassName),
	)
}

func (m *CIInfo) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:     internalSerialization.DeserializeStringFunc()(m.setSysId),
		nameKey:      internalSerialization.DeserializeStringFunc()(m.setName),
		classNameKey: internalSerialization.DeserializeStringFunc()(m.setClassName),
	}
}

func (m *CIInfo) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}

func (m *CIInfo) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}

func (m *CIInfo) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}

func (m *CIInfo) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}

func (m *CIInfo) GetClassName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), classNameKey)
}

func (m *CIInfo) setClassName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), classNameKey, val)
}

func CreateCIInfoFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCIInfo(), nil
}

// RelationshipInfo represents a relationship between CIs.
type RelationshipInfo struct {
	newInternal.BaseModel
}

func NewRelationshipInfo() *RelationshipInfo {
	return &RelationshipInfo{BaseModel: *newInternal.NewBaseModel()}
}

func (m *RelationshipInfo) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(parentKey)(m.GetParent),
		internalSerialization.SerializeStringFunc(childKey)(m.GetChild),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
	)
}

func (m *RelationshipInfo) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		parentKey: internalSerialization.DeserializeStringFunc()(m.setParent),
		childKey:  internalSerialization.DeserializeStringFunc()(m.setChild),
		typeKey:   internalSerialization.DeserializeStringFunc()(m.setType),
	}
}

func (m *RelationshipInfo) GetParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), parentKey)
}

func (m *RelationshipInfo) setParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), parentKey, val)
}

func (m *RelationshipInfo) GetChild() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), childKey)
}

func (m *RelationshipInfo) setChild(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), childKey, val)
}

func (m *RelationshipInfo) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}

func (m *RelationshipInfo) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}

func CreateRelationshipInfoFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRelationshipInfo(), nil
}

// GetContentResult represents the result details retrieved from the getContent endpoint.
type GetContentResult struct {
	newInternal.BaseModel
}

func NewGetContentResult() *GetContentResult {
	return &GetContentResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *GetContentResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[*CIInfo](cisKey)(m.GetCis),
		internalSerialization.SerializeCollectionOfObjectValuesFunc[*RelationshipInfo](relationsKey)(m.GetRelations),
	)
}

func (m *GetContentResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		cisKey:       internalSerialization.DeserializeCollectionOfObjectValuesFunc[*CIInfo](CreateCIInfoFromDiscriminatorValue)(m.setCis),
		relationsKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*RelationshipInfo](CreateRelationshipInfoFromDiscriminatorValue)(m.setRelations),
	}
}

func (m *GetContentResult) GetCis() ([]*CIInfo, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []*CIInfo](m.GetBackingStore(), cisKey)
}

func (m *GetContentResult) setCis(val []*CIInfo) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cisKey, val)
}

func (m *GetContentResult) GetRelations() ([]*RelationshipInfo, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []*RelationshipInfo](m.GetBackingStore(), relationsKey)
}

func (m *GetContentResult) setRelations(val []*RelationshipInfo) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), relationsKey, val)
}

func CreateGetContentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewGetContentResult(), nil
}

// GetContentResponse represents the response containing the application service content details.
type GetContentResponse interface {
	newInternal.ServiceNowItemResponse[*GetContentResult]
}

func CreateGetContentResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*GetContentResult](CreateGetContentResultFromDiscriminatorValue), nil
}
