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

// Attachment
type Attachment struct {
	internal.BackedModel
}

// NewAttachment creates a new instance of Attachment2Model
func NewAttachment2() *Attachment {
	return newAttachment2(internal.NewBaseModel())
}

// newAttachment2 creates a new instance of Attachment2Model with the provided model underlying it
func newAttachment2(model internal.BackedModel) *Attachment {
	return &Attachment{
		model,
	}
}

// CreateAttachment2FromDiscriminatorValue is a parsable factory for creating an Attachment2Model
func CreateAttachment2FromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAttachment2(), nil
}

// Serialize writes the objects properties to the current writer.
func (rE *Attachment) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(rE) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(tableSysIDKey)(rE.GetTableSysID),
		internalSerialization.SerializeStringToInt64Func(sizeBytesKey)(rE.GetSizeBytes),
		internalSerialization.SerializeStringFunc(downloadLinkKey)(rE.GetDownloadLink),
		internalSerialization.SerializeStringToTimeFunc(sysUpdatedOnKey, dateTimeFormat)(rE.GetSysUpdatedOn),
		internalSerialization.SerializeStringFunc(sysIDKey)(rE.GetSysID),
		internalSerialization.SerializeStringToFloat64Func(imageHeightKey)(rE.GetImageHeight),
		internalSerialization.SerializeStringToTimeFunc(sysCreatedOnKey, dateTimeFormat)(rE.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(fileNameKey)(rE.GetFileName),
		internalSerialization.SerializeStringFunc(sysCreatedByKey)(rE.GetSysCreatedBy),
		internalSerialization.SerializeStringToBoolFunc(compressedKey)(rE.GetCompressed),
		internalSerialization.SerializeStringFunc(averageImageColorKey)(rE.GetAverageImageColor),
		internalSerialization.SerializeStringFunc(sysUpdatedByKey)(rE.GetSysUpdatedBy),
		internalSerialization.SerializeStringToSliceFunc(sysTagsKey, " ")(rE.GetSysTags),
		internalSerialization.SerializeStringFunc(tableNameKey)(rE.GetTableName),
		internalSerialization.SerializeStringToFloat64Func(imageWidthKey)(rE.GetImageWidth),
		internalSerialization.SerializeStringToInt64Func(sysModCountKey)(rE.GetSysModCount),
		internalSerialization.SerializeStringFunc(contentTypeKey)(rE.GetContentType),
		internalSerialization.SerializeStringToInt64Func(sizeCompressedKey)(rE.GetSizeCompressed),
		internalSerialization.SerializeStringFunc(chunkSizeBytesKey)(rE.GetChunkSizeBytes),
		internalSerialization.SerializeStringFunc(hashKey)(rE.GetHash),
		internalSerialization.SerializeStringFunc(stateKey)(rE.GetState),
	)
}

// Attachment2Model returns the deserialization information for this object.
func (rE *Attachment) GetFieldDeserializers() map[string]func(serialization.ParseNode) error { //nolint:gocognit
	return map[string]func(serialization.ParseNode) error{
		tableSysIDKey:        internalSerialization.DeserializeStringFunc()(rE.setTableSysID),
		sizeBytesKey:         internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(rE.setSizeBytes),
		downloadLinkKey:      internalSerialization.DeserializeStringFunc()(rE.setDownloadLink),
		sysUpdatedOnKey:      internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr(dateTimeFormat))(rE.setSysUpdatedOn),
		sysIDKey:             internalSerialization.DeserializeStringFunc()(rE.setSysID),
		imageHeightKey:       internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(rE.setImageHeight),
		sysCreatedOnKey:      internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToTimePtr(dateTimeFormat))(rE.setSysCreatedOn),
		fileNameKey:          internalSerialization.DeserializeStringFunc()(rE.setFileName),
		sysCreatedByKey:      internalSerialization.DeserializeStringFunc()(rE.setSysCreatedBy),
		compressedKey:        internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(rE.setCompressed),
		averageImageColorKey: internalSerialization.DeserializeStringFunc()(rE.setAverageImageColor),
		sysUpdatedByKey:      internalSerialization.DeserializeStringFunc()(rE.setSysUpdatedBy),
		// TODO: figure out separator
		sysTagsKey:        internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToPrimitiveSlice(" ", func(s string) (string, error) { return s, nil }))(rE.setSysTags),
		tableNameKey:      internalSerialization.DeserializeStringFunc()(rE.setTableName),
		imageWidthKey:     internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(rE.setImageWidth),
		sysModCountKey:    internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(rE.setSysModCount),
		contentTypeKey:    internalSerialization.DeserializeStringFunc()(rE.setContentType),
		sizeCompressedKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(rE.setSizeCompressed),
		chunkSizeBytesKey: internalSerialization.DeserializeStringFunc()(rE.setChunkSizeBytes),
		hashKey:           internalSerialization.DeserializeStringFunc()(rE.setHash),
		stateKey:          internalSerialization.DeserializeStringFunc()(rE.setState),
	}
}

// GetTableSysID returns the table sys id
func (rE *Attachment) GetTableSysID() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, tableSysIDKey)
}

// setTableSysID sets the table sys id to the provide value
func (rE *Attachment) setTableSysID(tableSysID *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, tableSysIDKey, tableSysID)
}

// GetSizeBytes returns the attachment's size in bytes
func (rE *Attachment) GetSizeBytes() (*int64, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sizeBytesKey)
}

// setSizeBytes sets the size (in bytes) to the provided value
func (rE *Attachment) setSizeBytes(size *int64) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sizeBytesKey, size)
}

// GetDownloadLink returns the download link
func (rE *Attachment) GetDownloadLink() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, downloadLinkKey)
}

// setDownloadLink sets the download link to the provided value
func (rE *Attachment) setDownloadLink(link *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, downloadLinkKey, link)
}

// GetSysUpdatedOn return the last updated timestamp
func (rE *Attachment) GetSysUpdatedOn() (*time.Time, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sysUpdatedOnKey)
}

// setSysUpdatedOn sets the last updated timestamp to the provided value
func (rE *Attachment) setSysUpdatedOn(val *time.Time) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysUpdatedOnKey, val)
}

// GetSysID returns the sys id
func (rE *Attachment) GetSysID() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysIDKey)
}

// setSysID sets the sys id to the provide value
func (rE *Attachment) setSysID(val *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysIDKey, val)
}

// GetImageHeight returns the image's height, if the attachment is an image
func (rE *Attachment) GetImageHeight() (*float64, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, imageHeightKey)
}

// setImageHeight sets the image height to the provided value
func (rE *Attachment) setImageHeight(val *float64) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, imageHeightKey, val)
}

// GetSysCreatedOn returns the created on timestamp
func (rE *Attachment) GetSysCreatedOn() (*time.Time, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *time.Time](backingStore, sysCreatedOnKey)
}

// setSysCreatedOn sets the created on timestamp to the provided value
func (rE *Attachment) setSysCreatedOn(timestamp *time.Time) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysCreatedOnKey, timestamp)
}

// GetFileName returns the file name
func (rE *Attachment) GetFileName() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, fileNameKey)
}

// setFileName sets the file name to the provide value
func (rE *Attachment) setFileName(name *string) error {
	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, fileNameKey, name)
}

// GetSysCreatedBy returns the username of who created it
func (rE *Attachment) GetSysCreatedBy() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysCreatedByKey)
}

// setSysCreatedBy sets the username of who created it to the provided value
func (rE *Attachment) setSysCreatedBy(val *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysCreatedByKey, val)
}

// GetCompressed returns if the attachment is compressed
func (rE *Attachment) GetCompressed() (*bool, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, compressedKey)
}

// setCompressed sets if the attachment is compressed
func (rE *Attachment) setCompressed(compressed *bool) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, compressedKey, compressed)
}

// GetAverageImageColor returns the average image color, if an image
func (rE *Attachment) GetAverageImageColor() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, averageImageColorKey)
}

// setAverageImageColor sets the average image color to the provided value
func (rE *Attachment) setAverageImageColor(color *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, averageImageColorKey, color)
}

// GetSysUpdatedBy returns the username of the account that last updated the attachment
func (rE *Attachment) GetSysUpdatedBy() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sysUpdatedByKey)
}

// setSysUpdatedBy sets the username of the account that last updated the attachment to the provide value
func (rE *Attachment) setSysUpdatedBy(username *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysUpdatedByKey, username)
}

// GetSysTags returns slice of tags
func (rE *Attachment) GetSysTags() ([]string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []string](backingStore, sysTagsKey)
}

// setSysTags sets the sys tags to the provided values
func (rE *Attachment) setSysTags(tags []string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysTagsKey, tags)
}

// GetTableName returns associated table name
func (rE *Attachment) GetTableName() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, tableNameKey)
}

// setTableName sets table name to provided value
func (rE *Attachment) setTableName(name *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, tableNameKey, name)
}

// GetImageWidth returns the width, if attachment is an image
func (rE *Attachment) GetImageWidth() (*float64, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *float64](backingStore, imageWidthKey)
}

// setImageWidth sets the width to the provide value
func (rE *Attachment) setImageWidth(width *float64) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, imageWidthKey, width)
}

// GetSysModCount returns the mod count
func (rE *Attachment) GetSysModCount() (*int64, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sysModCountKey)
}

// setSysModCount sets the count to the provided value
func (rE *Attachment) setSysModCount(count *int64) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sysModCountKey, count)
}

// GetContentType returns the content type of the attachment
func (rE *Attachment) GetContentType() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, contentTypeKey)
}

// setContentType sets the content type to the provided value
func (rE *Attachment) setContentType(contentType *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, contentTypeKey, contentType)
}

// GetSizeCompressed returns compressed size of attachment
func (rE *Attachment) GetSizeCompressed() (*int64, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int64](backingStore, sizeCompressedKey)
}

// setSizeCompressed sets compressed size to provided value
func (rE *Attachment) setSizeCompressed(size *int64) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sizeCompressedKey, size)
}

// TODO: should be int64
// GetChunkSizeBytes returns chunk size (in bytes) of attachment
func (rE *Attachment) GetChunkSizeBytes() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, chunkSizeBytesKey)
}

// setChunkSizeBytes sets chunk size to provided value
func (rE *Attachment) setChunkSizeBytes(size *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, chunkSizeBytesKey, size)
}

// GetHash returns hash of attachment
func (rE *Attachment) GetHash() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, hashKey)
}

// setHash sets hash to provided value
func (rE *Attachment) setHash(hash *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, hashKey, hash)
}

// GetState returns the state
func (rE *Attachment) GetState() (*string, error) {
	if conversion.IsNil(rE) {
		return nil, nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, stateKey)
}

// setState sets the state
func (rE *Attachment) setState(val *string) error {
	if conversion.IsNil(rE) {
		return nil
	}

	backingStore := rE.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, stateKey, val)
}
