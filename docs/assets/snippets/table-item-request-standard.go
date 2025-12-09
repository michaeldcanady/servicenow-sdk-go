// Step 1: Create credentials
{%
	include-markdown 'assets/snippets/credentials.go'
%}

// Step 2: Initialize client
{%
	include-markdown 'assets/snippets/client.go'
%}

// Step 3: Define path parameters
pathParameters := map[string]string{
	"baseurl": "https://www.xSDK_SN_INSTANCEx.service-now.com/api/now",
	"table":   "xSDK_SN_TABLEx",
	"sysId":   "xSDK_SN_TABLE_SYS_IDx",
}

// Step 4: Build request
requestBuilder := tableapi.NewTableItemRequestBuilder2(client, pathParameters)