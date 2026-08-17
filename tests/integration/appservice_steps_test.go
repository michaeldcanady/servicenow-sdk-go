//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	appserviceapi "github.com/michaeldcanady/servicenow-sdk-go/appserviceapi"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type appServiceTestContext struct {
	client    *sdk.ServiceNowServiceClient
	response  interface{}
	err       error
	lastSysID string
}

func (c *appServiceTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../../.env")
	return nil
}

func (c *appServiceTestContext) iHaveInitializedTheServiceNowClient() error {
	instance := integrationInstance()

	cred, err := newIntegrationAuthProvider(instance)
	if err != nil {
		return err
	}

	opts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(cred),
		sdk.WithInstance(instance),
	}
	if httpClient := getHttpClient(); httpClient != nil {
		opts = append(opts, sdk.WithHTTPClient(httpClient))
	}

	client, err := sdk.NewServiceNowServiceClient(opts...)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *appServiceTestContext) iCreateAnApplicationServiceNamedWithComments(name, comments string) error {
	req := appserviceapi.NewCreateServiceRequest()
	req.SetName(&name)
	req.SetComments(&comments)

	resp, err := c.client.Now().Cmdb().AppService().Create().Post(context.Background(), req, nil)
	c.response = resp
	c.err = err

	if err == nil && resp != nil {
		result, err := resp.GetResult()
		if err == nil && result != nil {
			sysID, err := result.GetSysID()
			if err == nil && sysID != nil {
				c.lastSysID = *sysID
			}
		}
	}
	return nil
}

func (c *appServiceTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *appServiceTestContext) theCreatedServiceShouldHaveAValidSysID() error {
	resp, ok := c.response.(appserviceapi.CreateServiceResponse)
	if !ok || resp == nil {
		return fmt.Errorf("expected CreateServiceResponse, got %T", c.response)
	}
	result, err := resp.GetResult()
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("expected result inside CreateServiceResponse, got nil")
	}
	sysID, err := result.GetSysID()
	if err != nil {
		return err
	}
	if sysID == nil || *sysID == "" {
		return fmt.Errorf("expected a valid sys_id, got empty/nil")
	}
	c.lastSysID = *sysID
	return nil
}

func (c *appServiceTestContext) iRequestTheContentOfTheApplicationService() error {
	if c.lastSysID == "" {
		return fmt.Errorf("no application service sys_id set")
	}
	resp, err := c.client.Now().Cmdb().AppService().ByID(c.lastSysID).GetContent().Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *appServiceTestContext) theContentResponseShouldNotBeNil() error {
	resp, ok := c.response.(appserviceapi.GetContentResponse)
	if !ok || resp == nil {
		return fmt.Errorf("expected GetContentResponse, got %T", c.response)
	}
	result, err := resp.GetResult()
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("expected result inside GetContentResponse, got nil")
	}
	return nil
}

func (c *appServiceTestContext) iSearchForAnApplicationServiceByName(name string) error {
	cfg := &appserviceapi.FindServiceRequestConfiguration{
		QueryParameters: &appserviceapi.FindServiceQueryParameters{
			Name: internal.ToPointer(name),
		},
	}
	resp, err := c.client.Now().Cmdb().AppService().Csdm().FindService().Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *appServiceTestContext) theSearchResultsShouldContainAtLeastServiceNamed(minCount int, expectedName string) error {
	resp, ok := c.response.(appserviceapi.FindServiceResponse)
	if !ok || resp == nil {
		return fmt.Errorf("expected FindServiceResponse, got %T", c.response)
	}
	result, err := resp.GetResult()
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("expected result inside FindServiceResponse, got nil")
	}
	services, err := result.GetServices()
	if err != nil {
		return err
	}
	if len(services) < minCount {
		return fmt.Errorf("expected at least %d services, got %d", minCount, len(services))
	}
	found := false
	for _, s := range services {
		name, err := s.GetName()
		if err == nil && name != nil && *name == expectedName {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("could not find service named %q in search results", expectedName)
	}
	return nil
}

func InitializeAppServiceScenario(ctx *godog.ScenarioContext) {
	tc := &appServiceTestContext{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		setupGlobalMocks()
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		httpmock.DeactivateAndReset()
		return ctx, nil
	})

	ctx.Step(`^I have a valid ServiceNow instance and credentials$`, tc.iHaveAValidServiceNowInstanceAndCredentials)
	ctx.Step(`^I have initialized the ServiceNow client$`, tc.iHaveInitializedTheServiceNowClient)
	ctx.Step(`^I create an application service named "([^"]*)" with comments "([^"]*)"$`, tc.iCreateAnApplicationServiceNamedWithComments)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the created service should have a valid sys_id$`, tc.theCreatedServiceShouldHaveAValidSysID)
	ctx.Step(`^I request the content of the application service$`, tc.iRequestTheContentOfTheApplicationService)
	ctx.Step(`^the content response should not be nil$`, tc.theContentResponseShouldNotBeNil)
	ctx.Step(`^I search for an application service by name "([^"]*)"$`, tc.iSearchForAnApplicationServiceByName)
	ctx.Step(`^the search results should contain at least (\d+) service named "([^"]*)"$`, tc.theSearchResultsShouldContainAtLeastServiceNamed)
}

func TestAppServiceFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeAppServiceScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/appservice.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
