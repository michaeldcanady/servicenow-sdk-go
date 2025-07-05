package attachmentapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
)

//func TestNewMedia(t *testing.T) {
//	tests := []struct {
//		name string
//		test func(*testing.T)
//	}{}
//
//	for _, test := range tests {
//		t.Run(test.name, test.test)
//	}
//}

func TestNewFileWithContent(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				file := NewFileWithContent()

				assert.IsType(t, &FileWithContentModel{}, file)
				assert.IsType(t, &FileModel{}, file.(*FileWithContentModel).File)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateFileWithContentFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockParseNode := mocking.NewMockParseNode()

				model, err := CreateFileWithContentFromDiscriminatorValue(mockParseNode)

				assert.Nil(t, err)
				assert.IsType(t, &FileWithContentModel{}, model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestFileWithContent_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				file := mocking.NewMockFile()
				file.On("GetFieldDeserializers").Return(make(map[string]func(serialization.ParseNode) error, 0))

				fileWithContent := &FileWithContentModel{
					File: file,
				}

				fileWithContent.GetFieldDeserializers()

				file.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				fileWithContent := (*FileWithContentModel)(nil)

				fieldDeserializers := fileWithContent.GetFieldDeserializers()

				assert.Nil(t, fieldDeserializers)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: Write tests for TestFileWithContent_GetContent
func TestFileWithContent_GetContent(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: Write tests for TestFileWithContent_SetContent
func TestFileWithContent_SetContent(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
