package tests

import (
	"context"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/require"
)

func (c *sharedTestContext) correctBasicCredentialsAreSupplied(ctx context.Context) error {
	t := godog.T(ctx)

	username, ok := os.LookupEnv("SN_USERNAME")
	require.True(t, ok)
	require.NotEmpty(t, username)

	password, ok := os.LookupEnv("SN_PASSWORD")
	require.True(t, ok)
	require.NotEmpty(t, password)

	c.provider = credentials.NewBasicProvider(username, password)
	return nil
}

func (c *sharedTestContext) incorrectBasicCredentialsAreSupplied(ctx context.Context) error {
	c.provider = credentials.NewBasicProvider("invalid-user", "invalid-pass")
	return nil
}

func InitializeBasicAuthenticationScenario(ctx *godog.ScenarioContext) {
	feat := &sharedTestContext{}

	RegisterSharedSteps(ctx, feat)

	ctx.When(`^correct credentials are supplied$`, feat.correctBasicCredentialsAreSupplied)
	ctx.When(`^incorrect credentials are supplied$`, feat.incorrectBasicCredentialsAreSupplied)
}

func TestBasicAuthenticationFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeBasicAuthenticationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/authentication/basic_authentication.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
