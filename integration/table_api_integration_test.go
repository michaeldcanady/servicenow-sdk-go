package integration

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/require"
)

const (
	instance  = "fake"
	tableName = "table1"
	username  = "username"
	password  = "password"
)

func TestIntegrationTableCollection_Get2(t *testing.T) {
	httpmock.Activate(t)
	defer httpmock.DeactivateAndReset()

	type attr struct {
		key, expected string
	}

	tests := []struct {
		name          string
		statusCode    int
		responseBody  string
		expectErr     bool
		expectedErr   *core.ServiceNowError
		expectedAttrs []attr
	}{
		{
			name:       "success",
			statusCode: 200,
			responseBody: `{
              "result": [{
                "upon_approval": "proceed",
                "location": "",
                "expected_start": "",
                "reopen_count": "0",
                "close_notes": "",
                "impact": "2",
                "urgency": "2",
                "priority": "3",
                "number": "INC0010002",
                "category": "inquiry",
                "incident_state": "1",
                "active": "true",
                "short_description": "Unable to connect to office wifi",
                "sys_updated_on": "2016-01-22 14:28:24",
                "sys_updated_by": "admin",
                "sys_id": "c537bae64f411200adf9f8e18110c76e",
                "approval": "not requested",
                "severity": "3",
                "sys_created_by": "admin",
                "sys_created_on": "2016-01-22 14:28:24",
                "sys_class_name": "incident",
                "contact_type": "phone"
              }]
            }`,
			expectErr: false,
			expectedAttrs: []attr{
				{"upon_approval", "proceed"},
				{"location", ""},
				{"expected_start", ""},
				{"reopen_count", "0"},
				{"close_notes", ""},
				{"impact", "2"},
				{"urgency", "2"},
				{"priority", "3"},
				{"number", "INC0010002"},
				{"category", "inquiry"},
				{"incident_state", "1"},
				{"active", "true"},
				{"short_description", "Unable to connect to office wifi"},
				{"sys_updated_on", "2016-01-22 14:28:24"},
				{"sys_updated_by", "admin"},
				{"sys_id", "c537bae64f411200adf9f8e18110c76e"},
				{"approval", "not requested"},
				{"severity", "3"},
				{"sys_created_by", "admin"},
				{"sys_created_on", "2016-01-22 14:28:24"},
				{"sys_class_name", "incident"},
				{"contact_type", "phone"},
			},
		},
		{
			name:       "error",
			statusCode: 500,
			responseBody: `{
              "error": {
                "message": "Internal Server Error",
                "detail": "An unexpected error occurred."
              }
            }`,
			expectErr:   true,
			expectedErr: &core.ServiceNowError{Exception: core.Exception{Detail: "An unexpected error occurred.", Message: "Internal Server Error"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Reset()
			httpmock.RegisterResponder(
				"GET",
				fmt.Sprintf("https://%s.service-now.com/api/now/table/%s", instance, tableName),
				httpmock.NewStringResponder(tt.statusCode, tt.responseBody),
			)

			cred := credentials.NewUsernamePasswordCredential(username, password)
			client, err := servicenowsdkgo.NewServiceNowClient2WithHTTPClient(cred, instance, http.DefaultClient)
			require.NoError(t, err)

			ctx := context.Background()
			resp, err := client.Now2().Table2(tableName).Get2(ctx, nil)

			if tt.expectErr {
				require.Equal(t, tt.expectedErr, err)
				return
			}

			require.NoError(t, err)
			require.Len(t, resp.Result, 1)
			result := resp.Result[0]

			for _, attr := range tt.expectedAttrs {
				val, err := result.Value(attr.key).String()
				require.NoError(t, err)
				require.Equal(t, attr.expected, val)
			}
		})
	}
}
