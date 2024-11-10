package http

type AuthorizationProvider interface {
	AuthorizeRequest(request RequestInformation) error
}
