package attachmentapi

import (
	"testing"
)

// TODO: (TestCreateAttachmentCollectionResponse2FromDiscriminatorValue) Add tests
func TestCreateAttachmentCollectionResponse2FromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
