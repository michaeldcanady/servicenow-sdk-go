Feature: Table API POST Operations
  As a developer using the ServiceNow SDK
  I want to create records in ServiceNow tables
  So that I can add new data to the system

  Background:
    Given a client is available
    And a valid instance
    And correct credentials are supplied

  Scenario: Create a new record
    When I POST to the "incident" table with data:
      | short_description | POST Test Incident |
    And a request is sent
    Then the response should be successful
    And the response should contain the created record
