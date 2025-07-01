package attachmentapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestCreateAttachmentCollectionResponse2FromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockParseNode := mocking.NewMockParseNode()

				model, err := CreateAttachmentCollectionResponse2FromDiscriminatorValue(mockParseNode)

				assert.Nil(t, err)
				assert.IsType(t, &newInternal.BaseServiceNowCollectionResponse[Attachment2]{}, model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
