@integration @mock @appservice @api
Feature: ServiceNow Application Service API
  As a developer using the ServiceNow SDK
  I want to create, retrieve content, and find application services
  So that I can manage my application service graph

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @appservice @lifecycle
  Scenario: Successfully manage an application service lifecycle
    When I create an application service named "My App Service" with comments "Created via BDD integration test"
    Then the response should not be an error
    And the created service should have a valid sys_id
    When I request the content of the application service
    Then the response should not be an error
    And the content response should not be nil
    When I search for an application service by name "My App Service"
    Then the response should not be an error
    And the search results should contain at least 1 service named "My App Service"
