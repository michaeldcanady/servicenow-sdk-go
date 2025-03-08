package internal

import "github.com/stretchr/testify/mock"

type mockServiceNowClientOption struct {
	mock.Mock
}

func newMockServiceNowClientOption() *mockServiceNowClientOption {
	return &mockServiceNowClientOption{
		mock.Mock{},
	}
}

func (opt *mockServiceNowClientOption) ServiceNowClientOption(config *serviceNowClientConfig) error {
	args := opt.Called(config)
	return args.Error(0)
}

type mockServiceNowRequestAdapterOption struct {
	mock.Mock
}

func newMockServiceNowRequestAdapterOption() *mockServiceNowRequestAdapterOption {
	return &mockServiceNowRequestAdapterOption{
		mock.Mock{},
	}
}

func (opt *mockServiceNowRequestAdapterOption) ServiceNowClientOption(config *serviceNowRequestAdapterConfig) error {
	args := opt.Called(config)
	return args.Error(0)
}
