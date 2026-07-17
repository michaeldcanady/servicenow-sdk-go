package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

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
		internalSerialization.SerializeStringFunc(cdmSharedLibraryKey, m.GetCdmSharedLibrary),
		internalSerialization.SerializeStringFunc(descriptionKey, m.GetDescription),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(nodeKey, m.GetNode),
		internalSerialization.SerializeStringFunc(nodeMainKey, m.GetNodeMain),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(sysCreatedByKey, m.GetSysCreatedBy),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey, m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysIdKey, m.GetSysId),
		internalSerialization.SerializeStringFunc(sysUpdatedByKey, m.GetSysUpdatedBy),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey, m.GetSysUpdatedOn),
		internalSerialization.SerializeInt32Func(versionCounterKey, m.GetVersionCounter),
	)
}

func (m *ImpactedSharedComponentResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		cdmSharedLibraryKey: internalSerialization.DeserializeStringFunc(m.setCdmSharedLibrary),
		descriptionKey:      internalSerialization.DeserializeStringFunc(m.setDescription),
		nameKey:             internalSerialization.DeserializeStringFunc(m.setName),
		nodeKey:             internalSerialization.DeserializeStringFunc(m.setNode),
		nodeMainKey:         internalSerialization.DeserializeStringFunc(m.setNodeMain),
		stateKey:            internalSerialization.DeserializeStringFunc(m.setState),
		sysCreatedByKey:     internalSerialization.DeserializeStringFunc(m.setSysCreatedBy),
		sysCreatedOnKey:     internalSerialization.DeserializeStringFunc(m.setSysCreatedOn),
		sysIdKey:            internalSerialization.DeserializeStringFunc(m.setSysId),
		sysUpdatedByKey:     internalSerialization.DeserializeStringFunc(m.setSysUpdatedBy),
		sysUpdatedOnKey:     internalSerialization.DeserializeStringFunc(m.setSysUpdatedOn),
		versionCounterKey:   internalSerialization.DeserializeInt32Func(m.setVersionCounter),
	}
}

func (m *ImpactedSharedComponentResult) GetCdmSharedLibrary() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, cdmSharedLibraryKey)
}
func (m *ImpactedSharedComponentResult) setCdmSharedLibrary(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmSharedLibraryKey, val)
}
func (m *ImpactedSharedComponentResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, descriptionKey)
}
func (m *ImpactedSharedComponentResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}
func (m *ImpactedSharedComponentResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, nameKey)
}
func (m *ImpactedSharedComponentResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
func (m *ImpactedSharedComponentResult) GetNode() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, nodeKey)
}
func (m *ImpactedSharedComponentResult) setNode(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nodeKey, val)
}
func (m *ImpactedSharedComponentResult) GetNodeMain() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, nodeMainKey)
}
func (m *ImpactedSharedComponentResult) setNodeMain(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nodeMainKey, val)
}
func (m *ImpactedSharedComponentResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, stateKey)
}
func (m *ImpactedSharedComponentResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysCreatedByKey)
}
func (m *ImpactedSharedComponentResult) setSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedByKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysCreatedOnKey)
}
func (m *ImpactedSharedComponentResult) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysIdKey)
}
func (m *ImpactedSharedComponentResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysUpdatedByKey)
}
func (m *ImpactedSharedComponentResult) setSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedByKey, val)
}
func (m *ImpactedSharedComponentResult) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysUpdatedOnKey)
}
func (m *ImpactedSharedComponentResult) setSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedOnKey, val)
}
func (m *ImpactedSharedComponentResult) GetVersionCounter() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *int32](m, versionCounterKey)
}
func (m *ImpactedSharedComponentResult) setVersionCounter(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, versionCounterKey, val)
}

func CreateImpactedSharedComponentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewImpactedSharedComponentResult(), nil
}
