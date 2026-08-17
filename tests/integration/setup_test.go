//go:build integration

package integration

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/jarcoal/httpmock"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	attachmentMetadataHeader = "x-attachment-metadata"
)

func isOffline() bool {
	return os.Getenv("SN_OFFLINE") == "true"
}

func firstEnv(keys ...string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return ""
}

func integrationInstance() string {
	instance := firstEnv("SN_INSTANCE", "SNOW_INSTANCE")
	if instance == "" {
		instance = "mock_instance"
	}
	return instance
}

// newIntegrationAuthProvider picks client-credentials when SN_CLIENT_ID/SECRET (or SNOW_*)
// are set and the suite is not offline; otherwise basic auth (empty/mock values are fine
// under httpmock because responders never validate Authorization).
func newIntegrationAuthProvider(instance string) (authentication.AuthenticationProvider, error) {
	clientID := firstEnv("SN_CLIENT_ID", "SNOW_CLIENT_ID")
	clientSecret := firstEnv("SN_CLIENT_SECRET", "SNOW_CLIENT_SECRET")
	if !isOffline() && clientID != "" && clientSecret != "" {
		return credentials.NewClientCredentialsProvider(
			clientID,
			clientSecret,
			credentials.WithInstance(instance),
		)
	}

	username := firstEnv("SN_USERNAME", "SNOW_USERNAME")
	password := firstEnv("SN_PASSWORD", "SNOW_PASSWORD")
	if username == "" {
		username = "mock"
	}
	if password == "" {
		password = "mock"
	}
	return credentials.NewBasicProvider(username, password), nil
}

func setupGlobalMocks() {
	if !isOffline() {
		return
	}

	httpmock.Activate()

	httpmock.RegisterNoResponder(func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("no responder found for %s %s", req.Method, req.URL)
	})

	instance := integrationInstance()
	_ = os.Setenv("SN_INSTANCE", instance)

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
		resp.Header.Set(attachmentMetadataHeader, `{"sys_id":"mock_attach_id_1","file_name":"test.txt"}`)
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

	// Stats / Aggregation API Mocks
	statsBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/stats", instance)
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(statsBaseURL)+`/incident(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		query := req.URL.Query()
		// Live platform rejects requests with no aggregate operation parameters.
		if query.Get("sysparm_count") == "" &&
			query.Get("sysparm_sum_fields") == "" &&
			query.Get("sysparm_avg_fields") == "" &&
			query.Get("sysparm_min_fields") == "" &&
			query.Get("sysparm_max_fields") == "" {
			resp := httpmock.NewStringResponse(400, mockStatsNoAggregateParams)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}
		body := mockStatsCount
		if strings.Contains(query.Get("sysparm_query"), "sys_id=00000000000000000000000000000000") {
			body = mockStatsCountZero
		} else if query.Get("sysparm_sum_fields") != "" {
			body = mockStatsSum
		}
		resp := httpmock.NewStringResponse(200, body)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(statsBaseURL)+`/this_table_does_not_exist(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(400, mockStatsInvalidTable)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// Account API Mocks
	accountBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/account", instance)
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(accountBaseURL)+`(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		query := req.URL.Query()
		body := mockAccountList
		if strings.Contains(query.Get("sysparm_query"), "sys_id=00000000000000000000000000000000") {
			body = mockAccountListEmpty
		} else if query.Get("sysparm_limit") == "1" {
			body = mockAccountListSingle
		}
		resp := httpmock.NewStringResponse(200, body)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterResponder("GET", accountBaseURL+"/mock_account_sys_id_1", func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockAccountItem)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// CMDB Instance API Mocks
	cmdbBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/cmdb/instance", instance)
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(cmdbBaseURL)+`/cmdb_ci(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		query := req.URL.Query()
		body := mockCmdbList
		if strings.Contains(query.Get("sysparm_query"), "sys_id=00000000000000000000000000000000") {
			body = mockCmdbListEmpty
		} else if query.Get("sysparm_limit") == "1" {
			body = mockCmdbListSingle
		}
		resp := httpmock.NewStringResponse(200, body)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterResponder("GET", cmdbBaseURL+"/cmdb_ci/mock_cmdb_sys_id_1", func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockCmdbItem)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(cmdbBaseURL)+`/this_cmdb_class_does_not_exist(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(400, mockCmdbInvalidClass)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// Activity Subscriptions API Mocks
	actSubBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/actsub", instance)
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(actSubBaseURL)+`/activities(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		query := req.URL.Query()
		contextVal := query.Get("context")
		instanceVal := query.Get("context_instance")

		if contextVal == "" || instanceVal == "" {
			resp := httpmock.NewStringResponse(400, `{"error":{"message":"Invalid Query Parameters","detail":"context and context_instance are required"},"status":"failure"}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		}

		resp := httpmock.NewStringResponse(200, mockActivitySubscriptionItemResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(actSubBaseURL)+`/facets/[a-zA-Z0-9_]+/[a-zA-Z0-9_]+$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockFacetsInstanceResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// Appointment Booking API Mocks
	apptmntBookingBaseURL := fmt.Sprintf("https://%s.service-now.com/api/sn_apptmnt_booking/v1/appointment", instance)
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(apptmntBookingBaseURL)+`/configuration(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockAppointmentBookingConfigurationResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(apptmntBookingBaseURL)+`/calendar(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockAppointmentBookingCalendarResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(regexp.QuoteMeta(apptmntBookingBaseURL)+`/appointment(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockAppointmentResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(regexp.QuoteMeta(apptmntBookingBaseURL)+`/availability(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockAvailabilityResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(regexp.QuoteMeta(apptmntBookingBaseURL)+`/execute_rule_conditions(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockExecuteRuleConditionsResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("POST", regexp.MustCompile(regexp.QuoteMeta(apptmntBookingBaseURL)+`/userwindow(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockUserWindowResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// App Service API Mocks
	appServiceBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/cmdb/app_service", instance)
	httpmock.RegisterResponder("POST", appServiceBaseURL+"/create", func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockCreateServiceResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(appServiceBaseURL)+`/[a-zA-Z0-9_]+/getContent(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockGetContentResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	csdmAppServiceBaseURL := fmt.Sprintf("https://%s.service-now.com/api/now/cmdb/csdm/app_service", instance)
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(csdmAppServiceBaseURL)+`/find_service(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockFindServiceResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// Case API Mocks
	caseBaseURL := fmt.Sprintf("https://%s.service-now.com/api/sn_customerservice/v1/case", instance)

	// GET /case — list cases
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(caseBaseURL)+`(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockCaseList)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// POST /case — create a case
	httpmock.RegisterResponder("POST", caseBaseURL, func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(201, mockCreatedCase)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// GET /case/{id} — retrieve a case by ID
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(caseBaseURL)+`/[a-zA-Z0-9_]+(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockCaseItem)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// PUT /case/{id} — update a case by ID
	httpmock.RegisterRegexpResponder("PUT", regexp.MustCompile(regexp.QuoteMeta(caseBaseURL)+`/[a-zA-Z0-9_]+(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockUpdatedCase)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// GET /case/{id}/activities — list activities for a case
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(caseBaseURL)+`/[a-zA-Z0-9_]+/activities(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockCaseActivities)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// GET /case/field_values/{field_name} — retrieve field values for a case-level field
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(caseBaseURL)+`/field_values/[a-zA-Z0-9_]+(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockCaseFieldValues)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})

	// GET /case/{id}/field_values/{field_name} — retrieve field values for a field on a specific case
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(regexp.QuoteMeta(caseBaseURL)+`/[a-zA-Z0-9_]+/field_values/[a-zA-Z0-9_]+(?:\?.*)?$`), func(req *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, mockCaseFieldValues)
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
