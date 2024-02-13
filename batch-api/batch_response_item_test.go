package batchapi

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBatchResponseItem_GetBody(t *testing.T) {

}
func TestBatchResponseItem_GetType(t *testing.T) {

}
func TestBatchResponseItem_GetExecutionTime(t *testing.T) {

}
func TestBatchResponseItem_GetHeaders(t *testing.T) {

}
func TestBatchResponseItem_GetID(t *testing.T) {
	response := &batchResponse{}

	id := uuid.New().String()

	response.ID = &id

	assert.Equal(t, &id, response.GetID())
}
func TestBatchResponseItem_GetRedirectURL(t *testing.T) {

}
func TestBatchResponseItem_GetStatusCode(t *testing.T) {

}
