// Step 1: Create credentials
{%
	include-markdown 'assets/snippets/credentials.go'
%}

// Step 2: Initialize client
{%
	include-markdown 'assets/snippets/client.go'
%}

// Step 3: Define raw URL
rawURL := "https://www.xSDK_SN_INSTANCEx.service-now.com/api/now/v1/table/xSDK_SN_TABLEx/xSDK_SN_TABLE_SYS_IDx"

// Step 4: Build request
requestBuilder := tableapi.NewTableItemRequestBuilder3[*tableapi.TableRecord](rawURL, client.RequestAdapter, tableapi.CreateTableRecordFromDiscriminatorValue)