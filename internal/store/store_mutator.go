// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package store

import (
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type MutatorFunc[S store.BackingStore, T any] func(bs S, key string, value T) error

// DefaultStoreMutatorFunc[T] sets the [store.BackingStore] at the provided key to the provided value.
func DefaultStoreMutatorFunc[S store.BackingStore, T any](store store.BackingStore, key string, value T) error {
	if conversion.IsNil(store) {
		return snerrors.ErrNilStore
	}

	if strings.TrimSpace(key) == "" {
		return snerrors.ErrEmptyKey
	}

	return store.Set(key, value)
}
