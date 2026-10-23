@credential
Feature: Credential package providers
  The credential package provides typed authentication providers. Flow configs
  never contain a ServiceNow URL; the client supplies the instance URL.

  Scenario: Client credentials authenticate an API request
    Given I configure a new client credentials provider for the test instance
    When I initialize a client with the new credential provider
    Then authentication should succeed
    And I should be able to make an API request

  Scenario: A basic credential provider sets an authorization header
    Given I create a new basic credential provider
    When I authenticate a request to "https://mock_instance.service-now.com/api/now"
    Then the request should carry a header starting with "Basic "

  Scenario: An API key credential provider uses the default header
    Given I create a new API key credential provider
    When I authenticate a request to "https://mock_instance.service-now.com/api/now"
    Then the request should carry the header "x-sn-apikey" with value "integration-key"

  Scenario: The client supplies the instance URL to a credential provider
    Given I configure a new client credentials provider for the test instance
    When I initialize a client with the new credential provider
    Then the credential provider should be initialized with the test instance
