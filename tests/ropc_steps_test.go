package tests

import (
	"context"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/require"
)

func (c *sharedTestContext) correctROPCAreSupplied(ctx context.Context) error {
	t := godog.T(ctx)

	clientID, ok := os.LookupEnv("SN_CLIENT_ID")
	require.True(t, ok)
	require.NotEmpty(t, clientID)

	clientSecret, ok := os.LookupEnv("SN_CLIENT_SECRET")
	require.True(t, ok)
	require.NotEmpty(t, clientSecret)

	username, ok := os.LookupEnv("SN_USERNAME")
	require.True(t, ok)
	require.NotEmpty(t, username)

	password, ok := os.LookupEnv("SN_PASSWORD")
	require.True(t, ok)
	require.NotEmpty(t, password)

	provider, err := credentials.NewROPCProvider(clientID, clientSecret, username, password, credentials.WithInstance(c.instance))
	require.NoError(t, err)

	c.provider = provider
	return nil
}

func (c *sharedTestContext) incorrectROPCAreSupplied(ctx context.Context) error {
	provider, err := credentials.NewROPCProvider("invalid-id", "invalid-secret", "invalid-user", "invalid-pass", credentials.WithInstance(c.instance))
	if err == nil {
		c.provider = provider
	}
	return nil
}

func InitializeROPCAuthenticationScenario(ctx *godog.ScenarioContext) {
	feat := &sharedTestContext{}

	RegisterSharedSteps(ctx, feat)

	ctx.When(`^correct credentials are supplied$`, feat.correctROPCAreSupplied)
	ctx.When(`^incorrect credentials are supplied$`, feat.incorrectROPCAreSupplied)
}

func TestROPCAuthenticationFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeROPCAuthenticationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/authentication/ropc.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
