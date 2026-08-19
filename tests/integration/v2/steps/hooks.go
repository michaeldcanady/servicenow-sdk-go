package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
)

// Hooks are used by domain test files. Exported for reuse.

// BeforeScenario creates a fresh World for each scenario.
func BeforeScenario(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	w := support.NewWorld()
	return support.WithWorld(ctx, w), nil
}

// AfterScenario cleans up resources created during the scenario.
func AfterScenario(ctx context.Context, sc *godog.Scenario, scenarioErr error) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w != nil {
		for _, sysID := range w.CreatedIDs {
			support.CleanupResourceLog(ctx, w, sysID)
		}
	}
	return ctx, nil
}
