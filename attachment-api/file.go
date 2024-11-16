package attachmentapi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	createdByNameKey = "created_by_name"
	updatedByNameKey = "update_by_name"
)

// Fileable defines a serializable model object.
type Fileable interface {
	GetAverageImageColor() (*string, error)
	setAverageImageColor(*string) error
	GetCompressed() (*bool, error)
	setCompressed(*bool) error
	GetContentType() (*string, error)
	setContentType(*string) error
	GetCreatedByName() (*string, error)
	setCreatedByName(*string) error
	GetDownloadLink() (*string, error)
	setDownloadLink(*string) error
	GetFileName() (*string, error)
	setFileName(*string) error
	GetImageHeight() (*float64, error)
	setImageHeight(*float64) error
	GetImageWidth() (*float64, error)
	setImageWidth(*float64) error
	GetSizeBytes() (*int64, error)
	setSizeBytes(*int64) error
	GetSizeCompressed() (*int64, error)
	setSizeCompressed(*int64) error
	GetSysCreatedBy() (*string, error)
	setSysCreatedBy(*string) error
	GetSysCreatedOn() (*time.Time, error)
	setSysCreatedOn(*time.Time) error
	GetSysID() (*string, error)
	setSysID(*string) error
	GetSysModCount() (*int64, error)
	setSysModCount(*int64) error
	GetSysTags() ([]string, error)
	setSysTags([]string) error
	GetSysUpdatedBy() (*string, error)
	setSysUpdatedBy(*string) error
	GetSysUpdatedOn() (*time.Time, error)
	setSysUpdatedOn(*time.Time) error
	GetTableName() (*string, error)
	setTableName(*string) error
	GetTableSysID() (*string, error)
	setTableSysID(*string) error
	GetUpdatedByName() (*string, error)
	setUpdatedByName(*string) error
	serialization.Parsable
	store.BackedModel
}

type file struct {
	backingStore store.BackingStore
}

func NewFile() Fileable {
	return &file{
		backingStore: store.NewInMemoryBackingStore(),
	}
}

// CreateFileFromDiscriminatorValue is a parsable factory for creating a Fileable
func CreateFileFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewFile(), nil
}

// GetBackingStore retrieves the backing store for the model.
func (rE *file) GetBackingStore() store.BackingStore {
	if internal.IsNil(rE) {
		return nil
	}

	if internal.IsNil(rE.backingStore) {
		rE.backingStore = store.NewInMemoryBackingStore()
	}

	return rE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (rE *file) Serialize(writer serialization.SerializationWriter) error { //nolint:gocognit
	if internal.IsNil(rE) {
		return nil
	}

	fieldSerializers := map[string]func(serialization.SerializationWriter) error{
		averageImageColorKey: func(writer serialization.SerializationWriter) error {
			averageImageColor, err := rE.GetAverageImageColor()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(averageImageColorKey, averageImageColor)
		},
		compressedKey: func(writer serialization.SerializationWriter) error {
			compressed, err := rE.GetCompressed()
			if err != nil {
				return err
			}
			compressedString := fmt.Sprintf("%v", *compressed)

			return writer.WriteStringValue(compressedKey, &compressedString)
		},
		contentTypeKey: func(writer serialization.SerializationWriter) error {
			contentType, err := rE.GetContentType()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(contentTypeKey, contentType)
		},
		createdByNameKey: func(writer serialization.SerializationWriter) error {
			createdByName, err := rE.GetCreatedByName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(createdByNameKey, createdByName)
		},
		downloadLinkKey: func(writer serialization.SerializationWriter) error {
			downloadLink, err := rE.GetDownloadLink()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(createdByNameKey, downloadLink)
		},
		fileNameKey: func(writer serialization.SerializationWriter) error {
			fileName, err := rE.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(fileNameKey, fileName)
		},
		imageHeightKey: func(writer serialization.SerializationWriter) error {
			imageHeight, err := rE.GetImageHeight()
			if err != nil {
				return err
			}

			imageHeightString := strconv.FormatFloat(*imageHeight, 'f', -1, 64)

			return writer.WriteStringValue(imageHeightKey, &imageHeightString)
		},
		imageWidthKey: func(writer serialization.SerializationWriter) error {
			imageWidth, err := rE.GetImageWidth()
			if err != nil {
				return err
			}

			imageWidthString := strconv.FormatFloat(*imageWidth, 'f', -1, 64)

			return writer.WriteStringValue(imageHeightKey, &imageWidthString)
		},
		sizeBytesKey: func(writer serialization.SerializationWriter) error {
			sizeBytes, err := rE.GetSizeBytes()
			if err != nil {
				return err
			}

			sizeBytesString := fmt.Sprintf("%v", sizeBytes)

			return writer.WriteStringValue(sizeBytesKey, &sizeBytesString)
		},
		sizeCompressedKey: func(serialization.SerializationWriter) error {
			sizeCompressed, err := rE.GetSizeCompressed()
			if err != nil {
				return err
			}

			sizeCompressedString := fmt.Sprintf("%v", sizeCompressed)

			return writer.WriteStringValue(sizeBytesKey, &sizeCompressedString)
		},
		sysCreatedByKey: func(writer serialization.SerializationWriter) error {
			sysCreatedBy, err := rE.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(sysCreatedByKey, sysCreatedBy)
		},
		sysCreatedOnKey: func(writer serialization.SerializationWriter) error {
			sysCreatedOn, err := rE.GetSysCreatedOn()
			if err != nil {
				return err
			}

			sysCreatedOnString := sysCreatedOn.Format(time.RFC3339)

			return writer.WriteStringValue(sysCreatedOnKey, &sysCreatedOnString)
		},
		sysIDKey: func(writer serialization.SerializationWriter) error {
			sysID, err := rE.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(sysIDKey, sysID)
		},
		sysModCountKey: func(writer serialization.SerializationWriter) error {
			sysModCount, err := rE.GetSysModCount()
			if err != nil {
				return err
			}

			sysModCountString := fmt.Sprintf("%v", sysModCount)

			return writer.WriteStringValue(sizeBytesKey, &sysModCountString)
		},
		sysTagsKey: func(writer serialization.SerializationWriter) error {
			sysTags, err := rE.GetSysTags()
			if err != nil {
				return err
			}

			// TODO: confirm file separator
			sysTagsString := strings.Join(sysTags, " ")

			return writer.WriteStringValue(sysTagsKey, &sysTagsString)
		},
		sysUpdatedByKey: func(writer serialization.SerializationWriter) error {
			sysUpdatedBy, err := rE.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(sysUpdatedByKey, sysUpdatedBy)
		},
		sysUpdatedOnKey: func(writer serialization.SerializationWriter) error {
			sysUpdatedOn, err := rE.GetSysCreatedOn()
			if err != nil {
				return err
			}

			sysUpdatedOnString := sysUpdatedOn.Format(time.RFC3339)

			return writer.WriteStringValue(sysUpdatedOnKey, &sysUpdatedOnString)
		},
		tableNameKey: func(writer serialization.SerializationWriter) error {
			tableName, err := rE.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(tableNameKey, tableName)
		},
		tableSysIDKey: func(writer serialization.SerializationWriter) error {
			tableSysID, err := rE.GetFileName()
			if err != nil {
				return err
			}

			return writer.WriteStringValue(tableSysIDKey, tableSysID)
		},
		updatedByNameKey: func(writer serialization.SerializationWriter) error {
			updatedByName, err := rE.GetFileName()
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

// GetFieldDeserializers returns the deserialization information for this object.
func (rE *file) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { //nolint:gocognit
	if internal.IsNil(rE) {
		rE = NewFile().(*file)
	}

	return map[string]func(serialization.ParseNode) error{
		averageImageColorKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setAverageImageColor(val)
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

			return rE.setCompressed(&boolVal)
		},
		contentTypeKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setContentType(val)
		},
		createdByNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setCreatedByName(val)
		},
		downloadLinkKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setDownloadLink(val)
		},
		fileNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setFileName(val)
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

			return rE.setImageHeight(&floatVal)
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

			return rE.setImageWidth(&floatVal)
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

			return rE.setSizeBytes(&int64Val)
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

			return rE.setSizeCompressed(&int64Val)
		},
		sysCreatedByKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setSysCreatedBy(val)
		},
		sysCreatedOnKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			if internal.IsNil(val) || *val == "" {
				return rE.setSysUpdatedOn(nil)
			}

			dateTime, err := time.Parse("2006-01-02 15:04:05", *val)
			if err != nil {
				return err
			}

			return rE.setSysCreatedOn(&dateTime)
		},
		sysIDKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setSysID(val)
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

			return rE.setSysModCount(&int64Val)
		},
		sysTagsKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			// TODO: Figure out delimiter
			tags := strings.Split(*val, " ")

			return rE.setSysTags(tags)
		},
		sysUpdatedByKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setSysUpdatedBy(val)
		},
		sysUpdatedOnKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			if internal.IsNil(val) || *val == "" {
				return rE.setSysUpdatedOn(nil)
			}

			dateTime, err := time.Parse("2006-01-02 15:04:05", *val)
			if err != nil {
				return err
			}

			return rE.setSysUpdatedOn(&dateTime)
		},
		tableNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setTableName(val)
		},
		tableSysIDKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setTableSysID(val)
		},
		updatedByNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}

			return rE.setUpdatedByName(val)
		},
	}
}

// GetAverageImageColor returns, If the attachment is an image, the sum of all colors.
func (rE *file) GetAverageImageColor() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setAverageImageColor sets the sum of all colors.
func (rE *file) setAverageImageColor(averageImageColor *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(averageImageColorKey, averageImageColor)
}

// GetCompressed return flag that indicates whether the attachment file has been compressed.
func (rE *file) GetCompressed() (*bool, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*bool)
	if !ok {
		return nil, errors.New("val is not *bool")
	}

	return typedVal, nil
}

// setCompressed sets flag that indicates whether the attachment file has been compressed.
func (rE *file) setCompressed(compressed *bool) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(compressedKey, compressed)
}

// GetContentType returns content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (rE *file) GetContentType() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(contentTypeKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setContentType sets content-type of the associated attachment file, such as image or jpeg or application/x-shockwave-flash.
func (rE *file) setContentType(contentType *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(contentTypeKey, contentType)
}

// GetCreatedByName returns full name of entity that originally created the attachment file.
func (rE *file) GetCreatedByName() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(createdByNameKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setCreatedByName sets full name of entity that originally created the attachment file.
func (rE *file) setCreatedByName(createdByName *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(createdByNameKey, createdByName)
}

// GetDownloadLink returns download URL of the attachment on the ServiceNow instance.
func (rE *file) GetDownloadLink() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(downloadLinkKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setDownloadLink sets download URL of the attachment on the ServiceNow instance.
func (rE *file) setDownloadLink(downloadLink *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(downloadLinkKey, downloadLink)
}

// GetFileName returns the file name of the attachment.
func (rE *file) GetFileName() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(fileNameKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setFileName sets the file name of the attachment.
func (rE *file) setFileName(fileName *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(fileNameKey, fileName)
}

// GetImageHeight returns if an image file, the height of the image.
func (rE *file) GetImageHeight() (*float64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(imageHeightKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*float64)
	if !ok {
		return nil, errors.New("val is not *float64")
	}

	return typedVal, nil
}

// setImageHeight sets if an image file, the height of the image.
func (rE *file) setImageHeight(imageHeight *float64) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(imageHeightKey, imageHeight)
}

// GetImageWidth returns if an image file, the width of the image.
func (rE *file) GetImageWidth() (*float64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(imageWidthKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*float64)
	if !ok {
		return nil, errors.New("val is not *float64")
	}

	return typedVal, nil
}

// setImageWidth sets if an image file, the width of the image.
func (rE *file) setImageWidth(imageWidth *float64) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(imageWidthKey, imageWidth)
}

// GetSizeBytes returns size of the attachment.
func (rE *file) GetSizeBytes() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}

	return typedVal, nil
}

// setSizeBytes sets size of the attachment.
func (rE *file) setSizeBytes(sizeBytes *int64) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sizeBytesKey, sizeBytes)
}

// GetSizeCompressed returns size of the compressed attachment file. If the file is not compressed, empty.
func (rE *file) GetSizeCompressed() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}

	return typedVal, nil
}

// setSizeCompressed sets size of the compressed attachment file.
func (rE *file) setSizeCompressed(sizeCompressed *int64) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sizeCompressedKey, sizeCompressed)
}

// GetSysCreatedBy returns the entity that originally created the attachment file.
func (rE *file) GetSysCreatedBy() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setSysCreatedBy sets the entity that originally created the attachment file.
func (rE *file) setSysCreatedBy(sysCreatedBy *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sysCreatedByKey, sysCreatedBy)
}

// GetSysCreatedOn returns the date and time that the attachment file was initially saved to the instance.
func (rE *file) GetSysCreatedOn() (*time.Time, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*time.Time)
	if !ok {
		return nil, errors.New("val is not *time.Time")
	}

	return typedVal, nil
}

// setSysCreatedOn sets the date and time that the attachment file was initially saved to the instance.
func (rE *file) setSysCreatedOn(sysCreatedOn *time.Time) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sysCreatedOnKey, sysCreatedOn)
}

// GetSysID returns the sys_id of the attachment file. Read-Only.
func (rE *file) GetSysID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setSysID sets the sys_id of the attachment file.
func (rE *file) setSysID(sysID *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sysCreatedOnKey, sysID)
}

// GetSysModCount returns the number of times the attachment file has been modified (uploaded to the instance).
func (rE *file) GetSysModCount() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}

	return typedVal, nil
}

// setSysModCount sets the number of times the attachment file has been modified (uploaded to the instance).
func (rE *file) setSysModCount(sysModCount *int64) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sysModCountKey, sysModCount)
}

// GetSysTags returns any system tags associated with the attachment file.
func (rE *file) GetSysTags() ([]string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.([]string)
	if !ok {
		return nil, errors.New("val is not []string")
	}

	return typedVal, nil
}

// setSysTags sets any system tags associated with the attachment file.
func (rE *file) setSysTags(sysTags []string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sysTagsKey, sysTags)
}

// GetSysUpdatedBy returns the entity that last updated the attachment file.
func (rE *file) GetSysUpdatedBy() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setSysUpdatedBy sets the entity that last updated the attachment file.
func (rE *file) setSysUpdatedBy(sysUpdatedBy *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sysUpdatedByKey, sysUpdatedBy)
}

// GetSysUpdatedOn returns the date and time that the attachment file was last updated.
func (rE *file) GetSysUpdatedOn() (*time.Time, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*time.Time)
	if !ok {
		return nil, errors.New("val is not *time.Time")
	}

	return typedVal, nil
}

// setSysUpdatedOn sets the date and time that the attachment file was last updated.
func (rE *file) setSysUpdatedOn(sysUpdatedOn *time.Time) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(sysUpdatedOnKey, sysUpdatedOn)
}

// GetTableName returns the name of the table to which the attachment is associated.
func (rE *file) GetTableName() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setTableName sets the name of the table to which the attachment is associated.
func (rE *file) setTableName(tableName *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(tableNameKey, tableName)
}

// GetTableSysID returns the sys_id of the table associated with the attachment.
func (rE *file) GetTableSysID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setTableSysID sets the sys_id of the table associated with the attachment.
func (rE *file) setTableSysID(tableSysID *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(tableSysIDKey, tableSysID)
}

// GetUpdatedByName returns the full name of entity that last updated the attachment file.
func (rE *file) GetUpdatedByName() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	val, err := rE.GetBackingStore().Get(averageImageColorKey)
	if err != nil {
		return nil, err
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

// setUpdatedByName sets the full name of entity that last updated the attachment file.
func (rE *file) setUpdatedByName(updatedByName *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(updatedByNameKey, updatedByName)
}
