package attachmentapi

import "github.com/microsoft/kiota-abstractions-go/serialization"

var _ serialization.Parsable = (*Media)(nil)

// Media represents a HTTP media file
type Media struct {
	contentType string
	data        []byte
}

// NewMedia creates a new media
func NewMedia(contentType string, data []byte) *Media {
	return &Media{
		contentType: contentType,
		data:        data,
	}
}

// GetContentType returns the content-type of the media
func (P *Media) GetContentType() string {
	return P.contentType
}

// Serialize writes the objects properties to the current writer.
func (P *Media) Serialize(writer serialization.SerializationWriter) error {
	return writer.WriteByteArrayValue("", P.data)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (P *Media) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}
