// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

//go:build snippets

package snippets

// [START auth_imports]
import (
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/credential"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/credentials"
)

// [END auth_imports]

import "github.com/microsoft/kiota-abstractions-go/authentication"

func _() {
	// [START auth_basic_admin]
	credAdmin := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")
	// [END auth_basic_admin]

	// [START auth_basic_credential_package]
	basicCred, err := credential.NewBasicProvider(credential.BasicConfig{
		Username: "xSDK_USERNAMEx",
		Password: "xSDK_PASSWORDx",
	})
	if err != nil {
		log.Fatal(err)
	}
	// [END auth_basic_credential_package]

	// [START client_init_panic]
	clientPanic, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(credAdmin),
		servicenowsdkgo.WithURL("xSDK_SN_URLx"),
	)
	if err != nil {
		panic(err)
	}
	// [END client_init_panic]

	// [START auth_basic]
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")
	// [END auth_basic]

	// You'll need your Client ID.
	credToken, err := credentials.NewPublicAuthorizationCodeProvider(
		"your-client-id",
	)
	if err != nil {
		log.Fatal(err)
	}

	// [START client_init]
	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithURL("xSDK_SN_URLx"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// [END client_init]

	// [START auth_bearer_credential_package]
	bearerCred, err := credential.NewBearerProvider(credential.BearerConfig{Token: "xSDK_TOKENx"})
	if err != nil {
		log.Fatal(err)
	}
	// [END auth_bearer_credential_package]

	// [START auth_apikey_credential_package]
	apiKeyCred, err := credential.NewAPIKeyProvider(credential.APIKeyConfig{Key: "xSDK_APIKEYx"})
	if err != nil {
		log.Fatal(err)
	}
	// [END auth_apikey_credential_package]

	// [START client_init_credential_package]
	clientV2, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(bearerCred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// [END client_init_credential_package]

	_ = clientPanic
	_ = client
	_ = clientV2
	_ = credToken
	_ = basicCred
	_ = bearerCred
	_ = apiKeyCred
}

func authBasicFull() {
	// [START auth_basic_full]
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_basic_full]
}

func authBasicCredentialPackage() {
	// [START auth_basic_full_credential_package]
	cred, err := credential.NewBasicProvider(credential.BasicConfig{
		Username: "xSDK_USERNAMEx",
		Password: "xSDK_PASSWORDx",
	})
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_basic_full_credential_package]
}

func authROPC() {
	// [START auth_ropc]
	cred, err := credentials.NewROPCProvider(
		"{clientID}",
		"{clientSecret}",
		"xSDK_USERNAMEx",
		"xSDK_PASSWORDx",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_ropc]
}

func authClientCredentials() {
	// [START auth_client_credentials]
	cred, err := credentials.NewClientCredentialsProvider(
		"{clientID}",
		"{clientSecret}",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_client_credentials]
}

func authClientCredentialsCredentialPackage() {
	// [START auth_client_credentials_credential_package]
	cred, err := credential.NewClientCredentialsProvider(credential.ClientCredentialsConfig{
		ClientID:     "{clientID}",
		ClientSecret: "{clientSecret}",
	})
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_client_credentials_credential_package]
}

func authCodePrivate() {
	// [START auth_code_private]
	cred, err := credentials.NewPrivateAuthorizationCodeProvider(
		"{clientID}",
		"{clientSecret}",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_code_private]
}

func authCodePublic() {
	// [START auth_code_public]
	cred, err := credentials.NewPublicAuthorizationCodeProvider(
		"{clientID}",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_code_public]
}

func authExternalOIDC() {
	// [START auth_external_oidc]
	provider, err := credentials.NewExternalOIDCProvider(
		"{clientID}",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
		credentials.WithOIDCIssuer("https://login.microsoftonline.com/{tenantID}/v2.0"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(provider),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_external_oidc]
}

func authExternalOIDCScopes() {
	// [START auth_external_oidc_scopes]
	provider, err := credentials.NewExternalOIDCProvider(
		"{clientID}",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
		credentials.WithOIDCIssuer("https://login.microsoftonline.com/{tenantID}/v2.0"),
		credentials.WithScopes("api://{clientID}/.default"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// [END auth_external_oidc_scopes]

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(provider),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
}

func authJWT() {
	// tokenProvider must generate signed JWT assertions; it is user-provided
	// and must satisfy kiota's authentication.AccessTokenProvider.
	var tokenProvider authentication.AccessTokenProvider

	// [START auth_jwt]
	cred, err := credentials.NewJWTProvider(
		"{clientID}",
		"{clientSecret}",
		tokenProvider,
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_jwt]
}

func _() {
	authBasicFull()
	authROPC()
	authClientCredentials()
	authCodePrivate()
	authCodePublic()
	authExternalOIDC()
	authExternalOIDCScopes()
	authJWT()
}
