package internal

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewMainError(t *testing.T) {
	err := NewMainError()
	assert.NotNil(t, err)
}

func TestCreateMainErrorFromDiscriminatorValue(t *testing.T) {
	res, err := CreateMainErrorFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestMainError_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		model   *MainError
		wantErr bool
	}{
		{
			name:    "Successful",
			model:   NewMainError(),
			wantErr: false,
		},
		{
			name:    "nil model",
			model:   nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			err := tt.model.Serialize(writer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMainError_GetFieldDeserializers(t *testing.T) {
	m := NewMainError()
	deser := m.GetFieldDeserializers()
	assert.NotNil(t, deser[detailKey])
	assert.NotNil(t, deser[messageKey])
	assert.NotNil(t, deser[statusKey])
}

func TestMainError_Accessors(t *testing.T) {
	val := "test-value"
	tests := []struct {
		name     string
		model    *MainError
		setter   func(*MainError, *string) error
		getter   func(*MainError) (*string, error)
		expected *string
		wantErr  bool
	}{
		{
			name:     "Detail Success",
			model:    NewMainError(),
			setter:   func(m *MainError, v *string) error { return m.setDetail(v) },
			getter:   func(m *MainError) (*string, error) { return m.GetDetail() },
			expected: &val,
			wantErr:  false,
		},
		{
			name:     "Detail Nil Model",
			model:    nil,
			setter:   func(m *MainError, v *string) error { return m.setDetail(v) },
			getter:   func(m *MainError) (*string, error) { return m.GetDetail() },
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Message Success",
			model:    NewMainError(),
			setter:   func(m *MainError, v *string) error { return m.setMessage(v) },
			getter:   func(m *MainError) (*string, error) { return m.GetMessage() },
			expected: &val,
			wantErr:  false,
		},
		{
			name:     "Message Nil Model",
			model:    nil,
			setter:   func(m *MainError, v *string) error { return m.setMessage(v) },
			getter:   func(m *MainError) (*string, error) { return m.GetMessage() },
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Status Success",
			model:    NewMainError(),
			setter:   func(m *MainError, v *string) error { return m.setStatus(v) },
			getter:   func(m *MainError) (*string, error) { return m.GetStatus() },
			expected: &val,
			wantErr:  false,
		},
		{
			name:     "Status Nil Model",
			model:    nil,
			setter:   func(m *MainError, v *string) error { return m.setStatus(v) },
			getter:   func(m *MainError) (*string, error) { return m.GetStatus() },
			expected: nil,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setter != nil && tt.model != nil {
				err := tt.setter(tt.model, tt.expected)
				assert.NoError(t, err)
			}
			res, err := tt.getter(tt.model)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, res)
			}
		})
	}
}

func TestMainError_ErrorBranches(t *testing.T) {
	mNilBS := &MainError{Model: &mockNilBSModel{}}

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{"GetDetail Nil BS", func() error { _, err := mNilBS.GetDetail(); return err }, "backingStore is nil"},
		{"setDetail Nil BS", func() error { return mNilBS.setDetail(nil) }, "backingStore is nil"},
		{"GetMessage Nil BS", func() error { _, err := mNilBS.GetMessage(); return err }, "backingStore is nil"},
		{"setMessage Nil BS", func() error { return mNilBS.setMessage(nil) }, "backingStore is nil"},
		{"GetStatus Nil BS", func() error { _, err := mNilBS.GetStatus(); return err }, "backingStore is nil"},
		{"setStatus Nil BS", func() error { return mNilBS.setStatus(nil) }, "backingStore is nil"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			assert.Error(t, err)
			assert.Equal(t, tt.wantErr, err.Error())
		})
	}

	mWrongType := NewMainError()
	_ = mWrongType.GetBackingStore().Set(detailKey, 123)
	_, err := mWrongType.GetDetail()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot convert '123' to type *string")
}
