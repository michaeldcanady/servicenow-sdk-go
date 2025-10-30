package attachmentapi

import (
	"errors"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	internalStore "github.com/michaeldcanady/servicenow-sdk-go/internal/store"

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
		sysTagsKey:        internalSerialization.DeserializeMutatedStringFunc(rE.setSysTags, conversion.StringPtrToPrimitiveSlice[string](" ", func(s string) (string, error) { return s, nil })),
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
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, tableSysIDKey)
}

// setTableSysID sets the table sys id to the provide value
func (rE *Attachment2Model) setTableSysID(tableSysID *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, tableSysIDKey, tableSysID)
}

// GetSizeBytes returns the attachment's size in bytes
func (rE *Attachment2Model) GetSizeBytes() (*int64, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *int64](rE, sizeBytesKey)
}

// setSizeBytes sets the size (in bytes) to the provided value
func (rE *Attachment2Model) setSizeBytes(size *int64) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sizeBytesKey, size)
}

// GetDownloadLink returns the download link
func (rE *Attachment2Model) GetDownloadLink() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, downloadLinkKey)
}

// setDownloadLink sets the download link to the provided value
func (rE *Attachment2Model) setDownloadLink(link *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, downloadLinkKey, link)
}

// GetSysUpdatedOn return the last updated timestamp
func (rE *Attachment2Model) GetSysUpdatedOn() (*time.Time, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *time.Time](rE, sysUpdatedOnKey)
}

// setSysUpdatedOn sets the last updated timestamp to the provided value
func (rE *Attachment2Model) setSysUpdatedOn(val *time.Time) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sysUpdatedOnKey, val)
}

// GetSysID returns the sys id
func (rE *Attachment2Model) GetSysID() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, sysIDKey)
}

// setSysID sets the sys id to the provide value
func (rE *Attachment2Model) setSysID(val *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sysIDKey, val)
}

// GetImageHeight returns the image's height, if the attachment is an image
func (rE *Attachment2Model) GetImageHeight() (*float64, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *float64](rE, imageHeightKey)
}

// setImageHeight sets the image height to the provided value
func (rE *Attachment2Model) setImageHeight(val *float64) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, imageHeightKey, val)
}

// GetSysCreatedOn returns the created on timestamp
func (rE *Attachment2Model) GetSysCreatedOn() (*time.Time, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *time.Time](rE, sysCreatedOnKey)
}

// setSysCreatedOn sets the created on timestamp to the provided value
func (rE *Attachment2Model) setSysCreatedOn(timestamp *time.Time) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sysCreatedOnKey, timestamp)
}

// GetFileName returns the file name
func (rE *Attachment2Model) GetFileName() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, fileNameKey)
}

// setFileName sets the file name to the provide value
func (rE *Attachment2Model) setFileName(name *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, fileNameKey, name)
}

// GetSysCreatedBy returns the username of who created it
func (rE *Attachment2Model) GetSysCreatedBy() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, sysCreatedByKey)
}

// setSysCreatedBy sets the username of who created it to the provided value
func (rE *Attachment2Model) setSysCreatedBy(val *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sysCreatedByKey, val)
}

// GetCompressed returns if the attachment is compressed
func (rE *Attachment2Model) GetCompressed() (*bool, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *bool](rE, compressedKey)
}

// setCompressed sets if the attachment is compressed
func (rE *Attachment2Model) setCompressed(compressed *bool) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, compressedKey, compressed)
}

// GetAverageImageColor returns the average image color, if an image
func (rE *Attachment2Model) GetAverageImageColor() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, averageImageColorKey)
}

// setAverageImageColor sets the average image color to the provided value
func (rE *Attachment2Model) setAverageImageColor(color *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, averageImageColorKey, color)
}

// GetSysUpdatedBy returns the username of the account that last updated the attachment
func (rE *Attachment2Model) GetSysUpdatedBy() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, sysUpdatedByKey)
}

// setSysUpdatedBy sets the username of the account that last updated the attachment to the provide value
func (rE *Attachment2Model) setSysUpdatedBy(username *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sysUpdatedByKey, username)
}

// GetSysTags returns slice of tags
func (rE *Attachment2Model) GetSysTags() ([]string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, []string](rE, sysTagsKey)
}

// setSysTags sets the sys tags to the provided values
func (rE *Attachment2Model) setSysTags(tags []string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sysTagsKey, tags)
}

// GetTableName returns associated table name
func (rE *Attachment2Model) GetTableName() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, tableNameKey)
}

// setTableName sets table name to provided value
func (rE *Attachment2Model) setTableName(name *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, tableNameKey, name)
}

// GetImageWidth returns the width, if attachment is an image
func (rE *Attachment2Model) GetImageWidth() (*float64, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *float64](rE, imageWidthKey)
}

// setImageWidth sets the width to the provide value
func (rE *Attachment2Model) setImageWidth(width *float64) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, imageWidthKey, width)
}

// GetSysModCount returns the mod count
func (rE *Attachment2Model) GetSysModCount() (*int64, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *int64](rE, sysModCountKey)
}

// setSysModCount sets the count to the provided value
func (rE *Attachment2Model) setSysModCount(count *int64) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sysModCountKey, count)
}

// GetContentType returns the content type of the attachment
func (rE *Attachment2Model) GetContentType() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, contentTypeKey)
}

// setContentType sets the content type to the provided value
func (rE *Attachment2Model) setContentType(contentType *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, contentTypeKey, contentType)
}

// GetSizeCompressed returns compressed size of attachment
func (rE *Attachment2Model) GetSizeCompressed() (*int64, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *int64](rE, sizeCompressedKey)
}

// setSizeCompressed sets compressed size to provided value
func (rE *Attachment2Model) setSizeCompressed(size *int64) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, sizeCompressedKey, size)
}

// TODO: should be int64
// GetChunkSizeBytes returns chunk size (in bytes) of attachment
func (rE *Attachment2Model) GetChunkSizeBytes() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, chunkSizeBytesKey)
}

// setChunkSizeBytes sets chunk size to provided value
func (rE *Attachment2Model) setChunkSizeBytes(size *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, chunkSizeBytesKey, size)
}

// GetHash returns hash of attachment
func (rE *Attachment2Model) GetHash() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, hashKey)
}

// setHash sets hash to provided value
func (rE *Attachment2Model) setHash(hash *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, hashKey, hash)
}

// GetState returns the state
func (rE *Attachment2Model) GetState() (*string, error) {
	return internalStore.DefaultBackedModelAccessorFunc[*Attachment2Model, *string](rE, stateKey)
}

// setState sets the state
func (rE *Attachment2Model) setState(val *string) error {
	return internalStore.DefaultBackedModelMutatorFunc(rE, stateKey, val)
}
