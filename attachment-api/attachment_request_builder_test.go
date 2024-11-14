package attachmentapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/RecoLabs/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (c *MockClient) Send(ctx context.Context, requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
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

	pathParameters := map[string]string{internal.BasePathParameter: "https://instance.service-now.com/api/now"}

	req := NewAttachmentRequestBuilder(&client, pathParameters)

	assert.NotNil(t, req)
}

func TestAttachmentUrl(t *testing.T) {
	client := MockClient{}

	pathParameters := map[string]string{internal.BasePathParameter: "https://instance.service-now.com/api/now"}

	req := NewAttachmentRequestBuilder(&client, pathParameters)

	assert.Equal(t, req.PathParameters, pathParameters)

	if !reflect.DeepEqual(req.PathParameters, pathParameters) {
		t.Errorf("excepted: %s, got: %s", pathParameters, req.PathParameters)
	}
}

func TestAttachmentRequestBuilderGet(t *testing.T) {
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

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(responseJSON)) //nolint:errcheck
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error())) //nolint:errcheck
		}
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

	updatedOn, _ := time.Parse(DateTimeFormat, "2009-05-21 04:12:21")
	sysCreatedOn, _ := time.Parse(DateTimeFormat, "2009-05-21 04:12:21")

	(*expected.Result[0]).UpdatedOn = Time(updatedOn)
	(*expected.Result[0]).SysCreatedOn = Time(sysCreatedOn)

	parsedURL, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	pathParameters := map[string]string{internal.BasePathParameter: "http://" + parsedURL.Host, "table": parsedURL.Path}

	builder := NewAttachmentRequestBuilder(client, pathParameters)

	// Call the Get method
	resp, err := builder.Get(context.Background(), nil)

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

func TestAttachmentRequestBuilderFile(t *testing.T) {
	fakeUser := "fakeuser"
	today, _ := time.Parse(DateTimeFormat, "2009-05-21 04:12:21")
	formattedToday := today.Format(DateTimeFormat)
	tableName := "incident"
	tableSysID := "INC00000000"
	fileName := "testfile.txt"

	expected := &AttachmentItemResponse{
		Result: &Attachment{
			AverageImageColor: "String",
			Compressed:        false,
			ContentType:       "String",
			DownloadLink:      "String",
			FileName:          fileName,
			ImageHeight:       0,
			ImageWidth:        0,
			Size:              0,
			SizeCompressed:    0,
			SysCreatedBy:      fakeUser,
			SysCreatedOn:      Time(today),
			SysId:             "String",
			SysModCount:       0,
			SysTags:           "String",
			SysUpdatedBy:      fakeUser,
			UpdatedOn:         Time(today),
			TableName:         tableName,
			TableSysId:        tableSysID,
			//updated_by_name:   fakeUser,
		},
	}

	// Create an httptest.NewServer with a custom handler
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		//_, err := io.ReadAll(r.Body)
		//if err != nil {
		//	http.Error(w, "Error reading request body", http.StatusInternalServerError)
		//	return
		//}

		// Create a mock JSON response
		mockResponse := map[string]interface{}{
			"result": map[string]string{
				"average_image_color": "String",
				"compressed":          "false",
				"content_type":        "String",
				"created_by_name":     "String",
				"download_link":       "String",
				"file_name":           r.URL.Query()["file_name"][0],
				"image_height":        "0",
				"image_width":         "0",
				"size_bytes":          "0",
				"size_compressed":     "0",
				"sys_created_by":      fakeUser,
				"sys_created_on":      formattedToday,
				"sys_id":              "String",
				"sys_mod_count":       "0",
				"sys_tags":            "String",
				"sys_updated_by":      fakeUser,
				"sys_updated_on":      formattedToday,
				"table_name":          r.URL.Query()["table_name"][0],
				"table_sys_id":        r.URL.Query()["table_sys_id"][0],
				"updated_by_name":     fakeUser,
			},
		}

		// Convert the mock response to JSON
		responseJSON, err := json.Marshal(mockResponse)
		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		// Write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseJSON)
	}))
	defer mockServer.Close()

	// Set up a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) //nolint:errcheck

	// Write some data to the temporary file
	testData := []byte("test data")
	_, err = tempFile.Write(testData)
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	parsedURL, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	client := &MockClient{}

	pathParameters := map[string]string{internal.BasePathParameter: "http://" + parsedURL.Host}

	builder := NewAttachmentRequestBuilder(client, pathParameters)

	t.Run("Successful", func(t *testing.T) {
		params := &AttachmentRequestBuilderFileQueryParameters{
			FileName:   fileName,
			TableName:  tableName,
			TableSysId: tableSysID,
		}

		// Call the Get method
		resp, err := builder.File(context.Background(), tempFile.Name(), params)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Validate the response
		assert.Equal(t, expected, resp)
	})

	t.Run("Nil params", func(t *testing.T) {
		_, err := builder.File(context.Background(), tempFile.Name(), nil)
		assert.ErrorIs(t, ErrNilParams, err)
	})

	t.Run("Bad file", func(t *testing.T) {
		params := &AttachmentRequestBuilderFileQueryParameters{
			FileName:   fileName,
			TableName:  tableName,
			TableSysId: tableSysID,
		}

		_, err := builder.File(context.Background(), "bad-file.txt", params)
		assert.Error(t, err)
	})
}
