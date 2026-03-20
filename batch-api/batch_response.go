package batchapi

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
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
	kiotaStore.BackedModel
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
	if utils.IsNil(bR) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(batchRequestIDKey)(bR.GetBatchRequestID),
		internalSerialization.SerializeCollectionOfObjectValuesFunc[ServicedRequest](servicedRequestsKey)(bR.GetServicedRequests),
		internalSerialization.SerializeCollectionOfStringValuesFunc(unservicedRequestsKey)(bR.GetUnservicedRequests),
	)
}

// GetFieldDeserializers returns the deserialization information for this object
func (bR *BatchResponseModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if utils.IsNil(bR) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		batchRequestIDKey: internalSerialization.DeserializeStringFunc()(bR.setBatchRequestID),
		servicedRequestsKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetCollectionOfObjectValues(CreateServicedRequestFromDiscriminatorValue)
			if err != nil {
				return err
			}

			requests := make([]ServicedRequest, len(val))
			for index, value := range val {
				typedValue, ok := value.(ServicedRequest)
				if !ok {
					return fmt.Errorf("value is not ServicedRequest")
				}
				requests[index] = typedValue
			}

			return bR.setServicedRequests(requests)
		},
		unservicedRequestsKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetCollectionOfPrimitiveValues("string")
			if err != nil {
				return err
			}

			requests := make([]string, len(val))
			for index, value := range val {
				requests[index] = value.(string)
			}

			return bR.setUnservicedRequests(requests)
		},
	}
}

// GetBatchRequestID returns the id of the associated batch request
func (bR *BatchResponseModel) GetBatchRequestID() (*string, error) {
	if utils.IsNil(bR) {
		return nil, nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, batchRequestIDKey)
}

// GetServicedRequestByID returns the serviced request with the provided id
func (bR *BatchResponseModel) GetServicedRequestByID(id string) (ServicedRequest, error) {
	if internal.IsNil(bR) {
		return nil, nil
	}

	requests, err := bR.GetServicedRequests()
	if err != nil {
		return nil, err
	}

	for _, request := range requests {
		reqID, err := request.GetID()
		if err != nil {
			continue
		}

		if reqID != nil && *reqID == id {
			return request, nil
		}
	}

	return nil, nil
}

// setBatchRequestID sets the id of the associated batch request
func (bR *BatchResponseModel) setBatchRequestID(id *string) error {
	if utils.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, batchRequestIDKey, id)
}

// GetServicedRequests returns serviced requests
func (bR *BatchResponseModel) GetServicedRequests() ([]ServicedRequest, error) {
	if utils.IsNil(bR) {
		return nil, nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []ServicedRequest](backingStore, servicedRequestsKey)
}

// setServicedRequests sets the serviced requests to the provided values
func (bR *BatchResponseModel) setServicedRequests(requests []ServicedRequest) error {
	if utils.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, servicedRequestsKey, requests)
}

// GetUnservicedRequests returns the unserviced requests' id
func (bR *BatchResponseModel) GetUnservicedRequests() ([]string, error) {
	if utils.IsNil(bR) {
		return nil, nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []string](backingStore, unservicedRequestsKey)
}

// setUnservicedRequests sets the ids of the unserviced requests to the provided value
func (bR *BatchResponseModel) setUnservicedRequests(unservicedRequests []string) error {
	if utils.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, unservicedRequestsKey, unservicedRequests)
}
