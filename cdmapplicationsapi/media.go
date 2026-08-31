// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// Media represents an HTTP media file for upload.
type Media struct {
	contentType string
	data        []byte
}

// NewMedia instantiates a new Media.
func NewMedia(contentType string, data []byte) *Media {
	return &Media{
		contentType: contentType,
		data:        data,
	}
}

// GetContentType returns the content type.
func (p *Media) GetContentType() string {
	return p.contentType
}

// GetData returns the data.
func (p *Media) GetData() []byte {
	return p.data
}

// Serialize writes the object's properties to the given writer.
func (p *Media) Serialize(writer serialization.SerializationWriter) error {
	return writer.WriteByteArrayValue("", p.data)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (p *Media) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}
