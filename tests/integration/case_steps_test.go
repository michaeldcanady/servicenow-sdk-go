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
	"github.com/michaeldcanady/servicenow-sdk-go/caseapi"
)

type caseTestContext struct {
	client    *sdk.ServiceNowServiceClient
	response  interface{}
	err       error
	lastSysID string
}

func (c *caseTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../../.env")
	return nil
}

func (c *caseTestContext) iHaveInitializedTheServiceNowClient() error {
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

func (c *caseTestContext) iListAllCases() error {
	resp, err := c.client.CustomerService().Case().Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *caseTestContext) theCaseCollectionShouldContainAtLeast(minCount int) error {
	resp, ok := c.response.(caseapi.CaseCollectionResponse)
	if !ok || resp == nil {
		return fmt.Errorf("expected CaseCollectionResponse, got %T", c.response)
	}
	result, err := resp.GetResult()
	if err != nil {
		return err
	}
	if len(result) < minCount {
		return fmt.Errorf("expected at least %d cases, got %d", minCount, len(result))
	}
	return nil
}

func (c *caseTestContext) iCreateACaseWithShortDescriptionAndDescription(shortDesc, desc string) error {
	body := caseapi.NewCaseResult()
	if err := body.SetShortDescription(&shortDesc); err != nil {
		return err
	}
	if err := body.SetDescription(&desc); err != nil {
		return err
	}

	resp, err := c.client.CustomerService().Case().Post(context.Background(), body, nil)
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

func (c *caseTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *caseTestContext) theCreatedCaseShouldHaveAValidSysID() error {
	resp, ok := c.response.(caseapi.CaseItemResponse)
	if !ok || resp == nil {
		return fmt.Errorf("expected CaseItemResponse, got %T", c.response)
	}
	result, err := resp.GetResult()
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("expected result inside CaseItemResponse, got nil")
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

func (c *caseTestContext) iRetrieveTheCaseBySysID() error {
	if c.lastSysID == "" {
		return fmt.Errorf("no case sys_id set")
	}
	resp, err := c.client.CustomerService().Case().ByID(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *caseTestContext) theRetrievedCaseShouldHaveShortDescription(expectedShortDesc string) error {
	resp, ok := c.response.(caseapi.CaseItemResponse)
	if !ok || resp == nil {
		return fmt.Errorf("expected CaseItemResponse, got %T", c.response)
	}
	result, err := resp.GetResult()
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("expected result inside CaseItemResponse, got nil")
	}
	shortDesc, err := result.GetShortDescription()
	if err != nil {
		return err
	}
	if shortDesc == nil || *shortDesc != expectedShortDesc {
		got := "<nil>"
		if shortDesc != nil {
			got = *shortDesc
		}
		return fmt.Errorf("expected short_description %q, got %q", expectedShortDesc, got)
	}
	return nil
}

func (c *caseTestContext) iUpdateTheCaseShortDescriptionTo(newShortDesc string) error {
	if c.lastSysID == "" {
		return fmt.Errorf("no case sys_id set")
	}
	body := caseapi.NewCaseResult()
	if err := body.SetShortDescription(&newShortDesc); err != nil {
		return err
	}

	resp, err := c.client.CustomerService().Case().ByID(c.lastSysID).Put(context.Background(), body, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *caseTestContext) theUpdatedCaseShouldHaveShortDescription(expectedShortDesc string) error {
	if c.lastSysID == "" {
		return fmt.Errorf("no case sys_id set")
	}
	resp, err := c.client.CustomerService().Case().ByID(c.lastSysID).Get(context.Background(), nil)
	if err != nil {
		return err
	}
	c.response = resp
	c.err = err
	return c.theRetrievedCaseShouldHaveShortDescription(expectedShortDesc)
}

func (c *caseTestContext) iRetrieveActivitiesForTheCase() error {
	if c.lastSysID == "" {
		return fmt.Errorf("no case sys_id set")
	}
	resp, err := c.client.CustomerService().Case().ByID(c.lastSysID).Activities().Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *caseTestContext) theActivitiesResponseShouldNotBeNil() error {
	if c.response == nil {
		return fmt.Errorf("expected activities response, got nil")
	}
	resp, ok := c.response.(caseapi.ActivitiesResponse)
	if !ok {
		return fmt.Errorf("expected ActivitiesResponse, got %T", c.response)
	}
	if resp == nil {
		return fmt.Errorf("expected non-nil ActivitiesResponse")
	}
	return nil
}

func (c *caseTestContext) iRetrieveFieldValuesForField(fieldName string) error {
	resp, err := c.client.CustomerService().Case().FieldValues(fieldName).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *caseTestContext) iRetrieveFieldValuesForFieldOnTheCase(fieldName string) error {
	if c.lastSysID == "" {
		return fmt.Errorf("no case sys_id set")
	}
	resp, err := c.client.CustomerService().Case().ByID(c.lastSysID).FieldValues(fieldName).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *caseTestContext) theFieldValuesResponseShouldNotBeNil() error {
	if c.response == nil {
		return fmt.Errorf("expected field values response, got nil")
	}
	resp, ok := c.response.(caseapi.FieldValuesResponse)
	if !ok {
		return fmt.Errorf("expected FieldValuesResponse, got %T", c.response)
	}
	if resp == nil {
		return fmt.Errorf("expected non-nil FieldValuesResponse")
	}
	return nil
}

func InitializeCaseScenario(ctx *godog.ScenarioContext) {
	tc := &caseTestContext{}

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
	ctx.Step(`^I list all cases$`, tc.iListAllCases)
	ctx.Step(`^the case collection should contain at least (\d+) case$`, tc.theCaseCollectionShouldContainAtLeast)
	ctx.Step(`^I create a case with short description "([^"]*)" and description "([^"]*)"$`, tc.iCreateACaseWithShortDescriptionAndDescription)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the created case should have a valid sys_id$`, tc.theCreatedCaseShouldHaveAValidSysID)
	ctx.Step(`^I retrieve the case by its sys_id$`, tc.iRetrieveTheCaseBySysID)
	ctx.Step(`^the retrieved case should have short description "([^"]*)"$`, tc.theRetrievedCaseShouldHaveShortDescription)
	ctx.Step(`^I update the case short description to "([^"]*)"$`, tc.iUpdateTheCaseShortDescriptionTo)
	ctx.Step(`^the updated case should have short description "([^"]*)"$`, tc.theUpdatedCaseShouldHaveShortDescription)
	ctx.Step(`^I retrieve activities for the case$`, tc.iRetrieveActivitiesForTheCase)
	ctx.Step(`^the activities response should not be nil$`, tc.theActivitiesResponseShouldNotBeNil)
	ctx.Step(`^I retrieve field values for the "([^"]*)" field$`, tc.iRetrieveFieldValuesForField)
	ctx.Step(`^I retrieve field values for the "([^"]*)" field on the case$`, tc.iRetrieveFieldValuesForFieldOnTheCase)
	ctx.Step(`^the field values response should not be nil$`, tc.theFieldValuesResponseShouldNotBeNil)
}

func TestCaseFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeCaseScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/case.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
