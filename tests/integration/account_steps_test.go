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
	accountapi "github.com/michaeldcanady/servicenow-sdk-go/accountapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type accountTestContext struct {
	client    *sdk.ServiceNowServiceClient
	response  interface{}
	err       error
	lastSysID string
}

func (c *accountTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../../.env")
	return nil
}

func (c *accountTestContext) iHaveInitializedTheServiceNowClient() error {
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

func (c *accountTestContext) iRequestAllAccounts() error {
	resp, err := c.client.Now().Account().Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *accountTestContext) iRequestAccountsWithLimit(limit int) error {
	lim := int32(limit)
	cfg := &accountapi.AccountRequestBuilderGetRequestConfiguration{
		QueryParameters: &accountapi.AccountRequestBuilderGetQueryParameters{
			Limit: &lim,
		},
	}
	resp, err := c.client.Now().Account().Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *accountTestContext) iRequestAccountsWithQuery(query string) error {
	cfg := &accountapi.AccountRequestBuilderGetRequestConfiguration{
		QueryParameters: &accountapi.AccountRequestBuilderGetQueryParameters{
			Query: internal.ToPointer(query),
		},
	}
	resp, err := c.client.Now().Account().Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *accountTestContext) iHaveAtLeastAccountInTheInstance(minCount int) error {
	lim := int32(minCount)
	cfg := &accountapi.AccountRequestBuilderGetRequestConfiguration{
		QueryParameters: &accountapi.AccountRequestBuilderGetQueryParameters{
			Limit: &lim,
		},
	}
	resp, err := c.client.Now().Account().Get(context.Background(), cfg)
	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("expected an account collection response, but got nil")
	}
	results, err := resp.GetResult()
	if err != nil {
		return err
	}
	if len(results) < minCount {
		return fmt.Errorf("expected at least %d account(s), got %d — seed a customer_account on the instance to cover item GET", minCount, len(results))
	}
	sysID, err := results[0].GetSysID()
	if err != nil {
		return err
	}
	if sysID == nil || *sysID == "" {
		return fmt.Errorf("expected the first account to have a sys_id")
	}
	c.lastSysID = *sysID
	return nil
}

func (c *accountTestContext) iRequestTheAccountByItsSysID() error {
	if c.lastSysID == "" {
		return fmt.Errorf("no account sys_id captured from prior step")
	}
	resp, err := c.client.Now().Account().ByID(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *accountTestContext) iRequestTheAccountWithSysID(sysID string) error {
	if isOffline() {
		baseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/account", integrationInstance())
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", baseURL, sysID),
			httpmock.NewStringResponder(404, `{"error":{"message":"No Record found","detail":""},"status":"failure"}`))
	}
	resp, err := c.client.Now().Account().ByID(sysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *accountTestContext) theResponseShouldNotBeAnError() error {
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

func (c *accountTestContext) theResponseShouldBeAnAPIError() error {
	if c.err == nil {
		return fmt.Errorf("expected an API error, but got no error")
	}
	msg := c.err.Error()
	if strings.Contains(msg, "token acquisition") || strings.Contains(msg, "oauth2") {
		return fmt.Errorf("expected a ServiceNow API error, but auth failed first: %v", c.err)
	}
	return nil
}

func (c *accountTestContext) accountCollection() (accountapi.AccountCollectionResponse, error) {
	collection, ok := c.response.(accountapi.AccountCollectionResponse)
	if !ok || collection == nil {
		return nil, fmt.Errorf("expected an account collection response, but got %T", c.response)
	}
	return collection, nil
}

func (c *accountTestContext) theAccountResultsShouldContainAtLeastRecords(minCount int) error {
	collection, err := c.accountCollection()
	if err != nil {
		return err
	}
	results, err := collection.GetResult()
	if err != nil {
		return err
	}
	if len(results) < minCount {
		return fmt.Errorf("expected at least %d account records, got %d", minCount, len(results))
	}
	return nil
}

func (c *accountTestContext) theAccountResultsShouldContainAtMostRecords(maxCount int) error {
	collection, err := c.accountCollection()
	if err != nil {
		return err
	}
	results, err := collection.GetResult()
	if err != nil {
		return err
	}
	if len(results) > maxCount {
		return fmt.Errorf("expected at most %d account records, got %d", maxCount, len(results))
	}
	return nil
}

func (c *accountTestContext) theAccountResultsShouldContainExactlyRecords(count int) error {
	collection, err := c.accountCollection()
	if err != nil {
		return err
	}
	results, err := collection.GetResult()
	if err != nil {
		return err
	}
	if len(results) != count {
		return fmt.Errorf("expected exactly %d account records, got %d", count, len(results))
	}
	return nil
}

func (c *accountTestContext) theAccountResultShouldHaveTheCorrectSysID() error {
	item, ok := c.response.(accountapi.AccountItemResponse)
	if !ok || item == nil {
		return fmt.Errorf("expected an account item response, but got %T", c.response)
	}
	account, err := item.GetResult()
	if err != nil {
		return err
	}
	if account == nil {
		return fmt.Errorf("expected an account result, but got nil")
	}
	sysID, err := account.GetSysID()
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

func InitializeAccountScenario(ctx *godog.ScenarioContext) {
	tc := &accountTestContext{}

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
	ctx.Step(`^I request all accounts$`, tc.iRequestAllAccounts)
	ctx.Step(`^I request accounts with limit (\d+)$`, tc.iRequestAccountsWithLimit)
	ctx.Step(`^I request accounts with query "([^"]*)"$`, tc.iRequestAccountsWithQuery)
	ctx.Step(`^I have at least (\d+) account in the instance$`, tc.iHaveAtLeastAccountInTheInstance)
	ctx.Step(`^I request the account by its "sys_id"$`, tc.iRequestTheAccountByItsSysID)
	ctx.Step(`^I request the account with sys_id "([^"]*)"$`, tc.iRequestTheAccountWithSysID)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the response should be an API error$`, tc.theResponseShouldBeAnAPIError)
	ctx.Step(`^the account results should contain at least (\d+) records$`, tc.theAccountResultsShouldContainAtLeastRecords)
	ctx.Step(`^the account results should contain at most (\d+) records$`, tc.theAccountResultsShouldContainAtMostRecords)
	ctx.Step(`^the account results should contain exactly (\d+) records$`, tc.theAccountResultsShouldContainExactlyRecords)
	ctx.Step(`^the account result should have the correct "sys_id"$`, tc.theAccountResultShouldHaveTheCorrectSysID)
}

func TestAccountFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeAccountScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/account.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
