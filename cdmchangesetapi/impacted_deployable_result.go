package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

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
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *int32](m, cdiCountKey)
}
func (m *ImpactedDeployableResult) setCdiCount(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, cdiCountKey, val)
}
func (m *ImpactedDeployableResult) GetCdiUsage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, cdiUsageKey)
}
func (m *ImpactedDeployableResult) setCdiUsage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cdiUsageKey, val)
}
func (m *ImpactedDeployableResult) GetCdmApp() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *Reference](m, cdmAppKey)
}
func (m *ImpactedDeployableResult) setCdmApp(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmAppKey, val)
}
func (m *ImpactedDeployableResult) GetCdmCi() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *Reference](m, cdmCiKey)
}
func (m *ImpactedDeployableResult) setCdmCi(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmCiKey, val)
}
func (m *ImpactedDeployableResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, descriptionKey)
}
func (m *ImpactedDeployableResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}
func (m *ImpactedDeployableResult) GetEnvironmentType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, environmentTypeKey)
}
func (m *ImpactedDeployableResult) setEnvironmentType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, environmentTypeKey, val)
}
func (m *ImpactedDeployableResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, nameKey)
}
func (m *ImpactedDeployableResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
func (m *ImpactedDeployableResult) GetNode() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *Reference](m, nodeKey)
}
func (m *ImpactedDeployableResult) setNode(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, nodeKey, val)
}
func (m *ImpactedDeployableResult) GetSnapshotVersionCounter() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *int32](m, snapshotVersionCounterKey)
}
func (m *ImpactedDeployableResult) setSnapshotVersionCounter(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, snapshotVersionCounterKey, val)
}
func (m *ImpactedDeployableResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, stateKey)
}
func (m *ImpactedDeployableResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}
func (m *ImpactedDeployableResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, sysIdKey)
}
func (m *ImpactedDeployableResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}
func (m *ImpactedDeployableResult) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, sysCreatedByKey)
}
func (m *ImpactedDeployableResult) setSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedByKey, val)
}
func (m *ImpactedDeployableResult) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, sysCreatedOnKey)
}
func (m *ImpactedDeployableResult) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}
func (m *ImpactedDeployableResult) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, sysUpdatedByKey)
}
func (m *ImpactedDeployableResult) setSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedByKey, val)
}
func (m *ImpactedDeployableResult) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedDeployableResult, *string](m, sysUpdatedOnKey)
}
func (m *ImpactedDeployableResult) setSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedOnKey, val)
}

func CreateImpactedDeployableResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewImpactedDeployableResult(), nil
}
