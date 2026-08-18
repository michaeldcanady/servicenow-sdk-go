@integration @mock @caseapi
Feature: ServiceNow Customer Service Case API
  As a developer using the ServiceNow SDK
  I want to create, retrieve, update, and manage cases
  So that I can manage customer service cases programmatically

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: List all cases
    When I list all cases
    Then the response should not be an error
    And the case collection should contain at least 1 case

  @happy
  Scenario: Create case
    When I create a case with short description "BDD Test Case" and description "Test Description"
    Then the response should not be an error
    And the created case should have a valid sys_id

  @happy
  Scenario: Retrieve case by sys_id after create
    When I create a case with short description "BDD Test Case" and description "Retrieve desc"
    Then the response should not be an error
    When I retrieve the case by its sys_id
    Then the response should not be an error
    And the retrieved case should have short description "BDD Test Case"

  @happy
  Scenario: Update case
    When I create a case with short description "Update Test Case" and description "Update desc"
    Then the response should not be an error
    When I update the last case with PUT short description "Updated Case"
    Then the response should not be an error

  @happy
  Scenario: Case activities
    When I list all cases
    Then the response should not be an error
    When I retrieve activities for the last case
    Then the response should not be an error

  @happy
  Scenario: Case field values
    When I list all cases
    Then the response should not be an error
    When I retrieve field values "state" for the last case
    Then the response should not be an error

  @happy
  Scenario: Case lifecycle create retrieve update
    When I create a case with short description "Lifecycle Case" and description "Lifecycle desc"
    Then the response should not be an error
    When I retrieve the case by its sys_id
    Then the response should not be an error
    When I update the last case with PUT short description "Updated Lifecycle Case"
    Then the response should not be an error

  @error
  Scenario: Retrieve non-existent case returns error
    When I request the case with sys_id "00000000000000000000000000000000"
    Then the response should be an API error
