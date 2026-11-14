// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package core

import (
	"reflect"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
)

// maxEmbedDepth bounds how deep through anonymous struct/pointer embeddings the
// builder lookup follows. The module-wide convention is a direct embed
// (`struct { core.RequestBuilder }`), so 2 is generous; the cap exists so no
// adversarial nesting can turn this guard into a stack consumer.
const maxEmbedDepth = 2

// ValidateBuilder enforces the shared nil-preconditions that every request-builder
// verb and navigation method repeats before touching its receiver chain: a
// present receiver and a present embedded core.RequestBuilder (see CLAUDE.md's
// constructor-triad section, ADR-006, and ADR-022).
//
// Accepted inputs are pointers to structs and struct values alike. Anything else
// — untyped nil, typed-nil pointers (including typed nils hidden inside an
// interface), pointers to non-structs, multi-level pointers, structs without an
// embedded builder — reports snerrors.ErrNilRequestBuilder. The function is
// total: it never panics, including on inputs that make reflection traversal
// dangerous, on shapes where the legacy inline idiom itself panics, and on
// unexported embedding fields that reflect forbids reading through Value.Interface.
//
// The embedded builder is located as a field named "RequestBuilder" whose value
// is an interface holding a non-nil concrete value (the module-wide embedding
// contract: `struct { core.RequestBuilder }`). Lookup follows anonymous struct /
// non-nil anonymous-pointer promotions up to maxEmbedDepth levels; promotion
// through a NIL intermediate pointer is reported invalid rather than followed,
// since the legacy idiom panics on exactly those shapes. Fields of non-interface
// kinds (map, chan, func, concrete structs) do not satisfy the contract.
//
// Call sites:
//
//	if err := core.ValidateBuilder(rB); err != nil {
//	    return nil, err
//	}
//
// Returns the shared snerrors.ErrNilRequestBuilder by identity so errors.Is
// works uniformly across modules (ADR-006). The request-adapter precondition is
// deliberately NOT folded in here: several verbs interleave additional checks
// (requestConfiguration, query parameters) between the builder guard and the
// adapter guard, and folding them together would silently reorder which
// sentinel a caller holding multiple nils receives.
func ValidateBuilder(builder any) error {
	if conversion.IsNil(builder) {
		return snerrors.ErrNilRequestBuilder
	}

	value := reflect.ValueOf(builder)

	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return snerrors.ErrNilRequestBuilder
		}

		value = value.Elem()
	}

	if !hasEmbeddedRequestBuilder(value) {
		return snerrors.ErrNilRequestBuilder
	}

	return nil
}

// hasEmbeddedRequestBuilder reports whether value (a struct) carries an embedded
// core.RequestBuilder satisfying the contract described on ValidateBuilder.
//
// The field is resolved with reflect's FieldByName, which applies Go's own
// selector rules: the shallowest promoted RequestBuilder wins, and two at the
// same depth are ambiguous (not found). That keeps the guard in step with what
// rB.RequestBuilder actually selects, instead of accepting a deeper non-nil
// field the selector never reaches.
func hasEmbeddedRequestBuilder(value reflect.Value) bool {
	if value.Kind() != reflect.Struct {
		return false
	}

	field, ok := value.Type().FieldByName("RequestBuilder")
	if !ok || len(field.Index) > maxEmbedDepth+1 {
		return false
	}

	// FieldByIndexErr reports a nil embedded pointer on the path instead of panicking.
	resolved, err := value.FieldByIndexErr(field.Index)
	if err != nil {
		return false
	}

	return resolved.Kind() == reflect.Interface && !isNilValue(resolved)
}

// isNilValue reports whether v is nil at the value level: either a nilable-kind
// zero value or an interface holding one. Unlike boxing through Value.Interface
// (which panics on unexported fields), IsNil and Elem are always permitted.
func isNilValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer, reflect.Interface:
		if v.IsNil() {
			return true
		}

		if v.Kind() == reflect.Interface {
			return isNilValue(v.Elem())
		}

		return false
	default:
		return false
	}
}
