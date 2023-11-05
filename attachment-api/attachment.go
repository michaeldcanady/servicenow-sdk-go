package attachmentapi

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Attachment struct {
	TableSysId        string `json:"table_sys_id"`
	Size              Int    `json:"size_bytes"`
	DownloadLink      string `json:"download_link"`
	UpdatedOn         Time   `json:"sys_updated_on"`
	SysId             string `json:"sys_id"`
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

type Int int

func (i *Int) UnmarshalJSON(data []byte) error {

	cleanData := strings.Replace(string(data), "\"", "", -1)

	if cleanData == "" {
		cleanData = "0"
	}

	cleanInt, err := strconv.Atoi(cleanData)

	*i = Int(cleanInt)

	return err
}

type Bool bool

func (i *Bool) UnmarshalJSON(data []byte) error {

	cleanData := strings.Replace(string(data), "\"", "", -1)

	cleanInt, err := strconv.ParseBool(cleanData)

	*i = Bool(cleanInt)

	return err
}

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format("2006-01-02 15:04:05"))), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {

	parsedTime, err := time.Parse("2006-01-02 15:04:05", strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}
