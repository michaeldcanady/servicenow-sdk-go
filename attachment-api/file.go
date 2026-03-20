package attachmentapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	model "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	createdByNameKey = "created_by_name"
	updatedByNameKey = "update_by_name"
)

// File defines a serializable model object.
type File interface {
	GetAverageImageColor() (*string, error)
	SetAverageImageColor(*string) error
	GetCompressed() (*bool, error)
	SetCompressed(*bool) error
	GetContentType() (*string, error)
	SetContentType(*string) error
	GetCreatedByName() (*string, error)
	SetCreatedByName(*string) error
	GetCreatedBy() (*string, error)
	SetCreatedBy(*string) error
	GetDownloadLink() (*string, error)
	SetDownloadLink(*string) error
	GetFileName() (*string, error)
	SetFileName(*string) error
	GetImageHeight() (*float64, error)
	SetImageHeight(*float64) error
	GetImageWidth() (*float64, error)
	SetImageWidth(*float64) error
	GetSizeBytes() (*int64, error)
	SetSizeBytes(*int64) error
	GetSizeCompressed() (*int64, error)
	SetSizeCompressed(*int64) error
	GetSysCreatedBy() (*string, error)
	SetSysCreatedBy(*string) error
	GetSysCreatedOn() (*time.Time, error)
	SetSysCreatedOn(*time.Time) error
	GetSysID() (*string, error)
	SetSysID(*string) error
	GetSysModCount() (*int64, error)
	SetSysModCount(*int64) error
	GetSysTags() ([]string, error)
	SetSysTags([]string) error
	GetSysUpdatedBy() (*string, error)
	SetSysUpdatedBy(*string) error
	GetSysUpdatedOn() (*time.Time, error)
	SetSysUpdatedOn(*time.Time) error
	GetTableName() (*string, error)
	SetTableName(*string) error
	GetTableSysID() (*string, error)
	SetTableSysID(*string) error
	GetUpdatedByName() (*string, error)
	SetUpdatedByName(*string) error
	serialization.Parsable
	kiotaStore.BackedModel
}

type FileModel struct {
	model.Model
}

// NewFile creates a new instance of FileModel
func NewFile() *FileModel {
	return newFile(model.NewBaseModel())
}

// newFile creates a new instance of FileModel with the provided model underlying it
func newFile(model model.Model) *FileModel {
	return &FileModel{
		model,
	}
}

// CreateFileFromDiscriminatorValue is a parsable factory for creating a Fileable
func CreateFileFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFile(), nil
}

// Serialize writes the objects properties to the current writer.
func (f *FileModel) Serialize(writer serialization.SerializationWriter) error { //nolint:gocognit
	if utils.IsNil(f) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(averageImageColorKey)(f.GetAverageImageColor),
		internalSerialization.SerializeStringToBoolFunc(compressedKey)(f.GetCompressed),
		internalSerialization.SerializeStringFunc(contentTypeKey)(f.GetContentType),
		internalSerialization.SerializeStringFunc(createdByNameKey)(f.GetCreatedByName),
		internalSerialization.SerializeStringFunc(downloadLinkKey)(f.GetDownloadLink),
		internalSerialization.SerializeStringFunc(fileNameKey)(f.GetFileName),
		internalSerialization.SerializeStringToFloat64Func(imageHeightKey)(f.GetImageHeight),
		internalSerialization.SerializeStringToFloat64Func(imageWidthKey)(f.GetImageWidth),
		internalSerialization.SerializeStringToInt64Func(sizeBytesKey)(f.GetSizeBytes),
		internalSerialization.SerializeStringToInt64Func(sizeCompressedKey)(f.GetSizeCompressed),
		internalSerialization.SerializeStringFunc(sysCreatedByKey)(f.GetSysCreatedBy),
		internalSerialization.SerializeStringToTimeFunc(sysCreatedOnKey, time.RFC3339)(f.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysIDKey)(f.GetSysID),
		internalSerialization.SerializeStringToInt64Func(sysModCountKey)(f.GetSysModCount),
		internalSerialization.SerializeStringToSliceFunc(sysTagsKey, " ")(f.GetSysTags),
		internalSerialization.SerializeStringFunc(sysUpdatedByKey)(f.GetSysUpdatedBy),
		internalSerialization.SerializeStringToTimeFunc(sysUpdatedOnKey, time.RFC3339)(f.GetSysUpdatedOn),
		internalSerialization.SerializeStringFunc(tableNameKey)(f.GetTableName),
		internalSerialization.SerializeStringFunc(tableSysIDKey)(f.GetTableSysID),
		internalSerialization.SerializeStringFunc(updatedByNameKey)(f.GetUpdatedByName),
	)
}

// GetAverageImageColor returns, If the attachment is an image, the sum of all colors.
func (f *FileModel) GetAverageImageColor() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), averageImageColorKey)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (f *FileModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { //nolint:gocognit
	return map[string]func(serialization.ParseNode) error{
		averageImageColorKey: kiota.DeserializeStringFunc(f.SetAverageImageColor),
		compressedKey:        kiota.DeserializeMutatedStringFunc(utils.StringPtrToBoolPtr)(f.SetCompressed),
		contentTypeKey:       kiota.DeserializeStringFunc(f.SetContentType),
		createdByNameKey:     kiota.DeserializeStringFunc(f.SetCreatedByName),
		downloadLinkKey:      kiota.DeserializeStringFunc(f.SetDownloadLink),
		fileNameKey:          kiota.DeserializeStringFunc(f.SetFileName),
		imageHeightKey:       kiota.DeserializeMutatedStringFunc(utils.StringPtrToFloat64Ptr)(f.SetImageHeight),
		imageWidthKey:        kiota.DeserializeMutatedStringFunc(utils.StringPtrToFloat64Ptr)(f.SetImageWidth),
		sizeBytesKey:         kiota.DeserializeMutatedStringFunc(utils.StringPtrToInt64Ptr)(f.SetSizeBytes),
		sizeCompressedKey:    kiota.DeserializeMutatedStringFunc(utils.StringPtrToInt64Ptr)(f.SetSizeCompressed),
		sysCreatedByKey:      kiota.DeserializeStringFunc(f.SetSysCreatedBy),
		sysCreatedOnKey:      kiota.DeserializeMutatedStringFunc(utils.StringPtrToTimePtr("2006-01-02 15:04:05"))(f.SetSysCreatedOn),
		sysIDKey:             kiota.DeserializeStringFunc(f.SetSysID),
		sysModCountKey:       kiota.DeserializeMutatedStringFunc(utils.StringPtrToInt64Ptr)(f.SetSysModCount),
		sysTagsKey:           kiota.DeserializeMutatedStringFunc(utils.StringPtrToPrimitiveSlice(" ", func(s string) (string, error) { return s, nil }))(f.SetSysTags),
		sysUpdatedByKey:      kiota.DeserializeStringFunc(f.SetSysUpdatedBy),
		sysUpdatedOnKey:      kiota.DeserializeMutatedStringFunc(utils.StringPtrToTimePtr("2006-01-02 15:04:05"))(f.SetSysUpdatedOn),
		tableNameKey:         kiota.DeserializeStringFunc(f.SetTableName),
		tableSysIDKey:        kiota.DeserializeStringFunc(f.SetTableSysID),
		updatedByNameKey:     kiota.DeserializeStringFunc(f.SetUpdatedByName),
	}
}

// SetAverageImageColor Sets the sum of all colors.
func (f *FileModel) SetAverageImageColor(averageImageColor *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), averageImageColorKey, averageImageColor)
}

// GetCompressed return flag that indicates whether the attachment file has been compressed.
func (f *FileModel) GetCompressed() (*bool, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](f.GetBackingStore(), compressedKey)
}

// SetCompressed Sets flag that indicates whether the attachment file has been compressed.
func (f *FileModel) SetCompressed(compressed *bool) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), compressedKey, compressed)
}

// GetContentType returns content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *FileModel) GetContentType() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), contentTypeKey)
}

// SetContentType Sets content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *FileModel) SetContentType(contentType *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), contentTypeKey, contentType)
}

// GetCreatedByName returns full name of entity that originally created the attachment file.
func (f *FileModel) GetCreatedByName() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), createdByNameKey)
}

// SetCreatedByName Sets full name of entity that originally created the attachment file.
func (f *FileModel) SetCreatedByName(createdByName *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), createdByNameKey, createdByName)
}

// GetCreatedBy returns the entity that originally created the attachment file.
func (f *FileModel) GetCreatedBy() (*string, error) {
	return f.GetSysCreatedBy()
}

// SetCreatedBy Sets the entity that originally created the attachment file.
func (f *FileModel) SetCreatedBy(createdBy *string) error {
	return f.SetSysCreatedBy(createdBy)
}

// GetDownloadLink returns download URL of the attachment on the ServiceNow instance.
func (f *FileModel) GetDownloadLink() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), downloadLinkKey)
}

// SetDownloadLink Sets download URL of the attachment on the ServiceNow instance.
func (f *FileModel) SetDownloadLink(downloadLink *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), downloadLinkKey, downloadLink)
}

// GetFileName returns the file name of the attachment.
func (f *FileModel) GetFileName() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), fileNameKey)
}

// SetFileName Sets the file name of the attachment.
func (f *FileModel) SetFileName(fileName *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), fileNameKey, fileName)
}

// GetImageHeight returns if an image file, the height of the image.
func (f *FileModel) GetImageHeight() (*float64, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](f.GetBackingStore(), imageHeightKey)
}

// SetImageHeight Sets if an image file, the height of the image.
func (f *FileModel) SetImageHeight(imageHeight *float64) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), imageHeightKey, imageHeight)
}

// GetImageWidth returns if an image file, the width of the image.
func (f *FileModel) GetImageWidth() (*float64, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](f.GetBackingStore(), imageWidthKey)
}

// SetImageWidth Sets if an image file, the width of the image.
func (f *FileModel) SetImageWidth(imageWidth *float64) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), imageWidthKey, imageWidth)
}

// GetSizeBytes returns size of the attachment.
func (f *FileModel) GetSizeBytes() (*int64, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](f.GetBackingStore(), sizeBytesKey)
}

// SetSizeBytes Sets size of the attachment.
func (f *FileModel) SetSizeBytes(sizeBytes *int64) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sizeBytesKey, sizeBytes)
}

// GetSizeCompressed returns size of the compressed attachment file. If the file is not compressed, empty.
func (f *FileModel) GetSizeCompressed() (*int64, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](f.GetBackingStore(), sizeCompressedKey)
}

// SetSizeCompressed Sets size of the compressed attachment file.
func (f *FileModel) SetSizeCompressed(sizeCompressed *int64) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sizeCompressedKey, sizeCompressed)
}

// GetSysCreatedBy returns the entity that originally created the attachment file.
func (f *FileModel) GetSysCreatedBy() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), sysCreatedByKey)
}

// SetSysCreatedBy Sets the entity that originally created the attachment file.
func (f *FileModel) SetSysCreatedBy(sysCreatedBy *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sysCreatedByKey, sysCreatedBy)
}

// GetSysCreatedOn returns the date and time that the attachment file was initially saved to the instance.
func (f *FileModel) GetSysCreatedOn() (*time.Time, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](f.GetBackingStore(), sysCreatedOnKey)
}

// SetSysCreatedOn Sets the date and time that the attachment file was initially saved to the instance.
func (f *FileModel) SetSysCreatedOn(sysCreatedOn *time.Time) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sysCreatedOnKey, sysCreatedOn)
}

// GetSysID returns the sys_id of the attachment file. Read-Only.
func (f *FileModel) GetSysID() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), sysIDKey)
}

// SetSysID Sets the sys_id of the attachment file.
func (f *FileModel) SetSysID(sysID *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sysIDKey, sysID)
}

// GetSysModCount returns the number of times the attachment file has been modified (uploaded to the instance).
func (f *FileModel) GetSysModCount() (*int64, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](f.GetBackingStore(), sysModCountKey)
}

// SetSysModCount Sets the number of times the attachment file has been modified (uploaded to the instance).
func (f *FileModel) SetSysModCount(sysModCount *int64) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sysModCountKey, sysModCount)
}

// GetSysTags returns any system tags associated with the attachment file.
func (f *FileModel) GetSysTags() ([]string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []string](f.GetBackingStore(), sysTagsKey)
}

// SetSysTags Sets any system tags associated with the attachment file.
func (f *FileModel) SetSysTags(sysTags []string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sysTagsKey, sysTags)
}

// GetSysUpdatedBy returns the entity that last updated the attachment file.
func (f *FileModel) GetSysUpdatedBy() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), sysUpdatedByKey)
}

// SetSysUpdatedBy Sets the entity that last updated the attachment file.
func (f *FileModel) SetSysUpdatedBy(sysUpdatedBy *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sysUpdatedByKey, sysUpdatedBy)
}

// GetSysUpdatedOn returns the date and time that the attachment file was last updated.
func (f *FileModel) GetSysUpdatedOn() (*time.Time, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](f.GetBackingStore(), sysUpdatedOnKey)
}

// SetSysUpdatedOn Sets the date and time that the attachment file was last updated.
func (f *FileModel) SetSysUpdatedOn(sysUpdatedOn *time.Time) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), sysUpdatedOnKey, sysUpdatedOn)
}

// GetTableName returns the name of the table to which the attachment is associated.
func (f *FileModel) GetTableName() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), tableNameKey)
}

// SetTableName Sets the name of the table to which the attachment is associated.
func (f *FileModel) SetTableName(tableName *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), tableNameKey, tableName)
}

// GetTableSysID returns the sys_id of the table associated with the attachment.
func (f *FileModel) GetTableSysID() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), tableSysIDKey)
}

// SetTableSysID Sets the sys_id of the table associated with the attachment.
func (f *FileModel) SetTableSysID(tableSysID *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), tableSysIDKey, tableSysID)
}

// GetUpdatedByName returns the full name of entity that last updated the attachment file.
func (f *FileModel) GetUpdatedByName() (*string, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](f.GetBackingStore(), updatedByNameKey)
}

// SetUpdatedByName Sets the full name of entity that last updated the attachment file.
func (f *FileModel) SetUpdatedByName(updatedByName *string) error {
	if utils.IsNil(f) {
		return nil
	}

	return kiota.DefaultBackedModelMutatorFunc(f.GetBackingStore(), updatedByNameKey, updatedByName)
}
