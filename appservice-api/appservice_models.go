package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// CreateServiceRequest represents the request body for creating an application service.
type CreateServiceRequest struct {
	newInternal.BaseModel
}

func NewCreateServiceRequest() *CreateServiceRequest {
	return &CreateServiceRequest{BaseModel: *newInternal.NewBaseModel()}
}

func (m *CreateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey)(m.GetComments),
	)
}

func (m *CreateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:     internalSerialization.DeserializeStringFunc()(m.setName),
		commentsKey: internalSerialization.DeserializeStringFunc()(m.setComments),
	}
}

func (m *CreateServiceRequest) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}

func (m *CreateServiceRequest) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}

func (m *CreateServiceRequest) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), commentsKey)
}

func (m *CreateServiceRequest) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), commentsKey, val)
}

func CreateCreateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateServiceRequest(), nil
}

// CreateServiceResult represents the result details of a created application service.
type CreateServiceResult struct {
	newInternal.BaseModel
}

func NewCreateServiceResult() *CreateServiceResult {
	return &CreateServiceResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *CreateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey)(m.GetComments),
	)
}

func (m *CreateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:    internalSerialization.DeserializeStringFunc()(m.setSysId),
		nameKey:     internalSerialization.DeserializeStringFunc()(m.setName),
		commentsKey: internalSerialization.DeserializeStringFunc()(m.setComments),
	}
}

func (m *CreateServiceResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}

func (m *CreateServiceResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}

func (m *CreateServiceResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}

func (m *CreateServiceResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}

func (m *CreateServiceResult) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), commentsKey)
}

func (m *CreateServiceResult) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), commentsKey, val)
}

func CreateCreateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateServiceResult(), nil
}

// CreateServiceResponse represents the response containing the created application service details.
type CreateServiceResponse interface {
	newInternal.ServiceNowItemResponse[*CreateServiceResult]
}

func CreateCreateServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*CreateServiceResult](CreateCreateServiceResultFromDiscriminatorValue), nil
}

// FindServiceResult represents the result details retrieved from the find_service endpoint.
type FindServiceResult struct {
	newInternal.BaseModel
}

func NewFindServiceResult() *FindServiceResult {
	return &FindServiceResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *FindServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
		internalSerialization.SerializeStringFunc(environmentKey)(m.GetEnvironment),
		internalSerialization.SerializeStringFunc(versionKey)(m.GetVersion),
	)
}

func (m *FindServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:       internalSerialization.DeserializeStringFunc()(m.setSysId),
		nameKey:        internalSerialization.DeserializeStringFunc()(m.setName),
		numberKey:      internalSerialization.DeserializeStringFunc()(m.setNumber),
		environmentKey: internalSerialization.DeserializeStringFunc()(m.setEnvironment),
		versionKey:     internalSerialization.DeserializeStringFunc()(m.setVersion),
	}
}

func (m *FindServiceResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}

func (m *FindServiceResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}

func (m *FindServiceResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}

func (m *FindServiceResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}

func (m *FindServiceResult) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), numberKey)
}

func (m *FindServiceResult) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), numberKey, val)
}

func (m *FindServiceResult) GetEnvironment() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), environmentKey)
}

func (m *FindServiceResult) setEnvironment(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), environmentKey, val)
}

func (m *FindServiceResult) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), versionKey)
}

func (m *FindServiceResult) setVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), versionKey, val)
}

func CreateFindServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFindServiceResult(), nil
}

// FindServiceResponse represents the response containing the found application service details.
type FindServiceResponse interface {
	newInternal.ServiceNowItemResponse[*FindServiceResult]
}

func CreateFindServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*FindServiceResult](CreateFindServiceResultFromDiscriminatorValue), nil
}

// BasicDetails represents the basic details schema inside CSDM requests.
type BasicDetails struct {
	newInternal.BaseModel
}

func NewBasicDetails() *BasicDetails {
	return &BasicDetails{BaseModel: *newInternal.NewBaseModel()}
}

func (m *BasicDetails) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(environmentKey)(m.GetEnvironment),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(versionKey)(m.GetVersion),
		internalSerialization.SerializeStringFunc(businessAppKey)(m.GetBusinessApp),
		internalSerialization.SerializeStringFunc(businessServiceOfferingKey)(m.GetBusinessServiceOffering),
		internalSerialization.SerializeStringFunc(technicalServiceOfferingKey)(m.GetTechnicalServiceOffering),
	)
}

func (m *BasicDetails) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		environmentKey:              internalSerialization.DeserializeStringFunc()(m.setEnvironment),
		nameKey:                     internalSerialization.DeserializeStringFunc()(m.setName),
		versionKey:                  internalSerialization.DeserializeStringFunc()(m.setVersion),
		businessAppKey:              internalSerialization.DeserializeStringFunc()(m.setBusinessApp),
		businessServiceOfferingKey:  internalSerialization.DeserializeStringFunc()(m.setBusinessServiceOffering),
		technicalServiceOfferingKey: internalSerialization.DeserializeStringFunc()(m.setTechnicalServiceOffering),
	}
}

func (m *BasicDetails) GetEnvironment() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), environmentKey)
}

func (m *BasicDetails) setEnvironment(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), environmentKey, val)
}

func (m *BasicDetails) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}

func (m *BasicDetails) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}

func (m *BasicDetails) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), versionKey)
}

func (m *BasicDetails) setVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), versionKey, val)
}

func (m *BasicDetails) GetBusinessApp() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), businessAppKey)
}

func (m *BasicDetails) setBusinessApp(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), businessAppKey, val)
}

func (m *BasicDetails) GetBusinessServiceOffering() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), businessServiceOfferingKey)
}

func (m *BasicDetails) setBusinessServiceOffering(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), businessServiceOfferingKey, val)
}

func (m *BasicDetails) GetTechnicalServiceOffering() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), technicalServiceOfferingKey)
}

func (m *BasicDetails) setTechnicalServiceOffering(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), technicalServiceOfferingKey, val)
}

func CreateBasicDetailsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewBasicDetails(), nil
}

// RegisterServiceRequest represents the request body for registering a CSDM service.
type RegisterServiceRequest struct {
	newInternal.BaseModel
}

func NewRegisterServiceRequest() *RegisterServiceRequest {
	return &RegisterServiceRequest{BaseModel: *newInternal.NewBaseModel()}
}

func (m *RegisterServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*BasicDetails](basicDetailsKey)(m.GetBasicDetails),
	)
}

func (m *RegisterServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		basicDetailsKey: internalSerialization.DeserializeObjectValueFunc[*BasicDetails](CreateBasicDetailsFromDiscriminatorValue)(m.setBasicDetails),
	}
}

func (m *RegisterServiceRequest) GetBasicDetails() (*BasicDetails, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *BasicDetails](m.GetBackingStore(), basicDetailsKey)
}

func (m *RegisterServiceRequest) setBasicDetails(val *BasicDetails) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), basicDetailsKey, val)
}

func CreateRegisterServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRegisterServiceRequest(), nil
}

// RegisterServiceResult represents the result details of a registered CSDM service.
type RegisterServiceResult struct {
	newInternal.BaseModel
}

func NewRegisterServiceResult() *RegisterServiceResult {
	return &RegisterServiceResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *RegisterServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
	)
}

func (m *RegisterServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:   internalSerialization.DeserializeStringFunc()(m.setSysId),
		numberKey:  internalSerialization.DeserializeStringFunc()(m.setNumber),
		statusKey:  internalSerialization.DeserializeStringFunc()(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc()(m.setMessage),
	}
}

func (m *RegisterServiceResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}

func (m *RegisterServiceResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}

func (m *RegisterServiceResult) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), numberKey)
}

func (m *RegisterServiceResult) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), numberKey, val)
}

func (m *RegisterServiceResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), statusKey)
}

func (m *RegisterServiceResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), statusKey, val)
}

func (m *RegisterServiceResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), messageKey)
}

func (m *RegisterServiceResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), messageKey, val)
}

func CreateRegisterServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRegisterServiceResult(), nil
}

// RegisterServiceResponse represents the response containing the registered CSDM service details.
type RegisterServiceResponse interface {
	newInternal.ServiceNowItemResponse[*RegisterServiceResult]
}

func CreateRegisterServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*RegisterServiceResult](CreateRegisterServiceResultFromDiscriminatorValue), nil
}

// ServiceRelation represents a relationship between components inside Populate request.
type ServiceRelation struct {
	newInternal.BaseModel
}

func NewServiceRelation() *ServiceRelation {
	return &ServiceRelation{BaseModel: *newInternal.NewBaseModel()}
}

func (m *ServiceRelation) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(parentKey)(m.GetParent),
		internalSerialization.SerializeStringFunc(childKey)(m.GetChild),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
	)
}

func (m *ServiceRelation) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		parentKey: internalSerialization.DeserializeStringFunc()(m.setParent),
		childKey:  internalSerialization.DeserializeStringFunc()(m.setChild),
		typeKey:   internalSerialization.DeserializeStringFunc()(m.setType),
	}
}

func (m *ServiceRelation) GetParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), parentKey)
}

func (m *ServiceRelation) setParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), parentKey, val)
}

func (m *ServiceRelation) GetChild() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), childKey)
}

func (m *ServiceRelation) setChild(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), childKey, val)
}

func (m *ServiceRelation) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}

func (m *ServiceRelation) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}

func CreateServiceRelationFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceRelation(), nil
}

// PopulateServiceRequest represents the request body for populating a CSDM service.
type PopulateServiceRequest struct {
	newInternal.BaseModel
}

func NewPopulateServiceRequest() *PopulateServiceRequest {
	return &PopulateServiceRequest{BaseModel: *newInternal.NewBaseModel()}
}

func (m *PopulateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[*ServiceRelation](serviceRelationsKey)(m.GetServiceRelations),
	)
}

func (m *PopulateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		serviceRelationsKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*ServiceRelation](CreateServiceRelationFromDiscriminatorValue)(m.setServiceRelations),
	}
}

func (m *PopulateServiceRequest) GetServiceRelations() ([]*ServiceRelation, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []*ServiceRelation](m.GetBackingStore(), serviceRelationsKey)
}

func (m *PopulateServiceRequest) setServiceRelations(val []*ServiceRelation) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), serviceRelationsKey, val)
}

func CreatePopulateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPopulateServiceRequest(), nil
}

// PopulateServiceResult represents the result details of populating a service.
type PopulateServiceResult struct {
	newInternal.BaseModel
}

func NewPopulateServiceResult() *PopulateServiceResult {
	return &PopulateServiceResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *PopulateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
	)
}

func (m *PopulateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:  internalSerialization.DeserializeStringFunc()(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc()(m.setMessage),
	}
}

func (m *PopulateServiceResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), statusKey)
}

func (m *PopulateServiceResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), statusKey, val)
}

func (m *PopulateServiceResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), messageKey)
}

func (m *PopulateServiceResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), messageKey, val)
}

func CreatePopulateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPopulateServiceResult(), nil
}

// PopulateServiceResponse represents the response containing populate result details.
type PopulateServiceResponse interface {
	newInternal.ServiceNowItemResponse[*PopulateServiceResult]
}

func CreatePopulateServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*PopulateServiceResult](CreatePopulateServiceResultFromDiscriminatorValue), nil
}

// ServiceDetailsRequest represents the request body for modifying basic details of a CSDM service.
type ServiceDetailsRequest struct {
	newInternal.BaseModel
}

func NewServiceDetailsRequest() *ServiceDetailsRequest {
	return &ServiceDetailsRequest{BaseModel: *newInternal.NewBaseModel()}
}

func (m *ServiceDetailsRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*BasicDetails](basicDetailsKey)(m.GetBasicDetails),
	)
}

func (m *ServiceDetailsRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		basicDetailsKey: internalSerialization.DeserializeObjectValueFunc[*BasicDetails](CreateBasicDetailsFromDiscriminatorValue)(m.setBasicDetails),
	}
}

func (m *ServiceDetailsRequest) GetBasicDetails() (*BasicDetails, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *BasicDetails](m.GetBackingStore(), basicDetailsKey)
}

func (m *ServiceDetailsRequest) setBasicDetails(val *BasicDetails) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), basicDetailsKey, val)
}

func CreateServiceDetailsRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceDetailsRequest(), nil
}

// ServiceDetailsResult represents the result details of modifying service details.
type ServiceDetailsResult struct {
	newInternal.BaseModel
}

func NewServiceDetailsResult() *ServiceDetailsResult {
	return &ServiceDetailsResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *ServiceDetailsResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
	)
}

func (m *ServiceDetailsResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:  internalSerialization.DeserializeStringFunc()(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc()(m.setMessage),
	}
}

func (m *ServiceDetailsResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), statusKey)
}

func (m *ServiceDetailsResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), statusKey, val)
}

func (m *ServiceDetailsResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), messageKey)
}

func (m *ServiceDetailsResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), messageKey, val)
}

func CreateServiceDetailsResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceDetailsResult(), nil
}

// ServiceDetailsResponse represents the response containing service details update status.
type ServiceDetailsResponse interface {
	newInternal.ServiceNowItemResponse[*ServiceDetailsResult]
}

func CreateServiceDetailsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*ServiceDetailsResult](CreateServiceDetailsResultFromDiscriminatorValue), nil
}
