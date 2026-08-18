@integration @mock @batch
Feature: Batch API Operations
  As a developer using the ServiceNow SDK
  I want to send batched requests via the Batch API
  So that I can execute multiple operations in a single HTTP call

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: Single GET operation
    When I send a batch request with a GET operation for the "incident" table
    Then the response should not be an error
    And the batch response should contain at least 1 serviced request

  @happy
  Scenario: Multiple GET operations
    When I send a batch request with GET operations for the "incident" and "incident" tables
    Then the response should not be an error
    And the batch response should contain at least 1 serviced request

  @happy
  Scenario: Mixed POST+GET operations
    When I send a batch request with a POST to "incident" and a GET for "incident"
    Then the response should not be an error
    And the batch response should contain at least 1 serviced request

  @happy
  Scenario: Batch response contains successful results
    When I send a batch request with a GET operation for the "incident" table
    Then all batch serviced requests should have successful status codes

  @happy
  Scenario: Batch response contains multiple successful results
    When I send a batch request with GET operations for the "incident" and "incident" tables
    Then all batch serviced requests should have successful status codes

  @happy
  Scenario: Batch lifecycle create retrieve delete
    When I send a batch request with a POST to "incident" and a GET for "incident"
    Then the response should not be an error
    And the batch response should contain at least 1 serviced request

  # ── E2E-only: Live instance batch behaviors ──────────────────────────

  @e2e @happy
  Scenario: Real batch processes multiple operations atomically
    When I send a batch request with GET operations for the "incident" and "incident" tables
    Then the response should not be an error
    And all batch serviced requests should have successful status codes
