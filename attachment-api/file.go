package attachmentapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
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
	newInternal.Model
}

// NewFile creates a new instance of FileModel
func NewFile() *FileModel {
	return newFile(newInternal.NewBaseModel())
}

// newFile creates a new instance of FileModel with the provided model underlying it
func newFile(model newInternal.Model) *FileModel {
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
	if internal.IsNil(f) {
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
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, averageImageColorKey)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (f *FileModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { //nolint:gocognit
	return map[string]func(serialization.ParseNode) error{
		averageImageColorKey: internalSerialization.DeserializeStringFunc()(f.SetAverageImageColor),
		compressedKey:        internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(f.SetCompressed),
		contentTypeKey:       internalSerialization.DeserializeStringFunc()(f.SetContentType),
		createdByNameKey:     internalSerialization.DeserializeStringFunc()(f.SetCreatedByName),
		downloadLinkKey:      internalSerialization.DeserializeStringFunc()(f.SetDownloadLink),
		fileNameKey:          internalSerialization.DeserializeStringFunc()(f.SetFileName),
		imageHeightKey:       internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(f.SetImageHeight),
		imageWidthKey:        internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(f.SetImageWidth),
		sizeBytesKey:         internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(f.SetSizeBytes),
		sizeCompressedKey:    internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(f.SetSizeCompressed),
		sysCreatedByKey:      internalSerialization.DeserializeStringFunc()(f.SetSysCreatedBy),
		sysCreatedOnKey:      internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr("2006-01-02 15:04:05"))(f.SetSysCreatedOn),
		sysIDKey:             internalSerialization.DeserializeStringFunc()(f.SetSysID),
		sysModCountKey:       internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(f.SetSysModCount),
		sysTagsKey:           internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToPrimitiveSlice(" ", func(s string) (string, error) { return s, nil }))(f.SetSysTags),
		sysUpdatedByKey:      internalSerialization.DeserializeStringFunc()(f.SetSysUpdatedBy),
		sysUpdatedOnKey:      internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr("2006-01-02 15:04:05"))(f.SetSysUpdatedOn),
		tableNameKey:         internalSerialization.DeserializeStringFunc()(f.SetTableName),
		tableSysIDKey:        internalSerialization.DeserializeStringFunc()(f.SetTableSysID),
		updatedByNameKey:     internalSerialization.DeserializeStringFunc()(f.SetUpdatedByName),
	}
}

// SetAverageImageColor Sets the sum of all colors.
func (f *FileModel) SetAverageImageColor(averageImageColor *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, averageImageColorKey, averageImageColor)
}

// GetCompressed return flag that indicates whether the attachment file has been compressed.
func (f *FileModel) GetCompressed() (*bool, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, compressedKey)
}

// SetCompressed Sets flag that indicates whether the attachment file has been compressed.
func (f *FileModel) SetCompressed(compressed *bool) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, compressedKey, compressed)
}

// GetContentType returns content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *FileModel) GetContentType() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, contentTypeKey)
}

// SetContentType Sets content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *FileModel) SetContentType(contentType *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, contentTypeKey, contentType)
}

// GetCreatedByName returns full name of entity that originally created the attachment file.
func (f *FileModel) GetCreatedByName() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, createdByNameKey)
}

// SetCreatedByName Sets full name of entity that originally created the attachment file.
func (f *FileModel) SetCreatedByName(createdByName *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, createdByNameKey, createdByName)
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
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, downloadLinkKey)
}

// SetDownloadLink Sets download URL of the attachment on the ServiceNow instance.
func (f *FileModel) SetDownloadLink(downloadLink *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, downloadLinkKey, downloadLink)
}

// GetFileName returns the file name of the attachment.
func (f *FileModel) GetFileName() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, fileNameKey)
}

// SetFileName Sets the file name of the attachment.
func (f *FileModel) SetFileName(fileName *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, fileNameKey, fileName)
}

// GetImageHeight returns if an image file, the height of the image.
func (f *FileModel) GetImageHeight() (*float64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, imageHeightKey)
}

// SetImageHeight Sets if an image file, the height of the image.
func (f *FileModel) SetImageHeight(imageHeight *float64) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, imageHeightKey, imageHeight)
}

// GetImageWidth returns if an image file, the width of the image.
func (f *FileModel) GetImageWidth() (*float64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, imageWidthKey)
}

// SetImageWidth Sets if an image file, the width of the image.
func (f *FileModel) SetImageWidth(imageWidth *float64) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, imageWidthKey, imageWidth)
}

// GetSizeBytes returns size of the attachment.
func (f *FileModel) GetSizeBytes() (*int64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sizeBytesKey)
}

// SetSizeBytes Sets size of the attachment.
func (f *FileModel) SetSizeBytes(sizeBytes *int64) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sizeBytesKey, sizeBytes)
}

// GetSizeCompressed returns size of the compressed attachment file. If the file is not compressed, empty.
func (f *FileModel) GetSizeCompressed() (*int64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sizeCompressedKey)
}

// SetSizeCompressed Sets size of the compressed attachment file.
func (f *FileModel) SetSizeCompressed(sizeCompressed *int64) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sizeCompressedKey, sizeCompressed)
}

// GetSysCreatedBy returns the entity that originally created the attachment file.
func (f *FileModel) GetSysCreatedBy() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysCreatedByKey)
}

// SetSysCreatedBy Sets the entity that originally created the attachment file.
func (f *FileModel) SetSysCreatedBy(sysCreatedBy *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysCreatedByKey, sysCreatedBy)
}

// GetSysCreatedOn returns the date and time that the attachment file was initially saved to the instance.
func (f *FileModel) GetSysCreatedOn() (*time.Time, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sysCreatedOnKey)
}

// SetSysCreatedOn Sets the date and time that the attachment file was initially saved to the instance.
func (f *FileModel) SetSysCreatedOn(sysCreatedOn *time.Time) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysCreatedOnKey, sysCreatedOn)
}

// GetSysID returns the sys_id of the attachment file. Read-Only.
func (f *FileModel) GetSysID() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysIDKey)
}

// SetSysID Sets the sys_id of the attachment file.
func (f *FileModel) SetSysID(sysID *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysIDKey, sysID)
}

// GetSysModCount returns the number of times the attachment file has been modified (uploaded to the instance).
func (f *FileModel) GetSysModCount() (*int64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sysModCountKey)
}

// SetSysModCount Sets the number of times the attachment file has been modified (uploaded to the instance).
func (f *FileModel) SetSysModCount(sysModCount *int64) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysModCountKey, sysModCount)
}

// GetSysTags returns any system tags associated with the attachment file.
func (f *FileModel) GetSysTags() ([]string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []string](backingStore, sysTagsKey)
}

// SetSysTags Sets any system tags associated with the attachment file.
func (f *FileModel) SetSysTags(sysTags []string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysTagsKey, sysTags)
}

// GetSysUpdatedBy returns the entity that last updated the attachment file.
func (f *FileModel) GetSysUpdatedBy() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysUpdatedByKey)
}

// SetSysUpdatedBy Sets the entity that last updated the attachment file.
func (f *FileModel) SetSysUpdatedBy(sysUpdatedBy *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysUpdatedByKey, sysUpdatedBy)
}

// GetSysUpdatedOn returns the date and time that the attachment file was last updated.
func (f *FileModel) GetSysUpdatedOn() (*time.Time, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sysUpdatedOnKey)
}

// SetSysUpdatedOn Sets the date and time that the attachment file was last updated.
func (f *FileModel) SetSysUpdatedOn(sysUpdatedOn *time.Time) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysUpdatedOnKey, sysUpdatedOn)
}

// GetTableName returns the name of the table to which the attachment is associated.
func (f *FileModel) GetTableName() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, tableNameKey)
}

// SetTableName Sets the name of the table to which the attachment is associated.
func (f *FileModel) SetTableName(tableName *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, tableNameKey, tableName)
}

// GetTableSysID returns the sys_id of the table associated with the attachment.
func (f *FileModel) GetTableSysID() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, tableSysIDKey)
}

// SetTableSysID Sets the sys_id of the table associated with the attachment.
func (f *FileModel) SetTableSysID(tableSysID *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, tableSysIDKey, tableSysID)
}

// GetUpdatedByName returns the full name of entity that last updated the attachment file.
func (f *FileModel) GetUpdatedByName() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, updatedByNameKey)
}

// SetUpdatedByName Sets the full name of entity that last updated the attachment file.
func (f *FileModel) SetUpdatedByName(updatedByName *string) error {
	if internal.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, updatedByNameKey, updatedByName)
}
