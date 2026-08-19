package support

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/jarcoal/httpmock"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/tests/integration/v2/mockdata"
)

// RegisterAllMocks registers httpmock responders for all API surfaces.
// Called once in TestMain.
func RegisterAllMocks() {
	instance := IntegrationInstance()
	registerTableMocks(instance)
	registerAttachmentMocks(instance)
	registerBatchMocks(instance)
	registerAccountMocks(instance)
	registerCaseMocks(instance)
	registerCmdbMocks(instance)
	registerAggregationMocks(instance)
	registerActivitySubscriptionsMocks(instance)
	registerAppointmentBookingMocks(instance)
	registerAppServiceMocks(instance)
	registerOAuth2Mocks(instance)
	registerDocumentsMocks(instance)

	httpmock.RegisterNoResponder(func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("no responder found for %s %s", req.Method, req.URL)
	})
}

// ── Table ────────────────────────────────────────────────────────────────

func registerTableMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/now/v1/table", instance)

	httpmock.RegisterResponder("GET", base+"/incident",
		func(req *http.Request) (*http.Response, error) {
			query := req.URL.Query().Get("sysparm_query")
			data := mockdata.IncidentList
			if strings.Contains(query, "ORDERBYDESC") {
				data = mockdata.IncidentListSortedDesc
			}
			resp := httpmock.NewStringResponse(200, data)
			resp.Header.Set("Content-Type", "application/json")
			resp.Header.Set("Link", fmt.Sprintf(`<%s/incident?sysparm_limit=2&sysparm_offset=2>; rel="next"`, base))
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/incident",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mockdata.IncidentCreated)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(base+`/incident/[a-zA-Z0-9_]+$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.IncidentItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("PUT", regexp.MustCompile(base+`/incident/[a-zA-Z0-9_]+$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.IncidentUpdated)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("PATCH", regexp.MustCompile(base+`/incident/[a-zA-Z0-9_]+$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.IncidentPatched)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(base+`/incident/[a-zA-Z0-9_]+$`),
		httpmock.NewStringResponder(204, ""))

	httpmock.RegisterResponder("HEAD", base+"/incident",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "")
			resp.Header.Set("Content-Type", "application/json")
			resp.Header.Set("X-Total-Count", "1")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", base+"/incident/00000000000000000000000000000000",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(404, mockdata.SomeErrorJSON)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("DELETE", base+"/incident/00000000000000000000000000000000",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(404, mockdata.SomeErrorJSON)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── Attachment ───────────────────────────────────────────────────────────

func registerAttachmentMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/now/v1/attachment", instance)

	httpmock.RegisterResponder("GET", base,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.AttachmentList)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(base+`/(?:[a-zA-Z0-9_]+)?$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.AttachmentItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mockdata.AttachmentItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(base+`/[a-zA-Z0-9_]+$`),
		httpmock.NewStringResponder(204, ""))

	httpmock.RegisterResponder("DELETE", base+"/00000000000000000000000000000000",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(404, mockdata.SomeErrorJSON)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/file",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mockdata.AttachmentCreated)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(base+`/[a-zA-Z0-9_]+/file$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "mock file content")
			resp.Header.Set("Content-Type", "text/plain")
			return resp, nil
		})
}

// ── Batch ────────────────────────────────────────────────────────────────

func registerBatchMocks(instance string) {
	batchURL := fmt.Sprintf("https://%s.service-now.com/api/now/v1/batch", instance)
	tableBase := fmt.Sprintf("https://%s.service-now.com/api/now/v1/table", instance)

	httpmock.RegisterResponder("POST", batchURL,
		func(req *http.Request) (*http.Response, error) {
			var resp *http.Response
			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil || len(bodyBytes) == 0 {
				resp = httpmock.NewStringResponse(200, mockdata.BatchResponse)
			} else {
				bodyStr := string(bodyBytes)
				if strings.Contains(bodyStr, `"method":"INVALID"`) {
					resp = httpmock.NewStringResponse(400, `{"error":{"message":"Invalid operation in batch"},"status":"failure"}`)
				} else if strings.Contains(bodyStr, `"rest_requests":[]`) {
					resp = httpmock.NewStringResponse(400, `{"error":{"message":"Invalid batch request body"},"status":"failure"}`)
				} else if len(bodyBytes) > 2 {
					resp = httpmock.NewStringResponse(200, mockdata.BatchResponse)
				} else {
					resp = httpmock.NewStringResponse(400, `{"error":{"message":"Invalid batch request body"},"status":"failure"}`)
				}
			}
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", tableBase+"/incident",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mockdata.IncidentCreated)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", tableBase+"/incident/mock_sys_id_1",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.IncidentItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── Account ──────────────────────────────────────────────────────────────

func registerAccountMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/now/v1/account", instance)

	httpmock.RegisterResponder("GET", base,
		func(req *http.Request) (*http.Response, error) {
			query := req.URL.Query().Get("sysparm_query")
			data := mockdata.AccountList
			if query == "nameLIKEBDD" {
				data = mockdata.AccountListSingle
			} else if query == "sys_id=00000000000000000000000000000000" {
				data = mockdata.AccountListEmpty
			} else if strings.Contains(query, "sysparm_limit=1") {
				data = mockdata.AccountListSingle
			}
			resp := httpmock.NewStringResponse(200, data)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mockdata.AccountCreated)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", base+"/mock_account_sys_id_1",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.AccountItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("PATCH", regexp.MustCompile(base+`/[a-zA-Z0-9_]+$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.AccountUpdated)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(base+`/[a-zA-Z0-9_]+$`),
		httpmock.NewStringResponder(204, ""))
}

// ── Case ─────────────────────────────────────────────────────────────────

func registerCaseMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/sn_customerservice/v1/case", instance)

	httpmock.RegisterResponder("GET", base,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.CaseList)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mockdata.CaseCreated)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(base+`/[a-zA-Z0-9_]+$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.CaseItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("PUT", regexp.MustCompile(base+`/[a-zA-Z0-9_]+$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.CaseUpdated)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(base+`/[a-zA-Z0-9_]+$`),
		httpmock.NewStringResponder(204, ""))

	httpmock.RegisterResponder("GET", base+"/mock_case_sys_id_1/activities",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.CaseActivities)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", base+"/mock_case_sys_id_1/field_values/state",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.CaseFieldValues)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── CMDB ─────────────────────────────────────────────────────────────────

func registerCmdbMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/now/v1/cmdb/instance", instance)

	httpmock.RegisterResponder("GET", base+"/cmdb_ci",
		func(req *http.Request) (*http.Response, error) {
			query := req.URL.Query().Get("sysparm_query")
			data := mockdata.CmdbList
			if strings.Contains(query, "sysparm_limit=1") {
				data = mockdata.CmdbListSingle
			} else if query == "sys_id=00000000000000000000000000000000" {
				data = mockdata.CmdbListEmpty
			}
			resp := httpmock.NewStringResponse(200, data)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", base+"/cmdb_ci/mock_cmdb_sys_id_1",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.CmdbItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/cmdb_ci",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mockdata.CmdbItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("PATCH", regexp.MustCompile(base+`/cmdb_ci/[a-zA-Z0-9_]+$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.CmdbItem)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("DELETE", regexp.MustCompile(base+`/cmdb_ci/[a-zA-Z0-9_]+$`),
		httpmock.NewStringResponder(204, ""))

	httpmock.RegisterResponder("GET", base+"/this_cmdb_class_does_not_exist",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, mockdata.CmdbInvalidClass)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── Aggregation ──────────────────────────────────────────────────────────

func registerAggregationMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/now/stats/incident", instance)

	httpmock.RegisterResponder("GET", base,
		func(req *http.Request) (*http.Response, error) {
			sumParam := req.URL.Query().Get("sysparm_sum")
			if sumParam != "" {
				resp := httpmock.NewStringResponse(200, mockdata.StatsSum)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			}
			groupByParam := req.URL.Query().Get("sysparm_group_by")
			if groupByParam != "" {
				resp := httpmock.NewStringResponse(200, mockdata.StatsGroupBy)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			}
			count := req.URL.Query().Get("sysparm_count")
			if count == "" {
				resp := httpmock.NewStringResponse(400, mockdata.StatsNoAggregateParams)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			}
			resp := httpmock.NewStringResponse(200, mockdata.StatsCount)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://%s.service-now.com/api/now/stats/this_table_does_not_exist", instance),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, mockdata.StatsInvalidTable)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── Activity Subscriptions ───────────────────────────────────────────────

func registerActivitySubscriptionsMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/now/v1/actsub", instance)

	httpmock.RegisterResponder("GET", base+"/activities",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.ActivitySubscriptionItemResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", base+"/facets/mock_context_id/mock_instance_id",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.FacetsInstanceResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── Appointment Booking ──────────────────────────────────────────────────

func registerAppointmentBookingMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/sn_apptmnt_booking/v1/appointment", instance)

	httpmock.RegisterResponder("GET", base+"/configuration",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.AppointmentBookingConfigurationResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", base+"/calendar",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.AppointmentBookingCalendarResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/appointment",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.AppointmentResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/availability",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.AvailabilityResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/execute_rule_conditions",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.ExecuteRuleConditionsResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/userwindow",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.UserWindowResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── App Service ──────────────────────────────────────────────────────────

func registerAppServiceMocks(instance string) {
	appBase := fmt.Sprintf("https://%s.service-now.com/api/now/cmdb/app_service", instance)
	csdmBase := fmt.Sprintf("https://%s.service-now.com/api/now/cmdb/csdm/app_service", instance)

	httpmock.RegisterResponder("POST", appBase+"/create",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.CreateServiceResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", appBase+"/mock_app_service_sys_id_1/getContent",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.GetContentResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("GET", csdmBase+"/find_service",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.FindServiceResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── OAuth2 ───────────────────────────────────────────────────────────────

func registerOAuth2Mocks(instance string) {
	tokenURL := fmt.Sprintf("https://%s.service-now.com/oauth_token.do", instance)

	httpmock.RegisterResponder("POST", tokenURL,
		func(req *http.Request) (*http.Response, error) {
			if err := req.ParseForm(); err != nil {
				resp := httpmock.NewStringResponse(400, mockdata.TokenErrorResponse)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			}

			clientID := req.FormValue("client_id")
			clientSecret := req.FormValue("client_secret")
			grantType := req.FormValue("grant_type")

			if clientID == "invalid-client-id" || clientSecret == "invalid-secret" {
				resp := httpmock.NewStringResponse(401, mockdata.TokenErrorResponse)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			}

			switch grantType {
			case "client_credentials", "authorization_code":
				resp := httpmock.NewStringResponse(200, mockdata.TokenResponse)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			case "password":
				username := req.FormValue("username")
				password := req.FormValue("password")
				if username == "" || password == "" {
					resp := httpmock.NewStringResponse(400, mockdata.TokenErrorResponse)
					resp.Header.Set("Content-Type", "application/json")
					return resp, nil
				}
				resp := httpmock.NewStringResponse(200, mockdata.TokenResponse)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			case "refresh_token":
				resp := httpmock.NewStringResponse(200, mockdata.RefreshTokenResponse)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			case "urn:ietf:params:oauth:grant-type:jwt-bearer":
				resp := httpmock.NewStringResponse(200, mockdata.TokenResponse)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			default:
				resp := httpmock.NewStringResponse(400, mockdata.TokenErrorResponse)
				resp.Header.Set("Content-Type", "application/json")
				return resp, nil
			}
		})

	revocationURL := fmt.Sprintf("https://%s.service-now.com/oauth_revoke.do", instance)
	httpmock.RegisterResponder("POST", revocationURL,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mockdata.RevocationResponse)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}

// ── Documents ────────────────────────────────────────────────────────────

func registerDocumentsMocks(instance string) {
	base := fmt.Sprintf("https://%s.service-now.com/api/now/v1/documents", instance)

	httpmock.RegisterResponder("GET", base+"/explore",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"result": []}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/create",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, `{"result": {"sys_id": "mock_doc_sys_id_1"}}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/create_or_link",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, `{"result": {"sys_id": "mock_doc_sys_id_1"}}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("DELETE", base+"/delete",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(204, "")
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(base+`/versions/[a-zA-Z0-9_]+$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{"result": []}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(base+`/[a-zA-Z0-9_]+/content$`),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "mock document content")
			resp.Header.Set("Content-Type", "text/plain")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/sync_down",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, `{"error":{"message":"Invalid document"},"status":"failure"}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/attach",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, `{"error":{"message":"Missing provider"},"status":"failure"}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

	httpmock.RegisterResponder("POST", base+"/action",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(400, `{"error":{"message":"Invalid action"},"status":"failure"}`)
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})
}
