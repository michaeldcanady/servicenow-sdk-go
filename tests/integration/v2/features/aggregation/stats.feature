@integration @mock @aggregation
Feature: ServiceNow Aggregation (Stats) API
  As a developer using the ServiceNow SDK
  I want to retrieve aggregate statistics for tables
  So that I can analyze record data

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: Count aggregation
    When I request a count aggregation for the "incident" table
    Then the aggregation result should not be nil

  @happy @offline
  Scenario: Query aggregation
    When I request a query aggregation for the "incident" table with query "priority=1"
    Then the aggregation result should not be nil

  @error
  Scenario: Aggregation without parameters returns error
    When I request an aggregation without parameters for "incident"
    Then the response should be an API error

  @error
  Scenario: Aggregation for invalid table returns error
    When I request a count aggregation for the "this_table_does_not_exist" table
    Then the response should be an API error
