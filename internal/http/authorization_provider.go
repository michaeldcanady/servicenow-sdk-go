package http

const (
	authorizationHeader = "Authorization"
)

type AuthorizationProvider interface {
	AuthorizeRequest(request RequestInformation) error
}
