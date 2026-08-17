package mockdata

var IncidentList = `{
  "result": [
    {
      "sys_id": "mock_sys_id_1",
      "short_description": "Mock Incident 1",
      "active": "true",
      "priority": "3",
      "sys_created_on": "2023-01-01 00:00:00"
    },
    {
      "sys_id": "mock_sys_id_2",
      "short_description": "Mock Incident 2",
      "active": "true",
      "priority": "1",
      "sys_created_on": "2023-01-02 00:00:00"
    }
  ]
}`

var IncidentListSortedDesc = `{
  "result": [
    {
      "sys_id": "mock_sys_id_2",
      "short_description": "Mock Incident 2",
      "active": "true",
      "priority": "1",
      "sys_created_on": "2023-01-02 00:00:00"
    },
    {
      "sys_id": "mock_sys_id_1",
      "short_description": "Mock Incident 1",
      "active": "true",
      "priority": "3",
      "sys_created_on": "2023-01-01 00:00:00"
    }
  ]
}`

var IncidentItem = `{
  "result": {
    "sys_id": "mock_sys_id_1",
    "short_description": "Mock Incident 1",
    "active": "true",
    "priority": "3",
    "sys_created_on": "2023-01-01 00:00:00"
  }
}`

var IncidentCreated = `{
  "result": {
    "sys_id": "new_mock_sys_id",
    "short_description": "Created by integration test",
    "active": "true"
  }
}`

var IncidentUpdated = `{
  "result": {
    "sys_id": "new_mock_sys_id",
    "short_description": "Updated by integration test",
    "active": "true"
  }
}`

var IncidentPatched = `{
  "result": {
    "sys_id": "new_mock_sys_id",
    "short_description": "Patched by integration test",
    "active": "true"
  }
}`
