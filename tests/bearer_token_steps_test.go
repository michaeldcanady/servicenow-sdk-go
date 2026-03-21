package tests

import (
	"context"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/require"
)

func (c *sharedTestContext) correctBearerTokenAreSupplied(ctx context.Context) error {
	t := godog.T(ctx)

	token, ok := os.LookupEnv("SN_BEARER_TOKEN")
	require.True(t, ok)
	require.NotEmpty(t, token)

	c.provider = credentials.NewBearerTokenAuthenticationProvider(NewStaticTokenProvider(token))
	return nil
}

func (c *sharedTestContext) incorrectBearerTokenAreSupplied(ctx context.Context) error {
	c.provider = credentials.NewBearerTokenAuthenticationProvider(NewStaticTokenProvider("invalid-token"))
	return nil
}

func InitializeBearerTokenAuthenticationScenario(ctx *godog.ScenarioContext) {
	feat := &sharedTestContext{}

	RegisterSharedSteps(ctx, feat)

	ctx.When(`^correct credentials are supplied$`, feat.correctBearerTokenAreSupplied)
	ctx.When(`^incorrect credentials are supplied$`, feat.incorrectBearerTokenAreSupplied)
}

func TestBearerTokenAuthenticationFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeBearerTokenAuthenticationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/authentication/bearer_token.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
