package attachmentapi

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMultipartBody(t *testing.T) {
	fileName := "test.txt"
	contentType := "text/plain"
	fileContent := strings.NewReader("Hello, World!")
	tableName := "incident"
	tableSysID := "sysid123"

	body, contentTypeHeader, err := CreateMultipartBody(fileName, contentType, fileContent, tableName, tableSysID)

	assert.NoError(t, err)
	assert.NotNil(t, body)
	assert.Contains(t, contentTypeHeader, "multipart/form-data")
	assert.Contains(t, body.String(), "Hello, World!")
	assert.Contains(t, body.String(), "table_name")
	assert.Contains(t, body.String(), "incident")
	assert.Contains(t, body.String(), "table_sys_id")
	assert.Contains(t, body.String(), "sysid123")
}
