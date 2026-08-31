// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package policyapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// PoliciesMappingable defines the properties and methods of a PoliciesMapping.
type PoliciesMappingable interface {
	GetDescription() (*string, error)
	SetDescription(*string) error
	GetDocument() (*string, error)
	SetDocument(*string) error
	GetDocumentRef() (*Ref, error)
	SetDocumentRef(*Ref) error
	GetError() (*core.MainError, error)
	SetError(*core.MainError) error
	GetException() (*string, error)
	SetException(*string) error
	GetExceptionAllowed() (*bool, error)
	SetExceptionAllowed(*bool) error
	GetInputStatus() (*InputStatus, error)
	SetInputStatus(*InputStatus) error
	GetLastUpdatedBy() (*Ref, error)
	SetLastUpdatedBy(*Ref) error
	GetNumber() (*string, error)
	SetNumber(*string) error
	GetPolicy() (*Ref, error)
	SetPolicy(*Ref) error
	GetReason() (*string, error)
	SetReason(*string) error
	GetState() (*State, error)
	SetState(*State) error
	GetSysClassName() (*string, error)
	SetSysClassName(*string) error
	GetSysCreatedBy() (*string, error)
	SetSysCreatedBy(*string) error
	GetSysCreatedOn() (*time.Time, error)
	SetSysCreatedOn(*time.Time) error
	GetSysID() (*string, error)
	SetSysID(*string) error
	GetSysUpdatedBy() (*string, error)
	SetSysUpdatedBy(*string) error
	GetSysUpdatedOn() (*time.Time, error)
	SetSysUpdatedOn(*time.Time) error
	serialization.Parsable
	core.BackedModel
}

// PoliciesMapping represents a Service-Now policy mapping.
type PoliciesMapping struct {
	*core.BaseModel
}

// PoliciesMappingsResolved* are the field keys used to (de)serialize a PoliciesMapping.
const (
	PoliciesMappingsResolvedDescription      string = "description"
	PoliciesMappingsResolvedDocument         string = "document"
	PoliciesMappingsResolvedDocumentRef      string = "document_ref"
	PoliciesMappingsResolvedError            string = "error"
	PoliciesMappingsResolvedException        string = "exception"
	PoliciesMappingsResolvedExceptionAllowed string = "exception_allowed"
	PoliciesMappingsResolvedInputStatus      string = "input_status"
	PoliciesMappingsResolvedLastUpdatedBy    string = "last_updated_by"
	PoliciesMappingsResolvedNumber           string = "number"
	PoliciesMappingsResolvedPolicy           string = "policy"
	PoliciesMappingsResolvedReason           string = "reason"
	PoliciesMappingsResolvedState            string = "state"
	PoliciesMappingsResolvedSysClassName     string = "sys_class_name"
	PoliciesMappingsResolvedSysCreatedBy     string = "sys_created_by"
	PoliciesMappingsResolvedSysCreatedOn     string = "sys_created_on"
	PoliciesMappingsResolvedSysID            string = "sys_id"
	PoliciesMappingsResolvedSysUpdatedBy     string = "sys_updated_by"
	PoliciesMappingsResolvedSysUpdatedOn     string = "sys_updated_on"
)

// NewPoliciesMapping creates a new instance of PoliciesMappingsInput.
func NewPoliciesMapping() *PoliciesMapping {
	return &PoliciesMapping{
		BaseModel: core.NewBaseModel(),
	}
}

// CreatePoliciesMappingsInputFromDiscriminatorValue creates a new PoliciesMappingsInput from a ParseNode.
func CreatePoliciesMappingsInputFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPoliciesMapping(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (p *PoliciesMapping) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		PoliciesMappingsResolvedDescription:      internalSerialization.DeserializeStringFunc(p.SetDescription),
		PoliciesMappingsResolvedDocument:         internalSerialization.DeserializeStringFunc(p.SetDocument),
		PoliciesMappingsResolvedDocumentRef:      internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue, p.SetDocumentRef),
		PoliciesMappingsResolvedError:            internalSerialization.DeserializeObjectValueFunc[*core.MainError](core.CreateMainErrorFromDiscriminatorValue, p.SetError),
		PoliciesMappingsResolvedException:        internalSerialization.DeserializeStringFunc(p.SetException),
		PoliciesMappingsResolvedExceptionAllowed: internalSerialization.DeserializeBoolFunc(p.SetExceptionAllowed),
		PoliciesMappingsResolvedInputStatus:      internalSerialization.DeserializeEnumFunc[InputStatus](ParseInputStatus, p.SetInputStatus),
		PoliciesMappingsResolvedLastUpdatedBy:    internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue, p.SetLastUpdatedBy),
		PoliciesMappingsResolvedNumber:           internalSerialization.DeserializeStringFunc(p.SetNumber),
		PoliciesMappingsResolvedPolicy:           internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue, p.SetPolicy),
		PoliciesMappingsResolvedReason:           internalSerialization.DeserializeStringFunc(p.SetReason),
		PoliciesMappingsResolvedState:            internalSerialization.DeserializeEnumFunc[State](ParseState, p.SetState),
		PoliciesMappingsResolvedSysClassName:     internalSerialization.DeserializeStringFunc(p.SetSysClassName),
		PoliciesMappingsResolvedSysCreatedBy:     internalSerialization.DeserializeStringFunc(p.SetSysCreatedBy),
		PoliciesMappingsResolvedSysCreatedOn:     internalSerialization.DeserializeTimeFunc(p.SetSysCreatedOn),
		PoliciesMappingsResolvedSysID:            internalSerialization.DeserializeStringFunc(p.SetSysID),
		PoliciesMappingsResolvedSysUpdatedBy:     internalSerialization.DeserializeStringFunc(p.SetSysUpdatedBy),
		PoliciesMappingsResolvedSysUpdatedOn:     internalSerialization.DeserializeTimeFunc(p.SetSysUpdatedOn),
	}
}

// Serialize writes the objects properties to the current writer.
func (p *PoliciesMapping) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(p) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedDescription, p.GetDescription),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedDocument, p.GetDocument),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedDocumentRef, p.GetDocumentRef),
		internalSerialization.SerializeObjectValueFunc[*core.MainError](PoliciesMappingsResolvedError, p.GetError),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedException, p.GetException),
		internalSerialization.SerializeBoolFunc(PoliciesMappingsResolvedExceptionAllowed, p.GetExceptionAllowed),
		internalSerialization.SerializeEnumFunc[InputStatus](PoliciesMappingsResolvedInputStatus, p.GetInputStatus),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedLastUpdatedBy, p.GetLastUpdatedBy),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedNumber, p.GetNumber),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedPolicy, p.GetPolicy),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedReason, p.GetReason),
		internalSerialization.SerializeEnumFunc[State](PoliciesMappingsResolvedState, p.GetState),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedSysClassName, p.GetSysClassName),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedSysCreatedBy, p.GetSysCreatedBy),
		internalSerialization.SerializeTimeFunc(PoliciesMappingsResolvedSysCreatedOn, p.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedSysID, p.GetSysID),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedSysUpdatedBy, p.GetSysUpdatedBy),
		internalSerialization.SerializeTimeFunc(PoliciesMappingsResolvedSysUpdatedOn, p.GetSysUpdatedOn),
	)
}

// Getters and Setters

// GetDescription returns the description.
func (p *PoliciesMapping) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedDescription)
}

// SetDescription sets the description.
func (p *PoliciesMapping) SetDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedDescription, val)
}

// GetDocument returns the document.
func (p *PoliciesMapping) GetDocument() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedDocument)
}

// SetDocument sets the document.
func (p *PoliciesMapping) SetDocument(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedDocument, val)
}

// GetDocumentRef returns the document reference.
func (p *PoliciesMapping) GetDocumentRef() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *Ref](p, PoliciesMappingsResolvedDocumentRef)
}

// SetDocumentRef sets the document reference.
func (p *PoliciesMapping) SetDocumentRef(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedDocumentRef, val)
}

// GetError returns the error.
func (p *PoliciesMapping) GetError() (*core.MainError, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *core.MainError](p, PoliciesMappingsResolvedError)
}

// SetError sets the error.
func (p *PoliciesMapping) SetError(val *core.MainError) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedError, val)
}

// GetException returns the exception.
func (p *PoliciesMapping) GetException() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedException)
}

// SetException sets the exception.
func (p *PoliciesMapping) SetException(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedException, val)
}

// GetExceptionAllowed returns whether an exception is allowed.
func (p *PoliciesMapping) GetExceptionAllowed() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *bool](p, PoliciesMappingsResolvedExceptionAllowed)
}

// SetExceptionAllowed sets whether an exception is allowed.
func (p *PoliciesMapping) SetExceptionAllowed(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedExceptionAllowed, val)
}

// GetInputStatus returns the input status.
func (p *PoliciesMapping) GetInputStatus() (*InputStatus, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *InputStatus](p, PoliciesMappingsResolvedInputStatus)
}

// SetInputStatus sets the input status.
func (p *PoliciesMapping) SetInputStatus(val *InputStatus) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedInputStatus, val)
}

// GetLastUpdatedBy returns the reference to who last updated the mapping.
func (p *PoliciesMapping) GetLastUpdatedBy() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *Ref](p, PoliciesMappingsResolvedLastUpdatedBy)
}

// SetLastUpdatedBy sets the reference to who last updated the mapping.
func (p *PoliciesMapping) SetLastUpdatedBy(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedLastUpdatedBy, val)
}

// GetNumber returns the number.
func (p *PoliciesMapping) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedNumber)
}

// SetNumber sets the number.
func (p *PoliciesMapping) SetNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedNumber, val)
}

// GetPolicy returns the reference to the associated policy.
func (p *PoliciesMapping) GetPolicy() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *Ref](p, PoliciesMappingsResolvedPolicy)
}

// SetPolicy sets the reference to the associated policy.
func (p *PoliciesMapping) SetPolicy(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedPolicy, val)
}

// GetReason returns the reason.
func (p *PoliciesMapping) GetReason() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedReason)
}

// SetReason sets the reason.
func (p *PoliciesMapping) SetReason(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedReason, val)
}

// GetState returns the state.
func (p *PoliciesMapping) GetState() (*State, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *State](p, PoliciesMappingsResolvedState)
}

// SetState sets the state.
func (p *PoliciesMapping) SetState(val *State) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedState, val)
}

// GetSysClassName returns the sys_class_name.
func (p *PoliciesMapping) GetSysClassName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedSysClassName)
}

// SetSysClassName sets the sys_class_name.
func (p *PoliciesMapping) SetSysClassName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedSysClassName, val)
}

// GetSysCreatedBy returns who created the mapping.
func (p *PoliciesMapping) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedSysCreatedBy)
}

// SetSysCreatedBy sets who created the mapping.
func (p *PoliciesMapping) SetSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedSysCreatedBy, val)
}

// GetSysCreatedOn returns when the mapping was created.
func (p *PoliciesMapping) GetSysCreatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *time.Time](p, PoliciesMappingsResolvedSysCreatedOn)
}

// SetSysCreatedOn sets when the mapping was created.
func (p *PoliciesMapping) SetSysCreatedOn(val *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedSysCreatedOn, val)
}

// GetSysID returns the sys_id.
func (p *PoliciesMapping) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedSysID)
}

// SetSysID sets the sys_id.
func (p *PoliciesMapping) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedSysID, val)
}

// GetSysUpdatedBy returns who last updated the mapping.
func (p *PoliciesMapping) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *string](p, PoliciesMappingsResolvedSysUpdatedBy)
}

// SetSysUpdatedBy sets who last updated the mapping.
func (p *PoliciesMapping) SetSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedSysUpdatedBy, val)
}

// GetSysUpdatedOn returns when the mapping was last updated.
func (p *PoliciesMapping) GetSysUpdatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[*PoliciesMapping, *time.Time](p, PoliciesMappingsResolvedSysUpdatedOn)
}

// SetSysUpdatedOn sets when the mapping was last updated.
func (p *PoliciesMapping) SetSysUpdatedOn(val *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(p, PoliciesMappingsResolvedSysUpdatedOn, val)
}
