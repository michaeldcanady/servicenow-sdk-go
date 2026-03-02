package attachmentapi

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/store"
)

func TestNewFileWithContent(t *testing.T) {
	res := NewFileWithContent()
	if res == nil {
		t.Error("returned nil")
	}
}

func TestCreateFileWithContentFromDiscriminatorValue(t *testing.T) {
	res, err := CreateFileWithContentFromDiscriminatorValue(nil)
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}
	if res == nil {
		t.Error("returned nil")
	}
}

func TestFileWithContentModel_GetFieldDeserializers(t *testing.T) {
	m := NewFileWithContent()
	deser := m.GetFieldDeserializers()
	if deser == nil {
		t.Error("expected non-nil deser")
	}
	var nilM *FileWithContentModel
	if nilM.GetFieldDeserializers() != nil {
		t.Error("expected nil")
	}
}

func TestFileWithContentModel_GetContent(t *testing.T) {
	m := NewFileWithContent()
	data := []byte("test")
	_ = m.SetContent(data)
	var nilM *FileWithContentModel

	tests := []struct {
		name     string
		model    FileWithContent
		expected []byte
		err      bool
	}{
		{"Ok", m, data, false},
		{"NilM", nilM, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.model.GetContent()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if string(res) != string(tt.expected) {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestFileWithContentModel_SetContent(t *testing.T) {
	m := NewFileWithContent()
	data := []byte("test")
	var nilM *FileWithContentModel

	tests := []struct {
		name  string
		model FileWithContent
		err   bool
	}{
		{"Ok", m, false},
		{"NilM", nilM, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.SetContent(data)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestFileWithContentModel_ErrorBranches(t *testing.T) {
	mWrongType := NewFileWithContent()
	_ = mWrongType.GetBackingStore().Set(contentKey, 123)
	if _, err := mWrongType.GetContent(); err == nil || err.Error() != "content is not []byte" {
		t.Errorf("expected type error, got %v", err)
	}

	mNilBS := &FileWithContentModel{File: &mockNilBSFile{}}
	if _, err := mNilBS.GetContent(); err == nil || err.Error() != "store is nil" {
		t.Errorf("expected BS nil error in Get, got %v", err)
	}
	if err := mNilBS.SetContent(nil); err == nil || err.Error() != "store is nil" {
		t.Errorf("expected BS nil error in Set, got %v", err)
	}
}

type mockNilBSFile struct{ FileModel }

func (m *mockNilBSFile) GetBackingStore() store.BackingStore { return nil }
