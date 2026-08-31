// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package mockdata

var CreateServiceResponse = `{
  "result": {
    "sys_id": "mock_app_service_sys_id_1",
    "name": "My App Service",
    "comments": "Created via BDD integration test"
  }
}`

var GetContentResponse = `{
  "result": {
    "sys_id": "mock_app_service_sys_id_1"
  }
}`

var FindServiceResponse = `{
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
