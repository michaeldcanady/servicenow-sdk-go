@integration @mock @batch @delete
Feature: ServiceNow Batch API DELETE Operations
  As a developer using the ServiceNow SDK
  I want to be able to send batch requests with DELETE operations
  So that I can delete multiple records in a single call

  Background:
    Given I have a valid ServiceNow instance
    And I authenticate with Basic Auth

  @integration @batch @delete
  Scenario: Successfully send a batch request with a DELETE operation
    Given I create a new record in "incident" with description "Created for Batch DELETE"
    When I send a batch request with a DELETE operation for the created record
    Then the response should not be an error
    And the batch response should contain a successful result for the operation
    When I request the record by its "sys_id"
    Then the response should be a 404 error
