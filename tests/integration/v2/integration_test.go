// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

//go:build integration

package v2

import (
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
)

func TestMain(m *testing.M) {
	// Load .env from the repo root
	_ = godotenv.Load("../../../.env")

	if support.IsE2E() {
		// Validate credentials are present for online mode
		if err := support.RequireCredentials(); err != nil {
			fmt.Fprintf(os.Stderr, "skipping online integration tests: %v\n", err)
			os.Exit(0)
		}
	} else {
		// Mocked mode: activate httpmock unless the run explicitly asked for a
		// live instance (SN_OFFLINE=false). An unset SN_OFFLINE with credentials
		// present is ambiguous; resolve it toward mock and say so, rather than
		// hitting the real network without mocks.
		support.WarnAmbiguousMode()
		httpmock.Activate()
		support.RegisterAllMocks()
	}

	code := m.Run()

	if !support.IsE2E() {
		httpmock.DeactivateAndReset()
	}
	os.Exit(code)
}
