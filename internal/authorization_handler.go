package internal

type AuthorizationHandler struct {
	handler  BaseHandler
	provider AuthorizationProvider
}

func (a *AuthorizationHandler) Handle(request RequestInformation) error {
	err := a.provider.AuthorizeRequest(request)
	if err != nil {
		return err
	}
	return a.handler.Handle(request)
}
