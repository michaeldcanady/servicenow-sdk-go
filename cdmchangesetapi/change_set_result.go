package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangesetResult represents a changeset.
type ChangesetResult struct {
	core.BaseModel
}

func NewChangesetResult() *ChangesetResult {
	return &ChangesetResult{BaseModel: *core.NewBaseModel()}
}

func (m *ChangesetResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(autoValidateKey, m.GetAutoValidate),
		internalSerialization.SerializeObjectValueFunc[*Reference](cdmApplicationKey, m.GetCdmApplication),
		internalSerialization.SerializeStringFunc(committedAtKey, m.GetCommittedAt),
		internalSerialization.SerializeObjectValueFunc[*Reference](committedByKey, m.GetCommittedBy),
		internalSerialization.SerializeStringFunc(descriptionKey, m.GetDescription),
		internalSerialization.SerializeInt64Func(lastConflictDetectionTimeKey, m.GetLastConflictDetectionTime),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
		internalSerialization.SerializeStringFunc(publishOptionKey, m.GetPublishOption),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(sysIdKey, m.GetSysId),
		internalSerialization.SerializeStringFunc(titleKey, m.GetTitle),
	)
}

func (m *ChangesetResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		autoValidateKey:              internalSerialization.DeserializeBoolFunc(m.setAutoValidate),
		cdmApplicationKey:            internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue, m.setCdmApplication),
		committedAtKey:               internalSerialization.DeserializeStringFunc(m.setCommittedAt),
		committedByKey:               internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue, m.setCommittedBy),
		descriptionKey:               internalSerialization.DeserializeStringFunc(m.setDescription),
		lastConflictDetectionTimeKey: internalSerialization.DeserializeInt64Func(m.setLastConflictDetectionTime),
		numberKey:                    internalSerialization.DeserializeStringFunc(m.setNumber),
		publishOptionKey:             internalSerialization.DeserializeStringFunc(m.setPublishOption),
		stateKey:                     internalSerialization.DeserializeStringFunc(m.setState),
		sysIdKey:                     internalSerialization.DeserializeStringFunc(m.setSysId),
		titleKey:                     internalSerialization.DeserializeStringFunc(m.setTitle),
	}
}

func (m *ChangesetResult) GetAutoValidate() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *bool](m, autoValidateKey)
}
func (m *ChangesetResult) setAutoValidate(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, autoValidateKey, val)
}
func (m *ChangesetResult) GetCdmApplication() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *Reference](m, cdmApplicationKey)
}
func (m *ChangesetResult) setCdmApplication(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmApplicationKey, val)
}
func (m *ChangesetResult) GetCommittedAt() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *string](m, committedAtKey)
}
func (m *ChangesetResult) setCommittedAt(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, committedAtKey, val)
}
func (m *ChangesetResult) GetCommittedBy() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *Reference](m, committedByKey)
}
func (m *ChangesetResult) setCommittedBy(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, committedByKey, val)
}
func (m *ChangesetResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *string](m, descriptionKey)
}
func (m *ChangesetResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}
func (m *ChangesetResult) GetLastConflictDetectionTime() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *int64](m, lastConflictDetectionTimeKey)
}
func (m *ChangesetResult) setLastConflictDetectionTime(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(m, lastConflictDetectionTimeKey, val)
}
func (m *ChangesetResult) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *string](m, numberKey)
}
func (m *ChangesetResult) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}
func (m *ChangesetResult) GetPublishOption() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *string](m, publishOptionKey)
}
func (m *ChangesetResult) setPublishOption(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, publishOptionKey, val)
}
func (m *ChangesetResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *string](m, stateKey)
}
func (m *ChangesetResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}
func (m *ChangesetResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *string](m, sysIdKey)
}
func (m *ChangesetResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}
func (m *ChangesetResult) GetTitle() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetResult, *string](m, titleKey)
}
func (m *ChangesetResult) setTitle(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, titleKey, val)
}

func CreateChangesetResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewChangesetResult(), nil
}
