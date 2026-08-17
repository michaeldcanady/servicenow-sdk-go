@integration @mock @attachment
Feature: Attachment API Operations
  As a developer using the ServiceNow SDK
  I want to manage file attachments via the Attachment API
  So that I can upload, retrieve, list, and delete attachments programmatically

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: List all attachments
    When I request all attachments
    Then the response should not be an error

  @happy
  Scenario: Results contain at least 1 attachment
    Given I have at least 1 attachment in the instance
    Then the response should not be an error

  @happy
  Scenario: Get attachment by sys_id after upload
    When I upload a file to an incident
    Then the response should not be an error
    When I request the attachment by its sys_id
    Then the response should not be an error

  @happy
  Scenario: Upload file to incident
    When I upload a file to an incident
    Then the response should not be an error

  @happy
  Scenario: Get content of uploaded file
    When I upload a file to an incident
    Then the response should not be an error
    When I request the content of the created attachment
    Then the response should not be an error

  @happy
  Scenario: Delete attachment
    When I upload a file to an incident
    Then the response should not be an error
    When I delete the attachment
    Then the response should not be an error

  @error
  Scenario: Delete non-existent attachment returns error
    When I delete the attachment with sys_id "00000000000000000000000000000000"
    Then the response should be an API error

  @happy
  Scenario: Upload multiple files
    When I upload a file to an incident
    Then the response should not be an error
    When I upload a second file to an incident
    Then the response should not be an error

  @happy @offline
  Scenario: Verify attachment metadata
    Given I have at least 1 attachment in the instance
    When I request the attachment by its sys_id
    Then the response should not be an error
    And the attachment should have a file_name
    And the attachment should have a content_type

  @happy
  Scenario: Attachment lifecycle upload verify get content delete
    When I upload a file to an incident
    Then the response should not be an error
    When I request the attachment by its sys_id
    Then the response should not be an error
    When I request the content of the created attachment
    Then the response should not be an error
    When I delete the attachment
    Then the response should not be an error
