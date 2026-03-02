package internal

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/store"
)

func TestNewMainError(t *testing.T) {
	err := NewMainError()
	if err == nil {
		t.Fatal("NewMainError returned nil")
	}
}

func TestCreateMainErrorFromDiscriminatorValue(t *testing.T) {
	res, err := CreateMainErrorFromDiscriminatorValue(nil)
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}
	if res == nil {
		t.Error("returned nil")
	}
}

func TestMainError_Serialize(t *testing.T) {
	m := NewMainError()
	err := m.Serialize(nil)
	if err == nil || err.Error() != "unsupported" {
		t.Errorf("got err %v, expected unsupported", err)
	}
}

func TestMainError_GetFieldDeserializers(t *testing.T) {
	m := NewMainError()
	deser := m.GetFieldDeserializers()
	if deser[detailKey] == nil || deser[messageKey] == nil || deser[statusKey] == nil {
		t.Error("missing deserializers")
	}
}

func TestMainError_GetDetail(t *testing.T) {
	s := "d"
	m := NewMainError()
	_ = m.setDetail(&s)
	var nilM *MainError

	tests := []struct {
		name     string
		model    *MainError
		expected *string
		err      bool
	}{
		{"Ok", m, &s, false},
		{"NilM", nilM, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.model.GetDetail()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestMainError_setDetail(t *testing.T) {
	s := "d"
	m := NewMainError()
	var nilM *MainError

	tests := []struct {
		name  string
		model *MainError
		val   *string
		err   bool
	}{
		{"Ok", m, &s, false},
		{"NilM", nilM, &s, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.setDetail(tt.val)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestMainError_GetMessage(t *testing.T) {
	s := "m"
	m := NewMainError()
	_ = m.setMessage(&s)
	var nilM *MainError

	tests := []struct {
		name     string
		model    *MainError
		expected *string
		err      bool
	}{
		{"Ok", m, &s, false},
		{"NilM", nilM, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.model.GetMessage()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestMainError_setMessage(t *testing.T) {
	s := "m"
	m := NewMainError()
	var nilM *MainError

	tests := []struct {
		name  string
		model *MainError
		val   *string
		err   bool
	}{
		{"Ok", m, &s, false},
		{"NilM", nilM, &s, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.setMessage(tt.val)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestMainError_GetStatus(t *testing.T) {
	s := "s"
	m := NewMainError()
	_ = m.setStatus(&s)
	var nilM *MainError

	tests := []struct {
		name     string
		model    *MainError
		expected *string
		err      bool
	}{
		{"Ok", m, &s, false},
		{"NilM", nilM, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.model.GetStatus()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestMainError_setStatus(t *testing.T) {
	s := "s"
	m := NewMainError()
	var nilM *MainError

	tests := []struct {
		name  string
		model *MainError
		val   *string
		err   bool
	}{
		{"Ok", m, &s, false},
		{"NilM", nilM, &s, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.setStatus(tt.val)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestMainError_ErrorBranches(t *testing.T) {
	mNilBS := &MainError{Model: &mockNilBSModel{}}

	if _, err := mNilBS.GetDetail(); err == nil || err.Error() != "backingStore is nil" {
		t.Errorf("Expected backingStore is nil error, got %v", err)
	}
	if err := mNilBS.setDetail(nil); err == nil || err.Error() != "backingStore is nil" {
		t.Error("Expected BS nil error in setDetail")
	}
	if _, err := mNilBS.GetMessage(); err == nil || err.Error() != "backingStore is nil" {
		t.Error("Expected BS nil error in GetMessage")
	}
	if err := mNilBS.setMessage(nil); err == nil || err.Error() != "backingStore is nil" {
		t.Error("Expected BS nil error in setMessage")
	}
	if _, err := mNilBS.GetStatus(); err == nil || err.Error() != "backingStore is nil" {
		t.Error("Expected BS nil error in GetStatus")
	}
	if err := mNilBS.setStatus(nil); err == nil || err.Error() != "backingStore is nil" {
		t.Error("Expected BS nil error in setStatus")
	}

	mWrongType := NewMainError()
	_ = mWrongType.GetBackingStore().Set(detailKey, 123)
	if _, err := mWrongType.GetDetail(); err == nil || err.Error() != "detail is not *string" {
		t.Errorf("Expected wrong type error, got %v", err)
	}

	_ = mWrongType.GetBackingStore().Set(messageKey, 123)
	if _, err := mWrongType.GetMessage(); err == nil || err.Error() != "message is not *string" {
		t.Errorf("Expected wrong type error in Message, got %v", err)
	}

	_ = mWrongType.GetBackingStore().Set(statusKey, 123)
	if _, err := mWrongType.GetStatus(); err == nil || err.Error() != "status is not *string" {
		t.Errorf("Expected wrong type error in Status, got %v", err)
	}
}

type mockNilBSModel struct{ BaseModel }

func (m *mockNilBSModel) GetBackingStore() store.BackingStore { return nil }
