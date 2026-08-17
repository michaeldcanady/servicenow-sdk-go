@integration @mock @aggregation @stats
Feature: ServiceNow Stats API aggregate retrieval
  As a developer using the ServiceNow SDK
  I want to retrieve aggregate statistics for a table
  So that I can count and sum records without fetching every row

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @aggregation @stats @count
  Scenario: Successfully retrieve a record count for the incident table
    When I request the record count for the "incident" table
    Then the response should not be an error
    And the stats count should be present

  @integration @aggregation @stats @query
  Scenario: Successfully retrieve a filtered record count
    When I request the record count for the "incident" table with query "active=true"
    Then the response should not be an error
    And the stats count should be present

  @integration @aggregation @stats @sum
  Scenario: Successfully retrieve sum aggregates for a numeric field
    When I request the sum of "reassignment_count" for the "incident" table
    Then the response should not be an error
    And the stats sum for "reassignment_count" should be present

  @integration @aggregation @stats @empty
  Scenario: Return a zero count when the query matches no records
    When I request the record count for the "incident" table with query "sys_id=00000000000000000000000000000000"
    Then the response should not be an error
    And the stats count should be "0"

  @integration @aggregation @stats @error
  Scenario: Return an error for a non-existent table
    When I request the record count for the "this_table_does_not_exist" table
    Then the response should be an API error

  @integration @aggregation @stats @curveball
  Scenario: Return an error when no aggregate parameters are requested
    When I request stats for the "incident" table with no aggregate parameters
    Then the response should be an API error
