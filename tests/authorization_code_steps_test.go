package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/require"
)

func (c *sharedTestContext) correctAuthorizationCodeAreSupplied(ctx context.Context) error {
	t := godog.T(ctx)

	clientID, ok := os.LookupEnv("SN_CLIENT_ID")
	require.True(t, ok)
	require.NotEmpty(t, clientID)

	clientSecret, ok := os.LookupEnv("SN_CLIENT_SECRET")
	require.True(t, ok)
	require.NotEmpty(t, clientSecret)

	// 1. Activate httpmock to catch the token exchange
	httpmock.Activate()
	// Allow calls to the local listener (not mocked)
	httpmock.RegisterNoResponder(httpmock.InitialTransport.RoundTrip)

	// Mock the token response
	tokenURL := fmt.Sprintf("https://%s.service-now.com/oauth_token.do", c.instance)
	httpmock.RegisterResponder("POST", tokenURL,
		httpmock.NewStringResponder(200, `{
			"access_token": "mock_access_token",
			"refresh_token": "mock_refresh_token",
			"token_type": "Bearer",
			"expires_in": 3600
		}`))

	// Mock the incident table response
	incidentURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/table/incident", c.instance)
	httpmock.RegisterResponder("GET", incidentURL,
		httpmock.NewStringResponder(200, `{"result": []}`).
			HeaderSet(http.Header{"Content-Type": []string{"application/json"}}))

	opts := []credentials.AuthOption{
		credentials.WithInstance(c.instance),
		// 2. Custom URL Opener: Simulates the browser "Redirecting" to the local server
		credentials.WithURLOpener(func(authURL string) error {
			u, err := url.Parse(authURL)
			if err != nil {
				return err
			}

			state := u.Query().Get("state")
			redirectURI := u.Query().Get("redirect_uri")

			// Simulate the successful login redirect from ServiceNow back to our local listener
			callbackURL := fmt.Sprintf("%s?code=mock_code&state=%s", redirectURI, state)

			// We use the default client, but since we set httpmock.RegisterNoResponder
			// it will fall through to the real network for localhost
			resp, err := http.Get(callbackURL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			return nil
		}),
	}

	provider, err := credentials.NewAuthorizationCodeProvider(clientID, clientSecret, opts...)
	require.NoError(t, err)

	c.provider = provider
	return nil
}

func (c *sharedTestContext) incorrectAuthorizationCodeAreSupplied(ctx context.Context) error {
	t := godog.T(ctx)

	// Activate httpmock for the failure case as well
	httpmock.Activate()
	httpmock.RegisterNoResponder(httpmock.InitialTransport.RoundTrip)

	// Mock the token response as a failure
	tokenURL := fmt.Sprintf("https://%s.service-now.com/oauth_token.do", c.instance)
	httpmock.RegisterResponder("POST", tokenURL,
		httpmock.NewStringResponder(400, `{"error": "invalid_grant"}`))

	opts := []credentials.AuthOption{
		credentials.WithInstance(c.instance),
		// Mock opener that just succeeds without doing anything (no redirect)
		// This will cause the local server to time out or we can simulate a different error
		credentials.WithURLOpener(func(authURL string) error {
			u, err := url.Parse(authURL)
			if err != nil {
				return err
			}

			state := u.Query().Get("state")
			redirectURI := u.Query().Get("redirect_uri")

			// Simulate the successful login redirect from ServiceNow back to our local listener
			callbackURL := fmt.Sprintf("%s?code=mock_code&state=%s", redirectURI, state)

			// We use the default client, but since we set httpmock.RegisterNoResponder
			// it will fall through to the real network for localhost
			resp, err := http.Get(callbackURL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			return nil
		}),
	}

	provider, err := credentials.NewAuthorizationCodeProvider("invalid-id", "invalid-secret", opts...)
	require.NoError(t, err)

	c.provider = provider
	return nil
}

func (c *sharedTestContext) authenticationFlowIsCancelled(ctx context.Context) error {
	t := godog.T(ctx)

	// Activate httpmock for the cancellation case as well
	httpmock.Activate()
	httpmock.RegisterNoResponder(httpmock.InitialTransport.RoundTrip)

	opts := []credentials.AuthOption{
		credentials.WithInstance(c.instance),
		// Mock opener that just succeeds without doing anything (no redirect)
		// This simulates the user closing the browser or cancelling before login.
		credentials.WithURLOpener(func(authURL string) error {
			// Do nothing - simulated user closed the browser or cancelled
			return nil
		}),
	}

	// We can use any valid or invalid IDs here since it won't get to the token exchange
	provider, err := credentials.NewAuthorizationCodeProvider("any-id", "any-secret", opts...)
	require.NoError(t, err)

	c.provider = provider
	return nil
}

func InitializeAuthorizationCodeAuthenticationScenario(ctx *godog.ScenarioContext) {
	feat := &sharedTestContext{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Ensure httpmock is clean for each scenario
		httpmock.Reset()
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		httpmock.DeactivateAndReset()
		return ctx, nil
	})

	RegisterSharedSteps(ctx, feat)

	ctx.When(`^correct credentials are supplied$`, feat.correctAuthorizationCodeAreSupplied)
	ctx.When(`^incorrect credentials are supplied$`, feat.incorrectAuthorizationCodeAreSupplied)
	ctx.When(`^authentication flow is cancelled$`, feat.authenticationFlowIsCancelled)
}

func TestAuthorizationCodeAuthenticationFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeAuthorizationCodeAuthenticationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/authentication/authorization_code.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
