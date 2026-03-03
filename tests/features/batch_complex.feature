Feature: ServiceNow Batch API Complex Operations
  As a developer using the ServiceNow SDK
  I want to be able to send complex batch requests with multiple operations
  So that I can perform varied operations in a single call

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @batch @complex
  Scenario: Successfully send a batch request with multiple GET operations
    When I send a batch request with GET operations for "incident" and "change_request" tables
    Then the response should not be an error
    And the batch response should contain 2 successful results
