package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

var _ ActivitiesResultEntryAttachment = (*ActivitiesResultEntryAttachmentModel)(nil)

const (
	averageImageColorKey = "average_image_color"
	contentTypeKey       = "content_type"
	fileNameKey          = "file_name"
	imageHeightKey       = "image_height"
	imagePathKey         = "image_path"
	imageWidthKey        = "image_width"
	pathKey              = "path"
	sizeKey              = "size"
	sizeBytesKey         = "size_bytes"
)

type ActivitiesResultEntryAttachment interface {
	serialization.Parsable
	kiotaStore.BackedModel
	// GetAverageImageColor returns the hexadecimal representation of average color for image attachment.
	GetAverageImageColor() (*string, error)
	// SetAverageImageColor sets the hexadecimal representation of average color for image attachment.
	SetAverageImageColor(*string) error
	// GetContentType returns the MIME type for the attachment.
	GetContentType() (*string, error)
	// SetContentType sets the MIME type for the attachment.
	SetContentType(*string) error
	// GetFileName returns the file name for the attachment.
	GetFileName() (*string, error)
	// SetFileName sets the file name for the attachment.
	SetFileName(*string) error
	// GetImageHeight returns the height of image attachment in pixels.
	GetImageHeight() (*int64, error)
	// SetImageHeight sets the height of image attachment in pixels.
	SetImageHeight(*int64) error
	// TODO: is a url http.url
	// GetImagePath returns the direct download link for image attachment.
	GetImagePath() (*string, error)
	// SetImagePath sets the direct download link for image attachment.
	SetImagePath(*string) error
	// GetImageWidth returns the width of image attachment in pixels.
	GetImageWidth() (*int64, error)
	// SetImageWidth sets the width of image attachment in pixels.
	SetImageWidth(*int64) error
	// TODO: is a url http.url
	// GetPath returns the direct link to the attachment.
	GetPath() (*string, error)
	// SetPath sets the direct link to the attachment.
	SetPath(*string) error
	// TODO: how should formatted sizes be represented.
	// GetSize returns the size of the attachment.
	GetSize() (*string, error)
	// SetSize sets the size of the attachment.
	SetSize(*string) error
	// GetSizeBytes returns the size of the attachment in bytes.
	GetSizeBytes() (*int64, error)
	// SetSizeBytes sets the size of the attachment in bytes.
	SetSizeBytes(*int64) error
	// GetState returns the state of the attachment record.
	GetState() (*State, error)
	// SetState sets the state of the attachment record.
	SetState(*State) error
	// GetSysID returns the sys_id for the attachment.
	GetSysID() (*string, error)
	// SetSysID sets the sys_id for the attachment.
	SetSysID(*string) error

	// TODO: is url http.url
	// GetThumbnailPath returns the direct download link for thumbnail of image attachment.
	GetThumbnailPath() (*string, error)
	// SetThumbnailPath returns the direct download link for thumbnail of image attachment.
	SetThumbnailPath(*string) error
}

type ActivitiesResultEntryAttachmentModel struct {
	core.BaseModel
}

// NewActivitiesResultEntryAttachment creates a new instance of [ActivitiesResultEntryAttachmentModel].
func NewActivitiesResultEntryAttachment() *ActivitiesResultEntryAttachmentModel {
	return &ActivitiesResultEntryAttachmentModel{BaseModel: *core.NewBaseModel()}
}

// CreateActivitiesResultEntryAttachmentFromDiscriminatorValue creates a new instance of [ActivitiesResultEntryAttachment].
func CreateActivitiesResultEntryAttachmentFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResultEntryAttachment(), nil
}

// Serialize implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(a) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(averageImageColorKey, a.GetAverageImageColor),
		internalSerialization.SerializeStringFunc(contentTypeKey, a.GetContentType),
		internalSerialization.SerializeStringFunc(fileNameKey, a.GetFileName),
		internalSerialization.SerializeInt64Func(imageHeightKey, a.GetImageHeight),
		internalSerialization.SerializeStringFunc(imagePathKey, a.GetImagePath),
		internalSerialization.SerializeInt64Func(imageWidthKey, a.GetImageWidth),
		internalSerialization.SerializeStringFunc(pathKey, a.GetPath),
		internalSerialization.SerializeStringFunc(sizeKey, a.GetSize),
		internalSerialization.SerializeInt64Func(sizeBytesKey, a.GetSizeBytes),
		internalSerialization.SerializeEnumFunc(stateKey, a.GetState),
		internalSerialization.SerializeStringFunc(sysIDKey, a.GetSysID),
		internalSerialization.SerializeStringFunc(thumbnailPathKey, a.GetThumbnailPath),
	)
}

// GetFieldDeserializers implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		averageImageColorKey: internalSerialization.DeserializeStringFunc(a.SetAverageImageColor),
		contentTypeKey:       internalSerialization.DeserializeStringFunc(a.SetContentType),
		fileNameKey:          internalSerialization.DeserializeStringFunc(a.SetFileName),
		imageHeightKey:       internalSerialization.DeserializeInt64Func(a.SetImageHeight),
		imagePathKey:         internalSerialization.DeserializeStringFunc(a.SetImagePath),
		imageWidthKey:        internalSerialization.DeserializeInt64Func(a.SetImageWidth),
		pathKey:              internalSerialization.DeserializeStringFunc(a.SetPath),
		sizeKey:              internalSerialization.DeserializeStringFunc(a.SetSize),
		sizeBytesKey:         internalSerialization.DeserializeInt64Func(a.SetSizeBytes),
		stateKey:             internalSerialization.DeserializeEnumFunc(ParseState, a.SetState),
		sysIDKey:             internalSerialization.DeserializeStringFunc(a.SetSysID),
		thumbnailPathKey:     internalSerialization.DeserializeStringFunc(a.SetThumbnailPath),
	}
}

// GetAverageImageColor implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetAverageImageColor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *string](a, averageImageColorKey)
}

// GetContentType implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetContentType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *string](a, contentTypeKey)
}

// GetFileName implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetFileName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *string](a, fileNameKey)
}

// GetImageHeight implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetImageHeight() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *int64](a, imageHeightKey)
}

// GetImagePath implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetImagePath() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *string](a, imagePathKey)
}

// GetImageWidth implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetImageWidth() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *int64](a, imageWidthKey)
}

// GetPath implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetPath() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *string](a, pathKey)
}

// GetSize implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetSize() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *string](a, sizeKey)
}

// GetSizeBytes implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetSizeBytes() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *int64](a, sizeBytesKey)
}

// GetState implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetState() (*State, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *State](a, stateKey)
}

// GetSysID implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *string](a, sysIDKey)
}

// GetThumbnailPath implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) GetThumbnailPath() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultEntryAttachmentModel, *string](a, thumbnailPathKey)
}

// SetAverageImageColor implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetAverageImageColor(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, averageImageColorKey, val)
}

// SetContentType implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetContentType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, contentTypeKey, val)
}

// SetFileName implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetFileName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, fileNameKey, val)
}

// SetImageHeight implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetImageHeight(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(a, imageHeightKey, val)
}

// SetImagePath implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetImagePath(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, imagePathKey, val)
}

// SetImageWidth implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetImageWidth(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(a, imageWidthKey, val)
}

// SetPath implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetPath(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, pathKey, val)
}

// SetSize implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetSize(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, sizeKey, val)
}

// SetSizeBytes implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetSizeBytes(val *int64) error {
	return store.DefaultBackedModelMutatorFunc(a, sizeBytesKey, val)
}

// SetState implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetState(val *State) error {
	return store.DefaultBackedModelMutatorFunc(a, stateKey, val)
}

// SetSysID implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, sysIDKey, val)
}

// SetThumbnailPath implements [ActivitiesResultEntryAttachment].
func (a *ActivitiesResultEntryAttachmentModel) SetThumbnailPath(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, thumbnailPathKey, val)
}
