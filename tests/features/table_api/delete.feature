Feature: Table API DELETE Operations
  As a developer using the ServiceNow SDK
  I want to delete records from ServiceNow tables
  So that I can remove obsolete data

  Background:
    Given a client is available
    And a valid instance
    And correct credentials are supplied

  Scenario: Delete an existing record
    Given a record exists in the "incident" table
    When I DELETE the record from the "incident" table by Sys ID
    And a request is sent
    Then the response should be successful
