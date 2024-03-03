package batchapi

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBatchResponseItem_GetResponses(t *testing.T) {
	response := &batchResponse{
		Requests: make([]*batchResponseItem, 0),
	}

	response.Requests = append(response.Requests, &batchResponseItem{})

	assert.Len(t, response.GetResponses(), 1)
}

func TestBatchResponseItem_GetResponse(t *testing.T) {
	response := &batchResponse{
		Requests: make([]*batchResponseItem, 0),
	}

	item := &batchResponseItem{
		ID: uuid.New().String(),
	}

	response.Requests = append(response.Requests, item)

	assert.Equal(t, item, response.GetResponse(item.ID))
	assert.Equal(t, nil, response.GetResponse("1"))
}
