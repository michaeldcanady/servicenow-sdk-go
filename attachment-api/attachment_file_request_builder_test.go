package attachmentapi

import "testing"

// TODO: (TestNewAttachmentFileRequestBuilderInternal) Add tests
func TestNewAttachmentFileRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (NewAttachmentFileRequestBuilder) Add tests
func TestNewAttachmentFileRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (AttachmentFileRequestBuilder_Post) Add tests
func TestAttachmentFileRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (AttachmentFileRequestBuilder_ToPostRequestInformation) Add tests
func TestAttachmentFileRequestBuilder_ToPostRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
