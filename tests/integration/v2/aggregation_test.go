//go:build integration

package v2

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/steps"
)

func TestAggregationScenarios(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeAggregationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/aggregation"},
			Tags:     support.GodogTags(),
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("aggregation scenarios failed")
	}
}
