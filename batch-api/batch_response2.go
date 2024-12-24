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

// BatchResponseable representation of Service-Now Batch API response
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

// batchResponse2 implementation of BatchResponseable
type batchResponse2 struct {
	// backingStoreFactory factory to create backingStore
	backingStoreFactory store.BackingStoreFactory
	// backingStore the store backing the model
	backingStore store.BackingStore
}

// NewBatchResponse creates a new batch response
func NewBatchResponse() BatchResponseable {
	return &batchResponse2{
		backingStore:        store.NewInMemoryBackingStore(),
		backingStoreFactory: store.NewInMemoryBackingStore,
	}
}

// CreateBatchResponseFromDiscriminatorValue is a parsable factory for creating a BatchResponseable
func CreateBatchResponseFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchResponse(), nil
}

// GetBackingStore retrieves the backing store for the model
func (bR *batchResponse2) GetBackingStore() store.BackingStore {
	if internal.IsNil(bR) {
		return nil
	}

	if internal.IsNil(bR.backingStore) {
		bR.backingStore = bR.backingStoreFactory()
	}

	return bR.backingStore
}

// Serialize writes the objects properties to the current writer
func (bR *batchResponse2) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bR) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object
func (bR *batchResponse2) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(bR) {
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

// GetBatchRequestID returns the id of the associated batch request
func (bR *batchResponse2) GetBatchRequestID() (*string, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	id, err := bR.GetBackingStore().Get(batchRequestIDKey)
	if err != nil {
		return nil, err
	}

	strID, ok := id.(*string)
	if !ok {
		return nil, errors.New("id is not *string")
	}

	return strID, nil
}

// setBatchRequestID sets the id of the associated batch request
func (bR *batchResponse2) setBatchRequestID(id *string) error {
	if internal.IsNil(bR) {
		return nil
	}

	return bR.GetBackingStore().Set(batchRequestIDKey, id)
}

// GetServicedRequests returns serviced requests
func (bR *batchResponse2) GetServicedRequests() ([]ServicedRequestable, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}
	servicedRequests, err := bR.GetBackingStore().Get(servicedRequestsKey)
	if err != nil {
		return nil, err
	}

	typedServicedRequests, ok := servicedRequests.([]ServicedRequestable)
	if !ok {
		return nil, errors.New("id is not []ServicedRequestable")
	}

	return typedServicedRequests, nil
}

// GetServicedRequestByID returns the serviced request with the provided id
func (bR *batchResponse2) GetServicedRequestByID(id string) (ServicedRequestable, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	requests, err := bR.GetServicedRequests()
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

// setServicedRequests sets the serviced requests to the provided values
func (bR *batchResponse2) setServicedRequests(requests []ServicedRequestable) error {
	if internal.IsNil(bR) {
		return nil
	}

	return bR.GetBackingStore().Set(servicedRequestsKey, requests)
}

// GetUnservicedRequests returns the unserviced requests' id
func (bR *batchResponse2) GetUnservicedRequests() ([]string, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	unservicedRequests, err := bR.GetBackingStore().Get(unservicedRequestsKey)
	if err != nil {
		return nil, err
	}

	typedUnservicedRequests, ok := unservicedRequests.([]string)
	if !ok {
		return nil, errors.New("unservicedRequests is not []string")
	}

	return typedUnservicedRequests, nil
}

// setUnservicedRequests sets the ids of the unserviced requests to the provided value
func (bR *batchResponse2) setUnservicedRequests(unservicedRequests []string) error {
	if internal.IsNil(bR) {
		return nil
	}

	return bR.GetBackingStore().Set(unservicedRequestsKey, unservicedRequests)
}
