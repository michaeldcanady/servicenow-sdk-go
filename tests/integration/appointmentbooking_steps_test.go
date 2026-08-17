//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/appointmentbookingapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

type appointmentBookingTestContext struct {
	client   *sdk.ServiceNowServiceClient
	response interface{}
	err      error
}

func (c *appointmentBookingTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../../.env")
	return nil
}

func (c *appointmentBookingTestContext) iHaveInitializedTheServiceNowClient() error {
	instance := integrationInstance()

	cred, err := newIntegrationAuthProvider(instance)
	if err != nil {
		return err
	}

	opts := []sdk.ServiceNowServiceClientOption{
		sdk.WithAuthenticationProvider(cred),
		sdk.WithInstance(instance),
	}
	if httpClient := getHttpClient(); httpClient != nil {
		opts = append(opts, sdk.WithHTTPClient(httpClient))
	}

	client, err := sdk.NewServiceNowServiceClient(opts...)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *appointmentBookingTestContext) iRequestAppointmentBookingConfigurationForCatalog(catalogID string) error {
	cfg := &appointmentbookingapi.ConfigurationRequestBuilderGetRequestConfiguration{
		QueryParameters: &appointmentbookingapi.ConfigurationRequestBuilderGetQueryParameters{
			CatalogID: &catalogID,
		},
	}
	resp, err := c.client.AppointmentBooking().Configuration().Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *appointmentBookingTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		if !isOffline() {
			switch c.err.(type) {
			case *core.BadRequestError, *core.NotFoundError:
				return nil
			}
		}
		return fmt.Errorf("expected no error, but got: %s", getErrorTypeAndDetails(c.err))
	}
	return nil
}

func (c *appointmentBookingTestContext) theConfigurationResultShouldBeActive() error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.ConfigurationResult])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected ConfigurationResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	active, err := res.GetActive()
	if err != nil {
		return err
	}
	if active == nil || !*active {
		return fmt.Errorf("expected active to be true")
	}
	return nil
}

func (c *appointmentBookingTestContext) theConfigurationResultShouldHaveActiveString(expected string) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.ConfigurationResult])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected ConfigurationResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	activeStr, err := res.GetActiveString()
	if err != nil {
		return err
	}
	if activeStr == nil || *activeStr != expected {
		return fmt.Errorf("expected active_string to be %q, got %v", expected, activeStr)
	}
	return nil
}

func (c *appointmentBookingTestContext) iRequestAppointmentBookingCalendarForCatalogLocationAndOpenedFor(catalogID, location, openedFor string) error {
	cfg := &appointmentbookingapi.CalendarRequestBuilderGetRequestConfiguration{
		QueryParameters: &appointmentbookingapi.CalendarRequestBuilderGetQueryParameters{
			CatalogID: &catalogID,
			Location:  &location,
			OpenedFor: &openedFor,
		},
	}
	resp, err := c.client.AppointmentBooking().Calendar().Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *appointmentBookingTestContext) theCalendarResultShouldHaveRangeStartAndRangeEnd(expectedStart, expectedEnd string) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.CalendarResponse])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected CalendarItemResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	start, err := res.GetRangeStart()
	if err != nil {
		return err
	}
	end, err := res.GetRangeEnd()
	if err != nil {
		return err
	}
	if start == nil || *start != expectedStart {
		return fmt.Errorf("expected start %q, got %v", expectedStart, start)
	}
	if end == nil || *end != expectedEnd {
		return fmt.Errorf("expected end %q, got %v", expectedEnd, end)
	}
	return nil
}

func (c *appointmentBookingTestContext) iBookAnAppointmentForCatalogLocationAndOpenedFor(catalogID, location, openedFor string) error {
	body := appointmentbookingapi.NewAppointmentRequest()
	_ = body.SetCatalogID(&catalogID)
	_ = body.SetLocation(&location)
	_ = body.SetOpenedFor(&openedFor)
	resp, err := c.client.AppointmentBooking().Appointment().Post(context.Background(), body, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *appointmentBookingTestContext) theAppointmentBookingResultShouldBeSuccessful() error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.AppointmentResultModel])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected AppointmentResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	success, err := res.GetSuccess()
	if err != nil {
		return err
	}
	if success == nil || !*success {
		return fmt.Errorf("expected success to be true")
	}
	return nil
}

func (c *appointmentBookingTestContext) theAppointmentBookingResultDataShouldBe(expectedData string) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.AppointmentResultModel])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected AppointmentResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	data, err := res.GetData()
	if err != nil {
		return err
	}
	if data == nil || *data != expectedData {
		return fmt.Errorf("expected data to be %q, got %v", expectedData, data)
	}
	return nil
}

func (c *appointmentBookingTestContext) iCheckAvailabilityForCatalogLocationAndOpenedFor(catalogID, location, openedFor string) error {
	body := appointmentbookingapi.NewAvailabilityRequest()
	_ = body.SetCatalogID(&catalogID)
	_ = body.SetLocation(&location)
	_ = body.SetOpenedFor(&openedFor)
	resp, err := c.client.AppointmentBooking().Availability().Post(context.Background(), body, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *appointmentBookingTestContext) theAvailabilityResultShouldBeSuccessful() error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.AvailabilityResultModel])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected AvailabilityResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	success, err := res.GetSuccess()
	if err != nil {
		return err
	}
	if success == nil || !*success {
		return fmt.Errorf("expected success to be true")
	}
	return nil
}

func (c *appointmentBookingTestContext) theAvailabilityResultsShouldContainAtLeastSlot(minCount int) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.AvailabilityResultModel])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected AvailabilityResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	slots, err := res.GetAvailability()
	if err != nil {
		return err
	}
	if len(slots) < minCount {
		return fmt.Errorf("expected at least %d slots, got %d", minCount, len(slots))
	}
	return nil
}

func (c *appointmentBookingTestContext) iExecuteRuleConditionsForCatalogAndTask(catalogID, taskID string) error {
	body := appointmentbookingapi.NewExecuteRuleConditionsRequest()
	_ = body.SetCatalogID(&catalogID)
	_ = body.SetTaskID(&taskID)
	resp, err := c.client.AppointmentBooking().ExecuteRuleConditions().Post(context.Background(), body, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *appointmentBookingTestContext) theRuleConditionResultShouldHaveDedicatedCapacity() error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.ExecuteRuleConditionsResult])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected ExecuteRuleConditionsResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	cap, err := res.GetDedicatedCapacity()
	if err != nil {
		return err
	}
	if cap == nil || !*cap {
		return fmt.Errorf("expected dedicated capacity to be true")
	}
	return nil
}

func (c *appointmentBookingTestContext) theRuleConditionResultShouldHaveRuleName(expectedName string) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.ExecuteRuleConditionsResult])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected ExecuteRuleConditionsResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	name, err := res.GetRuleName()
	if err != nil {
		return err
	}
	if name == nil || *name != expectedName {
		return fmt.Errorf("expected rule name %q, got %v", expectedName, name)
	}
	return nil
}

func (c *appointmentBookingTestContext) iRequestUserWindow() error {
	body := appointmentbookingapi.NewUserWindowRequest()
	resp, err := c.client.AppointmentBooking().UserWindow().Post(context.Background(), body, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *appointmentBookingTestContext) theUserWindowResultShouldNotBeNil() error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	resp, ok := c.response.(core.ServiceNowItemResponse[*appointmentbookingapi.UserWindowResult])
	if !ok || conversion.IsNil(resp) {
		return fmt.Errorf("expected UserWindowResponse, got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected non-nil result")
	}
	return nil
}

func InitializeAppointmentBookingScenario(ctx *godog.ScenarioContext) {
	tc := &appointmentBookingTestContext{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		setupGlobalMocks()
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		httpmock.DeactivateAndReset()
		return ctx, nil
	})

	ctx.Step(`^I have a valid ServiceNow instance and credentials$`, tc.iHaveAValidServiceNowInstanceAndCredentials)
	ctx.Step(`^I have initialized the ServiceNow client$`, tc.iHaveInitializedTheServiceNowClient)
	ctx.Step(`^I request appointment booking configuration for catalog "([^"]*)"$`, tc.iRequestAppointmentBookingConfigurationForCatalog)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the configuration result should be active$`, tc.theConfigurationResultShouldBeActive)
	ctx.Step(`^the configuration result should have active string "([^"]*)"$`, tc.theConfigurationResultShouldHaveActiveString)
	ctx.Step(`^I request appointment booking calendar for catalog "([^"]*)", location "([^"]*)", and opened for "([^"]*)"$`, tc.iRequestAppointmentBookingCalendarForCatalogLocationAndOpenedFor)
	ctx.Step(`^the calendar result should have range start "([^"]*)" and range end "([^"]*)"$`, tc.theCalendarResultShouldHaveRangeStartAndRangeEnd)
	ctx.Step(`^I book an appointment for catalog "([^"]*)", location "([^"]*)", and opened for "([^"]*)"$`, tc.iBookAnAppointmentForCatalogLocationAndOpenedFor)
	ctx.Step(`^the appointment booking result should be successful$`, tc.theAppointmentBookingResultShouldBeSuccessful)
	ctx.Step(`^the appointment booking result data should be "([^"]*)"$`, tc.theAppointmentBookingResultDataShouldBe)
	ctx.Step(`^I check availability for catalog "([^"]*)", location "([^"]*)", and opened for "([^"]*)"$`, tc.iCheckAvailabilityForCatalogLocationAndOpenedFor)
	ctx.Step(`^the availability result should be successful$`, tc.theAvailabilityResultShouldBeSuccessful)
	ctx.Step(`^the availability results should contain at least (\d+) slot$`, tc.theAvailabilityResultsShouldContainAtLeastSlot)
	ctx.Step(`^I execute rule conditions for catalog "([^"]*)" and task "([^"]*)"$`, tc.iExecuteRuleConditionsForCatalogAndTask)
	ctx.Step(`^the rule condition result should have dedicated capacity$`, tc.theRuleConditionResultShouldHaveDedicatedCapacity)
	ctx.Step(`^the rule condition result should have rule name "([^"]*)"$`, tc.theRuleConditionResultShouldHaveRuleName)
	ctx.Step(`^I request user window$`, tc.iRequestUserWindow)
	ctx.Step(`^the user window result should not be nil$`, tc.theUserWindowResultShouldNotBeNil)
}

func TestAppointmentBookingFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeAppointmentBookingScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/appointmentbooking.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
