package mockdata

var ActivitySubscriptionItemResponse = `{
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

var FacetsInstanceResponse = `{
  "result": [
    {
      "status": 200,
      "message": "Facet 1",
      "stream": "facet_stream_1",
      "user": "mock_user"
    }
  ]
}`
