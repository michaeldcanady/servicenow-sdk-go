package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/require"
)

type jwtTestContext struct {
	*sharedTestContext
	server *httptest.Server
}

func (c *jwtTestContext) correctJWTAreSupplied(ctx context.Context) error {
	t := godog.T(ctx)

	var clientID, clientSecret, jwtToken string
	var ok bool

	if isOffline() {
		clientID = "mock-client-id"
		clientSecret = "mock-client-secret"
		jwtToken = "mock-jwt-token"
	} else {
		clientID, ok = os.LookupEnv("SN_CLIENT_ID")
		require.True(t, ok)
		require.NotEmpty(t, clientID)

		clientSecret, ok = os.LookupEnv("SN_CLIENT_SECRET")
		require.True(t, ok)
		require.NotEmpty(t, clientSecret)

		jwtToken, ok = os.LookupEnv("SN_JWT_TOKEN")
		require.True(t, ok)
		require.NotEmpty(t, jwtToken)
	}

	if isOffline() {
		c.instance = c.server.URL
	}

	provider, err := credentials.NewJWTProvider(clientID, clientSecret, NewStaticTokenProvider(jwtToken), credentials.WithURL(c.instance))
	require.NoError(t, err)

	c.provider = provider
	return nil
}

func (c *jwtTestContext) incorrectJWTAreSupplied(ctx context.Context) error {
	instance := c.instance
	if isOffline() {
		instance = c.server.URL
	}

	provider, err := credentials.NewJWTProvider("invalid-id", "invalid-secret", NewStaticTokenProvider("invalid-jwt"), credentials.WithURL(instance))
	if err == nil {
		c.provider = provider
	}
	return nil
}

func InitializeJWTAuthenticationScenario(ctx *godog.ScenarioContext) {
	feat := &jwtTestContext{
		sharedTestContext: &sharedTestContext{},
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		if isOffline() {
			mux := http.NewServeMux()
			mux.HandleFunc("/oauth_token.do", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `{
					"access_token": "mock_jwt_access_token",
					"token_type": "Bearer",
					"expires_in": 3600
				}`)
			})
			mux.HandleFunc("/api/now/v1/table/incident", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `{"result": []}`)
			})
			feat.server = httptest.NewServer(mux)
		}
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		if feat.server != nil {
			feat.server.Close()
		}
		return ctx, nil
	})

	RegisterSharedSteps(ctx, feat.sharedTestContext)

	ctx.When(`^correct credentials are supplied$`, feat.correctJWTAreSupplied)
	ctx.When(`^incorrect credentials are supplied$`, feat.incorrectJWTAreSupplied)
}

func TestJWTAuthenticationFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeJWTAuthenticationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/authentication/jwt.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
