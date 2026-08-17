@integration @documents @api
Feature: ServiceNow Documents API
  As a developer using the ServiceNow SDK
  I want to validate the behavior of the Documents endpoints
  So that the SDK matches the ServiceNow document API contract

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @documents @explore
  Scenario: Successfully explore documents
    When I explore documents
    Then the response should not be an error
    And the document collection should contain at least 0 documents

  @integration @documents @create
  Scenario: Create document without required payload fails with a service error
    When I create a document with name "BDD Document" and type "pdf"
    Then the response should be an API error

  @integration @documents @create_or_link
  Scenario: Create-or-link document without required payload fails with a service error
    When I create or link a document with name "BDD Document" and type "pdf"
    Then the response should be an API error

  @integration @documents @delete
  Scenario: Delete document without required query parameters fails with a service error
    When I delete a document without required query parameters
    Then the response should be an API error

  @integration @documents @versions
  Scenario: List document versions for a non-existent document fails with a service error
    When I request versions for document "00000000000000000000000000000000"
    Then the response should be an API error

  @integration @documents @version_state
  Scenario: Get version state for a non-existent version returns a response
    When I request the version state for version "00000000000000000000000000000000"
    Then the response should not be an error

  @integration @documents @content
  Scenario: Get document content for a non-existent document fails with a service error
    When I request content for document "00000000000000000000000000000000"
    Then the response should be an API error

  @integration @documents @syncdown
  Scenario: Sync down a non-existent document fails with a service error
    When I sync down document "00000000000000000000000000000000"
    Then the response should be an API error

  @integration @documents @attach
  Scenario: Attach a document using a missing provider fails with a service error
    When I attach a document using provider "missing-provider"
    Then the response should be an API error

  @integration @documents @action
  Scenario: Execute a document action against an invalid document version fails with a service error
    When I execute version action "checkout" for document "00000000000000000000000000000000" and version "00000000000000000000000000000000"
    Then the response should be an API error

  @integration @documents @explore @pagination
  Scenario: Explore documents with pagination parameters
    When I explore documents with limit 5
    Then the response should not be an error

  @integration @documents @explore @filtering
  Scenario: Explore documents with type filter
    When I explore documents with type filter "document"
    Then the response should not be an error

  @integration @documents @versions @missing_doc
  Scenario: List versions for document fails for non-existent document
    When I request versions for document "99999999999999999999999999999999"
    Then the response should be an API error

  @integration @documents @content @binary
  Scenario: Get binary content for non-existent document fails
    When I request content for document "99999999999999999999999999999999"
    Then the response should be an API error

  @integration @documents @delete @required_fields
  Scenario: Delete requires at least doc_sys_id query parameter
    When I delete a document with doc_sys_id "12345678901234567890123456789012"
    Then the response should be an API error

  @integration @documents @create @empty_payload
  Scenario: Create with empty payload object fails with validation error
    When I create a document with empty payload
    Then the response should be an API error

  @integration @documents @syncdown @invalid_state
  Scenario: Sync down with invalid document state fails
    When I sync down document "99999999999999999999999999999999"
    Then the response should be an API error

  @integration @documents @attach @no_auth
  Scenario: Attach with unauthorized provider fails
    When I attach a document using provider "unauthorized-provider"
    Then the response should be an API error

  @integration @documents @action @invalid_action
  Scenario: Execute invalid document action fails
    When I execute version action "invalid-action" for document "00000000000000000000000000000000" and version "00000000000000000000000000000000"
    Then the response should be an API error
