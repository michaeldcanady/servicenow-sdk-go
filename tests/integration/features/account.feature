@integration @mock @account @api
Feature: ServiceNow Account API retrieval
  As a developer using the ServiceNow SDK
  I want to list and fetch customer accounts
  So that I can work with CSM account records

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @account @collection
  Scenario: Successfully fetch the accounts collection
    When I request all accounts
    Then the response should not be an error
    And the account results should contain at least 0 records

  @integration @account @query
  Scenario: Successfully fetch accounts with a limit
    When I request accounts with limit 1
    Then the response should not be an error
    And the account results should contain at most 1 records

  @integration @account @empty
  Scenario: Return an empty collection when the query matches no accounts
    When I request accounts with query "sys_id=00000000000000000000000000000000"
    Then the response should not be an error
    And the account results should contain exactly 0 records

  @integration @account @item
  Scenario: Successfully fetch a single account by sys_id
    Given I have at least 1 account in the instance
    When I request the account by its "sys_id"
    Then the response should not be an error
    And the account result should have the correct "sys_id"

  @integration @account @error
  Scenario: Return an error for a non-existent account
    When I request the account with sys_id "00000000000000000000000000000000"
    Then the response should be an API error
