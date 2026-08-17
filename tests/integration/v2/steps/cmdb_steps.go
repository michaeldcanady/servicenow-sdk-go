package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/cmdbinstanceapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
)

type cmdbSteps struct{}

func (s *cmdbSteps) iRequestCIsForClassWithLimit(ctx context.Context, className string, limit int) (context.Context, error) {
	w := support.WorldFrom(ctx)

	l := int32(limit)
	cfg := &cmdbinstanceapi.CmdbClassRequestBuilderGetRequestConfiguration{
		QueryParameters: &cmdbinstanceapi.CmdbClassRequestBuilderGetQueryParameters{
			Limit: &l,
		},
	}

	resp, err := w.Client.Now().Cmdb().Instance().ByClass(className).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *cmdbSteps) iRequestCIsForClassWithQuery(ctx context.Context, className, query string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	cfg := &cmdbinstanceapi.CmdbClassRequestBuilderGetRequestConfiguration{
		QueryParameters: &cmdbinstanceapi.CmdbClassRequestBuilderGetQueryParameters{
			Query: &query,
		},
	}

	resp, err := w.Client.Now().Cmdb().Instance().ByClass(className).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *cmdbSteps) iHaveAtLeastCIInClass(ctx context.Context, count int, className string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	l := int32(count)
	cfg := &cmdbinstanceapi.CmdbClassRequestBuilderGetRequestConfiguration{
		QueryParameters: &cmdbinstanceapi.CmdbClassRequestBuilderGetQueryParameters{
			Limit: &l,
		},
	}

	resp, err := w.Client.Now().Cmdb().Instance().ByClass(className).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	results, err := resp.GetResult()
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	if len(results) < count {
		w.Err = fmt.Errorf("expected at least %d CIs, got %d", count, len(results))
		return support.WithWorld(ctx, w), nil
	}

	sysID, err := results[0].GetSysID()
	if err == nil && sysID != nil {
		w.LastSysID = *sysID
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *cmdbSteps) iRequestTheCIBySysIDInClass(ctx context.Context, field, className string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_cmdb_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no CI sys_id available — list CIs first")
		}
	}

	resp, err := w.Client.Now().Cmdb().Instance().ByClass(className).ByID(sysID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func ptrStr(s string) *string { return &s }

func (s *cmdbSteps) iCreateACIForClass(ctx context.Context, className string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	ci := cmdbinstanceapi.NewCmdbInstance()
	_ = ci.SetName(ptrStr("BDD Test CI"))
	_ = ci.SetClassName(ptrStr(className))

	resp, err := w.Client.Now().Cmdb().Instance().ByClass(className).Post(ctx, ci, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	result, err := resp.GetResult()
	if err == nil {
		sysID, err := result.GetSysID()
		if err == nil && sysID != nil && *sysID != "" {
			w.LastSysID = *sysID
		}
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *cmdbSteps) iUpdateLastCIForClassWithName(ctx context.Context, className, name string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_cmdb_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no CI sys_id available — create a CI first")
		}
	}

	ci := cmdbinstanceapi.NewCmdbInstance()
	_ = ci.SetName(ptrStr(name))

	resp, err := w.Client.Now().Cmdb().Instance().ByClass(className).ByID(sysID).Patch(ctx, ci, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *cmdbSteps) theCmdbResultsShouldContainAtLeastRecords(ctx context.Context, minCount int) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil, expected at least %d records", minCount)
	}

	colResp, ok := w.Response.(*core.BaseServiceNowCollectionResponse[cmdbinstanceapi.CmdbInstance])
	if !ok {
		return ctx, fmt.Errorf("response is not a CmdbCollectionResponse: %T", w.Response)
	}

	results, err := colResp.GetResult()
	if err != nil {
		return ctx, fmt.Errorf("failed to get results: %v", err)
	}

	if len(results) < minCount {
		return ctx, fmt.Errorf("expected at least %d records, got %d", minCount, len(results))
	}

	return ctx, nil
}

func (s *cmdbSteps) theCmdbResultShouldHaveTheCorrectSysID(ctx context.Context, field string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil")
	}

	itemResp, ok := w.Response.(*core.BaseServiceNowItemResponse[cmdbinstanceapi.CmdbInstance])
	if !ok {
		return ctx, fmt.Errorf("response is not a CmdbItemResponse: %T", w.Response)
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
		return ctx, fmt.Errorf("sys_id is nil or empty")
	}

	return ctx, nil
}

// InitializeCmdbScenario registers all CMDB step definitions.
func InitializeCmdbScenario(sc *godog.ScenarioContext) {
	s := &cmdbSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I request CIs for class "([^"]*)" with limit (\d+)$`, s.iRequestCIsForClassWithLimit)
	sc.Step(`^I request CIs for class "([^"]*)" with query "([^"]*)"$`, s.iRequestCIsForClassWithQuery)
	sc.Step(`^I have at least (\d+) CI in class "([^"]*)"$`, s.iHaveAtLeastCIInClass)
	sc.Step(`^I request the CI by its "([^"]*)" in class "([^"]*)"$`, s.iRequestTheCIBySysIDInClass)
	sc.Step(`^I create a CI for class "([^"]*)"$`, s.iCreateACIForClass)
	sc.Step(`^I update the last CI for class "([^"]*)" with name "([^"]*)"$`, s.iUpdateLastCIForClassWithName)
	sc.Step(`^the cmdb results should contain at least (\d+) record$`, s.theCmdbResultsShouldContainAtLeastRecords)
	sc.Step(`^the cmdb result should have the correct "([^"]*)"$`, s.theCmdbResultShouldHaveTheCorrectSysID)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
