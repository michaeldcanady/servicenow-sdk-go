package attachmentapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	contentKey = "content"
)

// FileWithContentable represents Service-Now File with it's content
type FileWithContentable interface {
	Fileable
	GetContent() ([]byte, error)
	setContent([]byte) error
}

// fileWithContent implementation of FileWithContentable
type fileWithContent struct {
	Fileable
}

// NewFileWithContent creates a new instance of FileWithContentable
func NewFileWithContent() FileWithContentable {
	return &fileWithContent{
		Fileable: NewFile(),
	}
}

// CreateFileWithContentFromDiscriminatorValue is a parsable factory for creating a FileWithContentable
func CreateFileWithContentFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewFileWithContent(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (f *fileWithContent) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(f) {
		f = NewFileWithContent().(*fileWithContent)
	}

	return f.Fileable.GetFieldDeserializers()
}

// GetContent returns contents of file
func (f *fileWithContent) GetContent() ([]byte, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	val, err := f.Fileable.GetBackingStore().Get(contentKey)
	if err != nil {
		return nil, err
	}

	content, ok := val.([]byte)
	if !ok {
		return nil, errors.New("content is not []byte")
	}

	return content, nil
}

// setContent sets the content to the provided value
func (f *fileWithContent) setContent(content []byte) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.Fileable.GetBackingStore().Set(contentKey, content)
}
