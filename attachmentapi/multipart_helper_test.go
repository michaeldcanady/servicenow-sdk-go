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

	body, err := CreateMultipartBody(fileName, contentType, fileContent, tableName, tableSysID)

	assert.NoError(t, err)
	assert.NotNil(t, body)
}
