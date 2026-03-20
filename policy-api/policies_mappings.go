package policyapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
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
	GetError() (*model.MainError, error)
	SetError(*model.MainError) error
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
	model.Model
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
	model.Model
}

// NewPoliciesMapping creates a new instance of PoliciesMappingsInput.
func NewPoliciesMapping() *PoliciesMapping {
	return &PoliciesMapping{
		Model: model.NewBaseModel(),
	}
}

// CreatePoliciesMappingsInputFromDiscriminatorValue creates a new PoliciesMappingsInput from a ParseNode.
func CreatePoliciesMappingsInputFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPoliciesMapping(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (p *PoliciesMapping) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		PoliciesMappingsResolvedDescription:      kiota.DeserializeStringFunc(p.SetDescription),
		PoliciesMappingsResolvedDocument:         kiota.DeserializeStringFunc(p.SetDocument),
		PoliciesMappingsResolvedDocumentRef:      kiota.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetDocumentRef),
		PoliciesMappingsResolvedError:            kiota.DeserializeObjectValueFunc[*model.MainError](model.CreateMainErrorFromDiscriminatorValue)(p.SetError),
		PoliciesMappingsResolvedException:        kiota.DeserializeStringFunc(p.SetException),
		PoliciesMappingsResolvedExceptionAllowed: kiota.DeserializeBoolFunc()(p.SetExceptionAllowed),
		PoliciesMappingsResolvedInputStatus:      kiota.DeserializeEnumFunc[InputStatus](ParseInputStatus)(p.SetInputStatus),
		PoliciesMappingsResolvedLastUpdatedBy:    kiota.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetLastUpdatedBy),
		PoliciesMappingsResolvedNumber:           kiota.DeserializeStringFunc(p.SetNumber),
		PoliciesMappingsResolvedPolicy:           kiota.DeserializeObjectValueFunc[*Ref](CreateRefFromDiscriminatorValue)(p.SetPolicy),
		PoliciesMappingsResolvedReason:           kiota.DeserializeStringFunc(p.SetReason),
		PoliciesMappingsResolvedState:            kiota.DeserializeEnumFunc[State](ParseState)(p.SetState),
		PoliciesMappingsResolvedSysClassName:     kiota.DeserializeStringFunc(p.SetSysClassName),
		PoliciesMappingsResolvedSysCreatedBy:     kiota.DeserializeStringFunc(p.SetSysCreatedBy),
		PoliciesMappingsResolvedSysCreatedOn:     kiota.DeserializeTimeFunc()(p.SetSysCreatedOn),
		PoliciesMappingsResolvedSysId:            kiota.DeserializeStringFunc(p.SetSysId),
		PoliciesMappingsResolvedSysUpdatedBy:     kiota.DeserializeStringFunc(p.SetSysUpdatedBy),
		PoliciesMappingsResolvedSysUpdatedOn:     kiota.DeserializeTimeFunc()(p.SetSysUpdatedOn),
	}
}

// Serialize writes the objects properties to the current writer.
func (p *PoliciesMapping) Serialize(writer serialization.SerializationWriter) error {
	if utils.IsNil(p) {
		return nil
	}

	return kiota.Serialize(writer,
		kiota.SerializeStringFunc(PoliciesMappingsResolvedDescription)(p.GetDescription),
		kiota.SerializeStringFunc(PoliciesMappingsResolvedDocument)(p.GetDocument),
		kiota.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedDocumentRef)(p.GetDocumentRef),
		kiota.SerializeObjectValueFunc[*model.MainError](PoliciesMappingsResolvedError)(p.GetError),
		kiota.SerializeStringFunc(PoliciesMappingsResolvedException)(p.GetException),
		kiota.SerializeBoolFunc(PoliciesMappingsResolvedExceptionAllowed)(p.GetExceptionAllowed),
		kiota.SerializeEnumFunc[InputStatus](PoliciesMappingsResolvedInputStatus)(p.GetInputStatus),
		kiota.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedLastUpdatedBy)(p.GetLastUpdatedBy),
		kiota.SerializeStringFunc(PoliciesMappingsResolvedNumber)(p.GetNumber),
		kiota.SerializeObjectValueFunc[*Ref](PoliciesMappingsResolvedPolicy)(p.GetPolicy),
		kiota.SerializeStringFunc(PoliciesMappingsResolvedReason)(p.GetReason),
		kiota.SerializeEnumFunc[State](PoliciesMappingsResolvedState)(p.GetState),
		kiota.SerializeStringFunc(PoliciesMappingsResolvedSysClassName)(p.GetSysClassName),
		kiota.SerializeStringFunc(PoliciesMappingsResolvedSysCreatedBy)(p.GetSysCreatedBy),
		kiota.SerializeTimeFunc(PoliciesMappingsResolvedSysCreatedOn)(p.GetSysCreatedOn),
		kiota.SerializeStringFunc(PoliciesMappingsResolvedSysId)(p.GetSysId),
		kiota.SerializeStringFunc(PoliciesMappingsResolvedSysUpdatedBy)(p.GetSysUpdatedBy),
		kiota.SerializeTimeFunc(PoliciesMappingsResolvedSysUpdatedOn)(p.GetSysUpdatedOn),
	)
}

// Getters and Setters

func (p *PoliciesMapping) GetDescription() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedDescription)
}

func (p *PoliciesMapping) SetDescription(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedDescription, val)
}

func (p *PoliciesMapping) GetDocument() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedDocument)
}

func (p *PoliciesMapping) SetDocument(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedDocument, val)
}

func (p *PoliciesMapping) GetDocumentRef() (*Ref, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsResolvedDocumentRef)
}

func (p *PoliciesMapping) SetDocumentRef(val *Ref) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedDocumentRef, val)
}

func (p *PoliciesMapping) GetError() (*model.MainError, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *model.MainError](p.GetBackingStore(), PoliciesMappingsResolvedError)
}

func (p *PoliciesMapping) SetError(val *model.MainError) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedError, val)
}

func (p *PoliciesMapping) GetException() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedException)
}

func (p *PoliciesMapping) SetException(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedException, val)
}

func (p *PoliciesMapping) GetExceptionAllowed() (*bool, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](p.GetBackingStore(), PoliciesMappingsResolvedExceptionAllowed)
}

func (p *PoliciesMapping) SetExceptionAllowed(val *bool) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedExceptionAllowed, val)
}

func (p *PoliciesMapping) GetInputStatus() (*InputStatus, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *InputStatus](p.GetBackingStore(), PoliciesMappingsResolvedInputStatus)
}

func (p *PoliciesMapping) SetInputStatus(val *InputStatus) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedInputStatus, val)
}

func (p *PoliciesMapping) GetLastUpdatedBy() (*Ref, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsResolvedLastUpdatedBy)
}

func (p *PoliciesMapping) SetLastUpdatedBy(val *Ref) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedLastUpdatedBy, val)
}

func (p *PoliciesMapping) GetNumber() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedNumber)
}

func (p *PoliciesMapping) SetNumber(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedNumber, val)
}

func (p *PoliciesMapping) GetPolicy() (*Ref, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Ref](p.GetBackingStore(), PoliciesMappingsResolvedPolicy)
}

func (p *PoliciesMapping) SetPolicy(val *Ref) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedPolicy, val)
}

func (p *PoliciesMapping) GetReason() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedReason)
}

func (p *PoliciesMapping) SetReason(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedReason, val)
}

func (p *PoliciesMapping) GetState() (*State, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *State](p.GetBackingStore(), PoliciesMappingsResolvedState)
}

func (p *PoliciesMapping) SetState(val *State) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedState, val)
}

func (p *PoliciesMapping) GetSysClassName() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedSysClassName)
}

func (p *PoliciesMapping) SetSysClassName(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysClassName, val)
}

func (p *PoliciesMapping) GetSysCreatedBy() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedSysCreatedBy)
}

func (p *PoliciesMapping) SetSysCreatedBy(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysCreatedBy, val)
}

func (p *PoliciesMapping) GetSysCreatedOn() (*time.Time, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](p.GetBackingStore(), PoliciesMappingsResolvedSysCreatedOn)
}

func (p *PoliciesMapping) SetSysCreatedOn(val *time.Time) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysCreatedOn, val)
}

func (p *PoliciesMapping) GetSysId() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedSysId)
}

func (p *PoliciesMapping) SetSysId(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysId, val)
}

func (p *PoliciesMapping) GetSysUpdatedBy() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](p.GetBackingStore(), PoliciesMappingsResolvedSysUpdatedBy)
}

func (p *PoliciesMapping) SetSysUpdatedBy(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysUpdatedBy, val)
}

func (p *PoliciesMapping) GetSysUpdatedOn() (*time.Time, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](p.GetBackingStore(), PoliciesMappingsResolvedSysUpdatedOn)
}

func (p *PoliciesMapping) SetSysUpdatedOn(val *time.Time) error {
	return kiota.DefaultBackedModelMutatorFunc(p.GetBackingStore(), PoliciesMappingsResolvedSysUpdatedOn, val)
}
