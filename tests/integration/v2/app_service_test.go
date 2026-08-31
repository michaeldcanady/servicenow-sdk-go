// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

//go:build integration

package v2

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/steps"
)

func TestAppServiceScenarios(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeAppServiceScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/app_service"},
			Tags:     support.GodogTags(),
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("app service scenarios failed")
	}
}
