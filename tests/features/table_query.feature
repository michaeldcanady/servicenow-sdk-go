Feature: ServiceNow Table API Filtering and Sorting
  As a developer using the ServiceNow SDK
  I want to be able to filter and sort records
  So that I can retrieve specific datasets efficiently

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @table @query
  Scenario: Filter incidents by active state and limit results
    When I request incidents with query "active=true" and limit 2
    Then the response should not be an error
    And the results should contain at most 2 records
    And each record should have "active" set to "true"

  @integration @table @query
  Scenario: Sort incidents by created date descending
    When I request incidents sorted by "sys_created_on" descending
    Then the response should not be an error
    And the records should be in descending order of "sys_created_on"
