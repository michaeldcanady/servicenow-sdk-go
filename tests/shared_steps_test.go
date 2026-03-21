//go:build integration

package tests

import (
	"context"
	"fmt"
	"os"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

type testContext struct {
	client   *sdk.ServiceNowServiceClient
	response interface{}
	err      error

	// Contextual data
	lastSysID     string
	incidentSysID string
	tableName     string

	// Request configuration
	pageSize int
	fields   []string
	query    string
	limit    int
}

func (c *testContext) reset() {
	c.response = nil
	c.err = nil
}

// --- Common Steps ---

func (c *testContext) iHaveAValidServiceNowInstance() error {
	_ = godotenv.Load("../.env")
	instance := os.Getenv("SN_INSTANCE")
	if instance == "" && isOffline() {
		instance = "mock_instance"
		_ = os.Setenv("SN_INSTANCE", instance)
	}
	if instance == "" {
		return fmt.Errorf("SN_INSTANCE environment variable is not set")
	}
	return nil
}

func (c *testContext) iAuthenticateWithBasicAuth() error {
	username := os.Getenv("SN_USERNAME")
	password := os.Getenv("SN_PASSWORD")

	if isOffline() {
		if username == "" {
			username = "admin"
		}
		if password == "" {
			password = "admin"
		}
	}

	cred := credentials.NewBasicProvider(username, password)
	instance := os.Getenv("SN_INSTANCE")

	clientOpts := []sdk.ServiceNowServiceClientOption{
		sdk.WithInstance(instance),
	}

	if isOffline() {
		clientOpts = append(clientOpts, sdk.WithHTTPClient(getHttpClient()))
	}

	client, err := sdk.NewServiceNowServiceClientWithOptions(
		cred,
		clientOpts...,
	)
	if err != nil {
		panic(err)
		return err
	}
	c.client = client
	return nil
}

func (c *testContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		if snErr, ok := c.err.(*model.ServicenowError); ok {
			inner, _ := snErr.GetError()
			msg, _ := inner.GetMessage()
			detail, _ := inner.GetDetail()
			status, _ := inner.GetStatus()

			return fmt.Errorf("expected no error, but got: %v (Message: %s, Detail: %s, Status: %s)",
				c.err, *msg, *detail, *status)
		}
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *testContext) theResponseShouldBeA404Error() error {
	if c.err == nil {
		return fmt.Errorf("expected a 404 error, but got no error")
	}
	return nil
}

// --- Shared Helper for Table Operations ---

func (c *testContext) iRequestTheRecordByItsSysID() error {
	if c.client == nil {
		return fmt.Errorf("client not initialized")
	}
	if c.tableName == "" {
		return fmt.Errorf("table name not set")
	}
	if c.lastSysID == "" {
		return fmt.Errorf("sys_id not set")
	}

	resp, err := c.client.Now().Table(c.tableName).ById(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theResultShouldHaveTheCorrectSysID() error {
	resp, ok := c.response.(model.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowItemResponse[*tableapi.TableRecord], got %T", c.response)
	}
	res, err := resp.GetResult()
	if err != nil {
		return err
	}
	sysID, _ := res.GetSysID()
	if *sysID != c.lastSysID {
		return fmt.Errorf("expected sys_id %s, got %s", c.lastSysID, *sysID)
	}
	return nil
}

func (c *testContext) iDeleteTheCreatedRecord() error {
	if c.client == nil {
		return fmt.Errorf("client not initialized")
	}
	err := c.client.Now().Table(c.tableName).ById(c.lastSysID).Delete(context.Background(), nil)
	c.err = err
	return nil
}

// registerSharedSteps registers steps that are common across all features.
func registerSharedSteps(ctx *godog.ScenarioContext, tc *testContext) {
	ctx.Step(`^I have a valid ServiceNow instance$`, tc.iHaveAValidServiceNowInstance)
	ctx.Step(`^I authenticate with Basic Auth$`, tc.iAuthenticateWithBasicAuth)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the response should be a 404 error$`, tc.theResponseShouldBeA404Error)

	// Reusable table steps
	ctx.Step(`^I request the record by its "sys_id"$`, tc.iRequestTheRecordByItsSysID)
	ctx.Step(`^the result should have the correct "sys_id"$`, tc.theResultShouldHaveTheCorrectSysID)
	ctx.Step(`^I delete the created record$`, tc.iDeleteTheCreatedRecord)
}
