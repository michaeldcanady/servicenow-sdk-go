package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type batchTestContext struct {
	client   *sdk.ServiceNowClient
	response *batchapi.BatchResponseModel
	err      error
}

func (c *batchTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../.env")
	required := []string{"SN_CLIENT_ID", "SN_USERNAME", "SN_PASSWORD", "SN_INSTANCE"}
	for _, env := range required {
		if os.Getenv(env) == "" {
			return fmt.Errorf("missing environment variable: %s", env)
		}
	}
	return nil
}

func (c *batchTestContext) iHaveInitializedTheServiceNowClient() error {
	instance := os.Getenv("SN_INSTANCE")
	authority := credentials.NewInstanceAuthority(instance)
	cred, err := credentials.NewROPCCredential(
		os.Getenv("SN_CLIENT_ID"),
		os.Getenv("SN_CLIENT_SECRET"),
		os.Getenv("SN_USERNAME"),
		os.Getenv("SN_PASSWORD"),
		authority,
		nil,
	)
	if err != nil {
		return err
	}

	client, err := sdk.NewServiceNowClient2(cred, instance)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *batchTestContext) iSendABatchRequestWithAGETOperationForTable(tableName string) error {
	batchReq := batchapi.NewBatchRequestModel()
	req := batchapi.NewRestRequest()

	id := "1"
	method := abstractions.GET
	url := "/api/now/v1/table/" + tableName + "?sysparm_limit=1"

	if err := req.SetID(&id); err != nil {
		return fmt.Errorf("failed to set id: %w", err)
	}
	if err := req.SetMethod(&method); err != nil {
		return fmt.Errorf("failed to set method: %w", err)
	}
	if err := req.SetURL(&url); err != nil {
		return fmt.Errorf("failed to set url: %w", err)
	}

	if err := batchReq.AddRequest(req); err != nil {
		return fmt.Errorf("failed to add request: %w", err)
	}

	resp, err := c.client.Now2().Batch().Post(context.Background(), batchReq, nil)
	c.response = resp
	if err != nil {
		c.err = fmt.Errorf("failed to send post request: %w", err)
	}
	return nil
}

func (c *batchTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *batchTestContext) theBatchResponseShouldContainASuccessfulResultForTheOperation() error {
	if c.response == nil {
		return fmt.Errorf("expected a batch response, but got nil")
	}
	servicedRequests, err := c.response.GetServicedRequests()
	if err != nil {
		return fmt.Errorf("failed to retrieve serviced requests: %w", err)
	}
	if len(servicedRequests) == 0 {
		return fmt.Errorf("expected at least 1 serviced request, got 0")
	}

	statusCode, err := servicedRequests[0].GetStatusCode()
	if err != nil {
		return fmt.Errorf("failed to retrieve status code: %w", err)
	}

	if *statusCode != 200 {
		return fmt.Errorf("expected status code 200, got %d", *statusCode)
	}
	return nil
}

func (c *batchTestContext) iSendABatchRequestWithGETOperationsForAndTables(table1, table2 string) error {
	batchReq := batchapi.NewBatchRequestModel()

	for i, tableName := range []string{table1, table2} {
		req := batchapi.NewRestRequest()
		id := fmt.Sprintf("%d", i+1)
		method := abstractions.GET
		url := "/api/now/v1/table/" + tableName + "?sysparm_limit=1"

		_ = req.SetID(&id)
		_ = req.SetMethod(&method)
		_ = req.SetURL(&url)

		if err := batchReq.AddRequest(req); err != nil {
			return fmt.Errorf("failed to add request %d: %w", i+1, err)
		}
	}

	resp, err := c.client.Now2().Batch().Post(context.Background(), batchReq, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *batchTestContext) theBatchResponseShouldContainSuccessfulResults(expectedCount int) error {
	if c.response == nil {
		return fmt.Errorf("expected a batch response, but got nil")
	}
	servicedRequests, err := c.response.GetServicedRequests()
	if err != nil {
		return fmt.Errorf("failed to retrieve serviced requests: %w", err)
	}
	if count := len(servicedRequests); count != expectedCount {
		return fmt.Errorf("expected %d serviced requests, got %d", expectedCount, count)
	}

	for i, req := range servicedRequests {
		statusCode, err := req.GetStatusCode()
		if err != nil {
			return fmt.Errorf("failed to retrieve status code for request %d: %w", i+1, err)
		}
		if *statusCode != 200 {
			return fmt.Errorf("expected status code 200 for request %d, got %d", i+1, *statusCode)
		}
	}
	return nil
}

func InitializeBatchScenario(ctx *godog.ScenarioContext) {
	tc := &batchTestContext{}

	ctx.Step(`^I have a valid ServiceNow instance and credentials$`, tc.iHaveAValidServiceNowInstanceAndCredentials)
	ctx.Step(`^I have initialized the ServiceNow client$`, tc.iHaveInitializedTheServiceNowClient)
	ctx.Step(`^I send a batch request with a GET operation for "([^"]*)" table$`, tc.iSendABatchRequestWithAGETOperationForTable)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the batch response should contain a successful result for the operation$`, tc.theBatchResponseShouldContainASuccessfulResultForTheOperation)

	ctx.Step(`^I send a batch request with GET operations for "([^"]*)" and "([^"]*)" tables$`, tc.iSendABatchRequestWithGETOperationsForAndTables)
	ctx.Step(`^the batch response should contain (\d+) successful results$`, tc.theBatchResponseShouldContainSuccessfulResults)
}

func TestBatchFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeBatchScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/batch_api.feature", "features/batch_complex.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
