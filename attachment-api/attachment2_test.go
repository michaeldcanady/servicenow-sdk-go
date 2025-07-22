package attachmentapi

import (
	"testing"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestNewAttachment2(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				value := NewAttachment2()

				assert.IsType(t, &Attachment2Model{}, value)
				assert.IsType(t, &newInternal.BaseModel{}, value.Model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestAttachment2_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*Attachment2Model)(nil)

				assert.Nil(t, model)

				_ = model.GetFieldDeserializers()

				assert.NotNil(t, model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
