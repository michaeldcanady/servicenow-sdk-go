// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package core_test

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dummyRequestBuilder struct {
	core.RequestBuilder
}

// Fixtures for Go's selector rule: outer.RequestBuilder resolves to the shallow
// (nil) field, so the deeper valid builder must not satisfy the guard.
type (
	selectorLeaf    struct{ core.RequestBuilder }
	selectorDeep    struct{ selectorLeaf }
	selectorShallow struct{ core.RequestBuilder }
	selectorOuter   struct {
		selectorDeep
		selectorShallow
	}
)

func newDummyRequestBuilder() *dummyRequestBuilder {
	return &dummyRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(nil, "{+baseurl}/test", nil),
	}
}

func TestValidateBuilder(t *testing.T) {
	t.Parallel()

	valid := newDummyRequestBuilder()

	tests := []struct {
		name    string
		builder any
		want    error
	}{
		{
			name:    "valid builder with embedded RequestBuilder",
			builder: valid,
			want:    nil,
		},
		{
			name:    "typed-nil receiver passed through interface{} (classic trap)",
			builder: (*dummyRequestBuilder)(nil),
			want:    snerrors.ErrNilRequestBuilder,
		},
		{
			name:    "untyped nil",
			builder: nil,
			want:    snerrors.ErrNilRequestBuilder,
		},
		{
			name: "nil embedded RequestBuilder interface",
			builder: &struct {
				core.RequestBuilder
			}{},
			want: snerrors.ErrNilRequestBuilder,
		},
		{
			name: "shallower nil builder wins over a deeper valid one (Go selector rule)",
			builder: selectorOuter{
				selectorDeep: selectorDeep{selectorLeaf{
					RequestBuilder: core.NewBaseRequestBuilder(nil, "{+baseurl}/t", nil),
				}},
			},
			want: snerrors.ErrNilRequestBuilder,
		},
		{
			name: "struct value (non-pointer) with valid embedded builder",
			builder: dummyRequestBuilder{
				RequestBuilder: core.NewBaseRequestBuilder(nil, "{+baseurl}/t", nil),
			},
			want: nil,
		},
		{
			name:    "pointer to non-struct (int)",
			builder: func() any { n := 42; return &n }(),
			want:    snerrors.ErrNilRequestBuilder,
		},
		{
			name: "field named RequestBuilder of non-interface kind",
			builder: &struct {
				RequestBuilder map[string]int
			}{RequestBuilder: map[string]int{"a": 1}},
			want: snerrors.ErrNilRequestBuilder,
		},
		{
			name: "struct without RequestBuilder field",
			builder: &struct {
				Name string
			}{Name: "not a builder"},
			want: snerrors.ErrNilRequestBuilder,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := core.ValidateBuilder(tt.builder)

			if tt.want == nil {
				require.NoError(t, err)
				return
			}

			require.Error(t, err)
			assert.ErrorIs(t, err, tt.want, "guard must return the shared sentinel by identity, not a lookalike")
		})
	}
}

func TestValidateBuilderSentinelIdentity(t *testing.T) {
	t.Parallel()

	err := core.ValidateBuilder((*dummyRequestBuilder)(nil))

	require.Error(t, err)
	assert.Same(t, snerrors.ErrNilRequestBuilder, err, "must reuse the shared sentinel by identity, never errors.New a lookalike")
	assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder, "exact identity is what makes errors.Is work across modules")
}

func TestValidateBuilderDoesNotPanicOnTypedNilChain(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("ValidateBuilder must not panic on typed-nil receivers, got panic: %v", r)
		}
	}()

	var nested *dummyRequestBuilder

	assert.ErrorIs(t, core.ValidateBuilder(nested), snerrors.ErrNilRequestBuilder)
}

func TestValidateBuilderWithAdapterBackedBuilder(t *testing.T) {
	t.Parallel()

	builder := &dummyRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(
			newMockAdapter(),
			"{+baseurl}/test",
			nil,
		),
	}

	require.NoError(t, core.ValidateBuilder(builder))
}

// newMockAdapter returns the smallest possible non-nil abstractions.RequestAdapter.
func newMockAdapter() abstractions.RequestAdapter {
	return &mockRequestAdapter{}
}

type mockRequestAdapter struct {
	abstractions.RequestAdapter
}
