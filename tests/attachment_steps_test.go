//go:build integration

package tests

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/jarcoal/httpmock"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
)

func (c *testContext) iRequestAllAttachments() error {
	resp, err := c.client.Now().Attachment().Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theAttachmentResultsShouldContainAtLeastRecords(minCount int) error {
	collection, ok := c.response.(attachmentapi.AttachmentCollectionResponse2)
	if !ok {
		return fmt.Errorf("resp is not attachmentapi.AttachmentCollectionResponse2, got %T", c.response)
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

func (c *testContext) iHaveAtLeastAttachmentInTheInstance(minCount int) error {
	resp, err := c.client.Now().Attachment().Get(context.Background(), nil)
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

func (c *testContext) iRequestTheAttachmentByItsSysID() error {
	resp, err := c.client.Now().Attachment().ByID(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theAttachmentResultShouldHaveTheCorrectSysID() error {
	response, ok := c.response.(model.ServiceNowItemResponse[attachmentapi.Attachment2])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowItemResponse[attachmentapi.Attachment2], got %T", c.response)
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

func (c *testContext) iUploadTheFileFromTheResourcesDirectoryToTheRecord(fileName string) error {
	var data []byte
	var err error

	if httpmock.Disabled() {
		path := "resources/" + fileName
		data, err = os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file from resources: %w", err)
		}
	} else {
		data = []byte("test content")
	}

	media := attachmentapi.NewMedia("text/plain", data)
	config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			FileName:   &fileName,
			TableName:  &c.tableName,
			TableSysID: &c.lastSysID,
		},
	}
	resp, err := c.client.Now().Attachment().File().Post(context.Background(), media, config)
	c.response = resp
	c.err = err
	if err == nil && !utils.IsNil(resp) {
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

func (c *testContext) theCreatedAttachmentShouldHaveAValidSysID() error {
	response, ok := c.response.(model.ServiceNowItemResponse[attachmentapi.File])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowItemResponse[attachmentapi.File], got %T", c.response)
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

func (c *testContext) theAttachmentFilenameShouldBe(fileName string) error {
	response, ok := c.response.(model.ServiceNowItemResponse[attachmentapi.File])
	if !ok {
		return fmt.Errorf("resp is not model.ServiceNowItemResponse[attachmentapi.File], got %T", c.response)
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

func (c *testContext) iRequestTheContentOfTheCreatedAttachment() error {
	resp, err := c.client.Now().Attachment().ByID(c.lastSysID).File().Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func (c *testContext) theRetrievedContentShouldMatchTheOriginalFile(fileName string) error {
	var originalData []byte
	var err error

	if httpmock.Disabled() {
		path := "resources/" + fileName
		originalData, err = os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read original file: %w", err)
		}
	} else {
		originalData = []byte("test content")
	}

	fileWithContent, ok := c.response.(attachmentapi.FileWithContent)
	if !ok {
		return fmt.Errorf("resp is not attachmentapi.FileWithContent, got %T", c.response)
	}

	retrievedData, err := fileWithContent.GetContent()
	if err != nil {
		return fmt.Errorf("failed to get retrieved content: %w", err)
	}

	if !bytes.Equal(originalData, retrievedData) {
		return fmt.Errorf("retrieved content does not match original file")
	}
	return nil
}

func (c *testContext) iDeleteTheCreatedAttachment() error {
	err := c.client.Now().Attachment().ByID(c.lastSysID).Delete(context.Background(), nil)
	c.err = err
	return nil
}

func (c *testContext) iRequestTheDeletedAttachmentByItsSysID() error {
	if isOffline() {
		instance := os.Getenv("SN_INSTANCE")
		baseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/attachment", instance)
		url := fmt.Sprintf("%s/%s", baseURL, c.lastSysID)
		httpmock.RegisterResponder("GET", url,
			httpmock.NewStringResponder(404, `{"error":{"message":"No Record found","detail":""},"status":"failure"}`))
	}
	resp, err := c.client.Now().Attachment().ByID(c.lastSysID).Get(context.Background(), nil)
	c.response = resp
	c.err = err
	return nil
}

func InitializeAttachmentScenario(ctx *godog.ScenarioContext, tc *testContext) {
	ctx.Step(`^I request all attachments$`, tc.iRequestAllAttachments)
	ctx.Step(`^the attachment results should contain at least (\d+) records$`, tc.theAttachmentResultsShouldContainAtLeastRecords)
	ctx.Step(`^I have at least (\d+) attachment in the instance$`, tc.iHaveAtLeastAttachmentInTheInstance)
	ctx.Step(`^I request the attachment by its "sys_id"$`, tc.iRequestTheAttachmentByItsSysID)
	ctx.Step(`^the attachment result should have the correct "sys_id"$`, tc.theAttachmentResultShouldHaveTheCorrectSysID)

	ctx.Step(`^I upload the file "([^"]*)" from the resources directory to the record$`, tc.iUploadTheFileFromTheResourcesDirectoryToTheRecord)
	ctx.Step(`^the created attachment should have a valid "sys_id"$`, tc.theCreatedAttachmentShouldHaveAValidSysID)
	ctx.Step(`^the attachment filename should be "([^"]*)"$`, tc.theAttachmentFilenameShouldBe)
	ctx.Step(`^I request the content of the created attachment$`, tc.iRequestTheContentOfTheCreatedAttachment)
	ctx.Step(`^the retrieved content should match the original file "([^"]*)"$`, tc.theRetrievedContentShouldMatchTheOriginalFile)
	ctx.Step(`^I delete the created attachment$`, tc.iDeleteTheCreatedAttachment)
	ctx.Step(`^I request the deleted attachment by its "sys_id"$`, tc.iRequestTheDeletedAttachmentByItsSysID)
}

func TestAttachmentFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/attachment_api.feature", "features/attachment_crud.feature", "features/attachment_content.feature"},
			Tags:     "integration",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
