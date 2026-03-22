Feature: Table API PATCH Operations
  As a developer using the ServiceNow SDK
  I want to partially update records in ServiceNow tables
  So that I can modify specific fields without replacing the entire record

  Background:
    Given a client is available
    And a valid instance
    And correct credentials are supplied

  Scenario: Partially update an existing record
    Given a record exists in the "incident" table
    When I PATCH the "incident" table by Sys ID with data:
      | short_description | PATCH Modified Incident |
    And a request is sent
    Then the response should be successful
    And the response should contain the updated record
