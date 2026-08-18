package mockdata

var AvailabilityResponse = `{
  "result": {
    "success": true,
    "has_more": false,
    "no_appt_available": false,
    "time_zone": "UTC",
    "availability": [
      {
        "date": "2023-10-01",
        "display_value": "2023-10-01 09:00:00",
        "value": "2023-10-01T09:00:00Z"
      }
    ]
  }
}`

var AppointmentResponse = `{
  "result": {
    "success": true,
    "message": "Appointment booked successfully",
    "data": "booking_sys_id_123"
  }
}`

var AppointmentBookingConfigurationResponse = `{
  "result": {
    "active": true,
    "active_string": "true",
    "advanced_calendar_view_portal": false,
    "auto_acceptance": true,
    "locale_language": "en"
  }
}`

var AppointmentBookingCalendarResponse = `{
  "result": {
    "range_end": "2023-01-31",
    "range_start": "2023-01-01"
  }
}`

var ExecuteRuleConditionsResponse = `{
  "result": {
    "dedicatedCapacity": true,
    "futureMaxBookableDays": "30",
    "ruleId": "rule_id_123",
    "ruleName": "Rule 1"
  }
}`

var UserWindowResponse = `{
  "result": {}
}`
