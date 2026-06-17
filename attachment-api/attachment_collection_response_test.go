package attachmentapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestCreateAttachmentCollectionResponseFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "SuccessfulCreation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockParseNode := mocking.NewMockParseNode()

			model, err := CreateAttachmentCollectionResponseFromDiscriminatorValue(mockParseNode)

			assert.NoError(t, err)
			assert.NotNil(t, model)
			assert.IsType(t, &internal.BaseServiceNowCollectionResponse[*Attachment]{}, model)
		})
	}
}
