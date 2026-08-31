// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package mockdata

var AccountList = `{
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

var AccountListSingle = `{
  "result": [
    {
      "sys_id": "mock_account_sys_id_1",
      "name": "Mock Account 1",
      "number": "ACCT0001"
    }
  ]
}`

var AccountListEmpty = `{
  "result": []
}`

var AccountItem = `{
  "result": {
    "sys_id": "mock_account_sys_id_1",
    "name": "Mock Account 1",
    "number": "ACCT0001"
  }
}`

var AccountCreated = `{
  "result": {
    "sys_id": "new_mock_account_sys_id",
    "name": "BDD Test Account",
    "number": "ACCT0003"
  }
}`

var AccountUpdated = `{
  "result": {
    "sys_id": "new_mock_account_sys_id",
    "name": "Updated Account",
    "number": "ACCT0003"
  }
}`
