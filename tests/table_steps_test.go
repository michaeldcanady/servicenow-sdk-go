//go:build integration

package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func (c *testContext) iRequestAllRecordsFromTheTable(tableName string) error {
	c.tableName = tableName
	resp, err := c.client.Now().Table(tableName).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	fmt.Printf("DEBUG: iRequestAllRecordsFromTheTable response type: %T\n", resp)
	return nil
}

func (c *testContext) theResultsShouldContainAtLeastRecord(minCount int) error {
	fmt.Printf("DEBUG: theResultsShouldContainAtLeastRecord c.response type: %T\n", c.response)
	collection, ok := c.response.(model.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowCollectionResponse[*tableapi.TableRecord], got %T", c.response)
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

func (c *testContext) eachRecordShouldHaveAValidSysID() error {
	collection, ok := c.response.(model.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowCollectionResponse[*tableapi.TableRecord], got %T", c.response)
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

func (c *testContext) iHaveAtLeastRecordInTheTable(minCount int, tableName string) error {
	c.tableName = tableName
	resp, err := c.client.Now().Table(tableName).Get(context.Background(), nil)
	if err != nil {
		return err
	}
	results, err := resp.GetResult()
	if err != nil {
		return err
	}
	if len(results) < minCount {
		return fmt.Errorf("not enough records found: %d", len(results))
	}
	sysID, _ := results[0].GetSysID()
	c.lastSysID = *sysID
	return nil
}

func (c *testContext) iCreateANewRecordWithDescription(tableName, description string) error {
	c.tableName = tableName
	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", description)

	resp, err := c.client.Now().Table(tableName).Post(context.Background(), record, nil)
	c.response = resp
	c.err = err
	if err == nil {
		result, _ := resp.GetResult()
		sysID, _ := result.GetSysID()
		c.lastSysID = *sysID
	}
	return nil
}

func (c *testContext) theCreatedRecordShouldHaveAValidSysID() error {
	item, ok := c.response.(model.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowItemResponse[*tableapi.TableRecord], got %T", c.response)
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

func (c *testContext) iUpdateTheRecordDescriptionTo(description string) error {
	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", description)

	resp, err := c.client.Now().Table(c.tableName).ById(c.lastSysID).Put(context.Background(), record, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) iPatchTheRecordDescriptionTo(description string) error {
	record := tableapi.NewTableRecord()
	_ = record.SetValue("short_description", description)

	resp, err := c.client.Now().Table(c.tableName).ById(c.lastSysID).Patch(context.Background(), record, nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theRecordShouldHaveDescription(description string) error {
	item, ok := c.response.(model.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowItemResponse[*tableapi.TableRecord], got %T", c.response)
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

func (c *testContext) iRequestRecordsWithQueryAndLimit(tableName, query string, limit int) error {
	c.tableName = tableName
	config := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Query: query,
			Limit: limit,
		},
	}
	resp, err := c.client.Now().Table(tableName).Get(context.Background(), config)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theResultsShouldContainAtMostRecords(maxCount int) error {
	collection, ok := c.response.(model.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowCollectionResponse[*tableapi.TableRecord], got %T", c.response)
	}
	results, _ := collection.GetResult()
	if len(results) > maxCount {
		return fmt.Errorf("expected at most %d records, got %d", maxCount, len(results))
	}
	return nil
}

func (c *testContext) eachRecordShouldHaveSetTo(field, value string) error {
	collection, ok := c.response.(model.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowCollectionResponse[*tableapi.TableRecord], got %T", c.response)
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

func (c *testContext) iRequestRecordsSortedByDescending(tableName, field string) error {
	c.tableName = tableName
	config := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Query: "ORDERBYDESC" + field,
		},
	}
	resp, err := c.client.Now().Table(tableName).Get(context.Background(), config)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theRecordsShouldBeInDescendingOrderOf(field string) error {
	collection, ok := c.response.(model.ServiceNowCollectionResponse[*tableapi.TableRecord])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowCollectionResponse[*tableapi.TableRecord], got %T", c.response)
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

func (c *testContext) iSetThePageSizeTo(pageSize int) error {
	c.pageSize = pageSize
	return nil
}

func (c *testContext) iUseTheTablePageIteratorToFetchRecords() error {
	config := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Limit:  c.pageSize,
			Fields: c.fields,
		},
	}
	resp, err := c.client.Now().Table(c.tableName).Get(context.Background(), config)
	if err != nil {
		return err
	}

	iterator, err := tableapi.NewDefaultTablePageIterator(resp, c.client.GetRequestAdapter())
	if err != nil {
		return err
	}

	pagesReached := 0
	totalItems := 0

	itemsInCurrentPage := 0
	err = iterator.Iterate(context.Background(), false, func(record *tableapi.TableRecord) bool {
		if itemsInCurrentPage == 0 {
			pagesReached++
		}
		totalItems++
		itemsInCurrentPage++
		// If we've processed all items in current page, reset itemsInCurrentPage
		if itemsInCurrentPage >= c.pageSize {
			itemsInCurrentPage = 0
			if pagesReached >= 2 {
				return false // Stop after 2 pages
			}
		}
		return true
	})

	if err == nil {
		if pagesReached < 2 {
			return fmt.Errorf("only reached %d pages, expected at least 2", pagesReached)
		}
	}

	return err
}

func (c *testContext) iShouldBeAbleToReachTheSecondPage() error {
	return nil
}

func InitializeTableScenario(ctx *godog.ScenarioContext, tc *testContext) {
	ctx.Step(`^I request all records from the "([^"]*)" table$`, tc.iRequestAllRecordsFromTheTable)
	ctx.Step(`^the results should contain at least (\d+) record$`, tc.theResultsShouldContainAtLeastRecord)
	ctx.Step(`^each record should have a valid "sys_id"$`, tc.eachRecordShouldHaveAValidSysID)
	ctx.Step(`^I have at least (\d+) record in the "([^"]*)" table$`, tc.iHaveAtLeastRecordInTheTable)

	ctx.Step(`^I create a new record in "([^"]*)" with description "([^"]*)"$`, tc.iCreateANewRecordWithDescription)
	ctx.Step(`^the created record should have a valid "sys_id"$`, tc.theCreatedRecordShouldHaveAValidSysID)
	ctx.Step(`^I update the record description to "([^"]*)"$`, tc.iUpdateTheRecordDescriptionTo)
	ctx.Step(`^I patch the record description to "([^"]*)"$`, tc.iPatchTheRecordDescriptionTo)
	ctx.Step(`^the record should have description "([^"]*)"$`, tc.theRecordShouldHaveDescription)

	ctx.Step(`^I request records from "([^"]*)" with query "([^"]*)" and limit (\d+)$`, tc.iRequestRecordsWithQueryAndLimit)
	ctx.Step(`^the results should contain at most (\d+) records$`, tc.theResultsShouldContainAtMostRecords)
	ctx.Step(`^each record should have "([^"]*)" set to "([^"]*)"$`, tc.eachRecordShouldHaveSetTo)
	ctx.Step(`^I request records from "([^"]*)" sorted by "([^"]*)" descending$`, tc.iRequestRecordsSortedByDescending)
	ctx.Step(`^the records should be in descending order of "([^"]*)"$`, tc.theRecordsShouldBeInDescendingOrderOf)

	ctx.Step(`^I set the page size to (\d+)$`, tc.iSetThePageSizeTo)
	ctx.Step(`^I use the Table PageIterator to fetch records from "([^"]*)"$`, func(tableName string) error {
		tc.tableName = tableName
		return tc.iUseTheTablePageIteratorToFetchRecords()
	})
	ctx.Step(`^I should be able to reach the second page$`, tc.iShouldBeAbleToReachTheSecondPage)
}

func TestTableFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features/table_query.feature"},
			//Paths:    []string{"features/incident_retrieval.feature", "features/table_crud.feature", "features/table_query.feature", "features/table_pagination.feature", "features/table_patch.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
