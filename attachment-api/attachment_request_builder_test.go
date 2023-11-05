package attachmentapi

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (c *MockClient) Send(requestInfo *core.RequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {

	req, err := requestInfo.ToRequest()
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func TestNewAttachmentRequestBuilder(t *testing.T) {

	client := MockClient{}

	pathParameters := map[string]string{"baseurl": "https://instance.service-now.com/api/now"}

	req := NewAttachmentRequestBuilder(&client, pathParameters)

	assert.NotNil(t, req)
}

func TestAttachmentUrl(t *testing.T) {

	client := MockClient{}

	pathParameters := map[string]string{"baseurl": "https://instance.service-now.com/api/now"}

	req := NewAttachmentRequestBuilder(&client, pathParameters)

	assert.Equal(t, req.PathParameters, pathParameters)

	if !reflect.DeepEqual(req.PathParameters, pathParameters) {
		t.Errorf("excepted: %s, got: %s", pathParameters, req.PathParameters)
	}
}

func TestAttachmentRequestBuilder_Get(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON
		responseJSON := `{
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
		  }`

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseJSON))
	}))

	client := &MockClient{}

	expected := &AttachmentCollectionResponse{
		Result: []*Attachment{
			{
				TableSysId:        "5054b6f8c0a800060056addcf551ecf8",
				Size:              462,
				DownloadLink:      "https://instance.service-now.com/api/now/attachment/615ea769c0a80166001cf5f2367302f5/file",
				SysId:             "615ea769c0a80166001cf5f2367302f5",
				ImageHeight:       0,
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
	}

	updatedOn, _ := time.Parse("2006-01-02 15:04:05", "2009-05-21 04:12:21")
	sysCreatedOn, _ := time.Parse("2006-01-02 15:04:05", "2009-05-21 04:12:21")

	(*expected.Result[0]).UpdatedOn = Time(updatedOn)
	(*expected.Result[0]).SysCreatedOn = Time(sysCreatedOn)

	parsedUrl, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	pathParameters := map[string]string{"baseurl": "http://" + parsedUrl.Host, "table": parsedUrl.Path}

	builder := NewAttachmentRequestBuilder(client, pathParameters)

	// Call the Get method
	resp, err := builder.Get(nil)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	// You can further validate the response if needed
	if resp == nil {
		t.Error("Expected a non-nil response, but got nil")
	}

	expectedType := reflect.TypeOf(&AttachmentCollectionResponse{})
	if reflect.TypeOf(resp) != expectedType {
		t.Errorf("Expected response of type %v, but got type %v", expectedType, reflect.TypeOf(resp))
	}

	if len(resp.Result) != 1 {
		t.Errorf("Expected response with 1 result, but got %v", len(resp.Result))
	}
	assert.Equal(t, expected, resp)
}
