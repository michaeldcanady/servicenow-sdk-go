package tableapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

const (
	fakeLinkKey          = "https://fake-link.com"
	fakeLinkWithLinks    = "https://fake-link1.com"
	fakeLinkWithLinksErr = "https://fake-link2.com"
	fakeLinkStatusFailed = "https://fake-link3.com"
	fakeLinkNilResponse  = "https://fake-link4.com"
	fakeLinkItemResult   = "https://fake-link5.com"

	fakeNextLink  = "https://fake-link.com?next"
	fakePrevLink  = "https://fake-link.com?prev"
	fakeFirstLink = "https://fake-link.com?first"
	fakeLastLink  = "https://fake-link.com?last"
)

var (
	fakeItemResult = map[string]interface{}{
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
	}

	fakeCollectionResponse = map[string]interface{}{
		"result": []map[string]interface{}{fakeItemResult},
	}

	fakeEntryResponse = map[string]interface{}{
		"result": fakeItemResult,
	}

	fakeEntry TableEntry = fakeItemResult

	expectedResult = PageResult{
		Result: []*TableEntry{
			&fakeEntry,
		},
		FirstPageLink:    fakeFirstLink,
		LastPageLink:     fakeLastLink,
		NextPageLink:     fakeNextLink,
		PreviousPageLink: fakePrevLink,
	}

	expectedIterator = PageIterator{
		currentPage: expectedResult,
		client:      &mockClient{},
		pauseIndex:  0,
	}

	fakeTableItemResponse = &TableItemResponse{
		Result: &fakeEntry,
	}
)

func getFakeJSON() []byte {
	jsonData, _ := json.Marshal(fakeCollectionResponse)

	return jsonData
}

// Mock client for testing
type mockClient struct{}

func (c *mockClient) Send(requestInformation core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	url, err := requestInformation.Url()
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: %s", err)
	}

	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{},
		Body:       http.NoBody,
		Request:    nil,
	}

	switch url {
	case fakeLinkStatusFailed:
		return nil, errors.New("net/http: nil Context")
	case fakeLinkWithLinks: // Adds headers
		header := http.Header{}

		header.Add("Link", "<"+fakeFirstLink+">;rel=\"first\"")
		header.Add("Link", "<"+fakeNextLink+">;rel=\"next\"")
		header.Add("Link", "<"+fakePrevLink+">;rel=\"prev\"")
		header.Add("Link", "<"+fakeLastLink+">;rel=\"last\"")
		resp.Header = header

		fallthrough
	case fakeLinkKey: // Adds body
		resp.Body = io.NopCloser(strings.NewReader(string(getFakeJSON())))

		fallthrough
	case fakeLinkWithLinksErr:
		return resp, nil
	case fakeLinkNilResponse:
		return nil, nil
	case fakeLinkItemResult:

		rawJSON, _ := json.Marshal(fakeEntryResponse)

		resp := &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     http.Header{},
			Body:       io.NopCloser(bytes.NewReader(rawJSON)),
			Request:    nil,
		}

		return resp, nil
	}

	return nil, nil
}
