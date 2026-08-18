@integration @mock @account
Feature: ServiceNow Account API
  As a developer using the ServiceNow SDK
  I want to list and fetch customer accounts
  So that I can work with CSM account records

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy @offline
  Scenario: List all accounts
    When I request all accounts
    Then the response should not be an error
    And the account results should contain at least 1 record

  @happy @offline
  Scenario: Get account by sys_id after listing
    Given I have at least 1 account in the instance
    When I request the account by its "sys_id"
    Then the response should not be an error
    And the account result should have the correct "sys_id"

  @happy @offline
  Scenario: Get account with query filter
    When I request accounts with query "nameLIKEBDD"
    Then the response should not be an error
    And the account results should contain at least 1 record

  @error
  Scenario: Get non-existent account returns error
    When I request the account with sys_id "00000000000000000000000000000000"
    Then the response should be an API error

  @happy @offline
  Scenario: Account limit parameter
    When I request accounts with limit 1
    Then the response should not be an error
    And the account results should contain at least 1 record
