package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/mock"
)

type MockParseNodeFactory struct {
	mock.Mock
}

func NewMockParseNodeFactory() *MockParseNodeFactory {
	return &MockParseNodeFactory{
		Mock: mock.Mock{},
	}
}

// GetValidContentType returns the content type this factory's parse nodes can deserialize.
func (mPF *MockParseNodeFactory) GetValidContentType() (string, error) {
	args := mPF.Called()
	return args.String(0), args.Error(1)
}

// GetRootParseNode return a new ParseNode instance that is the root of the content
func (mPF *MockParseNodeFactory) GetRootParseNode(contentType string, content []byte) (serialization.ParseNode, error) {
	args := mPF.Called(contentType, content)
	return args.Get(0).(serialization.ParseNode), args.Error(1)
}
