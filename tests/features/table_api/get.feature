Feature: Table API GET Operations
  As a developer using the ServiceNow SDK
  I want to retrieve records from ServiceNow tables
  So that I can read data with specific filters and constraints

  Background:
    Given a client is available
    And a valid instance
    And correct credentials are supplied

  Scenario: Retrieve a collection of records
    When I GET a collection from the "incident" table
    And a request is sent
    Then the response should be successful
    And the response should contain a list of records

  Scenario: Retrieve a collection with a limit
    When I GET a collection from the "incident" table with parameters:
      | sysparm_limit | 1 |
    And a request is sent
    Then the response should be successful
    And the response should contain 1 records

  Scenario: Retrieve a specific record by Sys ID
    Given a record exists in the "incident" table
    When I GET the record from the "incident" table by Sys ID
    And a request is sent
    Then the response should be successful
    And the response should contain the requested record
