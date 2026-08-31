// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package mockdata

var CmdbList = `{
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

var CmdbListSingle = `{
  "result": [
    {
      "sys_id": "mock_cmdb_sys_id_1",
      "name": "Mock CI 1",
      "className": "cmdb_ci"
    }
  ]
}`

var CmdbListEmpty = `{
  "result": []
}`

var CmdbItem = `{
  "result": {
    "attributes": {
      "sys_id": "mock_cmdb_sys_id_1",
      "name": "Mock CI 1"
    },
    "outbound_relations": [],
    "inbound_relations": []
  }
}`

var CmdbInvalidClass = `{
  "error": {
    "message": "Invalid class",
    "detail": "Class 'this_cmdb_class_does_not_exist' does not exist"
  },
  "status": "failure"
}`
