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

	if support.IsOffline() {
		httpmock.Activate()
		support.RegisterAllMocks()
	} else {
		// Validate credentials are present for online mode
		if err := support.RequireCredentials(); err != nil {
			fmt.Fprintf(os.Stderr, "skipping online integration tests: %v\n", err)
			os.Exit(0)
		}
	}

	code := m.Run()

	if support.IsOffline() {
		httpmock.DeactivateAndReset()
	}
	os.Exit(code)
}
