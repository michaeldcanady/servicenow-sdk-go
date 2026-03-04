@integration @mock @table @patch
Feature: ServiceNow Table API PATCH Operation
  As a developer using the ServiceNow SDK
  I want to be able to partially update records using PATCH
  So that I can efficiently update specific fields without sending the entire record

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @table @patch
  Scenario: Successfully perform PATCH operation on incident table
    Given I create a new incident with description "Created for PATCH"
    When I patch the incident description to "Patched by Godog"
    Then the response should not be an error
    And the record should have description "Patched by Godog"
    And I delete the created incident
