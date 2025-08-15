//go:build preview

package ast

import "github.com/stretchr/testify/mock"

type mockStringerWriter struct {
	mock.Mock
}

func newMockStringerWriter() *mockStringerWriter {
	return &mockStringerWriter{
		mock.Mock{},
	}
}

func (mock *mockStringerWriter) WriteString(s string) (n int, err error) {
	args := mock.Called(s)
	return args.Int(0), args.Error(1)
}

func (mock *mockStringerWriter) String() string {
	args := mock.Called()
	return args.String(0)
}
