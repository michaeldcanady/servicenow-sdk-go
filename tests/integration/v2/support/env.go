// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package support

import (
	"fmt"
	"os"
	"strings"
)

// IsOffline returns true when requests run against httpmock (no live instance).
// It is the exact negation of IsE2E, so the two never disagree: a run is mocked
// unless it explicitly sets SN_OFFLINE=false.
func IsOffline() bool {
	return !IsE2E()
}

// WarnAmbiguousMode writes a diagnostic to stderr when credentials are present
// but SN_OFFLINE is unset. Callers that resolve the ambiguity toward mock mode
// (the integration suite) use this so the choice is visible in the log.
func WarnAmbiguousMode() {
	if os.Getenv("SN_OFFLINE") != "" {
		return
	}
	if FirstEnv("SN_USERNAME", "SNOW_USERNAME") == "" {
		return
	}
	fmt.Fprintln(os.Stderr,
		"integration: SN_USERNAME is set but SN_OFFLINE is unset; running in "+
			"mocked mode. Set SN_OFFLINE=false to run against a live instance, "+
			"or SN_OFFLINE=true to silence this notice.")
}

// FirstEnv returns the first non-empty environment variable from the given names.
func FirstEnv(keys ...string) string {
	for _, k := range keys {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}

// IntegrationInstance returns the ServiceNow instance name from the environment,
// falling back to "mock_instance" for offline testing.
func IntegrationInstance() string {
	inst := FirstEnv("SN_INSTANCE", "SNOW_INSTANCE")
	if inst == "" {
		inst = "mock_instance"
	}
	return strings.TrimSpace(inst)
}

// ClientCredentials returns OAuth2 client credentials from the environment
// (SN_CLIENT_ID/SNOW_CLIENT_ID, SN_CLIENT_SECRET/SNOW_CLIENT_SECRET), falling
// back to the historical mock pair so offline/httpmock suites keep matching
// their recorded requests unchanged.
func ClientCredentials() (clientID, clientSecret string) {
	clientID = FirstEnv("SN_CLIENT_ID", "SNOW_CLIENT_ID")
	if clientID == "" {
		clientID = "mock-client-id"
	}
	clientSecret = FirstEnv("SN_CLIENT_SECRET", "SNOW_CLIENT_SECRET")
	if clientSecret == "" {
		clientSecret = "mock-client-secret"
	}
	return clientID, clientSecret
}

// RequireCredentials validates that required env vars are set for online mode.
// Returns an error if running online without credentials.
func RequireCredentials() error {
	if !IsE2E() {
		return nil
	}
	if FirstEnv("SN_USERNAME", "SNOW_USERNAME") == "" {
		return fmt.Errorf("online mode requires SN_USERNAME or SNOW_USERNAME environment variable")
	}
	if FirstEnv("SN_PASSWORD", "SNOW_PASSWORD") == "" {
		return fmt.Errorf("online mode requires SN_PASSWORD or SNOW_PASSWORD environment variable")
	}
	return nil
}

// IsE2E returns true when the environment explicitly asks to run against a
// live ServiceNow instance. It requires SN_OFFLINE=false, so an unset
// SN_OFFLINE never counts as live. The mocked integration suite therefore runs
// under httpmock unless you set SN_OFFLINE=false deliberately.
func IsE2E() bool {
	return os.Getenv("SN_OFFLINE") == "false"
}

// GodogTags returns the tag expression for godog integration (mocked) runs.
// Live mode: excludes scenarios that only make sense against httpmock.
// Mocked mode: includes @offline scenarios, excludes @e2e-only scenarios.
func GodogTags() string {
	if IsE2E() {
		return "integration && mock && ~offline && ~e2e"
	}
	return "integration && mock && ~e2e"
}

// E2EGodogTags returns the tag expression for godog E2E (live) runs.
// Runs all integration scenarios except those tagged @offline (mock-only edge cases).
func E2EGodogTags() string {
	return "integration && ~offline"
}
