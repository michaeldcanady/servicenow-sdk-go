package batchapi

import "net/http"

type BatchResponse struct {
	ID       string              `json:"batch_request_id"`
	Requests []BatchResponseItem `json:"serviced_requests"`
}

func (r *BatchResponse) ParseHeaders(http.Header) {}
