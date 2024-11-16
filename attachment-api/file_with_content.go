package attachmentapi

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type FileWithContentable interface {
	Fileable
	GetContent() ([]byte, error)
	setContent([]byte) error
}

type fileWithContent struct {
	Fileable
}

func NewFileWithContent() FileWithContentable {
	return &fileWithContent{
		Fileable: NewFile(),
	}
}

func CreateFileWithContentFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewFileWithContent(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (f *fileWithContent) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return f.Fileable.GetFieldDeserializers()
}

func (f *fileWithContent) GetContent() ([]byte, error) {
	val, err := f.Fileable.GetBackingStore().Get("content")
	if err != nil {
		return nil, err
	}

	content, ok := val.([]byte)
	if !ok {
		return nil, errors.New("content is not []byte")
	}

	return content, nil
}

func (f *fileWithContent) setContent(content []byte) error {
	return f.Fileable.GetBackingStore().Set("content", content)
}
