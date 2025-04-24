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

// BatchResponse representation of Service-Now Batch API response
type BatchResponse interface {
	GetBatchRequestID() (*string, error)
	GetServicedRequestByID(id string) (ServicedRequest, error)
	setBatchRequestID(*string) error
	GetServicedRequests() ([]ServicedRequest, error)
	setServicedRequests([]ServicedRequest) error
	GetUnservicedRequests() ([]string, error)
	setUnservicedRequests([]string) error
	serialization.Parsable
	store.BackedModel
}

// BatchResponseModel implementation of BatchResponse
type BatchResponseModel struct {
	newInternal.Model
}

// NewBatchResponse creates a new batch response
func NewBatchResponse() *BatchResponseModel {
	return &BatchResponseModel{
		newInternal.NewBaseModel(),
	}
}

// CreateBatchResponseFromDiscriminatorValue is a parsable factory for creating a BatchResponse
func CreateBatchResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchResponse(), nil
}

// Serialize writes the objects properties to the current writer
func (bR *BatchResponseModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bR) {
		return nil
	}

	return errors.New("Serialize not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object
func (bR *BatchResponseModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(bR) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		batchRequestIDKey: func(pn serialization.ParseNode) error {
			if internal.IsNil(pn) {
				return nil
			}

			value, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return bR.setBatchRequestID(value)
		},
		servicedRequestsKey: func(pn serialization.ParseNode) error {
			if internal.IsNil(pn) {
				return nil
			}

			values, err := pn.GetCollectionOfObjectValues(CreateServicedRequestFromDiscriminatorValue)
			if err != nil {
				return err
			}

			requests := make([]ServicedRequest, 0, len(values))
			for index, value := range values {
				typedValue, ok := value.(ServicedRequest)
				if !ok {
					return errors.New("value is not ServicedRequest")
				}
				requests[index] = typedValue
			}

			return bR.setServicedRequests(requests)
		},
		unservicedRequestsKey: func(pn serialization.ParseNode) error {
			if internal.IsNil(pn) {
				return nil
			}

			values, err := pn.GetCollectionOfPrimitiveValues("string")
			if err != nil {
				return err
			}

			requests := make([]string, 0, len(values))
			for index, value := range values {
				typedValue, ok := value.(string)
				if !ok {
					return errors.New("value is not string")
				}
				requests[index] = typedValue
			}

			return bR.setUnservicedRequests(requests)
		},
	}
}

// GetBatchRequestID returns the id of the associated batch request
func (bR *BatchResponseModel) GetBatchRequestID() (*string, error) {
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
func (bR *BatchResponseModel) setBatchRequestID(id *string) error {
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
func (bR *BatchResponseModel) GetServicedRequests() ([]ServicedRequest, error) {
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

	typedServicedRequests, ok := servicedRequests.([]ServicedRequest)
	if !ok {
		return nil, errors.New("servicedRequests is not []ServicedRequestable")
	}

	return typedServicedRequests, nil
}

// setServicedRequests sets the serviced requests to the provided values
func (bR *BatchResponseModel) setServicedRequests(requests []ServicedRequest) error {
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
func (bR *BatchResponseModel) GetUnservicedRequests() ([]string, error) {
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
func (bR *BatchResponseModel) setUnservicedRequests(unservicedRequests []string) error {
	if internal.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil
	}

	return backingStore.Set(unservicedRequestsKey, unservicedRequests)
}
