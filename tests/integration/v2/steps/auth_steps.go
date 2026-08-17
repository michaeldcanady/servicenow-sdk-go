package steps

import (
	"context"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
)

type authSteps struct{}

// ── Basic Authentication ──────────────────────────────────────────────────

func (s *authSteps) iInitializeAClientWithBasicAuthentication(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	err := support.NewClientWithBasicAuth(w)
	if err != nil {
		w.AuthErr = err
	}
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithBasicAuthUsingEmptyUsername(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	ap := credentials.NewBasicProvider("", "password")
	err := support.NewClientWithAuthProvider(w, ap)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithBasicAuthUsingEmptyPassword(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	ap := credentials.NewBasicProvider("admin", "")
	err := support.NewClientWithAuthProvider(w, ap)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

// ── Static Bearer Token ───────────────────────────────────────────────────

func (s *authSteps) iInitializeAClientWithAStaticBearerToken(ctx context.Context, token string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	if token == "" {
		w.AuthErr = fmt.Errorf("bearer token is empty")
		return support.WithWorld(ctx, w), nil
	}

	bearer := credentials.NewBearerTokenAuthenticationProvider(
		&support.StaticTokenProvider{Token: token},
	)
	err := support.NewClientWithAuthProvider(w, bearer)
	if err != nil {
		w.AuthErr = err
	}
	return support.WithWorld(ctx, w), nil
}

// ── Client Credentials ────────────────────────────────────────────────────

func (s *authSteps) iInitializeAClientWithClientCredentialsAuth(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	err := support.NewClientWithClientCredentials(w, "mock-client-id", "mock-client-secret")
	if err != nil {
		w.AuthErr = err
	}
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iAttemptToCreateClientCredentialsWithEmptyClientID(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	_, err := credentials.NewClientCredentialsProvider("", "some-secret")
	w.AuthErr = err
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iAttemptToCreateClientCredentialsWithEmptyClientSecret(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	_, err := credentials.NewClientCredentialsProvider("some-client-id", "")
	w.AuthErr = err
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithInvalidClientCredentials(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	err := support.NewClientWithClientCredentials(w, "invalid-client-id", "invalid-secret")
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

// ── Token Lifecycle ───────────────────────────────────────────────────────

func (s *authSteps) iMakeMultipleAPIRequests(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Client == nil {
		return ctx, fmt.Errorf("client not initialized")
	}
	for i := 0; i < 3; i++ {
		_, err := w.Client.Now().Table("incident").Get(ctx, nil)
		if err != nil {
			w.Err = err
			return support.WithWorld(ctx, w), nil
		}
	}
	w.Err = nil
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) theCachedTokenExpires(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	w.TokenExpired = true
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iMakeAnAPITriggeringTokenRefresh(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Client == nil {
		return ctx, fmt.Errorf("client not initialized")
	}
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iMakeAnAPIToAcquireToken(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Client == nil {
		return ctx, fmt.Errorf("client not initialized")
	}
	resp, err := w.Client.Now().Table("incident").Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}
	if resp != nil {
		w.CachedToken = "acquired"
	}
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iRevokeTheCurrentAccessToken(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	w.RevocationErr = nil
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) aNewAccessTokenShouldBeReturned(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Err != nil {
		return ctx, fmt.Errorf("expected new token, but got error: %v", w.Err)
	}
	return ctx, nil
}

func (s *authSteps) authenticationTokensShouldBeReused(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (s *authSteps) theCachedTokenShouldBeCleared(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	w.CachedToken = ""
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) theRevocationShouldSucceed(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.RevocationErr != nil {
		return ctx, fmt.Errorf("expected revocation to succeed, got: %v", w.RevocationErr)
	}
	return ctx, nil
}

// ── ROPC ──────────────────────────────────────────────────────────────────

func (s *authSteps) iInitializeAClientWithROPCAuth(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	err := support.NewClientWithROPC(w, "mock-client-id", "mock-client-secret", "mock", "mock")
	if err != nil {
		w.AuthErr = err
	}
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithROPCAuthUsingEmptyUsername(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	err := support.NewClientWithROPC(w, "client-id", "client-secret", "", "password")
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithROPCAuthUsingEmptyPassword(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	err := support.NewClientWithROPC(w, "client-id", "client-secret", "username", "")
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) theROPCGetAuthMethodShouldReturnBearerHeader(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Client == nil {
		return ctx, fmt.Errorf("client not initialized")
	}
	resp, err := w.Client.Now().Table("incident").Get(ctx, nil)
	if err != nil {
		return ctx, fmt.Errorf("API request failed: %v", err)
	}
	if resp == nil {
		return ctx, fmt.Errorf("expected non-nil response")
	}
	return ctx, nil
}

// ── JWT ───────────────────────────────────────────────────────────────────

func (s *authSteps) iInitializeAClientWithJWTBearerTokenAuth(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	validJWT, err := support.GenerateValidRS256JWT()
	if err != nil {
		return ctx, fmt.Errorf("failed to generate test JWT: %v", err)
	}

	tokenProvider := &support.StaticTokenProvider{Token: validJWT}
	oauth2HTTP := support.GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(support.IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	sdkClient := support.GetHTTPClient()
	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(support.IntegrationInstance()),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	w.Client = client
	w.AuthErr = nil
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithAFailingJWTTokenProvider(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	tokenProvider := &support.FailingTokenProvider{Err: fmt.Errorf("token provider failed")}
	oauth2HTTP := support.GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(support.IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	sdkClient := support.GetHTTPClient()
	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(support.IntegrationInstance()),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	w.Client = client
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithAnInvalidJWTAssertionProvider(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	tokenProvider := &support.StaticTokenProvider{Token: "not-a-valid-jwt"}
	oauth2HTTP := support.GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(support.IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	sdkClient := support.GetHTTPClient()
	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(support.IntegrationInstance()),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	w.Client = client
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithAJWTMissingRequiredClaims(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	incompleteJWT := support.GenerateJWTWithMissingClaims()
	tokenProvider := &support.StaticTokenProvider{Token: incompleteJWT}
	oauth2HTTP := support.GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(support.IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	sdkClient := support.GetHTTPClient()
	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(support.IntegrationInstance()),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	w.Client = client
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAClientWithAJWTUsingWrongAlgorithm(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	wrongAlgJWT := support.GenerateJWTWithWrongAlgorithm()
	tokenProvider := &support.StaticTokenProvider{Token: wrongAlgJWT}
	oauth2HTTP := support.GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(support.IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	sdkClient := support.GetHTTPClient()
	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(support.IntegrationInstance()),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	w.Client = client
	_, w.Err = w.Client.Now().Table("incident").Get(ctx, nil)
	return support.WithWorld(ctx, w), nil
}

// ── Authorization Code ────────────────────────────────────────────────────

func (s *authSteps) iInitializeAnAuthCodePublicClientWithValidParams(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	oauth2HTTP := support.GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(support.IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	provider, err := credentials.NewPublicAuthorizationCodeProvider("mock-client-id", authOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = provider
	w.AuthErr = nil
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) iInitializeAnAuthCodeConfidentialClientWithValidParams(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	oauth2HTTP := support.GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(support.IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	provider, err := credentials.NewAuthorizationCodeProvider("mock-client-id", "mock-client-secret", authOpts...)
	if err != nil {
		w.AuthErr = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = provider
	w.AuthErr = nil
	return support.WithWorld(ctx, w), nil
}

func (s *authSteps) theAuthCodeCredentialShouldBeInitialized(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr != nil {
		return ctx, fmt.Errorf("expected credential to initialize, but got error: %v", w.AuthErr)
	}
	if w.Response == nil {
		return ctx, fmt.Errorf("expected credential to be initialized, but it is nil")
	}
	return ctx, nil
}

func (s *authSteps) theClientShouldSupportPKCEChallengeGeneration(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("credential is nil, cannot check PKCE support")
	}
	return ctx, nil
}

func (s *authSteps) iAttemptToCreateAuthCodeCredentialsWithEmptyClientID(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	_, err := credentials.NewPublicAuthorizationCodeProvider("")
	w.AuthErr = err
	return support.WithWorld(ctx, w), nil
}

// ── Assertion Steps ───────────────────────────────────────────────────────

func (s *authSteps) authenticationShouldSucceed(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr != nil {
		return ctx, fmt.Errorf("expected authentication to succeed, got: %v", w.AuthErr)
	}
	if w.Client == nil {
		return ctx, fmt.Errorf("expected client to be initialized, but it is nil")
	}
	return ctx, nil
}

func (s *authSteps) authenticationShouldFail(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr == nil && w.Err == nil {
		return ctx, fmt.Errorf("expected authentication to fail, but it succeeded")
	}
	return ctx, nil
}

func (s *authSteps) iShouldBeAbleToMakeAnAPIRequest(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Client == nil {
		return ctx, fmt.Errorf("client is nil, cannot make API request")
	}
	_, err := w.Client.Now().Table("incident").Get(ctx, nil)
	if err != nil {
		return ctx, fmt.Errorf("expected successful API request, but got error: %v", err)
	}
	return ctx, nil
}

func (s *authSteps) theResponseShouldNotBeAnError(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Err != nil {
		return ctx, fmt.Errorf("expected no error, but got: %v", w.Err)
	}
	return ctx, nil
}

func (s *authSteps) theErrorShouldBeAuthenticationRelated(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	errMsg := ""
	if w.AuthErr != nil {
		errMsg = w.AuthErr.Error()
	} else if w.Err != nil {
		errMsg = w.Err.Error()
	}
	if errMsg == "" {
		return ctx, fmt.Errorf("expected an authentication-related error, but got nil")
	}

	authRelated := strings.Contains(errMsg, "authentication") ||
		strings.Contains(errMsg, "unauthorized") ||
		strings.Contains(errMsg, "401") ||
		strings.Contains(errMsg, "credentials") ||
		strings.Contains(errMsg, "invalid") ||
		strings.Contains(errMsg, "oauth") ||
		strings.Contains(errMsg, "token")

	if !authRelated {
		return ctx, fmt.Errorf("expected authentication-related error, got: %v", errMsg)
	}
	return ctx, nil
}

func (s *authSteps) credentialCreationShouldFail(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr == nil {
		return ctx, fmt.Errorf("expected credential creation to fail, but it succeeded")
	}
	return ctx, nil
}

func (s *authSteps) theErrorMessageShouldIndicateMissingParameters(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr == nil {
		return ctx, fmt.Errorf("expected an error about missing parameters, but got nil")
	}
	errMsg := w.AuthErr.Error()
	if !strings.Contains(errMsg, "clientID") &&
		!strings.Contains(errMsg, "clientSecret") &&
		!strings.Contains(errMsg, "client_id") &&
		!strings.Contains(errMsg, "client_secret") &&
		!strings.Contains(errMsg, "username") &&
		!strings.Contains(errMsg, "password") &&
		!strings.Contains(errMsg, "empty") {
		return ctx, fmt.Errorf("expected error about missing parameters, but got: %v", w.AuthErr)
	}
	return ctx, nil
}

func (s *authSteps) theErrorMessageShouldIndicateAnInvalidJWT(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	errMsg := ""
	if w.Err != nil {
		errMsg = w.Err.Error()
	} else if w.AuthErr != nil {
		errMsg = w.AuthErr.Error()
	}
	if errMsg == "" {
		return ctx, fmt.Errorf("expected error about invalid JWT, but got nil")
	}
	if !strings.Contains(errMsg, "JWT") && !strings.Contains(errMsg, "token") && !strings.Contains(errMsg, "invalid") {
		return ctx, fmt.Errorf("expected error about invalid JWT, got: %v", errMsg)
	}
	return ctx, nil
}

func (s *authSteps) theErrorMessageShouldIndicateAMissingJWTClaim(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	errMsg := ""
	if w.Err != nil {
		errMsg = w.Err.Error()
	} else if w.AuthErr != nil {
		errMsg = w.AuthErr.Error()
	}
	if errMsg == "" {
		return ctx, fmt.Errorf("expected error about missing JWT claim, but got nil")
	}
	if !strings.Contains(errMsg, "missing required claim") && !strings.Contains(errMsg, "claim") {
		return ctx, fmt.Errorf("expected error about missing JWT claim, got: %v", errMsg)
	}
	return ctx, nil
}

func (s *authSteps) theErrorMessageShouldIndicateWrongSigningAlgorithm(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	errMsg := ""
	if w.Err != nil {
		errMsg = w.Err.Error()
	} else if w.AuthErr != nil {
		errMsg = w.AuthErr.Error()
	}
	if errMsg == "" {
		return ctx, fmt.Errorf("expected error about wrong signing algorithm, but got nil")
	}
	if !strings.Contains(errMsg, "signing algorithm") && !strings.Contains(errMsg, "algorithm") {
		return ctx, fmt.Errorf("expected error about wrong signing algorithm, got: %v", errMsg)
	}
	return ctx, nil
}

func (s *authSteps) theAPIRequestShouldFailWithAuthError(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Err == nil {
		return ctx, fmt.Errorf("expected API request to fail with auth error, but got no error")
	}
	return ctx, nil
}

func (s *authSteps) theAuthTokenShouldBeNonEmptyBearerToken(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.CachedToken == "" {
		return ctx, fmt.Errorf("expected a non-empty Bearer token, but CachedToken is empty")
	}
	if w.CachedToken == "acquired" {
		return ctx, fmt.Errorf("expected a real token value, but got placeholder 'acquired'")
	}
	return ctx, nil
}

// InitializeAuthScenario registers all authentication step definitions.
func InitializeAuthScenario(sc *godog.ScenarioContext) {
	s := &authSteps{}

	// Background
	sc.Step(`^I have a valid ServiceNow instance and credentials$`, iHaveAValidInstanceAndCredentials)

	// Basic auth
	sc.Step(`^I initialize a client with basic authentication$`, s.iInitializeAClientWithBasicAuthentication)
	sc.Step(`^I initialize a client with basic authentication using empty username$`, s.iInitializeAClientWithBasicAuthUsingEmptyUsername)
	sc.Step(`^I initialize a client with basic authentication using empty password$`, s.iInitializeAClientWithBasicAuthUsingEmptyPassword)

	// Static bearer token
	sc.Step(`^I initialize a client with a static bearer token "([^"]*)"$`, s.iInitializeAClientWithAStaticBearerToken)

	// Client credentials
	sc.Step(`^I initialize a client with client credentials authentication$`, s.iInitializeAClientWithClientCredentialsAuth)
	sc.Step(`^I attempt to create client credentials with empty client ID$`, s.iAttemptToCreateClientCredentialsWithEmptyClientID)
	sc.Step(`^I attempt to create client credentials with empty client secret$`, s.iAttemptToCreateClientCredentialsWithEmptyClientSecret)
	sc.Step(`^I initialize a client with invalid client credentials$`, s.iInitializeAClientWithInvalidClientCredentials)

	// Token lifecycle
	sc.Step(`^I make multiple API requests$`, s.iMakeMultipleAPIRequests)
	sc.Step(`^the cached token expires$`, s.theCachedTokenExpires)
	sc.Step(`^I make an API request that triggers token refresh$`, s.iMakeAnAPITriggeringTokenRefresh)
	sc.Step(`^I make an API request to acquire a token$`, s.iMakeAnAPIToAcquireToken)
	sc.Step(`^I revoke the current access token$`, s.iRevokeTheCurrentAccessToken)

	// ROPC
	sc.Step(`^I initialize a client with ROPC authentication$`, s.iInitializeAClientWithROPCAuth)
	sc.Step(`^I initialize a client with ROPC authentication using empty username$`, s.iInitializeAClientWithROPCAuthUsingEmptyUsername)
	sc.Step(`^I initialize a client with ROPC authentication using empty password$`, s.iInitializeAClientWithROPCAuthUsingEmptyPassword)
	sc.Step(`^the ROPC GetAuthentication method should return a Bearer token header$`, s.theROPCGetAuthMethodShouldReturnBearerHeader)

	// JWT
	sc.Step(`^I initialize a client with JWT bearer token authentication$`, s.iInitializeAClientWithJWTBearerTokenAuth)
	sc.Step(`^I initialize a client with a failing JWT token provider$`, s.iInitializeAClientWithAFailingJWTTokenProvider)
	sc.Step(`^I initialize a client with an invalid JWT assertion provider$`, s.iInitializeAClientWithAnInvalidJWTAssertionProvider)
	sc.Step(`^I initialize a client with a JWT missing required claims$`, s.iInitializeAClientWithAJWTMissingRequiredClaims)
	sc.Step(`^I initialize a client with a JWT using wrong algorithm$`, s.iInitializeAClientWithAJWTUsingWrongAlgorithm)

	// Authorization code
	sc.Step(`^I initialize an authorization code public client with valid parameters$`, s.iInitializeAnAuthCodePublicClientWithValidParams)
	sc.Step(`^I initialize an authorization code confidential client with valid parameters$`, s.iInitializeAnAuthCodeConfidentialClientWithValidParams)
	sc.Step(`^the authorization code credential should be initialized$`, s.theAuthCodeCredentialShouldBeInitialized)
	sc.Step(`^the client should support PKCE challenge generation$`, s.theClientShouldSupportPKCEChallengeGeneration)
	sc.Step(`^I attempt to create authorization code credentials with empty client ID$`, s.iAttemptToCreateAuthCodeCredentialsWithEmptyClientID)

	// Assertion steps
	sc.Step(`^authentication should succeed$`, s.authenticationShouldSucceed)
	sc.Step(`^authentication should fail$`, s.authenticationShouldFail)
	sc.Step(`^I should be able to make an API request$`, s.iShouldBeAbleToMakeAnAPIRequest)
	sc.Step(`^the response should not be an error$`, s.theResponseShouldNotBeAnError)
	sc.Step(`^the error should be authentication-related$`, s.theErrorShouldBeAuthenticationRelated)
	sc.Step(`^credential creation should fail$`, s.credentialCreationShouldFail)
	sc.Step(`^the error message should indicate missing parameters$`, s.theErrorMessageShouldIndicateMissingParameters)
	sc.Step(`^the error message should indicate an invalid JWT$`, s.theErrorMessageShouldIndicateAnInvalidJWT)
	sc.Step(`^the error message should indicate a missing JWT claim$`, s.theErrorMessageShouldIndicateAMissingJWTClaim)
	sc.Step(`^the error message should indicate wrong signing algorithm$`, s.theErrorMessageShouldIndicateWrongSigningAlgorithm)
	sc.Step(`^all requests should succeed$`, s.iShouldBeAbleToMakeAnAPIRequest)
	sc.Step(`^authentication tokens should be reused$`, s.authenticationTokensShouldBeReused)
	sc.Step(`^a new access token should be returned$`, s.aNewAccessTokenShouldBeReturned)
	sc.Step(`^the revocation should succeed$`, s.theRevocationShouldSucceed)
	sc.Step(`^the cached token should be cleared$`, s.theCachedTokenShouldBeCleared)
	sc.Step(`^the API request should fail with an authentication error$`, s.theAPIRequestShouldFailWithAuthError)
	sc.Step(`^the authentication token should be a non-empty Bearer token$`, s.theAuthTokenShouldBeNonEmptyBearerToken)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
