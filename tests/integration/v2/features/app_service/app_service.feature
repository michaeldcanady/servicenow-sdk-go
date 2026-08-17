@integration @mock @app_service
Feature: ServiceNow App Service (CDM) API
  As a developer using the ServiceNow SDK
  I want to find, create, and manage application services
  So that I can work with CDM services

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: Find service
    When I find a service with query "test"
    Then the response should not be an error

  @happy
  Scenario: Create service
    When I create a service with name "BDD Test Service"
    Then the response should not be an error

  @happy
  Scenario: Get content of service
    When I create a service with name "Content Test Service"
    Then the response should not be an error
    When I get content of the last service
    Then the response should not be an error

  @error
  Scenario: Invalid service returns error
    When I get content of service "00000000000000000000000000000000"
    Then the response should be an API error
