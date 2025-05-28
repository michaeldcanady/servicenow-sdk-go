package attachmentapi

import "testing"

// TODO: (TestNewAttachmentItemFileRequestBuilderInternal) Add tests
func TestNewAttachmentItemFileRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestNewAttachmentItemFileRequestBuilder) Add tests
func TestNewAttachmentItemFileRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestAttachmentItemFileRequestBuilder_Get) Add tests
func TestAttachmentItemFileRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestAttachmentItemFileRequestBuilder_ToGetRequestInformation) Add tests
func TestAttachmentItemFileRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
