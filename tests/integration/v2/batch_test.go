//go:build integration

package v2

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/steps"
)

func TestBatchScenarios(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeBatchScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/batch"},
			Tags:     support.GodogTags(),
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("batch scenarios failed")
	}
}
