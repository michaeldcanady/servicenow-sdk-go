@integration @mock @caseapi @api
Feature: ServiceNow Customer Service Case API
  As a developer using the ServiceNow SDK
  I want to create, retrieve, update, and inspect cases
  So that I can manage customer service cases programmatically

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @caseapi @list
  Scenario: Successfully list cases
    When I list all cases
    Then the response should not be an error
    And the case collection should contain at least 1 case

  @integration @caseapi @lifecycle
  Scenario: Successfully manage a case lifecycle
    When I create a case with short description "BDD Test Case" and description "Created by integration test"
    Then the response should not be an error
    And the created case should have a valid sys_id
    When I retrieve the case by its sys_id
    Then the response should not be an error
    And the retrieved case should have short description "BDD Test Case"
    When I update the case short description to "Updated BDD Test Case"
    Then the response should not be an error
    And the updated case should have short description "Updated BDD Test Case"

  @integration @caseapi @activities
  Scenario: Successfully retrieve case activities
    When I create a case with short description "Activity Test Case" and description "Case for activity testing"
    Then the response should not be an error
    And the created case should have a valid sys_id
    When I retrieve activities for the case
    Then the response should not be an error
    And the activities response should not be nil

  @integration @caseapi @fieldvalues
  Scenario: Successfully retrieve case field values
    When I retrieve field values for the "state" field
    Then the response should not be an error
    And the field values response should not be nil

  @integration @caseapi @itemfieldvalues
  Scenario: Successfully retrieve field values for a specific case
    When I create a case with short description "Field Values Test Case" and description "Case for field values testing"
    Then the response should not be an error
    And the created case should have a valid sys_id
    When I retrieve field values for the "state" field on the case
    Then the response should not be an error
    And the field values response should not be nil
