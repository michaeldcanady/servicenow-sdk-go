package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

//Deprecated: deprecated since v{version}. Use `TableItemRequestBuilder2[T]` instead.
//TableItemRequestBuilder Represents the base of a Table Item Request
type TableItemRequestBuilder = TableItemRequestBuilder2[tableEntry]

// Deprecated: deprecated since v{version}. Use `NewTableItemRequestBuilder2[T]` instead.
// NewTableItemRequestBuilder creates a new instance of the TableItemRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableItemRequestBuilder.
func NewTableItemRequestBuilder(client core.Client, pathParameters map[string]string) *TableItemRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}",
		pathParameters,
	)
	return &TableItemRequestBuilder{
		*requestBuilder,
	}
}
