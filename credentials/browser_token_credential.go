package credentials

import "errors"

type BrowserTokenCredential struct {
	credential *TokenCredential2
	instance   string
}

type BrowserTokenOptions struct {
	options *TokenCredentialOptions
}

func NewBrowserTokenCredential(options *BrowserTokenOptions) (*BrowserTokenCredential, error) {

	var credential BrowserTokenCredential
	var err error

	credential.credential, err = NewTokenCredential2(options.options)
	if err != nil {
		return nil, err
	}

	credential.credential.prompt = credential.Prompt

	return &credential, nil
}

func (cred *BrowserTokenCredential) Prompt() (*AccessToken, error) {

	err := cred.credential.server.Start2()
	if err != nil {
		return nil, err
	}

	defer cred.credential.server.Stop()

	return nil, errors.New("Prompt not implemented")
}
