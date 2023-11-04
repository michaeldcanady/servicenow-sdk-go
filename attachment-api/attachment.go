package attachmentapi

import "time"

type Attachment struct {
	TableSysId        string    `json:"table_sys_id"`
	Size              int       `json:"size_bytes"`
	DownloadLink      string    `json:"download_link"`
	UpdatedOn         time.Time `json:"sys_updated_on"`
	SysId             string    `json:"sys_id"`
	ImageHeight       float64   `json:"image_height"`
	SysCreatedOn      time.Time `json:"sys_created_on"`
	FileName          string    `json:"file_name"`
	SysCreatedBy      string    `json:"sys_created_by"`
	Compressed        bool      `json:"compressed"`
	AverageImageColor string    `json:"average_image_color"`
	SysUpdatedBy      string    `json:"sys_updated_by"`
	SysTags           []string  `json:"sys_tags"`
	TableName         string    `json:"table_name"`
	ImageWidth        float64   `json:"image_width"`
	SysModCount       int       `json:"sys_mod_count"`
	ContentType       string    `json:"content_type"`
	SizeCompressed    int       `json:"size_compressed"`
}
