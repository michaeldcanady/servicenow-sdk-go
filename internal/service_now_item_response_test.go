package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBaseServiceNowItemResponse(t *testing.T) {
	res := NewBaseServiceNowItemResponse[serialization.Parsable](nil)
	if res == nil {
		t.Error("returned nil")
	}
}

func TestBaseServiceNowItemResponse_Serialize(t *testing.T) {
	res := NewBaseServiceNowItemResponse[serialization.Parsable](nil)
	err := res.Serialize(nil)
	assert.NoError(t, err)

	var nilR *BaseServiceNowItemResponse[serialization.Parsable]
	if err := nilR.Serialize(nil); err != nil {
		t.Error("nil receiver should return nil error")
	}
}

func TestBaseServiceNowItemResponse_GetFieldDeserializers(t *testing.T) {
	res := NewBaseServiceNowItemResponse[*MainError](CreateMainErrorFromDiscriminatorValue)
	deser := res.GetFieldDeserializers()
	if deser[resultKey] == nil {
		t.Error("missing deserializer")
	}
}

func TestBaseServiceNowItemResponse_GetBackingStore(t *testing.T) {
	res := NewBaseServiceNowItemResponse[serialization.Parsable](nil)
	res.backingStore = nil
	bs := res.GetBackingStore()
	if bs == nil {
		t.Error("failed to get backing store")
	}
	var nilR *BaseServiceNowItemResponse[serialization.Parsable]
	if bs := nilR.GetBackingStore(); bs != nil {
		t.Error("expected nil")
	}
}

func TestBaseServiceNowItemResponse_GetResult(t *testing.T) {
	res := NewBaseServiceNowItemResponse[*MainError](nil)
	me := NewMainError()
	_ = res.setResult(me)
	var nilR *BaseServiceNowItemResponse[*MainError]

	tests := []struct {
		name  string
		model *BaseServiceNowItemResponse[*MainError]
		err   bool
	}{
		{"Ok", res, false},
		{"NilM", nilR, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.model.GetResult()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestBaseServiceNowItemResponse_setResult(t *testing.T) {
	res := NewBaseServiceNowItemResponse[*MainError](nil)
	me := NewMainError()
	var nilR *BaseServiceNowItemResponse[*MainError]

	tests := []struct {
		name  string
		model *BaseServiceNowItemResponse[*MainError]
		val   *MainError
		err   bool
	}{
		{"Ok", res, me, false},
		{"NilM", nilR, me, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.setResult(tt.val)
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestBaseServiceNowItemResponse_ErrorBranches(t *testing.T) {
	res := NewBaseServiceNowItemResponse[*MainError](nil)
	_ = res.backingStore.Set(resultKey, "not-parsable")
	if _, err := res.GetResult(); err == nil {
		t.Error("expected error for wrong type in GetResult")
	}

	resNilF := NewBaseServiceNowItemResponse[*MainError](nil)
	deser := resNilF.GetFieldDeserializers()
	node := mocking.NewMockParseNode()
	node.On("GetObjectValue", mock.Anything).Return(nil, errors.New("factory is nil"))
	if err := deser[resultKey](node); err == nil || err.Error() != "factory is nil" {
		t.Errorf("expected factory nil error, got %v", err)
	}

	resNilBS := &BaseServiceNowItemResponse[*MainError]{backingStore: nil, backingStoreFactory: nil}
	if err := resNilBS.setResult(nil); err == nil || err.Error() != "backingStore is nil" {
		t.Errorf("expected store nil error, got %v", err)
	}
}
