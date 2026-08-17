@integration @mock @actsub @api
Feature: ServiceNow Activity Subscriptions API
  As a developer using the ServiceNow SDK
  I want to query activity subscriptions and their facets
  So that I can retrieve activity stream information

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @actsub @activities
  Scenario: Successfully fetch activities for a context and instance
    When I request activities for context "connect" and instance "mock_connect_instance"
    Then the response should not be an error
    And the activity subscription result should have status 200
    And the activity subscription result should have message "Success"
    And the activity subscription result should contain at least 1 activity

  @integration @actsub @activities @error
  Scenario: Fail to fetch activities without required query parameters
    When I request activities without query parameters
    Then the response should be an API error

  @integration @actsub @facets
  Scenario: Successfully fetch facets for context and instance
    When I request facets for context "connect" and instance "mock_connect_instance"
    Then the response should not be an error
    And the facets results should contain at least 1 record
