package attachmentapi

import (
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
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
	core.BackedModel
}

// NewAttachment creates a new instance of AttachmentModel
func NewAttachment() *Attachment {
	return newAttachment(core.NewBaseModel())
}

// newAttachment creates a new instance of AttachmentModel with the provided model underlying it
func newAttachment(model core.BackedModel) *Attachment {
	return &Attachment{
		model,
	}
}

// CreateAttachmentFromDiscriminatorValue is a parsable factory for creating an AttachmentModel
func CreateAttachmentFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAttachment(), nil
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
		internalSerialization.SerializeStringToInt64Func(chunkSizeBytesKey)(rE.GetChunkSizeBytes),
		internalSerialization.SerializeStringFunc(hashKey)(rE.GetHash),
		internalSerialization.SerializeStringFunc(stateKey)(rE.GetState),
	)
}

// AttachmentModel returns the deserialization information for this object.
func (rE *Attachment) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
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
		sysTagsKey:           internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToPrimitiveSlice(" ", func(s string) (string, error) { return s, nil }))(rE.setSysTags),
		tableNameKey:         internalSerialization.DeserializeStringFunc()(rE.setTableName),
		imageWidthKey:        internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToFloat64Ptr)(rE.setImageWidth),
		sysModCountKey:       internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(rE.setSysModCount),
		contentTypeKey:       internalSerialization.DeserializeStringFunc()(rE.setContentType),
		sizeCompressedKey:    internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(rE.setSizeCompressed),
		chunkSizeBytesKey:    internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToInt64Ptr)(rE.setChunkSizeBytes),
		hashKey:              internalSerialization.DeserializeStringFunc()(rE.setHash),
		stateKey:             internalSerialization.DeserializeStringFunc()(rE.setState),
	}
}

// GetTableSysID returns the table sys id
func (rE *Attachment) GetTableSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, tableSysIDKey)
}

// setTableSysID sets the table sys id to the provide value
func (rE *Attachment) setTableSysID(tableSysID *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, tableSysIDKey, tableSysID)
}

// GetSizeBytes returns the attachment's size in bytes
func (rE *Attachment) GetSizeBytes() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *int64](rE, sizeBytesKey)
}

// setSizeBytes sets the size (in bytes) to the provided value
func (rE *Attachment) setSizeBytes(size *int64) error {
	return store.DefaultBackedModelMutatorFunc(rE, sizeBytesKey, size)
}

// GetDownloadLink returns the download link
func (rE *Attachment) GetDownloadLink() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, downloadLinkKey)
}

// setDownloadLink sets the download link to the provided value
func (rE *Attachment) setDownloadLink(link *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, downloadLinkKey, link)
}

// GetSysUpdatedOn return the last updated timestamp
func (rE *Attachment) GetSysUpdatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *time.Time](rE, sysUpdatedOnKey)
}

// setSysUpdatedOn sets the last updated timestamp to the provided value
func (rE *Attachment) setSysUpdatedOn(val *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(rE, sysUpdatedOnKey, val)
}

// GetSysID returns the sys id
func (rE *Attachment) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, sysIDKey)
}

// setSysID sets the sys id to the provide value
func (rE *Attachment) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, sysIDKey, val)
}

// GetImageHeight returns the image's height, if the attachment is an image
func (rE *Attachment) GetImageHeight() (*float64, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *float64](rE, imageHeightKey)
}

// setImageHeight sets the image height to the provided value
func (rE *Attachment) setImageHeight(val *float64) error {
	return store.DefaultBackedModelMutatorFunc(rE, imageHeightKey, val)
}

// GetSysCreatedOn returns the created on timestamp
func (rE *Attachment) GetSysCreatedOn() (*time.Time, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *time.Time](rE, sysCreatedOnKey)
}

// setSysCreatedOn sets the created on timestamp to the provided value
func (rE *Attachment) setSysCreatedOn(timestamp *time.Time) error {
	return store.DefaultBackedModelMutatorFunc(rE, sysCreatedOnKey, timestamp)
}

// GetFileName returns the file name
func (rE *Attachment) GetFileName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, fileNameKey)
}

// setFileName sets the file name to the provide value
func (rE *Attachment) setFileName(name *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, fileNameKey, name)
}

// GetSysCreatedBy returns the username of who created it
func (rE *Attachment) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, sysCreatedByKey)
}

// setSysCreatedBy sets the username of who created it to the provided value
func (rE *Attachment) setSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, sysCreatedByKey, val)
}

// GetCompressed returns if the attachment is compressed
func (rE *Attachment) GetCompressed() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *bool](rE, compressedKey)
}

// setCompressed sets if the attachment is compressed
func (rE *Attachment) setCompressed(compressed *bool) error {
	return store.DefaultBackedModelMutatorFunc(rE, compressedKey, compressed)
}

// GetAverageImageColor returns the average image color, if an image
func (rE *Attachment) GetAverageImageColor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, averageImageColorKey)
}

// setAverageImageColor sets the average image color to the provided value
func (rE *Attachment) setAverageImageColor(color *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, averageImageColorKey, color)
}

// GetSysUpdatedBy returns the username of the account that last updated the attachment
func (rE *Attachment) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, sysUpdatedByKey)
}

// setSysUpdatedBy sets the username of the account that last updated the attachment to the provide value
func (rE *Attachment) setSysUpdatedBy(username *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, sysUpdatedByKey, username)
}

// GetSysTags returns slice of tags
func (rE *Attachment) GetSysTags() ([]string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, []string](rE, sysTagsKey)
}

// setSysTags sets the sys tags to the provided values
func (rE *Attachment) setSysTags(tags []string) error {
	return store.DefaultBackedModelMutatorFunc(rE, sysTagsKey, tags)
}

// GetTableName returns associated table name
func (rE *Attachment) GetTableName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, tableNameKey)
}

// setTableName sets table name to provided value
func (rE *Attachment) setTableName(name *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, tableNameKey, name)
}

// GetImageWidth returns the width, if attachment is an image
func (rE *Attachment) GetImageWidth() (*float64, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *float64](rE, imageWidthKey)
}

// setImageWidth sets the width to the provide value
func (rE *Attachment) setImageWidth(width *float64) error {
	return store.DefaultBackedModelMutatorFunc(rE, imageWidthKey, width)
}

// GetSysModCount returns the mod count
func (rE *Attachment) GetSysModCount() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *int64](rE, sysModCountKey)
}

// setSysModCount sets the count to the provided value
func (rE *Attachment) setSysModCount(count *int64) error {
	return store.DefaultBackedModelMutatorFunc(rE, sysModCountKey, count)
}

// GetContentType returns the content type of the attachment
func (rE *Attachment) GetContentType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, contentTypeKey)
}

// setContentType sets the content type to the provided value
func (rE *Attachment) setContentType(contentType *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, contentTypeKey, contentType)
}

// GetSizeCompressed returns compressed size of attachment
func (rE *Attachment) GetSizeCompressed() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *int64](rE, sizeCompressedKey)
}

// setSizeCompressed sets compressed size to provided value
func (rE *Attachment) setSizeCompressed(size *int64) error {
	return store.DefaultBackedModelMutatorFunc(rE, sizeCompressedKey, size)
}

// GetChunkSizeBytes returns chunk size (in bytes) of attachment
func (rE *Attachment) GetChunkSizeBytes() (*int64, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *int64](rE, chunkSizeBytesKey)
}

// setChunkSizeBytes sets chunk size to provided value
func (rE *Attachment) setChunkSizeBytes(size *int64) error {
	return store.DefaultBackedModelMutatorFunc(rE, chunkSizeBytesKey, size)
}

// GetHash returns hash of attachment
func (rE *Attachment) GetHash() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, hashKey)
}

// setHash sets hash to provided value
func (rE *Attachment) setHash(hash *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, hashKey, hash)
}

// GetState returns the state
func (rE *Attachment) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Attachment, *string](rE, stateKey)
}

// setState sets the state
func (rE *Attachment) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(rE, stateKey, val)
}
