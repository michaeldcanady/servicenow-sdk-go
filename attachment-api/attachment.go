package attachmentapi

// Deprecated: deprecated since v{unreleased}.
//
// Attachment
type Attachment struct {
	TableSysId        string `json:"table_sys_id"` //nolint:stylecheck
	Size              Int    `json:"size_bytes"`
	DownloadLink      string `json:"download_link"`
	UpdatedOn         Time   `json:"sys_updated_on"`
	SysId             string `json:"sys_id"` //nolint:stylecheck
	ImageHeight       Int    `json:"image_height"`
	SysCreatedOn      Time   `json:"sys_created_on"`
	FileName          string `json:"file_name"`
	SysCreatedBy      string `json:"sys_created_by"`
	Compressed        Bool   `json:"compressed"`
	AverageImageColor string `json:"average_image_color"`
	SysUpdatedBy      string `json:"sys_updated_by"`
	SysTags           string `json:"sys_tags"`
	TableName         string `json:"table_name"`
	ImageWidth        Int    `json:"image_width"`
	SysModCount       Int    `json:"sys_mod_count"`
	ContentType       string `json:"content_type"`
	SizeCompressed    Int    `json:"size_compressed"`
}
