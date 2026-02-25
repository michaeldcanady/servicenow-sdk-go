package attachmentapi

import (
	"testing"
)

func TestNewAttachment2(t *testing.T) {
	res := NewAttachment2()
	if res == nil {
		t.Error("returned nil")
	}
}

func TestAttachment2Model_GetFieldDeserializers(t *testing.T) {
	m := NewAttachment2()
	deser := m.GetFieldDeserializers()
	if deser[tableSysIDKey] == nil {
		t.Error("missing deserializer")
	}
}

func TestAttachment2Model_Serialize(t *testing.T) {
	m := NewAttachment2()
	err := m.Serialize(nil)
	if err == nil {
		t.Error("expected err")
	}
}

func TestAttachment2Model_GetTableSysID(t *testing.T) {
	s := "sid"
	m := NewAttachment2()
	_ = m.setTableSysID(&s)
	
	tests := []struct {
		name     string
		model    *Attachment2Model
		expected *string
	}{
		{"Ok", m, &s},
		{"NilM", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := tt.model.GetTableSysID()
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestAttachment2Model_setTableSysID(t *testing.T) {
	s := "sid"
	m := NewAttachment2()
	
	tests := []struct {
		name  string
		model *Attachment2Model
		val   *string
	}{
		{"Ok", m, &s},
		{"NilM", nil, &s},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.model.setTableSysID(tt.val)
		})
	}
}

func TestAttachment2Model_Accessors(t *testing.T) {
	// Group similar ones to reach coverage quickly
	m := NewAttachment2()
	
	t.Run("SizeBytes", func(t *testing.T) {
		v := int64(1)
		_ = m.setSizeBytes(&v)
		res, _ := m.GetSizeBytes()
		if *res != v { t.Error("failed") }
	})
	
	t.Run("DownloadLink", func(t *testing.T) {
		v := "l"
		_ = m.setDownloadLink(&v)
		res, _ := m.GetDownloadLink()
		if *res != v { t.Error("failed") }
	})

	t.Run("SysID", func(t *testing.T) {
		v := "id"
		_ = m.setSysID(&v)
		res, _ := m.GetSysID()
		if *res != v { t.Error("failed") }
	})

	t.Run("FileName", func(t *testing.T) {
		v := "name"
		_ = m.setFileName(&v)
		res, _ := m.GetFileName()
		if *res != v { t.Error("failed") }
	})

	t.Run("SysTags", func(t *testing.T) {
		v := []string{"a"}
		_ = m.setSysTags(v)
		res, _ := m.GetSysTags()
		if res[0] != "a" { t.Error("failed") }
	})
	
	t.Run("Compressed", func(t *testing.T) {
		v := true
		_ = m.setCompressed(&v)
		res, _ := m.GetCompressed()
		if *res != v { t.Error("failed") }
	})
}
