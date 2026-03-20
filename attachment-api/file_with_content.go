package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
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
	if utils.IsNil(f) {
		return nil
	}

	return f.File.GetFieldDeserializers()
}

// GetContent returns contents of file
func (f *FileWithContentModel) GetContent() ([]byte, error) {
	if utils.IsNil(f) {
		return nil, nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []byte](backingStore, contentKey)
}

// SetContent sets the content to the provided value
func (f *FileWithContentModel) SetContent(content []byte) error {
	if utils.IsNil(f) {
		return nil
	}

	backingStore := f.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, contentKey, content)
}
