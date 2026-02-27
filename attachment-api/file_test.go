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
	res := NewFile()
	assert.NotNil(t, res)
}

func TestFileModel_GetFieldDeserializers(t *testing.T) {
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
}

func TestFileModel_Serialize(t *testing.T) {
	m := NewFile()
	
	// Set all fields to cover all branches in Serialize
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

	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)

	err := m.Serialize(writer)
	assert.NoError(t, err)

	var nilM *FileModel
	err = nilM.Serialize(writer)
	assert.NoError(t, err)
}

func TestFileModel_Accessors(t *testing.T) {
	m := NewFile()
	
	t.Run("AverageImageColor", func(t *testing.T) {
		s := "red"
		_ = m.SetAverageImageColor(&s)
		res, _ := m.GetAverageImageColor()
		assert.Equal(t, s, *res)
	})

	t.Run("Compressed", func(t *testing.T) {
		v := true
		_ = m.SetCompressed(&v)
		res, _ := m.GetCompressed()
		assert.Equal(t, v, *res)
	})

	t.Run("ContentType", func(t *testing.T) {
		s := "text/plain"
		_ = m.SetContentType(&s)
		res, _ := m.GetContentType()
		assert.Equal(t, s, *res)
	})

	t.Run("CreatedByName", func(t *testing.T) {
		s := "admin"
		_ = m.SetCreatedByName(&s)
		res, _ := m.GetCreatedByName()
		assert.Equal(t, s, *res)
	})

	t.Run("DownloadLink", func(t *testing.T) {
		s := "http://link"
		_ = m.SetDownloadLink(&s)
		res, _ := m.GetDownloadLink()
		assert.Equal(t, s, *res)
	})

	t.Run("FileName", func(t *testing.T) {
		s := "file.txt"
		_ = m.SetFileName(&s)
		res, _ := m.GetFileName()
		assert.Equal(t, s, *res)
	})

	t.Run("ImageHeight", func(t *testing.T) {
		v := 100.0
		_ = m.SetImageHeight(&v)
		res, _ := m.GetImageHeight()
		assert.Equal(t, v, *res)
	})

	t.Run("ImageWidth", func(t *testing.T) {
		v := 200.0
		_ = m.SetImageWidth(&v)
		res, _ := m.GetImageWidth()
		assert.Equal(t, v, *res)
	})

	t.Run("SizeBytes", func(t *testing.T) {
		v := int64(1024)
		_ = m.SetSizeBytes(&v)
		res, _ := m.GetSizeBytes()
		assert.Equal(t, v, *res)
	})

	t.Run("SizeCompressed", func(t *testing.T) {
		v := int64(512)
		_ = m.SetSizeCompressed(&v)
		res, _ := m.GetSizeCompressed()
		assert.Equal(t, v, *res)
	})

	t.Run("SysCreatedBy", func(t *testing.T) {
		s := "user"
		_ = m.SetSysCreatedBy(&s)
		res, _ := m.GetSysCreatedBy()
		assert.Equal(t, s, *res)
	})

	t.Run("SysCreatedOn", func(t *testing.T) {
		v := time.Now()
		_ = m.SetSysCreatedOn(&v)
		res, _ := m.GetSysCreatedOn()
		assert.True(t, res.Equal(v))
	})

	t.Run("SysID", func(t *testing.T) {
		s := "sys_id"
		_ = m.SetSysID(&s)
		res, _ := m.GetSysID()
		assert.Equal(t, s, *res)
	})

	t.Run("SysModCount", func(t *testing.T) {
		v := int64(5)
		_ = m.SetSysModCount(&v)
		res, _ := m.GetSysModCount()
		assert.Equal(t, v, *res)
	})

	t.Run("SysTags", func(t *testing.T) {
		v := []string{"tag1"}
		_ = m.SetSysTags(v)
		res, _ := m.GetSysTags()
		assert.Equal(t, v[0], res[0])
	})

	t.Run("SysUpdatedBy", func(t *testing.T) {
		s := "updater"
		_ = m.SetSysUpdatedBy(&s)
		res, _ := m.GetSysUpdatedBy()
		assert.Equal(t, s, *res)
	})

	t.Run("SysUpdatedOn", func(t *testing.T) {
		v := time.Now()
		_ = m.SetSysUpdatedOn(&v)
		res, _ := m.GetSysUpdatedOn()
		assert.True(t, res.Equal(v))
	})

	t.Run("TableName", func(t *testing.T) {
		s := "incident"
		_ = m.SetTableName(&s)
		res, _ := m.GetTableName()
		assert.Equal(t, s, *res)
	})

	t.Run("TableSysID", func(t *testing.T) {
		s := "table_sid"
		_ = m.SetTableSysID(&s)
		res, _ := m.GetTableSysID()
		assert.Equal(t, s, *res)
	})

	t.Run("UpdatedByName", func(t *testing.T) {
		s := "updated_by"
		_ = m.SetUpdatedByName(&s)
		res, _ := m.GetUpdatedByName()
		assert.Equal(t, s, *res)
	})
}

func TestFileModel_ErrorBranches(t *testing.T) {
	m := NewFile()
	
	t.Run("AverageImageColor_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(averageImageColorKey, 123)
		_, err := m.GetAverageImageColor()
		assert.Error(t, err)
	})

	t.Run("Compressed_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(compressedKey, 123)
		_, err := m.GetCompressed()
		assert.Error(t, err)
	})

	t.Run("ContentType_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(contentTypeKey, 123)
		_, err := m.GetContentType()
		assert.Error(t, err)
	})

	t.Run("CreatedByName_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(createdByNameKey, 123)
		_, err := m.GetCreatedByName()
		assert.Error(t, err)
	})

	t.Run("DownloadLink_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(downloadLinkKey, 123)
		_, err := m.GetDownloadLink()
		assert.Error(t, err)
	})

	t.Run("FileName_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(fileNameKey, 123)
		_, err := m.GetFileName()
		assert.Error(t, err)
	})

	t.Run("ImageHeight_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(imageHeightKey, 123)
		_, err := m.GetImageHeight()
		assert.Error(t, err)
	})

	t.Run("ImageWidth_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(imageWidthKey, 123)
		_, err := m.GetImageWidth()
		assert.Error(t, err)
	})

	t.Run("SizeBytes_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sizeBytesKey, 123)
		_, err := m.GetSizeBytes()
		assert.Error(t, err)
	})

	t.Run("SizeCompressed_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sizeCompressedKey, 123)
		_, err := m.GetSizeCompressed()
		assert.Error(t, err)
	})

	t.Run("SysCreatedBy_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sysCreatedByKey, 123)
		_, err := m.GetSysCreatedBy()
		assert.Error(t, err)
	})

	t.Run("SysCreatedOn_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sysCreatedOnKey, 123)
		_, err := m.GetSysCreatedOn()
		assert.Error(t, err)
	})

	t.Run("SysID_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sysIDKey, 123)
		_, err := m.GetSysID()
		assert.Error(t, err)
	})

	t.Run("SysModCount_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sysModCountKey, 123)
		_, err := m.GetSysModCount()
		assert.Error(t, err)
	})

	t.Run("SysTags_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sysTagsKey, 123)
		_, err := m.GetSysTags()
		assert.Error(t, err)
	})

	t.Run("SysUpdatedBy_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sysUpdatedByKey, 123)
		_, err := m.GetSysUpdatedBy()
		assert.Error(t, err)
	})

	t.Run("SysUpdatedOn_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(sysUpdatedOnKey, 123)
		_, err := m.GetSysUpdatedOn()
		assert.Error(t, err)
	})

	t.Run("TableName_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(tableNameKey, 123)
		_, err := m.GetTableName()
		assert.Error(t, err)
	})

	t.Run("TableSysID_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(tableSysIDKey, 123)
		_, err := m.GetTableSysID()
		assert.Error(t, err)
	})

	t.Run("UpdatedByName_WrongType", func(t *testing.T) {
		_ = m.GetBackingStore().Set(updatedByNameKey, 123)
		_, err := m.GetUpdatedByName()
		assert.Error(t, err)
	})
	
	t.Run("NilReceiver_Accessors", func(t *testing.T) {
		var nilM *FileModel
		v1, _ := nilM.GetAverageImageColor()
		assert.Nil(t, v1)
		v2, _ := nilM.GetCompressed()
		assert.Nil(t, v2)
		v3, _ := nilM.GetContentType()
		assert.Nil(t, v3)
		v4, _ := nilM.GetCreatedByName()
		assert.Nil(t, v4)
		v5, _ := nilM.GetDownloadLink()
		assert.Nil(t, v5)
		v6, _ := nilM.GetFileName()
		assert.Nil(t, v6)
		v7, _ := nilM.GetImageHeight()
		assert.Nil(t, v7)
		v8, _ := nilM.GetImageWidth()
		assert.Nil(t, v8)
		v9, _ := nilM.GetSizeBytes()
		assert.Nil(t, v9)
		v10, _ := nilM.GetSizeCompressed()
		assert.Nil(t, v10)
		v11, _ := nilM.GetSysCreatedBy()
		assert.Nil(t, v11)
		v12, _ := nilM.GetSysCreatedOn()
		assert.Nil(t, v12)
		v13, _ := nilM.GetSysID()
		assert.Nil(t, v13)
		v14, _ := nilM.GetSysModCount()
		assert.Nil(t, v14)
		v15, _ := nilM.GetSysTags()
		assert.Nil(t, v15)
		v16, _ := nilM.GetSysUpdatedBy()
		assert.Nil(t, v16)
		v17, _ := nilM.GetSysUpdatedOn()
		assert.Nil(t, v17)
		v18, _ := nilM.GetTableName()
		assert.Nil(t, v18)
		v19, _ := nilM.GetTableSysID()
		assert.Nil(t, v19)
		v20, _ := nilM.GetUpdatedByName()
		assert.Nil(t, v20)
		
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
	writer := mocking.NewMockSerializationWriter()

	keys := []string{
		averageImageColorKey, compressedKey, contentTypeKey, createdByNameKey,
		downloadLinkKey, fileNameKey, imageHeightKey, imageWidthKey,
		sizeBytesKey, sizeCompressedKey, sysCreatedByKey, sysCreatedOnKey,
		sysIDKey, sysModCountKey, sysTagsKey, sysUpdatedByKey,
		sysUpdatedOnKey, tableNameKey, tableSysIDKey, updatedByNameKey,
	}

	for _, key := range keys {
		t.Run(key, func(t *testing.T) {
			m := NewFile()
			_ = m.GetBackingStore().Set(key, 123) // Poison with wrong type
			err := m.Serialize(writer)
			assert.Error(t, err)
		})
	}
}
