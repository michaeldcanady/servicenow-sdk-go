Feature: Table API PUT Operations
  As a developer using the ServiceNow SDK
  I want to replace records in ServiceNow tables
  So that I can update data completely

  Background:
    Given a client is available
    And a valid instance
    And correct credentials are supplied

  Scenario: Replace an existing record
    Given a record exists in the "incident" table
    When I PUT to the "incident" table by Sys ID with data:
      | short_description | PUT Replaced Incident |
    And a request is sent
    Then the response should be successful
    And the response should contain the updated record
