package tableapi

type requestBuilder interface {
}

type TableItemRequestBuilder2 struct {
    requestBuilder
}

func NewTableItemRequestBuilder2(client core.Client, pathParameters map[string]string) *TableItemRequestBuilder2 {
   return &TableItemRequestBuilder2{
        core.NewRequestBuilder(
		client,
		"{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}",
		pathParameters,
	)
   }
}


