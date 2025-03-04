package mocking

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockContext struct {
	mock.Mock
}

func NewMockContext() *MockContext {
	return &MockContext{
		mock.Mock{},
	}
}

func (c *MockContext) Deadline() (deadline time.Time, ok bool) {
	_ = c.Called(deadline, ok)
	return
}

func (c *MockContext) Done() <-chan struct{} {
	args := c.Called()
	return args.Get(0).(<-chan struct{})
}

func (c *MockContext) Err() error {
	args := c.Called()
	return args.Error(0)
}

func (c *MockContext) Value(key any) any {
	args := c.Called(key)
	return args.Get(0)
}
