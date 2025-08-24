//go:build preview.query

package query

import "github.com/stretchr/testify/mock"

type mockErrorAdder struct {
	mock.Mock
}

func newMockErrorAdder() *mockErrorAdder {
	return &mockErrorAdder{
		mock.Mock{},
	}
}

func (mock *mockErrorAdder) addErrors(errs ...error) {
	_ = mock.Called(errs)
}
