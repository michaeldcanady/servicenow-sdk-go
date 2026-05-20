package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// Reference represents a link and value pair.
type Reference struct {
	newInternal.BaseModel
}

func NewReference() *Reference {
	return &Reference{BaseModel: *newInternal.NewBaseModel()}
}

func (m *Reference) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(linkKey)(m.GetLink),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
	)
}

func (m *Reference) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		linkKey:  internalSerialization.DeserializeStringFunc()(m.setLink),
		valueKey: internalSerialization.DeserializeStringFunc()(m.setValue),
	}
}

func (m *Reference) GetLink() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), linkKey)
}
func (m *Reference) setLink(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), linkKey, val)
}
func (m *Reference) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), valueKey)
}
func (m *Reference) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, val)
}

func CreateReferenceFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewReference(), nil
}

// ChangesetResult represents a changeset.
type ChangesetResult struct {
	newInternal.BaseModel
}

func NewChangesetResult() *ChangesetResult {
	return &ChangesetResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *ChangesetResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(autoValidateKey)(m.GetAutoValidate),
		internalSerialization.SerializeObjectValueFunc[*Reference](cdmApplicationKey)(m.GetCdmApplication),
		internalSerialization.SerializeStringFunc(committedAtKey)(m.GetCommittedAt),
		internalSerialization.SerializeObjectValueFunc[*Reference](committedByKey)(m.GetCommittedBy),
		internalSerialization.SerializeStringFunc(descriptionKey)(m.GetDescription),
		internalSerialization.SerializeInt64Func(lastConflictDetectionTimeKey)(m.GetLastConflictDetectionTime),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
		internalSerialization.SerializeStringFunc(publishOptionKey)(m.GetPublishOption),
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(titleKey)(m.GetTitle),
	)
}

func (m *ChangesetResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		autoValidateKey:              internalSerialization.DeserializeBoolFunc()(m.setAutoValidate),
		cdmApplicationKey:            internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue)(m.setCdmApplication),
		committedAtKey:               internalSerialization.DeserializeStringFunc()(m.setCommittedAt),
		committedByKey:               internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue)(m.setCommittedBy),
		descriptionKey:               internalSerialization.DeserializeStringFunc()(m.setDescription),
		lastConflictDetectionTimeKey: internalSerialization.DeserializeInt64Func()(m.setLastConflictDetectionTime),
		numberKey:                    internalSerialization.DeserializeStringFunc()(m.setNumber),
		publishOptionKey:             internalSerialization.DeserializeStringFunc()(m.setPublishOption),
		stateKey:                     internalSerialization.DeserializeStringFunc()(m.setState),
		sysIdKey:                     internalSerialization.DeserializeStringFunc()(m.setSysId),
		titleKey:                     internalSerialization.DeserializeStringFunc()(m.setTitle),
	}
}

func (m *ChangesetResult) GetAutoValidate() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), autoValidateKey)
}
func (m *ChangesetResult) setAutoValidate(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), autoValidateKey, val)
}
func (m *ChangesetResult) GetCdmApplication() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Reference](m.GetBackingStore(), cdmApplicationKey)
}
func (m *ChangesetResult) setCdmApplication(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cdmApplicationKey, val)
}
func (m *ChangesetResult) GetCommittedAt() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), committedAtKey)
}
func (m *ChangesetResult) setCommittedAt(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), committedAtKey, val)
}
func (m *ChangesetResult) GetCommittedBy() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Reference](m.GetBackingStore(), committedByKey)
}
func (m *ChangesetResult) setCommittedBy(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), committedByKey, val)
}
func (m *ChangesetResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), descriptionKey)
}
func (m *ChangesetResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), descriptionKey, val)
}
func (m *ChangesetResult) GetLastConflictDetectionTime() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](m.GetBackingStore(), lastConflictDetectionTimeKey)
}
func (m *ChangesetResult) setLastConflictDetectionTime(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), lastConflictDetectionTimeKey, val)
}
func (m *ChangesetResult) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), numberKey)
}
func (m *ChangesetResult) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), numberKey, val)
}
func (m *ChangesetResult) GetPublishOption() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), publishOptionKey)
}
func (m *ChangesetResult) setPublishOption(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), publishOptionKey, val)
}
func (m *ChangesetResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}
func (m *ChangesetResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}
func (m *ChangesetResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *ChangesetResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *ChangesetResult) GetTitle() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), titleKey)
}
func (m *ChangesetResult) setTitle(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), titleKey, val)
}

func CreateChangesetResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewChangesetResult(), nil
}

// ChangesetActivityResult represents a changeset activity.
type ChangesetActivityResult struct {
	newInternal.BaseModel
}

func NewChangesetActivityResult() *ChangesetActivityResult {
	return &ChangesetActivityResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *ChangesetActivityResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*Reference](changesetIdKey)(m.GetChangesetId),
		internalSerialization.SerializeBoolFunc(conflictKey)(m.GetConflict),
		internalSerialization.SerializeStringFunc(namePathKey)(m.GetNamePath),
		internalSerialization.SerializeStringFunc(newNameKey)(m.GetNewName),
		internalSerialization.SerializeStringFunc(oldNameKey)(m.GetOldName),
		internalSerialization.SerializeStringFunc(newValueKey)(m.GetNewValue),
		internalSerialization.SerializeStringFunc(oldValueKey)(m.GetOldValue),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeBoolFunc(secureKey)(m.GetSecure),
	)
}

func (m *ChangesetActivityResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		changesetIdKey: internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue)(m.setChangesetId),
		conflictKey:    internalSerialization.DeserializeBoolFunc()(m.setConflict),
		namePathKey:    internalSerialization.DeserializeStringFunc()(m.setNamePath),
		newNameKey:     internalSerialization.DeserializeStringFunc()(m.setNewName),
		oldNameKey:     internalSerialization.DeserializeStringFunc()(m.setOldName),
		newValueKey:    internalSerialization.DeserializeStringFunc()(m.setNewValue),
		oldValueKey:    internalSerialization.DeserializeStringFunc()(m.setOldValue),
		typeKey:        internalSerialization.DeserializeStringFunc()(m.setType),
		secureKey:      internalSerialization.DeserializeBoolFunc()(m.setSecure),
	}
}

func (m *ChangesetActivityResult) GetChangesetId() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Reference](m.GetBackingStore(), changesetIdKey)
}
func (m *ChangesetActivityResult) setChangesetId(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), changesetIdKey, val)
}
func (m *ChangesetActivityResult) GetConflict() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), conflictKey)
}
func (m *ChangesetActivityResult) setConflict(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), conflictKey, val)
}
func (m *ChangesetActivityResult) GetNamePath() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), namePathKey)
}
func (m *ChangesetActivityResult) setNamePath(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), namePathKey, val)
}
func (m *ChangesetActivityResult) GetNewName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), newNameKey)
}
func (m *ChangesetActivityResult) setNewName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), newNameKey, val)
}
func (m *ChangesetActivityResult) GetOldName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), oldNameKey)
}
func (m *ChangesetActivityResult) setOldName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), oldNameKey, val)
}
func (m *ChangesetActivityResult) GetNewValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), newValueKey)
}
func (m *ChangesetActivityResult) setNewValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), newValueKey, val)
}
func (m *ChangesetActivityResult) GetOldValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), oldValueKey)
}
func (m *ChangesetActivityResult) setOldValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), oldValueKey, val)
}
func (m *ChangesetActivityResult) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}
func (m *ChangesetActivityResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}
func (m *ChangesetActivityResult) GetSecure() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), secureKey)
}
func (m *ChangesetActivityResult) setSecure(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), secureKey, val)
}

func CreateChangesetActivityResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewChangesetActivityResult(), nil
}

// CommitStatusResult represents a commit status.
type CommitStatusResult struct {
	newInternal.BaseModel
}

func NewCommitStatusResult() *CommitStatusResult {
	return &CommitStatusResult{BaseModel: *newInternal.NewBaseModel()}
}

func (m *CommitStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
	)
}

func (m *CommitStatusResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		stateKey: internalSerialization.DeserializeStringFunc()(m.setState),
	}
}

func (m *CommitStatusResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}
func (m *CommitStatusResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}

func CreateCommitStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCommitStatusResult(), nil
}
