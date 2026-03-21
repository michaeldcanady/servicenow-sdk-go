package tests

import (
	"context"
	"os"
	"strings"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/stretchr/testify/require"
)

func (c *sharedTestContext) aClientIsAvailable(_ context.Context) error {
	_ = godotenv.Load("../.env")
	c.client = nil
	return nil
}

func (c *sharedTestContext) aValidInstance(ctx context.Context) error {
	t := godog.T(ctx)

	instance, ok := os.LookupEnv("SN_INSTANCE")
	require.True(t, ok)
	require.NotEmpty(t, instance)

	c.instance = instance
	return nil
}

func (c *sharedTestContext) aRequestIsSent(ctx context.Context) error {
	t := godog.T(ctx)

	var instanceOpt servicenowsdkgo.ServiceNowServiceClientOption
	if strings.HasPrefix(c.instance, "http") {
		instanceOpt = servicenowsdkgo.WithURL(c.instance)
	} else {
		instanceOpt = servicenowsdkgo.WithInstance(c.instance)
	}

	opts := []servicenowsdkgo.ServiceNowServiceClientOption{
		servicenowsdkgo.WithAuthenticationProvider(c.provider),
		instanceOpt,
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(opts...)
	require.NoError(t, err)
	require.NotNil(t, client)

	c.client = client

	// Use a small, common request to verify connectivity/auth
	c.resp, c.err = client.Now().Table("incident").Get(context.Background(), nil)

	return nil
}

func (c *sharedTestContext) theResponseShouldBeSuccessful(ctx context.Context) error {
	t := godog.T(ctx)
	require.NoError(t, c.err)
	require.NotNil(t, c.resp)
	return nil
}

func (c *sharedTestContext) anAuthenticationErrorMessageIsShown(ctx context.Context) error {
	t := godog.T(ctx)
	require.Error(t, c.err)
	// You might want to match a specific error message here depending on the auth type
	return nil
}

func RegisterSharedSteps(ctx *godog.ScenarioContext, feat *sharedTestContext) {
	ctx.Given(`^a client is available$`, feat.aClientIsAvailable)
	ctx.Given(`^a valid instance$`, feat.aValidInstance)
	ctx.When(`^a request is sent$`, feat.aRequestIsSent)
	ctx.Then(`^the response should be successful$`, feat.theResponseShouldBeSuccessful)
	ctx.Then(`^an authentication error message is shown$`, feat.anAuthenticationErrorMessageIsShown)
}
