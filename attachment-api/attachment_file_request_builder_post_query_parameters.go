package attachmentapi

type AttachmentFileRequestBuilderPostQueryParameters struct {
	EncryptionContext string `uri:"encryption_context"`
	FileName          string `uri:"file_name"`
	TableName         string `uri:"table_name"`
	TableSysID        string `uri:"table_sys_id"`
}
