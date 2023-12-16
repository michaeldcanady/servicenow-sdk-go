package tableapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

const (
	fakeLinkKey          = "https://fake-link.com"
	fakeLinkWithLinks    = "https://fake-link1.com"
	fakeLinkWithLinksErr = "https://fake-link2.com"
	fakeLinkStatusFailed = "https://fake-link3.com"
	fakeLinkNilResponse  = "https://fake-link4.com"

	fakeNextLink  = "https://fake-link.com?next"
	fakePrevLink  = "https://fake-link.com?prev"
	fakeFirstLink = "https://fake-link.com?first"
	fakeLastLink  = "https://fake-link.com?last"
)

var (
	fakeResult = map[string]interface{}{
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

	fakeEntry TableEntry = fakeResult

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
)

func getFakeJSON() []byte {
	rawJSON := map[string]interface{}{
		"result": []map[string]interface{}{
			fakeResult,
		},
	}

	jsonData, _ := json.Marshal(rawJSON)

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
	}

	return nil, nil
}

func TestNewPageIteratorWithClient(t *testing.T) {
	// Mock client and current page
	client := &mockClient{}
	currentPage := TableCollectionResponse{
		Result:           []*TableEntry{&fakeEntry},
		FirstPageLink:    fakeFirstLink,
		LastPageLink:     fakeLastLink,
		NextPageLink:     fakeNextLink,
		PreviousPageLink: fakePrevLink,
	}

	pageIterator, err := NewPageIterator(currentPage, client)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if pageIterator == nil {
		t.Error("Expected PageIterator, but got nil")
	}

	assert.Equal(t, &expectedIterator, pageIterator)
}

func TestNewPageIteratorWithoutClient(t *testing.T) {
	// Mock current page
	currentPage := TableCollectionResponse{
		// Initialize with test data
	}

	pageIterator, err := NewPageIterator(currentPage, nil)

	assert.Equal(t, (*PageIterator)(nil), pageIterator)
	assert.Equal(t, ErrNilClient, err)
}

func TestNewPageIteratorNilCurrentPageWithClient(t *testing.T) {
	client := &mockClient{}
	pageIterator, err := NewPageIterator(nil, client)

	assert.Equal(t, (*PageIterator)(nil), pageIterator)
	assert.Equal(t, ErrNilResponse, err)
}

func TestPageIteratorNextWithLinkNoError(t *testing.T) {
	currentPage := TableCollectionResponse{
		NextPageLink: fakeLinkWithLinks,
	}

	client := &mockClient{}

	pageIterator, err := NewPageIterator(currentPage, client)
	assert.Nil(t, err)

	page, err := pageIterator.next()
	assert.Nil(t, err)

	assert.Equal(t, expectedResult, page)
}

func TestPageIteratorEnumerateAll(t *testing.T) {
	pageIterator := PageIterator{
		currentPage: expectedResult,
	}

	enumCount := 0

	keepIterating := pageIterator.enumerate(func(item *TableEntry) bool {
		index := pageIterator.pauseIndex

		result := expectedResult.Result[index]

		assert.Equal(t, result, item)

		enumCount += 1

		return true
	})
	assert.Equal(t, true, keepIterating)
	assert.Equal(t, len(expectedResult.Result), enumCount)
}

func TestPageIteratorEnumerateOnce(t *testing.T) {
	pageIterator := PageIterator{
		currentPage: expectedResult,
	}

	enumCount := 0

	keepIterating := pageIterator.enumerate(func(item *TableEntry) bool {
		index := pageIterator.pauseIndex

		result := expectedResult.Result[index]

		assert.Equal(t, result, item)

		enumCount += 1

		return false
	})
	assert.Equal(t, false, keepIterating)
	assert.Equal(t, 1, enumCount)
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
	assert.ErrorIs(t, err, ErrNilCallback)
}

func TestPageIteratorIterateSinglePageWithCallback(t *testing.T) {
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
	assert.Nil(t, err)
}

func TestPageIteratorIterateMultiplePagesWithCallback(t *testing.T) {
	// Mock PageIterator
	pageIterator := &PageIterator{
		currentPage: PageResult{
			NextPageLink: fakeLinkWithLinks,
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
	assert.Nil(t, err)
}

func TestPageIteratorLast(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			LastPageLink: fakeNextLink,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	_, err := pageIterator.Last()

	assert.Error(t, err)
}

func TestPageFetchPageSendErr(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			LastPageLink: fakeLinkStatusFailed,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	_, err := pageIterator.fetchPage(fakeLinkStatusFailed)
	assert.Error(t, err)
}

func TestPageFetchPageEmptyUri(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			LastPageLink: fakeLinkStatusFailed,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	_, err := pageIterator.fetchPage("")
	assert.ErrorIs(t, err, ErrEmptyURI)
}

func TestPageIteratorFetchAndConvertPageWithLinkErrNilResponseBody(t *testing.T) {
	currentPage := TableCollectionResponse{
		NextPageLink: fakeLinkWithLinksErr,
	}

	client := &mockClient{}

	pageIterator, err := NewPageIterator(currentPage, client)
	assert.Nil(t, err)

	page, err := pageIterator.fetchAndConvertPage(pageIterator.currentPage.NextPageLink)
	assert.ErrorIs(t, err, core.ErrNilResponseBody)

	assert.Equal(t, PageResult{}, page)
}

func TestPageIteratorFetchAndConvertPageWithoutLink(t *testing.T) {
	currentPage := TableCollectionResponse{
		NextPageLink: fakeLinkNilResponse,
	}

	client := &mockClient{}

	pageIterator, err := NewPageIterator(currentPage, client)
	assert.Nil(t, err)

	page, err := pageIterator.fetchAndConvertPage(pageIterator.currentPage.NextPageLink)
	assert.ErrorIs(t, err, core.ErrNilResponse)

	assert.Equal(t, PageResult{}, page)
}
