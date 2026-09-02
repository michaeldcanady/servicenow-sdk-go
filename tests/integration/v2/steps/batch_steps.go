// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/batchapi"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type batchSteps struct{}

func (s *batchSteps) iSendABatchRequestWithAGETOperationForTable(ctx context.Context, table string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	batchReq := batchapi.NewBatchRequestModel()

	restReq := batchapi.NewRestRequest()
	id := "1"
	method := abstractions.GET
	url := fmt.Sprintf("/api/now/v1/table/%s?sysparm_limit=1", table)
	_ = restReq.SetID(&id)
	_ = restReq.SetMethod(&method)
	_ = restReq.SetURL(&url)
	_ = batchReq.AddRequest(restReq)

	resp, err := w.Client.Now().Batch().Post(ctx, batchReq, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *batchSteps) iSendABatchRequestWithGETOperationsForAndTables(ctx context.Context, table1, table2 string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	batchReq := batchapi.NewBatchRequestModel()

	restReq1 := batchapi.NewRestRequest()
	id1 := "1"
	method1 := abstractions.GET
	url1 := fmt.Sprintf("/api/now/v1/table/%s?sysparm_limit=1", table1)
	_ = restReq1.SetID(&id1)
	_ = restReq1.SetMethod(&method1)
	_ = restReq1.SetURL(&url1)
	_ = batchReq.AddRequest(restReq1)

	restReq2 := batchapi.NewRestRequest()
	id2 := "2"
	method2 := abstractions.GET
	url2 := fmt.Sprintf("/api/now/v1/table/%s?sysparm_limit=1", table2)
	_ = restReq2.SetID(&id2)
	_ = restReq2.SetMethod(&method2)
	_ = restReq2.SetURL(&url2)
	_ = batchReq.AddRequest(restReq2)

	resp, err := w.Client.Now().Batch().Post(ctx, batchReq, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *batchSteps) iSendABatchRequestWithAPOSTToAndAGETFor(ctx context.Context, postTable, getTable string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	batchReq := batchapi.NewBatchRequestModel()

	postReq := batchapi.NewRestRequest()
	postID := "1"
	postMethod := abstractions.POST
	postURL := fmt.Sprintf("/api/now/v1/table/%s", postTable)
	_ = postReq.SetID(&postID)
	_ = postReq.SetMethod(&postMethod)
	_ = postReq.SetURL(&postURL)
	_ = postReq.SetBody([]byte(`{"short_description":"Batch test"}`))
	_ = batchReq.AddRequest(postReq)

	getReq := batchapi.NewRestRequest()
	getID := "2"
	getMethod := abstractions.GET
	getURL := fmt.Sprintf("/api/now/v1/table/%s?sysparm_limit=1", getTable)
	_ = getReq.SetID(&getID)
	_ = getReq.SetMethod(&getMethod)
	_ = getReq.SetURL(&getURL)
	_ = batchReq.AddRequest(getReq)

	resp, err := w.Client.Now().Batch().Post(ctx, batchReq, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *batchSteps) iSendABatchRequestWithAnInvalidOperation(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	batchReq := batchapi.NewBatchRequestModel()

	restReq := batchapi.NewRestRequest()
	id := "1"
	method := abstractions.GET
	url := "/api/now/v1/table/invalid_endpoint"
	_ = restReq.SetID(&id)
	_ = restReq.SetMethod(&method)
	_ = restReq.SetURL(&url)
	_ = batchReq.AddRequest(restReq)

	resp, err := w.Client.Now().Batch().Post(ctx, batchReq, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *batchSteps) iSendABatchRequestWithAnEmptyBody(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	batchReq := batchapi.NewBatchRequestModel()

	resp, err := w.Client.Now().Batch().Post(ctx, batchReq, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *batchSteps) theBatchResponseShouldContainAtLeastServicedRequests(ctx context.Context, minCount int) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil, expected at least %d serviced requests", minCount)
	}

	batchResp, ok := w.Response.(*batchapi.BatchResponseModel)
	if !ok {
		return ctx, fmt.Errorf("response is not a *batchapi.BatchResponseModel")
	}

	serviced, err := batchResp.GetServicedRequests()
	if err != nil {
		return ctx, fmt.Errorf("failed to get serviced requests: %v", err)
	}

	if len(serviced) < minCount {
		return ctx, fmt.Errorf("expected at least %d serviced requests, got %d", minCount, len(serviced))
	}

	return ctx, nil
}

func (s *batchSteps) allBatchServicedRequestsShouldHaveSuccessfulStatusCodes(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil")
	}

	batchResp, ok := w.Response.(*batchapi.BatchResponseModel)
	if !ok {
		return ctx, fmt.Errorf("response is not a *batchapi.BatchResponseModel")
	}

	serviced, err := batchResp.GetServicedRequests()
	if err != nil {
		return ctx, fmt.Errorf("failed to get serviced requests: %v", err)
	}

	for i, req := range serviced {
		code, err := req.GetStatusCode()
		if err != nil {
			return ctx, fmt.Errorf("failed to get status code for request %d: %v", i, err)
		}
		if code == nil {
			return ctx, fmt.Errorf("status code is nil for request %d", i)
		}
		if *code < 200 || *code >= 300 {
			return ctx, fmt.Errorf("request %d has non-successful status code: %d", i, *code)
		}
	}

	return ctx, nil
}

// InitializeBatchScenario registers all batch step definitions.
func InitializeBatchScenario(sc *godog.ScenarioContext) {
	s := &batchSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I send a batch request with a GET operation for the "([^"]*)" table$`, s.iSendABatchRequestWithAGETOperationForTable)
	sc.Step(`^I send a batch request with GET operations for the "([^"]*)" and "([^"]*)" tables$`, s.iSendABatchRequestWithGETOperationsForAndTables)
	sc.Step(`^I send a batch request with a POST to "([^"]*)" and a GET for "([^"]*)"$`, s.iSendABatchRequestWithAPOSTToAndAGETFor)
	sc.Step(`^I send a batch request with an invalid operation$`, s.iSendABatchRequestWithAnInvalidOperation)
	sc.Step(`^I send a batch request with an empty body$`, s.iSendABatchRequestWithAnEmptyBody)
	sc.Step(`^the batch response should contain at least (\d+) serviced request$`, s.theBatchResponseShouldContainAtLeastServicedRequests)
	sc.Step(`^all batch serviced requests should have successful status codes$`, s.allBatchServicedRequestsShouldHaveSuccessfulStatusCodes)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
