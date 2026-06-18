package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// Reference represents a link and value pair.
type Reference struct {
	core.BaseModel
}

func NewReference() *Reference {
	return &Reference{BaseModel: *core.NewBaseModel()}
}

func (m *Reference) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
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
	core.BaseModel
}

func NewChangesetActivityResult() *ChangesetActivityResult {
	return &ChangesetActivityResult{BaseModel: *core.NewBaseModel()}
}

func (m *ChangesetActivityResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
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
	core.BaseModel
}

func NewCommitStatusResult() *CommitStatusResult {
	return &CommitStatusResult{BaseModel: *core.NewBaseModel()}
}

func (m *CommitStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
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

// ImpactedSharedComponentResult represents an impacted shared component.
type ImpactedSharedComponentResult struct {
	core.BaseModel
}

func NewImpactedSharedComponentResult() *ImpactedSharedComponentResult {
	return &ImpactedSharedComponentResult{BaseModel: *core.NewBaseModel()}
}

func (m *ImpactedSharedComponentResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(cdmSharedLibraryKey)(m.GetCdmSharedLibrary),
		internalSerialization.SerializeStringFunc(descriptionKey)(m.GetDescription),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(nodeKey)(m.GetNode),
		internalSerialization.SerializeStringFunc(nodeMainKey)(m.GetNodeMain),
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeStringFunc(sysCreatedByKey)(m.GetSysCreatedBy),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey)(m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(sysUpdatedByKey)(m.GetSysUpdatedBy),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey)(m.GetSysUpdatedOn),
		internalSerialization.SerializeInt32Func(versionCounterKey)(m.GetVersionCounter),
	)
}

func (m *ImpactedSharedComponentResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		cdmSharedLibraryKey: internalSerialization.DeserializeStringFunc()(m.setCdmSharedLibrary),
		descriptionKey:      internalSerialization.DeserializeStringFunc()(m.setDescription),
		nameKey:             internalSerialization.DeserializeStringFunc()(m.setName),
		nodeKey:             internalSerialization.DeserializeStringFunc()(m.setNode),
		nodeMainKey:         internalSerialization.DeserializeStringFunc()(m.setNodeMain),
		stateKey:            internalSerialization.DeserializeStringFunc()(m.setState),
		sysCreatedByKey:     internalSerialization.DeserializeStringFunc()(m.setSysCreatedBy),
		sysCreatedOnKey:     internalSerialization.DeserializeStringFunc()(m.setSysCreatedOn),
		sysIdKey:            internalSerialization.DeserializeStringFunc()(m.setSysId),
		sysUpdatedByKey:     internalSerialization.DeserializeStringFunc()(m.setSysUpdatedBy),
		sysUpdatedOnKey:     internalSerialization.DeserializeStringFunc()(m.setSysUpdatedOn),
		versionCounterKey:   internalSerialization.DeserializeInt32Func()(m.setVersionCounter),
	}
}

func (m *ImpactedSharedComponentResult) GetCdmSharedLibrary() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), cdmSharedLibraryKey)
}
func (m *ImpactedSharedComponentResult) setCdmSharedLibrary(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cdmSharedLibraryKey, val)
}
func (m *ImpactedSharedComponentResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), descriptionKey)
}
func (m *ImpactedSharedComponentResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), descriptionKey, val)
}
func (m *ImpactedSharedComponentResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *ImpactedSharedComponentResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}
func (m *ImpactedSharedComponentResult) GetNode() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nodeKey)
}
func (m *ImpactedSharedComponentResult) setNode(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nodeKey, val)
}
func (m *ImpactedSharedComponentResult) GetNodeMain() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nodeMainKey)
}
func (m *ImpactedSharedComponentResult) setNodeMain(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nodeMainKey, val)
}
func (m *ImpactedSharedComponentResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}
func (m *ImpactedSharedComponentResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysCreatedByKey)
}
func (m *ImpactedSharedComponentResult) setSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysCreatedByKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysCreatedOnKey)
}
func (m *ImpactedSharedComponentResult) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysCreatedOnKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *ImpactedSharedComponentResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysUpdatedByKey)
}
func (m *ImpactedSharedComponentResult) setSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysUpdatedByKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysUpdatedOnKey)
}
func (m *ImpactedSharedComponentResult) setSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysUpdatedOnKey, val)
}
func (m *ImpactedSharedComponentResult) GetVersionCounter() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int32](m.GetBackingStore(), versionCounterKey)
}
func (m *ImpactedSharedComponentResult) setVersionCounter(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), versionCounterKey, val)
}

func CreateImpactedSharedComponentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewImpactedSharedComponentResult(), nil
}

// ImpactedDeployableResult represents an impacted deployable (query-based).
type ImpactedDeployableResult struct {
	core.BaseModel
}

func NewImpactedDeployableResult() *ImpactedDeployableResult {
	return &ImpactedDeployableResult{BaseModel: *core.NewBaseModel()}
}

func (m *ImpactedDeployableResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeInt32Func(cdiCountKey)(m.GetCdiCount),
		internalSerialization.SerializeStringFunc(cdiUsageKey)(m.GetCdiUsage),
		internalSerialization.SerializeObjectValueFunc[*Reference](cdmAppKey)(m.GetCdmApp),
		internalSerialization.SerializeObjectValueFunc[*Reference](cdmCiKey)(m.GetCdmCi),
		internalSerialization.SerializeStringFunc(descriptionKey)(m.GetDescription),
		internalSerialization.SerializeStringFunc(environmentTypeKey)(m.GetEnvironmentType),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeObjectValueFunc[*Reference](nodeKey)(m.GetNode),
		internalSerialization.SerializeInt32Func(snapshotVersionCounterKey)(m.GetSnapshotVersionCounter),
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(sysCreatedByKey)(m.GetSysCreatedBy),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey)(m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysUpdatedByKey)(m.GetSysUpdatedBy),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey)(m.GetSysUpdatedOn),
	)
}

func (m *ImpactedDeployableResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		cdiCountKey:               internalSerialization.DeserializeInt32Func()(m.setCdiCount),
		cdiUsageKey:               internalSerialization.DeserializeStringFunc()(m.setCdiUsage),
		cdmAppKey:                 internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue)(m.setCdmApp),
		cdmCiKey:                  internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue)(m.setCdmCi),
		descriptionKey:            internalSerialization.DeserializeStringFunc()(m.setDescription),
		environmentTypeKey:        internalSerialization.DeserializeStringFunc()(m.setEnvironmentType),
		nameKey:                   internalSerialization.DeserializeStringFunc()(m.setName),
		nodeKey:                   internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue)(m.setNode),
		snapshotVersionCounterKey: internalSerialization.DeserializeInt32Func()(m.setSnapshotVersionCounter),
		stateKey:                  internalSerialization.DeserializeStringFunc()(m.setState),
		sysIdKey:                  internalSerialization.DeserializeStringFunc()(m.setSysId),
		sysCreatedByKey:           internalSerialization.DeserializeStringFunc()(m.setSysCreatedBy),
		sysCreatedOnKey:           internalSerialization.DeserializeStringFunc()(m.setSysCreatedOn),
		sysUpdatedByKey:           internalSerialization.DeserializeStringFunc()(m.setSysUpdatedBy),
		sysUpdatedOnKey:           internalSerialization.DeserializeStringFunc()(m.setSysUpdatedOn),
	}
}

func (m *ImpactedDeployableResult) GetCdiCount() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int32](m.GetBackingStore(), cdiCountKey)
}
func (m *ImpactedDeployableResult) setCdiCount(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cdiCountKey, val)
}
func (m *ImpactedDeployableResult) GetCdiUsage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), cdiUsageKey)
}
func (m *ImpactedDeployableResult) setCdiUsage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cdiUsageKey, val)
}
func (m *ImpactedDeployableResult) GetCdmApp() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Reference](m.GetBackingStore(), cdmAppKey)
}
func (m *ImpactedDeployableResult) setCdmApp(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cdmAppKey, val)
}
func (m *ImpactedDeployableResult) GetCdmCi() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Reference](m.GetBackingStore(), cdmCiKey)
}
func (m *ImpactedDeployableResult) setCdmCi(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cdmCiKey, val)
}
func (m *ImpactedDeployableResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), descriptionKey)
}
func (m *ImpactedDeployableResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), descriptionKey, val)
}
func (m *ImpactedDeployableResult) GetEnvironmentType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), environmentTypeKey)
}
func (m *ImpactedDeployableResult) setEnvironmentType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), environmentTypeKey, val)
}
func (m *ImpactedDeployableResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *ImpactedDeployableResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}
func (m *ImpactedDeployableResult) GetNode() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *Reference](m.GetBackingStore(), nodeKey)
}
func (m *ImpactedDeployableResult) setNode(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nodeKey, val)
}
func (m *ImpactedDeployableResult) GetSnapshotVersionCounter() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int32](m.GetBackingStore(), snapshotVersionCounterKey)
}
func (m *ImpactedDeployableResult) setSnapshotVersionCounter(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), snapshotVersionCounterKey, val)
}
func (m *ImpactedDeployableResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}
func (m *ImpactedDeployableResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}
func (m *ImpactedDeployableResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *ImpactedDeployableResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *ImpactedDeployableResult) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysCreatedByKey)
}
func (m *ImpactedDeployableResult) setSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysCreatedByKey, val)
}
func (m *ImpactedDeployableResult) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysCreatedOnKey)
}
func (m *ImpactedDeployableResult) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysCreatedOnKey, val)
}
func (m *ImpactedDeployableResult) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysUpdatedByKey)
}
func (m *ImpactedDeployableResult) setSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysUpdatedByKey, val)
}
func (m *ImpactedDeployableResult) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysUpdatedOnKey)
}
func (m *ImpactedDeployableResult) setSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysUpdatedOnKey, val)
}

func CreateImpactedDeployableResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewImpactedDeployableResult(), nil
}

// ImpactedDeployableBySysIdResult represents an impacted deployable (path-based).
type ImpactedDeployableBySysIdResult struct {
	core.BaseModel
}

func NewImpactedDeployableBySysIdResult() *ImpactedDeployableBySysIdResult {
	return &ImpactedDeployableBySysIdResult{BaseModel: *core.NewBaseModel()}
}

func (m *ImpactedDeployableBySysIdResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(changesetIdKey)(m.GetChangesetId),
		internalSerialization.SerializeBoolFunc(conflictKey)(m.GetConflict),
		internalSerialization.SerializeStringFunc(conflictTypeKey)(m.GetConflictType),
		internalSerialization.SerializeStringFunc(descriptionKey)(m.GetDescription),
		internalSerialization.SerializeStringFunc(effectiveFromKey)(m.GetEffectiveFrom),
		internalSerialization.SerializeStringFunc(effectiveToKey)(m.GetEffectiveTo),
		internalSerialization.SerializeInt32Func(levelKey)(m.GetLevel),
		internalSerialization.SerializeStringFunc(linkedToKey)(m.GetLinkedTo),
		internalSerialization.SerializeStringFunc(mainIdKey)(m.GetMainId),
		internalSerialization.SerializeStringFunc(mainIdEncodedKey)(m.GetMainIdEncoded),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(nodeClassifierKey)(m.GetNodeClassifier),
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
		internalSerialization.SerializeStringFunc(secureValueKey)(m.GetSecureValue),
	)
}

func (m *ImpactedDeployableBySysIdResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		changesetIdKey:    internalSerialization.DeserializeStringFunc()(m.setChangesetId),
		conflictKey:       internalSerialization.DeserializeBoolFunc()(m.setConflict),
		conflictTypeKey:   internalSerialization.DeserializeStringFunc()(m.setConflictType),
		descriptionKey:    internalSerialization.DeserializeStringFunc()(m.setDescription),
		effectiveFromKey:  internalSerialization.DeserializeStringFunc()(m.setEffectiveFrom),
		effectiveToKey:    internalSerialization.DeserializeStringFunc()(m.setEffectiveTo),
		levelKey:          internalSerialization.DeserializeInt32Func()(m.setLevel),
		linkedToKey:       internalSerialization.DeserializeStringFunc()(m.setLinkedTo),
		mainIdKey:         internalSerialization.DeserializeStringFunc()(m.setMainId),
		mainIdEncodedKey:  internalSerialization.DeserializeStringFunc()(m.setMainIdEncoded),
		nameKey:           internalSerialization.DeserializeStringFunc()(m.setName),
		nodeClassifierKey: internalSerialization.DeserializeStringFunc()(m.setNodeClassifier),
		statusKey:         internalSerialization.DeserializeStringFunc()(m.setStatus),
		sysIdKey:          internalSerialization.DeserializeStringFunc()(m.setSysId),
		typeKey:           internalSerialization.DeserializeStringFunc()(m.setType),
		valueKey:          internalSerialization.DeserializeStringFunc()(m.setValue),
		secureValueKey:    internalSerialization.DeserializeStringFunc()(m.setSecureValue),
	}
}

func (m *ImpactedDeployableBySysIdResult) GetChangesetId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), changesetIdKey)
}
func (m *ImpactedDeployableBySysIdResult) setChangesetId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), changesetIdKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetConflict() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), conflictKey)
}
func (m *ImpactedDeployableBySysIdResult) setConflict(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), conflictKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetConflictType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), conflictTypeKey)
}
func (m *ImpactedDeployableBySysIdResult) setConflictType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), conflictTypeKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), descriptionKey)
}
func (m *ImpactedDeployableBySysIdResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), descriptionKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetEffectiveFrom() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), effectiveFromKey)
}
func (m *ImpactedDeployableBySysIdResult) setEffectiveFrom(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), effectiveFromKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetEffectiveTo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), effectiveToKey)
}
func (m *ImpactedDeployableBySysIdResult) setEffectiveTo(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), effectiveToKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetLevel() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int32](m.GetBackingStore(), levelKey)
}
func (m *ImpactedDeployableBySysIdResult) setLevel(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), levelKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetLinkedTo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), linkedToKey)
}
func (m *ImpactedDeployableBySysIdResult) setLinkedTo(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), linkedToKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetMainId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), mainIdKey)
}
func (m *ImpactedDeployableBySysIdResult) setMainId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), mainIdKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetMainIdEncoded() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), mainIdEncodedKey)
}
func (m *ImpactedDeployableBySysIdResult) setMainIdEncoded(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), mainIdEncodedKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *ImpactedDeployableBySysIdResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetNodeClassifier() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nodeClassifierKey)
}
func (m *ImpactedDeployableBySysIdResult) setNodeClassifier(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nodeClassifierKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), statusKey)
}
func (m *ImpactedDeployableBySysIdResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), statusKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *ImpactedDeployableBySysIdResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}
func (m *ImpactedDeployableBySysIdResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), valueKey)
}
func (m *ImpactedDeployableBySysIdResult) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetSecureValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), secureValueKey)
}
func (m *ImpactedDeployableBySysIdResult) setSecureValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), secureValueKey, val)
}

func CreateImpactedDeployableBySysIdResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewImpactedDeployableBySysIdResult(), nil
}
