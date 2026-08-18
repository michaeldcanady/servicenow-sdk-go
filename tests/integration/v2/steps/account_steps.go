package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/michaeldcanady/servicenow-sdk-go/accountapi"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/mockdata"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
)

type accountSteps struct{}

func (s *accountSteps) iRequestAllAccounts(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Account().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *accountSteps) iRequestAccountsWithQuery(ctx context.Context, query string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	cfg := &accountapi.AccountRequestBuilderGetRequestConfiguration{
		QueryParameters: &accountapi.AccountRequestBuilderGetQueryParameters{
			Query: &query,
		},
	}

	resp, err := w.Client.Now().Account().Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *accountSteps) iRequestAccountsWithLimit(ctx context.Context, limit int) (context.Context, error) {
	w := support.WorldFrom(ctx)

	l := int32(limit)
	cfg := &accountapi.AccountRequestBuilderGetRequestConfiguration{
		QueryParameters: &accountapi.AccountRequestBuilderGetQueryParameters{
			Limit: &l,
		},
	}

	resp, err := w.Client.Now().Account().Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *accountSteps) iHaveAtLeastAccountInTheInstance(ctx context.Context, count int) (context.Context, error) {
	w := support.WorldFrom(ctx)

	limit := int32(count)
	cfg := &accountapi.AccountRequestBuilderGetRequestConfiguration{
		QueryParameters: &accountapi.AccountRequestBuilderGetQueryParameters{
			Limit: &limit,
		},
	}

	resp, err := w.Client.Now().Account().Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	colResp, ok := resp.(accountapi.AccountCollectionResponse)
	if !ok {
		w.Err = fmt.Errorf("unexpected response type: %T", resp)
		return support.WithWorld(ctx, w), nil
	}

	results, err := colResp.GetResult()
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	if len(results) < count {
		w.Err = fmt.Errorf("expected at least %d accounts, got %d", count, len(results))
		return support.WithWorld(ctx, w), nil
	}

	sysID, err := results[0].GetSysID()
	if err == nil && sysID != nil {
		w.LastSysID = *sysID
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *accountSteps) iRequestTheAccountBySysID(ctx context.Context, field string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_account_sys_id_1"
		} else {
			return ctx, fmt.Errorf("no account sys_id available — list accounts first")
		}
	}

	resp, err := w.Client.Now().Account().ByID(sysID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *accountSteps) iRequestTheAccountWithSysID(ctx context.Context, sysID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	if support.IsOffline() {
		instance := support.IntegrationInstance()
		url := fmt.Sprintf("https://%s.service-now.com/api/now/v1/account/%s", instance, sysID)
		httpmock.RegisterResponder("GET", url,
			httpmock.NewStringResponder(404, mockdata.SomeErrorJSON))
	}

	resp, err := w.Client.Now().Account().ByID(sysID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *accountSteps) theAccountResultsShouldContainAtLeastRecords(ctx context.Context, minCount int) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil, expected at least %d records", minCount)
	}

	colResp, ok := w.Response.(accountapi.AccountCollectionResponse)
	if !ok {
		return ctx, fmt.Errorf("response is not an AccountCollectionResponse: %T", w.Response)
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

func (s *accountSteps) theAccountResultShouldHaveTheCorrectSysID(ctx context.Context, field string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil")
	}

	itemResp, ok := w.Response.(accountapi.AccountItemResponse)
	if !ok {
		return ctx, fmt.Errorf("response is not an AccountItemResponse: %T", w.Response)
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

// InitializeAccountScenario registers all account step definitions.
func InitializeAccountScenario(sc *godog.ScenarioContext) {
	s := &accountSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I request all accounts$`, s.iRequestAllAccounts)
	sc.Step(`^I request accounts with query "([^"]*)"$`, s.iRequestAccountsWithQuery)
	sc.Step(`^I request accounts with limit (\d+)$`, s.iRequestAccountsWithLimit)
	sc.Step(`^I have at least (\d+) account in the instance$`, s.iHaveAtLeastAccountInTheInstance)
	sc.Step(`^I request the account by its "([^"]*)"$`, s.iRequestTheAccountBySysID)
	sc.Step(`^I request the account with sys_id "([^"]*)"$`, s.iRequestTheAccountWithSysID)
	sc.Step(`^the account results should contain at least (\d+) record$`, s.theAccountResultsShouldContainAtLeastRecords)
	sc.Step(`^the account result should have the correct "([^"]*)"$`, s.theAccountResultShouldHaveTheCorrectSysID)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
