//go:build integration

package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

func (c *testContext) iSendABatchRequestWithAGETOperationForTable(tableName string) error {
	batchReq := batchapi.NewBatchRequestModel()
	req := batchapi.NewRestRequest()

	id := "1"
	method := abstractions.GET
	url := "/api/now/v1/table/" + tableName + "?sysparm_limit=1"

	_ = req.SetID(&id)
	_ = req.SetMethod(&method)
	_ = req.SetURL(&url)

	_ = batchReq.AddRequest(req)

	resp, err := c.client.Now().Batch().Post(context.Background(), batchReq, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) iSendABatchRequestWithAPOSTToAndAGETFor(postTable, getTable string) error {
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

	_ = batchReq.AddRequest(postReq)

	// GET request
	getReq := batchapi.NewRestRequest()
	id2 := "2"
	getMethod := abstractions.GET
	getUrl := "/api/now/v1/table/" + getTable + "?sysparm_limit=1"

	_ = getReq.SetID(&id2)
	_ = getReq.SetMethod(&getMethod)
	_ = getReq.SetURL(&getUrl)

	_ = batchReq.AddRequest(getReq)

	resp, err := c.client.Now().Batch().Post(context.Background(), batchReq, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theBatchResponseShouldContainASuccessfulResultForTheOperation() error {
	response, ok := c.response.(*batchapi.BatchResponseModel)
	if !ok {
		return fmt.Errorf("expected a batch response, but got %T", c.response)
	}
	servicedRequests, err := response.GetServicedRequests()
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

func (c *testContext) iSendABatchRequestWithGETOperationsForAndTables(table1, table2 string) error {
	batchReq := batchapi.NewBatchRequestModel()

	for i, tableName := range []string{table1, table2} {
		req := batchapi.NewRestRequest()
		id := fmt.Sprintf("%d", i+1)
		method := abstractions.GET
		url := "/api/now/v1/table/" + tableName + "?sysparm_limit=1"

		_ = req.SetID(&id)
		_ = req.SetMethod(&method)
		_ = req.SetURL(&url)

		_ = batchReq.AddRequest(req)
	}

	resp, err := c.client.Now().Batch().Post(context.Background(), batchReq, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theBatchResponseShouldContainSuccessfulResults(expectedCount int) error {
	response, ok := c.response.(*batchapi.BatchResponseModel)
	if !ok {
		return fmt.Errorf("expected a batch response, but got %T", c.response)
	}
	servicedRequests, err := response.GetServicedRequests()
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

func (c *testContext) iSendABatchRequestWithADELETEOperationForTheCreatedRecord() error {
	batchReq := batchapi.NewBatchRequestModel()
	req := batchapi.NewRestRequest()

	id := "1"
	method := abstractions.DELETE
	url := "/api/now/v1/table/" + c.tableName + "/" + c.lastSysID

	_ = req.SetID(&id)
	_ = req.SetMethod(&method)
	_ = req.SetURL(&url)

	_ = batchReq.AddRequest(req)

	resp, err := c.client.Now().Batch().Post(context.Background(), batchReq, nil)
	c.response = resp
	c.err = err
	return nil
}

func InitializeBatchScenario(ctx *godog.ScenarioContext, tc *testContext) {
	ctx.Step(`^I send a batch request with a GET operation for "([^"]*)" table$`, tc.iSendABatchRequestWithAGETOperationForTable)
	ctx.Step(`^the batch response should contain a successful result for the operation$`, tc.theBatchResponseShouldContainASuccessfulResultForTheOperation)
	ctx.Step(`^I send a batch request with GET operations for "([^"]*)" and "([^"]*)" tables$`, tc.iSendABatchRequestWithGETOperationsForAndTables)
	ctx.Step(`^the batch response should contain (\d+) successful results$`, tc.theBatchResponseShouldContainSuccessfulResults)
	ctx.Step(`^I send a batch request with a POST to "([^"]*)" and a GET for "([^"]*)"$`, tc.iSendABatchRequestWithAPOSTToAndAGETFor)
	ctx.Step(`^I send a batch request with a DELETE operation for the created record$`, tc.iSendABatchRequestWithADELETEOperationForTheCreatedRecord)
}

func TestBatchFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
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
