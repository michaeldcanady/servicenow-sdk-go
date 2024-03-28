package internal

type AuthorizationHandler struct {
	*RequestHandler
	provider AuthorizationProvider
}

func NewAuthorizationHandler(provider AuthorizationProvider) *AuthorizationHandler {
	return &AuthorizationHandler{
		RequestHandler: NewRequestHandler(),
		provider:       provider,
	}
}

func (a *AuthorizationHandler) Handle(request RequestInformation) error {
	err := a.provider.AuthorizeRequest(request)
	if err != nil {
		return err
	}
	return a.BaseHandler.Handle(request)
}
