package support

import (
	"fmt"
	"os"
	"strings"
)

// IsOffline returns true when running with httpmock (no live instance).
func IsOffline() bool {
	return os.Getenv("SN_OFFLINE") == "true"
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
	if IsOffline() {
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

// IsE2E returns true when running against a live ServiceNow instance (not mocked).
func IsE2E() bool {
	return !IsOffline()
}

// GodogTags returns the tag expression for godog integration (mocked) runs.
// Offline: includes @offline scenarios, excludes @e2e-only scenarios.
// Online: excludes them (~ is NOT in godog).
func GodogTags() string {
	if IsOffline() {
		return "integration && mock && ~e2e"
	}
	return "integration && mock && ~offline && ~e2e"
}

// E2EGodogTags returns the tag expression for godog E2E (live) runs.
// Runs all integration scenarios except those tagged @offline (mock-only edge cases).
func E2EGodogTags() string {
	return "integration && ~offline"
}
