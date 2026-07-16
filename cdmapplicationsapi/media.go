package cdmapplicationsapi

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// Media represents an HTTP media file for upload.
type Media struct {
	contentType string
	data        []byte
}

func NewMedia(contentType string, data []byte) *Media {
	return &Media{
		contentType: contentType,
		data:        data,
	}
}

func (p *Media) GetContentType() string {
	return p.contentType
}

func (p *Media) GetData() []byte {
	return p.data
}

func (p *Media) Serialize(writer serialization.SerializationWriter) error {
	return writer.WriteByteArrayValue("", p.data)
}

func (p *Media) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}
