package attachmentapi

import "testing"

// TODO: (TestNewAttachmentUploadRequestBuilderInternal) Add tests
func TestNewAttachmentUploadRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestNewAttachmentUploadRequestBuilder) Add tests
func TestNewAttachmentUploadRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestAttachmentUploadRequestBuilder_Post) Add tests
func TestAttachmentUploadRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestAttachmentUploadRequestBuilder_ToPostRequestInformation) Add tests
func TestAttachmentUploadRequestBuilder_ToPostRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
