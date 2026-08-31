// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package mocking

import (
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/mock"
)

type MockBackingStoreFactorySetter struct {
	mock.Mock
}

func NewMockBackingStoreFactorySetter() *MockBackingStoreFactorySetter {
	return &MockBackingStoreFactorySetter{
		mock.Mock{},
	}
}

func (m *MockBackingStoreFactorySetter) SetBackingStoreFactory(factory store.BackingStoreFactory) error {
	args := m.Called(factory)
	return args.Error(0)
}
