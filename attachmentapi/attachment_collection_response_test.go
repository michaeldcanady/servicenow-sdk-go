package attachmentapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
)

func TestCreateAttachmentCollectionResponseFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name         string
		parseNode    serialization.ParseNode
		wantErr      bool
		expectedType interface{}
	}{
		{
			name:         "nil parse node",
			parseNode:    nil,
			wantErr:      false,
			expectedType: &core.BaseServiceNowCollectionResponse[*Attachment]{},
		},
		{
			name:         "valid parse node",
			parseNode:    mocking.NewMockParseNode(),
			wantErr:      false,
			expectedType: &core.BaseServiceNowCollectionResponse[*Attachment]{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAttachmentCollectionResponseFromDiscriminatorValue(tt.parseNode)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.IsType(t, tt.expectedType, got)

				// Verify it has the expected field deserializers
				deserializers := got.GetFieldDeserializers()
				assert.NotNil(t, deserializers)
				assert.Contains(t, deserializers, "result")
				assert.Contains(t, deserializers, "next")
				assert.Contains(t, deserializers, "previous")
				assert.Contains(t, deserializers, "first")
				assert.Contains(t, deserializers, "last")
			}
		})
	}
}
