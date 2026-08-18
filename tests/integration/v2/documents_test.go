//go:build integration

package v2

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/steps"
)

func TestDocumentsScenarios(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeDocumentsScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/documents"},
			Tags:     support.GodogTags(),
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("documents scenarios failed")
	}
}
