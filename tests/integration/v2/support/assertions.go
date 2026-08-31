// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package support

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
)

// RequireNoError asserts that the world's error is nil.
func RequireNoError(t *testing.T, w *World) bool {
	t.Helper()
	if w.Err != nil {
		t.Errorf("expected no error, got: %v", w.Err)
		return false
	}
	return true
}

// RequireNoErrors asserts both Err and AuthErr are nil.
func RequireNoErrors(t *testing.T, w *World) bool {
	t.Helper()
	if w.AuthErr != nil {
		t.Errorf("expected no auth error, got: %v", w.AuthErr)
		return false
	}
	if w.Err != nil {
		t.Errorf("expected no error, got: %v", w.Err)
		return false
	}
	return true
}

// RequireError asserts that the world has a non-nil error.
func RequireError(t *testing.T, w *World) bool {
	t.Helper()
	if w.Err == nil && w.AuthErr == nil {
		t.Error("expected an error, but got none")
		return false
	}
	return true
}

// RequireAuthError asserts that the world has a non-nil auth error.
func RequireAuthError(t *testing.T, w *World) bool {
	t.Helper()
	if w.AuthErr == nil {
		t.Error("expected an auth error, but got none")
		return false
	}
	return true
}

// ExtractErrorMessage returns a human-readable string from an error.
// For ServiceNowError, this uses the Error() method which extracts the message.
func ExtractErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	// ServiceNowError implements error — its Error() method extracts the message
	var snErr *core.ServiceNowError
	if errors.As(err, &snErr) {
		return snErr.Error()
	}

	return err.Error()
}

// RequireErrorMessageContains asserts the error message contains the given substring.
func RequireErrorMessageContains(t *testing.T, w *World, substr string) bool {
	t.Helper()
	err := w.Err
	if err == nil {
		err = w.AuthErr
	}
	if err == nil {
		t.Errorf("expected error containing %q, but got no error", substr)
		return false
	}
	msg := ExtractErrorMessage(err)
	if !strings.Contains(msg, substr) {
		t.Errorf("expected error message to contain %q, got %q", substr, msg)
		return false
	}
	return true
}

// CleanupResource attempts to DELETE a resource by sys_id using the table API.
// Returns any error encountered during deletion.
func CleanupResource(ctx context.Context, w *World, sysID string) error {
	if w.Client == nil || sysID == "" {
		return nil
	}
	return w.Client.Now().Table("incident").ByID(sysID).Delete(ctx, nil)
}

// CleanupResourceLog attempts to DELETE a resource, logging a warning on failure
// instead of returning an error. Intended for AfterScenario hooks where cleanup
// failure should not abort the test run but should be visible in output.
func CleanupResourceLog(ctx context.Context, w *World, sysID string) {
	if err := CleanupResource(ctx, w, sysID); err != nil && IsE2E() {
		fmt.Fprintf(os.Stderr, "WARNING: cleanup failed for sys_id %s: %v\n", sysID, err)
	}
}

// RequireRevocationSuccess asserts the revocation succeeded.
func RequireRevocationSuccess(ctx context.Context, w *World) (context.Context, error) {
	if w.RevocationErr != nil {
		return ctx, fmt.Errorf("expected revocation to succeed, got: %v", w.RevocationErr)
	}
	return ctx, nil
}
