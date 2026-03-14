package policyapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

type PoliciesMappingable interface {
	GetDescription() (*string, error)
	SetDescription(*string) error
	GetDocument() (*string, error)
	SetDocument(*string) error
	GetDocumentRef() (*Ref, error)
	SetDocumentRef(*Ref) error
	GetError() (*newInternal.MainError, error)
	SetError(*newInternal.MainError) error
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
	GetSysId() (*string, error)
	SetSysId(*string) error
	GetSysUpdatedBy() (*string, error)
	SetSysUpdatedBy(*string) error
	GetSysUpdatedOn() (*time.Time, error)
	SetSysUpdatedOn(*time.Time) error
	serialization.Parsable
	newInternal.Model
}

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
	PoliciesMappingsResolvedSysId            string = "sys_id"
	PoliciesMappingsResolvedSysUpdatedBy     string = "sys_updated_by"
	PoliciesMappingsResolvedSysUpdatedOn     string = "sys_updated_on"
)

type PoliciesMapping struct {
	newInternal.Model
}

// NewPoliciesMapping creates a new instance of PoliciesMappingsInput.
func NewPoliciesMapping() *PoliciesMapping {
	return &PoliciesMapping{
		Model: newInternal.NewBaseModel(),
	}
}

// CreatePoliciesMappingsInputFromDiscriminatorValue creates a new PoliciesMappingsInput from a ParseNode.
func CreatePoliciesMappingsInputFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPoliciesMapping(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (p *PoliciesMapping) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		PoliciesMappingsResolvedDescription:      internalSerialization.DeserializeStringFunc()(p.SetDescription),
		PoliciesMappingsResolvedDocument:         internalSerialization.DeserializeStringFunc()(p.SetDocument),
		PoliciesMappingsResolvedDocumentRef:      internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetDocumentRef),
		PoliciesMappingsResolvedError:            internalSerialization.DeserializeObjectValueFunc[*newInternal.MainError](newInternal.CreateMainErrorFromDiscriminatorValue)(p.SetError),
		PoliciesMappingsResolvedException:        internalSerialization.DeserializeStringFunc()(p.SetException),
		PoliciesMappingsResolvedExceptionAllowed: internalSerialization.DeserializeBoolFunc()(p.SetExceptionAllowed),
		PoliciesMappingsResolvedInputStatus:      internalSerialization.DeserializeEnumFunc[InputStatus](ParseInputStatus)(p.SetInputStatus),
		PoliciesMappingsResolvedLastUpdatedBy:    internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetLastUpdatedBy),
		PoliciesMappingsResolvedNumber:           internalSerialization.DeserializeStringFunc()(p.SetNumber),
		PoliciesMappingsResolvedPolicy:           internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetPolicy),
		PoliciesMappingsResolvedReason:           internalSerialization.DeserializeStringFunc()(p.SetReason),
		PoliciesMappingsResolvedState:            internalSerialization.DeserializeEnumFunc[State](ParseState)(p.SetState),
		PoliciesMappingsResolvedSysClassName:     internalSerialization.DeserializeStringFunc()(p.SetSysClassName),
		PoliciesMappingsResolvedSysCreatedBy:     internalSerialization.DeserializeStringFunc()(p.SetSysCreatedBy),
		PoliciesMappingsResolvedSysCreatedOn:     internalSerialization.DeserializeTimeFunc()(p.SetSysCreatedOn),
		PoliciesMappingsResolvedSysId:            internalSerialization.DeserializeStringFunc()(p.SetSysId),
		PoliciesMappingsResolvedSysUpdatedBy:     internalSerialization.DeserializeStringFunc()(p.SetSysUpdatedBy),
		PoliciesMappingsResolvedSysUpdatedOn:     internalSerialization.DeserializeTimeFunc()(p.SetSysUpdatedOn),
	}
}

// Serialize writes the objects properties to the current writer.
func (p *PoliciesMapping) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(p) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedDescription)(p.GetDescription),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedDocument)(p.GetDocument),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedDocumentRef)(p.GetDocumentRef),
		internalSerialization.SerializeObjectValueFunc[*newInternal.MainError](PoliciesMappingsResolvedError)(p.GetError),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedException)(p.GetException),
		internalSerialization.SerializeBoolFunc(PoliciesMappingsResolvedExceptionAllowed)(p.GetExceptionAllowed),
		internalSerialization.SerializeEnumFunc[InputStatus](PoliciesMappingsResolvedInputStatus)(p.GetInputStatus),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedLastUpdatedBy)(p.GetLastUpdatedBy),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedNumber)(p.GetNumber),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedPolicy)(p.GetPolicy),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedReason)(p.GetReason),
		internalSerialization.SerializeEnumFunc[State](PoliciesMappingsResolvedState)(p.GetState),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedSysClassName)(p.GetSysClassName),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedSysCreatedBy)(p.GetSysCreatedBy),
		internalSerialization.SerializeTimeFunc(PoliciesMappingsResolvedSysCreatedOn)(p.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedSysId)(p.GetSysId),
		internalSerialization.SerializeStringFunc(PoliciesMappingsResolvedSysUpdatedBy)(p.GetSysUpdatedBy),
		internalSerialization.SerializeTimeFunc(PoliciesMappingsResolvedSysUpdatedOn)(p.GetSysUpdatedOn),
	)
}

// Getters and Setters

func (p *PoliciesMapping) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedDescription)
}

func (p *PoliciesMapping) SetDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedDescription, val)
}

func (p *PoliciesMapping) GetDocument() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedDocument)
}

func (p *PoliciesMapping) SetDocument(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedDocument, val)
}

func (p *PoliciesMapping) GetDocumentRef() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsResolvedDocumentRef)
}

func (p *PoliciesMapping) SetDocumentRef(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedDocumentRef, val)
}

func (p *PoliciesMapping) GetError() (*newInternal.MainError, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *newInternal.MainError](p.GetBackingStore(), PoliciesMappingsResolvedError)
}

func (p *PoliciesMapping) SetError(val *newInternal.MainError) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedError, val)
}

func (p *PoliciesMapping) GetException() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedException)
}

func (p *PoliciesMapping) SetException(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedException, val)
}

func (p *PoliciesMapping) GetExceptionAllowed() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](p.GetBackingStore(), PoliciesMappingsResolvedExceptionAllowed)
}

func (p *PoliciesMapping) SetExceptionAllowed(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedExceptionAllowed, val)
}

func (p *PoliciesMapping) GetInputStatus() (*InputStatus, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *InputStatus](p.GetBackingStore(), PoliciesMappingsResolvedInputStatus)
}

func (p *PoliciesMapping) SetInputStatus(val *InputStatus) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedInputStatus, val)
}

func (p *PoliciesMapping) GetLastUpdatedBy() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsResolvedLastUpdatedBy)
}

func (p *PoliciesMapping) SetLastUpdatedBy(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedLastUpdatedBy, val)
}

func (p *PoliciesMapping) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedNumber)
}

func (p *PoliciesMapping) SetNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedNumber, val)
}

func (p *PoliciesMapping) GetPolicy() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsResolvedPolicy)
}

func (p *PoliciesMapping) SetPolicy(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedPolicy, val)
}

func (p *PoliciesMapping) GetReason() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedReason)
}

func (p *PoliciesMapping) SetReason(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedReason, val)
}

func (p *PoliciesMapping) GetState() (*State, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *State](p.GetBackingStore(), PoliciesMappingsResolvedState)
}

func (p *PoliciesMapping) SetState(val *State) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedState, val)
}

func (p *PoliciesMapping) GetSysClassName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedSysClassName)
}

func (p *PoliciesMapping) SetSysClassName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysClassName, val)
}

func (p *PoliciesMapping) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedSysCreatedBy)
}

func (p *PoliciesMapping) SetSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysCreatedBy, val)
}

func (p *PoliciesMapping) GetSysCreatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](p.GetBackingStore(), PoliciesMappingsResolvedSysCreatedOn)
}

func (p *PoliciesMapping) SetSysCreatedOn(val *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysCreatedOn, val)
}

func (p *PoliciesMapping) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedSysId)
}

func (p *PoliciesMapping) SetSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysId, val)
}

func (p *PoliciesMapping) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedSysUpdatedBy)
}

func (p *PoliciesMapping) SetSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysUpdatedBy, val)
}

func (p *PoliciesMapping) GetSysUpdatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](p.GetBackingStore(), PoliciesMappingsResolvedSysUpdatedOn)
}

func (p *PoliciesMapping) SetSysUpdatedOn(val *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysUpdatedOn, val)
}
