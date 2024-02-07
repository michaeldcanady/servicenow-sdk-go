package attachmentapi

type AttachmentRequestBuilderGetQueryParameters struct {
	// Limit Limit to be applied on pagination.
	Limit int `url:"sysparm_limit"`
	// Offset Number of records to exclude from the query. Use this parameter to get more records than specified in sysparm_limit. For example, if sysparm_limit is set to 500, but there are additional records you want to query, you can specify a sysparm_offset value of 500 to get the second set of records.
	Offset int `url:"sysparm_offset"`
	// Query Encoded query. Queries for the Attachment API are relative to the Attachments [sys_attachment] table.
	Query string `url:"sysparm_query"`
}
