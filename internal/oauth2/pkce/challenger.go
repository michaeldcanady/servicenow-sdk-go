package pkce

import "errors"

type Challenger interface {
	Name() Method
	Challenge(verifier string) (string, error)
}

func ChallengerFor(method Method) (Challenger, error) {
	switch method {
	case MethodS256:
		return S256Challenger{}, nil
	case MethodPlain:
		return PlainChallenger{}, nil
	default:
		return nil, errors.New("unsupported PKCE method")
	}
}

func NewPKCEChallenge(method Method, verifier string) (string, error) {
	if verifier == "" {
		return "", errors.New("verifier cannot be empty")
	}

	challenger, err := ChallengerFor(method)
	if err != nil {
		return "", err
	}

	return challenger.Challenge(verifier)
}
