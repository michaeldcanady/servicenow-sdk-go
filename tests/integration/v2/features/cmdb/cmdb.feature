@integration @mock @cmdb
Feature: ServiceNow CMDB Instance API
  As a developer using the ServiceNow SDK
  I want to query and manage configuration items by class
  So that I can inspect and manipulate CMDB records

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: List CIs for class with limit
    When I request CIs for class "cmdb_ci" with limit 1
    Then the response should not be an error
    And the cmdb results should contain at least 1 record

  @happy
  Scenario: Get CI by sys_id after listing
    Given I have at least 1 CI in class "cmdb_ci"
    When I request the CI by its "sys_id" in class "cmdb_ci"
    Then the response should not be an error
    And the cmdb result should have the correct "sys_id"

  @error
  Scenario: Get CI for invalid class returns error
    When I request CIs for class "this_cmdb_class_does_not_exist" with limit 1
    Then the response should be an API error

  @happy @offline
  Scenario: List CIs with query filter
    When I request CIs for class "cmdb_ci" with query "nameLIKEMock"
    Then the response should not be an error
    And the cmdb results should contain at least 1 record

  @happy @offline
  Scenario: Create CI
    When I create a CI for class "cmdb_ci"
    Then the response should not be an error

  @happy @offline
  Scenario: Update CI using patch
    When I create a CI for class "cmdb_ci"
    Then the response should not be an error
    When I update the last CI for class "cmdb_ci" with name "Updated CI"
    Then the response should not be an error

  @happy @offline
  Scenario: CI lifecycle create update
    When I create a CI for class "cmdb_ci"
    Then the response should not be an error
    When I update the last CI for class "cmdb_ci" with name "Lifecycle CI"
    Then the response should not be an error
