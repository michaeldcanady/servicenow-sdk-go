package batchapi

import (
	"errors"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	restRequestsKey = "rest_requests"
)

type BatchRequestable interface {
	GetBatchRequestID() (*string, error)
	SetBatchRequestID(*string) error
	GetRestRequests() ([]RestRequestable, error)
	SetRestRequests([]RestRequestable) error
	AddRequest(RestRequestable) error
	serialization.Parsable
	store.BackedModel
}

type batchRequest2 struct {
	backingStore store.BackingStore
}

// NewBatchRequest2 creates a new batch request.
func NewBatchRequest2() BatchRequestable {
	request := &batchRequest2{
		backingStore: store.NewInMemoryBackingStore(),
	}

	id := uuid.NewString()

	_ = request.SetBatchRequestID(&id)

	return request
}

// CreateBatchRequest2FromDiscriminatorValue is a parsable factory for creating a BatchRequestable
func CreateBatchRequest2FromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchRequest2(), nil
}

// GetBackingStore retrieves the backing store for the model.
func (rE *batchRequest2) GetBackingStore() store.BackingStore {
	if internal.IsNil(rE) {
		return nil
	}

	if internal.IsNil(rE.backingStore) {
		rE.backingStore = store.NewInMemoryBackingStore()
	}

	return rE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (rE *batchRequest2) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(rE) {
		return nil
	}

	serializers := []func(serialization.SerializationWriter) error{
		func(sw serialization.SerializationWriter) error {
			id, err := rE.GetBatchRequestID()
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
			requests, err := rE.GetRestRequests()
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
func (rE *batchRequest2) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(rE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		batchRequestIDKey: func(pn serialization.ParseNode) error {
			id, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return rE.SetBatchRequestID(id)
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

			return rE.SetRestRequests(typedRequest)
		},
	}
}

// AddRequest add request to the slice of requests.
func (rE *batchRequest2) AddRequest(request RestRequestable) error {
	if internal.IsNil(rE) {
		return nil
	}

	requests, err := rE.GetRestRequests()
	if err != nil {
		return err
	}

	requests = append(requests, request)

	return rE.SetRestRequests(requests)
}

// GetBatchRequestID returns the id of the request.
func (rE *batchRequest2) GetBatchRequestID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	id, err := rE.GetBackingStore().Get(batchRequestIDKey)
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
func (rE *batchRequest2) SetBatchRequestID(id *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(batchRequestIDKey, id)
}

// GetRestRequests returns batched requests.
func (rE *batchRequest2) GetRestRequests() ([]RestRequestable, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	requests, err := rE.GetBackingStore().Get(restRequestsKey)
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
func (rE *batchRequest2) SetRestRequests(requests []RestRequestable) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(restRequestsKey, requests)
}
