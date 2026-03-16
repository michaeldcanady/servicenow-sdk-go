package pkce

import (
	"crypto/sha256"
	"encoding/base64"
)

type S256Challenger struct{}

func (S256Challenger) Name() Method { return MethodS256 }

func (S256Challenger) Challenge(verifier string) (string, error) {
	sum := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(sum[:]), nil
}
