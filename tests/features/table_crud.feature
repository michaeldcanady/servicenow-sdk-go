Feature: ServiceNow Table API CRUD Operations
  As a developer using the ServiceNow SDK
  I want to be able to create, read, update, and delete records
  So that I can manage data in my instance

  Background:
    Given I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client

  @integration @table @crud
  Scenario: Successfully perform CRUD operations on incident table
    When I create a new incident with description "Created by Godog"
    Then the response should not be an error
    And the created record should have a valid "sys_id"
    
    When I update the incident description to "Updated by Godog"
    Then the response should not be an error
    And the record should have description "Updated by Godog"

    When I delete the created incident
    Then the response should not be an error

    When I request the deleted incident by its "sys_id"
    Then the response should be a 404 error
