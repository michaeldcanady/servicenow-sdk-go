//go:build integration

package v2

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/steps"
)

func TestTableScenarios(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeTableScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/table"},
			Tags:     support.GodogTags(),
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("table scenarios failed")
	}
}
