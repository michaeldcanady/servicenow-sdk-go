// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package mockdata

var StatsCount = `{
  "result": {
    "stats": {
      "count": "42"
    }
  }
}`

var StatsCountZero = `{
  "result": {
    "stats": {
      "count": "0"
    }
  }
}`

var StatsSum = `{
  "result": {
    "stats": {
      "count": "42",
      "sum": {
        "reassignment_count": "17"
      }
    }
  }
}`

var StatsEmptyResult = `{
  "result": {
    "stats": {}
  }
}`

var StatsNoAggregateParams = `{
  "error": {
    "message": "No request parameter for aggregate operation specified",
    "detail": ""
  },
  "status": "failure"
}`

var StatsInvalidTable = `{
  "error": {
    "message": "Invalid table",
    "detail": "Table 'this_table_does_not_exist' does not exist"
  },
  "status": "failure"
}`

var StatsGroupBy = `{
  "result": {
    "stats": {
      "count": "42",
      "group_by": {
        "priority": {
          "1": "5",
          "2": "10",
          "3": "27"
        }
      }
    }
  }
}`
