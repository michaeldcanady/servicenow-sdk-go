package attachmentapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewMedia(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				contentType := "application/json"
				data := []byte("I am data")

				media := NewMedia(contentType, data)

				assert.Equal(t, contentType, media.contentType)
				assert.Equal(t, data, media.data)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

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

func TestMedia_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockData := []byte("I am a string")

				mockWriter := mocking.NewMockSerializationWriter()
				mockWriter.On("WriteByteArrayValue", "", mockData).Return(nil)

				media := &Media{
					data: mockData,
				}

				err := media.Serialize(mockWriter)

				assert.Nil(t, err)
			},
		},
		{
			name: "Error",
			test: func(t *testing.T) {
				mockData := []byte("I am a string")

				mockWriter := mocking.NewMockSerializationWriter()
				mockWriter.On("WriteByteArrayValue", "", mockData).Return(errors.New("errored"))

				media := &Media{
					data: mockData,
				}

				err := media.Serialize(mockWriter)

				assert.Equal(t, errors.New("errored"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestMedia_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				media := &Media{}

				fieldDeserializers := media.GetFieldDeserializers()

				assert.Nil(t, fieldDeserializers)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
