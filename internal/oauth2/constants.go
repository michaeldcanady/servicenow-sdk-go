package oauth2

// --------------------
// Constants
// --------------------

// HTTP header keys and content types.
const (
	ContentTypeKey            = "Content-Type"
	AcceptKey                 = "Accept"
	FormURLEncodedContentType = "application/x-www-form-urlencoded"
	JSONContentType           = "application/json"
)

// Common OAuth2 parameter keys.
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
	ErrorKey               = "error"
	ErrorDescriptionKey    = "error_description"

	// Device Flow specific keys.
	DeviceCodeKey      = "device_code"
	UserCodeKey        = "user_code"
	VerificationURIKey = "verification_uri"
	ExpiresInKey       = "expires_in"
	IntervalKey        = "interval"

	// Revocation and Introspection specific keys.
	TokenKey         = "token"
	TokenTypeHintKey = "token_type_hint"
	ActiveKey        = "active"
)

// Authorization URL query parameters and header keys.
const (
	ResponseTypeKey  = "response_type"
	ResponseTypeCode = "code"
	AuthorizationKey = "Authorization"
)

// Standard OAuth2 grant type values.
const (
	GrantTypeAuthCode     = "authorization_code"
	GrantTypeClientCreds  = "client_credentials"
	GrantTypeRefreshToken = "refresh_token"
	GrantTypePassword     = "password"
	GrantTypeJWTBearer    = "urn:ietf:params:oauth:grant-type:jwt-bearer" //nolint:gosec
	GrantTypeDeviceCode   = "urn:ietf:params:oauth:grant-type:device_code"
)
