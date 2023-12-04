package tableapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

const fakeLinkKey = "https://fake-link.com"

// Mock client for testing
type mockClient struct{}

func (c *mockClient) Send(requestInformation core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {

	url, err := requestInformation.Url()
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: %s", err)
	}

	if url == fakeLinkKey {

		rawJson := map[string]interface{}{
			"result": []map[string]interface{}{
				{
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
			},
		}

		jsonData, _ := json.Marshal(rawJson)

		// Create a http.Response with the JSON data as the body
		resp := &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     http.Header{},
			Body:       io.NopCloser(strings.NewReader(string(jsonData))), // We will replace this later
			Request:    nil,                                               // We don't need this for the example
		}

		return resp, nil
	}

	return nil, nil
}

func TestNewPageIteratorWithClient(t *testing.T) {
	// Mock client and current page
	client := &mockClient{}
	currentPage := TableCollectionResponse{
		// Initialize with test data
	}

	pageIterator, err := NewPageIterator(currentPage, client)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if pageIterator == nil {
		t.Error("Expected PageIterator, but got nil")
	}
}

func TestIterateWithNoCallback(t *testing.T) {
	// Mock PageIterator
	pageIterator := &PageIterator{
		currentPage: PageResult{
			// Initialize with test data
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	err := pageIterator.Iterate(nil)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestNewPageIteratorIterateSinglePageWithCallback(t *testing.T) {
	// Mock PageIterator
	pageIterator := &PageIterator{
		currentPage: PageResult{
			// Initialize with test data
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	// Mock callback function
	callback := func(pageItem *TableEntry) bool {
		// Implement your callback logic for testing
		return true
	}

	err := pageIterator.Iterate(callback)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
