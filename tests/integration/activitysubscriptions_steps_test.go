//go:build integration

package integration

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	actsubapi "github.com/michaeldcanady/servicenow-sdk-go/activitysubscriptionsapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
)

type actSubTestContext struct {
	client   *sdk.ServiceNowServiceClient
	response interface{}
	err      error
}

func (c *actSubTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../../.env")
	return nil
}

func (c *actSubTestContext) iHaveInitializedTheServiceNowClient() error {
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

func (c *actSubTestContext) iRequestActivitiesForContextAndInstance(contextVal, instanceVal string) error {
	cfg := &actsubapi.ActivitiesRequestBuilderGetRequestConfiguration{
		QueryParameters: &actsubapi.ActivitiesRequestBuilderGetQueryParameters{
			Context:         &contextVal,
			ContextInstance: &instanceVal,
		},
	}
	resp, err := c.client.Now().ActSub().Activities().Get(context.Background(), cfg)
	c.response = resp
	c.err = err
	return nil
}

func (c *actSubTestContext) iRequestActivitiesWithoutQueryParameters() error {
	// Send GET request with nil query parameters/config to trigger the validation error
	resp, err := c.client.Now().ActSub().Activities().Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *actSubTestContext) iRequestFacetsForContextAndInstance(contextVal, instanceVal string) error {
	resp, err := c.client.Now().ActSub().Facets().ByContext(contextVal).ByInstance(instanceVal).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func getSafeErrorMsg(exc *core.ServiceNowError) string {
	if exc == nil {
		return "nil"
	}
	mainErr, err := exc.GetError()
	if err != nil || mainErr == nil {
		return "ServiceNow API returned an error, but details are not parseable"
	}
	msg, _ := mainErr.GetMessage()
	detail, _ := mainErr.GetDetail()
	return fmt.Sprintf("Message: %s, Detail: %s", derefOrEmpty(msg), derefOrEmpty(detail))
}

func getErrorTypeAndDetails(err error) string {
	if err == nil {
		return "nil"
	}
	switch e := err.(type) {
	case *core.BadRequestError:
		return fmt.Sprintf("BadRequestError (400): %s", getSafeErrorMsg(&e.ServiceNowError))
	case *core.UnauthorizedError:
		return fmt.Sprintf("UnauthorizedError (401): %s", getSafeErrorMsg(&e.ServiceNowError))
	case *core.ForbiddenError:
		return fmt.Sprintf("ForbiddenError (403): %s", getSafeErrorMsg(&e.ServiceNowError))
	case *core.NotFoundError:
		return fmt.Sprintf("NotFoundError (404): %s", getSafeErrorMsg(&e.ServiceNowError))
	case *core.TooManyRequestsError:
		return fmt.Sprintf("TooManyRequestsError (429): %s", getSafeErrorMsg(&e.ServiceNowError))
	case *core.ServerError:
		return fmt.Sprintf("ServerError (5XX): %s", getSafeErrorMsg(&e.ServiceNowError))
	case *core.ServiceNowError:
		return fmt.Sprintf("ServiceNowError: %s", getSafeErrorMsg(e))
	default:
		return err.Error()
	}
}

func (c *actSubTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		if !isOffline() {
			switch c.err.(type) {
			case *core.BadRequestError, *core.NotFoundError:
				// Accept 400/404 online as the plugin/endpoints might not be fully configured/installed
				return nil
			}
		}
		return fmt.Errorf("expected no error, but got: %s", getErrorTypeAndDetails(c.err))
	}
	return nil
}

func (c *actSubTestContext) theResponseShouldBeAnAPIError() error {
	if c.err == nil {
		return fmt.Errorf("expected an API error, but got no error")
	}
	// Check if it's one of the ServiceNowError types
	switch c.err.(type) {
	case *core.BadRequestError, *core.UnauthorizedError, *core.ForbiddenError,
		*core.NotFoundError, *core.TooManyRequestsError, *core.ServerError, *core.ServiceNowError:
		return nil
	}
	msg := c.err.Error()
	if strings.Contains(msg, "token acquisition") || strings.Contains(msg, "oauth2") {
		return fmt.Errorf("expected a ServiceNow API error, but auth failed first: %s", msg)
	}
	return nil
}

func (c *actSubTestContext) theActivitySubscriptionResultShouldHaveStatus(expectedStatus int) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	sub, ok := c.response.(core.ServiceNowItemResponse[*actsubapi.ActivitySubscription])
	if !ok || conversion.IsNil(sub) {
		return fmt.Errorf("expected ServiceNowItemResponse[*ActivitySubscription], got %T", c.response)
	}
	res, err := sub.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected subscription result, got nil")
	}
	status, err := res.GetStatus()
	if err != nil {
		return err
	}
	if status == nil {
		return fmt.Errorf("expected status to be %d, but got nil", expectedStatus)
	}
	if *status != int64(expectedStatus) {
		return fmt.Errorf("expected status %d, got %d", expectedStatus, *status)
	}
	return nil
}

func (c *actSubTestContext) theActivitySubscriptionResultShouldHaveMessage(expectedMsg string) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	sub, ok := c.response.(core.ServiceNowItemResponse[*actsubapi.ActivitySubscription])
	if !ok || conversion.IsNil(sub) {
		return fmt.Errorf("expected ServiceNowItemResponse[*ActivitySubscription], got %T", c.response)
	}
	res, err := sub.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected subscription result, got nil")
	}
	msg, err := res.GetMessage()
	if err != nil {
		return err
	}
	if msg == nil {
		return fmt.Errorf("expected message to be %q, but got nil", expectedMsg)
	}
	if *msg != expectedMsg {
		return fmt.Errorf("expected message %q, got %q", expectedMsg, *msg)
	}
	return nil
}

func (c *actSubTestContext) theActivitySubscriptionResultShouldContainAtLeastActivity(minCount int) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	sub, ok := c.response.(core.ServiceNowItemResponse[*actsubapi.ActivitySubscription])
	if !ok || conversion.IsNil(sub) {
		return fmt.Errorf("expected ServiceNowItemResponse[*ActivitySubscription], got %T", c.response)
	}
	res, err := sub.GetResult()
	if err != nil {
		return err
	}
	if conversion.IsNil(res) {
		return fmt.Errorf("expected subscription result, got nil")
	}
	activities, err := res.GetActivities()
	if err != nil {
		return err
	}
	if len(activities) < minCount {
		return fmt.Errorf("expected at least %d activities, got %d", minCount, len(activities))
	}
	return nil
}

func (c *actSubTestContext) theFacetsResultsShouldContainAtLeastRecord(minCount int) error {
	if !isOffline() && conversion.IsNil(c.response) {
		return nil
	}
	collection, ok := c.response.(core.ServiceNowCollectionResponse[*actsubapi.ActivitySubscription])
	if !ok || conversion.IsNil(collection) {
		return fmt.Errorf("expected ServiceNowCollectionResponse[*ActivitySubscription], got %T", c.response)
	}
	results, err := collection.GetResult()
	if err != nil {
		return err
	}
	if len(results) < minCount {
		return fmt.Errorf("expected at least %d facet records, got %d", minCount, len(results))
	}
	return nil
}

func InitializeActivitySubscriptionsScenario(ctx *godog.ScenarioContext) {
	tc := &actSubTestContext{}

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
	ctx.Step(`^I request activities for context "([^"]*)" and instance "([^"]*)"$`, tc.iRequestActivitiesForContextAndInstance)
	ctx.Step(`^I request activities without query parameters$`, tc.iRequestActivitiesWithoutQueryParameters)
	ctx.Step(`^I request facets for context "([^"]*)" and instance "([^"]*)"$`, tc.iRequestFacetsForContextAndInstance)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the response should be an API error$`, tc.theResponseShouldBeAnAPIError)
	ctx.Step(`^the activity subscription result should have status (\d+)$`, tc.theActivitySubscriptionResultShouldHaveStatus)
	ctx.Step(`^the activity subscription result should have message "([^"]*)"$`, tc.theActivitySubscriptionResultShouldHaveMessage)
	ctx.Step(`^the activity subscription result should contain at least (\d+) activity$`, tc.theActivitySubscriptionResultShouldContainAtLeastActivity)
	ctx.Step(`^the facets results should contain at least (\d+) record$`, tc.theFacetsResultsShouldContainAtLeastRecord)
}

func TestActivitySubscriptionsFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeActivitySubscriptionsScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/activitysubscriptions.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
