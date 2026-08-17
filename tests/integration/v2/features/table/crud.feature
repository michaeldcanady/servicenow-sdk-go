@integration @mock @table
Feature: Table API CRUD Operations
  As a developer using the ServiceNow SDK
  I want to create, read, update, and delete table records
  So that I can manage data lifecycle programmatically

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: Create record and verify all fields
    When I create a new incident with description "BDD Test Incident"
    Then the response should not be an error
    And the created record should have a valid "sys_id"

  @happy
  Scenario: Create record with minimal fields
    When I create a new incident with description "Minimal"
    Then the response should not be an error
    And the created record should have a valid "sys_id"

  @happy
  Scenario: Retrieve record by sys_id
    When I create a new incident with description "Retrieve Test"
    Then the response should not be an error
    When I request the incident by its "sys_id"
    Then the response should not be an error

  @error
  Scenario: Retrieve non-existent record returns error
    When I request the incident with sys_id "00000000000000000000000000000000"
    Then the response should be an API error

  @happy
  Scenario: Fetch collection of records
    When I request all incidents from the "incident" table
    Then the response should not be an error
    And the results should contain at least 1 record

  @happy @offline
  Scenario: Fetch collection with query filter
    When I request all incidents from the "incident" table with query "ORDERBYDESCsys_created_on"
    Then the response should not be an error
    And the results should contain at least 1 record

  @happy @offline
  Scenario: Fetch collection with display_value parameter
    When I request all incidents from the "incident" table with display_value "all"
    Then the response should not be an error
    And the results should contain at least 1 record

  @happy
  Scenario: Update record using PUT
    When I create a new incident with description "PUT Test"
    Then the response should not be an error
    When I update the last incident with PUT description "Updated via PUT"
    Then the response should not be an error

  @happy
  Scenario: Partial update using PATCH
    When I create a new incident with description "PATCH Test"
    Then the response should not be an error
    When I patch the last incident description to "Patched Description"
    Then the response should not be an error

  @happy
  Scenario: Delete record
    When I create a new incident with description "Delete Test"
    Then the response should not be an error
    When I delete the last incident
    Then the response should not be an error

  @error
  Scenario: Delete non-existent record returns error
    When I delete the incident "00000000000000000000000000000000"
    Then the response should be an API error

  @happy
  Scenario: HEAD request returns headers
    When I create a new incident with description "HEAD Test"
    Then the response should not be an error
    When I send a HEAD request for the last incident
    Then the response should not be an error

  @happy
  Scenario: Create, retrieve, update, delete full lifecycle
    When I create a new incident with description "Lifecycle Test"
    Then the response should not be an error
    When I request the incident by its "sys_id"
    Then the response should not be an error
    When I update the last incident with PUT description "Updated Lifecycle"
    Then the response should not be an error
    When I delete the last incident
    Then the response should not be an error

  @happy
  Scenario: Create record with special characters in description
    When I create a new incident with description "Special chars: @#$%^&*()"
    Then the response should not be an error
    And the created record should have a valid "sys_id"

  @happy @offline
  Scenario: Request with sysparm_fields to limit returned fields
    When I request all incidents from the "incident" table with fields "sys_id,short_description"
    Then the response should not be an error
    And the results should contain at least 1 record

  @happy @offline
  Scenario: Request with sysparm_limit and sysparm_offset
    When I request all incidents from the "incident" table with limit 1 and offset 0
    Then the response should not be an error
    And the results should contain at least 1 record

  # ── E2E-only: Live instance behaviors ────────────────────────────────

  @e2e @happy
  Scenario: Real pagination returns multiple pages
    When I request all incidents from the "incident" table
    Then the response should not be an error
    And the results should contain at least 1 record

  @e2e @error
  Scenario: Real error response has ServiceNow error structure
    When I request the incident with sys_id "00000000000000000000000000000000"
    Then the response should be an API error

  @e2e @happy
  Scenario: Create and retrieve round-trips real fields
    When I create a new incident with description "E2E Round-trip Test"
    Then the response should not be an error
    And the created record should have a valid "sys_id"
    When I request the incident by its "sys_id"
    Then the response should not be an error
    And the retrieved record should have field "short_description" containing "E2E Round-trip Test"
