package attachmentapi

type AttachmentRequestBuilderFileQueryParameters struct {
	//EncryptionContext Sys_id of an encryption context record.
	//Specify this parameter to allow only users with the specified encryption context to access the attachment.
	//For additional information on encryption context records,
	// see [Encryption Support]:https://docs.servicenow.com/csh?topicname=c_EncryptionSupport&version=vancouver&pubname=vancouver-platform-security.
	EncryptionContext string `url:"encryption_context"`
	//FileName Name to give the attachment.
	FileName string `url:"file_name"`
	//TableName Name of the table to attach the file to.
	TableName string `url:"table_name"`
	//TableSysId Sys_id of the record in the table specified in table_name that you want to attach the file to.
	TableSysId string `url:"table_sys_id"` //nolint:stylecheck
}
