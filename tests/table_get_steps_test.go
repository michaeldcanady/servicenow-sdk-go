package tests

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/require"
)

func (c *sharedTestContext) iGETACollectionFromTheTable(ctx context.Context, tableName string) error {
	c.tableName = tableName
	c.executeReq = func() (serialization.Parsable, error) {
		return c.client.Now().Table(c.tableName).Get(context.Background(), nil)
	}
	return nil
}

func (c *sharedTestContext) iGETACollectionFromTheTableWithParameters(ctx context.Context, tableName string, table *godog.Table) error {
	t := godog.T(ctx)
	c.tableName = tableName

	config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{},
	}

	for _, row := range table.Rows {
		key := row.Cells[0].Value
		value := row.Cells[1].Value

		switch key {
		case "sysparm_limit":
			limit, err := strconv.Atoi(value)
			require.NoError(t, err)

			config.QueryParameters.Limit = limit
		case "sysparm_fields":
			fields := strings.Split(value, ",")
			config.QueryParameters.Fields = fields
		}
	}

	c.executeReq = func() (serialization.Parsable, error) {
		return c.client.Now().Table(c.tableName).Get(context.Background(), config)
	}
	return nil
}

func (c *sharedTestContext) theResponseShouldContainNRecords(ctx context.Context, count int) error {
	t := godog.T(ctx)
	colResp, ok := c.resp.(internal.ServiceNowCollectionResponse[*tableapi.TableRecord])
	require.True(t, ok)

	results, err := colResp.GetResult()
	require.NoError(t, err)

	require.Len(t, results, count)
	return nil
}

func (c *sharedTestContext) iGETTheRecordFromTheTableBySysID(ctx context.Context, tableName string) error {
	c.tableName = tableName
	c.executeReq = func() (serialization.Parsable, error) {
		return c.client.Now().Table(c.tableName).ById(c.sysID).Get(context.Background(), nil)
	}
	return nil
}

func (c *sharedTestContext) theResponseShouldContainTheRequestedRecord(ctx context.Context) error {
	t := godog.T(ctx)
	itemResp, ok := c.resp.(internal.ServiceNowItemResponse[*tableapi.TableRecord])
	require.True(t, ok)

	result, err := itemResp.GetResult()
	require.NoError(t, err)

	sysID, err := result.GetSysID()
	require.NoError(t, err)
	require.NotNil(t, sysID)

	require.Equal(t, c.sysID, *sysID)
	return nil
}

func InitializeGETOperationsScenario(ctx *godog.ScenarioContext) {
	feat := &sharedTestContext{}

	RegisterSharedSteps(ctx, feat)
	InitializeBasicAuthenticationScenario(ctx)

	ctx.Given(`^a record exists in the "([^"]*)" table$`, feat.aRecordExistsInTheTable)
	ctx.When(`^I GET a collection from the "([^"]*)" table$`, feat.iGETACollectionFromTheTable)
	ctx.When(`^I GET a collection from the "([^"]*)" table with parameters:$`, feat.iGETACollectionFromTheTableWithParameters)
	ctx.Then(`^the response should contain (\d+) records$`, feat.theResponseShouldContainNRecords)
	ctx.When(`^I GET the record from the "([^"]*)" table by Sys ID$`, feat.iGETTheRecordFromTheTableBySysID)
	ctx.Then(`^the response should contain the requested record$`, feat.theResponseShouldContainTheRequestedRecord)
	// Reuse step from collection operations if needed
	//ctx.Then(`^the response should contain a list of records$`, feat.theResponseShouldContainAListOfRecords)
}

func TestGETOperationsFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeGETOperationsScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/table_api/get.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
