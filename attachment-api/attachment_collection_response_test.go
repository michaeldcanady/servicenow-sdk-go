package attachmentapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestCreateAttachmentCollectionResponseFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockParseNode := mocking.NewMockParseNode()

				model, err := CreateAttachmentCollectionResponseFromDiscriminatorValue(mockParseNode)

				assert.Nil(t, err)
				assert.IsType(t, &internal.BaseServiceNowCollectionResponse[*Attachment]{}, model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
