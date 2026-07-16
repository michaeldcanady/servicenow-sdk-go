package attachmentapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
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
	core.BackedModel
}

// NewFile creates a new instance of FileModel
func NewFile() *File {
	return newFile(core.NewBaseModel())
}

// newFile creates a new instance of FileModel with the provided model underlying it
func newFile(model core.BackedModel) *File {
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
func (f *File) GetAverageImageColor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, averageImageColorKey)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (f *File) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
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
func (f *File) SetAverageImageColor(averageImageColor *string) error {
	return store.DefaultBackedModelMutatorFunc(f, averageImageColorKey, averageImageColor)
}

// GetCompressed return flag that indicates whether the attachment file has been compressed.
func (f *File) GetCompressed() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *bool](f, compressedKey)
}

// SetCompressed Sets flag that indicates whether the attachment file has been compressed.
func (f *File) SetCompressed(compressed *bool) error {
	return store.DefaultBackedModelMutatorFunc(f, compressedKey, compressed)
}

// GetContentType returns content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *File) GetContentType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, contentTypeKey)
}

// SetContentType Sets content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *File) SetContentType(contentType *string) error {
	return store.DefaultBackedModelMutatorFunc(f, contentTypeKey, contentType)
}

// GetCreatedByName returns full name of entity that originally created the attachment file.
func (f *File) GetCreatedByName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, createdByNameKey)
}

// SetCreatedByName Sets full name of entity that originally created the attachment file.
func (f *File) SetCreatedByName(createdByName *string) error {
	return store.DefaultBackedModelMutatorFunc(f, createdByNameKey, createdByName)
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
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, downloadLinkKey)
}

// SetDownloadLink Sets download URL of the attachment on the ServiceNow instance.
func (f *File) SetDownloadLink(downloadLink *string) error {
	return store.DefaultBackedModelMutatorFunc(f, downloadLinkKey, downloadLink)
}

// GetFileName returns the file name of the attachment.
func (f *File) GetFileName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, fileNameKey)
}

// SetFileName Sets the file name of the attachment.
func (f *File) SetFileName(fileName *string) error {
	return store.DefaultBackedModelMutatorFunc(f, fileNameKey, fileName)
}

// GetImageHeight returns if an image file, the height of the image.
func (f *File) GetImageHeight() (*float64, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *float64](f, imageHeightKey)
}

// SetImageHeight Sets if an image file, the height of the image.
func (f *File) SetImageHeight(imageHeight *float64) error {
	return store.DefaultBackedModelMutatorFunc(f, imageHeightKey, imageHeight)
}

// GetImageWidth returns if an image file, the width of the image.
func (f *File) GetImageWidth() (*float64, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *float64](f, imageWidthKey)
}

// SetImageWidth Sets if an image file, the width of the image.
func (f *File) SetImageWidth(imageWidth *float64) error {
	return store.DefaultBackedModelMutatorFunc(f, imageWidthKey, imageWidth)
}

// GetSizeBytes returns size of the attachment.
func (f *File) GetSizeBytes() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *int64](f, sizeBytesKey)
}

// SetSizeBytes Sets size of the attachment.
func (f *File) SetSizeBytes(sizeBytes *int64) error {
	return store.DefaultBackedModelMutatorFunc(f, sizeBytesKey, sizeBytes)
}

// GetSizeCompressed returns size of the compressed attachment file. If the file is not compressed, empty.
func (f *File) GetSizeCompressed() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *int64](f, sizeCompressedKey)
}

// SetSizeCompressed Sets size of the compressed attachment file.
func (f *File) SetSizeCompressed(sizeCompressed *int64) error {
	return store.DefaultBackedModelMutatorFunc(f, sizeCompressedKey, sizeCompressed)
}

// GetSysCreatedBy returns the entity that originally created the attachment file.
func (f *File) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, sysCreatedByKey)
}

// SetSysCreatedBy Sets the entity that originally created the attachment file.
func (f *File) SetSysCreatedBy(sysCreatedBy *string) error {
	return store.DefaultBackedModelMutatorFunc(f, sysCreatedByKey, sysCreatedBy)
}

// GetSysCreatedOn returns the date and time that the attachment file was initially saved to the instance.
func (f *File) GetSysCreatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *time.Time](f, sysCreatedOnKey)
}

// SetSysCreatedOn Sets the date and time that the attachment file was initially saved to the instance.
func (f *File) SetSysCreatedOn(sysCreatedOn *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(f, sysCreatedOnKey, sysCreatedOn)
}

// GetSysID returns the sys_id of the attachment file. Read-Only.
func (f *File) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, sysIDKey)
}

// SetSysID Sets the sys_id of the attachment file.
func (f *File) SetSysID(sysID *string) error {
	return store.DefaultBackedModelMutatorFunc(f, sysIDKey, sysID)
}

// GetSysModCount returns the number of times the attachment file has been modified (uploaded to the instance).
func (f *File) GetSysModCount() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *int64](f, sysModCountKey)
}

// SetSysModCount Sets the number of times the attachment file has been modified (uploaded to the instance).
func (f *File) SetSysModCount(sysModCount *int64) error {
	return store.DefaultBackedModelMutatorFunc(f, sysModCountKey, sysModCount)
}

// GetSysTags returns any system tags associated with the attachment file.
func (f *File) GetSysTags() ([]string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, []string](f, sysTagsKey)
}

// SetSysTags Sets any system tags associated with the attachment file.
func (f *File) SetSysTags(sysTags []string) error {
	return store.DefaultBackedModelMutatorFunc(f, sysTagsKey, sysTags)
}

// GetSysUpdatedBy returns the entity that last updated the attachment file.
func (f *File) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, sysUpdatedByKey)
}

// SetSysUpdatedBy Sets the entity that last updated the attachment file.
func (f *File) SetSysUpdatedBy(sysUpdatedBy *string) error {
	return store.DefaultBackedModelMutatorFunc(f, sysUpdatedByKey, sysUpdatedBy)
}

// GetSysUpdatedOn returns the date and time that the attachment file was last updated.
func (f *File) GetSysUpdatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *time.Time](f, sysUpdatedOnKey)
}

// SetSysUpdatedOn Sets the date and time that the attachment file was last updated.
func (f *File) SetSysUpdatedOn(sysUpdatedOn *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(f, sysUpdatedOnKey, sysUpdatedOn)
}

// GetTableName returns the name of the table to which the attachment is associated.
func (f *File) GetTableName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, tableNameKey)
}

// SetTableName Sets the name of the table to which the attachment is associated.
func (f *File) SetTableName(tableName *string) error {
	return store.DefaultBackedModelMutatorFunc(f, tableNameKey, tableName)
}

// GetTableSysID returns the sys_id of the table associated with the attachment.
func (f *File) GetTableSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, tableSysIDKey)
}

// SetTableSysID Sets the sys_id of the table associated with the attachment.
func (f *File) SetTableSysID(tableSysID *string) error {
	return store.DefaultBackedModelMutatorFunc(f, tableSysIDKey, tableSysID)
}

// GetUpdatedByName returns the full name of entity that last updated the attachment file.
func (f *File) GetUpdatedByName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*File, *string](f, updatedByNameKey)
}

// SetUpdatedByName Sets the full name of entity that last updated the attachment file.
func (f *File) SetUpdatedByName(updatedByName *string) error {
	return store.DefaultBackedModelMutatorFunc(f, updatedByNameKey, updatedByName)
}
