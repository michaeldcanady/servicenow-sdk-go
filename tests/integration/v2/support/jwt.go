// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package support

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

// GenerateValidRS256JWT creates a valid RS256-signed JWT with all required claims.
func GenerateValidRS256JWT() (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := map[string]interface{}{
		"iss": "test-issuer",
		"sub": "test-subject",
		"aud": "test-audience",
		"exp": now.Add(1 * time.Hour).Unix(),
		"iat": now.Unix(),
		"jti": "test-jti-12345",
	}

	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`))
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)

	signingInput := header + "." + claimsEncoded
	hash := sha256.Sum256([]byte(signingInput))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	sigEncoded := base64.RawURLEncoding.EncodeToString(signature)

	return signingInput + "." + sigEncoded, nil
}

// GenerateJWTWithMissingClaims creates a JWT missing the "aud" claim.
func GenerateJWTWithMissingClaims() string {
	now := time.Now()
	claims := map[string]interface{}{
		"iss": "test-issuer",
		"sub": "test-subject",
		"exp": now.Add(1 * time.Hour).Unix(),
		"iat": now.Unix(),
		"jti": "test-jti-12345",
	}

	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`))
	claimsJSON, _ := json.Marshal(claims)
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)

	return header + "." + claimsEncoded + ".fake-signature"
}

// GenerateJWTWithWrongAlgorithm creates a JWT with HS256 algorithm (wrong for RS256 validation).
func GenerateJWTWithWrongAlgorithm() string {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	now := time.Now()
	claims := map[string]interface{}{
		"iss": "test-issuer",
		"sub": "test-subject",
		"aud": "test-audience",
		"exp": now.Add(1 * time.Hour).Unix(),
		"iat": now.Unix(),
		"jti": "test-jti-12345",
	}
	claimsJSON, _ := json.Marshal(claims)
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)

	return header + "." + claimsEncoded + ".fake-signature"
}
