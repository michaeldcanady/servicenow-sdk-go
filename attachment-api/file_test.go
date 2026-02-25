package attachmentapi

import (
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {
	res := NewFile()
	if res == nil {
		t.Error("returned nil")
	}
}

func TestFileModel_GetFieldDeserializers(t *testing.T) {
	m := NewFile()
	deser := m.GetFieldDeserializers()
	if deser[averageImageColorKey] == nil {
		t.Error("missing deserializer")
	}
	var nilM *FileModel
	if nilM.GetFieldDeserializers() == nil {
		t.Error("nil receiver should return new file deser")
	}
}

func TestFileModel_Serialize(t *testing.T) {
	// Minimal test for coverage
	m := NewFile()
	_ = m.Serialize(nil)
	var nilM *FileModel
	_ = nilM.Serialize(nil)
}

func TestFileModel_Accessors(t *testing.T) {
	m := NewFile()
	
	t.Run("AverageImageColor", func(t *testing.T) {
		s := "red"
		_ = m.SetAverageImageColor(&s)
		res, _ := m.GetAverageImageColor()
		if *res != s { t.Error("failed") }
	})

	t.Run("Compressed", func(t *testing.T) {
		v := true
		_ = m.SetCompressed(&v)
		res, _ := m.GetCompressed()
		if *res != v { t.Error("failed") }
	})

	t.Run("ContentType", func(t *testing.T) {
		s := "text/plain"
		_ = m.SetContentType(&s)
		res, _ := m.GetContentType()
		if *res != s { t.Error("failed") }
	})

	t.Run("CreatedByName", func(t *testing.T) {
		s := "admin"
		_ = m.SetCreatedByName(&s)
		res, _ := m.GetCreatedByName()
		if *res != s { t.Error("failed") }
	})

	t.Run("DownloadLink", func(t *testing.T) {
		s := "http://link"
		_ = m.SetDownloadLink(&s)
		res, _ := m.GetDownloadLink()
		if *res != s { t.Error("failed") }
	})

	t.Run("FileName", func(t *testing.T) {
		s := "file.txt"
		_ = m.SetFileName(&s)
		res, _ := m.GetFileName()
		if *res != s { t.Error("failed") }
	})

	t.Run("ImageHeight", func(t *testing.T) {
		v := 100.0
		_ = m.SetImageHeight(&v)
		res, _ := m.GetImageHeight()
		if *res != v { t.Error("failed") }
	})

	t.Run("ImageWidth", func(t *testing.T) {
		v := 200.0
		_ = m.SetImageWidth(&v)
		res, _ := m.GetImageWidth()
		if *res != v { t.Error("failed") }
	})

	t.Run("SizeBytes", func(t *testing.T) {
		v := int64(1024)
		_ = m.SetSizeBytes(&v)
		res, _ := m.GetSizeBytes()
		if *res != v { t.Error("failed") }
	})

	t.Run("SizeCompressed", func(t *testing.T) {
		v := int64(512)
		_ = m.SetSizeCompressed(&v)
		res, _ := m.GetSizeCompressed()
		if *res != v { t.Error("failed") }
	})

	t.Run("SysCreatedBy", func(t *testing.T) {
		s := "user"
		_ = m.SetSysCreatedBy(&s)
		res, _ := m.GetSysCreatedBy()
		if *res != s { t.Error("failed") }
	})

	t.Run("SysCreatedOn", func(t *testing.T) {
		v := time.Now()
		_ = m.SetSysCreatedOn(&v)
		res, _ := m.GetSysCreatedOn()
		if !res.Equal(v) { t.Error("failed") }
	})

	t.Run("SysID", func(t *testing.T) {
		s := "sys_id"
		_ = m.SetSysID(&s)
		res, _ := m.GetSysID()
		if *res != s { t.Error("failed") }
	})

	t.Run("SysModCount", func(t *testing.T) {
		v := int64(5)
		_ = m.SetSysModCount(&v)
		res, _ := m.GetSysModCount()
		if *res != v { t.Error("failed") }
	})

	t.Run("SysTags", func(t *testing.T) {
		v := []string{"tag1"}
		_ = m.SetSysTags(v)
		res, _ := m.GetSysTags()
		if res[0] != v[0] { t.Error("failed") }
	})

	t.Run("SysUpdatedBy", func(t *testing.T) {
		s := "updater"
		_ = m.SetSysUpdatedBy(&s)
		res, _ := m.GetSysUpdatedBy()
		if *res != s { t.Error("failed") }
	})

	t.Run("SysUpdatedOn", func(t *testing.T) {
		v := time.Now()
		_ = m.SetSysUpdatedOn(&v)
		res, _ := m.GetSysUpdatedOn()
		if !res.Equal(v) { t.Error("failed") }
	})

	t.Run("TableName", func(t *testing.T) {
		s := "incident"
		_ = m.SetTableName(&s)
		res, _ := m.GetTableName()
		if *res != s { t.Error("failed") }
	})

	t.Run("TableSysID", func(t *testing.T) {
		s := "table_sid"
		_ = m.SetTableSysID(&s)
		res, _ := m.GetTableSysID()
		if *res != s { t.Error("failed") }
	})

	t.Run("UpdatedByName", func(t *testing.T) {
		s := "updated_by"
		_ = m.SetUpdatedByName(&s)
		res, _ := m.GetUpdatedByName()
		if *res != s { t.Error("failed") }
	})
}

func TestFileModel_ErrorBranches(t *testing.T) {
	m := NewFile()
	_ = m.GetBackingStore().Set(averageImageColorKey, 123)
	if _, err := m.GetAverageImageColor(); err == nil {
		t.Error("expected err for wrong type")
	}
}
