package servicenowsdkgo

import "github.com/stretchr/testify/mock"

type mockServiceNowClientOption struct {
	mock.Mock
}

func newMockServiceNowClientOption() *mockServiceNowClientOption {
	return &mockServiceNowClientOption{
		mock.Mock{},
	}
}

func (sCO *mockServiceNowClientOption) ServiceNowServiceClientOption(config *serviceNowServiceClientConfig) error {
	args := sCO.Called(config)

	return args.Error(0)
}
