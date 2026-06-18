package attachmentapi

import (
	"io"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateMultipartBody creates a multipart/form-data body for attachment uploads.
// It maps the file data to the expected ServiceNow attachment API fields.
func CreateMultipartBody(fileName string, contentType string, fileContent io.Reader, tableName string, tableSysID string) (serialization.Parsable, error) {
	body := abstractions.NewMultipartBody()

	// Add file content
	body.AddOrReplacePart("file", fileName, fileContent)

	// Add table details
	body.AddOrReplacePart("table_name", "", tableName)
	body.AddOrReplacePart("table_sys_id", "", tableSysID)
	
	return body, nil
}
