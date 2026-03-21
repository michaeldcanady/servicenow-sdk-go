package tests

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/require"
)

func (c *sharedTestContext) correctClientCredentialsAreSupplied(ctx context.Context) error {
	t := godog.T(ctx)

	clientID, ok := os.LookupEnv("SN_CLIENT_ID")
	require.True(t, ok)
	require.NotEmpty(t, clientID)

	clientSecret, ok := os.LookupEnv("SN_CLIENT_SECRET")
	require.True(t, ok)
	require.NotEmpty(t, clientSecret)

	fullURL := c.instance
	if !strings.HasPrefix(fullURL, "http") {
		fullURL = "https://" + fullURL + ".service-now.com"
	}
	provider, err := credentials.NewClientCredentialsProvider(clientID, clientSecret, credentials.WithURL(fullURL))
	require.NoError(t, err)

	c.provider = provider
	return nil
}

func (c *sharedTestContext) incorrectClientCredentialsAreSupplied(ctx context.Context) error {
	provider, err := credentials.NewClientCredentialsProvider("invalid-id", "invalid-secret", credentials.WithInstance(c.instance))
	if err == nil {
		c.provider = provider
	}
	return nil
}

func InitializeClientCredentialsAuthenticationScenario(ctx *godog.ScenarioContext) {
	feat := &sharedTestContext{}

	RegisterSharedSteps(ctx, feat)

	ctx.When(`^correct credentials are supplied$`, feat.correctClientCredentialsAreSupplied)
	ctx.When(`^incorrect credentials are supplied$`, feat.incorrectClientCredentialsAreSupplied)
}

func TestClientCredentialsAuthenticationFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeClientCredentialsAuthenticationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/authentication/client_credentials.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
