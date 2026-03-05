package attachmentapi

import (
	"errors"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create file",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := NewFile()
			assert.NotNil(t, res)
		})
	}
}

func TestFileModel_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Standard flow",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewFile()
			deser := m.GetFieldDeserializers()
			assert.NotNil(t, deser[averageImageColorKey])

			for key, fn := range deser {
				node := mocking.NewMockParseNode()
				s := "test"
				switch key {
				case sizeBytesKey, sizeCompressedKey, sysModCountKey:
					s = "1"
				case compressedKey:
					s = "true"
				case sysCreatedOnKey, sysUpdatedOnKey:
					s = "2006-01-02 15:04:05"
				case imageHeightKey, imageWidthKey:
					s = "1.2"
				}
				node.On("GetStringValue").Return(&s, nil)
				_ = fn(node)
			}

			// Test error branches
			for _, fn := range deser {
				// Test Read Error
				nodeReadError := mocking.NewMockParseNode()
				nodeReadError.On("GetStringValue").Return((*string)(nil), errors.New("read error"))
				_ = fn(nodeReadError)

				// Test Parse Error
				nodeParseError := mocking.NewMockParseNode()
				s := "not-a-value"
				nodeParseError.On("GetStringValue").Return(&s, nil)
				_ = fn(nodeParseError)
			}

			var nilM *FileModel
			assert.NotNil(t, nilM.GetFieldDeserializers())
		})
	}
}

func TestFileModel_Serialize(t *testing.T) {
	tests := []struct {
		name string
		m    *FileModel
	}{
		{
			name: "Standard serialize",
			m: func() *FileModel {
				m := NewFile()
				s := "test"
				b := true
				f := 1.2
				i := int64(10)
				tm := time.Now()
				tags := []string{"tag"}

				_ = m.SetAverageImageColor(&s)
				_ = m.SetCompressed(&b)
				_ = m.SetContentType(&s)
				_ = m.SetCreatedByName(&s)
				_ = m.SetDownloadLink(&s)
				_ = m.SetFileName(&s)
				_ = m.SetImageHeight(&f)
				_ = m.SetImageWidth(&f)
				_ = m.SetSizeBytes(&i)
				_ = m.SetSizeCompressed(&i)
				_ = m.SetSysCreatedBy(&s)
				_ = m.SetSysCreatedOn(&tm)
				_ = m.SetSysID(&s)
				_ = m.SetSysModCount(&i)
				_ = m.SetSysTags(tags)
				_ = m.SetSysUpdatedBy(&s)
				_ = m.SetSysUpdatedOn(&tm)
				_ = m.SetTableName(&s)
				_ = m.SetTableSysID(&s)
				_ = m.SetUpdatedByName(&s)
				return m
			}(),
		},
		{
			name: "Nil model",
			m:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			writer.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
			writer.On("WriteCollectionOfPrimitiveValues", mock.Anything, mock.Anything).Return(nil)
			writer.On("WriteObjectValue", mock.Anything, mock.Anything).Return(nil)
			writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

			err := test.m.Serialize(writer)
			assert.NoError(t, err)
		})
	}
}

func TestFileModel_Accessors(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name   string
		set    func(m *FileModel)
		get    func(m *FileModel) (interface{}, error)
		expect interface{}
	}{
		{"AverageImageColor", func(m *FileModel) { s := "red"; _ = m.SetAverageImageColor(&s) }, func(m *FileModel) (interface{}, error) { return m.GetAverageImageColor() }, "red"},
		{"Compressed", func(m *FileModel) { v := true; _ = m.SetCompressed(&v) }, func(m *FileModel) (interface{}, error) { return m.GetCompressed() }, true},
		{"ContentType", func(m *FileModel) { s := "text/plain"; _ = m.SetContentType(&s) }, func(m *FileModel) (interface{}, error) { return m.GetContentType() }, "text/plain"},
		{"CreatedByName", func(m *FileModel) { s := "admin"; _ = m.SetCreatedByName(&s) }, func(m *FileModel) (interface{}, error) { return m.GetCreatedByName() }, "admin"},
		{"DownloadLink", func(m *FileModel) { s := "http://link"; _ = m.SetDownloadLink(&s) }, func(m *FileModel) (interface{}, error) { return m.GetDownloadLink() }, "http://link"},
		{"FileName", func(m *FileModel) { s := "file.txt"; _ = m.SetFileName(&s) }, func(m *FileModel) (interface{}, error) { return m.GetFileName() }, "file.txt"},
		{"ImageHeight", func(m *FileModel) { v := 100.0; _ = m.SetImageHeight(&v) }, func(m *FileModel) (interface{}, error) { return m.GetImageHeight() }, 100.0},
		{"ImageWidth", func(m *FileModel) { v := 200.0; _ = m.SetImageWidth(&v) }, func(m *FileModel) (interface{}, error) { return m.GetImageWidth() }, 200.0},
		{"SizeBytes", func(m *FileModel) { v := int64(1024); _ = m.SetSizeBytes(&v) }, func(m *FileModel) (interface{}, error) { return m.GetSizeBytes() }, int64(1024)},
		{"SizeCompressed", func(m *FileModel) { v := int64(512); _ = m.SetSizeCompressed(&v) }, func(m *FileModel) (interface{}, error) { return m.GetSizeCompressed() }, int64(512)},
		{"SysCreatedBy", func(m *FileModel) { s := "user"; _ = m.SetSysCreatedBy(&s) }, func(m *FileModel) (interface{}, error) { return m.GetSysCreatedBy() }, "user"},
		{"SysCreatedOn", func(m *FileModel) { _ = m.SetSysCreatedOn(&now) }, func(m *FileModel) (interface{}, error) { return m.GetSysCreatedOn() }, &now},
		{"SysID", func(m *FileModel) { s := "sys_id"; _ = m.SetSysID(&s) }, func(m *FileModel) (interface{}, error) { return m.GetSysID() }, "sys_id"},
		{"SysModCount", func(m *FileModel) { v := int64(5); _ = m.SetSysModCount(&v) }, func(m *FileModel) (interface{}, error) { return m.GetSysModCount() }, int64(5)},
		{"SysTags", func(m *FileModel) { v := []string{"tag1"}; _ = m.SetSysTags(v) }, func(m *FileModel) (interface{}, error) { return m.GetSysTags() }, []string{"tag1"}},
		{"SysUpdatedBy", func(m *FileModel) { s := "updater"; _ = m.SetSysUpdatedBy(&s) }, func(m *FileModel) (interface{}, error) { return m.GetSysUpdatedBy() }, "updater"},
		{"SysUpdatedOn", func(m *FileModel) { _ = m.SetSysUpdatedOn(&now) }, func(m *FileModel) (interface{}, error) { return m.GetSysUpdatedOn() }, &now},
		{"TableName", func(m *FileModel) { s := "incident"; _ = m.SetTableName(&s) }, func(m *FileModel) (interface{}, error) { return m.GetTableName() }, "incident"},
		{"TableSysID", func(m *FileModel) { s := "table_sid"; _ = m.SetTableSysID(&s) }, func(m *FileModel) (interface{}, error) { return m.GetTableSysID() }, "table_sid"},
		{"UpdatedByName", func(m *FileModel) { s := "updated_by"; _ = m.SetUpdatedByName(&s) }, func(m *FileModel) (interface{}, error) { return m.GetUpdatedByName() }, "updated_by"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewFile()
			test.set(m)
			res, err := test.get(m)
			assert.NoError(t, err)
			switch v := res.(type) {
			case *string:
				assert.Equal(t, test.expect, *v)
			case *bool:
				assert.Equal(t, test.expect, *v)
			case *float64:
				assert.Equal(t, test.expect, *v)
			case *int64:
				assert.Equal(t, test.expect, *v)
			case *time.Time:
				assert.True(t, v.Equal(*(test.expect.(*time.Time))))
			case []string:
				assert.Equal(t, test.expect, v)
			default:
				t.Fatalf("unexpected type %T", v)
			}
		})
	}
}

func TestFileModel_ErrorBranches(t *testing.T) {
	tests := []struct {
		name string
		key  string
		get  func(m *FileModel) (interface{}, error)
	}{
		{"AverageImageColor", averageImageColorKey, func(m *FileModel) (interface{}, error) { return m.GetAverageImageColor() }},
		{"Compressed", compressedKey, func(m *FileModel) (interface{}, error) { return m.GetCompressed() }},
		{"ContentType", contentTypeKey, func(m *FileModel) (interface{}, error) { return m.GetContentType() }},
		{"CreatedByName", createdByNameKey, func(m *FileModel) (interface{}, error) { return m.GetCreatedByName() }},
		{"DownloadLink", downloadLinkKey, func(m *FileModel) (interface{}, error) { return m.GetDownloadLink() }},
		{"FileName", fileNameKey, func(m *FileModel) (interface{}, error) { return m.GetFileName() }},
		{"ImageHeight", imageHeightKey, func(m *FileModel) (interface{}, error) { return m.GetImageHeight() }},
		{"ImageWidth", imageWidthKey, func(m *FileModel) (interface{}, error) { return m.GetImageWidth() }},
		{"SizeBytes", sizeBytesKey, func(m *FileModel) (interface{}, error) { return m.GetSizeBytes() }},
		{"SizeCompressed", sizeCompressedKey, func(m *FileModel) (interface{}, error) { return m.GetSizeCompressed() }},
		{"SysCreatedBy", sysCreatedByKey, func(m *FileModel) (interface{}, error) { return m.GetSysCreatedBy() }},
		{"SysCreatedOn", sysCreatedOnKey, func(m *FileModel) (interface{}, error) { return m.GetSysCreatedOn() }},
		{"SysID", sysIDKey, func(m *FileModel) (interface{}, error) { return m.GetSysID() }},
		{"SysModCount", sysModCountKey, func(m *FileModel) (interface{}, error) { return m.GetSysModCount() }},
		{"SysTags", sysTagsKey, func(m *FileModel) (interface{}, error) { return m.GetSysTags() }},
		{"SysUpdatedBy", sysUpdatedByKey, func(m *FileModel) (interface{}, error) { return m.GetSysUpdatedBy() }},
		{"SysUpdatedOn", sysUpdatedOnKey, func(m *FileModel) (interface{}, error) { return m.GetSysUpdatedOn() }},
		{"TableName", tableNameKey, func(m *FileModel) (interface{}, error) { return m.GetTableName() }},
		{"TableSysID", tableSysIDKey, func(m *FileModel) (interface{}, error) { return m.GetTableSysID() }},
		{"UpdatedByName", updatedByNameKey, func(m *FileModel) (interface{}, error) { return m.GetUpdatedByName() }},
	}

	for _, test := range tests {
		t.Run(test.name+"_WrongType", func(t *testing.T) {
			m := NewFile()
			_ = m.GetBackingStore().Set(test.key, 123)
			_, err := test.get(m)
			assert.Error(t, err)
		})
	}

	t.Run("NilReceiver_Accessors", func(t *testing.T) {
		var nilM *FileModel
		for _, test := range tests {
			res, err := test.get(nilM)
			assert.NoError(t, err)
			assert.Nil(t, res)
		}

		assert.Nil(t, nilM.SetAverageImageColor(nil))
		assert.Nil(t, nilM.SetCompressed(nil))
		assert.Nil(t, nilM.SetContentType(nil))
		assert.Nil(t, nilM.SetCreatedByName(nil))
		assert.Nil(t, nilM.SetDownloadLink(nil))
		assert.Nil(t, nilM.SetFileName(nil))
		assert.Nil(t, nilM.SetImageHeight(nil))
		assert.Nil(t, nilM.SetImageWidth(nil))
		assert.Nil(t, nilM.SetSizeBytes(nil))
		assert.Nil(t, nilM.SetSizeCompressed(nil))
		assert.Nil(t, nilM.SetSysCreatedBy(nil))
		assert.Nil(t, nilM.SetSysCreatedOn(nil))
		assert.Nil(t, nilM.SetSysID(nil))
		assert.Nil(t, nilM.SetSysModCount(nil))
		assert.Nil(t, nilM.SetSysTags(nil))
		assert.Nil(t, nilM.SetSysUpdatedBy(nil))
		assert.Nil(t, nilM.SetSysUpdatedOn(nil))
		assert.Nil(t, nilM.SetTableName(nil))
		assert.Nil(t, nilM.SetTableSysID(nil))
		assert.Nil(t, nilM.SetUpdatedByName(nil))
	})
}

func TestFileModel_Serialize_Errors(t *testing.T) {
	keys := []string{
		averageImageColorKey, compressedKey, contentTypeKey, createdByNameKey,
		downloadLinkKey, fileNameKey, imageHeightKey, imageWidthKey,
		sizeBytesKey, sizeCompressedKey, sysCreatedByKey, sysCreatedOnKey,
		sysIDKey, sysModCountKey, sysTagsKey, sysUpdatedByKey,
		sysUpdatedOnKey, tableNameKey, tableSysIDKey, updatedByNameKey,
	}

	for _, key := range keys {
		t.Run(key, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			writer.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
			writer.On("WriteCollectionOfPrimitiveValues", mock.Anything, mock.Anything).Return(nil)
			writer.On("WriteObjectValue", mock.Anything, mock.Anything).Return(nil)
			writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

			m := NewFile()
			_ = m.GetBackingStore().Set(key, 123) // Poison with wrong type
			err := m.Serialize(writer)
			assert.Error(t, err)
		})
	}
}
