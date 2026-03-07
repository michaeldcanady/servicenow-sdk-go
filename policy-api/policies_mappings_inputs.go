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

type PoliciesMappingsInputable interface {
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
	PoliciesMappingsInputsResolvedDescription      string = "description"
	PoliciesMappingsInputsResolvedDocument         string = "document"
	PoliciesMappingsInputsResolvedDocumentRef      string = "document_ref"
	PoliciesMappingsInputsResolvedError            string = "error"
	PoliciesMappingsInputsResolvedException        string = "exception"
	PoliciesMappingsInputsResolvedExceptionAllowed string = "exception_allowed"
	PoliciesMappingsInputsResolvedInputStatus      string = "input_status"
	PoliciesMappingsInputsResolvedLastUpdatedBy    string = "last_updated_by"
	PoliciesMappingsInputsResolvedNumber           string = "number"
	PoliciesMappingsInputsResolvedPolicy           string = "policy"
	PoliciesMappingsInputsResolvedReason           string = "reason"
	PoliciesMappingsInputsResolvedState            string = "state"
	PoliciesMappingsInputsResolvedSysClassName     string = "sys_class_name"
	PoliciesMappingsInputsResolvedSysCreatedBy     string = "sys_created_by"
	PoliciesMappingsInputsResolvedSysCreatedOn     string = "sys_created_on"
	PoliciesMappingsInputsResolvedSysId            string = "sys_id"
	PoliciesMappingsInputsResolvedSysUpdatedBy     string = "sys_updated_by"
	PoliciesMappingsInputsResolvedSysUpdatedOn     string = "sys_updated_on"
)

type PoliciesMappingsInput struct {
	newInternal.Model
}

// NewPoliciesMappingsInput creates a new instance of PoliciesMappingsInput.
func NewPoliciesMappingsInput() *PoliciesMappingsInput {
	return &PoliciesMappingsInput{
		Model: newInternal.NewBaseModel(),
	}
}

// CreatePoliciesMappingsInputFromDiscriminatorValue creates a new PoliciesMappingsInput from a ParseNode.
func CreatePoliciesMappingsInputFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPoliciesMappingsInput(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (p *PoliciesMappingsInput) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		PoliciesMappingsInputsResolvedDescription:      internalSerialization.DeserializeStringFunc()(p.SetDescription),
		PoliciesMappingsInputsResolvedDocument:         internalSerialization.DeserializeStringFunc()(p.SetDocument),
		PoliciesMappingsInputsResolvedDocumentRef:      internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetDocumentRef),
		PoliciesMappingsInputsResolvedError:            internalSerialization.DeserializeObjectValueFunc[*newInternal.MainError](newInternal.CreateMainErrorFromDiscriminatorValue)(p.SetError),
		PoliciesMappingsInputsResolvedException:        internalSerialization.DeserializeStringFunc()(p.SetException),
		PoliciesMappingsInputsResolvedExceptionAllowed: internalSerialization.DeserializeBoolFunc()(p.SetExceptionAllowed),
		PoliciesMappingsInputsResolvedInputStatus:      internalSerialization.DeserializeEnumFunc[InputStatus](ParseInputStatus)(p.SetInputStatus),
		PoliciesMappingsInputsResolvedLastUpdatedBy:    internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetLastUpdatedBy),
		PoliciesMappingsInputsResolvedNumber:           internalSerialization.DeserializeStringFunc()(p.SetNumber),
		PoliciesMappingsInputsResolvedPolicy:           internalSerialization.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetPolicy),
		PoliciesMappingsInputsResolvedReason:           internalSerialization.DeserializeStringFunc()(p.SetReason),
		PoliciesMappingsInputsResolvedState:            internalSerialization.DeserializeEnumFunc[State](ParseState)(p.SetState),
		PoliciesMappingsInputsResolvedSysClassName:     internalSerialization.DeserializeStringFunc()(p.SetSysClassName),
		PoliciesMappingsInputsResolvedSysCreatedBy:     internalSerialization.DeserializeStringFunc()(p.SetSysCreatedBy),
		PoliciesMappingsInputsResolvedSysCreatedOn:     internalSerialization.DeserializeTimeFunc()(p.SetSysCreatedOn),
		PoliciesMappingsInputsResolvedSysId:            internalSerialization.DeserializeStringFunc()(p.SetSysId),
		PoliciesMappingsInputsResolvedSysUpdatedBy:     internalSerialization.DeserializeStringFunc()(p.SetSysUpdatedBy),
		PoliciesMappingsInputsResolvedSysUpdatedOn:     internalSerialization.DeserializeTimeFunc()(p.SetSysUpdatedOn),
	}
}

// Serialize writes the objects properties to the current writer.
func (p *PoliciesMappingsInput) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(p) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedDescription)(p.GetDescription),
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedDocument)(p.GetDocument),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsInputsResolvedDocumentRef)(p.GetDocumentRef),
		internalSerialization.SerializeObjectValueFunc[*newInternal.MainError](PoliciesMappingsInputsResolvedError)(p.GetError),
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedException)(p.GetException),
		internalSerialization.SerializeBoolFunc(PoliciesMappingsInputsResolvedExceptionAllowed)(p.GetExceptionAllowed),
		internalSerialization.SerializeEnumFunc[InputStatus](PoliciesMappingsInputsResolvedInputStatus)(p.GetInputStatus),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsInputsResolvedLastUpdatedBy)(p.GetLastUpdatedBy),
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedNumber)(p.GetNumber),
		internalSerialization.SerializeObjectValueFunc[*Ref](PoliciesMappingsInputsResolvedPolicy)(p.GetPolicy),
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedReason)(p.GetReason),
		internalSerialization.SerializeEnumFunc[State](PoliciesMappingsInputsResolvedState)(p.GetState),
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedSysClassName)(p.GetSysClassName),
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedSysCreatedBy)(p.GetSysCreatedBy),
		internalSerialization.SerializeTimeFunc(PoliciesMappingsInputsResolvedSysCreatedOn)(p.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedSysId)(p.GetSysId),
		internalSerialization.SerializeStringFunc(PoliciesMappingsInputsResolvedSysUpdatedBy)(p.GetSysUpdatedBy),
		internalSerialization.SerializeTimeFunc(PoliciesMappingsInputsResolvedSysUpdatedOn)(p.GetSysUpdatedOn),
	)
}

// Getters and Setters

func (p *PoliciesMappingsInput) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedDescription)
}

func (p *PoliciesMappingsInput) SetDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedDescription, val)
}

func (p *PoliciesMappingsInput) GetDocument() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedDocument)
}

func (p *PoliciesMappingsInput) SetDocument(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedDocument, val)
}

func (p *PoliciesMappingsInput) GetDocumentRef() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsInputsResolvedDocumentRef)
}

func (p *PoliciesMappingsInput) SetDocumentRef(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedDocumentRef, val)
}

func (p *PoliciesMappingsInput) GetError() (*newInternal.MainError, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *newInternal.MainError](p.GetBackingStore(), PoliciesMappingsInputsResolvedError)
}

func (p *PoliciesMappingsInput) SetError(val *newInternal.MainError) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedError, val)
}

func (p *PoliciesMappingsInput) GetException() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedException)
}

func (p *PoliciesMappingsInput) SetException(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedException, val)
}

func (p *PoliciesMappingsInput) GetExceptionAllowed() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](p.GetBackingStore(), PoliciesMappingsInputsResolvedExceptionAllowed)
}

func (p *PoliciesMappingsInput) SetExceptionAllowed(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedExceptionAllowed, val)
}

func (p *PoliciesMappingsInput) GetInputStatus() (*InputStatus, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *InputStatus](p.GetBackingStore(), PoliciesMappingsInputsResolvedInputStatus)
}

func (p *PoliciesMappingsInput) SetInputStatus(val *InputStatus) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedInputStatus, val)
}

func (p *PoliciesMappingsInput) GetLastUpdatedBy() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsInputsResolvedLastUpdatedBy)
}

func (p *PoliciesMappingsInput) SetLastUpdatedBy(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedLastUpdatedBy, val)
}

func (p *PoliciesMappingsInput) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedNumber)
}

func (p *PoliciesMappingsInput) SetNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedNumber, val)
}

func (p *PoliciesMappingsInput) GetPolicy() (*Ref, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsInputsResolvedPolicy)
}

func (p *PoliciesMappingsInput) SetPolicy(val *Ref) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedPolicy, val)
}

func (p *PoliciesMappingsInput) GetReason() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedReason)
}

func (p *PoliciesMappingsInput) SetReason(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedReason, val)
}

func (p *PoliciesMappingsInput) GetState() (*State, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *State](p.GetBackingStore(), PoliciesMappingsInputsResolvedState)
}

func (p *PoliciesMappingsInput) SetState(val *State) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedState, val)
}

func (p *PoliciesMappingsInput) GetSysClassName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedSysClassName)
}

func (p *PoliciesMappingsInput) SetSysClassName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedSysClassName, val)
}

func (p *PoliciesMappingsInput) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedSysCreatedBy)
}

func (p *PoliciesMappingsInput) SetSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedSysCreatedBy, val)
}

func (p *PoliciesMappingsInput) GetSysCreatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](p.GetBackingStore(), PoliciesMappingsInputsResolvedSysCreatedOn)
}

func (p *PoliciesMappingsInput) SetSysCreatedOn(val *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedSysCreatedOn, val)
}

func (p *PoliciesMappingsInput) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedSysId)
}

func (p *PoliciesMappingsInput) SetSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedSysId, val)
}

func (p *PoliciesMappingsInput) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsInputsResolvedSysUpdatedBy)
}

func (p *PoliciesMappingsInput) SetSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedSysUpdatedBy, val)
}

func (p *PoliciesMappingsInput) GetSysUpdatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](p.GetBackingStore(), PoliciesMappingsInputsResolvedSysUpdatedOn)
}

func (p *PoliciesMappingsInput) SetSysUpdatedOn(val *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsInputsResolvedSysUpdatedOn, val)
}
