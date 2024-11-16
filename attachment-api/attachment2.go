package attachmentapi

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
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

type Attachmentable interface {
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

// attachment...
type attachment struct {
	backingStore store.BackingStore
}

func NewAttachment() Attachmentable {
	return &attachment{
		backingStore: store.NewInMemoryBackingStore(),
	}
}

// CreateAttachmentFromDiscriminatorValue is a parsable factory for creating a Attachmentable
func CreateAttachmentFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewAttachment(), nil
}

// GetBackingStore retrieves the backing store for the model.
func (rE *attachment) GetBackingStore() store.BackingStore {
	if internal.IsNil(rE) {
		return nil
	}

	if internal.IsNil(rE.backingStore) {
		rE.backingStore = store.NewInMemoryBackingStore()
	}

	return rE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (rE *attachment) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (rE *attachment) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { //nolint:gocognit
	if internal.IsNil(rE) {
		rE = NewAttachment().(*attachment)
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

func (rE *attachment) GetTableSysID() (*string, error) {
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

func (rE *attachment) setTableSysID(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(tableSysIDKey, val)
}

func (rE *attachment) GetSizeBytes() (*int64, error) {
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

func (rE *attachment) setSizeBytes(val *int64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sizeBytesKey, val)
}

func (rE *attachment) GetDownloadLink() (*string, error) {
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

func (rE *attachment) setDownloadLink(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(downloadLinkKey, val)
}

func (rE *attachment) GetSysUpdatedOn() (*time.Time, error) {
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

func (rE *attachment) setSysUpdatedOn(val *time.Time) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysUpdatedOnKey, val)
}

func (rE *attachment) GetSysID() (*string, error) {
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

func (rE *attachment) setSysID(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysIDKey, val)
}

func (rE *attachment) GetImageHeight() (*float64, error) {
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

func (rE *attachment) setImageHeight(val *float64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(imageHeightKey, val)
}

func (rE *attachment) GetSysCreatedOn() (*time.Time, error) {
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

func (rE *attachment) setSysCreatedOn(val *time.Time) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysCreatedOnKey, val)
}

func (rE *attachment) GetFileName() (*string, error) {
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

func (rE *attachment) setFileName(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(fileNameKey, val)
}

func (rE *attachment) GetSysCreatedBy() (*string, error) {
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

func (rE *attachment) setSysCreatedBy(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysCreatedByKey, val)
}

func (rE *attachment) GetCompressed() (*bool, error) {
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

func (rE *attachment) setCompressed(val *bool) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(compressedKey, val)
}

func (rE *attachment) GetAverageImageColor() (*string, error) {
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

func (rE *attachment) setAverageImageColor(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(averageImageColorKey, val)
}

func (rE *attachment) GetSysUpdatedBy() (*string, error) {
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

func (rE *attachment) setSysUpdatedBy(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysUpdatedByKey, val)
}

func (rE *attachment) GetSysTags() ([]string, error) {
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

func (rE *attachment) setSysTags(val []string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysTagsKey, val)
}

func (rE *attachment) GetTableName() (*string, error) {
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

func (rE *attachment) setTableName(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(tableNameKey, val)
}

func (rE *attachment) GetImageWidth() (*float64, error) {
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

func (rE *attachment) setImageWidth(val *float64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(imageWidthKey, val)
}

func (rE *attachment) GetSysModCount() (*int64, error) {
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

func (rE *attachment) setSysModCount(val *int64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sysModCountKey, val)
}

func (rE *attachment) GetContentType() (*string, error) {
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

func (rE *attachment) setContentType(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(contentTypeKey, val)
}

func (rE *attachment) GetSizeCompressed() (*int64, error) {
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

func (rE *attachment) setSizeCompressed(val *int64) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(sizeCompressedKey, val)
}

func (rE *attachment) GetChunkSizeBytes() (*string, error) {
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
func (rE *attachment) setChunkSizeBytes(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(chunkSizeBytesKey, val)
}

func (rE *attachment) GetHash() (*string, error) {
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

func (rE *attachment) setHash(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(hashKey, val)
}

func (rE *attachment) GetState() (*string, error) {
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

func (rE *attachment) setState(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}
	return rE.GetBackingStore().Set(hashKey, val)
}
