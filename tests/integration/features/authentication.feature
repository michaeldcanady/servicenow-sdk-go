@integration @authentication @credentials
Feature: ServiceNow SDK Credential Authentication
  As a developer using the ServiceNow SDK
  I want to authenticate using various credential types
  So that I can securely access ServiceNow APIs with different authentication methods

  Background:
    And I have a valid ServiceNow instance and credentials

  # ── Basic Authentication ──────────────────────────────────────────────

  @integration @authentication @basic
  Scenario: Basic authentication succeeds and makes an API request
    When I initialize a client with basic authentication
    Then authentication should succeed
    And I should be able to make an API request

  @integration @authentication @basic @negative
  Scenario: Basic authentication fails with empty username
    When I initialize a client with basic authentication using empty username
    Then the API request should fail with an authentication error

  @integration @authentication @basic @negative
  Scenario: Basic authentication fails with empty password
    When I initialize a client with basic authentication using empty password
    Then the API request should fail with an authentication error

  # ── Static Bearer Token ───────────────────────────────────────────────

  @integration @authentication @bearer
  Scenario: Static bearer token succeeds and makes an API request
    When I initialize a client with a static bearer token "valid-token-abc123"
    Then authentication should succeed
    And I should be able to make an API request

  @integration @authentication @bearer @negative
  Scenario: Static bearer token provider rejects empty token
    When I initialize a client with a static bearer token ""
    Then authentication should fail
    And the error message should indicate missing parameters

  # ── Client Credentials ────────────────────────────────────────────────

  @integration @authentication @client_credentials
  Scenario: Client credentials authentication succeeds and makes an API request
    When I initialize a client with client credentials authentication
    Then authentication should succeed
    And I should be able to make an API request

  @integration @authentication @client_credentials @negative
  Scenario: Client credentials provider rejects empty client ID
    When I attempt to create client credentials with empty client ID
    Then credential creation should fail
    And the error message should indicate missing parameters

  @integration @authentication @client_credentials @negative
  Scenario: Client credentials provider rejects empty client secret
    When I attempt to create client credentials with empty client secret
    Then credential creation should fail
    And the error message should indicate missing parameters

  @integration @authentication @client_credentials @negative
  Scenario: Client credentials authentication fails with invalid credentials
    When I initialize a client with invalid client credentials
    Then authentication should fail
    And the error should be authentication-related

  @integration @authentication @client_credentials @token_lifecycle
  Scenario: Client credentials token is cached across multiple requests
    When I initialize a client with client credentials authentication
    And I make multiple API requests
    Then all requests should succeed
    And authentication tokens should be reused

  @integration @authentication @client_credentials @token_lifecycle
  Scenario: Client credentials token is refreshed when expired
    When I initialize a client with client credentials authentication
    And the cached token expires
    And I make an API request that triggers token refresh
    Then the response should not be an error
    And a new access token should be returned

  @integration @authentication @client_credentials @token_lifecycle
  Scenario: Client credentials token can be revoked
    When I initialize a client with client credentials authentication
    And I make an API request to acquire a token
    And I revoke the current access token
    Then the revocation should succeed
    And the cached token should be cleared

  # ── Resource Owner Password Credentials (ROPC) ───────────────────────

  @integration @authentication @ropc
  Scenario: ROPC authentication succeeds and makes an API request
    When I initialize a client with ROPC authentication
    Then authentication should succeed
    And I should be able to make an API request

  @integration @authentication @ropc @negative
  Scenario: ROPC authentication fails with empty username
    When I initialize a client with ROPC authentication using empty username
    Then the API request should fail with an authentication error

  @integration @authentication @ropc @negative
  Scenario: ROPC authentication fails with empty password
    When I initialize a client with ROPC authentication using empty password
    Then the API request should fail with an authentication error

  @integration @authentication @ropc
  Scenario: ROPC GetAuthentication returns a Bearer header
    When I initialize a client with ROPC authentication
    Then the ROPC GetAuthentication method should return a Bearer token header

  # ── JWT Bearer Token ──────────────────────────────────────────────────

  @integration @authentication @jwt
  Scenario: JWT bearer token authentication succeeds
    When I initialize a client with JWT bearer token authentication
    Then authentication should succeed
    And I should be able to make an API request

  @integration @authentication @jwt @negative
  Scenario: JWT bearer token fails when the token provider errors
    When I initialize a client with a failing JWT token provider
    Then authentication should fail
    And the error should be authentication-related

  @integration @authentication @jwt @negative
  Scenario: JWT bearer token fails with invalid JWT assertion
    When I initialize a client with an invalid JWT assertion provider
    Then authentication should fail
    And the error message should indicate an invalid JWT

  @integration @authentication @jwt @negative
  Scenario: JWT bearer token fails with missing required claims
    When I initialize a client with a JWT missing required claims
    Then authentication should fail
    And the error message should indicate a missing JWT claim

  @integration @authentication @jwt @negative
  Scenario: JWT bearer token fails with wrong signing algorithm
    When I initialize a client with a JWT using wrong algorithm
    Then authentication should fail
    And the error message should indicate wrong signing algorithm

  # ── Authorization Code ────────────────────────────────────────────────

  @integration @authentication @authorization_code
  Scenario: Authorization code public client initializes with PKCE
    When I initialize an authorization code public client with valid parameters
    Then the authorization code credential should be initialized
    And the client should support PKCE challenge generation

  @integration @authentication @authorization_code
  Scenario: Authorization code confidential client initializes without PKCE
    When I initialize an authorization code confidential client with valid parameters
    Then the authorization code credential should be initialized

  @integration @authentication @authorization_code @negative
  Scenario: Authorization code provider rejects empty client ID
    When I attempt to create authorization code credentials with empty client ID
    Then credential creation should fail
    And the error message should indicate missing parameters

  # ── Cross-cutting: Invalid token endpoint ─────────────────────────────

  @integration @authentication @negative
  Scenario: OAuth2 token endpoint returns an error for invalid credentials
    When I initialize a client with invalid client credentials
    Then authentication should fail
    And the error should be authentication-related
