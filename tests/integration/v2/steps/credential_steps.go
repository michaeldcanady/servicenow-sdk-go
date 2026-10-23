// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

//go:build integration

package steps

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/credential"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type recordingCredentialProvider struct {
	baseURL string
}

func (p *recordingCredentialProvider) AuthenticateRequest(context.Context, *abstractions.RequestInformation, map[string]interface{}) error {
	return nil
}

func (p *recordingCredentialProvider) Initialize(baseURL string) error {
	p.baseURL = baseURL
	return nil
}

func registerCredentialSteps(sc *godog.ScenarioContext) {
	sc.Step(`^I configure a new client credentials provider for the test instance$`, iConfigureANewClientCredentialsProvider)
	sc.Step(`^I configure a recording credential provider$`, iConfigureARecordingCredentialProvider)
	sc.Step(`^I initialize a client with the new credential provider$`, iInitializeAClientWithTheNewCredentialProvider)
	sc.Step(`^the credential provider should be initialized with the test instance$`, theCredentialProviderShouldBeInitializedWithTheTestInstance)
	sc.Step(`^I create a new basic credential provider$`, iCreateANewBasicCredentialProvider)
	sc.Step(`^I create a new API key credential provider$`, iCreateANewAPIKeyCredentialProvider)
	sc.Step(`^I authenticate a request to "([^"]*)"$`, iAuthenticateARequestTo)
	sc.Step(`^the request should carry a header starting with "([^"]*)"$`, theRequestShouldCarryAHeaderStartingWith)
	sc.Step(`^the request should carry the header "([^"]*)" with value "([^"]*)"$`, theRequestShouldCarryTheHeaderWithValue)
}

func iConfigureANewClientCredentialsProvider(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	clientID, clientSecret := support.ClientCredentials()

	config := credential.ClientCredentialsConfig{
		ProviderConfig: credential.ProviderConfig{
			HTTPClient:   support.GetOAuth2HTTPClient(),
			AllowedHosts: []string{fmt.Sprintf("%s.service-now.com", support.IntegrationInstance())},
		},
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	provider, err := credential.NewClientCredentialsProvider(config)
	w.AuthErr = err
	if err == nil {
		w.CredentialProvider = provider
	}
	return ctx, nil
}

func iConfigureARecordingCredentialProvider(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	w.CredentialProvider = &recordingCredentialProvider{}
	return ctx, nil
}

func iInitializeAClientWithTheNewCredentialProvider(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr != nil {
		return ctx, w.AuthErr
	}
	if w.CredentialProvider == nil {
		return ctx, fmt.Errorf("credential provider is nil")
	}

	w.AuthErr = support.NewClientWithAuthProvider(w, w.CredentialProvider)
	return ctx, nil
}

func theCredentialProviderShouldBeInitializedWithTheTestInstance(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	provider, ok := w.CredentialProvider.(*recordingCredentialProvider)
	if !ok {
		return ctx, fmt.Errorf("expected a recording credential provider")
	}

	expected := fmt.Sprintf("https://%s.service-now.com", support.IntegrationInstance())
	if provider.baseURL != expected {
		return ctx, fmt.Errorf("expected provider base URL %q, got %q", expected, provider.baseURL)
	}
	return ctx, nil
}

func iCreateANewBasicCredentialProvider(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	username := support.FirstEnv("SN_USERNAME", "SNOW_USERNAME")
	password := support.FirstEnv("SN_PASSWORD", "SNOW_PASSWORD")
	if support.IsOffline() {
		if username == "" {
			username = "mock"
		}
		if password == "" {
			password = "mock"
		}
	}

	provider, err := credential.NewBasicProvider(credential.BasicConfig{
		Username: username,
		Password: password,
	})
	w.AuthErr = err
	if err == nil {
		w.CredentialProvider = provider
	}
	return ctx, nil
}

func iCreateANewAPIKeyCredentialProvider(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	provider, err := credential.NewAPIKeyProvider(credential.APIKeyConfig{Key: "integration-key"})
	w.AuthErr = err
	if err == nil {
		w.CredentialProvider = provider
	}
	return ctx, nil
}

func iAuthenticateARequestTo(ctx context.Context, rawURL string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.AuthErr != nil {
		return ctx, w.AuthErr
	}
	if w.CredentialProvider == nil {
		return ctx, fmt.Errorf("credential provider is nil")
	}

	uri, err := url.Parse(rawURL)
	if err != nil {
		return ctx, err
	}

	request := abstractions.NewRequestInformation()
	request.SetUri(*uri)
	w.CredentialRequest = request
	w.AuthErr = w.CredentialProvider.AuthenticateRequest(ctx, request, nil)
	return ctx, nil
}

func theRequestShouldCarryAHeaderStartingWith(ctx context.Context, prefix string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.CredentialRequest == nil || w.CredentialRequest.Headers == nil {
		return ctx, fmt.Errorf("credential request has no headers")
	}

	for _, value := range w.CredentialRequest.Headers.Get("Authorization") {
		if strings.HasPrefix(value, prefix) {
			return ctx, nil
		}
	}
	return ctx, fmt.Errorf("expected Authorization header to start with %q, got %v", prefix, w.CredentialRequest.Headers.Get("Authorization"))
}

func theRequestShouldCarryTheHeaderWithValue(ctx context.Context, name, expected string) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.CredentialRequest == nil || w.CredentialRequest.Headers == nil {
		return ctx, fmt.Errorf("credential request has no headers")
	}

	values := w.CredentialRequest.Headers.Get(name)
	for _, value := range values {
		if value == expected {
			return ctx, nil
		}
	}
	return ctx, fmt.Errorf("expected header %q to equal %q, got %v", name, expected, values)
}
