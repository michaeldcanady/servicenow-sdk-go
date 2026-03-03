Feature: ServiceNow Attachment API CRUD Operations
  As a developer using the ServiceNow SDK
  I want to be able to upload, retrieve, and delete attachments
  So that I can manage files in my instance

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @attachment @crud
  Scenario: Successfully upload and delete an attachment
    Given I have an incident record in the "incident" table
    When I upload the file "test.txt" from the resources directory to the incident
    Then the response should not be an error
    And the created attachment should have a valid "sys_id"
    And the attachment filename should be "test.txt"

    When I request the attachment by its "sys_id"
    Then the response should not be an error
    And the result should have the correct "sys_id"

    When I delete the created attachment
    Then the response should not be an error

    When I request the deleted attachment by its "sys_id"
    Then the response should be a 404 error
