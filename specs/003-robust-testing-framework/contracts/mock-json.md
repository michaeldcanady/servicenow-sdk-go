# Mock JSON Contract

Mock JSON files must represent the exact structure of a ServiceNow REST API response body.

## Example: Incident Record
```json
{
  "result": {
    "sys_id": "8d259e831b201110b4460a661a4bcbe1",
    "short_description": "Network connectivity issue",
    "number": "INC0000001"
  }
}
```

## Example: Error Response
```json
{
  "error": {
    "message": "Record not found",
    "detail": "The sys_id provided does not match any existing record."
  },
  "status": "failure"
}
```
