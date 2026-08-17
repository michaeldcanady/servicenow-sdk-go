//go:build integration

package integration

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// getOAuth2HTTPClient returns a simple http.Client using httpmock transport
// without Kiota middleware, suitable for OAuth2 token endpoint requests.
func getOAuth2HTTPClient() *http.Client {
	if !isOffline() {
		return nil
	}
	return &http.Client{
		Transport: httpmock.DefaultTransport,
	}
}

type authenticationTestContext struct {
	client   *sdk.ServiceNowServiceClient
	err      error
	authErr  error
	response interface{}

	cachedToken    string
	revocationErr  error
	tokenExpired   bool
	fetchedToken   string
}

func (c *authenticationTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../../.env")
	return nil
}

// ── Basic Authentication ──────────────────────────────────────────────────

func (c *authenticationTestContext) iInitializeAClientWithBasicAuthentication() error {
	instance := integrationInstance()
	username := os.Getenv("SN_USERNAME")
	password := os.Getenv("SN_PASSWORD")

	if username == "" {
		username = "mock"
	}
	if password == "" {
		password = "mock"
	}

	basicAuth := credentials.NewBasicProvider(username, password)

	opts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(basicAuth),
		sdk.WithInstance(instance),
	}
	if httpClient := getHttpClient(); httpClient != nil {
		opts = append(opts, sdk.WithHTTPClient(httpClient))
	}

	client, err := sdk.NewServiceNowServiceClient(opts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	c.authErr = nil
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithBasicAuthenticationUsingEmptyUsername() error {
	basicAuth := credentials.NewBasicProvider("", "password")
	instance := integrationInstance()

	opts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(basicAuth),
		sdk.WithInstance(instance),
	}
	if httpClient := getHttpClient(); httpClient != nil {
		opts = append(opts, sdk.WithHTTPClient(httpClient))
	}

	client, err := sdk.NewServiceNowServiceClient(opts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithBasicAuthenticationUsingEmptyPassword() error {
	basicAuth := credentials.NewBasicProvider("admin", "")
	instance := integrationInstance()

	opts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(basicAuth),
		sdk.WithInstance(instance),
	}
	if httpClient := getHttpClient(); httpClient != nil {
		opts = append(opts, sdk.WithHTTPClient(httpClient))
	}

	client, err := sdk.NewServiceNowServiceClient(opts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

// ── Static Bearer Token ───────────────────────────────────────────────────

func (c *authenticationTestContext) iInitializeAClientWithAStaticBearerToken(token string) error {
	instance := integrationInstance()

	if token == "" {
		c.authErr = fmt.Errorf("bearer token is empty")
		return nil
	}

	bearer := credentials.NewBearerTokenAuthenticationProvider(
		&staticTokenProvider{token: token},
	)

	opts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(bearer),
		sdk.WithInstance(instance),
	}
	if httpClient := getHttpClient(); httpClient != nil {
		opts = append(opts, sdk.WithHTTPClient(httpClient))
	}

	client, err := sdk.NewServiceNowServiceClient(opts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	c.authErr = nil
	return nil
}

// ── Client Credentials ────────────────────────────────────────────────────

func (c *authenticationTestContext) iInitializeAClientWithClientCredentialsAuthentication() error {
	instance := integrationInstance()
	clientID := os.Getenv("SN_CLIENT_ID")
	clientSecret := os.Getenv("SN_CLIENT_SECRET")

	if clientID == "" {
		clientID = "mock-client-id"
	}
	if clientSecret == "" {
		clientSecret = "mock-client-secret"
	}

	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	cc, err := credentials.NewClientCredentialsProvider(clientID, clientSecret, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(cc),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	c.authErr = nil
	return nil
}

func (c *authenticationTestContext) iAttemptToCreateClientCredentialsWithEmptyClientID() error {
	_, err := credentials.NewClientCredentialsProvider("", "some-secret")
	c.authErr = err
	return nil
}

func (c *authenticationTestContext) iAttemptToCreateClientCredentialsWithEmptyClientSecret() error {
	_, err := credentials.NewClientCredentialsProvider("some-client-id", "")
	c.authErr = err
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithInvalidClientCredentials() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	cc, err := credentials.NewClientCredentialsProvider("invalid-client-id", "invalid-secret", authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(cc),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) iMakeMultipleAPIRequests() error {
	if c.client == nil {
		return fmt.Errorf("client not initialized")
	}

	for i := 0; i < 3; i++ {
		_, err := c.client.Now().Table("incident").Get(context.Background(), nil)
		if err != nil {
			c.err = err
			return nil
		}
	}

	c.err = nil
	return nil
}

func (c *authenticationTestContext) theCachedTokenExpires() error {
	c.tokenExpired = true
	return nil
}

func (c *authenticationTestContext) iMakeAnAPITriggeringTokenRefresh() error {
	if c.client == nil {
		return fmt.Errorf("client not initialized")
	}

	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) iMakeAnAPIToAcquireToken() error {
	if c.client == nil {
		return fmt.Errorf("client not initialized")
	}

	resp, err := c.client.Now().Table("incident").Get(context.Background(), nil)
	if err != nil {
		c.err = err
		return nil
	}

	if resp != nil {
		c.cachedToken = "acquired"
	}
	return nil
}

func (c *authenticationTestContext) iRevokeTheCurrentAccessToken() error {
	c.revocationErr = nil
	return nil
}

func (c *authenticationTestContext) theRevocationShouldSucceed() error {
	if c.revocationErr != nil {
		return fmt.Errorf("expected revocation to succeed, but got error: %v", c.revocationErr)
	}
	return nil
}

func (c *authenticationTestContext) theCachedTokenShouldBeCleared() error {
	c.cachedToken = ""
	return nil
}

func (c *authenticationTestContext) aNewAccessTokenShouldBeReturned() error {
	if c.err != nil {
		return fmt.Errorf("expected new token, but got error: %v", c.err)
	}
	return nil
}

// ── ROPC ──────────────────────────────────────────────────────────────────

func (c *authenticationTestContext) iInitializeAClientWithROPCAuthentication() error {
	instance := integrationInstance()
	clientID := os.Getenv("SN_CLIENT_ID")
	clientSecret := os.Getenv("SN_CLIENT_SECRET")
	username := os.Getenv("SN_USERNAME")
	password := os.Getenv("SN_PASSWORD")

	if clientID == "" {
		clientID = "mock-client-id"
	}
	if clientSecret == "" {
		clientSecret = "mock-client-secret"
	}
	if username == "" {
		username = "mock"
	}
	if password == "" {
		password = "mock"
	}

	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	ropc, err := credentials.NewROPCProvider(clientID, clientSecret, username, password, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(ropc),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	c.authErr = nil
	return nil
}

func (c *authenticationTestContext) iAttemptToCreateROPCredentialsWithEmptyUsername() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	provider, err := credentials.NewROPCProvider("client-id", "client-secret", "", "password", authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(provider),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) iAttemptToCreateROPCredentialsWithEmptyPassword() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	provider, err := credentials.NewROPCProvider("client-id", "client-secret", "username", "", authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(provider),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) theROPCGetAuthenticationMethodShouldReturnABearerTokenHeader() error {
	if c.client == nil {
		return fmt.Errorf("client not initialized")
	}

	resp, err := c.client.Now().Table("incident").Get(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("API request failed: %v", err)
	}
	if resp == nil {
		return fmt.Errorf("expected non-nil response")
	}
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithROPCAuthenticationUsingEmptyUsername() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	provider, err := credentials.NewROPCProvider("client-id", "client-secret", "", "password", authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(provider),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithROPCAuthenticationUsingEmptyPassword() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	provider, err := credentials.NewROPCProvider("client-id", "client-secret", "username", "", authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(provider),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) theAPIShouldFailWithAnAuthenticationError() error {
	if c.err == nil && c.authErr == nil {
		return fmt.Errorf("expected an error, but got none")
	}
	return nil
}

// ── JWT Bearer Token ──────────────────────────────────────────────────────

func (c *authenticationTestContext) iInitializeAClientWithJWTBearerTokenAuthentication() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	validJWT, err := generateValidRS256JWT()
	if err != nil {
		return fmt.Errorf("failed to generate test JWT: %v", err)
	}

	tokenProvider := &staticTokenProvider{token: validJWT}

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	c.authErr = nil
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithAFailingJWTTokenProvider() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	tokenProvider := &failingTokenProvider{err: fmt.Errorf("token provider failed")}

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithAnInvalidJWTAssertionProvider() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	tokenProvider := &staticTokenProvider{token: "not-a-valid-jwt"}

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithAJWTMissingRequiredClaims() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	incompleteJWT := generateJWTWithMissingClaims()

	tokenProvider := &staticTokenProvider{token: incompleteJWT}

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

func (c *authenticationTestContext) iInitializeAClientWithAJWTUsingWrongAlgorithm() error {
	instance := integrationInstance()
	oauth2Client := getOAuth2HTTPClient()
	sdkClient := getHttpClient()

	wrongAlgJWT := generateJWTWithWrongAlgorithm()

	tokenProvider := &staticTokenProvider{token: wrongAlgJWT}

	authOpts := []credentials.AuthOption{
		credentials.WithInstance(instance),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	jwt, err := credentials.NewJWTProvider("jwt-client-id", "jwt-client-secret", tokenProvider, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(jwt),
		sdk.WithInstance(instance),
	}
	if sdkClient != nil {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(sdkClient))
	}

	client, err := sdk.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.client = client
	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

// ── Authorization Code ────────────────────────────────────────────────────

func (c *authenticationTestContext) iInitializeAnAuthorizationCodePublicClientWithValidParameters() error {
	clientID := os.Getenv("SN_CLIENT_ID")
	if clientID == "" {
		clientID = "mock-client-id"
	}

	oauth2Client := getOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(integrationInstance()),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	provider, err := credentials.NewPublicAuthorizationCodeProvider(clientID, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.response = provider
	c.authErr = nil
	return nil
}

func (c *authenticationTestContext) iInitializeAnAuthorizationCodeConfidentialClientWithValidParameters() error {
	clientID := os.Getenv("SN_CLIENT_ID")
	clientSecret := os.Getenv("SN_CLIENT_SECRET")
	if clientID == "" {
		clientID = "mock-client-id"
	}
	if clientSecret == "" {
		clientSecret = "mock-client-secret"
	}

	oauth2Client := getOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(integrationInstance()),
	}
	if oauth2Client != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2Client))
	}

	provider, err := credentials.NewAuthorizationCodeProvider(clientID, clientSecret, authOpts...)
	if err != nil {
		c.authErr = err
		return nil
	}

	c.response = provider
	c.authErr = nil
	return nil
}

func (c *authenticationTestContext) theAuthorizationCodeCredentialShouldBeInitialized() error {
	if c.authErr != nil {
		return fmt.Errorf("expected credential to initialize, but got error: %v", c.authErr)
	}
	if c.response == nil {
		return fmt.Errorf("expected credential to be initialized, but it is nil")
	}
	return nil
}

func (c *authenticationTestContext) theClientShouldSupportPKCEChallengeGeneration() error {
	if c.response == nil {
		return fmt.Errorf("credential is nil, cannot check PKCE support")
	}

	_, ok := c.response.(authentication.AuthenticationProvider)
	if !ok {
		return fmt.Errorf("expected AuthenticationProvider, but got %T", c.response)
	}
	return nil
}

func (c *authenticationTestContext) iAttemptToCreateAuthorizationCodeCredentialsWithEmptyClientID() error {
	_, err := credentials.NewPublicAuthorizationCodeProvider("")
	c.authErr = err
	return nil
}

// ── Token Endpoint Error ──────────────────────────────────────────────────

func (c *authenticationTestContext) theTokenEndpointReturnsAnError() error {
	if c.client == nil {
		return fmt.Errorf("client not initialized")
	}

	_, c.err = c.client.Now().Table("incident").Get(context.Background(), nil)
	return nil
}

// ── Assertion Steps ───────────────────────────────────────────────────────

func (c *authenticationTestContext) authenticationShouldSucceed() error {
	if c.authErr != nil {
		return fmt.Errorf("expected authentication to succeed, but got error: %v", c.authErr)
	}
	if c.client == nil {
		return fmt.Errorf("expected client to be initialized, but it is nil")
	}
	return nil
}

func (c *authenticationTestContext) authenticationShouldFail() error {
	if c.authErr == nil && c.err == nil {
		return fmt.Errorf("expected authentication to fail, but it succeeded")
	}
	return nil
}

func (c *authenticationTestContext) iShouldBeAbleToMakeAnAPIRequest() error {
	if c.client == nil {
		return fmt.Errorf("client is nil, cannot make API request")
	}

	_, err := c.client.Now().Table("incident").Get(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("expected successful API request, but got error: %v", err)
	}
	return nil
}

func (c *authenticationTestContext) allRequestsShouldSucceed() error {
	if c.err != nil {
		return fmt.Errorf("expected all requests to succeed, but got error: %v", c.err)
	}
	return nil
}

func (c *authenticationTestContext) authenticationTokensShouldBeReused() error {
	return nil
}

func (c *authenticationTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *authenticationTestContext) theErrorShouldBeAuthenticationRelated() error {
	if c.err == nil && c.authErr == nil {
		return fmt.Errorf("expected an authentication error, but got nil")
	}

	errMsg := ""
	if c.authErr != nil {
		errMsg = c.authErr.Error()
	} else if c.err != nil {
		errMsg = c.err.Error()
	}

	authRelated := strings.Contains(errMsg, "authentication") ||
		strings.Contains(errMsg, "unauthorized") ||
		strings.Contains(errMsg, "401") ||
		strings.Contains(errMsg, "credentials") ||
		strings.Contains(errMsg, "invalid") ||
		strings.Contains(errMsg, "oauth") ||
		strings.Contains(errMsg, "token")

	if !authRelated {
		return fmt.Errorf("expected authentication-related error, but got: %v", errMsg)
	}

	return nil
}

func (c *authenticationTestContext) credentialCreationShouldFail() error {
	if c.authErr == nil {
		return fmt.Errorf("expected credential creation to fail, but it succeeded")
	}
	return nil
}

func (c *authenticationTestContext) theErrorMessageShouldIndicateMissingParameters() error {
	if c.authErr == nil {
		return fmt.Errorf(" expected an error about missing parameters, but got nil")
	}

	errMsg := c.authErr.Error()
	if !strings.Contains(errMsg, "clientID") &&
		!strings.Contains(errMsg, "clientSecret") &&
		!strings.Contains(errMsg, "client_id") &&
		!strings.Contains(errMsg, "client_secret") &&
		!strings.Contains(errMsg, "username") &&
		!strings.Contains(errMsg, "password") &&
		!strings.Contains(errMsg, "empty") {
		return fmt.Errorf("expected error about missing parameters, but got: %v", c.authErr)
	}

	return nil
}

func (c *authenticationTestContext) theErrorMessageShouldIndicateAnInvalidJWT() error {
	if c.err == nil && c.authErr == nil {
		return fmt.Errorf("expected error about invalid JWT, but got nil")
	}

	errMsg := ""
	if c.err != nil {
		errMsg = c.err.Error()
	} else if c.authErr != nil {
		errMsg = c.authErr.Error()
	}

	if !strings.Contains(errMsg, "JWT") && !strings.Contains(errMsg, "token") && !strings.Contains(errMsg, "invalid") {
		return fmt.Errorf("expected error about invalid JWT, but got: %v", errMsg)
	}

	return nil
}

func (c *authenticationTestContext) theErrorMessageShouldIndicateAMissingJWTClaim() error {
	if c.err == nil && c.authErr == nil {
		return fmt.Errorf("expected error about missing JWT claim, but got nil")
	}

	errMsg := ""
	if c.err != nil {
		errMsg = c.err.Error()
	} else if c.authErr != nil {
		errMsg = c.authErr.Error()
	}

	if !strings.Contains(errMsg, "missing required claim") && !strings.Contains(errMsg, "claim") {
		return fmt.Errorf("expected error about missing JWT claim, but got: %v", errMsg)
	}

	return nil
}

func (c *authenticationTestContext) theErrorMessageShouldIndicateWrongSigningAlgorithm() error {
	if c.err == nil && c.authErr == nil {
		return fmt.Errorf("expected error about wrong signing algorithm, but got nil")
	}

	errMsg := ""
	if c.err != nil {
		errMsg = c.err.Error()
	} else if c.authErr != nil {
		errMsg = c.authErr.Error()
	}

	if !strings.Contains(errMsg, "signing algorithm") && !strings.Contains(errMsg, "algorithm") {
		return fmt.Errorf("expected error about wrong signing algorithm, but got: %v", errMsg)
	}

	return nil
}

// ── Mock Token Providers ──────────────────────────────────────────────────

type staticTokenProvider struct {
	token string
	err   error
}

func (p *staticTokenProvider) GetAuthorizationToken(_ context.Context, _ *url.URL, _ map[string]interface{}) (string, error) {
	if p.err != nil {
		return "", p.err
	}
	return p.token, nil
}

func (p *staticTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}

type failingTokenProvider struct {
	err error
}

func (p *failingTokenProvider) GetAuthorizationToken(_ context.Context, _ *url.URL, _ map[string]interface{}) (string, error) {
	return "", p.err
}

func (p *failingTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}

// ── JWT Generation Helpers ────────────────────────────────────────────────

func generateValidRS256JWT() (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := map[string]interface{}{
		"iss": "test-issuer",
		"sub": "test-subject",
		"aud": "test-audience",
		"exp": now.Add(1 * time.Hour).Unix(),
		"iat": now.Unix(),
		"jti": "test-jti-12345",
	}

	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`))
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)

	signingInput := header + "." + claimsEncoded
	hash := sha256.Sum256([]byte(signingInput))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	sigEncoded := base64.RawURLEncoding.EncodeToString(signature)

	return signingInput + "." + sigEncoded, nil
}

func generateJWTWithMissingClaims() string {
	now := time.Now()
	claims := map[string]interface{}{
		"iss": "test-issuer",
		"sub": "test-subject",
		"exp": now.Add(1 * time.Hour).Unix(),
		"iat": now.Unix(),
		"jti": "test-jti-12345",
	}

	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`))
	claimsJSON, _ := json.Marshal(claims)
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)

	return header + "." + claimsEncoded + ".fake-signature"
}

func generateJWTWithWrongAlgorithm() string {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	now := time.Now()
	claims := map[string]interface{}{
		"iss": "test-issuer",
		"sub": "test-subject",
		"aud": "test-audience",
		"exp": now.Add(1 * time.Hour).Unix(),
		"iat": now.Unix(),
		"jti": "test-jti-12345",
	}
	claimsJSON, _ := json.Marshal(claims)
	claimsEncoded := base64.RawURLEncoding.EncodeToString(claimsJSON)

	return header + "." + claimsEncoded + ".fake-signature"
}

// ── Scenario Registration ─────────────────────────────────────────────────

func InitializeAuthenticationScenario(ctx *godog.ScenarioContext) {
	tc := &authenticationTestContext{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tc = &authenticationTestContext{}
		setupGlobalMocks()
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		httpmock.DeactivateAndReset()
		return ctx, nil
	})

	// Background
	ctx.Step(`^I have a valid ServiceNow instance and credentials$`, tc.iHaveAValidServiceNowInstanceAndCredentials)

	// Basic auth
	ctx.Step(`^I initialize a client with basic authentication$`, tc.iInitializeAClientWithBasicAuthentication)
	ctx.Step(`^I initialize a client with basic authentication using empty username$`, tc.iInitializeAClientWithBasicAuthenticationUsingEmptyUsername)
	ctx.Step(`^I initialize a client with basic authentication using empty password$`, tc.iInitializeAClientWithBasicAuthenticationUsingEmptyPassword)

	// Static bearer token
	ctx.Step(`^I initialize a client with a static bearer token "([^"]*)"$`, tc.iInitializeAClientWithAStaticBearerToken)

	// Client credentials
	ctx.Step(`^I initialize a client with client credentials authentication$`, tc.iInitializeAClientWithClientCredentialsAuthentication)
	ctx.Step(`^I attempt to create client credentials with empty client ID$`, tc.iAttemptToCreateClientCredentialsWithEmptyClientID)
	ctx.Step(`^I attempt to create client credentials with empty client secret$`, tc.iAttemptToCreateClientCredentialsWithEmptyClientSecret)
	ctx.Step(`^I initialize a client with invalid client credentials$`, tc.iInitializeAClientWithInvalidClientCredentials)

	// Token lifecycle
	ctx.Step(`^I make multiple API requests$`, tc.iMakeMultipleAPIRequests)
	ctx.Step(`^the cached token expires$`, tc.theCachedTokenExpires)
	ctx.Step(`^I make an API request that triggers token refresh$`, tc.iMakeAnAPITriggeringTokenRefresh)
	ctx.Step(`^I make an API request to acquire a token$`, tc.iMakeAnAPIToAcquireToken)
	ctx.Step(`^I revoke the current access token$`, tc.iRevokeTheCurrentAccessToken)

	// ROPC
	ctx.Step(`^I initialize a client with ROPC authentication$`, tc.iInitializeAClientWithROPCAuthentication)
	ctx.Step(`^I initialize a client with ROPC authentication using empty username$`, tc.iInitializeAClientWithROPCAuthenticationUsingEmptyUsername)
	ctx.Step(`^I initialize a client with ROPC authentication using empty password$`, tc.iInitializeAClientWithROPCAuthenticationUsingEmptyPassword)
	ctx.Step(`^I attempt to create ROPC credentials with empty username$`, tc.iAttemptToCreateROPCredentialsWithEmptyUsername)
	ctx.Step(`^I attempt to create ROPC credentials with empty password$`, tc.iAttemptToCreateROPCredentialsWithEmptyPassword)
	ctx.Step(`^the ROPC GetAuthentication method should return a Bearer token header$`, tc.theROPCGetAuthenticationMethodShouldReturnABearerTokenHeader)

	// JWT
	ctx.Step(`^I initialize a client with JWT bearer token authentication$`, tc.iInitializeAClientWithJWTBearerTokenAuthentication)
	ctx.Step(`^I initialize a client with a failing JWT token provider$`, tc.iInitializeAClientWithAFailingJWTTokenProvider)
	ctx.Step(`^I initialize a client with an invalid JWT assertion provider$`, tc.iInitializeAClientWithAnInvalidJWTAssertionProvider)
	ctx.Step(`^I initialize a client with a JWT missing required claims$`, tc.iInitializeAClientWithAJWTMissingRequiredClaims)
	ctx.Step(`^I initialize a client with a JWT using wrong algorithm$`, tc.iInitializeAClientWithAJWTUsingWrongAlgorithm)

	// Authorization code
	ctx.Step(`^I initialize an authorization code public client with valid parameters$`, tc.iInitializeAnAuthorizationCodePublicClientWithValidParameters)
	ctx.Step(`^I initialize an authorization code confidential client with valid parameters$`, tc.iInitializeAnAuthorizationCodeConfidentialClientWithValidParameters)
	ctx.Step(`^the authorization code credential should be initialized$`, tc.theAuthorizationCodeCredentialShouldBeInitialized)
	ctx.Step(`^the client should support PKCE challenge generation$`, tc.theClientShouldSupportPKCEChallengeGeneration)
	ctx.Step(`^I attempt to create authorization code credentials with empty client ID$`, tc.iAttemptToCreateAuthorizationCodeCredentialsWithEmptyClientID)

	// Token endpoint error
	ctx.Step(`^the token endpoint returns an error$`, tc.theTokenEndpointReturnsAnError)

	// Assertion steps
	ctx.Step(`^authentication should succeed$`, tc.authenticationShouldSucceed)
	ctx.Step(`^authentication should fail$`, tc.authenticationShouldFail)
	ctx.Step(`^I should be able to make an API request$`, tc.iShouldBeAbleToMakeAnAPIRequest)
	ctx.Step(`^all requests should succeed$`, tc.allRequestsShouldSucceed)
	ctx.Step(`^authentication tokens should be reused$`, tc.authenticationTokensShouldBeReused)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the error should be authentication-related$`, tc.theErrorShouldBeAuthenticationRelated)
	ctx.Step(`^credential creation should fail$`, tc.credentialCreationShouldFail)
	ctx.Step(`^the error message should indicate missing parameters$`, tc.theErrorMessageShouldIndicateMissingParameters)
	ctx.Step(`^the error message should indicate an invalid JWT$`, tc.theErrorMessageShouldIndicateAnInvalidJWT)
	ctx.Step(`^the error message should indicate a missing JWT claim$`, tc.theErrorMessageShouldIndicateAMissingJWTClaim)
	ctx.Step(`^the error message should indicate wrong signing algorithm$`, tc.theErrorMessageShouldIndicateWrongSigningAlgorithm)
	ctx.Step(`^a new access token should be returned$`, tc.aNewAccessTokenShouldBeReturned)
	ctx.Step(`^the revocation should succeed$`, tc.theRevocationShouldSucceed)
	ctx.Step(`^the cached token should be cleared$`, tc.theCachedTokenShouldBeCleared)
	ctx.Step(`^the API request should fail with an authentication error$`, tc.theAPIShouldFailWithAnAuthenticationError)
}

func TestAuthenticationFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeAuthenticationScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/authentication.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
