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

func TestAppointmentBookingScenarios(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeAppointmentBookingScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/appointment_booking"},
			Tags:     support.GodogTags(),
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("appointment booking scenarios failed")
	}
}
