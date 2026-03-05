package internal

import (
	"errors"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseModel(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Initialization",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model := NewBaseModel()
			assert.NotNil(t, model)
			assert.NotNil(t, model.backingStoreFactory)
		})
	}
}

func TestBaseModel_SetBackingStoreFactory(t *testing.T) {
	var nilM *BaseModel
	m := &BaseModel{}
	f := store.NewInMemoryBackingStore

	tests := []struct {
		name    string
		model   *BaseModel
		factory store.BackingStoreFactory
		err     error
	}{
		{"Ok", m, f, nil},
		{"NilFactory", m, nil, errors.New("factory is nil")},
		{"NilModel", nilM, f, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.SetBackingStoreFactory(tt.factory)
			if tt.err != nil {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("got err %v, expected err %v", err, tt.err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected err %v", err)
				}
			}
		})
	}
}

func TestBaseModel_GetBackingStore(t *testing.T) {
	var nilM *BaseModel
	m := NewBaseModel()
	bs := store.NewInMemoryBackingStore()
	mWithBS := &BaseModel{backingStore: bs}

	tests := []struct {
		name  string
		model *BaseModel
		isNil bool
	}{
		{"FromFactory", m, false},
		{"Existing", mWithBS, false},
		{"NilModel", nilM, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.model.GetBackingStore()
			if tt.isNil {
				if res != nil {
					t.Error("expected nil")
				}
			} else {
				if res == nil {
					t.Error("expected non-nil")
				}
			}
		})
	}
}
