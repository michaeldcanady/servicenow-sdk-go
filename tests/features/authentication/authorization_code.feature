Feature: Authorization Code Authentication
  As a developer using the ServiceNow SDK
  I want to authenticate using OAuth Authorization Code flow
  So that I can securely interact with the ServiceNow API

  Scenario: Successful authentication with valid authorization code
    Given a client is available
    And a valid instance
    When correct credentials are supplied
    And a request is sent
    Then the response should be successful

  Scenario: Authentication server times out
    Given a client is available
    And a valid instance
    When authentication flow is cancelled
    And a request is sent
    Then an authentication error message is shown
