package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

type tableTestContext struct {
	client    *sdk.ServiceNowClient
	response  interface{} // Generic response to support either collection or item
	err       error
	lastSysID string
}

func (c *tableTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../.env")
	required := []string{"SN_CLIENT_ID", "SN_USERNAME", "SN_PASSWORD", "SN_INSTANCE"}
	for _, env := range required {
		if os.Getenv(env) == "" {
			return fmt.Errorf("missing environment variable: %s", env)
		}
	}
	return nil
}

func (c *tableTestContext) iHaveInitializedTheServiceNowClient() error {
	instance := os.Getenv("SN_INSTANCE")
	authority := credentials.NewInstanceAuthority(instance)
	cred, err := credentials.NewROPCCredential(
		os.Getenv("SN_CLIENT_ID"),
		os.Getenv("SN_CLIENT_SECRET"),
		os.Getenv("SN_USERNAME"),
		os.Getenv("SN_PASSWORD"),
		authority,
		nil,
	)
	if err != nil {
		return err
	}

	client, err := sdk.NewServiceNowClient2(cred, instance)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *tableTestContext) iRequestAllIncidentsFromTheTable(tableName string) error {
	resp, err := c.client.Now2().TableV2(tableName).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *tableTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *tableTestContext) theResultsShouldContainAtLeastRecord(minCount int) error {
	collection, ok := c.response.(newInternal.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("expected a collection response, but got %T", c.response)
	}
	results, err := collection.GetResult()
	if err != nil {
		return err
	}

	if len(results) < minCount {
		return fmt.Errorf("expected at least %d records, got %d", minCount, len(results))
	}
	return nil
}

func (c *tableTestContext) eachRecordShouldHaveAValidSysID() error {
	collection, ok := c.response.(newInternal.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("expected a collection response, but got %T", c.response)
	}
	results, _ := collection.GetResult()

	for _, entry := range results {
		sysID, err := entry.GetSysID()
		if err != nil || sysID == nil || *sysID == "" {
			return fmt.Errorf("found record with missing or invalid sys_id")
		}
	}
	return nil
}

func (c *tableTestContext) iHaveAtLeastIncidentInTheTable(minCount int, tableName string) error {
	resp, err := c.client.Now2().TableV2(tableName).Get(context.Background(), nil)
	if err != nil {
		return err
	}
	results, err := resp.GetResult()
	if err != nil {
		return err
	}
	if len(results) < minCount {
		return fmt.Errorf("not enough incidents found: %d", len(results))
	}
	sysID, _ := results[0].GetSysID()
	c.lastSysID = *sysID
	return nil
}

func (c *tableTestContext) iRequestTheIncidentByItsSysID() error {
	resp, err := c.client.Now2().TableV2("incident").ById(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *tableTestContext) theResultShouldHaveTheCorrectSysID() error {
	item, ok := c.response.(newInternal.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("expected an item response, but got %T", c.response)
	}
	result, err := item.GetResult()
	if err != nil {
		return err
	}
	sysID, _ := result.GetSysID()
	if *sysID != c.lastSysID {
		return fmt.Errorf("expected sys_id %s, got %s", c.lastSysID, *sysID)
	}
	return nil
}

func (c *tableTestContext) iCreateANewIncidentWithDescription(description string) error {
	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", description)

	resp, err := c.client.Now2().TableV2("incident").Post(context.Background(), record, nil)
	c.response = resp
	c.err = err
	if err == nil {
		result, _ := resp.GetResult()
		sysID, _ := result.GetSysID()
		c.lastSysID = *sysID
	}
	return nil
}

func (c *tableTestContext) theCreatedRecordShouldHaveAValidSysID() error {
	item, ok := c.response.(newInternal.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("expected an item response, but got %T", c.response)
	}
	result, err := item.GetResult()
	if err != nil {
		return err
	}
	sysID, err := result.GetSysID()
	if err != nil || sysID == nil || *sysID == "" {
		return fmt.Errorf("missing or invalid sys_id in created record")
	}
	return nil
}

func (c *tableTestContext) iUpdateTheIncidentDescriptionTo(description string) error {
	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", description)

	resp, err := c.client.Now2().TableV2("incident").ById(c.lastSysID).Put(context.Background(), record, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *tableTestContext) theRecordShouldHaveDescription(description string) error {
	item, ok := c.response.(newInternal.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("expected an item response, but got %T", c.response)
	}
	result, _ := item.GetResult()
	element, err := result.Get("short_description")
	if err != nil {
		return err
	}
	val, _ := element.GetValue()
	strVal, _ := val.GetStringValue()
	if *strVal != description {
		return fmt.Errorf("expected description %s, got %s", description, *strVal)
	}
	return nil
}

func (c *tableTestContext) iDeleteTheCreatedIncident() error {
	err := c.client.Now2().TableV2("incident").ById(c.lastSysID).Delete(context.Background(), nil)
	c.err = err
	return nil
}

func (c *tableTestContext) iRequestTheDeletedIncidentByItsSysID() error {
	resp, err := c.client.Now2().TableV2("incident").ById(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *tableTestContext) theResponseShouldBeA404Error() error {
	if c.err == nil {
		return fmt.Errorf("expected a 404 error, but got no error")
	}
	return nil
}

func InitializeTableScenario(ctx *godog.ScenarioContext) {
	tc := &tableTestContext{}

	ctx.Step(`^I have a valid ServiceNow instance and credentials$`, tc.iHaveAValidServiceNowInstanceAndCredentials)
	ctx.Step(`^I have initialized the ServiceNow client$`, tc.iHaveInitializedTheServiceNowClient)
	ctx.Step(`^I request all incidents from the "([^"]*)" table$`, tc.iRequestAllIncidentsFromTheTable)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the results should contain at least (\d+) record$`, tc.theResultsShouldContainAtLeastRecord)
	ctx.Step(`^each record should have a valid "sys_id"$`, tc.eachRecordShouldHaveAValidSysID)
	ctx.Step(`^I have at least (\d+) incident in the "([^"]*)" table$`, tc.iHaveAtLeastIncidentInTheTable)
	ctx.Step(`^I request the incident by its "sys_id"$`, tc.iRequestTheIncidentByItsSysID)
	ctx.Step(`^the result should have the correct "sys_id"$`, tc.theResultShouldHaveTheCorrectSysID)

	ctx.Step(`^I create a new incident with description "([^"]*)"$`, tc.iCreateANewIncidentWithDescription)
	ctx.Step(`^the created record should have a valid "sys_id"$`, tc.theCreatedRecordShouldHaveAValidSysID)
	ctx.Step(`^I update the incident description to "([^"]*)"$`, tc.iUpdateTheIncidentDescriptionTo)
	ctx.Step(`^the record should have description "([^"]*)"$`, tc.theRecordShouldHaveDescription)
	ctx.Step(`^I delete the created incident$`, tc.iDeleteTheCreatedIncident)
	ctx.Step(`^I request the deleted incident by its "sys_id"$`, tc.iRequestTheDeletedIncidentByItsSysID)
	ctx.Step(`^the response should be a 404 error$`, tc.theResponseShouldBeA404Error)
}

func TestTableFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeTableScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/incident_retrieval.feature", "features/table_crud.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
