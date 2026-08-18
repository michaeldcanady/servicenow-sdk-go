//go:build integration

package v2

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/steps"
)

func TestCaseScenarios(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeCaseScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/case"},
			Tags:     support.GodogTags(),
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("case scenarios failed")
	}
}
