package cdmapplicationsapi // nolint:dupl // shares field-count shape with ComponentUploadRequest/ExportStatusResult by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CollectionUploadRequest represents the body for uploading collections.
type CollectionUploadRequest struct {
	core.BaseModel
}

// NewCollectionUploadRequest instantiates a new CollectionUploadRequest.
func NewCollectionUploadRequest() *CollectionUploadRequest {
	return &CollectionUploadRequest{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *CollectionUploadRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey, m.GetAppName),
		internalSerialization.SerializeStringFunc(collectionNameKey, m.GetCollectionName),
		internalSerialization.SerializeStringFunc(dataKey, m.GetData),
		internalSerialization.SerializeStringFunc(formatKey, m.GetFormat),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *CollectionUploadRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:        internalSerialization.DeserializeStringFunc(m.setAppName),
		collectionNameKey: internalSerialization.DeserializeStringFunc(m.setCollectionName),
		dataKey:           internalSerialization.DeserializeStringFunc(m.setData),
		formatKey:         internalSerialization.DeserializeStringFunc(m.setFormat),
	}
}

// GetAppName returns the app name.
func (m *CollectionUploadRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CollectionUploadRequest, *string](m, appNameKey)
}
func (m *CollectionUploadRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appNameKey, val)
}

// GetCollectionName returns the collection name.
func (m *CollectionUploadRequest) GetCollectionName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CollectionUploadRequest, *string](m, collectionNameKey)
}
func (m *CollectionUploadRequest) setCollectionName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, collectionNameKey, val)
}

// GetData returns the data.
func (m *CollectionUploadRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CollectionUploadRequest, *string](m, dataKey)
}
func (m *CollectionUploadRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dataKey, val)
}

// GetFormat returns the format.
func (m *CollectionUploadRequest) GetFormat() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CollectionUploadRequest, *string](m, formatKey)
}
func (m *CollectionUploadRequest) setFormat(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, formatKey, val)
}

// CreateCollectionUploadRequestFromDiscriminatorValue creates a new CollectionUploadRequest from a ParseNode.
func CreateCollectionUploadRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCollectionUploadRequest(), nil
}
