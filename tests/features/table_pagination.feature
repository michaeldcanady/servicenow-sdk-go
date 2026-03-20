@integration @mock @table @pagination
Feature: ServiceNow Table API Pagination
  As a developer using the ServiceNow SDK
  I want to be able to paginate through records
  So that I can handle large datasets efficiently

  Background:
    And I have a valid ServiceNow instance and credentials
    And I have initialized the ServiceNow client


  @integration @table @pagination
  Scenario: Successfully iterate through multiple pages of records
    Given I set the page size to 2
    When I use the Table PageIterator to fetch records
    Then I should be able to reach the second page
    And the total count of records retrieved should be greater than 2
