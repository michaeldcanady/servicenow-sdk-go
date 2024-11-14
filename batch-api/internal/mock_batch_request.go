package internal

import (
	"github.com/RecoLabs/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/mock"
)

type MockBatchRequest struct {
	mock.Mock
}

func (m *MockBatchRequest) AddRequest(req internal.RequestInformation, flag bool) error {
	args := m.Called(req, flag)
	return args.Error(0)
}
