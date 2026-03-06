package attachmentapi

import (
	"fmt"
	"strconv"
	"strings"
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

	fieldSerializers := map[string]func(serialization.SerializationWriter) error{
		averageImageColorKey: func(writer serialization.SerializationWriter) error {
			averageImageColor, err := f.GetAverageImageColor()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(averageImageColorKey, averageImageColor)
		},
		compressedKey: func(writer serialization.SerializationWriter) error {
			compressed, err := f.GetCompressed()
			if err != nil {
				return err
			}

			if compressed == nil {
				return nil
			}

			compressedString := fmt.Sprintf("%v", *compressed)

			return writer.WriteStringValue(compressedKey, &compressedString)
		},
		contentTypeKey: func(writer serialization.SerializationWriter) error {
			contentType, err := f.GetContentType()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(contentTypeKey, contentType)
		},
		createdByNameKey: func(writer serialization.SerializationWriter) error {
			createdByName, err := f.GetCreatedByName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(createdByNameKey, createdByName)
		},
		downloadLinkKey: func(writer serialization.SerializationWriter) error {
			downloadLink, err := f.GetDownloadLink()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(downloadLinkKey, downloadLink)
		},
		fileNameKey: func(writer serialization.SerializationWriter) error {
			fileName, err := f.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(fileNameKey, fileName)
		},
		imageHeightKey: func(writer serialization.SerializationWriter) error {
			imageHeight, err := f.GetImageHeight()
			if err != nil {
				return err
			}

			if imageHeight == nil {
				return nil
			}

			imageHeightString := strconv.FormatFloat(*imageHeight, 'f', -1, 64)

			return writer.WriteStringValue(imageHeightKey, &imageHeightString)
		},
		imageWidthKey: func(writer serialization.SerializationWriter) error {
			imageWidth, err := f.GetImageWidth()
			if err != nil {
				return err
			}

			if imageWidth == nil {
				return nil
			}

			imageWidthString := strconv.FormatFloat(*imageWidth, 'f', -1, 64)

			return writer.WriteStringValue(imageWidthKey, &imageWidthString)
		},
		sizeBytesKey: func(writer serialization.SerializationWriter) error {
			sizeBytes, err := f.GetSizeBytes()
			if err != nil {
				return err
			}

			if sizeBytes == nil {
				return nil
			}

			sizeBytesString := fmt.Sprintf("%v", *sizeBytes)

			return writer.WriteStringValue(sizeBytesKey, &sizeBytesString)
		},
		sizeCompressedKey: func(writer serialization.SerializationWriter) error {
			sizeCompressed, err := f.GetSizeCompressed()
			if err != nil {
				return err
			}

			if sizeCompressed == nil {
				return nil
			}

			sizeCompressedString := fmt.Sprintf("%v", *sizeCompressed)

			return writer.WriteStringValue(sizeCompressedKey, &sizeCompressedString)
		},
		sysCreatedByKey: func(writer serialization.SerializationWriter) error {
			sysCreatedBy, err := f.GetSysCreatedBy()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(sysCreatedByKey, sysCreatedBy)
		},
		sysCreatedOnKey: func(writer serialization.SerializationWriter) error {
			sysCreatedOn, err := f.GetSysCreatedOn()
			if err != nil {
				return err
			}

			if sysCreatedOn == nil {
				return nil
			}

			sysCreatedOnString := sysCreatedOn.Format(time.RFC3339)

			return writer.WriteStringValue(sysCreatedOnKey, &sysCreatedOnString)
		},
		sysIDKey: func(writer serialization.SerializationWriter) error {
			sysID, err := f.GetSysID()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(sysIDKey, sysID)
		},
		sysModCountKey: func(writer serialization.SerializationWriter) error {
			sysModCount, err := f.GetSysModCount()
			if err != nil {
				return err
			}

			if sysModCount == nil {
				return nil
			}

			sysModCountString := fmt.Sprintf("%v", *sysModCount)

			return writer.WriteStringValue(sysModCountKey, &sysModCountString)
		},
		sysTagsKey: func(writer serialization.SerializationWriter) error {
			sysTags, err := f.GetSysTags()
			if err != nil {
				return err
			}

			// TODO: confirm file separator
			sysTagsString := strings.Join(sysTags, " ")

			return writer.WriteStringValue(sysTagsKey, &sysTagsString)
		},
		sysUpdatedByKey: func(writer serialization.SerializationWriter) error {
			sysUpdatedBy, err := f.GetSysUpdatedBy()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(sysUpdatedByKey, sysUpdatedBy)
		},
		sysUpdatedOnKey: func(writer serialization.SerializationWriter) error {
			sysUpdatedOn, err := f.GetSysUpdatedOn()
			if err != nil {
				return err
			}

			if sysUpdatedOn == nil {
				return nil
			}

			sysUpdatedOnString := sysUpdatedOn.Format(time.RFC3339)

			return writer.WriteStringValue(sysUpdatedOnKey, &sysUpdatedOnString)
		},
		tableNameKey: func(writer serialization.SerializationWriter) error {
			tableName, err := f.GetTableName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(tableNameKey, tableName)
		},
		tableSysIDKey: func(writer serialization.SerializationWriter) error {
			tableSysID, err := f.GetTableSysID()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(tableSysIDKey, tableSysID)
		},
		updatedByNameKey: func(writer serialization.SerializationWriter) error {
			updatedByName, err := f.GetUpdatedByName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(updatedByNameKey, updatedByName)
		},
	}

	for _, serializer := range fieldSerializers {
		if err := serializer(writer); err != nil {
			return err
		}
	}

	return nil
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
		averageImageColorKey: internalSerialization.DeserializeStringFunc(f.SetAverageImageColor),
		compressedKey:        internalSerialization.DeserializeMutatedStringFunc(f.SetCompressed, conversion.StringPtrToBoolPtr),
		contentTypeKey:       internalSerialization.DeserializeStringFunc(f.SetContentType),
		createdByNameKey:     internalSerialization.DeserializeStringFunc(f.SetCreatedByName),
		downloadLinkKey:      internalSerialization.DeserializeStringFunc(f.SetDownloadLink),
		fileNameKey:          internalSerialization.DeserializeStringFunc(f.SetFileName),
		imageHeightKey:       internalSerialization.DeserializeMutatedStringFunc(f.SetImageHeight, conversion.StringPtrToFloat64Ptr),
		imageWidthKey:        internalSerialization.DeserializeMutatedStringFunc(f.SetImageWidth, conversion.StringPtrToFloat64Ptr),
		sizeBytesKey:         internalSerialization.DeserializeMutatedStringFunc(f.SetSizeBytes, conversion.StringPtrToInt64Ptr),
		sizeCompressedKey:    internalSerialization.DeserializeMutatedStringFunc(f.SetSizeCompressed, conversion.StringPtrToInt64Ptr),
		sysCreatedByKey:      internalSerialization.DeserializeStringFunc(f.SetSysCreatedBy),
		sysCreatedOnKey:      internalSerialization.DeserializeMutatedStringFunc(f.SetSysCreatedOn, conversion.StringPtrToTimePtr("2006-01-02 15:04:05")),
		sysIDKey:             internalSerialization.DeserializeStringFunc(f.SetSysID),
		sysModCountKey:       internalSerialization.DeserializeMutatedStringFunc(f.SetSysModCount, conversion.StringPtrToInt64Ptr),
		sysTagsKey:           internalSerialization.DeserializeMutatedStringFunc(f.SetSysTags, conversion.StringPtrToPrimitiveSlice(" ", func(s string) (string, error) { return s, nil })),
		sysUpdatedByKey:      internalSerialization.DeserializeStringFunc(f.SetSysUpdatedBy),
		sysUpdatedOnKey:      internalSerialization.DeserializeMutatedStringFunc(f.SetSysUpdatedOn, conversion.StringPtrToTimePtr("2006-01-02 15:04:05")),
		tableNameKey:         internalSerialization.DeserializeStringFunc(f.SetTableName),
		tableSysIDKey:        internalSerialization.DeserializeStringFunc(f.SetTableSysID),
		updatedByNameKey:     internalSerialization.DeserializeStringFunc(f.SetUpdatedByName),
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
