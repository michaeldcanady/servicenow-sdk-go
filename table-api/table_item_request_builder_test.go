package tableapi

import (
	"encoding/json"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (c *MockClient) Send(requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {

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
		_, _ = w.Write([]byte(responseJSON)) //nolint:errcheck
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

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestTableItemRequestBuilder_Delete(t *testing.T) {
	client := &MockClient{}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON
		responseJSON := ``

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseJSON)) //nolint:errcheck
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

func TestTableItemRequestBuilder_Put(t *testing.T) {
	// Create a mock client for testing
	client := &MockClient{}

	// Create a mock server for testing
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		resp := map[string]interface{}{
			"result": map[string]interface{}{
				"parent":           "",
				"made_sla":         "true",
				"watch_list":       "",
				"upon_reject":      "cancel",
				"sys_updated_on":   "2016-01-19 04:52:04",
				"approval_history": "",
				"number":           "PRB0000050",
				"sys_updated_by":   "glide.maint",
				"opened_by": map[string]interface{}{
					"link":  "https://instance.servicenow.com/api/now/table/sys_user/glide.maint",
					"value": "glide.maint",
				},
				"user_input":     "",
				"sys_created_on": "2016-01-19 04:51:19",
				"sys_domain": map[string]interface{}{
					"link":  "https://instance.servicenow.com/api/now/table/sys_user_group/global",
					"value": "global",
				},
				"state":          "4",
				"sys_created_by": "glide.maint",
				"knowledge":      "false",
				"order":          "",
				"closed_at":      "2016-01-19 04:52:04",
				"cmdb_ci": map[string]interface{}{
					"link":  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
					"value": "55b35562c0a8010e01cff22378e0aea9",
				},
				"delivery_plan":            "",
				"impact":                   "3",
				"active":                   "false",
				"work_notes_list":          "",
				"business_service":         "",
				"priority":                 "4",
				"sys_domain_path":          "/",
				"time_worked":              "",
				"expected_start":           "",
				"rejection_goto":           "",
				"opened_at":                "2016-01-19 04:49:47",
				"business_duration":        "1970-01-01 00:00:00",
				"group_list":               "",
				"work_end":                 "",
				"approval_set":             "",
				"wf_activity":              "",
				"work_notes":               "",
				"short_description":        "Switch occasionally drops connections",
				"correlation_display":      "",
				"delivery_task":            "",
				"work_start":               "",
				"assignment_group":         "",
				"additional_assignee_list": "",
				"description":              "Switch occasionally drops connections",
				"calendar_duration":        "1970-01-01 00:02:17",
				"close_notes":              "updated firmware",
				"sys_class_name":           "problem",
				"closed_by":                "",
				"follow_up":                "",
				"sys_id":                   "04ce72c9c0a8016600b5b7f75ac67b5b",
				"contact_type":             "phone",
				"urgency":                  "3",
				"company":                  "",
				"reassignment_count":       "",
				"activity_due":             "",
				"assigned_to":              "",
				"comments":                 "",
				"approval":                 "not requested",
				"sla_due":                  "",
				"comments_and_work_notes":  "",
				"due_date":                 "",
				"sys_mod_count":            "1",
				"sys_tags":                 "",
				"escalation":               "0",
				"upon_approval":            "proceed",
				"correlation_id":           "",
				"location":                 "",
			},
		}

		resp = maps.Clone[map[string]interface{}](resp)

		switch r.Method {
		case http.MethodGet:
			break
		case http.MethodPut:
			body := make(map[string]interface{})

			// Read the request body into a []byte
			requestBody, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = json.Unmarshal(requestBody, &body) // Use &body to correctly update the 'body' map
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			result := resp["result"].(map[string]interface{})

			for key, value := range body {
				result[key] = value
			}

			resp["result"] = result // Update the 'result' in 'resp'

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		jsonData, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(jsonData)) //nolint:errcheck
	}))

	// Parse the URL of the mock server
	parsedURL, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	// Prepare path parameters
	pathParameters := map[string]string{
		"baseurl": "http://" + parsedURL.Host,
		"table":   parsedURL.Path,
		"sysId":   "sysid",
	}

	// Create a request builder
	req := NewTableItemRequestBuilder(client, pathParameters)

	// Prepare values to update the record
	values := map[string]string{
		// Provide values to change in the record here
		"location": "home",
	}

	// Send the PUT request to update the record
	updatedRecord, err := req.Put(values, nil)

	// Perform assertions and test the response
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	expected := &TableItemResponse{
		Result: &TableEntry{
			"parent":           "",
			"made_sla":         "true",
			"watch_list":       "",
			"upon_reject":      "cancel",
			"sys_updated_on":   "2016-01-19 04:52:04",
			"approval_history": "",
			"number":           "PRB0000050",
			"sys_updated_by":   "glide.maint",
			"opened_by": map[string]interface{}{
				"link":  "https://instance.servicenow.com/api/now/table/sys_user/glide.maint",
				"value": "glide.maint",
			},
			"user_input":     "",
			"sys_created_on": "2016-01-19 04:51:19",
			"sys_domain": map[string]interface{}{
				"link":  "https://instance.servicenow.com/api/now/table/sys_user_group/global",
				"value": "global",
			},
			"state":          "4",
			"sys_created_by": "glide.maint",
			"knowledge":      "false",
			"order":          "",
			"closed_at":      "2016-01-19 04:52:04",
			"cmdb_ci": map[string]interface{}{
				"link":  "https://instance.servicenow.com/api/now/table/cmdb_ci/55b35562c0a8010e01cff22378e0aea9",
				"value": "55b35562c0a8010e01cff22378e0aea9",
			},
			"delivery_plan":            "",
			"impact":                   "3",
			"active":                   "false",
			"work_notes_list":          "",
			"business_service":         "",
			"priority":                 "4",
			"sys_domain_path":          "/",
			"time_worked":              "",
			"expected_start":           "",
			"rejection_goto":           "",
			"opened_at":                "2016-01-19 04:49:47",
			"business_duration":        "1970-01-01 00:00:00",
			"group_list":               "",
			"work_end":                 "",
			"approval_set":             "",
			"wf_activity":              "",
			"work_notes":               "",
			"short_description":        "Switch occasionally drops connections",
			"correlation_display":      "",
			"delivery_task":            "",
			"work_start":               "",
			"assignment_group":         "",
			"additional_assignee_list": "",
			"description":              "Switch occasionally drops connections",
			"calendar_duration":        "1970-01-01 00:02:17",
			"close_notes":              "updated firmware",
			"sys_class_name":           "problem",
			"closed_by":                "",
			"follow_up":                "",
			"sys_id":                   "04ce72c9c0a8016600b5b7f75ac67b5b",
			"contact_type":             "phone",
			"urgency":                  "3",
			"company":                  "",
			"reassignment_count":       "",
			"activity_due":             "",
			"assigned_to":              "",
			"comments":                 "",
			"approval":                 "not requested",
			"sla_due":                  "",
			"comments_and_work_notes":  "",
			"due_date":                 "",
			"sys_mod_count":            "1",
			"sys_tags":                 "",
			"escalation":               "0",
			"upon_approval":            "proceed",
			"correlation_id":           "",
			"location":                 "home",
		},
	}

	// You should assert that the updatedRecord matches your expected response here.
	// You may need to unmarshal the JSON response and compare specific fields.
	// For example: assert.Equal(t, expectedValue, updatedRecord.Field)

	assert.Equal(t, expected, updatedRecord)
}
