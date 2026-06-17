package attachmentapi

import (
	"bytes"
	"io"
	"mime/multipart"
)

// CreateMultipartBody creates a multipart/form-data body for attachment uploads.
// It maps the file data to the expected ServiceNow attachment API fields.
func CreateMultipartBody(fileName string, contentType string, fileContent io.Reader, tableName string, tableSysID string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file content
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, "", err
	}
	if _, err := io.Copy(part, fileContent); err != nil {
		return nil, "", err
	}

	// Add table details
	if err := writer.WriteField("table_name", tableName); err != nil {
		return nil, "", err
	}
	if err := writer.WriteField("table_sys_id", tableSysID); err != nil {
		return nil, "", err
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}
