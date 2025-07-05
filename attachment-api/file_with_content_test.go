package attachmentapi

import (
	"errors"
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

func TestFileWithContent_GetContent(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockContent := []byte{}

				mockBackingStore := mocking.NewMockBackingStore()
				mockBackingStore.On("Get", "content").Return(mockContent, nil)

				mockModel := mocking.NewMockModel()
				mockModel.On("GetBackingStore").Return(mockBackingStore)

				file := &FileWithContentModel{File: &FileModel{Model: mockModel}}

				content, err := file.GetContent()

				assert.Equal(t, mockContent, content)
				assert.Nil(t, err)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				mockContent := "bad type"

				mockBackingStore := mocking.NewMockBackingStore()
				mockBackingStore.On("Get", "content").Return(mockContent, nil)

				mockModel := mocking.NewMockModel()
				mockModel.On("GetBackingStore").Return(mockBackingStore)

				file := &FileWithContentModel{File: &FileModel{Model: mockModel}}

				content, err := file.GetContent()

				assert.Nil(t, content)
				assert.Equal(t, errors.New("content is not []byte"), err)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				mockBackingStore := mocking.NewMockBackingStore()
				mockBackingStore.On("Get", "content").Return(nil, errors.New("retrieval error"))

				mockModel := mocking.NewMockModel()
				mockModel.On("GetBackingStore").Return(mockBackingStore)

				file := &FileWithContentModel{File: &FileModel{Model: mockModel}}

				content, err := file.GetContent()

				assert.Nil(t, content)
				assert.Equal(t, errors.New("retrieval error"), err)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				mockContent := []byte{}

				mockBackingStore := mocking.NewMockBackingStore()
				mockBackingStore.On("Get", "content").Return(mockContent, nil)

				mockModel := mocking.NewMockModel()
				mockModel.On("GetBackingStore").Return(nil)

				file := &FileWithContentModel{File: &FileModel{Model: mockModel}}

				content, err := file.GetContent()

				assert.Nil(t, content)
				assert.Equal(t, errors.New("store is nil"), err)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				file := (*FileWithContentModel)(nil)

				content, err := file.GetContent()

				assert.Nil(t, content)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestFileWithContent_SetContent(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockContent := []byte{}

				mockBackingStore := mocking.NewMockBackingStore()
				mockBackingStore.On("Set", "content", mockContent).Return(nil)

				mockModel := mocking.NewMockModel()
				mockModel.On("GetBackingStore").Return(mockBackingStore)

				file := &FileWithContentModel{File: &FileModel{Model: mockModel}}

				err := file.SetContent(mockContent)

				assert.Nil(t, err)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				mockContent := []byte{}

				mockModel := mocking.NewMockModel()
				mockModel.On("GetBackingStore").Return(nil)

				file := &FileWithContentModel{File: &FileModel{Model: mockModel}}

				err := file.SetContent(mockContent)

				assert.Equal(t, errors.New("store is nil"), err)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				mockContent := []byte{}

				file := (*FileWithContentModel)(nil)

				err := file.SetContent(mockContent)

				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
