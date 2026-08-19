# Appointment Booking API

The Appointment Booking API provides endpoints for managing appointment scheduling, availability, calendar lookups, and configuration.

## \[POST\] /sn_apptmnt_booking/v1/appointment/appointment

Creates a new appointment.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/appointmentbookingapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &appointmentbookingapi.AppointmentRequest{}
	response, err := client.Now().AppointmentBooking().Appointment().Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[POST\] /sn_apptmnt_booking/v1/appointment/availability

Checks appointment availability.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/appointmentbookingapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &appointmentbookingapi.AvailabilityRequest{}
	response, err := client.Now().AppointmentBooking().Availability().Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /sn_apptmnt_booking/v1/appointment/calendar

Retrieves calendar data for appointment scheduling.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().AppointmentBooking().Calendar().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /sn_apptmnt_booking/v1/appointment/configuration

Retrieves appointment booking configuration.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().AppointmentBooking().Configuration().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
