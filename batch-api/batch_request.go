package batchapi

import (
	"errors"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	restRequestsKey = "rest_requests"
)

// BatchRequestable represents a Service-Now Batch API request
type BatchRequestable interface {
	GetBatchRequestID() (*string, error)
	SetBatchRequestID(*string) error
	GetRestRequests() ([]RestRequestable, error)
	SetRestRequests([]RestRequestable) error
	AddRequest(RestRequestable) error
	serialization.Parsable
	store.BackedModel
}

// BatchRequest implementation of BatchRequestable
type BatchRequest struct {
	newInternal.Model
}

// NewBatchRequest2 creates a new BatchRequestable.
func NewBatchRequest2() *BatchRequest {
	request := &BatchRequest{
		newInternal.NewBaseModel(),
	}

	id := uuid.NewString()

	_ = request.SetBatchRequestID(&id)

	return request
}

// CreateBatchRequest2FromDiscriminatorValue is a parsable factory for creating a BatchRequestable
func CreateBatchRequest2FromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchRequest2(), nil
}

// Serialize writes the objects properties to the current writer.
func (bR *BatchRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bR) {
		return nil
	}

	serializers := []func(serialization.SerializationWriter) error{
		func(sw serialization.SerializationWriter) error {
			id, err := bR.GetBatchRequestID()
			if err != nil {
				return err
			}

			// ensure request has an id BEFORE serializing
			if internal.IsNil(id) || *id == "" {
				idString := uuid.NewString()
				id = &idString
			}

			return sw.WriteStringValue(batchRequestIDKey, id)
		},
		func(sw serialization.SerializationWriter) error {
			requests, err := bR.GetRestRequests()
			if err != nil {
				return err
			}

			// Create a new slice of serialization.Parsable
			parsableRequests := make([]serialization.Parsable, len(requests))
			for i, header := range requests {
				parsableRequests[i] = header
			}

			return sw.WriteCollectionOfObjectValues(restRequestsKey, parsableRequests)
		},
	}

	for _, serializer := range serializers {
		if err := serializer(writer); err != nil {
			return err
		}
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (bR *BatchRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(bR) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		batchRequestIDKey: func(pn serialization.ParseNode) error {
			id, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return bR.SetBatchRequestID(id)
		},
		restRequestsKey: func(pn serialization.ParseNode) error {
			requests, err := pn.GetCollectionOfObjectValues(CreateRestRequestFromDiscriminatorValue)
			if err != nil {
				return err
			}

			typedRequest := make([]RestRequestable, 0, len(requests))

			for _, request := range requests {
				request, ok := request.(RestRequestable)
				if !ok {
					return errors.New("request is not RestRequestable")
				}
				typedRequest = append(typedRequest, request)
			}

			return bR.SetRestRequests(typedRequest)
		},
	}
}

// AddRequest add request to the slice of requests.
func (bR *BatchRequest) AddRequest(request RestRequestable) error {
	if internal.IsNil(bR) {
		return nil
	}

	requests, err := bR.GetRestRequests()
	if err != nil {
		return err
	}

	requests = append(requests, request)

	return bR.SetRestRequests(requests)
}

// GetBatchRequestID returns the id of the request.
func (bR *BatchRequest) GetBatchRequestID() (*string, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	id, err := bR.GetBackingStore().Get(batchRequestIDKey)
	if err != nil {
		return nil, err
	}

	typedID, ok := id.(*string)
	if !ok {
		return nil, errors.New("id is not *string")
	}

	return typedID, nil
}

// SetBatchRequestID sets the id of the request.
func (bR *BatchRequest) SetBatchRequestID(id *string) error {
	if internal.IsNil(bR) {
		return nil
	}

	return bR.GetBackingStore().Set(batchRequestIDKey, id)
}

// GetRestRequests returns batched requests.
func (bR *BatchRequest) GetRestRequests() ([]RestRequestable, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	requests, err := bR.GetBackingStore().Get(restRequestsKey)
	if err != nil {
		return nil, err
	}

	typedRequests, ok := requests.([]RestRequestable)
	if !ok {
		return nil, errors.New("requests is not []RestRequestable")
	}

	return typedRequests, nil
}

// SetRestRequests sets the batched requests.
func (bR *BatchRequest) SetRestRequests(requests []RestRequestable) error {
	if internal.IsNil(bR) {
		return nil
	}

	return bR.GetBackingStore().Set(restRequestsKey, requests)
}
