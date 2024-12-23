package attachmentapi

// AttachmentFileRequestBuilderPostQueryParameters represents attachment file post request query parameters
type AttachmentFileRequestBuilderPostQueryParameters struct {
	// EncryptionContext Sys_id of an encryption context record.
	// Specify this parameter to allow only users with the specified encryption context to access the attachment.
	EncryptionContext *string `uri:"encryption_context,omitempty"`
	// FileName Name to give the attachment.
	FileName *string `uri:"file_name,omitempty"`
	// TableName Name of the table to attach the file to.
	TableName *string `uri:"table_name,omitempty"`
	// TableSysID Sys_id of the record in the table specified in table_name that you want to attach the file to.
	TableSysID *string `uri:"table_sys_id,omitempty"`
}
