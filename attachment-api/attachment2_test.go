package attachmentapi

import (
	"errors"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
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
	assert.NotNil(t, deser[tableSysIDKey])

	for key, fn := range deser {
		node := mocking.NewMockParseNode()
		s := "test"
		switch key {
		case sizeBytesKey, sysModCountKey, sizeCompressedKey, chunkSizeBytesKey:
			s = "1"
		case compressedKey:
			s = "true"
		case sysUpdatedOnKey, sysCreatedOnKey:
			s = "2006-01-02 15:04:05"
		case imageHeightKey, imageWidthKey:
			s = "1.2"
		}

		node.On("GetStringValue").Return(&s, nil)
		_ = fn(node)
	}

	// Test error branches
	for _, fn := range deser {
		node := mocking.NewMockParseNode()
		node.On("GetStringValue").Return((*string)(nil), errors.New("read error"))
		_ = fn(node)
	}
}

func TestAttachment2Model_Serialize(t *testing.T) {
	m := NewAttachment2()
	err := m.Serialize(nil)
	assert.NoError(t, err)

	var nilM *Attachment2Model
	err = nilM.Serialize(nil)
	assert.NoError(t, err)
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

func TestCreateAttachment2FromDiscriminatorValue(t *testing.T) {
	res, err := CreateAttachment2FromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestAttachment2Model_BasicAccessors(t *testing.T) {
	m := NewAttachment2()

	t.Run("SizeBytes", func(t *testing.T) {
		v := int64(1)
		_ = m.setSizeBytes(&v)
		res, _ := m.GetSizeBytes()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("DownloadLink", func(t *testing.T) {
		v := "l"
		_ = m.setDownloadLink(&v)
		res, _ := m.GetDownloadLink()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("SysID", func(t *testing.T) {
		v := "id"
		_ = m.setSysID(&v)
		res, _ := m.GetSysID()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("FileName", func(t *testing.T) {
		v := "name"
		_ = m.setFileName(&v)
		res, _ := m.GetFileName()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("SysTags", func(t *testing.T) {
		v := []string{"a"}
		_ = m.setSysTags(v)
		res, _ := m.GetSysTags()
		if res[0] != "a" {
			t.Error("failed")
		}
	})

	t.Run("Compressed", func(t *testing.T) {
		v := true
		_ = m.setCompressed(&v)
		res, _ := m.GetCompressed()
		if *res != v {
			t.Error("failed")
		}
	})
}

func TestAttachment2Model_AuditAccessors(t *testing.T) {
	m := NewAttachment2()

	t.Run("SysUpdatedOn", func(t *testing.T) {
		v := time.Now()
		_ = m.setSysUpdatedOn(&v)
		res, _ := m.GetSysUpdatedOn()
		if !res.Equal(v) {
			t.Error("failed")
		}
	})

	t.Run("SysCreatedOn", func(t *testing.T) {
		v := time.Now()
		_ = m.setSysCreatedOn(&v)
		res, _ := m.GetSysCreatedOn()
		if !res.Equal(v) {
			t.Error("failed")
		}
	})

	t.Run("SysCreatedBy", func(t *testing.T) {
		v := "user"
		_ = m.setSysCreatedBy(&v)
		res, _ := m.GetSysCreatedBy()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("SysUpdatedBy", func(t *testing.T) {
		v := "user2"
		_ = m.setSysUpdatedBy(&v)
		res, _ := m.GetSysUpdatedBy()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("SysModCount", func(t *testing.T) {
		v := int64(5)
		_ = m.setSysModCount(&v)
		res, _ := m.GetSysModCount()
		if *res != v {
			t.Error("failed")
		}
	})
}

func TestAttachment2Model_ImageAccessors(t *testing.T) {
	m := NewAttachment2()

	t.Run("ImageHeight", func(t *testing.T) {
		v := float64(100)
		_ = m.setImageHeight(&v)
		res, _ := m.GetImageHeight()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("AverageImageColor", func(t *testing.T) {
		v := "red"
		_ = m.setAverageImageColor(&v)
		res, _ := m.GetAverageImageColor()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("ImageWidth", func(t *testing.T) {
		v := float64(200)
		_ = m.setImageWidth(&v)
		res, _ := m.GetImageWidth()
		if *res != v {
			t.Error("failed")
		}
	})
}

func TestAttachment2Model_MetadataAccessors(t *testing.T) {
	m := NewAttachment2()

	t.Run("TableName", func(t *testing.T) {
		v := "table"
		_ = m.setTableName(&v)
		res, _ := m.GetTableName()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("ContentType", func(t *testing.T) {
		v := "image/png"
		_ = m.setContentType(&v)
		res, _ := m.GetContentType()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("SizeCompressed", func(t *testing.T) {
		v := int64(50)
		_ = m.setSizeCompressed(&v)
		res, _ := m.GetSizeCompressed()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("ChunkSizeBytes", func(t *testing.T) {
		v := "1024"
		_ = m.setChunkSizeBytes(&v)
		res, _ := m.GetChunkSizeBytes()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("Hash", func(t *testing.T) {
		v := "hash"
		_ = m.setHash(&v)
		res, _ := m.GetHash()
		if *res != v {
			t.Error("failed")
		}
	})

	t.Run("State", func(t *testing.T) {
		v := "available"
		_ = m.setState(&v)
		res, _ := m.GetState()
		if *res != v {
			t.Error("failed")
		}
	})
}
