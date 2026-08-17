//go:build integration

package integration

var mockIncidentList = `{
  "result": [
    {
      "sys_id": "mock_sys_id_1",
      "short_description": "Mock Incident 1",
      "active": "true",
      "sys_created_on": "2023-01-01 00:00:00"
    },
    {
      "sys_id": "mock_sys_id_2",
      "short_description": "Mock Incident 2",
      "active": "true",
      "sys_created_on": "2023-01-02 00:00:00"
    }
  ]
}`

var mockIncidentListSortedDesc = `{
  "result": [
    {
      "sys_id": "mock_sys_id_2",
      "short_description": "Mock Incident 2",
      "active": "true",
      "sys_created_on": "2023-01-02 00:00:00"
    },
    {
      "sys_id": "mock_sys_id_1",
      "short_description": "Mock Incident 1",
      "active": "true",
      "sys_created_on": "2023-01-01 00:00:00"
    }
  ]
}`

var mockIncidentItem = `{
  "result": {
    "sys_id": "mock_sys_id_1",
    "short_description": "Mock Incident 1",
    "active": "true",
    "sys_created_on": "2023-01-01 00:00:00"
  }
}`

var mockCreatedIncident = `{
  "result": {
    "sys_id": "new_mock_sys_id",
    "short_description": "Created by Godog",
    "active": "true"
  }
}`

var mockUpdatedIncident = `{
  "result": {
    "sys_id": "new_mock_sys_id",
    "short_description": "Updated by Godog",
    "active": "true"
  }
}`

var mockPatchedIncident = `{
  "result": {
    "sys_id": "new_mock_sys_id",
    "short_description": "Patched by Godog",
    "active": "true"
  }
}`

var mockAttachmentList = `{
  "result": [
    {
      "sys_id": "mock_attach_id_1",
      "file_name": "test.txt",
      "table_name": "incident",
      "table_sys_id": "mock_sys_id_1"
    }
  ]
}`

var mockAttachmentItem = `{
  "result": {
    "sys_id": "mock_attach_id_1",
    "file_name": "test.txt",
    "table_name": "incident",
    "table_sys_id": "mock_sys_id_1"
  }
}`

var mockBatchResponse = `{
  "serviced_requests": [
    {
      "id": "1",
      "status_code": 200,
      "body": "eyJyZXN1bHQiOlt7InN5c19pZCI6Im1vY2tfc3lzX2lkXzEifV19"
    }
  ],
  "unserviced_requests": []
}`

var mockBatchMultiResponse = `{
  "serviced_requests": [
    {
      "id": "1",
      "status_code": 200,
      "body": "eyJyZXN1bHQiOlt7InN5c19pZCI6Im1vY2tfc3lzX2lkXzEifV19"
    },
    {
      "id": "2",
      "status_code": 200,
      "body": "eyJyZXN1bHQiOlt7InN5c19pZCI6Im1vY2tfc3lzX2lkXzIifV19"
    }
  ],
  "unserviced_requests": []
}`

var mockStatsCount = `{
  "result": {
    "stats": {
      "count": "42"
    }
  }
}`

var mockStatsCountZero = `{
  "result": {
    "stats": {
      "count": "0"
    }
  }
}`

var mockStatsSum = `{
  "result": {
    "stats": {
      "count": "42",
      "sum": {
        "reassignment_count": "17"
      }
    }
  }
}`

var mockStatsEmptyResult = `{
  "result": {
    "stats": {}
  }
}`

var mockStatsNoAggregateParams = `{
  "error": {
    "message": "No request parameter for aggregate operation specified",
    "detail": ""
  },
  "status": "failure"
}`

var mockStatsInvalidTable = `{
  "error": {
    "message": "Invalid table",
    "detail": "Table 'this_table_does_not_exist' does not exist"
  },
  "status": "failure"
}`

var mockAccountList = `{
  "result": [
    {
      "sys_id": "mock_account_sys_id_1",
      "name": "Mock Account 1",
      "number": "ACCT0001"
    },
    {
      "sys_id": "mock_account_sys_id_2",
      "name": "Mock Account 2",
      "number": "ACCT0002"
    }
  ]
}`

var mockAccountListSingle = `{
  "result": [
    {
      "sys_id": "mock_account_sys_id_1",
      "name": "Mock Account 1",
      "number": "ACCT0001"
    }
  ]
}`

var mockAccountListEmpty = `{
  "result": []
}`

var mockAccountItem = `{
  "result": {
    "sys_id": "mock_account_sys_id_1",
    "name": "Mock Account 1",
    "number": "ACCT0001"
  }
}`

var mockCmdbList = `{
  "result": [
    {
      "sys_id": "mock_cmdb_sys_id_1",
      "name": "Mock CI 1",
      "className": "cmdb_ci"
    },
    {
      "sys_id": "mock_cmdb_sys_id_2",
      "name": "Mock CI 2",
      "className": "cmdb_ci"
    }
  ]
}`

var mockCmdbListSingle = `{
  "result": [
    {
      "sys_id": "mock_cmdb_sys_id_1",
      "name": "Mock CI 1",
      "className": "cmdb_ci"
    }
  ]
}`

var mockCmdbListEmpty = `{
  "result": []
}`

var mockCmdbItem = `{
  "result": {
    "attributes": {
      "sys_id": "mock_cmdb_sys_id_1",
      "name": "Mock CI 1"
    },
    "outbound_relations": [],
    "inbound_relations": []
  }
}`

var mockCmdbInvalidClass = `{
  "error": {
    "message": "Invalid class",
    "detail": "Class 'this_cmdb_class_does_not_exist' does not exist"
  },
  "status": "failure"
}`

var mockActivitySubscriptionItemResponse = `{
  "result": {
    "status": 200,
    "message": "Success",
    "stream": "mock_stream",
    "user": "mock_user",
    "activities": [
      {
        "sys_id": "activity_1",
        "title": "Mock Activity 1",
        "activity_type_id": "type_1",
        "source_table_name": "incident",
        "subobject_table_name": "incident",
        "subobject_sys_id": "incident_sys_id"
      }
    ]
  }
}`

var mockFacetsInstanceResponse = `{
  "result": [
    {
      "status": 200,
      "message": "Facet 1",
      "stream": "facet_stream_1",
      "user": "mock_user"
    }
  ]
}`

var mockAvailabilityResponse = `{
  "result": {
    "success": true,
    "has_more": false,
    "no_appt_available": false,
    "time_zone": "UTC",
    "availability": [
      {
        "date": "2023-10-01",
        "display_value": "2023-10-01 09:00:00",
        "value": "2023-10-01T09:00:00Z"
      }
    ]
  }
}`

var mockAppointmentResponse = `{
  "result": {
    "success": true,
    "message": "Appointment booked successfully",
    "data": "booking_sys_id_123"
  }
}`

var mockAppointmentBookingConfigurationResponse = `{
  "result": {
    "active": true,
    "active_string": "true",
    "advanced_calendar_view_portal": false,
    "auto_acceptance": true,
    "locale_language": "en"
  }
}`

var mockAppointmentBookingCalendarResponse = `{
  "result": {
    "range_end": "2023-01-31",
    "range_start": "2023-01-01"
  }
}`

var mockExecuteRuleConditionsResponse = `{
  "result": {
    "dedicatedCapacity": true,
    "futureMaxBookableDays": "30",
    "ruleId": "rule_id_123",
    "ruleName": "Rule 1"
  }
}`

var mockUserWindowResponse = `{
  "result": {}
}`

var mockCreateServiceResponse = `{
  "result": {
    "sys_id": "mock_app_service_sys_id_1",
    "name": "My App Service",
    "comments": "Created via BDD integration test"
  }
}`

var mockGetContentResponse = `{
  "result": {
    "sys_id": "mock_app_service_sys_id_1"
  }
}`

var mockFindServiceResponse = `{
  "result": {
    "services": [
      {
        "sys_id": "mock_app_service_sys_id_1",
        "name": "My App Service",
        "number": "APP0001",
        "relationships": [],
        "environment": "production",
        "version": "1.0"
      }
    ]
  }
}`

// Case API mock data

var mockCaseList = `{
  "result": [
    {
      "sys_id": "mock_case_sys_id_1",
      "number": "CS0001001",
      "short_description": "BDD Test Case",
      "description": "Created by integration test",
      "state": "1"
    },
    {
      "sys_id": "mock_case_sys_id_2",
      "number": "CS0001002",
      "short_description": "Another Case",
      "description": "Another integration test case",
      "state": "1"
    }
  ]
}`

var mockCaseItem = `{
  "result": {
    "sys_id": "mock_case_sys_id_1",
    "number": "CS0001001",
    "short_description": "BDD Test Case",
    "description": "Created by integration test",
    "state": "1"
  }
}`

var mockCreatedCase = `{
  "result": {
    "sys_id": "new_mock_case_sys_id",
    "number": "CS0001003",
    "short_description": "BDD Test Case",
    "description": "Created by integration test",
    "state": "1"
  }
}`

var mockUpdatedCase = `{
  "result": {
    "sys_id": "new_mock_case_sys_id",
    "number": "CS0001003",
    "short_description": "Updated BDD Test Case",
    "description": "Created by integration test",
    "state": "1"
  }
}`

var mockCaseActivities = `{
  "result": {
    "entries": [
      {
        "type": "comment",
        "value": "Case created",
        "user": {
          "link": "https://mock_instance.service-now.com/api/now/v1/table/sys_user/mock_user_id",
          "value": "mock_user_id"
        }
      }
    ]
  }
}`

var mockCaseFieldValues = `{
  "result": {
    "label": "State",
    "sequence": "1",
    "dependent_value": ""
  }
}`

// Policy Mappings API mock data

var mockPolicyMappingItem = `{
  "result": {
    "sys_id": "mock_policy_mapping_sys_id_1",
    "number": "CDM0001001",
    "state": "proposed"
  }
}`
