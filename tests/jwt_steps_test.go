package tests

import (
	"context"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/require"
)

func (c *sharedTestContext) correctJWTAreSupplied(ctx context.Context) error {
	t := godog.T(ctx)

	clientID, ok := os.LookupEnv("SN_CLIENT_ID")
	require.True(t, ok)
	require.NotEmpty(t, clientID)

	clientSecret, ok := os.LookupEnv("SN_CLIENT_SECRET")
	require.True(t, ok)
	require.NotEmpty(t, clientSecret)

	jwtToken, ok := os.LookupEnv("SN_JWT_TOKEN")
	require.True(t, ok)
	require.NotEmpty(t, jwtToken)

	provider, err := credentials.NewJWTProvider(clientID, clientSecret, NewStaticTokenProvider(jwtToken), credentials.WithInstance(c.instance))
	require.NoError(t, err)

	c.provider = provider
	return nil
}

func (c *sharedTestContext) incorrectJWTAreSupplied(ctx context.Context) error {
	provider, err := credentials.NewJWTProvider("invalid-id", "invalid-secret", NewStaticTokenProvider("invalid-jwt"), credentials.WithInstance(c.instance))
	if err == nil {
		c.provider = provider
	}
	return nil
}

func InitializeJWTAuthenticationScenario(ctx *godog.ScenarioContext) {
	feat := &sharedTestContext{}

	RegisterSharedSteps(ctx, feat)

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
