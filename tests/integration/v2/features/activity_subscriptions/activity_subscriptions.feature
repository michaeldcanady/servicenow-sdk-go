@integration @mock @activity_subscriptions @offline
Feature: ServiceNow Activity Subscriptions API
  As a developer using the ServiceNow SDK
  I want to retrieve activities for a stream
  So that I can monitor activity subscriptions

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: Retrieve activities for stream
    When I retrieve activities for stream "mock_stream_id"
    Then the response should not be an error

  @happy
  Scenario: Retrieve facets for context
    When I retrieve facets for context "mock_context_id" and instance "mock_instance_id"
    Then the response should not be an error

  @error
  Scenario: Invalid stream returns error
    When I retrieve activities for stream ""
    Then the response should be an API error
