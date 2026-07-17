package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServiceNowError(t *testing.T) {
	err := NewServiceNowError()
	if err == nil {
		t.Fatal("NewServiceNowError returned nil")
	}
}

func TestCreateServiceNowErrorFromDiscriminatorValue(t *testing.T) {
	res, err := CreateServiceNowErrorFromDiscriminatorValue(nil)
	if err != nil {
		t.Errorf("unexpected err %v", err)
	}
	if res == nil {
		t.Error("returned nil")
	}
}

func TestServiceNowError_Serialize(t *testing.T) {
	err := NewServiceNowError().Serialize(nil)
	assert.NoError(t, err)

	var nilE *ServiceNowError
	err = nilE.Serialize(nil)
	assert.NoError(t, err)
}

func TestServiceNowError_GetFieldDeserializers(t *testing.T) {
	deser := NewServiceNowError().GetFieldDeserializers()
	if deser[errorKey] == nil {
		t.Error("missing deserializer")
	}
}

func TestServiceNowError_GetError(t *testing.T) {
	me := NewMainError()
	e := NewServiceNowError()
	_ = e.setError(me)
	var nilE *ServiceNowError

	tests := []struct {
		name     string
		model    *ServiceNowError
		expected MainErrorable
		err      bool
	}{
		{"Ok", e, me, false},
		{"NilE", nilE, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.model.GetError()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
		})
	}
}

func TestServiceNowError_setError(t *testing.T) {
	me := NewMainError()
	e := NewServiceNowError()
	var nilE *ServiceNowError

	tests := []struct {
		name  string
		model *ServiceNowError
		val   MainErrorable
		err   bool
	}{
		{"Ok", e, me, false},
		{"NilE", nilE, me, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.setError(tt.val)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

// TODO: improve test table design
func TestServiceNowError_ErrorBranches(t *testing.T) {
	eWrongType := NewServiceNowError()
	assert.Nil(t, eWrongType.GetBackingStore().Set(errorKey, 123))
	val, err := eWrongType.GetError()
	assert.Nil(t, val)
	assert.Equal(t, errors.New("cannot convert '123' to type core.MainErrorable"), err)

	eNilBS := &ServiceNowError{BackedModel: &mockNilBSModel{}}
	val, err = eNilBS.GetError()
	assert.Nil(t, val)
	assert.Equal(t, errors.New("store is nil"), err)
}
