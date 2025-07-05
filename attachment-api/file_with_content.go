package attachmentapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	contentKey = "content"
)

// FileWithContent represents Service-Now File with it's content
type FileWithContent interface {
	File
	GetContent() ([]byte, error)
	SetContent([]byte) error
}

// FileWithContentModel implementation of FileWithContent
type FileWithContentModel struct {
	File
}

// NewFileWithContent creates a new instance of FileWithContent
func newFileWithContent(file File) FileWithContent {
	return &FileWithContentModel{
		file,
	}
}

// NewFileWithContent creates a new instance of FileWithContent
func NewFileWithContent() FileWithContent {
	return newFileWithContent(
		NewFile(),
	)
}

// CreateFileWithContentFromDiscriminatorValue is a parsable factory for creating a FileWithContent
func CreateFileWithContentFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewFileWithContent(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (f *FileWithContentModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(f) {
		return nil
	}

	return f.File.GetFieldDeserializers()
}

// GetContent returns contents of file
func (f *FileWithContentModel) GetContent() ([]byte, error) {
	if internal.IsNil(f) {
		return nil, nil
	}

	store := f.GetBackingStore()
	if internal.IsNil(store) {
		return nil, errors.New("store is nil")
	}

	val, err := store.Get(contentKey)
	if err != nil {
		return nil, err
	}

	content, ok := val.([]byte)
	if !ok {
		return nil, errors.New("content is not []byte")
	}

	return content, nil
}

// SetContent sets the content to the provided value
func (f *FileWithContentModel) SetContent(content []byte) error {
	if internal.IsNil(f) {
		return nil
	}

	store := f.GetBackingStore()
	if internal.IsNil(store) {
		return errors.New("store is nil")
	}

	return store.Set(contentKey, content)
}
