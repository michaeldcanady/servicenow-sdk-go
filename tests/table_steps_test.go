//go:build integration

package tests

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

type tableTestContext struct {
	client       *sdk.ServiceNowClient
	response     interface{} // Generic response to support either collection or item
	err          error
	lastSysID    string
	pageSize     int
	fields       []string
	pagesReached int
	totalItems   int
}

func (c *tableTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../.env")
	return nil
}

func (c *tableTestContext) iHaveInitializedTheServiceNowClient() error {
	instance := os.Getenv("SN_INSTANCE")
	if instance == "" {
		instance = "mock_instance"
	}

	cred := credentials.NewUsernamePasswordCredential(
		os.Getenv("SN_USERNAME"),
		os.Getenv("SN_PASSWORD"),
	)

	client, err := sdk.NewServiceNowClient2WithHTTPClient(cred, instance, getHttpClient())
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
		if snErr, ok := c.err.(*core.ServiceNowError); ok {
			return fmt.Errorf("expected no error, but got: %v (Message: %s, Detail: %s, Status: %s)",
				c.err, snErr.Exception.Message, snErr.Exception.Detail, snErr.Status)
		}
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

func (c *tableTestContext) iPatchTheIncidentDescriptionTo(description string) error {
	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", description)

	resp, err := c.client.Now2().TableV2("incident").ById(c.lastSysID).Patch(context.Background(), record, nil)
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
	if isOffline() {
		baseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/table/incident/", os.Getenv("SN_INSTANCE"))
		httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(baseURL+`[a-zA-Z0-9_]+$`),
			httpmock.NewStringResponder(404, `{"error":{"message":"No Record found","detail":""},"status":"failure"}`))
	}

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

func (c *tableTestContext) iRequestIncidentsWithQueryAndLimit(query string, limit int) error {
	config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
			Query: query,
			Limit: limit,
		},
	}
	resp, err := c.client.Now2().TableV2("incident").Get(context.Background(), config)
	c.response = resp
	c.err = err
	return nil
}

func (c *tableTestContext) theResultsShouldContainAtMostRecords(maxCount int) error {
	collection, ok := c.response.(newInternal.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("expected a collection response, but got %T", c.response)
	}
	results, _ := collection.GetResult()
	if len(results) > maxCount {
		return fmt.Errorf("expected at most %d records, got %d", maxCount, len(results))
	}
	return nil
}

func (c *tableTestContext) eachRecordShouldHaveSetTo(field, value string) error {
	collection, ok := c.response.(newInternal.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("expected a collection response, but got %T", c.response)
	}
	results, _ := collection.GetResult()
	for _, record := range results {
		elem, err := record.Get(field)
		if err != nil {
			return err
		}
		val, _ := elem.GetValue()
		strVal, _ := val.GetStringValue()
		if *strVal != value {
			return fmt.Errorf("expected %s to be %s, got %s", field, value, *strVal)
		}
	}
	return nil
}

func (c *tableTestContext) iRequestIncidentsSortedByDescending(field string) error {
	config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
			Query: "ORDERBYDESC" + field,
		},
	}
	resp, err := c.client.Now2().TableV2("incident").Get(context.Background(), config)
	c.response = resp
	c.err = err
	return nil
}

func (c *tableTestContext) theRecordsShouldBeInDescendingOrderOf(field string) error {
	collection, ok := c.response.(newInternal.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("expected a collection response, but got %T", c.response)
	}
	results, _ := collection.GetResult()
	if len(results) < 2 {
		return nil // Can't check order with less than 2
	}
	for i := 0; i < len(results)-1; i++ {
		elem1, _ := results[i].Get(field)
		val1, _ := elem1.GetValue()
		str1, _ := val1.GetStringValue()

		elem2, _ := results[i+1].Get(field)
		val2, _ := elem2.GetValue()
		str2, _ := val2.GetStringValue()

		if *str1 < *str2 {
			return fmt.Errorf("records not in descending order: %s < %s", *str1, *str2)
		}
	}
	return nil
}

func (c *tableTestContext) iSetThePageSizeTo(pageSize int) error {
	c.pageSize = pageSize
	return nil
}

func (c *tableTestContext) iSetTheFieldsTo(fields []string) error {
	c.fields = fields
	return nil
}

func (c *tableTestContext) iUseTheTablePageIteratorToFetchRecords() error {
	config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
			Limit:  c.pageSize,
			Fields: c.fields,
		},
	}
	resp, err := c.client.Now2().TableV2("incident").Get(context.Background(), config)
	if err != nil {
		return err
	}

	iterator, err := tableapi.NewDefaultTablePageIterator(resp, c.client.RequestAdapter)
	if err != nil {
		return err
	}

	c.pagesReached = 0
	c.totalItems = 0

	itemsInCurrentPage := 0
	err = iterator.Iterate(context.Background(), false, func(record *tableapi.TableRecord) bool {
		if itemsInCurrentPage == 0 {
			c.pagesReached++
		}
		c.totalItems++
		itemsInCurrentPage++
		// If we've processed all items in current page, reset itemsInCurrentPage
		if itemsInCurrentPage >= c.pageSize {
			itemsInCurrentPage = 0
			if c.pagesReached >= 2 {
				return false // Stop after 2 pages
			}
		}
		return true
	})

	return err
}

func (c *tableTestContext) iShouldBeAbleToReachTheSecondPage() error {
	if c.pagesReached < 2 {
		return fmt.Errorf("only reached %d pages, expected at least 2", c.pagesReached)
	}
	return nil
}

func (c *tableTestContext) theTotalCountOfRecordsRetrievedShouldBeGreaterThan(count int) error {
	if c.totalItems <= count {
		return fmt.Errorf("expected more than %d items, got %d", count, c.totalItems)
	}
	return nil
}

func InitializeTableScenario(ctx *godog.ScenarioContext) {
	tc := &tableTestContext{}

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
	ctx.Step(`^I patch the incident description to "([^"]*)"$`, tc.iPatchTheIncidentDescriptionTo)
	ctx.Step(`^the record should have description "([^"]*)"$`, tc.theRecordShouldHaveDescription)
	ctx.Step(`^I delete the created incident$`, tc.iDeleteTheCreatedIncident)
	ctx.Step(`^I request the deleted incident by its "sys_id"$`, tc.iRequestTheDeletedIncidentByItsSysID)
	ctx.Step(`^the response should be a 404 error$`, tc.theResponseShouldBeA404Error)

	ctx.Step(`^I request incidents with query "([^"]*)" and limit (\d+)$`, tc.iRequestIncidentsWithQueryAndLimit)
	ctx.Step(`^the results should contain at most (\d+) records$`, tc.theResultsShouldContainAtMostRecords)
	ctx.Step(`^each record should have "([^"]*)" set to "([^"]*)"$`, tc.eachRecordShouldHaveSetTo)
	ctx.Step(`^I request incidents sorted by "([^"]*)" descending$`, tc.iRequestIncidentsSortedByDescending)
	ctx.Step(`^the records should be in descending order of "([^"]*)"$`, tc.theRecordsShouldBeInDescendingOrderOf)

	ctx.Step(`^I set the page size to (\d+)$`, tc.iSetThePageSizeTo)
	ctx.Step(`^I use the Table PageIterator to fetch records$`, tc.iUseTheTablePageIteratorToFetchRecords)
	ctx.Step(`^I should be able to reach the second page$`, tc.iShouldBeAbleToReachTheSecondPage)
	ctx.Step(`^the total count of records retrieved should be greater than (\d+)$`, tc.theTotalCountOfRecordsRetrievedShouldBeGreaterThan)
}

func TestTableFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeTableScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/incident_retrieval.feature", "features/table_crud.feature", "features/table_query.feature", "features/table_pagination.feature", "features/table_patch.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
