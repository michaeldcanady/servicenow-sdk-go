package attachmentapi

// AttachmentRequestBuilder2GetQueryParameters represents attachment get request query parameters
type AttachmentRequestBuilder2GetQueryParameters struct {
	// SysparmLimit Limit to be applied on pagination.
	SysparmLimit int `uri:"sysparm_limit,omitempty"`
	// SysparmOffset Number of records to exclude from the query. Use this parameter to get more records than specified in sysparm_limit.
	SysparmOffset int `uri:"sysparm_offset,omitempty"`
	// SysparmQuery Encoded query. Queries for the Attachment API are relative to the Attachments [sys_attachment] table.
	SysparmQuery string `uri:"sysparm_query,omitempty"`
}
