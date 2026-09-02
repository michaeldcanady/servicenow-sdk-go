// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package store

import (
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// AccessorFunc[S,T] defines a generic function signature for retrieving a value from a backing store
// using a specified key and converting it to a desired type.
type AccessorFunc[S store.BackingStore, T any] func(S, string) (T, error)

// DefaultStoreAccessorFunc[S, T] is a generic implementation of AccessorFunc that retrieves a value
// from a backing store and attempts to convert it to the specified type.
func DefaultStoreAccessorFunc[S store.BackingStore, T any](store S, key string) (T, error) {
	var result T

	if conversion.IsNil(store) {
		return result, snerrors.ErrNilStore
	}
	if strings.TrimSpace(key) == "" {
		return result, snerrors.ErrEmptyKey
	}

	val, err := store.Get(key)
	if err != nil {
		return result, err
	}

	if err := conversion.As2(val, &result, true); err != nil {
		return result, err
	}

	return result, nil
}
