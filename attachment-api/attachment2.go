package attachmentapi

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	sysIDKey             = "sys_id"
	tableSysIDKey        = "table_sys_id"
	sizeBytesKey         = "size_bytes"
	downloadLinkKey      = "download_link"
	sysUpdatedOnKey      = "sys_updated_on"
	imageHeightKey       = "image_height"
	sysCreatedOnKey      = "sys_created_on"
	fileNameKey          = "file_name"
	sysCreatedByKey      = "sys_created_by"
	compressedKey        = "compressed"
	averageImageColorKey = "average_image_color"
	sysUpdatedByKey      = "sys_updated_by"
	sysTagsKey           = "sys_tags"
	tableNameKey         = "table_name"
	imageWidthKey        = "image_width"
	sysModCountKey       = "sys_mod_count"
	contentTypeKey       = "content_type"
	sizeCompressedKey    = "size_compressed"
	chunkSizeBytesKey    = "chunk_size_bytes"
	hashKey              = "hash"
	stateKey             = "state"
)

// Attachment2 represents Service-Now attachment
type Attachment2 interface {
	GetTableSysID() (*string, error)
	setTableSysID(*string) error
	GetSizeBytes() (*int64, error)
	setSizeBytes(*int64) error
	GetDownloadLink() (*string, error)
	setDownloadLink(*string) error
	GetSysUpdatedOn() (*time.Time, error)
	setSysUpdatedOn(*time.Time) error
	GetSysID() (*string, error)
	setSysID(*string) error
	GetImageHeight() (*float64, error)
	setImageHeight(*float64) error
	GetSysCreatedOn() (*time.Time, error)
	setSysCreatedOn(*time.Time) error
	GetFileName() (*string, error)
	setFileName(*string) error
	GetSysCreatedBy() (*string, error)
	setSysCreatedBy(*string) error
	GetCompressed() (*bool, error)
	setCompressed(*bool) error
	GetAverageImageColor() (*string, error)
	setAverageImageColor(*string) error
	GetSysUpdatedBy() (*string, error)
	setSysUpdatedBy(*string) error
	GetSysTags() ([]string, error)
	setSysTags([]string) error
	GetTableName() (*string, error)
	setTableName(*string) error
	GetImageWidth() (*float64, error)
	setImageWidth(*float64) error
	GetSysModCount() (*int64, error)
	setSysModCount(*int64) error
	GetContentType() (*string, error)
	setContentType(*string) error
	GetSizeCompressed() (*int64, error)
	setSizeCompressed(*int64) error
	GetChunkSizeBytes() (*string, error)
	setChunkSizeBytes(*string) error
	GetHash() (*string, error)
	setHash(*string) error
	GetState() (*string, error)
	setState(*string) error

	serialization.Parsable
	store.BackedModel
}

// Attachment2Model implementation of Attachment2
type Attachment2Model struct {
	newInternal.Model
}

// NewAttachment creates a new instance of Attachment2Model
func NewAttachment2() *Attachment2Model {
	return newAttachment2(newInternal.NewBaseModel())
}

// newAttachment2 creates a new instance of Attachment2Model with the provided model underlying it
func newAttachment2(model newInternal.Model) *Attachment2Model {
	return &Attachment2Model{
		model,
	}
}

// CreateAttachment2FromDiscriminatorValue is a parsable factory for creating an Attachment2Model
func CreateAttachment2FromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAttachment2(), nil
}

// Serialize writes the objects properties to the current writer.
func (rE *Attachment2Model) Serialize(_ serialization.SerializationWriter) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("serialization not supported")
}

// Attachment2Model returns the deserialization information for this object.
func (rE *Attachment2Model) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { //nolint:gocognit
	if internal.IsNil(rE) {
		rE = NewAttachment2()
	}

	return map[string]func(serialization.ParseNode) error{
		tableSysIDKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setTableSysID(val)
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
		downloadLinkKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setDownloadLink(val)
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
		sysIDKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setSysID(val)
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
		fileNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setFileName(val)
		},
		sysCreatedByKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setSysCreatedBy(val)
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
		averageImageColorKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setAverageImageColor(val)
		},
		sysUpdatedByKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setSysUpdatedBy(val)
		},
		sysTagsKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			if internal.IsNil(val) || *val == "" {
				return rE.setSysTags(nil)
			}

			// TODO: figure out separator
			sliceVal := strings.Split(*val, " ")
			return rE.setSysTags(sliceVal)
		},
		tableNameKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setTableName(val)
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
		contentTypeKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setContentType(val)
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
		chunkSizeBytesKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setChunkSizeBytes(val)
		},
		hashKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setHash(val)
		},
		stateKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setState(val)
		},
	}
}

// GetTableSysID returns the table sys id
func (rE *Attachment2Model) GetTableSysID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(tableSysIDKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}
	return typedVal, nil
}

// setTableSysID sets the table sys id to the provide value
func (rE *Attachment2Model) setTableSysID(tableSysID *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(tableSysIDKey, tableSysID)
}

// GetSizeBytes returns the attachment's size in bytes
func (rE *Attachment2Model) GetSizeBytes() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sizeBytesKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}
	return typedVal, nil
}

// setSizeBytes sets the size (in bytes) to the provided value
func (rE *Attachment2Model) setSizeBytes(size *int64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sizeBytesKey, size)
}

// GetDownloadLink returns the download link
func (rE *Attachment2Model) GetDownloadLink() (*string, error) {
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

// setDownloadLink sets the download link to the provided value
func (rE *Attachment2Model) setDownloadLink(link *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(downloadLinkKey, link)
}

// GetSysUpdatedOn return the last updated timestamp
func (rE *Attachment2Model) GetSysUpdatedOn() (*time.Time, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sysUpdatedOnKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*time.Time)
	if !ok {
		return nil, errors.New("val is not *time.Time")
	}
	return typedVal, nil
}

// setSysUpdatedOn sets the last updated timestamp to the provided value
func (rE *Attachment2Model) setSysUpdatedOn(val *time.Time) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysUpdatedOnKey, val)
}

// GetSysID returns the sys id
func (rE *Attachment2Model) GetSysID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sysIDKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}
	return typedVal, nil
}

// setSysID sets the sys id to the provide value
func (rE *Attachment2Model) setSysID(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysIDKey, val)
}

// GetImageHeight returns the image's height, if the attachment is an image
func (rE *Attachment2Model) GetImageHeight() (*float64, error) {
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

// setImageHeight sets the image height to the provided value
func (rE *Attachment2Model) setImageHeight(val *float64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(imageHeightKey, val)
}

// GetSysCreatedOn returns the created on timestamp
func (rE *Attachment2Model) GetSysCreatedOn() (*time.Time, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sysCreatedOnKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*time.Time)
	if !ok {
		return nil, errors.New("val is not *time.Time")
	}
	return typedVal, nil
}

// setSysCreatedOn sets the created on timestamp to the provided value
func (rE *Attachment2Model) setSysCreatedOn(timestamp *time.Time) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysCreatedOnKey, timestamp)
}

// GetFileName returns the file name
func (rE *Attachment2Model) GetFileName() (*string, error) {
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

// setFileName sets the file name to the provide value
func (rE *Attachment2Model) setFileName(name *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(fileNameKey, name)
}

// GetSysCreatedBy returns the username of who created it
func (rE *Attachment2Model) GetSysCreatedBy() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sysCreatedByKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}
	return typedVal, nil
}

// setSysCreatedBy sets the username of who created it to the provided value
func (rE *Attachment2Model) setSysCreatedBy(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysCreatedByKey, val)
}

// GetCompressed returns if the attachment is compressed
func (rE *Attachment2Model) GetCompressed() (*bool, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(compressedKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*bool)
	if !ok {
		return nil, errors.New("val is not *bool")
	}
	return typedVal, nil
}

// setCompressed sets if the attachment is compressed
func (rE *Attachment2Model) setCompressed(compressed *bool) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(compressedKey, compressed)
}

// GetAverageImageColor returns the average image color, if an image
func (rE *Attachment2Model) GetAverageImageColor() (*string, error) {
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

// setAverageImageColor sets the average image color to the provided value
func (rE *Attachment2Model) setAverageImageColor(color *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(averageImageColorKey, color)
}

// GetSysUpdatedBy returns the username of the account that last updated the attachment
func (rE *Attachment2Model) GetSysUpdatedBy() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sysUpdatedByKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}
	return typedVal, nil
}

// setSysUpdatedBy sets the username of the account that last updated the attachment to the provide value
func (rE *Attachment2Model) setSysUpdatedBy(username *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysUpdatedByKey, username)
}

// GetSysTags returns slice of tags
func (rE *Attachment2Model) GetSysTags() ([]string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sysTagsKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.([]string)
	if !ok {
		return nil, errors.New("val is not []string")
	}
	return typedVal, nil
}

// setSysTags sets the sys tags to the provided values
func (rE *Attachment2Model) setSysTags(tags []string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysTagsKey, tags)
}

// GetTableName returns associated table name
func (rE *Attachment2Model) GetTableName() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(tableNameKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}
	return typedVal, nil
}

// setTableName sets table name to provided value
func (rE *Attachment2Model) setTableName(name *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(tableNameKey, name)
}

// GetImageWidth returns the width, if attachment is an image
func (rE *Attachment2Model) GetImageWidth() (*float64, error) {
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

// setImageWidth sets the width to the provide value
func (rE *Attachment2Model) setImageWidth(width *float64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(imageWidthKey, width)
}

// GetSysModCount returns the mod count
func (rE *Attachment2Model) GetSysModCount() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sysModCountKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}
	return typedVal, nil
}

// setSysModCount sets the count to the provided value
func (rE *Attachment2Model) setSysModCount(count *int64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysModCountKey, count)
}

// GetContentType returns the content type of the attachment
func (rE *Attachment2Model) GetContentType() (*string, error) {
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

// setContentType sets the content type to the provided value
func (rE *Attachment2Model) setContentType(contentType *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(contentTypeKey, contentType)
}

// GetSizeCompressed returns compressed size of attachment
func (rE *Attachment2Model) GetSizeCompressed() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(sizeCompressedKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*int64)
	if !ok {
		return nil, errors.New("val is not *int64")
	}
	return typedVal, nil
}

// setSizeCompressed sets compressed size to provided value
func (rE *Attachment2Model) setSizeCompressed(size *int64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sizeCompressedKey, size)
}

// GetChunkSizeBytes returns chunk size (in bytes) of attachment
func (rE *Attachment2Model) GetChunkSizeBytes() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(chunkSizeBytesKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}
	return typedVal, nil
}

// setChunkSizeBytes sets chunk size to provided value
func (rE *Attachment2Model) setChunkSizeBytes(size *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(chunkSizeBytesKey, size)
}

// GetHash returns hash of attachment
func (rE *Attachment2Model) GetHash() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(hashKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}
	return typedVal, nil
}

// setHash sets has to provided value
func (rE *Attachment2Model) setHash(hash *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(hashKey, hash)
}

func (rE *Attachment2Model) GetState() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	val, err := rE.GetBackingStore().Get(stateKey)
	if err != nil {
		return nil, err
	}
	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}
	return typedVal, nil
}

func (rE *Attachment2Model) setState(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(hashKey, val)
}
