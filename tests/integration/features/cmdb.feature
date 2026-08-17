@integration @mock @cmdb @api
Feature: ServiceNow CMDB Instance API retrieval
  As a developer using the ServiceNow SDK
  I want to query configuration items by class
  So that I can inspect CMDB records

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @cmdb @collection
  Scenario: Successfully fetch CIs for a class
    When I request CIs for class "cmdb_ci" with limit 5
    Then the response should not be an error
    And the cmdb results should contain at least 0 records

  @integration @cmdb @query
  Scenario: Successfully fetch CIs with a limit
    When I request CIs for class "cmdb_ci" with limit 1
    Then the response should not be an error
    And the cmdb results should contain at most 1 records

  @integration @cmdb @empty
  Scenario: Return an empty collection when the query matches no CIs
    When I request CIs for class "cmdb_ci" with query "sys_id=00000000000000000000000000000000"
    Then the response should not be an error
    And the cmdb results should contain exactly 0 records

  @integration @cmdb @item
  Scenario: Successfully fetch a single CI by sys_id
    Given I have at least 1 CI in class "cmdb_ci"
    When I request the CI by its "sys_id" in class "cmdb_ci"
    Then the response should not be an error
    And the cmdb result should have the correct "sys_id"

  @integration @cmdb @error
  Scenario: Return an error for a non-existent CI
    When I request the CI with sys_id "00000000000000000000000000000000" in class "cmdb_ci"
    Then the response should be an API error

  @integration @cmdb @error
  Scenario: Return an error for a non-existent CMDB class
    When I request CIs for class "this_cmdb_class_does_not_exist" with limit 1
    Then the response should be an API error
