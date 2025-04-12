package batchapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
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

// BatchResponse implementation of BatchResponseable
type BatchResponse struct {
	newInternal.Model
}

// NewBatchResponse creates a new batch response
func NewBatchResponse() *BatchResponse {
	return &BatchResponse{
		newInternal.NewBaseModel(),
	}
}

// CreateBatchResponseFromDiscriminatorValue is a parsable factory for creating a BatchResponseable
func CreateBatchResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchResponse(), nil
}

// Serialize writes the objects properties to the current writer
func (bR *BatchResponse) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bR) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object
func (bR *BatchResponse) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
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
func (bR *BatchResponse) GetBatchRequestID() (*string, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	backingStore := bR.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, nil
	}

	id, err := backingStore.Get(batchRequestIDKey)
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
func (bR *BatchResponse) setBatchRequestID(id *string) error {
	if internal.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil
	}

	return backingStore.Set(batchRequestIDKey, id)
}

// GetServicedRequests returns serviced requests
func (bR *BatchResponse) GetServicedRequests() ([]ServicedRequestable, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	backingStore := bR.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, nil
	}

	servicedRequests, err := backingStore.Get(servicedRequestsKey)
	if err != nil {
		return nil, err
	}

	typedServicedRequests, ok := servicedRequests.([]ServicedRequestable)
	if !ok {
		return nil, errors.New("servicedRequests is not []ServicedRequestable")
	}

	return typedServicedRequests, nil
}

// setServicedRequests sets the serviced requests to the provided values
func (bR *BatchResponse) setServicedRequests(requests []ServicedRequestable) error {
	if internal.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil
	}

	return backingStore.Set(servicedRequestsKey, requests)
}

// GetUnservicedRequests returns the unserviced requests' id
func (bR *BatchResponse) GetUnservicedRequests() ([]string, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	backingStore := bR.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, nil
	}

	unservicedRequests, err := backingStore.Get(unservicedRequestsKey)
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
func (bR *BatchResponse) setUnservicedRequests(unservicedRequests []string) error {
	if internal.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil
	}

	return backingStore.Set(unservicedRequestsKey, unservicedRequests)
}
