// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package mockdata

var BatchResponse = `{
  "serviced_requests": [
    {
      "id": "1",
      "status_code": 200,
      "body": "eyJyZXN1bHQiOlt7InN5c19pZCI6Im1vY2tfc3lzX2lkXzEifV19"
    }
  ],
  "unserviced_requests": []
}`

var BatchMultiResponse = `{
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
