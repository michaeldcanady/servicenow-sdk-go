@integration @mock @batch @multi-method
Feature: ServiceNow Batch API Multi-Method Operations
  As a developer using the ServiceNow SDK
  I want to be able to send batch requests with different HTTP methods
  So that I can perform mixed operations in a single call

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @batch @multi-method
  Scenario: Successfully send a batch request with POST and GET operations
    When I send a batch request with a POST to "incident" and a GET for "incident"
    Then the response should not be an error
    And the batch response should contain 2 successful results
