package credentials

type TokenCredentialOptions struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
}

// validateOptions validates the token credential options
func validateOptions(options *TokenCredentialOptions) error {

	if options.ClientID == "" {
		return ErrMissingClientID
	}

	if options.ClientSecret == "" {
		return ErrMissingClientSecret
	}

	return nil
}
