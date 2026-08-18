@integration @mock @appointment_booking @offline
Feature: ServiceNow Appointment Booking API
  As a developer using the ServiceNow SDK
  I want to manage appointment bookings
  So that I can configure, schedule, and check availability for appointments

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  # ── Happy Path ──────────────────────────────────────────────────────

  @happy
  Scenario: Retrieve booking configuration
    When I retrieve the appointment booking configuration
    Then the response should not be an error

  @happy
  Scenario: Retrieve calendar
    When I retrieve the appointment booking calendar
    Then the response should not be an error

  @happy
  Scenario: Create appointment
    When I create an appointment
    Then the response should not be an error

  @happy
  Scenario: Check availability
    When I check appointment availability
    Then the response should not be an error

  @happy
  Scenario: Get user window
    When I retrieve the user window
    Then the response should not be an error
