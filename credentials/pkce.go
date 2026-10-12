// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"

// generatePKCEChallenge generates a PKCE code verifier and its challenge for the
// given method. A pkce.MethodUnset method produces empty values, matching a
// client that does not use PKCE. The verifier is returned even when challenge
// generation fails so callers can preserve it.
func generatePKCEChallenge(method pkce.Method) (challenge, verifier string, err error) {
	if method == pkce.MethodUnset {
		return "", "", nil
	}

	verifier, err = pkce.NewVerifierGenerator(pkce.DefaultVerifierEntropy).Generate()
	if err != nil {
		return "", "", err
	}

	challenge, err = pkce.NewPKCEChallenge(method, verifier)
	if err != nil {
		return "", verifier, err
	}

	return challenge, verifier, nil
}
