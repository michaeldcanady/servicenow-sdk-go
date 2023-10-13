package tableapi

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTableRequestBuilder(t *testing.T) {

	client := MockClient{}

	pathParameters := map[string]string{"baseurl": "https://instance.service-now.com/api/now", "table": "table1"}

	req := NewTableItemRequestBuilder(&client, pathParameters)

	assert.NotNil(t, req)
}

func TestTableUrl(t *testing.T) {

	client := MockClient{}

	pathParameters := map[string]string{"baseurl": "https://instance.service-now.com/api/now", "table": "table1"}

	req := NewTableItemRequestBuilder(&client, pathParameters)

	assert.Equal(t, req.PathParameters, pathParameters)

	if !reflect.DeepEqual(req.PathParameters, pathParameters) {
		t.Errorf("excepted: %s, got: %s", pathParameters, req.PathParameters)
	}
}

func TestTableRequestBuilder_Get(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON
		responseJSON := `{
		  "result": [
		    {
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
		  ]
		}`

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseJSON))
	}))

	client := &MockClient{}

	parsedUrl, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	pathParameters := map[string]string{"baseurl": "http://" + parsedUrl.Host, "table": parsedUrl.Path}

	builder := NewTableRequestBuilder(client, pathParameters)

	// Call the Get method
	resp, err := builder.Get(nil)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	// You can further validate the response if needed
	if resp == nil {
		t.Error("Expected a non-nil response, but got nil")
	}

	expectedType := reflect.TypeOf(&TableCollectionResponse{})
	if reflect.TypeOf(resp) != expectedType {
		t.Errorf("Expected response of type %v, but got type %v", expectedType, reflect.TypeOf(resp))
	}

	if len(resp.Result) != 1 {
		t.Errorf("Expected response with 1 result, but got %v", len(resp.Result))
	}
}

func TestTableRequestBuilder_Count(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON
		responseJSON := ``

		w.WriteHeader(http.StatusOK)
		w.Header().Add("X-Total-Count", "1")
		_, _ = w.Write([]byte(responseJSON))
	}))
	defer mockServer.Close()

	client := &MockClient{}

	// Create an instance of TableRequestBuilder using the mock server URL

	parsedUrl, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	pathParameters := map[string]string{"baseurl": "http://" + parsedUrl.Host, "table": parsedUrl.Path}

	builder := NewTableRequestBuilder(client, pathParameters)

	// Call the Get method
	count, err := builder.Count()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	// You can further validate the response if needed
	if count == -1 {
		t.Error("Expected a non-nil response, but got nil")
	}

	expectedType := reflect.Int

	if reflect.TypeOf(count).Kind() != expectedType {
		t.Errorf("Expected response of type %v, but got type %v", expectedType, reflect.TypeOf(count))
	}
}
