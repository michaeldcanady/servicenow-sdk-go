// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package mockdata

var CaseList = `{
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

var CaseItem = `{
  "result": {
    "sys_id": "mock_case_sys_id_1",
    "number": "CS0001001",
    "short_description": "BDD Test Case",
    "description": "Created by integration test",
    "state": "1"
  }
}`

var CaseCreated = `{
  "result": {
    "sys_id": "new_mock_case_sys_id",
    "number": "CS0001003",
    "short_description": "BDD Test Case",
    "description": "Created by integration test",
    "state": "1"
  }
}`

var CaseUpdated = `{
  "result": {
    "sys_id": "new_mock_case_sys_id",
    "number": "CS0001003",
    "short_description": "Updated BDD Test Case",
    "description": "Created by integration test",
    "state": "1"
  }
}`

var CaseActivities = `{
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

var CaseFieldValues = `{
  "result": {
    "label": "State",
    "sequence": "1",
    "dependent_value": ""
  }
}`
