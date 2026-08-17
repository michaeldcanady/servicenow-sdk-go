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
	aggregationapi "github.com/michaeldcanady/servicenow-sdk-go/aggregationapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type aggregationTestContext struct {
	client   *sdk.ServiceNowServiceClient
	response core.ServiceNowItemResponse[*aggregationapi.StatsResultModel]
	err      error
}

func (c *aggregationTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../../.env")
	return nil
}

func (c *aggregationTestContext) iHaveInitializedTheServiceNowClient() error {
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

func (c *aggregationTestContext) iRequestTheRecordCountForTheTable(tableName string) error {
	count := true
	cfg := &aggregationapi.StatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &aggregationapi.StatsRequestBuilderGetQueryParameters{
			Count: &count,
		},
	}
	resp, err := c.client.Now().Stats(tableName).Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *aggregationTestContext) iRequestTheRecordCountForTheTableWithQuery(tableName, query string) error {
	count := true
	cfg := &aggregationapi.StatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &aggregationapi.StatsRequestBuilderGetQueryParameters{
			Count: &count,
			Query: internal.ToPointer(query),
		},
	}
	resp, err := c.client.Now().Stats(tableName).Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *aggregationTestContext) iRequestTheSumOfForTheTable(fieldName, tableName string) error {
	count := true
	cfg := &aggregationapi.StatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &aggregationapi.StatsRequestBuilderGetQueryParameters{
			Count:     &count,
			SumFields: []string{fieldName},
		},
	}
	resp, err := c.client.Now().Stats(tableName).Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *aggregationTestContext) iRequestStatsForTheTableWithNoAggregateParameters(tableName string) error {
	resp, err := c.client.Now().Stats(tableName).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *aggregationTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		if snErr, ok := c.err.(*core.ServiceNowError); ok {
			mainErr, err := snErr.GetError()
			if err != nil {
				return fmt.Errorf("failed to retrieve error details: %w", err)
			}

			msg, err := mainErr.GetMessage()
			if err != nil {
				return fmt.Errorf("failed to retrieve error message: %w", err)
			}

			detail, err := mainErr.GetDetail()
			if err != nil {
				return fmt.Errorf("failed to retrieve error detail: %w", err)
			}

			status, err := mainErr.GetStatus()
			if err != nil {
				return fmt.Errorf("failed to retrieve error status: %w", err)
			}

			return fmt.Errorf("expected no error, but got: %v (Message: %s, Detail: %s, Status: %s)",
				c.err, derefOrEmpty(msg), derefOrEmpty(detail), derefOrEmpty(status))
		}
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *aggregationTestContext) theResponseShouldBeAnAPIError() error {
	if c.err == nil {
		return fmt.Errorf("expected an API error, but got no error")
	}
	// Auth/token failures are setup problems, not the endpoint behavior under test.
	msg := c.err.Error()
	if strings.Contains(msg, "token acquisition") || strings.Contains(msg, "oauth2") {
		return fmt.Errorf("expected a ServiceNow API error, but auth failed first: %v", c.err)
	}
	return nil
}

func (c *aggregationTestContext) stats() (aggregationapi.Stats, error) {
	if c.response == nil {
		return nil, fmt.Errorf("expected a stats response, but got nil")
	}
	result, err := c.response.GetResult()
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("expected a stats result, but got nil")
	}
	return result.GetStats()
}

func (c *aggregationTestContext) theStatsCountShouldBePresent() error {
	stats, err := c.stats()
	if err != nil {
		return err
	}
	if stats == nil {
		return fmt.Errorf("expected stats payload, but got nil")
	}
	count, err := stats.GetCount()
	if err != nil {
		return err
	}
	if count == nil || *count == "" {
		return fmt.Errorf("expected a stats count, but it was absent")
	}
	return nil
}

func (c *aggregationTestContext) theStatsCountShouldBe(expected string) error {
	stats, err := c.stats()
	if err != nil {
		return err
	}
	if stats == nil {
		return fmt.Errorf("expected stats payload, but got nil")
	}
	count, err := stats.GetCount()
	if err != nil {
		return err
	}
	if count == nil {
		return fmt.Errorf("expected stats count %q, but count was nil", expected)
	}
	if *count != expected {
		return fmt.Errorf("expected stats count %q, got %q", expected, *count)
	}
	return nil
}

func (c *aggregationTestContext) theStatsCountShouldBeAbsent() error {
	stats, err := c.stats()
	if err != nil {
		return err
	}
	if stats == nil {
		return nil
	}
	count, err := stats.GetCount()
	if err != nil {
		return err
	}
	if count != nil {
		return fmt.Errorf("expected stats count to be absent, got %q", *count)
	}
	return nil
}

func (c *aggregationTestContext) theStatsSumForShouldBePresent(fieldName string) error {
	stats, err := c.stats()
	if err != nil {
		return err
	}
	if stats == nil {
		return fmt.Errorf("expected stats payload, but got nil")
	}
	sum, err := stats.GetSum()
	if err != nil {
		return err
	}
	if sum == nil {
		return fmt.Errorf("expected a stats sum map, but got nil")
	}
	if _, ok := sum[fieldName]; !ok {
		return fmt.Errorf("expected sum entry for %q, keys=%v", fieldName, keysOf(sum))
	}
	return nil
}

func derefOrEmpty(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func keysOf(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func InitializeAggregationScenario(ctx *godog.ScenarioContext) {
	tc := &aggregationTestContext{}

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
	ctx.Step(`^I request the record count for the "([^"]*)" table$`, tc.iRequestTheRecordCountForTheTable)
	ctx.Step(`^I request the record count for the "([^"]*)" table with query "([^"]*)"$`, tc.iRequestTheRecordCountForTheTableWithQuery)
	ctx.Step(`^I request the sum of "([^"]*)" for the "([^"]*)" table$`, tc.iRequestTheSumOfForTheTable)
	ctx.Step(`^I request stats for the "([^"]*)" table with no aggregate parameters$`, tc.iRequestStatsForTheTableWithNoAggregateParameters)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the response should be an API error$`, tc.theResponseShouldBeAnAPIError)
	ctx.Step(`^the stats count should be present$`, tc.theStatsCountShouldBePresent)
	ctx.Step(`^the stats count should be "([^"]*)"$`, tc.theStatsCountShouldBe)
	ctx.Step(`^the stats count should be absent$`, tc.theStatsCountShouldBeAbsent)
	ctx.Step(`^the stats sum for "([^"]*)" should be present$`, tc.theStatsSumForShouldBePresent)
}

func TestAggregationFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeAggregationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/stats.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
