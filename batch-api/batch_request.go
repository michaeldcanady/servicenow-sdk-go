package batchapi

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type BatchRequest interface {
	AddRequest(internal.RequestInformation, bool) error
}

type batchRequest struct {
	ID       string      `json:"batch_request_id"`
	Requests []BatchItem `json:"rest_requests"`
	client   core.Client2
}

func NewBatchRequest(client core.Client2) BatchRequest {
	return &batchRequest{
		ID:       uuid.New().String(),
		Requests: make([]BatchItem, 0),
		client:   client,
	}
}

func (r *batchRequest) AddRequest(requestInfo internal.RequestInformation, excludeResponseHeaders bool) error {
	item, err := r.toBatchItem(requestInfo, excludeResponseHeaders)
	if err != nil {
		return err
	}

	r.Requests = append(r.Requests, item)
	return nil
}

func (r *batchRequest) toBatchItem(requestInfo internal.RequestInformation, excludeResponseHeaders bool) (BatchItem, error) {
	// make sure base url is set accordingly

	uri, err := requestInfo.Url()
	if err != nil {
		return nil, err
	}

	var body internal.RequestBody
	if requestInfo.GetContent() != nil {
		err = json.Unmarshal(requestInfo.GetContent(), &body)
		if err != nil {
			return nil, err
		}
	}

	newID := uuid.NewString()
	method := requestInfo.GetMethod()

	request := NewBatchItem(excludeResponseHeaders)
	request.SetID(&newID)
	request.SetMethod(&method)
	request.SetBody(body)
	request.SetHeaders(requestInfo.GetHeaders())

	baseURI, err := getBaseURL(r.client)
	if err != nil {
		return nil, err
	}

	var finalURL = strings.Replace(uri, baseURI.String(), "", 1)
	// /api is required and as of 1.0, it is appended to the client base path.
	if !strings.HasPrefix(finalURL, "/api") {
		finalURL = "/api" + finalURL
	}
	request.SetURL(&finalURL)

	return request, nil
}
