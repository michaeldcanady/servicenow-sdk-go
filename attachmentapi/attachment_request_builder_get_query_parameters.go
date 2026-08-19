package attachmentapi

// AttachmentRequestBuilderGetQueryParameters represents attachment get request query parameters
type AttachmentRequestBuilderGetQueryParameters struct {
	// SysparmLimit Limit to be applied on pagination.
	SysparmLimit *int32 `uriparametername:"sysparm_limit"`
	// SysparmOffset Number of records to exclude from the query. Use this parameter to get more records than specified in sysparm_limit.
	SysparmOffset *int32 `uriparametername:"sysparm_offset"`
	// SysparmQuery Encoded query. Queries for the Attachment API are relative to the Attachments [sys_attachment] table.
	SysparmQuery *string `uriparametername:"sysparm_query"`
}
