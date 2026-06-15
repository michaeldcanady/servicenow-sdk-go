package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// UploadStatusOutput represents the output metadata from an upload.
type UploadStatusOutput struct {
	internal.BaseModel
}

// NewUploadStatusOutput instantiates a new UploadStatusOutput.
func NewUploadStatusOutput() *UploadStatusOutput {
	return &UploadStatusOutput{BaseModel: *internal.NewBaseModel()}
}

// Serialize serializes information the current object.
func (m *UploadStatusOutput) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
	)
}

// GetFieldDeserializers the deserialization information for the current model.
func (m *UploadStatusOutput) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:  internalSerialization.DeserializeStringFunc()(m.setSysId),
		numberKey: internalSerialization.DeserializeStringFunc()(m.setNumber),
	}
}

// GetSysId gets the sys_id property value.
func (m *UploadStatusOutput) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}

// setSysId sets the sys_id property value.
func (m *UploadStatusOutput) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}

// GetNumber gets the number property value.
func (m *UploadStatusOutput) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), numberKey)
}

// setNumber sets the number property value.
func (m *UploadStatusOutput) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), numberKey, val)
}

// CreateUploadStatusOutputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value.
func CreateUploadStatusOutputFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUploadStatusOutput(), nil
}

// UploadStatusResult represents the status response of an upload.
type UploadStatusResult struct {
	internal.BaseModel
}

// NewUploadStatusResult instantiates a new UploadStatusResult.
func NewUploadStatusResult() *UploadStatusResult {
	return &UploadStatusResult{BaseModel: *internal.NewBaseModel()}
}

// Serialize serializes information the current object.
func (m *UploadStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeObjectValueFunc[*UploadStatusOutput](outputKey)(m.GetOutput),
	)
}

// GetFieldDeserializers the deserialization information for the current model.
func (m *UploadStatusResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		typeKey:   internalSerialization.DeserializeStringFunc()(m.setType),
		stateKey:  internalSerialization.DeserializeStringFunc()(m.setState),
		outputKey: internalSerialization.DeserializeObjectValueFunc[*UploadStatusOutput](CreateUploadStatusOutputFromDiscriminatorValue)(m.setOutput),
	}
}

// GetType gets the type property value.
func (m *UploadStatusResult) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}

// setType sets the type property value.
func (m *UploadStatusResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}

// GetState gets the state property value.
func (m *UploadStatusResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}

// setState sets the state property value.
func (m *UploadStatusResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}

// GetOutput gets the output property value.
func (m *UploadStatusResult) GetOutput() (*UploadStatusOutput, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *UploadStatusOutput](m.GetBackingStore(), outputKey)
}

// setOutput sets the output property value.
func (m *UploadStatusResult) setOutput(val *UploadStatusOutput) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), outputKey, val)
}

// CreateUploadStatusResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value.
func CreateUploadStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUploadStatusResult(), nil
}

// ExportResult represents an export result.
type ExportResult struct {
	internal.BaseModel
}

func NewExportResult() *ExportResult {
	return &ExportResult{BaseModel: *internal.NewBaseModel()}
}

func (m *ExportResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
	)
}

func (m *ExportResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:   internalSerialization.DeserializeStringFunc()(m.setSysId),
		nameKey:    internalSerialization.DeserializeStringFunc()(m.setName),
		stateKey:   internalSerialization.DeserializeStringFunc()(m.setState),
		statusKey:  internalSerialization.DeserializeStringFunc()(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc()(m.setMessage),
	}
}

func (m *ExportResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *ExportResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *ExportResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *ExportResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}
func (m *ExportResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}
func (m *ExportResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}
func (m *ExportResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), statusKey)
}
func (m *ExportResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), statusKey, val)
}
func (m *ExportResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), messageKey)
}
func (m *ExportResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), messageKey, val)
}

func CreateExportResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExportResult(), nil
}

// ExportStatusResult represents the status of an export.
type ExportStatusResult struct {
	internal.BaseModel
}

func NewExportStatusResult() *ExportStatusResult {
	return &ExportStatusResult{BaseModel: *internal.NewBaseModel()}
}

func (m *ExportStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
		internalSerialization.SerializeStringFunc(progressKey)(m.GetProgress),
	)
}

func (m *ExportStatusResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		stateKey:    internalSerialization.DeserializeStringFunc()(m.setState),
		statusKey:   internalSerialization.DeserializeStringFunc()(m.setStatus),
		messageKey:  internalSerialization.DeserializeStringFunc()(m.setMessage),
		progressKey: internalSerialization.DeserializeStringFunc()(m.setProgress),
	}
}

func (m *ExportStatusResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}
func (m *ExportStatusResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}
func (m *ExportStatusResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), statusKey)
}
func (m *ExportStatusResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), statusKey, val)
}
func (m *ExportStatusResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), messageKey)
}
func (m *ExportStatusResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), messageKey, val)
}
func (m *ExportStatusResult) GetProgress() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), progressKey)
}
func (m *ExportStatusResult) setProgress(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), progressKey, val)
}

func CreateExportStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExportStatusResult(), nil
}

// SharedLibraryComponentApplication represents an application associated with shared libraries.
type SharedLibraryComponentApplication struct {
	internal.BaseModel
}

func NewSharedLibraryComponentApplication() *SharedLibraryComponentApplication {
	return &SharedLibraryComponentApplication{BaseModel: *internal.NewBaseModel()}
}

func (m *SharedLibraryComponentApplication) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(versionKey)(m.GetVersion),
		internalSerialization.SerializeStringFunc(descriptionKey)(m.GetDescription),
		internalSerialization.SerializeStringFunc(appNameKey)(m.GetAppName),
	)
}

func (m *SharedLibraryComponentApplication) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:       internalSerialization.DeserializeStringFunc()(m.setSysId),
		nameKey:        internalSerialization.DeserializeStringFunc()(m.setName),
		versionKey:     internalSerialization.DeserializeStringFunc()(m.setVersion),
		descriptionKey: internalSerialization.DeserializeStringFunc()(m.setDescription),
		appNameKey:     internalSerialization.DeserializeStringFunc()(m.setAppName),
	}
}

func (m *SharedLibraryComponentApplication) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *SharedLibraryComponentApplication) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *SharedLibraryComponentApplication) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *SharedLibraryComponentApplication) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}
func (m *SharedLibraryComponentApplication) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), versionKey)
}
func (m *SharedLibraryComponentApplication) setVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), versionKey, val)
}
func (m *SharedLibraryComponentApplication) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), descriptionKey)
}
func (m *SharedLibraryComponentApplication) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), descriptionKey, val)
}
func (m *SharedLibraryComponentApplication) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appNameKey)
}
func (m *SharedLibraryComponentApplication) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appNameKey, val)
}

func CreateSharedLibraryComponentApplicationFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewSharedLibraryComponentApplication(), nil
}

// ComponentUploadRequest represents the body for uploading components.
type ComponentUploadRequest struct {
	internal.BaseModel
}

func NewComponentUploadRequest() *ComponentUploadRequest {
	return &ComponentUploadRequest{BaseModel: *internal.NewBaseModel()}
}

func (m *ComponentUploadRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey)(m.GetAppName),
		internalSerialization.SerializeStringFunc(componentNameKey)(m.GetComponentName),
		internalSerialization.SerializeStringFunc(dataKey)(m.GetData),
		internalSerialization.SerializeStringFunc(formatKey)(m.GetFormat),
	)
}

func (m *ComponentUploadRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:       internalSerialization.DeserializeStringFunc()(m.setAppName),
		componentNameKey: internalSerialization.DeserializeStringFunc()(m.setComponentName),
		dataKey:          internalSerialization.DeserializeStringFunc()(m.setData),
		formatKey:        internalSerialization.DeserializeStringFunc()(m.setFormat),
	}
}

func (m *ComponentUploadRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appNameKey)
}
func (m *ComponentUploadRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appNameKey, val)
}
func (m *ComponentUploadRequest) GetComponentName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), componentNameKey)
}
func (m *ComponentUploadRequest) setComponentName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), componentNameKey, val)
}
func (m *ComponentUploadRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), dataKey)
}
func (m *ComponentUploadRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), dataKey, val)
}
func (m *ComponentUploadRequest) GetFormat() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), formatKey)
}
func (m *ComponentUploadRequest) setFormat(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), formatKey, val)
}

func CreateComponentUploadRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewComponentUploadRequest(), nil
}

// ComponentVarsUploadRequest represents the body for uploading component variables.
type ComponentVarsUploadRequest struct {
	internal.BaseModel
}

func NewComponentVarsUploadRequest() *ComponentVarsUploadRequest {
	return &ComponentVarsUploadRequest{BaseModel: *internal.NewBaseModel()}
}

func (m *ComponentVarsUploadRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey)(m.GetAppName),
		internalSerialization.SerializeStringFunc(componentNameKey)(m.GetComponentName),
		internalSerialization.SerializeAnyFunc(varsKey)(m.GetVars),
	)
}

func (m *ComponentVarsUploadRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:       internalSerialization.DeserializeStringFunc()(m.setAppName),
		componentNameKey: internalSerialization.DeserializeStringFunc()(m.setComponentName),
		varsKey:          internalSerialization.DeserializeAnyFunc()(m.setVars),
	}
}

func (m *ComponentVarsUploadRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appNameKey)
}
func (m *ComponentVarsUploadRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appNameKey, val)
}
func (m *ComponentVarsUploadRequest) GetComponentName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), componentNameKey)
}
func (m *ComponentVarsUploadRequest) setComponentName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), componentNameKey, val)
}
func (m *ComponentVarsUploadRequest) GetVars() (any, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), varsKey)
}
func (m *ComponentVarsUploadRequest) setVars(val any) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), varsKey, val)
}

func CreateComponentVarsUploadRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewComponentVarsUploadRequest(), nil
}

// CollectionUploadRequest represents the body for uploading collections.
type CollectionUploadRequest struct {
	internal.BaseModel
}

func NewCollectionUploadRequest() *CollectionUploadRequest {
	return &CollectionUploadRequest{BaseModel: *internal.NewBaseModel()}
}

func (m *CollectionUploadRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey)(m.GetAppName),
		internalSerialization.SerializeStringFunc(collectionNameKey)(m.GetCollectionName),
		internalSerialization.SerializeStringFunc(dataKey)(m.GetData),
		internalSerialization.SerializeStringFunc(formatKey)(m.GetFormat),
	)
}

func (m *CollectionUploadRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:        internalSerialization.DeserializeStringFunc()(m.setAppName),
		collectionNameKey: internalSerialization.DeserializeStringFunc()(m.setCollectionName),
		dataKey:           internalSerialization.DeserializeStringFunc()(m.setData),
		formatKey:         internalSerialization.DeserializeStringFunc()(m.setFormat),
	}
}

func (m *CollectionUploadRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appNameKey)
}
func (m *CollectionUploadRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appNameKey, val)
}
func (m *CollectionUploadRequest) GetCollectionName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), collectionNameKey)
}
func (m *CollectionUploadRequest) setCollectionName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), collectionNameKey, val)
}
func (m *CollectionUploadRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), dataKey)
}
func (m *CollectionUploadRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), dataKey, val)
}
func (m *CollectionUploadRequest) GetFormat() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), formatKey)
}
func (m *CollectionUploadRequest) setFormat(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), formatKey, val)
}

func CreateCollectionUploadRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCollectionUploadRequest(), nil
}

// DeployableUpdateRequest represents the body for updating deployables.
type DeployableUpdateRequest struct {
	internal.BaseModel
}

func NewDeployableUpdateRequest() *DeployableUpdateRequest {
	return &DeployableUpdateRequest{BaseModel: *internal.NewBaseModel()}
}

func (m *DeployableUpdateRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey)(m.GetAppName),
		internalSerialization.SerializeStringFunc(deployableNameKey)(m.GetDeployableName),
		internalSerialization.SerializeStringFunc(dataKey)(m.GetData),
	)
}

func (m *DeployableUpdateRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:        internalSerialization.DeserializeStringFunc()(m.setAppName),
		deployableNameKey: internalSerialization.DeserializeStringFunc()(m.setDeployableName),
		dataKey:           internalSerialization.DeserializeStringFunc()(m.setData),
	}
}

func (m *DeployableUpdateRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appNameKey)
}
func (m *DeployableUpdateRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appNameKey, val)
}
func (m *DeployableUpdateRequest) GetDeployableName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), deployableNameKey)
}
func (m *DeployableUpdateRequest) setDeployableName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), deployableNameKey, val)
}
func (m *DeployableUpdateRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), dataKey)
}
func (m *DeployableUpdateRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), dataKey, val)
}

func CreateDeployableUpdateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewDeployableUpdateRequest(), nil
}

// SharedComponentUpdateRequest represents the body for updating shared components.
type SharedComponentUpdateRequest struct {
	internal.BaseModel
}

func NewSharedComponentUpdateRequest() *SharedComponentUpdateRequest {
	return &SharedComponentUpdateRequest{BaseModel: *internal.NewBaseModel()}
}

func (m *SharedComponentUpdateRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey)(m.GetAppName),
		internalSerialization.SerializeStringFunc(sharedComponentNameKey)(m.GetSharedComponentName),
		internalSerialization.SerializeStringFunc(dataKey)(m.GetData),
	)
}

func (m *SharedComponentUpdateRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:             internalSerialization.DeserializeStringFunc()(m.setAppName),
		sharedComponentNameKey: internalSerialization.DeserializeStringFunc()(m.setSharedComponentName),
		dataKey:                internalSerialization.DeserializeStringFunc()(m.setData),
	}
}

func (m *SharedComponentUpdateRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), appNameKey)
}
func (m *SharedComponentUpdateRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), appNameKey, val)
}
func (m *SharedComponentUpdateRequest) GetSharedComponentName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sharedComponentNameKey)
}
func (m *SharedComponentUpdateRequest) setSharedComponentName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sharedComponentNameKey, val)
}
func (m *SharedComponentUpdateRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), dataKey)
}
func (m *SharedComponentUpdateRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), dataKey, val)
}

func CreateSharedComponentUpdateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewSharedComponentUpdateRequest(), nil
}

// Media represents an HTTP media file for upload.
type Media struct {
	contentType string
	data        []byte
}

func NewMedia(contentType string, data []byte) *Media {
	return &Media{
		contentType: contentType,
		data:        data,
	}
}

func (p *Media) GetContentType() string {
	return p.contentType
}

func (p *Media) GetData() []byte {
	return p.data
}

func (p *Media) Serialize(writer serialization.SerializationWriter) error {
	return writer.WriteByteArrayValue("", p.data)
}

func (p *Media) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}
