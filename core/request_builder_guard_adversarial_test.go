// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package core_test

// Adversarial suite for core.ValidateBuilder (request_builder_guard.go).
// Claims under attack:
//  1. returns exactly snerrors.ErrNilRequestBuilder or nil (by identity / errors.Is);
//  2. never panics for ANY input;
//  3. differential oracle vs legacy idiom conversion.IsNil(x) || conversion.IsNil(x.RequestBuilder).
//
// Invariant classes: totality/no-panic, codomain+identity, differential,
// determinism, documentation conformance.

import (
	"testing"
	"unsafe"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type advEmbedder struct {
	core.RequestBuilder
}

type advNoField struct {
	Name string
}

type advEmptyStruct struct{}

type advMapField struct {
	RequestBuilder map[string]int
}

type advFuncField struct {
	RequestBuilder func()
}

type advChanField struct {
	RequestBuilder chan int
}

type advIndirect struct {
	*advEmbedder // promotes a RequestBuilder-named field through a pointer
}

func advValidEmbedded() core.RequestBuilder {
	return core.NewBaseRequestBuilder(newMockAdapter(), "{+baseurl}/adversarial", nil)
}

// assertOnlySentinelOrFail enforces claim 1's codomain.
func assertOnlySentinelOrFail(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		return
	}
	assert.Same(t, snerrors.ErrNilRequestBuilder, err, "codomain violation: not the shared sentinel")
	assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder, "errors.Is must reach the shared sentinel")
}

// TestAdversarialValidateBuilderNeverPanics attacks claim 2 (totality).
// Predicted fault: guard line `value.Elem().FieldByName("RequestBuilder")` runs
// whenever input is a non-nil pointer whose element is not a struct, and
// reflect.Value.FieldByName panics on every non-struct Kind.
func TestAdversarialValidateBuilderNeverPanics(t *testing.T) {
	t.Parallel()

	n := 42
	s := "str"
	f := 3.14
	b := true
	arr := [2]int{1, 2}
	m := map[string]int{"k": 1}
	sl := []int{1}
	ch := make(chan int)
	fn := func() {}
	var iBoxingNil any
	embed := newDummyRequestBuilder()
	pp := &embed
	ppp := &pp
	var iany any

	tests := []struct {
		name  string
		input any
	}{
		{name: "pointer to int", input: &n},
		{name: "pointer to string", input: &s},
		{name: "pointer to float64", input: &f},
		{name: "pointer to bool", input: &b},
		{name: "pointer to array", input: &arr},
		{name: "pointer to map", input: &m},
		{name: "pointer to slice", input: &sl},
		{name: "pointer to chan", input: &ch},
		{name: "pointer to func", input: &fn},
		{name: "pointer to interface boxing nil", input: &iBoxingNil},
		{name: "pointer to empty struct", input: &advEmptyStruct{}},
		{name: "pointer to struct lacking field", input: &advNoField{Name: "x"}},
		{name: "pointer-to-pointer **T outer non-nil", input: pp},
		{name: "triple pointer ***T", input: ppp},
		{name: "pointer to any holding nil", input: &iany},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.NotPanics(t, func() {
				_ = core.ValidateBuilder(tt.input)
			}, "claim 2 violated: ValidateBuilder panicked on %T", tt.input)
		})
	}
}

// TestAdversarialValidateBuilderCodomainOnlySentinelOrNil attacks claims 1+2 on
// terminating inputs: raw nilables, kinds conversion.IsNil names explicitly
// (UnsafePointer), reflect traps. Result must be nil or THE sentinel.
func TestAdversarialValidateBuilderCodomainOnlySentinelOrNil(t *testing.T) {
	t.Parallel()

	nilMap := map[string]int(nil)
	nilSlice := []int(nil)
	nilChan := chan int(nil)
	upNil := unsafe.Pointer(nil)
	type tiny struct{ _ byte }
	upSet := unsafe.Pointer(new(tiny))

	tests := []struct {
		name  string
		input any
	}{
		{name: "untyped nil", input: nil},
		{name: "typed nil *dummyRequestBuilder", input: (*dummyRequestBuilder)(nil)},
		{name: "nil map", input: nilMap},
		{name: "nil slice", input: nilSlice},
		{name: "nil chan", input: nilChan},
		{name: "nil func", input: func() {}},
		{name: "empty non-nil slice", input: []int{}},
		{name: "empty non-nil map", input: map[string]int{}},
		{name: "non-nil func", input: func() {}},
		{name: "plain int", input: 7},
		{name: "uintptr zero", input: uintptr(0)},
		{name: "nil unsafe.Pointer", input: upNil},
		{name: "set unsafe.Pointer", input: upSet},
		{name: "struct value lacking field", input: advNoField{Name: "x"}},
		{name: "ptr to struct with nil map field named RequestBuilder", input: &advMapField{}},
		{name: "ptr to struct with nil chan field named RequestBuilder", input: &advChanField{}},
		{name: "*BaseRequestBuilder itself (no RequestBuilder-named field)", input: core.NewBaseRequestBuilder(newMockAdapter(), "{+baseurl}/t", nil)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var err error
			require.NotPanics(t, func() {
				err = core.ValidateBuilder(tt.input)
			})
			assertOnlySentinelOrFail(t, err)
		})
	}
}

// TestAdversarialDifferentialAgainstLegacyIdiom attacks claim 3 head-on:
// guard(x)==nil <=> !(IsNil(x) || IsNil(x.RequestBuilder)), bidirectional, plus
// determinism. Shapes where legacy cannot compile are oracle-excluded.
func TestAdversarialDifferentialAgainstLegacyIdiom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		build          func() any
		legacy         func(x any) bool
		legacyMayPanic bool
	}{
		{
			name:  "valid *embedder",
			build: func() any { return &advEmbedder{RequestBuilder: advValidEmbedded()} },
			legacy: func(x any) bool {
				p := x.(*advEmbedder)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
		},
		{
			name:  "typed-nil receiver",
			build: func() any { return (*advEmbedder)(nil) },
			legacy: func(x any) bool {
				p := x.(*advEmbedder)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
		},
		{
			name:  "receiver ok, embedded interface nil",
			build: func() any { return &advEmbedder{} },
			legacy: func(x any) bool {
				p := x.(*advEmbedder)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
		},
		{
			name:  "embedded interface holding typed-nil *core.BaseRequestBuilder",
			build: func() any { return &advEmbedder{RequestBuilder: (*core.BaseRequestBuilder)(nil)} },
			legacy: func(x any) bool {
				p := x.(*advEmbedder)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
		},
		{
			name:  "embedded interface holding typed-nil self-referential builder",
			build: func() any { return &advEmbedder{RequestBuilder: (*advEmbedder)(nil)} },
			legacy: func(x any) bool {
				p := x.(*advEmbedder)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
		},
		{
			name:  "VALUE receiver with valid embedded builder",
			build: func() any { return advEmbedder{RequestBuilder: advValidEmbedded()} },
			legacy: func(x any) bool {
				p := x.(advEmbedder)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
		},
		{
			name:  "VALUE receiver with nil embedded builder",
			build: func() any { return advEmbedder{} },
			legacy: func(x any) bool {
				p := x.(advEmbedder)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
		},
		{
			name:  "indirect embedding, inner present",
			build: func() any { return &advIndirect{advEmbedder: &advEmbedder{RequestBuilder: advValidEmbedded()}} },
			legacy: func(x any) bool {
				p := x.(*advIndirect)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
		},
		{
			name:  "indirect embedding, inner nil (legacy crashes)",
			build: func() any { return &advIndirect{} },
			legacy: func(x any) bool {
				p := x.(*advIndirect)
				return conversion.IsNil(p) || conversion.IsNil(p.RequestBuilder)
			},
			legacyMayPanic: true,
		},
		{
			name:   "struct without RequestBuilder field (oracle N/A)",
			build:  func() any { return &advNoField{Name: "x"} },
			legacy: nil,
		},
		{
			name:   "**T to valid builder (oracle N/A, totality only)",
			build:  func() any { b := newDummyRequestBuilder(); return &b },
			legacy: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			x := tt.build()

			var guardErr error
			require.NotPanics(t, func() {
				guardErr = core.ValidateBuilder(x)
			}, "guard panicked where the oracle requires a verdict")

			guardErr2 := core.ValidateBuilder(x)
			assert.Equal(t, guardErr == nil, guardErr2 == nil, "nondeterministic guard verdict")

			assertOnlySentinelOrFail(t, guardErr)

			if tt.legacy == nil {
				return
			}

			var legacyInvalid, legacyPanicked bool
			func() {
				defer func() {
					if r := recover(); r != nil {
						legacyPanicked = true
					}
				}()
				legacyInvalid = tt.legacy(x)
			}()

			if legacyPanicked {
				require.True(t, tt.legacyMayPanic, "legacy idiom crashed unexpectedly")
				t.Logf("oracle note: legacy idiom panics; guard returned err=%v (safer, logged not counted)", guardErr)
				return
			}

			assert.Equal(t, !legacyInvalid, guardErr == nil,
				"differential oracle violated: legacy invalid=%v, guard err==%v (%T)",
				legacyInvalid, guardErr, x)
		})
	}
}

// TestAdversarialDocumentationLies verifies the docstring lie found by this
// suite is FIXED: a struct whose RequestBuilder-named field is merely a non-nil
// map/chan/func implements nothing and must now receive the sentinel (the
// original implementation accepted these impostors with nil).
func TestAdversarialDocumentationLies(t *testing.T) {
	t.Parallel()

	lies := []struct {
		name    string
		builder any
	}{
		{name: "non-nil map field impersonates builder", builder: &advMapField{RequestBuilder: map[string]int{"a": 1}}},
		{name: "non-nil chan field impersonates builder", builder: &advChanField{RequestBuilder: make(chan int, 1)}},
		{name: "non-nil func field impersonates builder", builder: &advFuncField{RequestBuilder: func() {}}},
	}

	for _, tt := range lies {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := core.ValidateBuilder(tt.builder)
			assertOnlySentinelOrFail(t, err)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder,
				"documented behavior: impostor fields get the sentinel, not acceptance")
		})
	}
}
