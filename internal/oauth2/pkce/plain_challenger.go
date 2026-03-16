package pkce

type PlainChallenger struct{}

func (PlainChallenger) Name() Method { return MethodPlain }

func (PlainChallenger) Challenge(verifier string) (string, error) {
	return verifier, nil
}
