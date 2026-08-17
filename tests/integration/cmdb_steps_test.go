//go:build integration

package integration

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	cmdbinstanceapi "github.com/michaeldcanady/servicenow-sdk-go/cmdbinstanceapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type cmdbTestContext struct {
	client    *sdk.ServiceNowServiceClient
	response  interface{}
	err       error
	lastSysID string
	lastClass string
}

func (c *cmdbTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../../.env")
	return nil
}

func (c *cmdbTestContext) iHaveInitializedTheServiceNowClient() error {
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

func (c *cmdbTestContext) iRequestCIsForClassWithLimit(className string, limit int) error {
	lim := int32(limit)
	cfg := &cmdbinstanceapi.CmdbClassRequestBuilderGetRequestConfiguration{
		QueryParameters: &cmdbinstanceapi.CmdbClassRequestBuilderGetQueryParameters{
			Limit: &lim,
		},
	}
	resp, err := c.client.Now().Cmdb().Instance().ByClass(className).Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	c.lastClass = className
	return nil
}

func (c *cmdbTestContext) iRequestCIsForClassWithQuery(className, query string) error {
	cfg := &cmdbinstanceapi.CmdbClassRequestBuilderGetRequestConfiguration{
		QueryParameters: &cmdbinstanceapi.CmdbClassRequestBuilderGetQueryParameters{
			Query: internal.ToPointer(query),
		},
	}
	resp, err := c.client.Now().Cmdb().Instance().ByClass(className).Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	c.lastClass = className
	return nil
}

func (c *cmdbTestContext) iHaveAtLeastCIInClass(minCount int, className string) error {
	lim := int32(minCount)
	cfg := &cmdbinstanceapi.CmdbClassRequestBuilderGetRequestConfiguration{
		QueryParameters: &cmdbinstanceapi.CmdbClassRequestBuilderGetQueryParameters{
			Limit: &lim,
		},
	}
	resp, err := c.client.Now().Cmdb().Instance().ByClass(className).Get(context.Background(), cfg)
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("expected a cmdb collection response, but got nil")
	}
	results, err := resp.GetResult()
	if err != nil {
		return err
	}
	if len(results) < minCount {
		return fmt.Errorf("expected at least %d CI(s) in class %q, got %d", minCount, className, len(results))
	}
	sysID, err := results[0].GetSysID()
	if err != nil {
		return err
	}
	if sysID == nil || *sysID == "" {
		return fmt.Errorf("expected the first CI to have a sys_id")
	}
	c.lastSysID = *sysID
	c.lastClass = className
	return nil
}

func (c *cmdbTestContext) iRequestTheCIByItsSysIDInClass(className string) error {
	if c.lastSysID == "" {
		return fmt.Errorf("no CI sys_id captured from prior step")
	}
	resp, err := c.client.Now().Cmdb().Instance().ByClass(className).ByID(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	c.lastClass = className
	return nil
}

func (c *cmdbTestContext) iRequestTheCIWithSysIDInClass(sysID, className string) error {
	if isOffline() {
		baseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/cmdb/instance/%s", integrationInstance(), className)
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", baseURL, sysID),
			httpmock.NewStringResponder(404, `{"error":{"message":"No Record found","detail":""},"status":"failure"}`))
	}
	resp, err := c.client.Now().Cmdb().Instance().ByClass(className).ByID(sysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	c.lastClass = className
	return nil
}

func (c *cmdbTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		if snErr, ok := c.err.(*core.ServiceNowError); ok {
			mainErr, err := snErr.GetError()
			if err != nil {
				return fmt.Errorf("failed to retrieve error details: %w", err)
			}
			msg, _ := mainErr.GetMessage()
			detail, _ := mainErr.GetDetail()
			status, _ := mainErr.GetStatus()
			return fmt.Errorf("expected no error, but got: %v (Message: %s, Detail: %s, Status: %s)",
				c.err, derefOrEmpty(msg), derefOrEmpty(detail), derefOrEmpty(status))
		}
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *cmdbTestContext) theResponseShouldBeAnAPIError() error {
	if c.err == nil {
		return fmt.Errorf("expected an API error, but got no error")
	}
	msg := c.err.Error()
	if strings.Contains(msg, "token acquisition") || strings.Contains(msg, "oauth2") {
		return fmt.Errorf("expected a ServiceNow API error, but auth failed first: %v", c.err)
	}
	return nil
}

func (c *cmdbTestContext) cmdbCollection() (*core.BaseServiceNowCollectionResponse[cmdbinstanceapi.CmdbInstance], error) {
	collection, ok := c.response.(*core.BaseServiceNowCollectionResponse[cmdbinstanceapi.CmdbInstance])
	if !ok || collection == nil {
		return nil, fmt.Errorf("expected a cmdb collection response, but got %T", c.response)
	}
	return collection, nil
}

func (c *cmdbTestContext) theCmdbResultsShouldContainAtLeastRecords(minCount int) error {
	collection, err := c.cmdbCollection()
	if err != nil {
		return err
	}
	results, err := collection.GetResult()
	if err != nil {
		return err
	}
	if len(results) < minCount {
		return fmt.Errorf("expected at least %d cmdb records, got %d", minCount, len(results))
	}
	return nil
}

func (c *cmdbTestContext) theCmdbResultsShouldContainAtMostRecords(maxCount int) error {
	collection, err := c.cmdbCollection()
	if err != nil {
		return err
	}
	results, err := collection.GetResult()
	if err != nil {
		return err
	}
	if len(results) > maxCount {
		return fmt.Errorf("expected at most %d cmdb records, got %d", maxCount, len(results))
	}
	return nil
}

func (c *cmdbTestContext) theCmdbResultsShouldContainExactlyRecords(count int) error {
	collection, err := c.cmdbCollection()
	if err != nil {
		return err
	}
	results, err := collection.GetResult()
	if err != nil {
		return err
	}
	if len(results) != count {
		return fmt.Errorf("expected exactly %d cmdb records, got %d", count, len(results))
	}
	return nil
}

func (c *cmdbTestContext) theCmdbResultShouldHaveTheCorrectSysID() error {
	item, ok := c.response.(*core.BaseServiceNowItemResponse[cmdbinstanceapi.CmdbInstance])
	if !ok || item == nil {
		return fmt.Errorf("expected a cmdb item response, but got %T", c.response)
	}
	ci, err := item.GetResult()
	if err != nil {
		return err
	}
	if ci == nil {
		return fmt.Errorf("expected a cmdb result, but got nil")
	}
	sysID, err := ci.GetSysID()
	if err != nil {
		return err
	}
	if sysID == nil || *sysID != c.lastSysID {
		got := "<nil>"
		if sysID != nil {
			got = *sysID
		}
		return fmt.Errorf("expected sys_id %q, got %q", c.lastSysID, got)
	}
	return nil
}

func InitializeCmdbScenario(ctx *godog.ScenarioContext) {
	tc := &cmdbTestContext{}

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
	ctx.Step(`^I request CIs for class "([^"]*)" with limit (\d+)$`, tc.iRequestCIsForClassWithLimit)
	ctx.Step(`^I request CIs for class "([^"]*)" with query "([^"]*)"$`, tc.iRequestCIsForClassWithQuery)
	ctx.Step(`^I have at least (\d+) CI in class "([^"]*)"$`, tc.iHaveAtLeastCIInClass)
	ctx.Step(`^I request the CI by its "sys_id" in class "([^"]*)"$`, tc.iRequestTheCIByItsSysIDInClass)
	ctx.Step(`^I request the CI with sys_id "([^"]*)" in class "([^"]*)"$`, tc.iRequestTheCIWithSysIDInClass)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the response should be an API error$`, tc.theResponseShouldBeAnAPIError)
	ctx.Step(`^the cmdb results should contain at least (\d+) records$`, tc.theCmdbResultsShouldContainAtLeastRecords)
	ctx.Step(`^the cmdb results should contain at most (\d+) records$`, tc.theCmdbResultsShouldContainAtMostRecords)
	ctx.Step(`^the cmdb results should contain exactly (\d+) records$`, tc.theCmdbResultsShouldContainExactlyRecords)
	ctx.Step(`^the cmdb result should have the correct "sys_id"$`, tc.theCmdbResultShouldHaveTheCorrectSysID)
}

func TestCmdbFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeCmdbScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/cmdb.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
