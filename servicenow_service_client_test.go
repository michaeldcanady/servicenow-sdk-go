package servicenowsdkgo

import (
	"testing"
)

// TODO: add tests
// unsure how to test :(
func TestNewServiceNowServiceClientWithOptions(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
// unsure how to test :(
func TestNewServiceNowServiceClient(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
