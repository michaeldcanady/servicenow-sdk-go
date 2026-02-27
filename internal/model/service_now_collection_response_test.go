package model

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

func TestNewBaseServiceNowCollectionResponse(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	if res == nil {
		t.Error("returned nil")
	}
}

func TestBaseServiceNowCollectionResponse_Serialize(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	err := res.Serialize(nil)
	if err == nil || err.Error() != "Serialize not implemented" {
		t.Errorf("got %v, expected Serialize not implemented", err)
	}
	var nilR *BaseServiceNowCollectionResponse[serialization.Parsable]
	if err := nilR.Serialize(nil); err != nil {
		t.Error("nil receiver should return nil error")
	}
}

func TestBaseServiceNowCollectionResponse_GetFieldDeserializers(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	if res.GetFieldDeserializers() != nil {
		t.Error("expected nil")
	}
}

func TestBaseServiceNowCollectionResponse_GetBackingStore(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	res.backingStore = nil
	bs, err := res.GetBackingStore()
	if err != nil || bs == nil {
		t.Error("failed to get backing store")
	}
}

func TestBaseServiceNowCollectionResponse_GetResult(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[*MainError](nil)
	me := NewMainError()
	_ = res.backingStore.Set(resultKey, []any{me})

	tests := []struct {
		name  string
		model *BaseServiceNowCollectionResponse[*MainError]
		err   bool
	}{
		{"Ok", res, false},
		{"NilM", nil, false},
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

func TestBaseServiceNowCollectionResponse_GetLinks(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	link := "http://link"
	_ = res.backingStore.Set(nextKey, &link)
	_ = res.backingStore.Set(previousKey, &link)
	_ = res.backingStore.Set(firstKey, &link)
	_ = res.backingStore.Set(lastKey, &link)

	t.Run("Next", func(t *testing.T) {
		l, _ := res.GetNextLink()
		if *l != link {
			t.Error("failed")
		}
	})
	t.Run("Previous", func(t *testing.T) {
		l, _ := res.GetPreviousLink()
		if *l != link {
			t.Error("failed")
		}
	})
	t.Run("First", func(t *testing.T) {
		l, _ := res.GetFirstLink()
		if *l != link {
			t.Error("failed")
		}
	})
	t.Run("Last", func(t *testing.T) {
		l, _ := res.GetLastLink()
		if *l != link {
			t.Error("failed")
		}
	})

	var nilR *BaseServiceNowCollectionResponse[serialization.Parsable]
	t.Run("Nil", func(t *testing.T) {
		if l, _ := nilR.GetNextLink(); l != nil {
			t.Error("expected nil")
		}
	})
}

func TestBaseServiceNowCollectionResponse_ErrorBranches(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	_ = res.backingStore.Set(resultKey, "not-a-slice")
	if _, err := res.GetResult(); err == nil || err.Error() != "val is not slice" {
		t.Errorf("expected slice error, got %v", err)
	}

	_ = res.backingStore.Set(nextKey, 123)
	if _, err := res.GetNextLink(); err == nil || err.Error() != "val is not *string" {
		t.Error("expected string error in next link")
	}

	_ = res.backingStore.Set(previousKey, 123)
	if _, err := res.GetPreviousLink(); err == nil || err.Error() != "val is not *string" {
		t.Error("expected string error in prev link")
	}

	_ = res.backingStore.Set(firstKey, 123)
	if _, err := res.GetFirstLink(); err == nil || err.Error() != "val is not *string" {
		t.Error("expected string error in first link")
	}

	_ = res.backingStore.Set(lastKey, 123)
	if _, err := res.GetLastLink(); err == nil || err.Error() != "val is not *string" {
		t.Error("expected string error in last link")
	}
}
