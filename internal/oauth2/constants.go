package oauth2

// --------------------
// Constants
// --------------------

// Content types
const (
	ContentTypeKey            = "Content-Type"
	AcceptKey                 = "Accept"
	FormURLEncodedContentType = "application/x-www-form-urlencoded"
	JSONContentType           = "application/json"
)

// Common OAuth2 keys
const (
	ClientIDKey            = "client_id"
	ClientSecretKey        = "client_secret"
	GrantTypeKey           = "grant_type"
	ScopeKey               = "scope"
	CodeKey                = "code"
	RefreshTokenKey        = "refresh_token"
	UsernameKey            = "username"
	PasswordKey            = "password"
	RedirectURIKey         = "redirect_uri"
	CodeVerifierKey        = "code_verifier"
	CodeChallengeKey       = "code_challenge"
	CodeChallengeMethodKey = "code_challenge_method"
	AssertionKey           = "assertion"
	StateKey               = "state"
)

// Auth URL query params
const (
	ResponseTypeKey  = "response_type"
	ResponseTypeCode = "code"
	AuthorizationKey = "Authorization"
)

// Grant type values
const (
	GrantTypeAuthCode     = "authorization_code"
	GrantTypeClientCreds  = "client_credentials"
	GrantTypeRefreshToken = "refresh_token"
	GrantTypePassword     = "password"
	GrantTypeJWTBearer    = "urn:ietf:params:oauth:grant-type:jwt-bearer" //nolint:gosec
)
