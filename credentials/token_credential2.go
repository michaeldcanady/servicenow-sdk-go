package credentials

// TokenCredential2 used for token authetication
type TokenCredential2 struct {
	options *TokenCredentialOptions
	prompt  func() (*AccessToken, error)
	server  *HTTPServer
}

// NewTokenCredential2 creates a new TokenCredential2 instance with the specified options
func NewTokenCredential2(options *TokenCredentialOptions) (*TokenCredential2, error) {

	var credential TokenCredential2
	var serverConfig serverConfig

	err := validateOptions(options)
	if err != nil {
		return nil, err
	}

	credential.options = options

	validateServerConfig(&serverConfig)

	credential.server, err = startLocalServer(options.RedirectUrl, &serverConfig)
	if err != nil {
		return nil, err
	}
	return &credential, nil
}
