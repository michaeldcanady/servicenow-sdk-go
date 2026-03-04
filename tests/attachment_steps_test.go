package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
	sdk "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
)

type attachmentTestContext struct {
	client        *sdk.ServiceNowClient
	response      interface{} // Generic response to support either collection or item
	err           error
	lastSysID     string
	incidentSysID string
}

func (c *attachmentTestContext) iHaveAValidServiceNowInstanceAndCredentials() error {
	_ = godotenv.Load("../.env")
	required := []string{"SN_CLIENT_ID", "SN_USERNAME", "SN_PASSWORD", "SN_INSTANCE"}
	for _, env := range required {
		if os.Getenv(env) == "" {
			return fmt.Errorf("missing environment variable: %s", env)
		}
	}
	return nil
}

func (c *attachmentTestContext) iHaveInitializedTheServiceNowClient() error {
	instance := os.Getenv("SN_INSTANCE")

	cred := credentials.NewUsernamePasswordCredential(
		os.Getenv("SN_USERNAME"),
		os.Getenv("SN_PASSWORD"),
	)

	client, err := sdk.NewServiceNowClient2(cred, instance)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *attachmentTestContext) iRequestAllAttachments() error {
	resp, err := c.client.Now2().Attachment2().Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *attachmentTestContext) theResponseShouldNotBeAnError() error {
	if c.err != nil {
		return fmt.Errorf("expected no error, but got: %v", c.err)
	}
	return nil
}

func (c *attachmentTestContext) theResultsShouldContainAtLeastRecords(minCount int) error {
	collection, ok := c.response.(*attachmentapi.AttachmentCollectionResponse2Model)
	if !ok {
		return fmt.Errorf("expected a collection response, but got %T", c.response)
	}
	results, err := collection.GetResult()
	if err != nil {
		return fmt.Errorf("failed to retrieve results: %w", err)
	}

	if len(results) < minCount {
		return fmt.Errorf("expected at least %d records, got %d", minCount, len(results))
	}
	return nil
}

func (c *attachmentTestContext) iHaveAtLeastAttachmentInTheInstance(minCount int) error {
	resp, err := c.client.Now2().Attachment2().Get(context.Background(), nil)
	if err != nil {
		return err
	}

	results, err := resp.GetResult()
	if err != nil {
		return fmt.Errorf("failed to retrieve results: %w", err)
	}

	if count := len(results); count < minCount {
		return fmt.Errorf("not enough attachments found: %d", count)
	}
	sysID, err := results[0].GetSysID()
	if err != nil {
		return fmt.Errorf("failed to retrieve sys_id: %w", err)
	}

	c.lastSysID = *sysID
	return nil
}

func (c *attachmentTestContext) iRequestTheAttachmentByItsSysID() error {
	resp, err := c.client.Now2().Attachment2().ByID(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *attachmentTestContext) theResultShouldHaveTheCorrectSysID() error {
	response, ok := c.response.(newInternal.ServiceNowItemResponse[attachmentapi.Attachment2])
	if !ok {
		return fmt.Errorf("expected a ServiceNowItemResponse[Attachment2], but got %T", c.response)
	}

	item, err := response.GetResult()
	if err != nil {
		return fmt.Errorf("failed to retrieve result: %w", err)
	}

	sysID, err := item.GetSysID()
	if err != nil {
		return fmt.Errorf("failed to retrieve sys_id: %w", err)
	}

	if sysID == nil {
		return fmt.Errorf("sys_id is nil")
	}

	if *sysID != c.lastSysID {
		return fmt.Errorf("expected sys_id %s, got %s", c.lastSysID, *sysID)
	}
	return nil
}

func (c *attachmentTestContext) iHaveAnIncidentRecordInTheTable(tableName string) error {
	resp, err := c.client.Now2().TableV2(tableName).Get(context.Background(), nil)
	if err != nil {
		return err
	}
	results, err := resp.GetResult()
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return fmt.Errorf("no incidents found in table %s", tableName)
	}
	sysID, _ := results[0].GetSysID()
	c.incidentSysID = *sysID
	return nil
}

func (c *attachmentTestContext) iUploadTheFileFromTheResourcesDirectoryToTheIncident(fileName string) error {
	path := "resources/" + fileName
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file from resources: %w", err)
	}

	media := attachmentapi.NewMedia("text/plain", data)
	config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			FileName:   &fileName,
			TableName:  ptr("incident"),
			TableSysID: &c.incidentSysID,
		},
	}
	resp, err := c.client.Now2().Attachment2().File().Post(context.Background(), media, config)
	c.response = resp
	c.err = err
	if err == nil && !newInternal.IsNil(resp) {
		item, err := resp.GetResult()
		if err != nil {
			return fmt.Errorf("failed to get result from response: %w", err)
		}
		sysID, err := item.GetSysID()
		if err != nil {
			return fmt.Errorf("failed to get sys_id: %w", err)
		}
		if sysID == nil {
			return fmt.Errorf("sys_id is nil in response")
		}
		c.lastSysID = *sysID
	}
	return nil
}

func (c *attachmentTestContext) theCreatedAttachmentShouldHaveAValidSysID() error {
	response, ok := c.response.(newInternal.ServiceNowItemResponse[attachmentapi.File])
	if !ok {
		return fmt.Errorf("expected a ServiceNowItemResponse[File], but got %T", c.response)
	}
	item, err := response.GetResult()
	if err != nil {
		return fmt.Errorf("failed to get result: %w", err)
	}
	sysID, err := item.GetSysID()
	if err != nil || sysID == nil || *sysID == "" {
		return fmt.Errorf("missing or invalid sys_id in created attachment")
	}
	return nil
}

func (c *attachmentTestContext) theAttachmentFilenameShouldBe(fileName string) error {
	response, ok := c.response.(newInternal.ServiceNowItemResponse[attachmentapi.File])
	if !ok {
		return fmt.Errorf("expected a ServiceNowItemResponse[File], but got %T", c.response)
	}
	item, err := response.GetResult()
	if err != nil {
		return fmt.Errorf("failed to get result: %w", err)
	}
	name, _ := item.GetFileName()
	if *name != fileName {
		return fmt.Errorf("expected filename %s, got %s", fileName, *name)
	}
	return nil
}

func (c *attachmentTestContext) iDeleteTheCreatedAttachment() error {
	err := c.client.Now2().Attachment2().ByID(c.lastSysID).Delete(context.Background(), nil)
	c.err = err
	return nil
}

func (c *attachmentTestContext) iRequestTheDeletedAttachmentByItsSysID() error {
	resp, err := c.client.Now2().Attachment2().ByID(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *attachmentTestContext) theResponseShouldBeA404Error() error {
	if c.err == nil {
		return fmt.Errorf("expected a 404 error, but got no error")
	}
	return nil
}

func ptr[T any](v T) *T {
	return &v
}

func InitializeAttachmentScenario(ctx *godog.ScenarioContext) {
	tc := &attachmentTestContext{}

	ctx.Step(`^I have a valid ServiceNow instance and credentials$`, tc.iHaveAValidServiceNowInstanceAndCredentials)
	ctx.Step(`^I have initialized the ServiceNow client$`, tc.iHaveInitializedTheServiceNowClient)
	ctx.Step(`^I request all attachments$`, tc.iRequestAllAttachments)
	ctx.Step(`^the response should not be an error$`, tc.theResponseShouldNotBeAnError)
	ctx.Step(`^the results should contain at least (\d+) records$`, tc.theResultsShouldContainAtLeastRecords)
	ctx.Step(`^I have at least (\d+) attachment in the instance$`, tc.iHaveAtLeastAttachmentInTheInstance)
	ctx.Step(`^I request the attachment by its "sys_id"$`, tc.iRequestTheAttachmentByItsSysID)
	ctx.Step(`^the result should have the correct "sys_id"$`, tc.theResultShouldHaveTheCorrectSysID)

	ctx.Step(`^I have an incident record in the "([^"]*)" table$`, tc.iHaveAnIncidentRecordInTheTable)
	ctx.Step(`^I upload the file "([^"]*)" from the resources directory to the incident$`, tc.iUploadTheFileFromTheResourcesDirectoryToTheIncident)
	ctx.Step(`^the created attachment should have a valid "sys_id"$`, tc.theCreatedAttachmentShouldHaveAValidSysID)
	ctx.Step(`^the attachment filename should be "([^"]*)"$`, tc.theAttachmentFilenameShouldBe)
	ctx.Step(`^I delete the created attachment$`, tc.iDeleteTheCreatedAttachment)
	ctx.Step(`^I request the deleted attachment by its "sys_id"$`, tc.iRequestTheDeletedAttachmentByItsSysID)
	ctx.Step(`^the response should be a 404 error$`, tc.theResponseShouldBeA404Error)
}

func TestAttachmentFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeAttachmentScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/attachment_api.feature", "features/attachment_crud.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
