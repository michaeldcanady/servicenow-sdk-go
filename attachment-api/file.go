package attachmentapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	createdByNameKey = "created_by_name"
	updatedByNameKey = "update_by_name"
)

// file defines a serializable model object.
type file interface {
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

type File struct {
	internal.BackedModel
}

// NewFile creates a new instance of FileModel
func NewFile() *File {
	return newFile(internal.NewBaseModel())
}

// newFile creates a new instance of FileModel with the provided model underlying it
func newFile(model internal.BackedModel) *File {
	return &File{
		model,
	}
}

// CreateFileFromDiscriminatorValue is a parsable factory for creating a Fileable
func CreateFileFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFile(), nil
}

// Serialize writes the objects properties to the current writer.
func (f *File) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(f) {
		return nil
	}

	return internalSerialization.Serialize(writer, f.getSerializationFields()...)
}

func (f *File) getSerializationFields() []internalSerialization.WriterFunc {
	return []internalSerialization.WriterFunc{
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
	}
}

// GetAverageImageColor returns, If the attachment is an image, the sum of all colors.
func (f *File) GetAverageImageColor() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, averageImageColorKey)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (f *File) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	deserializers := map[string]func(serialization.ParseNode) error{
		averageImageColorKey: internalSerialization.DeserializeStringFunc()(f.SetAverageImageColor),
		compressedKey:        internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(f.SetCompressed),
		contentTypeKey:       internalSerialization.DeserializeStringFunc()(f.SetContentType),
		createdByNameKey:     internalSerialization.DeserializeStringFunc()(f.SetCreatedByName),
		downloadLinkKey:      internalSerialization.DeserializeStringFunc()(f.SetDownloadLink),
	}

	for k, v := range f.getAdditionalFieldDeserializers() {
		deserializers[k] = v
	}
	return deserializers
}

func (f *File) getAdditionalFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		fileNameKey:        internalSerialization.DeserializeStringFunc()(f.SetFileName),
		imageHeightKey:     internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(f.SetImageHeight),
		imageWidthKey:      internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(f.SetImageWidth),
		sizeBytesKey:       internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(f.SetSizeBytes),
		sizeCompressedKey:  internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(f.SetSizeCompressed),
		sysCreatedByKey:    internalSerialization.DeserializeStringFunc()(f.SetSysCreatedBy),
		sysCreatedOnKey:    internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr("2006-01-02 15:04:05"))(f.SetSysCreatedOn),
		sysIDKey:           internalSerialization.DeserializeStringFunc()(f.SetSysID),
		sysModCountKey:     internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(f.SetSysModCount),
		sysTagsKey:         internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToPrimitiveSlice(" ", func(s string) (string, error) { return s, nil }))(f.SetSysTags),
		sysUpdatedByKey:    internalSerialization.DeserializeStringFunc()(f.SetSysUpdatedBy),
		sysUpdatedOnKey:    internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr("2006-01-02 15:04:05"))(f.SetSysUpdatedOn),
		tableNameKey:       internalSerialization.DeserializeStringFunc()(f.SetTableName),
		tableSysIDKey:      internalSerialization.DeserializeStringFunc()(f.SetTableSysID),
		updatedByNameKey:   internalSerialization.DeserializeStringFunc()(f.SetUpdatedByName),
	}
}

// SetAverageImageColor Sets the sum of all colors.
func (f *File) SetAverageImageColor(averageImageColor *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, averageImageColorKey, averageImageColor)
}

// GetCompressed return flag that indicates whether the attachment file has been compressed.
func (f *File) GetCompressed() (*bool, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, compressedKey)
}

// SetCompressed Sets flag that indicates whether the attachment file has been compressed.
func (f *File) SetCompressed(compressed *bool) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, compressedKey, compressed)
}

// GetContentType returns content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *File) GetContentType() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, contentTypeKey)
}

// SetContentType Sets content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *File) SetContentType(contentType *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, contentTypeKey, contentType)
}

// GetCreatedByName returns full name of entity that originally created the attachment file.
func (f *File) GetCreatedByName() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, createdByNameKey)
}

// SetCreatedByName Sets full name of entity that originally created the attachment file.
func (f *File) SetCreatedByName(createdByName *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, createdByNameKey, createdByName)
}

// GetCreatedBy returns the entity that originally created the attachment file.
func (f *File) GetCreatedBy() (*string, error) {
	return f.GetSysCreatedBy()
}

// SetCreatedBy Sets the entity that originally created the attachment file.
func (f *File) SetCreatedBy(createdBy *string) error {
	return f.SetSysCreatedBy(createdBy)
}

// GetDownloadLink returns download URL of the attachment on the ServiceNow instance.
func (f *File) GetDownloadLink() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, downloadLinkKey)
}

// SetDownloadLink Sets download URL of the attachment on the ServiceNow instance.
func (f *File) SetDownloadLink(downloadLink *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, downloadLinkKey, downloadLink)
}

// GetFileName returns the file name of the attachment.
func (f *File) GetFileName() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, fileNameKey)
}

// SetFileName Sets the file name of the attachment.
func (f *File) SetFileName(fileName *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, fileNameKey, fileName)
}

// GetImageHeight returns if an image file, the height of the image.
func (f *File) GetImageHeight() (*float64, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, imageHeightKey)
}

// SetImageHeight Sets if an image file, the height of the image.
func (f *File) SetImageHeight(imageHeight *float64) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, imageHeightKey, imageHeight)
}

// GetImageWidth returns if an image file, the width of the image.
func (f *File) GetImageWidth() (*float64, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, imageWidthKey)
}

// SetImageWidth Sets if an image file, the width of the image.
func (f *File) SetImageWidth(imageWidth *float64) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, imageWidthKey, imageWidth)
}

// GetSizeBytes returns size of the attachment.
func (f *File) GetSizeBytes() (*int64, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sizeBytesKey)
}

// SetSizeBytes Sets size of the attachment.
func (f *File) SetSizeBytes(sizeBytes *int64) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sizeBytesKey, sizeBytes)
}

// GetSizeCompressed returns size of the compressed attachment file. If the file is not compressed, empty.
func (f *File) GetSizeCompressed() (*int64, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sizeCompressedKey)
}

// SetSizeCompressed Sets size of the compressed attachment file.
func (f *File) SetSizeCompressed(sizeCompressed *int64) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sizeCompressedKey, sizeCompressed)
}

// GetSysCreatedBy returns the entity that originally created the attachment file.
func (f *File) GetSysCreatedBy() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysCreatedByKey)
}

// SetSysCreatedBy Sets the entity that originally created the attachment file.
func (f *File) SetSysCreatedBy(sysCreatedBy *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysCreatedByKey, sysCreatedBy)
}

// GetSysCreatedOn returns the date and time that the attachment file was initially saved to the instance.
func (f *File) GetSysCreatedOn() (*time.Time, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sysCreatedOnKey)
}

// SetSysCreatedOn Sets the date and time that the attachment file was initially saved to the instance.
func (f *File) SetSysCreatedOn(sysCreatedOn *time.Time) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysCreatedOnKey, sysCreatedOn)
}

// GetSysID returns the sys_id of the attachment file. Read-Only.
func (f *File) GetSysID() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysIDKey)
}

// SetSysID Sets the sys_id of the attachment file.
func (f *File) SetSysID(sysID *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysIDKey, sysID)
}

// GetSysModCount returns the number of times the attachment file has been modified (uploaded to the instance).
func (f *File) GetSysModCount() (*int64, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sysModCountKey)
}

// SetSysModCount Sets the number of times the attachment file has been modified (uploaded to the instance).
func (f *File) SetSysModCount(sysModCount *int64) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysModCountKey, sysModCount)
}

// GetSysTags returns any system tags associated with the attachment file.
func (f *File) GetSysTags() ([]string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []string](backingStore, sysTagsKey)
}

// SetSysTags Sets any system tags associated with the attachment file.
func (f *File) SetSysTags(sysTags []string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysTagsKey, sysTags)
}

// GetSysUpdatedBy returns the entity that last updated the attachment file.
func (f *File) GetSysUpdatedBy() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysUpdatedByKey)
}

// SetSysUpdatedBy Sets the entity that last updated the attachment file.
func (f *File) SetSysUpdatedBy(sysUpdatedBy *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysUpdatedByKey, sysUpdatedBy)
}

// GetSysUpdatedOn returns the date and time that the attachment file was last updated.
func (f *File) GetSysUpdatedOn() (*time.Time, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sysUpdatedOnKey)
}

// SetSysUpdatedOn Sets the date and time that the attachment file was last updated.
func (f *File) SetSysUpdatedOn(sysUpdatedOn *time.Time) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysUpdatedOnKey, sysUpdatedOn)
}

// GetTableName returns the name of the table to which the attachment is associated.
func (f *File) GetTableName() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, tableNameKey)
}

// SetTableName Sets the name of the table to which the attachment is associated.
func (f *File) SetTableName(tableName *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, tableNameKey, tableName)
}

// GetTableSysID returns the sys_id of the table associated with the attachment.
func (f *File) GetTableSysID() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, tableSysIDKey)
}

// SetTableSysID Sets the sys_id of the table associated with the attachment.
func (f *File) SetTableSysID(tableSysID *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, tableSysIDKey, tableSysID)
}

// GetUpdatedByName returns the full name of entity that last updated the attachment file.
func (f *File) GetUpdatedByName() (*string, error) {
	if conversion.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, updatedByNameKey)
}

// SetUpdatedByName Sets the full name of entity that last updated the attachment file.
func (f *File) SetUpdatedByName(updatedByName *string) error {
	if conversion.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, updatedByNameKey, updatedByName)
}
