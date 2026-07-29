package core

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestParseHeaders(t *testing.T) {
	tests := []struct {
		name       string
		headers    *abstractions.ResponseHeaders
		collection func() *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]
		verify     func(*testing.T, *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable])
	}{
		{
			name: "single link header sets all rels",
			headers: func() *abstractions.ResponseHeaders {
				h := abstractions.NewResponseHeaders()
				h.Add(linkHeaderKey, `<https://example.com?page=2>; rel="next", <https://example.com?page=1>; rel="previous", <https://example.com?page=0>; rel="first", <https://example.com?page=9>; rel="last"`)
				return h
			}(),
			collection: func() *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable] {
				c := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				c.On("SetNextLink", internal.ToPointer("https://example.com?page=2")).Return(nil)
				c.On("SetPreviousLink", internal.ToPointer("https://example.com?page=1")).Return(nil)
				c.On("SetFirstLink", internal.ToPointer("https://example.com?page=0")).Return(nil)
				c.On("SetLastLink", internal.ToPointer("https://example.com?page=9")).Return(nil)
				return c
			},
			verify: func(t *testing.T, c *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]) {
				c.AssertCalled(t, "SetNextLink", internal.ToPointer("https://example.com?page=2"))
				c.AssertCalled(t, "SetPreviousLink", internal.ToPointer("https://example.com?page=1"))
				c.AssertCalled(t, "SetFirstLink", internal.ToPointer("https://example.com?page=0"))
				c.AssertCalled(t, "SetLastLink", internal.ToPointer("https://example.com?page=9"))
			},
		},
		{
			name: "malformed link header sets nothing",
			headers: func() *abstractions.ResponseHeaders {
				h := abstractions.NewResponseHeaders()
				h.Add(linkHeaderKey, "not-a-valid-link-header")
				return h
			}(),
			collection: func() *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable] {
				return &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			},
			verify: func(t *testing.T, c *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]) {
				c.AssertExpectations(t)
			},
		},
		{
			name:    "no link header",
			headers: abstractions.NewResponseHeaders(),
			collection: func() *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable] {
				return &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			},
			verify: func(t *testing.T, c *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]) {
				c.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection := tt.collection()
			ParseHeaders(collection, tt.headers)
			tt.verify(t, collection)
		})
	}
}

func TestParseHeaders_NilGuards(t *testing.T) {
	t.Run("nil collection", func(t *testing.T) {
		assert.NotPanics(t, func() {
			ParseHeaders(nil, abstractions.NewResponseHeaders())
		})
	})

	t.Run("nil headers", func(t *testing.T) {
		collection := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
		assert.NotPanics(t, func() {
			ParseHeaders(collection, nil)
		})
	})
}
