package internal

type AuthorizationHandler struct {
	*BaseHandler
	provider AuthorizationProvider
}

func NewAuthorizationHandler(provider AuthorizationProvider) *AuthorizationHandler {
	return &AuthorizationHandler{
		BaseHandler: NewBaseHandler(),
		provider:    provider,
	}
}

func (a *AuthorizationHandler) Handle(request RequestInformation) error {
	err := a.provider.AuthorizeRequest(request)
	if err != nil {
		return err
	}
	return a.BaseHandler.Handle(request)
}
