@integration @mock @attachment @api
Feature: ServiceNow Attachment API
  As a developer using the ServiceNow SDK
  I want to be able to manage attachments
  So that I can handle files in my instance

  Background:
    Given I have a valid ServiceNow instance
    And I authenticate with Basic Auth

  @integration @attachment @collection
  Scenario: Successfully fetch attachments collection
    When I request all attachments
    Then the response should not be an error
    And the attachment results should contain at least 0 records

  @integration @attachment @item
  Scenario: Successfully fetch a single attachment
    Given I have at least 1 attachment in the instance
    When I request the attachment by its "sys_id"
    Then the response should not be an error
    And the attachment result should have the correct "sys_id"
