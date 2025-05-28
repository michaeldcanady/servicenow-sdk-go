package mocking

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/mock"
)

type MockRequestOption struct {
	mock.Mock
}

func NewMockRequestOption() *MockRequestOption {
	return &MockRequestOption{
		mock.Mock{},
	}
}

// GetKey returns the key to store the current option under.
func (m *MockRequestOption) GetKey() abstractions.RequestOptionKey {
	args := m.Called()
	return args.Get(0).(abstractions.RequestOptionKey)
}
