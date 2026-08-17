@integration @mock @documents @offline
Feature: ServiceNow Documents API
  As a developer using the ServiceNow SDK
  I want to explore, create, and manage documents
  So that I can work with document metadata and content

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: Explore documents
    When I explore available documents
    Then the response should not be an error

  @happy
  Scenario: Create document
    When I create a document
    Then the response should not be an error

  @happy
  Scenario: Delete document
    When I create a document
    Then the response should not be an error
    When I delete a document
    Then the response should not be an error

  @happy
  Scenario: Get versions
    When I get versions for document "mock_doc_sys_id_1"
    Then the response should not be an error

  @happy
  Scenario: Get content
    When I get content for document "mock_doc_sys_id_1"
    Then the response should not be an error
