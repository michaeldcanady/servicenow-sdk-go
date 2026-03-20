package attachmentapi

import (
	"errors"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"

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
	if utils.IsNil(rE) {
		return nil
	}

	return errors.New("serialization not supported")
}

// Attachment2Model returns the deserialization information for this object.
func (rE *Attachment2Model) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { //nolint:gocognit
	return map[string]func(serialization.ParseNode) error{
		tableSysIDKey:        kiota.DeserializeStringFunc(rE.setTableSysID),
		sizeBytesKey:         kiota.DeserializeMutatedStringFunc(utils.StringPtrToInt64Ptr)(rE.setSizeBytes),
		downloadLinkKey:      kiota.DeserializeStringFunc(rE.setDownloadLink),
		sysUpdatedOnKey:      kiota.DeserializeMutatedStringFunc(utils.StringPtrToTimePtr(dateTimeFormat))(rE.setSysUpdatedOn),
		sysIDKey:             kiota.DeserializeStringFunc(rE.setSysID),
		imageHeightKey:       kiota.DeserializeMutatedStringFunc(utils.StringPtrToFloat64Ptr)(rE.setImageHeight),
		sysCreatedOnKey:      kiota.DeserializeMutatedStringFunc(utils.StringPtrToTimePtr(dateTimeFormat))(rE.setSysCreatedOn),
		fileNameKey:          kiota.DeserializeStringFunc(rE.setFileName),
		sysCreatedByKey:      kiota.DeserializeStringFunc(rE.setSysCreatedBy),
		compressedKey:        kiota.DeserializeMutatedStringFunc(utils.StringPtrToBoolPtr)(rE.setCompressed),
		averageImageColorKey: kiota.DeserializeStringFunc(rE.setAverageImageColor),
		sysUpdatedByKey:      kiota.DeserializeStringFunc(rE.setSysUpdatedBy),
		// TODO: figure out separator
		sysTagsKey:        internalSerialization.DeserializeMutatedStringFunc(utils.StringPtrToPrimitiveSlice(" ", func(s string) (string, error) { return s, nil }))(rE.setSysTags),
		tableNameKey:      internalSerialization.DeserializeStringFunc(rE.setTableName),
		imageWidthKey:     internalSerialization.DeserializeMutatedStringFunc(utils.StringPtrToFloat64Ptr)(rE.setImageWidth),
		sysModCountKey:    internalSerialization.DeserializeMutatedStringFunc(utils.StringPtrToInt64Ptr)(rE.setSysModCount),
		contentTypeKey:    internalSerialization.DeserializeStringFunc(rE.setContentType),
		sizeCompressedKey: internalSerialization.DeserializeMutatedStringFunc(utils.StringPtrToInt64Ptr)(rE.setSizeCompressed),
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
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), tableSysIDKey)
}

// setTableSysID sets the table sys id to the provide value
func (rE *Attachment2Model) setTableSysID(tableSysID *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), tableSysIDKey, tableSysID)
}

// GetSizeBytes returns the attachment's size in bytes
func (rE *Attachment2Model) GetSizeBytes() (*int64, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *int64](rE.GetBackingStore(), sizeBytesKey)
}

// setSizeBytes sets the size (in bytes) to the provided value
func (rE *Attachment2Model) setSizeBytes(size *int64) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sizeBytesKey, size)
}

// GetDownloadLink returns the download link
func (rE *Attachment2Model) GetDownloadLink() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), downloadLinkKey)
}

// setDownloadLink sets the download link to the provided value
func (rE *Attachment2Model) setDownloadLink(link *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), downloadLinkKey, link)
}

// GetSysUpdatedOn return the last updated timestamp
func (rE *Attachment2Model) GetSysUpdatedOn() (*time.Time, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *time.Time](rE.GetBackingStore(), sysUpdatedOnKey)
}

// setSysUpdatedOn sets the last updated timestamp to the provided value
func (rE *Attachment2Model) setSysUpdatedOn(val *time.Time) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sysUpdatedOnKey, val)
}

// GetSysID returns the sys id
func (rE *Attachment2Model) GetSysID() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), sysIDKey)
}

// setSysID sets the sys id to the provide value
func (rE *Attachment2Model) setSysID(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sysIDKey, val)
}

// GetImageHeight returns the image's height, if the attachment is an image
func (rE *Attachment2Model) GetImageHeight() (*float64, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *float64](rE.GetBackingStore(), imageHeightKey)
}

// setImageHeight sets the image height to the provided value
func (rE *Attachment2Model) setImageHeight(val *float64) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), imageHeightKey, val)
}

// GetSysCreatedOn returns the created on timestamp
func (rE *Attachment2Model) GetSysCreatedOn() (*time.Time, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *time.Time](rE.GetBackingStore(), sysCreatedOnKey)
}

// setSysCreatedOn sets the created on timestamp to the provided value
func (rE *Attachment2Model) setSysCreatedOn(timestamp *time.Time) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sysCreatedOnKey, timestamp)
}

// GetFileName returns the file name
func (rE *Attachment2Model) GetFileName() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), fileNameKey)
}

// setFileName sets the file name to the provide value
func (rE *Attachment2Model) setFileName(name *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), fileNameKey, name)
}

// GetSysCreatedBy returns the username of who created it
func (rE *Attachment2Model) GetSysCreatedBy() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), sysCreatedByKey)
}

// setSysCreatedBy sets the username of who created it to the provided value
func (rE *Attachment2Model) setSysCreatedBy(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sysCreatedByKey, val)
}

// GetCompressed returns if the attachment is compressed
func (rE *Attachment2Model) GetCompressed() (*bool, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *bool](rE.GetBackingStore(), compressedKey)
}

// setCompressed sets if the attachment is compressed
func (rE *Attachment2Model) setCompressed(compressed *bool) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), compressedKey, compressed)
}

// GetAverageImageColor returns the average image color, if an image
func (rE *Attachment2Model) GetAverageImageColor() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), averageImageColorKey)
}

// setAverageImageColor sets the average image color to the provided value
func (rE *Attachment2Model) setAverageImageColor(color *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), averageImageColorKey, color)
}

// GetSysUpdatedBy returns the username of the account that last updated the attachment
func (rE *Attachment2Model) GetSysUpdatedBy() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), sysUpdatedByKey)
}

// setSysUpdatedBy sets the username of the account that last updated the attachment to the provide value
func (rE *Attachment2Model) setSysUpdatedBy(username *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sysUpdatedByKey, username)
}

// GetSysTags returns slice of tags
func (rE *Attachment2Model) GetSysTags() ([]string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, []string](rE.GetBackingStore(), sysTagsKey)
}

// setSysTags sets the sys tags to the provided values
func (rE *Attachment2Model) setSysTags(tags []string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sysTagsKey, tags)
}

// GetTableName returns associated table name
func (rE *Attachment2Model) GetTableName() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), tableNameKey)
}

// setTableName sets table name to provided value
func (rE *Attachment2Model) setTableName(name *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), tableNameKey, name)
}

// GetImageWidth returns the width, if attachment is an image
func (rE *Attachment2Model) GetImageWidth() (*float64, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *float64](rE.GetBackingStore(), imageWidthKey)
}

// setImageWidth sets the width to the provide value
func (rE *Attachment2Model) setImageWidth(width *float64) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), imageWidthKey, width)
}

// GetSysModCount returns the mod count
func (rE *Attachment2Model) GetSysModCount() (*int64, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *int64](rE.GetBackingStore(), sysModCountKey)
}

// setSysModCount sets the count to the provided value
func (rE *Attachment2Model) setSysModCount(count *int64) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sysModCountKey, count)
}

// GetContentType returns the content type of the attachment
func (rE *Attachment2Model) GetContentType() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), contentTypeKey)
}

// setContentType sets the content type to the provided value
func (rE *Attachment2Model) setContentType(contentType *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), contentTypeKey, contentType)
}

// GetSizeCompressed returns compressed size of attachment
func (rE *Attachment2Model) GetSizeCompressed() (*int64, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *int64](rE.GetBackingStore(), sizeCompressedKey)
}

// setSizeCompressed sets compressed size to provided value
func (rE *Attachment2Model) setSizeCompressed(size *int64) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), sizeCompressedKey, size)
}

// TODO: should be int64
// GetChunkSizeBytes returns chunk size (in bytes) of attachment
func (rE *Attachment2Model) GetChunkSizeBytes() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), chunkSizeBytesKey)
}

// setChunkSizeBytes sets chunk size to provided value
func (rE *Attachment2Model) setChunkSizeBytes(size *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), chunkSizeBytesKey, size)
}

// GetHash returns hash of attachment
func (rE *Attachment2Model) GetHash() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), hashKey)
}

// setHash sets hash to provided value
func (rE *Attachment2Model) setHash(hash *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), hashKey, hash)
}

// GetState returns the state
func (rE *Attachment2Model) GetState() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[store.BackingStore, *string](rE.GetBackingStore(), stateKey)
}

// setState sets the state
func (rE *Attachment2Model) setState(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(rE.GetBackingStore(), stateKey, val)
}
