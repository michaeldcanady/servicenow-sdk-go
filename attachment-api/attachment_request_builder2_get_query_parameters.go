package attachmentapi

type AttachmentRequestBuilder2GetQueryParameters struct {
	SysparmLimit  int    `uri:"sysparm_limit"`
	SysparmOffset int    `uri:"sysparm_offset"`
	SysparmQuery  string `uri:"sysparm_query"`
}
