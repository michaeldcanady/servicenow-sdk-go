package tableapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/abstraction"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (c *MockClient) Send(requestInfo *abstraction.RequestInformation, errorMapping abstraction.ErrorMapping) (*http.Response, error) {

	req, err := requestInfo.ToRequest()
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func TestNewTableItemRequestBuilder(t *testing.T) {
	client := &MockClient{}

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1", "sysId": "sysid"}

	req := NewTableItemRequestBuilder(client, pathParameters)

	assert.NotNil(t, req)
}

func TestTableItemRequestBuilder_Get(t *testing.T) {
	client := &MockClient{}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON
		responseJSON := `{
			"result": {
			  "parent": "",
			  "made_sla": "true",
			  "watch_list": "",
			  "upon_reject": "cancel",
			  "sys_updated_on": "2016-01-19 04:52:04",
			  "approval_history": "",
			  "number": "PRB0000050",
			  "sys_updated_by": "glide.maint",
			  "opened_by": {
				"link": "https://instance.servicenow.com/api/now/table/sys_user/glide.maint",
				"value": "glide.maint"
			  },
			  "user_input": "",
			  "sys_created_on": "2016-01-19 04:51:19",
			  "sys_domain": {
				"link": "https://instance.servicenow.com/api/now/table/sys_user_group/global",
				"value": "global"
			  },
			  "state": "4",
			  "sys_created_by": "glide.maint",
			  "knowledge": "false",
			  "order": "",
			  "closed_at": "2016-01-19 04:52:04",
			  "cmdb_ci": {
				"link": "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				"value": "55b35562c0a8010e01cff22378e0aea9"
			  },
			  "delivery_plan": "",
			  "impact": "3",
			  "active": "false",
			  "work_notes_list": "",
			  "business_service": "",
			  "priority": "4",
			  "sys_domain_path": "/",
			  "time_worked": "",
			  "expected_start": "",
			  "rejection_goto": "",
			  "opened_at": "2016-01-19 04:49:47",
			  "business_duration": "1970-01-01 00:00:00",
			  "group_list": "",
			  "work_end": "",
			  "approval_set": "",
			  "wf_activity": "",
			  "work_notes": "",
			  "short_description": "Switch occasionally drops connections",
			  "correlation_display": "",
			  "delivery_task": "",
			  "work_start": "",
			  "assignment_group": "",
			  "additional_assignee_list": "",
			  "description": "Switch occasionally drops connections",
			  "calendar_duration": "1970-01-01 00:02:17",
			  "close_notes": "updated firmware",
			  "sys_class_name": "problem",
			  "closed_by": "",
			  "follow_up": "",
			  "sys_id": "04ce72c9c0a8016600b5b7f75ac67b5b",
			  "contact_type": "phone",
			  "urgency": "3",
			  "company": "",
			  "reassignment_count": "",
			  "activity_due": "",
			  "assigned_to": "",
			  "comments": "",
			  "approval": "not requested",
			  "sla_due": "",
			  "comments_and_work_notes": "",
			  "due_date": "",
			  "sys_mod_count": "1",
			  "sys_tags": "",
			  "escalation": "0",
			  "upon_approval": "proceed",
			  "correlation_id": "",
			  "location": ""
			}
		  }`

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseJSON))
	}))

	parsedUrl, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	pathParameters := map[string]string{"baseurl": "http://" + parsedUrl.Host, "table": parsedUrl.Path, "sysId": "sysid"}

	req := NewTableItemRequestBuilder(client, pathParameters)

	params := &TableItemRequestBuilderGetQueryParameters{
		DisplayValue:         "true",
		ExcludeReferenceLink: true,
		Fields:               []string{"field1", "field2"},
		QueryNoDomain:        true,
		View:                 "desktop",
	}

	response, err := req.Get(params)

	fmt.Println(response)

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestTableItemRequestBuilder_Delete(t *testing.T) {
	client := &MockClient{}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON
		responseJSON := ``

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseJSON))
	}))

	parsedUrl, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	pathParameters := map[string]string{"baseurl": "http://" + parsedUrl.Host, "table": parsedUrl.Path, "sysId": "sysid"}
	req := NewTableItemRequestBuilder(client, pathParameters)

	params := &TableItemRequestBuilderDeleteQueryParameters{
		QueryNoDomain: true,
	}

	err = req.Delete(params)

	assert.Nil(t, err)
}
