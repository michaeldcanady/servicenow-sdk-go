// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

//go:build e2e

package v2

import (
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/steps"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load("../../../.env")

	if support.IsOffline() {
		fmt.Fprintln(os.Stderr, "e2e: SN_OFFLINE=true — skipping E2E tests (use integration build tag for mocked tests)")
		os.Exit(0)
	}

	if err := support.RequireCredentials(); err != nil {
		fmt.Fprintf(os.Stderr, "e2e: %v\n", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestE2ETableScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/table", steps.InitializeTableScenario)
}

func TestE2EAttachmentScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/attachment", steps.InitializeAttachmentScenario)
}

func TestE2EBatchScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/batch", steps.InitializeBatchScenario)
}

func TestE2EAccountScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/account", steps.InitializeAccountScenario)
}

func TestE2ECaseScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/case", steps.InitializeCaseScenario)
}

func TestE2ECmdbScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/cmdb", steps.InitializeCmdbScenario)
}

func TestE2EAggregationScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/aggregation", steps.InitializeAggregationScenario)
}

func TestE2EAuthScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/authentication", steps.InitializeAuthScenario)
}

func TestE2EDocumentsScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/documents", steps.InitializeDocumentsScenario)
}

func TestE2EActivitySubscriptionsScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/activity_subscriptions", steps.InitializeActivitySubscriptionsScenario)
}

func TestE2EAppointmentBookingScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/appointment_booking", steps.InitializeAppointmentBookingScenario)
}

func TestE2EAppServiceScenarios(t *testing.T) {
	if skipE2E(t) {
		return
	}
	runE2ESuite(t, "features/app_service", steps.InitializeAppServiceScenario)
}

// skipE2E returns true if the test should be skipped (e.g. missing env vars mid-run).
func skipE2E(t *testing.T) bool {
	t.Helper()
	if support.IsOffline() {
		t.Skip("Skipping E2E test: SN_OFFLINE=true")
		return true
	}
	return false
}

// runE2ESuite executes a godog suite for E2E with live instance tags.
func runE2ESuite(t *testing.T, featurePath string, initializer func(*godog.ScenarioContext)) {
	t.Helper()
	suite := godog.TestSuite{
		ScenarioInitializer: initializer,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{featurePath},
			Tags:     support.E2EGodogTags(),
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatalf("E2E scenarios failed: %s", featurePath)
	}
}
