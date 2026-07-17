package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

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
		internalSerialization.SerializeStringFunc(changesetIdKey, m.GetChangesetId),
		internalSerialization.SerializeBoolFunc(conflictKey, m.GetConflict),
		internalSerialization.SerializeStringFunc(conflictTypeKey, m.GetConflictType),
		internalSerialization.SerializeStringFunc(descriptionKey, m.GetDescription),
		internalSerialization.SerializeStringFunc(effectiveFromKey, m.GetEffectiveFrom),
		internalSerialization.SerializeStringFunc(effectiveToKey, m.GetEffectiveTo),
		internalSerialization.SerializeInt32Func(levelKey, m.GetLevel),
		internalSerialization.SerializeStringFunc(linkedToKey, m.GetLinkedTo),
		internalSerialization.SerializeStringFunc(mainIdKey, m.GetMainId),
		internalSerialization.SerializeStringFunc(mainIdEncodedKey, m.GetMainIdEncoded),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(nodeClassifierKey, m.GetNodeClassifier),
		internalSerialization.SerializeStringFunc(statusKey, m.GetStatus),
		internalSerialization.SerializeStringFunc(sysIdKey, m.GetSysId),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
		internalSerialization.SerializeStringFunc(secureValueKey, m.GetSecureValue),
	)
}

func (m *ImpactedDeployableBySysIdResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		changesetIdKey:    internalSerialization.DeserializeStringFunc(m.setChangesetId),
		conflictKey:       internalSerialization.DeserializeBoolFunc(m.setConflict),
		conflictTypeKey:   internalSerialization.DeserializeStringFunc(m.setConflictType),
		descriptionKey:    internalSerialization.DeserializeStringFunc(m.setDescription),
		effectiveFromKey:  internalSerialization.DeserializeStringFunc(m.setEffectiveFrom),
		effectiveToKey:    internalSerialization.DeserializeStringFunc(m.setEffectiveTo),
		levelKey:          internalSerialization.DeserializeInt32Func(m.setLevel),
		linkedToKey:       internalSerialization.DeserializeStringFunc(m.setLinkedTo),
		mainIdKey:         internalSerialization.DeserializeStringFunc(m.setMainId),
		mainIdEncodedKey:  internalSerialization.DeserializeStringFunc(m.setMainIdEncoded),
		nameKey:           internalSerialization.DeserializeStringFunc(m.setName),
		nodeClassifierKey: internalSerialization.DeserializeStringFunc(m.setNodeClassifier),
		statusKey:         internalSerialization.DeserializeStringFunc(m.setStatus),
		sysIdKey:          internalSerialization.DeserializeStringFunc(m.setSysId),
		typeKey:           internalSerialization.DeserializeStringFunc(m.setType),
		valueKey:          internalSerialization.DeserializeStringFunc(m.setValue),
		secureValueKey:    internalSerialization.DeserializeStringFunc(m.setSecureValue),
	}
}

func (m *ImpactedDeployableBySysIdResult) GetChangesetId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, changesetIdKey)
}
func (m *ImpactedDeployableBySysIdResult) setChangesetId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, changesetIdKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetConflict() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *bool](m, conflictKey)
}
func (m *ImpactedDeployableBySysIdResult) setConflict(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, conflictKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetConflictType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, conflictTypeKey)
}
func (m *ImpactedDeployableBySysIdResult) setConflictType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, conflictTypeKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, descriptionKey)
}
func (m *ImpactedDeployableBySysIdResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetEffectiveFrom() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, effectiveFromKey)
}
func (m *ImpactedDeployableBySysIdResult) setEffectiveFrom(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, effectiveFromKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetEffectiveTo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, effectiveToKey)
}
func (m *ImpactedDeployableBySysIdResult) setEffectiveTo(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, effectiveToKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetLevel() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *int32](m, levelKey)
}
func (m *ImpactedDeployableBySysIdResult) setLevel(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, levelKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetLinkedTo() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, linkedToKey)
}
func (m *ImpactedDeployableBySysIdResult) setLinkedTo(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, linkedToKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetMainId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, mainIdKey)
}
func (m *ImpactedDeployableBySysIdResult) setMainId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, mainIdKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetMainIdEncoded() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, mainIdEncodedKey)
}
func (m *ImpactedDeployableBySysIdResult) setMainIdEncoded(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, mainIdEncodedKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, nameKey)
}
func (m *ImpactedDeployableBySysIdResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetNodeClassifier() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, nodeClassifierKey)
}
func (m *ImpactedDeployableBySysIdResult) setNodeClassifier(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nodeClassifierKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, statusKey)
}
func (m *ImpactedDeployableBySysIdResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, sysIdKey)
}
func (m *ImpactedDeployableBySysIdResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, typeKey)
}
func (m *ImpactedDeployableBySysIdResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, valueKey)
}
func (m *ImpactedDeployableBySysIdResult) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}
func (m *ImpactedDeployableBySysIdResult) GetSecureValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableBySysIdResult, *string](m, secureValueKey)
}
func (m *ImpactedDeployableBySysIdResult) setSecureValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, secureValueKey, val)
}

func CreateImpactedDeployableBySysIdResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewImpactedDeployableBySysIdResult(), nil
}
