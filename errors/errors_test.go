package errors_test

import (
	"fmt"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
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

func TestNewValidationError(t *testing.T) {
	err := snerrors.NewValidationError("foo")
	assert.EqualError(t, err, "foo cannot be nil")
}
