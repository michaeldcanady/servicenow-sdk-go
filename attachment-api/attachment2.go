package attachmentapi

import (
	"errors"
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
	dateTimeFormat       = "2006-01-02 15:04:05"
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
	kiotaStore.BackedModel
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
	return map[string]func(serialization.ParseNode) error{
		tableSysIDKey:        internalSerialization.DeserializeStringFunc(rE.setTableSysID),
		sizeBytesKey:         internalSerialization.DeserializeMutatedStringFunc(rE.setSizeBytes, conversion.StringPtrToInt64Ptr),
		downloadLinkKey:      internalSerialization.DeserializeStringFunc(rE.setDownloadLink),
		sysUpdatedOnKey:      internalSerialization.DeserializeMutatedStringFunc(rE.setSysUpdatedOn, conversion.StringPtrToTimePtr(dateTimeFormat)),
		sysIDKey:             internalSerialization.DeserializeStringFunc(rE.setSysID),
		imageHeightKey:       internalSerialization.DeserializeMutatedStringFunc(rE.setImageHeight, conversion.StringPtrToFloat64Ptr),
		sysCreatedOnKey:      internalSerialization.DeserializeMutatedStringFunc(rE.setSysCreatedOn, conversion.StringPtrToTimePtr(dateTimeFormat)),
		fileNameKey:          internalSerialization.DeserializeStringFunc(rE.setFileName),
		sysCreatedByKey:      internalSerialization.DeserializeStringFunc(rE.setSysCreatedBy),
		compressedKey:        internalSerialization.DeserializeMutatedStringFunc(rE.setCompressed, conversion.StringPtrToBoolPtr),
		averageImageColorKey: internalSerialization.DeserializeStringFunc(rE.setAverageImageColor),
		sysUpdatedByKey:      internalSerialization.DeserializeStringFunc(rE.setSysUpdatedBy),
		// TODO: figure out separator
		sysTagsKey:        internalSerialization.DeserializeMutatedStringFunc(rE.setSysTags, conversion.StringPtrToPrimitiveSlice(" ", func(s string) (string, error) { return s, nil })),
		tableNameKey:      internalSerialization.DeserializeStringFunc(rE.setTableName),
		imageWidthKey:     internalSerialization.DeserializeMutatedStringFunc(rE.setImageWidth, conversion.StringPtrToFloat64Ptr),
		sysModCountKey:    internalSerialization.DeserializeMutatedStringFunc(rE.setSysModCount, conversion.StringPtrToInt64Ptr),
		contentTypeKey:    internalSerialization.DeserializeStringFunc(rE.setContentType),
		sizeCompressedKey: internalSerialization.DeserializeMutatedStringFunc(rE.setSizeCompressed, conversion.StringPtrToInt64Ptr),
		chunkSizeBytesKey: func(node serialization.ParseNode) error {
			val, err := node.GetStringValue()
			if err != nil {
				return err
			}
			return rE.setChunkSizeBytes(val)
		},
		hashKey:  internalSerialization.DeserializeStringFunc(rE.setHash),
		stateKey: internalSerialization.DeserializeStringFunc(rE.setState),
	}
}

// GetTableSysID returns the table sys id
func (rE *Attachment2Model) GetTableSysID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, tableSysIDKey)
}

// setTableSysID sets the table sys id to the provide value
func (rE *Attachment2Model) setTableSysID(tableSysID *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, tableSysIDKey, tableSysID)
}

// GetSizeBytes returns the attachment's size in bytes
func (rE *Attachment2Model) GetSizeBytes() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sizeBytesKey)
}

// setSizeBytes sets the size (in bytes) to the provided value
func (rE *Attachment2Model) setSizeBytes(size *int64) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sizeBytesKey, size)
}

// GetDownloadLink returns the download link
func (rE *Attachment2Model) GetDownloadLink() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, downloadLinkKey)
}

// setDownloadLink sets the download link to the provided value
func (rE *Attachment2Model) setDownloadLink(link *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, downloadLinkKey, link)
}

// GetSysUpdatedOn return the last updated timestamp
func (rE *Attachment2Model) GetSysUpdatedOn() (*time.Time, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sysUpdatedOnKey)
}

// setSysUpdatedOn sets the last updated timestamp to the provided value
func (rE *Attachment2Model) setSysUpdatedOn(val *time.Time) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysUpdatedOnKey, val)
}

// GetSysID returns the sys id
func (rE *Attachment2Model) GetSysID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysIDKey)
}

// setSysID sets the sys id to the provide value
func (rE *Attachment2Model) setSysID(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysIDKey, val)
}

// GetImageHeight returns the image's height, if the attachment is an image
func (rE *Attachment2Model) GetImageHeight() (*float64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, imageHeightKey)
}

// setImageHeight sets the image height to the provided value
func (rE *Attachment2Model) setImageHeight(val *float64) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, imageHeightKey, val)
}

// GetSysCreatedOn returns the created on timestamp
func (rE *Attachment2Model) GetSysCreatedOn() (*time.Time, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sysCreatedOnKey)
}

// setSysCreatedOn sets the created on timestamp to the provided value
func (rE *Attachment2Model) setSysCreatedOn(timestamp *time.Time) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysCreatedOnKey, timestamp)
}

// GetFileName returns the file name
func (rE *Attachment2Model) GetFileName() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, fileNameKey)
}

// setFileName sets the file name to the provide value
func (rE *Attachment2Model) setFileName(name *string) error {
	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, fileNameKey, name)
}

// GetSysCreatedBy returns the username of who created it
func (rE *Attachment2Model) GetSysCreatedBy() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysCreatedByKey)
}

// setSysCreatedBy sets the username of who created it to the provided value
func (rE *Attachment2Model) setSysCreatedBy(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysCreatedByKey, val)
}

// GetCompressed returns if the attachment is compressed
func (rE *Attachment2Model) GetCompressed() (*bool, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, compressedKey)
}

// setCompressed sets if the attachment is compressed
func (rE *Attachment2Model) setCompressed(compressed *bool) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, compressedKey, compressed)
}

// GetAverageImageColor returns the average image color, if an image
func (rE *Attachment2Model) GetAverageImageColor() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, averageImageColorKey)
}

// setAverageImageColor sets the average image color to the provided value
func (rE *Attachment2Model) setAverageImageColor(color *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, averageImageColorKey, color)
}

// GetSysUpdatedBy returns the username of the account that last updated the attachment
func (rE *Attachment2Model) GetSysUpdatedBy() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysUpdatedByKey)
}

// setSysUpdatedBy sets the username of the account that last updated the attachment to the provide value
func (rE *Attachment2Model) setSysUpdatedBy(username *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysUpdatedByKey, username)
}

// GetSysTags returns slice of tags
func (rE *Attachment2Model) GetSysTags() ([]string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []string](backingStore, sysTagsKey)
}

// setSysTags sets the sys tags to the provided values
func (rE *Attachment2Model) setSysTags(tags []string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysTagsKey, tags)
}

// GetTableName returns associated table name
func (rE *Attachment2Model) GetTableName() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, tableNameKey)
}

// setTableName sets table name to provided value
func (rE *Attachment2Model) setTableName(name *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, tableNameKey, name)
}

// GetImageWidth returns the width, if attachment is an image
func (rE *Attachment2Model) GetImageWidth() (*float64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, imageWidthKey)
}

// setImageWidth sets the width to the provide value
func (rE *Attachment2Model) setImageWidth(width *float64) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, imageWidthKey, width)
}

// GetSysModCount returns the mod count
func (rE *Attachment2Model) GetSysModCount() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sysModCountKey)
}

// setSysModCount sets the count to the provided value
func (rE *Attachment2Model) setSysModCount(count *int64) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysModCountKey, count)
}

// GetContentType returns the content type of the attachment
func (rE *Attachment2Model) GetContentType() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, contentTypeKey)
}

// setContentType sets the content type to the provided value
func (rE *Attachment2Model) setContentType(contentType *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, contentTypeKey, contentType)
}

// GetSizeCompressed returns compressed size of attachment
func (rE *Attachment2Model) GetSizeCompressed() (*int64, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sizeCompressedKey)
}

// setSizeCompressed sets compressed size to provided value
func (rE *Attachment2Model) setSizeCompressed(size *int64) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sizeCompressedKey, size)
}

// TODO: should be int64
// GetChunkSizeBytes returns chunk size (in bytes) of attachment
func (rE *Attachment2Model) GetChunkSizeBytes() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, chunkSizeBytesKey)
}

// setChunkSizeBytes sets chunk size to provided value
func (rE *Attachment2Model) setChunkSizeBytes(size *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, chunkSizeBytesKey, size)
}

// GetHash returns hash of attachment
func (rE *Attachment2Model) GetHash() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, hashKey)
}

// setHash sets hash to provided value
func (rE *Attachment2Model) setHash(hash *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, hashKey, hash)
}

// GetState returns the state
func (rE *Attachment2Model) GetState() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, stateKey)
}

// setState sets the state
func (rE *Attachment2Model) setState(val *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, stateKey, val)
}
