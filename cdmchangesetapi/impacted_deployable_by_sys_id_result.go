package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImpactedDeployableBySysIDResult represents an impacted deployable (path-based).
type ImpactedDeployableBySysIDResult struct {
	core.BaseModel
}

// NewImpactedDeployableBySysIDResult instantiates a new ImpactedDeployableBySysIDResult.
func NewImpactedDeployableBySysIDResult() *ImpactedDeployableBySysIDResult {
	return &ImpactedDeployableBySysIDResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *ImpactedDeployableBySysIDResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(changesetIDKey, m.GetChangesetID),
		internalSerialization.SerializeBoolFunc(conflictKey, m.GetConflict),
		internalSerialization.SerializeStringFunc(conflictTypeKey, m.GetConflictType),
		internalSerialization.SerializeStringFunc(descriptionKey, m.GetDescription),
		internalSerialization.SerializeStringFunc(effectiveFromKey, m.GetEffectiveFrom),
		internalSerialization.SerializeStringFunc(effectiveToKey, m.GetEffectiveTo),
		internalSerialization.SerializeInt32Func(levelKey, m.GetLevel),
		internalSerialization.SerializeStringFunc(linkedToKey, m.GetLinkedTo),
		internalSerialization.SerializeStringFunc(mainIDKey, m.GetMainID),
		internalSerialization.SerializeStringFunc(mainIDEncodedKey, m.GetMainIDEncoded),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(nodeClassifierKey, m.GetNodeClassifier),
		internalSerialization.SerializeStringFunc(statusKey, m.GetStatus),
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
		internalSerialization.SerializeStringFunc(secureValueKey, m.GetSecureValue),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *ImpactedDeployableBySysIDResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		changesetIDKey:    internalSerialization.DeserializeStringFunc(m.setChangesetID),
		conflictKey:       internalSerialization.DeserializeBoolFunc(m.setConflict),
		conflictTypeKey:   internalSerialization.DeserializeStringFunc(m.setConflictType),
		descriptionKey:    internalSerialization.DeserializeStringFunc(m.setDescription),
		effectiveFromKey:  internalSerialization.DeserializeStringFunc(m.setEffectiveFrom),
		effectiveToKey:    internalSerialization.DeserializeStringFunc(m.setEffectiveTo),
		levelKey:          internalSerialization.DeserializeInt32Func(m.setLevel),
		linkedToKey:       internalSerialization.DeserializeStringFunc(m.setLinkedTo),
		mainIDKey:         internalSerialization.DeserializeStringFunc(m.setMainID),
		mainIDEncodedKey:  internalSerialization.DeserializeStringFunc(m.setMainIDEncoded),
		nameKey:           internalSerialization.DeserializeStringFunc(m.setName),
		nodeClassifierKey: internalSerialization.DeserializeStringFunc(m.setNodeClassifier),
		statusKey:         internalSerialization.DeserializeStringFunc(m.setStatus),
		sysIDKey:          internalSerialization.DeserializeStringFunc(m.setSysID),
		typeKey:           internalSerialization.DeserializeStringFunc(m.setType),
		valueKey:          internalSerialization.DeserializeStringFunc(m.setValue),
		secureValueKey:    internalSerialization.DeserializeStringFunc(m.setSecureValue),
	}
}

// GetChangesetID returns the changeset id.
func (m *ImpactedDeployableBySysIDResult) GetChangesetID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, changesetIDKey)
}
func (m *ImpactedDeployableBySysIDResult) setChangesetID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, changesetIDKey, val)
}

// GetConflict returns the conflict.
func (m *ImpactedDeployableBySysIDResult) GetConflict() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *bool](m, conflictKey)
}
func (m *ImpactedDeployableBySysIDResult) setConflict(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, conflictKey, val)
}

// GetConflictType returns the conflict type.
func (m *ImpactedDeployableBySysIDResult) GetConflictType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, conflictTypeKey)
}
func (m *ImpactedDeployableBySysIDResult) setConflictType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, conflictTypeKey, val)
}

// GetDescription returns the description.
func (m *ImpactedDeployableBySysIDResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, descriptionKey)
}
func (m *ImpactedDeployableBySysIDResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}

// GetEffectiveFrom returns the effective from.
func (m *ImpactedDeployableBySysIDResult) GetEffectiveFrom() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, effectiveFromKey)
}
func (m *ImpactedDeployableBySysIDResult) setEffectiveFrom(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, effectiveFromKey, val)
}

// GetEffectiveTo returns the effective to.
func (m *ImpactedDeployableBySysIDResult) GetEffectiveTo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, effectiveToKey)
}
func (m *ImpactedDeployableBySysIDResult) setEffectiveTo(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, effectiveToKey, val)
}

// GetLevel returns the level.
func (m *ImpactedDeployableBySysIDResult) GetLevel() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *int32](m, levelKey)
}
func (m *ImpactedDeployableBySysIDResult) setLevel(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, levelKey, val)
}

// GetLinkedTo returns the linked to.
func (m *ImpactedDeployableBySysIDResult) GetLinkedTo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, linkedToKey)
}
func (m *ImpactedDeployableBySysIDResult) setLinkedTo(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, linkedToKey, val)
}

// GetMainID returns the main id.
func (m *ImpactedDeployableBySysIDResult) GetMainID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, mainIDKey)
}
func (m *ImpactedDeployableBySysIDResult) setMainID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, mainIDKey, val)
}

// GetMainIDEncoded returns the main id encoded.
func (m *ImpactedDeployableBySysIDResult) GetMainIDEncoded() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, mainIDEncodedKey)
}
func (m *ImpactedDeployableBySysIDResult) setMainIDEncoded(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, mainIDEncodedKey, val)
}

// GetName returns the name.
func (m *ImpactedDeployableBySysIDResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, nameKey)
}
func (m *ImpactedDeployableBySysIDResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetNodeClassifier returns the node classifier.
func (m *ImpactedDeployableBySysIDResult) GetNodeClassifier() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, nodeClassifierKey)
}
func (m *ImpactedDeployableBySysIDResult) setNodeClassifier(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nodeClassifierKey, val)
}

// GetStatus returns the status.
func (m *ImpactedDeployableBySysIDResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, statusKey)
}
func (m *ImpactedDeployableBySysIDResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}

// GetSysID returns the sys id.
func (m *ImpactedDeployableBySysIDResult) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, sysIDKey)
}
func (m *ImpactedDeployableBySysIDResult) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetType returns the type.
func (m *ImpactedDeployableBySysIDResult) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, typeKey)
}
func (m *ImpactedDeployableBySysIDResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetValue returns the value.
func (m *ImpactedDeployableBySysIDResult) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, valueKey)
}
func (m *ImpactedDeployableBySysIDResult) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

// GetSecureValue returns the secure value.
func (m *ImpactedDeployableBySysIDResult) GetSecureValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIDResult, *string](m, secureValueKey)
}
func (m *ImpactedDeployableBySysIDResult) setSecureValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, secureValueKey, val)
}

// CreateImpactedDeployableBySysIDResultFromDiscriminatorValue creates a new ImpactedDeployableBySysIDResult from a ParseNode.
func CreateImpactedDeployableBySysIDResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewImpactedDeployableBySysIDResult(), nil
}
