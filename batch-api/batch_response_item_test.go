package batchapi

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	intBatch "github.com/RecoLabs/servicenow-sdk-go/batch-api/internal"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBatchResponseItem_GetBody(t *testing.T) {
	tests := []intBatch.Test[map[string]interface{}]{
		{
			Title:    "Successful",
			Input:    "{\"name\":\"jeff\"}",
			Expected: map[string]interface{}{"name": any("jeff")},
		},
		{
			Title:       "invalid JSON",
			Input:       "invalid JSON",
			Expected:    nil,
			ExpectedErr: &json.SyntaxError{},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			response := &batchResponseItem{
				Body: base64.StdEncoding.EncodeToString([]byte(test.Input.(string))),
			}
			body, err := response.GetBody()
			if test.ExpectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			if body != nil {
				assert.Equal(t, test.Expected, *body)
			} else {
				assert.Nil(t, test.Expected)
			}
		})
	}
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
