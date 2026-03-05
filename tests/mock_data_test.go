//go:build integration

package tests

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
