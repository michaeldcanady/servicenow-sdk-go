package credentials

import "errors"

type CustomPromptTokenCredential struct {
	credential *TokenCredential2
	instance   string
	options    *CustomPromptTokenOptions
}

type CustomPromptTokenOptions struct {
	options *TokenCredentialOptions
	prompt  func() (string, string, error)
}

func NewCustomPromptTokenCredential(options *CustomPromptTokenOptions) (*CustomPromptTokenCredential, error) {

	var credential CustomPromptTokenCredential
	var err error

	credential.credential, err = NewTokenCredential2(options.options)
	if err != nil {
		return nil, err
	}

	credential.credential.prompt = credential.Prompt

	return &credential, nil
}

func (cred *CustomPromptTokenCredential) Prompt() (*AccessToken, error) {

	err := cred.credential.server.Start2()
	if err != nil {
		return nil, err
	}

	defer cred.credential.server.Stop()

	return nil, errors.New("Prompt not implemented")
}
