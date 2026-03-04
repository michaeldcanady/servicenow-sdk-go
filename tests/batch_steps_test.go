//go:build integration

package tests

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type batchTestContext struct {
	client    *sdk.ServiceNowClient
	response  *batchapi.BatchResponseModel
	err       error
	lastSysID string
}

func (c *batchTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../.env")
	return nil
}

func (c *batchTestContext) iHaveInitializedTheServiceNowClient() error {
	instance := os.Getenv("SN_INSTANCE")
	if instance == "" {
		instance = "mock_instance"
	}

	var cred credentials.Credential
	if os.Getenv("SN_USERNAME") != "" {
		cred = credentials.NewUsernamePasswordCredential(
			os.Getenv("SN_USERNAME"),
			os.Getenv("SN_PASSWORD"),
		)
	} else {
		cred = credentials.NewUsernamePasswordCredential("mock", "mock")
	}

	client, err := sdk.NewServiceNowClient2WithHTTPClient(cred, instance, getHttpClient())
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

func (c *batchTestContext) iSendABatchRequestWithAPOSTToAndAGETFor(postTable, getTable string) error {
	batchReq := batchapi.NewBatchRequestModel()

	// POST request
	postReq := batchapi.NewRestRequest()
	id1 := "1"
	postMethod := abstractions.POST
	postUrl := "/api/now/v1/table/" + postTable

	_ = postReq.SetID(&id1)
	_ = postReq.SetMethod(&postMethod)
	_ = postReq.SetURL(&postUrl)

	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", "Batch POST test")
	if err := postReq.SetBodyFromParsable("application/json", record); err != nil {
		return fmt.Errorf("failed to set body: %w", err)
	}

	if err := batchReq.AddRequest(postReq); err != nil {
		return fmt.Errorf("failed to add post request: %w", err)
	}

	// GET request
	getReq := batchapi.NewRestRequest()
	id2 := "2"
	getMethod := abstractions.GET
	getUrl := "/api/now/v1/table/" + getTable + "?sysparm_limit=1"

	_ = getReq.SetID(&id2)
	_ = getReq.SetMethod(&getMethod)
	_ = getReq.SetURL(&getUrl)

	if err := batchReq.AddRequest(getReq); err != nil {
		return fmt.Errorf("failed to add get request: %w", err)
	}

	resp, err := c.client.Now2().Batch().Post(context.Background(), batchReq, nil)
	c.response = resp
	c.err = err
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

	if *statusCode < 200 || *statusCode >= 300 {
		return fmt.Errorf("expected successful status code, got %d", *statusCode)
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
		if *statusCode < 200 || *statusCode >= 300 {
			return fmt.Errorf("expected successful status code for request %d, got %d", i+1, *statusCode)
		}
	}
	return nil
}

func (c *batchTestContext) iCreateANewIncidentWithDescription(description string) error {
	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", description)

	resp, err := c.client.Now2().TableV2("incident").Post(context.Background(), record, nil)
	if err != nil {
		return err
	}
	result, _ := resp.GetResult()
	sysID, _ := result.GetSysID()
	c.lastSysID = *sysID
	return nil
}

func (c *batchTestContext) iSendABatchRequestWithADELETEOperationForTheCreatedIncident() error {
	batchReq := batchapi.NewBatchRequestModel()
	req := batchapi.NewRestRequest()

	id := "1"
	method := abstractions.DELETE
	url := "/api/now/v1/table/incident/" + c.lastSysID

	_ = req.SetID(&id)
	_ = req.SetMethod(&method)
	_ = req.SetURL(&url)

	if err := batchReq.AddRequest(req); err != nil {
		return fmt.Errorf("failed to add delete request: %w", err)
	}

	resp, err := c.client.Now2().Batch().Post(context.Background(), batchReq, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *batchTestContext) iRequestTheDeletedIncidentByItsSysID() error {
	if isOffline() {
		baseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/table/incident/", os.Getenv("SN_INSTANCE"))
		httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(baseURL+`[a-zA-Z0-9_]+$`),
			httpmock.NewStringResponder(404, `{"error":{"message":"No Record found","detail":""},"status":"failure"}`))
	}
	_, err := c.client.Now2().TableV2("incident").ById(c.lastSysID).Get(context.Background(), nil)
	c.err = err
	return nil
}

func (c *batchTestContext) theResponseShouldBeA404Error() error {
	if c.err == nil {
		return fmt.Errorf("expected a 404 error, but got no error")
	}
	return nil
}

func InitializeBatchScenario(ctx *godog.ScenarioContext) {
	tc := &batchTestContext{}

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
	ctx.Step(`^I send a batch request with a GET operation for "([^"]*)" table$`, tc.iSendABatchRequestWithAGETOperationForTable)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the batch response should contain a successful result for the operation$`, tc.theBatchResponseShouldContainASuccessfulResultForTheOperation)

	ctx.Step(`^I send a batch request with GET operations for "([^"]*)" and "([^"]*)" tables$`, tc.iSendABatchRequestWithGETOperationsForAndTables)
	ctx.Step(`^the batch response should contain (\d+) successful results$`, tc.theBatchResponseShouldContainSuccessfulResults)

	ctx.Step(`^I send a batch request with a POST to "([^"]*)" and a GET for "([^"]*)"$`, tc.iSendABatchRequestWithAPOSTToAndAGETFor)

	ctx.Step(`^I create a new incident with description "([^"]*)"$`, tc.iCreateANewIncidentWithDescription)
	ctx.Step(`^I send a batch request with a DELETE operation for the created incident$`, tc.iSendABatchRequestWithADELETEOperationForTheCreatedIncident)
	ctx.Step(`^I request the deleted incident by its "sys_id"$`, tc.iRequestTheDeletedIncidentByItsSysID)
	ctx.Step(`^the response should be a 404 error$`, tc.theResponseShouldBeA404Error)
}

func TestBatchFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeBatchScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/batch_api.feature", "features/batch_complex.feature", "features/batch_multi_method.feature", "features/batch_delete.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
