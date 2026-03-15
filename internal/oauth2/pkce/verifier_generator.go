package pkce

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type VerifierGenerator struct {
	EntropyBytes int
	Rand         io.Reader
}

const DefaultVerifierEntropy = 32

func NewVerifierGenerator(entropyBytes int) VerifierGenerator {
	return VerifierGenerator{
		EntropyBytes: entropyBytes,
		Rand:         rand.Reader,
	}
}

func (g VerifierGenerator) Generate() (string, error) {
	if g.EntropyBytes <= 0 {
		return "", errors.New("entropy must be > 0")
	}
	if g.Rand == nil {
		return "", errors.New("rand reader cannot be nil")
	}

	buf := make([]byte, g.EntropyBytes)
	if _, err := io.ReadFull(g.Rand, buf); err != nil {
		return "", err
	}

	// RFC 7636 requires URL-safe base64 without padding
	return base64.RawURLEncoding.EncodeToString(buf), nil
}
