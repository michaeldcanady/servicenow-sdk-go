@integration @mock @batch @delete
Feature: ServiceNow Batch API DELETE Operations
  As a developer using the ServiceNow SDK
  I want to be able to send batch requests with DELETE operations
  So that I can delete multiple records in a single call

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @batch @delete
  Scenario: Successfully send a batch request with a DELETE operation
    Given I create a new incident with description "Created for Batch DELETE"
    When I send a batch request with a DELETE operation for the created incident
    Then the response should not be an error
    And the batch response should contain a successful result for the operation
    When I request the deleted incident by its "sys_id"
    Then the response should be a 404 error
