package attachmentapi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
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
	store.BackedModel
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

			return writer.WriteStringValue(createdByNameKey, downloadLink)
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

			imageHeightString := strconv.FormatFloat(*imageHeight, 'f', -1, 64)

			return writer.WriteStringValue(imageHeightKey, &imageHeightString)
		},
		imageWidthKey: func(writer serialization.SerializationWriter) error {
			imageWidth, err := f.GetImageWidth()
			if err != nil {
				return err
			}

			imageWidthString := strconv.FormatFloat(*imageWidth, 'f', -1, 64)

			return writer.WriteStringValue(imageHeightKey, &imageWidthString)
		},
		sizeBytesKey: func(writer serialization.SerializationWriter) error {
			sizeBytes, err := f.GetSizeBytes()
			if err != nil {
				return err
			}

			sizeBytesString := fmt.Sprintf("%v", sizeBytes)

			return writer.WriteStringValue(sizeBytesKey, &sizeBytesString)
		},
		sizeCompressedKey: func(serialization.SerializationWriter) error {
			sizeCompressed, err := f.GetSizeCompressed()
			if err != nil {
				return err
			}

			sizeCompressedString := fmt.Sprintf("%v", sizeCompressed)

			return writer.WriteStringValue(sizeBytesKey, &sizeCompressedString)
		},
		sysCreatedByKey: func(writer serialization.SerializationWriter) error {
			sysCreatedBy, err := f.GetFileName()
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

			sysCreatedOnString := sysCreatedOn.Format(time.RFC3339)

			return writer.WriteStringValue(sysCreatedOnKey, &sysCreatedOnString)
		},
		sysIDKey: func(writer serialization.SerializationWriter) error {
			sysID, err := f.GetFileName()
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

			sysModCountString := fmt.Sprintf("%v", sysModCount)

			return writer.WriteStringValue(sizeBytesKey, &sysModCountString)
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
			sysUpdatedBy, err := f.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(sysUpdatedByKey, sysUpdatedBy)
		},
		sysUpdatedOnKey: func(writer serialization.SerializationWriter) error {
			sysUpdatedOn, err := f.GetSysCreatedOn()
			if err != nil {
				return err
			}

			sysUpdatedOnString := sysUpdatedOn.Format(time.RFC3339)

			return writer.WriteStringValue(sysUpdatedOnKey, &sysUpdatedOnString)
		},
		tableNameKey: func(writer serialization.SerializationWriter) error {
			tableName, err := f.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(tableNameKey, tableName)
		},
		tableSysIDKey: func(writer serialization.SerializationWriter) error {
			tableSysID, err := f.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(tableSysIDKey, tableSysID)
		},
		updatedByNameKey: func(writer serialization.SerializationWriter) error {
			updatedByName, err := f.GetFileName()
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

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (f *FileModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { //nolint:gocognit
	if internal.IsNil(f) {
		f = NewFile()
	}

	return map[string]func(serialization.ParseNode) error{
		averageImageColorKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetAverageImageColor(val)
		},
		compressedKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			boolVal, err := strconv.ParseBool(*val)
			if err != nil {
				return err
			}

			return f.SetCompressed(&boolVal)
		},
		contentTypeKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetContentType(val)
		},
		createdByNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetCreatedByName(val)
		},
		downloadLinkKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetDownloadLink(val)
		},
		fileNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetFileName(val)
		},
		imageHeightKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			floatVal, err := strconv.ParseFloat(*val, 64)
			if err != nil {
				return err
			}

			return f.SetImageHeight(&floatVal)
		},
		imageWidthKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			floatVal, err := strconv.ParseFloat(*val, 64)
			if err != nil {
				return err
			}

			return f.SetImageWidth(&floatVal)
		},
		sizeBytesKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			intVal, err := strconv.Atoi(*val)
			if err != nil {
				return err
			}
			int64Val := int64(intVal)

			return f.SetSizeBytes(&int64Val)
		},
		sizeCompressedKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			intVal, err := strconv.Atoi(*val)
			if err != nil {
				return err
			}
			int64Val := int64(intVal)

			return f.SetSizeCompressed(&int64Val)
		},
		sysCreatedByKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetSysCreatedBy(val)
		},
		sysCreatedOnKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			if internal.IsNil(val) || *val == "" {
				return f.SetSysUpdatedOn(nil)
			}

			dateTime, err := time.Parse("2006-01-02 15:04:05", *val)
			if err != nil {
				return err
			}

			return f.SetSysCreatedOn(&dateTime)
		},
		sysIDKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetSysID(val)
		},
		sysModCountKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			intVal, err := strconv.Atoi(*val)
			if err != nil {
				return err
			}
			int64Val := int64(intVal)

			return f.SetSysModCount(&int64Val)
		},
		sysTagsKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			// TODO: Figure out delimiter
			tags := strings.Split(*val, " ")

			return f.SetSysTags(tags)
		},
		sysUpdatedByKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetSysUpdatedBy(val)
		},
		sysUpdatedOnKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			if internal.IsNil(val) || *val == "" {
				return f.SetSysUpdatedOn(nil)
			}

			dateTime, err := time.Parse("2006-01-02 15:04:05", *val)
			if err != nil {
				return err
			}

			return f.SetSysUpdatedOn(&dateTime)
		},
		tableNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetTableName(val)
		},
		tableSysIDKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetTableSysID(val)
		},
		updatedByNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return f.SetUpdatedByName(val)
		},
	}
}

// SetAverageImageColor Sets the sum of all colors.
func (f *FileModel) SetAverageImageColor(averageImageColor *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(averageImageColorKey, averageImageColor)
}

// GetCompressed return flag that indicates whether the attachment file has been compressed.
func (f *FileModel) GetCompressed() (*bool, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*bool)
	if !ok {
		return nil, errors.New("val is not *bool")
	}

	return typedVal, nil
}

// SetCompressed Sets flag that indicates whether the attachment file has been compressed.
func (f *FileModel) SetCompressed(compressed *bool) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(compressedKey, compressed)
}

// GetContentType returns content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *FileModel) GetContentType() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(contentTypeKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetContentType Sets content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (f *FileModel) SetContentType(contentType *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(contentTypeKey, contentType)
}

// GetCreatedByName returns full name of entity that originally created the attachment file.
func (f *FileModel) GetCreatedByName() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(createdByNameKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetCreatedByName Sets full name of entity that originally created the attachment file.
func (f *FileModel) SetCreatedByName(createdByName *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(createdByNameKey, createdByName)
}

// GetDownloadLink returns download URL of the attachment on the ServiceNow instance.
func (f *FileModel) GetDownloadLink() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(downloadLinkKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetDownloadLink Sets download URL of the attachment on the ServiceNow instance.
func (f *FileModel) SetDownloadLink(downloadLink *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(downloadLinkKey, downloadLink)
}

// GetFileName returns the file name of the attachment.
func (f *FileModel) GetFileName() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(fileNameKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetFileName Sets the file name of the attachment.
func (f *FileModel) SetFileName(fileName *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(fileNameKey, fileName)
}

// GetImageHeight returns if an image file, the height of the image.
func (f *FileModel) GetImageHeight() (*float64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(imageHeightKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*float64)
	if !ok {
		return nil, errors.New("val is not *float64")
	}

	return typedVal, nil
}

// SetImageHeight Sets if an image file, the height of the image.
func (f *FileModel) SetImageHeight(imageHeight *float64) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(imageHeightKey, imageHeight)
}

// GetImageWidth returns if an image file, the width of the image.
func (f *FileModel) GetImageWidth() (*float64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(imageWidthKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*float64)
	if !ok {
		return nil, errors.New("val is not *float64")
	}

	return typedVal, nil
}

// SetImageWidth Sets if an image file, the width of the image.
func (f *FileModel) SetImageWidth(imageWidth *float64) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(imageWidthKey, imageWidth)
}

// GetSizeBytes returns size of the attachment.
func (f *FileModel) GetSizeBytes() (*int64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}

	return typedVal, nil
}

// SetSizeBytes Sets size of the attachment.
func (f *FileModel) SetSizeBytes(sizeBytes *int64) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sizeBytesKey, sizeBytes)
}

// GetSizeCompressed returns size of the compressed attachment file. If the file is not compressed, empty.
func (f *FileModel) GetSizeCompressed() (*int64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}

	return typedVal, nil
}

// SetSizeCompressed Sets size of the compressed attachment file.
func (f *FileModel) SetSizeCompressed(sizeCompressed *int64) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sizeCompressedKey, sizeCompressed)
}

// GetSysCreatedBy returns the entity that originally created the attachment file.
func (f *FileModel) GetSysCreatedBy() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetSysCreatedBy Sets the entity that originally created the attachment file.
func (f *FileModel) SetSysCreatedBy(sysCreatedBy *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sysCreatedByKey, sysCreatedBy)
}

// GetSysCreatedOn returns the date and time that the attachment file was initially saved to the instance.
func (f *FileModel) GetSysCreatedOn() (*time.Time, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*time.Time)
	if !ok {
		return nil, errors.New("val is not *time.Time")
	}

	return typedVal, nil
}

// SetSysCreatedOn Sets the date and time that the attachment file was initially saved to the instance.
func (f *FileModel) SetSysCreatedOn(sysCreatedOn *time.Time) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sysCreatedOnKey, sysCreatedOn)
}

// GetSysID returns the sys_id of the attachment file. Read-Only.
func (f *FileModel) GetSysID() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetSysID Sets the sys_id of the attachment file.
func (f *FileModel) SetSysID(sysID *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sysCreatedOnKey, sysID)
}

// GetSysModCount returns the number of times the attachment file has been modified (uploaded to the instance).
func (f *FileModel) GetSysModCount() (*int64, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}

	return typedVal, nil
}

// SetSysModCount Sets the number of times the attachment file has been modified (uploaded to the instance).
func (f *FileModel) SetSysModCount(sysModCount *int64) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sysModCountKey, sysModCount)
}

// GetSysTags returns any system tags associated with the attachment file.
func (f *FileModel) GetSysTags() ([]string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.([]string)
	if !ok {
		return nil, errors.New("val is not []string")
	}

	return typedVal, nil
}

// SetSysTags Sets any system tags associated with the attachment file.
func (f *FileModel) SetSysTags(sysTags []string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sysTagsKey, sysTags)
}

// GetSysUpdatedBy returns the entity that last updated the attachment file.
func (f *FileModel) GetSysUpdatedBy() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetSysUpdatedBy Sets the entity that last updated the attachment file.
func (f *FileModel) SetSysUpdatedBy(sysUpdatedBy *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sysUpdatedByKey, sysUpdatedBy)
}

// GetSysUpdatedOn returns the date and time that the attachment file was last updated.
func (f *FileModel) GetSysUpdatedOn() (*time.Time, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*time.Time)
	if !ok {
		return nil, errors.New("val is not *time.Time")
	}

	return typedVal, nil
}

// SetSysUpdatedOn Sets the date and time that the attachment file was last updated.
func (f *FileModel) SetSysUpdatedOn(sysUpdatedOn *time.Time) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(sysUpdatedOnKey, sysUpdatedOn)
}

// GetTableName returns the name of the table to which the attachment is associated.
func (f *FileModel) GetTableName() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetTableName Sets the name of the table to which the attachment is associated.
func (f *FileModel) SetTableName(tableName *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(tableNameKey, tableName)
}

// GetTableSysID returns the sys_id of the table associated with the attachment.
func (f *FileModel) GetTableSysID() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetTableSysID Sets the sys_id of the table associated with the attachment.
func (f *FileModel) SetTableSysID(tableSysID *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(tableSysIDKey, tableSysID)
}

// GetUpdatedByName returns the full name of entity that last updated the attachment file.
func (f *FileModel) GetUpdatedByName() (*string, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// SetUpdatedByName Sets the full name of entity that last updated the attachment file.
func (f *FileModel) SetUpdatedByName(updatedByName *string) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.GetBackingStore().Set(updatedByNameKey, updatedByName)
}
