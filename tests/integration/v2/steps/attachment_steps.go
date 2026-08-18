package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/michaeldcanady/servicenow-sdk-go/attachmentapi"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/tableapi"
	"github.com/michaeldcanady/servicenow-sdk-go/tests/integration/v2/support"
)

type attachmentSteps struct{}

func (s *attachmentSteps) iRequestAllAttachments(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Attachment().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *attachmentSteps) iHaveAtLeastAttachmentInTheInstance(ctx context.Context, minCount int) (context.Context, error) {
	w := support.WorldFrom(ctx)

	resp, err := w.Client.Now().Attachment().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	results, err := resp.GetResult()
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	if len(results) < minCount {
		w.Err = fmt.Errorf("expected at least %d attachments, got %d", minCount, len(results))
		return support.WithWorld(ctx, w), nil
	}

	if len(results) > 0 {
		sysID, err := results[0].GetSysID()
		if err == nil && sysID != nil && *sysID != "" {
			w.LastSysID = *sysID
		}
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *attachmentSteps) iRequestTheAttachmentBySysID(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_attach_id_1"
		} else {
			return ctx, fmt.Errorf("no attachment sys_id available — upload or list first")
		}
	}

	resp, err := w.Client.Now().Attachment().ByID(sysID).Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *attachmentSteps) iUploadAFileToAnIncident(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	incident := tableapi.NewTableRecord()
	desc := "Attachment Test Incident"
	incident.SetValue("short_description", desc)
	incidentResp, err := w.Client.Now().Table("incident").Post(ctx, incident, nil)
	if err != nil {
		w.Err = fmt.Errorf("failed to create test incident: %w", err)
		return support.WithWorld(ctx, w), nil
	}
	itemResp, ok := incidentResp.(core.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		w.Err = fmt.Errorf("unexpected incident response type: %T", incidentResp)
		return support.WithWorld(ctx, w), nil
	}
	incidentRecord, err := itemResp.GetResult()
	if err != nil {
		w.Err = fmt.Errorf("failed to get incident result: %w", err)
		return support.WithWorld(ctx, w), nil
	}
	incidentSysID, err := incidentRecord.GetSysID()
	if err != nil {
		w.Err = fmt.Errorf("failed to get incident sys_id: %w", err)
		return support.WithWorld(ctx, w), nil
	}
	if incidentSysID == nil || *incidentSysID == "" {
		w.Err = fmt.Errorf("created incident has no sys_id")
		return support.WithWorld(ctx, w), nil
	}
	w.TrackCreation(*incidentSysID)

	media := attachmentapi.NewMedia("text/plain", []byte("test content"))
	tableName := "incident"
	tableSysID := *incidentSysID
	fileName := "test.txt"

	config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			TableName:  &tableName,
			TableSysID: &tableSysID,
			FileName:   &fileName,
		},
	}

	resp, err := w.Client.Now().Attachment().File().Post(ctx, media, config)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp

	fileResp, err := resp.GetResult()
	if err == nil && fileResp != nil {
		sysID, err := fileResp.GetSysID()
		if err == nil && sysID != nil && *sysID != "" {
			w.LastSysID = *sysID
		}
	}

	return support.WithWorld(ctx, w), nil
}

func (s *attachmentSteps) iUploadASecondFileToAnIncident(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	incident := tableapi.NewTableRecord()
	desc := "Attachment Test Incident 2"
	incident.SetValue("short_description", desc)
	incidentResp, err := w.Client.Now().Table("incident").Post(ctx, incident, nil)
	if err != nil {
		w.Err = fmt.Errorf("failed to create test incident: %w", err)
		return support.WithWorld(ctx, w), nil
	}
	itemResp, ok := incidentResp.(core.ServiceNowItemResponse[*tableapi.TableRecord])
	if !ok {
		w.Err = fmt.Errorf("unexpected incident response type: %T", incidentResp)
		return support.WithWorld(ctx, w), nil
	}
	incidentRecord, err := itemResp.GetResult()
	if err != nil {
		w.Err = fmt.Errorf("failed to get incident result: %w", err)
		return support.WithWorld(ctx, w), nil
	}
	incidentSysID, err := incidentRecord.GetSysID()
	if err != nil {
		w.Err = fmt.Errorf("failed to get incident sys_id: %w", err)
		return support.WithWorld(ctx, w), nil
	}
	if incidentSysID == nil || *incidentSysID == "" {
		w.Err = fmt.Errorf("created incident has no sys_id")
		return support.WithWorld(ctx, w), nil
	}
	w.TrackCreation(*incidentSysID)

	media := attachmentapi.NewMedia("text/plain", []byte("second test content"))
	tableName := "incident"
	tableSysID := *incidentSysID
	fileName := "test2.txt"

	config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			TableName:  &tableName,
			TableSysID: &tableSysID,
			FileName:   &fileName,
		},
	}

	resp, err := w.Client.Now().Attachment().File().Post(ctx, media, config)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *attachmentSteps) iRequestContentOfCreatedAttachment(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_attach_id_1"
		} else {
			return ctx, fmt.Errorf("no attachment sys_id available — upload first")
		}
	}

	resp, err := w.Client.Now().Attachment().ByID(sysID).File().Get(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	w.Response = resp
	return support.WithWorld(ctx, w), nil
}

func (s *attachmentSteps) iDeleteTheAttachment(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)

	sysID := w.LastSysID
	if sysID == "" {
		if support.IsOffline() {
			sysID = "mock_attach_id_1"
		} else {
			return ctx, fmt.Errorf("no attachment sys_id available — upload first")
		}
	}

	err := w.Client.Now().Attachment().ByID(sysID).Delete(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	return support.WithWorld(ctx, w), nil
}

func (s *attachmentSteps) iDeleteTheAttachmentWithSysID(ctx context.Context, sysID string) (context.Context, error) {
	w := support.WorldFrom(ctx)

	err := w.Client.Now().Attachment().ByID(sysID).Delete(ctx, nil)
	if err != nil {
		w.Err = err
		return support.WithWorld(ctx, w), nil
	}

	return support.WithWorld(ctx, w), nil
}

func (s *attachmentSteps) theResponseShouldContainAtLeastAttachments(ctx context.Context, minCount int) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil, expected at least %d attachments", minCount)
	}

	colResp, ok := w.Response.(core.ServiceNowCollectionResponse[*attachmentapi.Attachment])
	if !ok {
		return ctx, fmt.Errorf("response is not an attachment collection response")
	}

	results, err := colResp.GetResult()
	if err != nil {
		return ctx, fmt.Errorf("failed to get results: %v", err)
	}

	if len(results) < minCount {
		return ctx, fmt.Errorf("expected at least %d attachments, got %d", minCount, len(results))
	}

	return ctx, nil
}

func (s *attachmentSteps) theAttachmentShouldHaveAFileName(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil")
	}
	return ctx, nil
}

func (s *attachmentSteps) theAttachmentShouldHaveAContentType(ctx context.Context) (context.Context, error) {
	w := support.WorldFrom(ctx)
	if w.Response == nil {
		return ctx, fmt.Errorf("response is nil")
	}
	return ctx, nil
}

// InitializeAttachmentScenario registers all attachment step definitions.
func InitializeAttachmentScenario(sc *godog.ScenarioContext) {
	s := &attachmentSteps{}

	RegisterSharedSteps(sc)

	sc.Step(`^I request all attachments$`, s.iRequestAllAttachments)
	sc.Step(`^I have at least (\d+) attachment in the instance$`, s.iHaveAtLeastAttachmentInTheInstance)
	sc.Step(`^I request the attachment by its sys_id$`, s.iRequestTheAttachmentBySysID)
	sc.Step(`^I upload a file to an incident$`, s.iUploadAFileToAnIncident)
	sc.Step(`^I upload a second file to an incident$`, s.iUploadASecondFileToAnIncident)
	sc.Step(`^I request the content of the created attachment$`, s.iRequestContentOfCreatedAttachment)
	sc.Step(`^I delete the attachment$`, s.iDeleteTheAttachment)
	sc.Step(`^I delete the attachment with sys_id "([^"]*)"$`, s.iDeleteTheAttachmentWithSysID)
	sc.Step(`^the response should contain at least (\d+) attachment$`, s.theResponseShouldContainAtLeastAttachments)
	sc.Step(`^the attachment should have a file_name$`, s.theAttachmentShouldHaveAFileName)
	sc.Step(`^the attachment should have a content_type$`, s.theAttachmentShouldHaveAContentType)

	sc.Before(BeforeScenario)
	sc.After(AfterScenario)
}
