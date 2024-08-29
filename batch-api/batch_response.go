package batchapi

import "net/http"

type BatchResponse interface {
	GetID() *string
	GetResponses() []BatchResponseItem[any]
	GetResponse(id string) BatchResponseItem[any]
}

type batchResponse struct {
	ID       *string              `json:"batch_request_id"`
	Requests []*batchResponseItem `json:"serviced_requests"`
}

func (r *batchResponse) ParseHeaders(http.Header) {}

func (r *batchResponse) GetID() *string {
	return r.ID
}

func (r *batchResponse) GetResponses() []BatchResponseItem[any] {
	var ret []BatchResponseItem[any]

	for _, request := range r.Requests {
		ret = append(ret, request)
	}

	return ret
}

func (r *batchResponse) GetResponse(id string) BatchResponseItem[any] {
	for _, response := range r.Requests {
		if response.GetID() == id {
			return response
		}
	}
	return nil
}
