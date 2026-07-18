# Appointment Booking API overview

The Appointment Booking API books and manages service appointments: check
availability windows, create appointments, and read booking configuration.

## Basic usage

This module hangs directly off the client root:

```go
booking := client.AppointmentBooking()

// Check availability
availability, err := booking.Availability().Post(context.Background(), availabilityRequest, nil)

// Book an appointment
appointment, err := booking.Appointment().Post(context.Background(), appointmentRequest, nil)
```

## Available operations

- **Book appointment** — `Appointment().Post(ctx, body, config)`.
- **Check availability** — `Availability().Post(ctx, body, config)`.
- **Calendar** — `Calendar().Get(ctx, config)` reads calendar data.
- **Configuration** — `Configuration().Get(ctx, config)` reads booking configuration.
- **Rule conditions** — `ExecuteRuleConditions().Post(ctx, body, config)` evaluates booking rules.
- **User window** — `UserWindow()` reads the user's booking window.
