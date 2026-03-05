package attachmentapi

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAttachmentCollectionRequestMarshal(t *testing.T) {
	updatedOn, _ := time.Parse("2006-01-02 15:04:05", "2009-05-21 04:12:21")
	sysCreatedOn, _ := time.Parse("2006-01-02 15:04:05", "2009-05-21 04:12:21")

	tests := []struct {
		name     string
		json     string
		expected AttachmentCollectionResponse
	}{
		{
			name: "Valid collection response",
			json: `{
				"result": [
				  {
					"table_sys_id": "5054b6f8c0a800060056addcf551ecf8",
					"size_bytes": "462",
					"download_link": "https://instance.service-now.com/api/now/attachment/615ea769c0a80166001cf5f2367302f5/file",
					"sys_updated_on": "2009-05-21 04:12:21",
					"sys_id": "615ea769c0a80166001cf5f2367302f5",
					"image_height": "",
					"sys_created_on": "2009-05-21 04:12:21",
					"file_name": "blocks.swf",
					"sys_created_by": "glide.maint",
					"compressed": "true",
					"average_image_color": "",
					"sys_updated_by": "glide.maint",
					"sys_tags": "",
					"table_name": "content_block_programmatic",
					"image_width": "",
					"sys_mod_count": "0",
					"content_type": "application/x-shockwave-flash",
					"size_compressed": "485"
				  }
				]
			  }`,
			expected: AttachmentCollectionResponse{
				Result: []*Attachment{
					{
						TableSysId:        "5054b6f8c0a800060056addcf551ecf8",
						Size:              462,
						DownloadLink:      "https://instance.service-now.com/api/now/attachment/615ea769c0a80166001cf5f2367302f5/file",
						UpdatedOn:         Time(updatedOn),
						SysId:             "615ea769c0a80166001cf5f2367302f5",
						ImageHeight:       0,
						SysCreatedOn:      Time(sysCreatedOn),
						FileName:          "blocks.swf",
						SysCreatedBy:      "glide.maint",
						Compressed:        true,
						AverageImageColor: "",
						SysUpdatedBy:      "glide.maint",
						SysTags:           "",
						TableName:         "content_block_programmatic",
						ImageWidth:        0,
						ContentType:       "application/x-shockwave-flash",
						SizeCompressed:    485,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var attachmentCollectionResponse AttachmentCollectionResponse
			err := json.Unmarshal([]byte(test.json), &attachmentCollectionResponse)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, attachmentCollectionResponse)
		})
	}
}

func TestAttachmentCollectionResponse_ParseHeaders(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Nil headers",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var attachmentCollectionResponse AttachmentCollectionResponse
			attachmentCollectionResponse.ParseHeaders(nil)
		})
	}
}
