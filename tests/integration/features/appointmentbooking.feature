@integration @mock @appointmentbooking @api
Feature: ServiceNow Appointment Booking API
  As a developer using the ServiceNow SDK
  I want to query appointment booking configurations, calendars, slot availability, rule conditions, and book appointments
  So that I can retrieve scheduling information

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @appointmentbooking @configuration
  Scenario: Successfully fetch appointment booking configuration
    When I request appointment booking configuration for catalog "mock_catalog_id"
    Then the response should not be an error
    And the configuration result should be active
    And the configuration result should have active string "true"

  @integration @appointmentbooking @calendar
  Scenario: Successfully fetch appointment booking calendar
    When I request appointment booking calendar for catalog "mock_catalog_id", location "mock_location", and opened for "mock_user"
    Then the response should not be an error
    And the calendar result should have range start "2023-01-01" and range end "2023-01-31"

  @integration @appointmentbooking @appointment
  Scenario: Successfully book an appointment
    When I book an appointment for catalog "mock_catalog_id", location "mock_location", and opened for "mock_user"
    Then the response should not be an error
    And the appointment booking result should be successful
    And the appointment booking result data should be "booking_sys_id_123"

  @integration @appointmentbooking @availability
  Scenario: Successfully check appointment availability
    When I check availability for catalog "mock_catalog_id", location "mock_location", and opened for "mock_user"
    Then the response should not be an error
    And the availability result should be successful
    And the availability results should contain at least 1 slot

  @integration @appointmentbooking @executeruleconditions
  Scenario: Successfully execute rule conditions
    When I execute rule conditions for catalog "mock_catalog_id" and task "mock_task_id"
    Then the response should not be an error
    And the rule condition result should have dedicated capacity
    And the rule condition result should have rule name "Rule 1"

  @integration @appointmentbooking @userwindow
  Scenario: Successfully retrieve user window
    When I request user window
    Then the response should not be an error
    And the user window result should not be nil
