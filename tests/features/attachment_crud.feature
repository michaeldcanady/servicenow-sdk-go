@integration @mock @attachment @crud
Feature: ServiceNow Attachment API CRUD Operations
  As a developer using the ServiceNow SDK
  I want to be able to upload, retrieve, and delete attachments
  So that I can manage files in my instance

  Background:
    Given I have a valid ServiceNow instance
    And I authenticate with Basic Auth

  @integration @attachment @crud
  Scenario: Successfully upload and delete an attachment
    Given I have at least 1 record in the "incident" table
    When I upload the file "test.txt" from the resources directory to the record
    Then the response should not be an error
    And the created attachment should have a valid "sys_id"
    And the attachment filename should be "test.txt"

    When I request the attachment by its "sys_id"
    Then the response should not be an error
    And the attachment result should have the correct "sys_id"

    When I delete the created attachment
    Then the response should not be an error

    When I request the record by its "sys_id"
    And I request the deleted attachment by its "sys_id"
    Then the response should be a 404 error
