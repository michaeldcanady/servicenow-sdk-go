package batchapi

import (
	"encoding/base64"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBatchResponseItem_GetBody(t *testing.T) {
	data := "{\"name\":\"jeff\"}"

	response := &batchResponseItem{
		Body: base64.StdEncoding.EncodeToString([]byte(data)),
	}

	body, err := response.GetBody()

	expected := any(data)

	assert.Nil(t, err)
	assert.Equal(t, &expected, body)
}
func TestBatchResponseItem_GetType(t *testing.T) {
	response := &batchResponseItem{
		Body: "{\"name\":\"jeff\"}",
	}

	assert.Equal(t, reflect.TypeOf(map[string]interface{}{}), response.GetType())
}
func TestBatchResponseItem_GetExecutionTime(t *testing.T) {
	response := &batchResponseItem{
		ExecutionTime: 100,
	}

	assert.Equal(t, time.Duration(100), response.GetExecutionTime())
}
func TestBatchResponseItem_GetHeaders(t *testing.T) {
	response := &batchResponseItem{
		Headers: []batchHeader{
			{
				Name:  "Content-Type",
				Value: "application/json",
			},
		},
	}

	assert.Equal(t, []batchHeader{
		{
			Name:  "Content-Type",
			Value: "application/json",
		},
	}, response.GetHeaders())
}
func TestBatchResponseItem_GetID(t *testing.T) {
	response := &batchResponseItem{}

	id := uuid.New().String()

	response.ID = id

	assert.Equal(t, id, response.GetID())
}
func TestBatchResponseItem_GetRedirectURL(t *testing.T) {
	response := &batchResponseItem{
		RedirectURL: "https://TestURL.com",
	}

	assert.Equal(t, "https://TestURL.com", response.GetRedirectURL())
}
func TestBatchResponseItem_GetStatusCode(t *testing.T) {
	response := &batchResponseItem{
		StatusCode: 400,
	}

	assert.Equal(t, 400, response.GetStatusCode())
}
