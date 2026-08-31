// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package pkce

type PlainChallenger struct{}

func (PlainChallenger) Name() Method { return MethodPlain }

func (PlainChallenger) Challenge(verifier string) (string, error) {
	return verifier, nil
}
