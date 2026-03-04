@integration @mock @table @collection
Feature: ServiceNow Table API
  As a developer using the ServiceNow SDK
  I want to be able to interact with tables
  So that I can manage records in my instance

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @table @collection
  Scenario: Successfully fetch incidents collection
    When I request all incidents from the "incident" table
    Then the response should not be an error
    And the results should contain at least 1 record
    And each record should have a valid "sys_id"

  @integration @table @item
  Scenario: Successfully fetch a single incident
    Given I have at least 1 incident in the "incident" table
    When I request the incident by its "sys_id"
    Then the response should not be an error
    And the result should have the correct "sys_id"
