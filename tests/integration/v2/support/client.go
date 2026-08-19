package support

import (
	"fmt"
	"net/http"

	sdk "github.com/michaeldcanady/servicenow-sdk-go/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/credentials"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	nethttplibrary "github.com/microsoft/kiota-http-go"

	"github.com/jarcoal/httpmock"
)

// NewSDKClient creates a ServiceNowServiceClient with the given options.
// It automatically injects the HTTP client (httpmock when offline) and instance.
func NewSDKClient(w *World, opts ...sdk.ServiceNowServiceClientOption) error {
	instance := IntegrationInstance()
	httpClient := GetHTTPClient()

	baseOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithInstance(instance),
	}
	if httpClient != nil {
		baseOpts = append(baseOpts, sdk.WithHTTPClient(httpClient))
	}
	baseOpts = append(baseOpts, opts...)

	client, err := sdk.NewServiceNowServiceClient(baseOpts...)
	if err != nil {
		return err
	}
	w.Client = client
	return nil
}

// NewClientWithAuthProvider creates a client using the given authentication provider.
func NewClientWithAuthProvider(w *World, ap authentication.AuthenticationProvider) error {
	return NewSDKClient(w, sdk.WithAuthenticationProvider(ap))
}

// NewClientWithBasicAuth creates a client using basic (username/password) auth.
func NewClientWithBasicAuth(w *World) error {
	ap, err := newBasicAuthProvider()
	if err != nil {
		return err
	}
	return NewClientWithAuthProvider(w, ap)
}

// NewClientWithClientCredentials creates a client using OAuth2 client credentials.
func NewClientWithClientCredentials(w *World, clientID, clientSecret string) error {
	oauth2HTTP := GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	cc, err := credentials.NewClientCredentialsProvider(clientID, clientSecret, authOpts...)
	if err != nil {
		return err
	}
	return NewClientWithAuthProvider(w, cc)
}

// NewClientWithROPC creates a client using OAuth2 ROPC.
func NewClientWithROPC(w *World, clientID, clientSecret, username, password string) error {
	oauth2HTTP := GetOAuth2HTTPClient()
	authOpts := []credentials.AuthOption{
		credentials.WithInstance(IntegrationInstance()),
	}
	if oauth2HTTP != nil {
		authOpts = append(authOpts, credentials.WithHTTPClient(oauth2HTTP))
	}

	ropc, err := credentials.NewROPCProvider(clientID, clientSecret, username, password, authOpts...)
	if err != nil {
		return err
	}
	return NewClientWithAuthProvider(w, ropc)
}

// GetHTTPClient returns an *http.Client that uses httpmock transport (when offline)
// wrapped with Kiota middleware, suitable for SDK API calls.
func GetHTTPClient() *http.Client {
	if !IsOffline() {
		return nil
	}
	transport := nethttplibrary.NewCustomTransportWithParentTransport(
		httpmock.DefaultTransport, nethttplibrary.GetDefaultMiddlewares()...,
	)
	return &http.Client{Transport: transport}
}

// GetOAuth2HTTPClient returns a plain http.Client using httpmock transport
// without Kiota middleware, suitable for OAuth2 token endpoint requests.
func GetOAuth2HTTPClient() *http.Client {
	if !IsOffline() {
		return nil
	}
	return &http.Client{
		Transport: httpmock.DefaultTransport,
	}
}

// newBasicAuthProvider creates a basic auth provider from environment variables.
// When online, returns an error if credentials are missing.
func newBasicAuthProvider() (*credentials.BasicAuthenticationProvider, error) {
	username := FirstEnv("SN_USERNAME", "SNOW_USERNAME")
	password := FirstEnv("SN_PASSWORD", "SNOW_PASSWORD")

	if !IsOffline() {
		if username == "" {
			return nil, fmt.Errorf("online mode requires SN_USERNAME or SNOW_USERNAME environment variable")
		}
		if password == "" {
			return nil, fmt.Errorf("online mode requires SN_PASSWORD or SNOW_PASSWORD environment variable")
		}
	} else {
		if username == "" {
			username = "mock"
		}
		if password == "" {
			password = "mock"
		}
	}
	return credentials.NewBasicProvider(username, password), nil
}
