package batchapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	batchRequestIDKey     = "batch_request_id"
	servicedRequestsKey   = "serviced_requests"
	unservicedRequestsKey = "unserviced_requests"
)

type BatchResponseable interface {
	GetBatchRequestID() (*string, error)
	GetServicedRequestByID(id string) (ServicedRequestable, error)
	setBatchRequestID(*string) error
	GetServicedRequests() ([]ServicedRequestable, error)
	setServicedRequests([]ServicedRequestable) error
	GetUnservicedRequests() ([]string, error)
	setUnservicedRequests([]string) error
	serialization.Parsable
	store.BackedModel
}

type batchResponse2 struct {
	backingStore store.BackingStore
}

// NewBatchResponse creates a new batch response.
func NewBatchResponse() BatchResponseable {
	return &batchResponse2{
		backingStore: store.NewInMemoryBackingStore(),
	}
}

// CreateBatchResponseFromDiscriminatorValue is a parsable factory for creating a BatchResponseable
func CreateBatchResponseFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchResponse(), nil
}

// GetBackingStore retrieves the backing store for the model.
func (rE *batchResponse2) GetBackingStore() store.BackingStore {
	if internal.IsNil(rE) {
		return nil
	}

	if internal.IsNil(rE.backingStore) {
		rE.backingStore = store.NewInMemoryBackingStore()
	}

	return rE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (rE *batchResponse2) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(rE) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (rE *batchResponse2) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(rE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		batchRequestIDKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (batchRequestIDKey) not implemented")
		},
		servicedRequestsKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (servicedRequestsKey) not implemented")
		},
		unservicedRequestsKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (unservicedRequestsKey) not implemented")
		},
	}
}

func (rE *batchResponse2) GetBatchRequestID() (*string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	id, err := rE.GetBackingStore().Get(batchRequestIDKey)
	if err != nil {
		return nil, err
	}

	strID, ok := id.(*string)
	if !ok {
		return nil, errors.New("id is not *string")
	}

	return strID, nil
}

func (rE *batchResponse2) setBatchRequestID(id *string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(batchRequestIDKey, id)
}

func (rE *batchResponse2) GetServicedRequests() ([]ServicedRequestable, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}
	servicedRequests, err := rE.GetBackingStore().Get(servicedRequestsKey)
	if err != nil {
		return nil, err
	}

	typedServicedRequests, ok := servicedRequests.([]ServicedRequestable)
	if !ok {
		return nil, errors.New("id is not []ServicedRequestable")
	}

	return typedServicedRequests, nil
}

func (rE *batchResponse2) GetServicedRequestByID(id string) (ServicedRequestable, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	requests, err := rE.GetServicedRequests()
	if err != nil {
		return nil, err
	}

	for _, req := range requests {
		reqID, err := req.GetID()
		if err != nil {
			return nil, err
		}
		if *reqID == id {
			return req, nil
		}
	}

	return nil, errors.New("no requests with id")
}

func (rE *batchResponse2) setServicedRequests(requests []ServicedRequestable) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(servicedRequestsKey, requests)
}

func (rE *batchResponse2) GetUnservicedRequests() ([]string, error) {
	if internal.IsNil(rE) {
		return nil, nil
	}

	unservicedRequests, err := rE.GetBackingStore().Get(unservicedRequestsKey)
	if err != nil {
		return nil, err
	}

	typedUnservicedRequests, ok := unservicedRequests.([]string)
	if !ok {
		return nil, errors.New("unservicedRequests is not []string")
	}

	return typedUnservicedRequests, nil
}

func (rE *batchResponse2) setUnservicedRequests(unservicedRequests []string) error {
	if internal.IsNil(rE) {
		return nil
	}

	return rE.GetBackingStore().Set(unservicedRequestsKey, unservicedRequests)
}
