//go:build integration

package tests

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/jarcoal/httpmock"
	nethttplibrary "github.com/microsoft/kiota-http-go"
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
			query := req.URL.Query().Get("sysparm_query")
			data := mockIncidentList
			if strings.Contains(query, "ORDERBYDESC") {
				data = mockIncidentListSortedDesc
			}
			resp := httpmock.NewStringResponse(200, data)
			resp.Header.Set("Content-Type", "application/json")
			resp.Header.Set("Link", fmt.Sprintf("<%s/incident?sysparm_limit=2&sysparm_offset=2>; rel=\"next\"", tableBaseURL))
			return resp, nil
		})
	httpmock.RegisterResponder("POST", tableBaseURL+"/incident",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mockCreatedIncident)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	// Exact match for the default mock incident
	httpmock.RegisterResponder("GET", tableBaseURL+"/incident/mock_sys_id_1",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockIncidentItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("PUT", regexp.MustCompile(tableBaseURL+`/incident/[a-zA-Z0-9_]+$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockUpdatedIncident)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("PATCH", regexp.MustCompile(tableBaseURL+`/incident/[a-zA-Z0-9_]+$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockPatchedIncident)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(tableBaseURL+`/incident/[a-zA-Z0-9_]+$`), httpmock.NewStringResponder(204, ""))

	// Attachment API Mocks
	attachBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/attachment", instance)
	httpmock.RegisterResponder("GET", attachBaseURL, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockAttachmentList)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	attachIdRegex := regexp.MustCompile(attachBaseURL + `(?:/([a-zA-Z0-9_]+))?$`)
	httpmock.RegisterRegexpResponder("GET", attachIdRegex, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockAttachmentItem)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(attachBaseURL+`/+file`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(201, mockAttachmentItem)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	fileRegex := regexp.MustCompile(attachBaseURL + `(?:/([a-zA-Z0-9_]+))?/file`)
	httpmock.RegisterRegexpResponder("GET", fileRegex, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, "test content")
		resp.Header.Set("x-attachment-metadata", `{"sys_id":"mock_attach_id_1","file_name":"test.txt"}`)
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("DELETE", attachIdRegex, httpmock.NewStringResponder(204, ""))

	// Batch API Mocks
	batchBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1", instance)
	httpmock.RegisterResponder("POST", batchBaseURL+"/batch", func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockBatchMultiResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterResponder("POST", batchBaseURL+"/table/incident", func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(201, mockCreatedIncident)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(batchBaseURL+`/+table/+incident/+(?:[a-zA-Z0-9_]+)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockIncidentItem)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
}

func getHttpClient() *http.Client {
	if isOffline() {
		// Wrap httpmock transport with Kiota middleware to ensure HeadersInspectionOptions are populated
		transport := nethttplibrary.NewCustomTransportWithParentTransport(httpmock.DefaultTransport, nethttplibrary.GetDefaultMiddlewares()...)
		return &http.Client{
			Transport: transport,
		}
	}
	return nil
}
