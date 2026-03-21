@integration @mock @attachment @content
Feature: ServiceNow Attachment Content Retrieval
  As a developer using the ServiceNow SDK
  I want to be able to retrieve the content of an attachment
  So that I can download and process files

  Background:
    Given I have a valid ServiceNow instance
    And I authenticate with Basic Auth

  @integration @attachment @content
  Scenario: Successfully retrieve attachment content
    Given I have at least 1 record in the "incident" table
    And I upload the file "test.txt" from the resources directory to the record
    When I request the content of the created attachment
    Then the response should not be an error
    And the retrieved content should match the original file "test.txt"
    And I delete the created attachment
