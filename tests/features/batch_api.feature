@integration @mock @batch @api
Feature: ServiceNow Batch API
  As a developer using the ServiceNow SDK
  I want to be able to send batch requests
  So that I can perform multiple operations in a single call

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @batch
  Scenario: Successfully send a batch request
    When I send a batch request with a GET operation for "incident" table
    Then the response should not be an error
    And the batch response should contain a successful result for the operation
