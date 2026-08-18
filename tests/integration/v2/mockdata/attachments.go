package mockdata

var AttachmentList = `{
  "result": [
    {
      "sys_id": "mock_attach_id_1",
      "file_name": "test.txt",
      "table_name": "incident",
      "table_sys_id": "mock_sys_id_1",
      "content_type": "text/plain",
      "size": "1024"
    }
  ]
}`

var AttachmentItem = `{
  "result": {
    "sys_id": "mock_attach_id_1",
    "file_name": "test.txt",
    "table_name": "incident",
    "table_sys_id": "mock_sys_id_1",
    "content_type": "text/plain",
    "size": "1024"
  }
}`

var AttachmentCreated = `{
  "result": {
    "sys_id": "new_mock_attach_id",
    "file_name": "test.txt",
    "table_name": "incident",
    "table_sys_id": "mock_sys_id_1",
    "content_type": "text/plain",
    "size": "1024"
  }
}`
