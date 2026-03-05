package internal

import (
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
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
	deserializers := res.GetFieldDeserializers()
	if deserializers == nil {
		t.Error("expected non-nil deserializers")
	}
	expectedKeys := []string{resultKey, nextKey, previousKey, firstKey, lastKey}
	for _, key := range expectedKeys {
		if _, ok := deserializers[key]; !ok {
			t.Errorf("expected key %s in deserializers", key)
		}
	}
}

func TestBaseServiceNowCollectionResponse_GetBackingStore(t *testing.T) {
	res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
	res.backingStore = nil
	bs := res.GetBackingStore()
	if bs == nil {
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

func TestBaseServiceNowCollectionResponse_ParseHeaders(t *testing.T) {
	tests := []struct {
		name     string
		headers  func() *abstractions.ResponseHeaders
		expected struct {
			next  string
			prev  string
			first string
			last  string
		}
	}{
		{
			name: "All links with quotes",
			headers: func() *abstractions.ResponseHeaders {
				h := abstractions.NewResponseHeaders()
				h.Add("Link", `<url1>;rel="first", <url2>;rel="prev", <url3>;rel="next", <url4>;rel="last"`)
				return h
			},
			expected: struct {
				next  string
				prev  string
				first string
				last  string
			}{
				next:  "url3",
				prev:  "url2",
				first: "url1",
				last:  "url4",
			},
		},
		{
			name: "Links without quotes",
			headers: func() *abstractions.ResponseHeaders {
				h := abstractions.NewResponseHeaders()
				h.Add("Link", `<url1>;rel=first, <url2>;rel=prev, <url3>;rel=next, <url4>;rel=last`)
				return h
			},
			expected: struct {
				next  string
				prev  string
				first string
				last  string
			}{
				next:  "url3",
				prev:  "url2",
				first: "url1",
				last:  "url4",
			},
		},
		{
			name: "Links with spaces",
			headers: func() *abstractions.ResponseHeaders {
				h := abstractions.NewResponseHeaders()
				h.Add("Link", `<url1> ; rel="first" , <url2> ; rel="prev"`)
				return h
			},
			expected: struct {
				next  string
				prev  string
				first string
				last  string
			}{
				prev:  "url2",
				first: "url1",
			},
		},
		{
			name: "Lowercase link header",
			headers: func() *abstractions.ResponseHeaders {
				h := abstractions.NewResponseHeaders()
				h.Add("link", `<url1>;rel="next"`)
				return h
			},
			expected: struct {
				next  string
				prev  string
				first string
				last  string
			}{
				next: "url1",
			},
		},
		{
			name: "Mixed parameters",
			headers: func() *abstractions.ResponseHeaders {
				h := abstractions.NewResponseHeaders()
				h.Add("Link", `<url1>; foo=bar; rel=next; baz=qux`)
				return h
			},
			expected: struct {
				next  string
				prev  string
				first string
				last  string
			}{
				next: "url1",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := NewBaseServiceNowCollectionResponse[serialization.Parsable](nil)
			res.ParseHeaders(test.headers())

			if test.expected.next != "" {
				l, _ := res.GetNextLink()
				if assert.NotNil(t, l, "NextLink should not be nil") {
					assert.Equal(t, test.expected.next, *l)
				}
			}
			if test.expected.prev != "" {
				l, _ := res.GetPreviousLink()
				if assert.NotNil(t, l, "PrevLink should not be nil") {
					assert.Equal(t, test.expected.prev, *l)
				}
			}
			if test.expected.first != "" {
				l, _ := res.GetFirstLink()
				if assert.NotNil(t, l, "FirstLink should not be nil") {
					assert.Equal(t, test.expected.first, *l)
				}
			}
			if test.expected.last != "" {
				l, _ := res.GetLastLink()
				if assert.NotNil(t, l, "LastLink should not be nil") {
					assert.Equal(t, test.expected.last, *l)
				}
			}
		})
	}
}
