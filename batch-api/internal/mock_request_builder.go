package internal

import (
	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/mock"
)

type MockRequestBuilder struct {
	mock.Mock
}

func (m *MockRequestBuilder) SendPost3(config *core.RequestConfiguration) error {
	args := m.Called(config)
	return args.Error(0)
}
