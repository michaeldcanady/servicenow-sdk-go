// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/caseapi"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/mockdata"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
)

type caseSteps struct{}

func (s *caseSteps) iListAllCases(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.CustomerService().Case().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	colResp, ok := resp.(caseapi.CaseCollectionResponse)
	if ok {
		results, err := colResp.GetResult()
		if err == nil && len(results) > 0 {
			sysID, err := results[0].GetSysID()
			if err == nil && sysID != nil && *sysID != "" {
				w.LastSysID = *sysID
			}
		}
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *caseSteps) iCreateACaseWithShortDescriptionAndDescription(ctx context.Context, shortDesc, description string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	record := caseapi.NewCaseResult()
	_ = record.SetShortDescription(&shortDesc)
	_ = record.SetDescription(&description)

	resp, err := w.Client.CustomerService().Case().Post(ctx, record, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	itemResp, ok := resp.(caseapi.CaseItemResponse)
	if !ok {
		w.Err = fmt.Errorf("unexpected response type: %T", resp)
		return support.WithWorld(ctx, w), nil
	}

	result, err := itemResp.GetResult()
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	sysID, err := result.GetSysID()
	if err == nil && sysID != nil && *sysID != "" {
		w.LastSysID = *sysID
		w.TrackCreation(*sysID)
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *caseSteps) iRetrieveTheCaseBySysID(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_case_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no case sys_id available — create a case first")
		}
	}

	resp, err := w.Client.CustomerService().Case().ByID(sysID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *caseSteps) iRequestTheCaseWithSysID(ctx context.Context, sysID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	if support.IsOffline() {
		instance := support.IntegrationInstance()
		url := fmt.Sprintf("https://%s.service-now.com/api/sn_customerservice/v1/case/%s", instance, sysID)
		httpmock.RegisterResponder("GET", url,
			httpmock.NewStringResponder(404, mockdata.SomeErrorJSON))
	}

	resp, err := w.Client.CustomerService().Case().ByID(sysID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *caseSteps) iUpdateLastCaseWithPUTShortDescription(ctx context.Context, shortDesc string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_case_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no case sys_id available — create a case first")
		}
	}

	record := caseapi.NewCaseResult()
	_ = record.SetShortDescription(&shortDesc)

	resp, err := w.Client.CustomerService().Case().ByID(sysID).Put(ctx, record, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *caseSteps) iRetrieveActivitiesForTheLastCase(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_case_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no case sys_id available — list or create a case first")
		}
	}

	resp, err := w.Client.CustomerService().Case().ByID(sysID).Activities().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *caseSteps) iRetrieveFieldValuesForTheLastCase(ctx context.Context, fieldName string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_case_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no case sys_id available — list or create a case first")
		}
	}

	resp, err := w.Client.CustomerService().Case().ByID(sysID).FieldValues(fieldName).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *caseSteps) theCaseCollectionShouldContainAtLeast(ctx context.Context, minCount int) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil, expected at least %d cases", minCount)
	}

	colResp, ok := w.Response.(caseapi.CaseCollectionResponse)
	if !ok {
		return ctx, fmt.Errorf("response is not a CaseCollectionResponse: %T", w.Response)
	}

	results, err := colResp.GetResult()
	if err != nil {
		return ctx, fmt.Errorf("failed to get results: %v", err)
	}

	if len(results) < minCount {
		return ctx, fmt.Errorf("expected at least %d cases, got %d", minCount, len(results))
	}

	return ctx, nil
}

func (s *caseSteps) theCreatedCaseShouldHaveAValidSysID(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil")
	}

	itemResp, ok := w.Response.(caseapi.CaseItemResponse)
	if !ok {
		return ctx, fmt.Errorf("response is not a CaseItemResponse: %T", w.Response)
	}

	result, err := itemResp.GetResult()
	if err != nil {
		return ctx, fmt.Errorf("failed to get result: %v", err)
	}

	sysID, err := result.GetSysID()
	if err != nil {
		return ctx, fmt.Errorf("failed to get sys_id: %v", err)
	}

	if sysID == nil || *sysID == "" {
		return ctx, fmt.Errorf("created case has nil or empty sys_id")
	}

	return ctx, nil
}

func (s *caseSteps) theRetrievedCaseShouldHaveShortDescription(ctx context.Context, expected string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil")
	}

	itemResp, ok := w.Response.(caseapi.CaseItemResponse)
	if !ok {
		return ctx, fmt.Errorf("response is not a CaseItemResponse: %T", w.Response)
	}

	result, err := itemResp.GetResult()
	if err != nil {
		return ctx, fmt.Errorf("failed to get result: %v", err)
	}

	shortDesc, err := result.GetShortDescription()
	if err != nil {
		return ctx, fmt.Errorf("failed to get short_description: %v", err)
	}

	if shortDesc == nil || *shortDesc != expected {
		return ctx, fmt.Errorf("expected short_description %q, got %v", expected, shortDesc)
	}

	return ctx, nil
}

// InitializeCaseScenario registers all case step definitions.
func InitializeCaseScenario(sc *godog.ScenarioContext) {
	s := &caseSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I list all cases$`, s.iListAllCases)
	sc.Step(`^I create a case with short description "([^"]*)" and description "([^"]*)"$`, s.iCreateACaseWithShortDescriptionAndDescription)
	sc.Step(`^I retrieve the case by its sys_id$`, s.iRetrieveTheCaseBySysID)
	sc.Step(`^I request the case with sys_id "([^"]*)"$`, s.iRequestTheCaseWithSysID)
	sc.Step(`^I update the last case with PUT short description "([^"]*)"$`, s.iUpdateLastCaseWithPUTShortDescription)
	sc.Step(`^I retrieve activities for the last case$`, s.iRetrieveActivitiesForTheLastCase)
	sc.Step(`^I retrieve field values "([^"]*)" for the last case$`, s.iRetrieveFieldValuesForTheLastCase)
	sc.Step(`^the case collection should contain at least (\d+) case$`, s.theCaseCollectionShouldContainAtLeast)
	sc.Step(`^the created case should have a valid sys_id$`, s.theCreatedCaseShouldHaveAValidSysID)
	sc.Step(`^the retrieved case should have short description "([^"]*)"$`, s.theRetrievedCaseShouldHaveShortDescription)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
