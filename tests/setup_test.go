//go:build integration

package tests

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/jarcoal/httpmock"
)

func isOffline() bool {
	return os.Getenv("SN_OFFLINE") == "true"
}

func setupGlobalMocks() {
	if !isOffline() {
		return
	}

	httpmock.Activate()

	httpmock.RegisterNoResponder(func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("no responder found for %s %s", req.Method, req.URL)
	})

	instance := os.Getenv("SN_INSTANCE")
	if instance == "" {
		instance = "mock_instance"
		_ = os.Setenv("SN_INSTANCE", instance)
	}

	// Table API Mocks
	tableBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/table", instance)
	httpmock.RegisterResponder("GET", tableBaseURL+"/incident",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockIncidentList)
			resp.Header.Set("Link", fmt.Sprintf("<%s/incident?sysparm_limit=2&sysparm_offset=2>; rel=\"next\"", tableBaseURL))
			return resp, nil
		})
	httpmock.RegisterResponder("POST", tableBaseURL+"/incident",
		httpmock.NewStringResponder(201, mockCreatedIncident))

	tableIdRegex := regexp.MustCompile(tableBaseURL + `/+incident(?:/+(?:[a-zA-Z0-9_]+)?)?$`)
	httpmock.RegisterRegexpResponder("GET", tableIdRegex, httpmock.NewStringResponder(200, mockIncidentItem))
	httpmock.RegisterRegexpResponder("PUT", tableIdRegex, httpmock.NewStringResponder(200, mockUpdatedIncident))
	httpmock.RegisterRegexpResponder("PATCH", tableIdRegex, httpmock.NewStringResponder(200, mockPatchedIncident))
	httpmock.RegisterRegexpResponder("DELETE", tableIdRegex, httpmock.NewStringResponder(204, ""))

	// Attachment API Mocks
	attachBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/attachment", instance)
	httpmock.RegisterResponder("GET", attachBaseURL, httpmock.NewStringResponder(200, mockAttachmentList))

	attachIdRegex := regexp.MustCompile(attachBaseURL + `/?([a-zA-Z0-9_]+)?$`)
	httpmock.RegisterRegexpResponder("GET", attachIdRegex, httpmock.NewStringResponder(200, mockAttachmentItem))
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(attachBaseURL+`/+file`), httpmock.NewStringResponder(201, mockAttachmentItem))

	fileRegex := regexp.MustCompile(attachBaseURL + `/?([a-zA-Z0-9_]+)?/file`)
	httpmock.RegisterRegexpResponder("GET", fileRegex, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, "test content")
		resp.Header.Set("x-attachment-metadata", `{"sys_id":"mock_attach_id_1","file_name":"test.txt"}`)
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("DELETE", attachIdRegex, httpmock.NewStringResponder(204, ""))

	// Batch API Mocks
	batchBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1", instance)
	httpmock.RegisterResponder("POST", batchBaseURL+"/batch", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(200, mockBatchMultiResponse), nil
	})
	httpmock.RegisterResponder("POST", batchBaseURL+"/table/incident", httpmock.NewStringResponder(201, mockCreatedIncident))
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(batchBaseURL+`/+table/+incident/+(?:[a-zA-Z0-9_]+)?$`), httpmock.NewStringResponder(200, mockIncidentItem))
}

func getHttpClient() *http.Client {
	if isOffline() {
		return &http.Client{
			Transport: httpmock.DefaultTransport,
		}
	}
	return http.DefaultClient
}
