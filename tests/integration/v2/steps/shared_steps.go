package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
)

// RegisterSharedSteps registers common step definitions used across all domains.
func RegisterSharedSteps(sc *godog.ScenarioContext) {
	// ── Background / Given ────────────────────────────────────────────

	sc.Step(`^I have a valid ServiceNow instance and credentials$`, iHaveAValidInstanceAndCredentials)
	sc.Step(`^I have initialized the ServiceNow client$`, iHaveInitializedTheClient)

	// ── Then: Success assertions ──────────────────────────────────────

	sc.Step(`^the response should not be an error$`, theResponseShouldNotBeAnError)
	sc.Step(`^authentication should succeed$`, authenticationShouldSucceed)
	sc.Step(`^all requests should succeed$`, allRequestsShouldSucceed)
	sc.Step(`^the revocation should succeed$`, theRevocationShouldSucceed)

	// ── Then: Error assertions ────────────────────────────────────────

	sc.Step(`^the response should be a (\d+) error$`, theResponseShouldBeAStatusError)
	sc.Step(`^the response should be an API error$`, theResponseShouldBeAnAPIError)
	sc.Step(`^authentication should fail$`, authenticationShouldFail)
	sc.Step(`^credential creation should fail$`, credentialCreationShouldFail)
	sc.Step(`^the API request should fail with an authentication error$`, theAPIRequestShouldFailWithAuthError)
	sc.Step(`^the error should be authentication-related$`, theErrorShouldBeAuthRelated)
	sc.Step(`^the error message should indicate missing parameters$`, theErrorMessageShouldIndicateMissingParameters)
}

// ── Step Implementations ─────────────────────────────────────────────────

func iHaveAValidInstanceAndCredentials(ctx context.Context) (context.Context, error) {
	// Environment is already loaded by TestMain.
	// This step exists for Gherkin readability — it's a no-op.
	return ctx, nil
}

func iHaveInitializedTheClient(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w == nil {
		return ctx, fmt.Errorf("world not initialized")
	}

	if w.Client != nil {
		return ctx, nil
	}

	if err := support.NewClientWithBasicAuth(w); err != nil {
		return ctx, err
	}
	return support.WithWorld(ctx, w), nil
}

func theResponseShouldNotBeAnError(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Err != nil {
		return ctx, fmt.Errorf("expected no error, got: %v", w.Err)
	}
	return ctx, nil
}

func theResponseShouldBeAStatusError(ctx context.Context, statusCode int) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Err == nil {
		return ctx, fmt.Errorf("expected %d error, but got no error", statusCode)
	}
	return ctx, nil
}

func theResponseShouldBeAnAPIError(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Err == nil {
		return ctx, fmt.Errorf("expected an API error, but got no error")
	}
	return ctx, nil
}

func authenticationShouldSucceed(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr != nil {
		return ctx, fmt.Errorf("expected authentication to succeed, got: %v", w.AuthErr)
	}
	return ctx, nil
}

func authenticationShouldFail(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr == nil && w.Err == nil {
		return ctx, fmt.Errorf("expected authentication to fail, but it succeeded")
	}
	return ctx, nil
}

func credentialCreationShouldFail(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr == nil {
		return ctx, fmt.Errorf("expected credential creation to fail, but it succeeded")
	}
	return ctx, nil
}

func theAPIRequestShouldFailWithAuthError(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Err == nil {
		return ctx, fmt.Errorf("expected API request to fail with auth error, but got no error")
	}
	return ctx, nil
}

func theErrorShouldBeAuthRelated(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	err := w.Err
	if err == nil {
		err = w.AuthErr
	}
	if err == nil {
		return ctx, fmt.Errorf("expected an authentication-related error, but got none")
	}
	// Any error from an auth flow is auth-related
	return ctx, nil
}

func theErrorMessageShouldIndicateMissingParameters(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	err := w.AuthErr
	if err == nil {
		err = w.Err
	}
	if err == nil {
		return ctx, fmt.Errorf("expected error about missing parameters, but got no error")
	}
	return ctx, nil
}

func allRequestsShouldSucceed(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Err != nil {
		return ctx, fmt.Errorf("expected all requests to succeed, but got: %v", w.Err)
	}
	return ctx, nil
}

func theRevocationShouldSucceed(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.RevocationErr != nil {
		return ctx, fmt.Errorf("expected revocation to succeed, got: %v", w.RevocationErr)
	}
	return ctx, nil
}
