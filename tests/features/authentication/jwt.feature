Feature: JWT Authentication
  As a developer using the ServiceNow SDK
  I want to authenticate using a JSON Web Token (JWT)
  So that I can securely interact with the ServiceNow API

  Scenario: Successful authentication with valid JWT
    Given a client is available
    And a valid instance
    When correct credentials are supplied
    And a request is sent
    Then the response should be successful

  Scenario: Unsuccessful authentication with invalid JWT
    Given a client is available
    And a valid instance
    When incorrect credentials are supplied
    And a request is sent
    Then an authentication error message is shown
