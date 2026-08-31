// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/appointmentbookingapi"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/support"
)

type appointmentBookingSteps struct{}

func (s *appointmentBookingSteps) iRetrieveTheAppointmentBookingConfiguration(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.AppointmentBooking().Configuration().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *appointmentBookingSteps) iRetrieveTheAppointmentBookingCalendar(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.AppointmentBooking().Calendar().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *appointmentBookingSteps) iCreateAnAppointment(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	body := appointmentbookingapi.NewAppointmentRequest()
	_ = body.SetLocation(ptrStr("Test Location"))
	_ = body.SetEndDateUTC(ptrStr("2023-10-01T10:00:00Z"))
	_ = body.SetStartDateUTC(ptrStr("2023-10-01T09:00:00Z"))

	resp, err := w.Client.AppointmentBooking().Appointment().Post(ctx, body, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *appointmentBookingSteps) iCheckAppointmentAvailability(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	body := appointmentbookingapi.NewAvailabilityRequest()
	_ = body.SetLocation(ptrStr("Test Location"))
	_ = body.SetStartDate(ptrStr("2023-10-01T00:00:00Z"))
	_ = body.SetEndDate(ptrStr("2023-10-01T23:59:59Z"))

	resp, err := w.Client.AppointmentBooking().Availability().Post(ctx, body, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *appointmentBookingSteps) iRetrieveTheUserWindow(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	body := appointmentbookingapi.NewUserWindowRequest()

	resp, err := w.Client.AppointmentBooking().UserWindow().Post(ctx, body, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

// InitializeAppointmentBookingScenario registers all appointment booking step definitions.
func InitializeAppointmentBookingScenario(sc *godog.ScenarioContext) {
	s := &appointmentBookingSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I retrieve the appointment booking configuration$`, s.iRetrieveTheAppointmentBookingConfiguration)
	sc.Step(`^I retrieve the appointment booking calendar$`, s.iRetrieveTheAppointmentBookingCalendar)
	sc.Step(`^I create an appointment$`, s.iCreateAnAppointment)
	sc.Step(`^I check appointment availability$`, s.iCheckAppointmentAvailability)
	sc.Step(`^I retrieve the user window$`, s.iRetrieveTheUserWindow)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
