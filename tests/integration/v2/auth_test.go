//go:build integration

package v2

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/steps"
)

func TestAuthenticationScenarios(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeAuthScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/authentication"},
			Tags:     support.GodogTags(),
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run authentication feature tests")
	}
}
