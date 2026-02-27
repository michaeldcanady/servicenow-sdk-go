package model

import (
	"testing"
)

func TestNewServicenowError(t *testing.T) {
	err := NewServicenowError()
	if err == nil {
		t.Fatal("NewServicenowError returned nil")
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

func TestServicenowError_Serialize(t *testing.T) {
	err := NewServicenowError().Serialize(nil)
	if err == nil || err.Error() != "unsupported" {
		t.Errorf("got err %v, expected unsupported", err)
	}
}

func TestServicenowError_GetFieldDeserializers(t *testing.T) {
	deser := NewServicenowError().GetFieldDeserializers()
	if deser[errorKey] == nil {
		t.Error("missing deserializer")
	}
}

func TestServicenowError_GetError(t *testing.T) {
	me := NewMainError()
	e := NewServicenowError()
	_ = e.setError(me)
	var nilE *ServicenowError

	tests := []struct {
		name     string
		model    *ServicenowError
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

func TestServicenowError_setError(t *testing.T) {
	me := NewMainError()
	e := NewServicenowError()
	var nilE *ServicenowError

	tests := []struct {
		name  string
		model *ServicenowError
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

func TestServicenowError_ErrorBranches(t *testing.T) {
	eWrongType := NewServicenowError()
	_ = eWrongType.GetBackingStore().Set(errorKey, 123)
	if _, err := eWrongType.GetError(); err == nil || err.Error() != "rawMainErr is not MainErrorable" {
		t.Errorf("Expected wrong type error, got %v", err)
	}

	eNilBS := &ServicenowError{Model: &mockNilBSModel{}}
	if err := eNilBS.setError(nil); err == nil || err.Error() != "backingStore is nil" {
		t.Errorf("Expected BS nil error, got %v", err)
	}
}
