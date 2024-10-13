package attachmentapi

type AttachmentFileRequestBuilderPostQueryParameter struct {
	EncryptionContext string `url:"encryption_context"`
	FileName          string `url:"file_name"`
	TableName         string `url:"table_name"`
	TableSysID        string `url:"table_sys_id"`
}
