package internal

import (
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
		{"NilE", nilE, nil, false},
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
		{"NilE", nilE, me, false},
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

func TestServiceNowError_ErrorBranches(t *testing.T) {
	eWrongType := NewServiceNowError()
	_ = eWrongType.GetBackingStore().Set(errorKey, 123)
	if _, err := eWrongType.GetError(); err == nil || err.Error() != "cannot convert '123' to type internal.MainErrorable" {
		t.Errorf("Expected wrong type error, got %v", err)
	}

	eNilBS := &ServiceNowError{BackedModel: &mockNilBSModel{}}
	if err := eNilBS.setError(nil); err == nil || err.Error() != "backingStore is nil" {
		t.Errorf("Expected BS nil error, got %v", err)
	}
}
