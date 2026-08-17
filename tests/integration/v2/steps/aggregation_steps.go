package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/aggregationapi"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
)

type aggregationSteps struct{}

func (s *aggregationSteps) iRequestACountAggregationForTheTable(ctx context.Context, tableName string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	count := true
	cfg := &aggregationapi.StatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &aggregationapi.StatsRequestBuilderGetQueryParameters{
			Count: &count,
		},
	}

	resp, err := w.Client.Now().Stats(tableName).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *aggregationSteps) iRequestAQueryAggregationForTheTableWithQuery(ctx context.Context, tableName, query string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	count := true
	cfg := &aggregationapi.StatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &aggregationapi.StatsRequestBuilderGetQueryParameters{
			Count: &count,
			Query: &query,
		},
	}

	resp, err := w.Client.Now().Stats(tableName).Get(ctx, cfg)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *aggregationSteps) iRequestAnAggregationWithoutParametersFor(ctx context.Context, tableName string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Stats(tableName).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *aggregationSteps) theAggregationResultShouldNotBeNil(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("expected aggregation result to not be nil")
	}
	return ctx, nil
}

// InitializeAggregationScenario registers all aggregation step definitions.
func InitializeAggregationScenario(sc *godog.ScenarioContext) {
	s := &aggregationSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I request a count aggregation for the "([^"]*)" table$`, s.iRequestACountAggregationForTheTable)
	sc.Step(`^I request a query aggregation for the "([^"]*)" table with query "([^"]*)"$`, s.iRequestAQueryAggregationForTheTableWithQuery)
	sc.Step(`^I request an aggregation without parameters for "([^"]*)"$`, s.iRequestAnAggregationWithoutParametersFor)
	sc.Step(`^the aggregation result should not be nil$`, s.theAggregationResultShouldNotBeNil)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
