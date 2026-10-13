// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package errors_test

import (
	"fmt"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/stretchr/testify/assert"
)

func TestSentinels_ErrorsIs(t *testing.T) {
	sentinels := map[string]error{
		"ErrNilRequestAdapter":       snerrors.ErrNilRequestAdapter,
		"ErrNilRequestBuilder":       snerrors.ErrNilRequestBuilder,
		"ErrNilResponse":             snerrors.ErrNilResponse,
		"ErrNilContext":              snerrors.ErrNilContext,
		"ErrNilConfig":               snerrors.ErrNilConfig,
		"ErrNilBody":                 snerrors.ErrNilBody,
		"ErrNilInput":                snerrors.ErrNilInput,
		"ErrNilRequestConfiguration": snerrors.ErrNilRequestConfiguration,
		"ErrNilQueryParameters":      snerrors.ErrNilQueryParameters,
		"ErrNilFactory":              snerrors.ErrNilFactory,
		"ErrNilStore":                snerrors.ErrNilStore,
		"ErrNilPathParameters":       snerrors.ErrNilPathParameters,
		"ErrEmptyPathParameters":     snerrors.ErrEmptyPathParameters,
		"ErrNilMutator":              snerrors.ErrNilMutator,
		"ErrNilModel":                snerrors.ErrNilModel,
		"ErrEmptyMiddleware":         snerrors.ErrEmptyMiddleware,
		"ErrEmptyKey":                snerrors.ErrEmptyKey,
		"ErrNilRequestInfo":          snerrors.ErrNilRequestInfo,
		"ErrNilClient":               snerrors.ErrNilClient,
		"ErrNilResult":               snerrors.ErrNilResult,
		"ErrWrongResponseType":       snerrors.ErrWrongResponseType,
		"ErrParsing":                 snerrors.ErrParsing,
		"ErrEmptyURI":                snerrors.ErrEmptyURI,
		"ErrEmptyBaseURL":            snerrors.ErrEmptyBaseURL,
		"ErrEmptyClientID":           snerrors.ErrEmptyClientID,
		"ErrEmptyClientSecret":       snerrors.ErrEmptyClientSecret,
		"ErrEmptyUsername":           snerrors.ErrEmptyUsername,
		"ErrEmptyPassword":           snerrors.ErrEmptyPassword,
		"ErrNilCallback":             snerrors.ErrNilCallback,
		"ErrNilParams":               snerrors.ErrNilParams,
	}

	for name, sentinel := range sentinels {
		t.Run(name, func(t *testing.T) {
			wrapped := fmt.Errorf("context: %w", sentinel)
			assert.ErrorIsf(t, wrapped, sentinel, "wrapped error should satisfy errors.Is against %s", name)
		})
	}
}

func TestSentinels_Distinct(t *testing.T) {
	sentinels := []error{
		snerrors.ErrNilRequestAdapter,
		snerrors.ErrNilRequestBuilder,
		snerrors.ErrNilResponse,
		snerrors.ErrNilContext,
		snerrors.ErrNilConfig,
		snerrors.ErrNilBody,
		snerrors.ErrNilInput,
		snerrors.ErrNilRequestConfiguration,
		snerrors.ErrNilQueryParameters,
		snerrors.ErrNilFactory,
		snerrors.ErrNilStore,
		snerrors.ErrNilPathParameters,
		snerrors.ErrEmptyPathParameters,
		snerrors.ErrNilMutator,
		snerrors.ErrNilModel,
		snerrors.ErrEmptyMiddleware,
		snerrors.ErrEmptyKey,
		snerrors.ErrNilRequestInfo,
		snerrors.ErrNilClient,
		snerrors.ErrNilResult,
		snerrors.ErrWrongResponseType,
		snerrors.ErrParsing,
		snerrors.ErrEmptyURI,
		snerrors.ErrEmptyBaseURL,
		snerrors.ErrEmptyClientID,
		snerrors.ErrEmptyClientSecret,
		snerrors.ErrEmptyUsername,
		snerrors.ErrEmptyPassword,
		snerrors.ErrNilCallback,
		snerrors.ErrNilParams,
	}

	for i, a := range sentinels {
		for j, b := range sentinels {
			if i == j {
				continue
			}
			assert.NotErrorIsf(t, a, b, "sentinel %d (%v) should not satisfy errors.Is against distinct sentinel %d (%v)", i, a, j, b)
		}
	}
}

func TestSentinels_EmptyCredentialMessages(t *testing.T) {
	want := map[string]string{
		"ErrEmptyClientID":     "clientId is empty",
		"ErrEmptyClientSecret": "clientSecret is empty",
		"ErrEmptyBaseURL":      "baseURL is empty",
		"ErrEmptyUsername":     "username is empty",
		"ErrEmptyPassword":     "password is empty",
	}
	sentinels := map[string]error{
		"ErrEmptyClientID":     snerrors.ErrEmptyClientID,
		"ErrEmptyClientSecret": snerrors.ErrEmptyClientSecret,
		"ErrEmptyBaseURL":      snerrors.ErrEmptyBaseURL,
		"ErrEmptyUsername":     snerrors.ErrEmptyUsername,
		"ErrEmptyPassword":     snerrors.ErrEmptyPassword,
	}

	for name, sentinel := range sentinels {
		t.Run(name, func(t *testing.T) {
			assert.EqualError(t, sentinel, want[name])
		})
	}
}

func TestNewValidationError(t *testing.T) {
	err := snerrors.NewValidationError("foo")
	assert.EqualError(t, err, "foo cannot be nil")
}
