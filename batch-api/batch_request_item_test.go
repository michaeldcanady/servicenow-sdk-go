package batchapi

import (
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestBatchItem_GetID(t *testing.T) {
	// Create a batchItem instance
	item := NewBatchItem(false)
	id := "testID"
	item.SetID(&id)

	assert.Equal(t, &id, item.GetID())
}

func TestBatchItem_GetMethod(t *testing.T) {
	// Create a batchItem instance
	item := NewBatchItem(false)
	id := "GET"
	item.SetMethod(&id)

	assert.Equal(t, &id, item.GetMethod())
}

func TestBatchItem_GetURL(t *testing.T) {
	// Create a batchItem instance
	item := NewBatchItem(false)
	id := "GET"
	item.SetURL(&id)

	assert.Equal(t, &id, item.GetURL())
}

func TestBatchItem_GetHeaders(t *testing.T) {
	// Create a batchItem instance
	item := NewBatchItem(false)
	header := internal.NewRequestHeader()
	header.Set(internal.ContentTypeHeader, internal.JSONContentType)

	expectedHeader := header
	expectedHeader.Set(internal.AcceptHeader, internal.JSONContentType)

	item.SetHeaders(header)

	assert.Equal(t, expectedHeader, item.GetHeaders())
}

func TestBatchItem_SetHeaders(t *testing.T) {
	item := NewBatchItem(false)
	header := internal.NewRequestHeader()
	header.Set(internal.ContentTypeHeader, "")
	item.SetHeaders(header)

	expectedHeader := header
	expectedHeader.Set(internal.AcceptHeader, internal.JSONContentType)

	assert.Equal(t, expectedHeader, item.GetHeaders())
}

func TestBatchItem_GetBody(t *testing.T) {
	// Create a batchItem instance
	item := NewBatchItem(false)

	body := map[string]interface{}{}
	item.SetBody(body)

	assert.Equal(t, internal.RequestBody(body), item.GetBody())
}

func TestBatchItem_GetExcludeResponseHeaders(t *testing.T) {
	// Create a batchItem instance
	item := NewBatchItem(false)

	item.SetExcludeResponseHeaders(true)

	assert.Equal(t, true, item.GetExcludeResponseHeaders())
}
