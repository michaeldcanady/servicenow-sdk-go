package batchapi

import (
	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	restRequestsKey = "rest_requests"
)

// BatchRequest represents a Service-Now Batch API request
type BatchRequest interface {
	GetBatchRequestID() (*string, error)
	SetBatchRequestID(*string) error
	GetRestRequests() ([]RestRequest, error)
	SetRestRequests([]RestRequest) error
	AddRequest(RestRequest) error
	serialization.Parsable
	kiotaStore.BackedModel
}

// BatchRequestModel implementation of BatchRequest
type BatchRequestModel struct {
	internal.BackedModel
}

// NewBatchRequestModel creates a new BatchRequest.
func NewBatchRequestModel() *BatchRequestModel {
	request := &BatchRequestModel{
		internal.NewBaseModel(),
	}

	id := uuid.NewString()

	_ = request.SetBatchRequestID(&id)

	return request
}

// CreateBatchRequestFromDiscriminatorValue is a parsable factory for creating a BatchRequest
func CreateBatchRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchRequestModel(), nil
}

// Serialize writes the objects properties to the current writer.
func (bR *BatchRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(bR) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(batchRequestIDKey)(func() (*string, error) {
			id, err := bR.GetBatchRequestID()
			if err != nil {
				return nil, err
			}

			// ensure request has an id BEFORE serializing
			if conversion.IsNil(id) || *id == "" {
				idString := uuid.NewString()
				id = &idString
			}

			return id, nil
		}),
		internalSerialization.SerializeCollectionOfObjectValuesFunc[RestRequest](restRequestsKey)(bR.GetRestRequests),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (bR *BatchRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if conversion.IsNil(bR) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		batchRequestIDKey: internalSerialization.DeserializeStringFunc()(bR.SetBatchRequestID),
		restRequestsKey:   internalSerialization.DeserializeCollectionOfObjectValuesFunc[RestRequest](CreateRestRequestFromDiscriminatorValue)(bR.SetRestRequests),
	}
}

// AddRequest add request to the slice of requests.
func (bR *BatchRequestModel) AddRequest(request RestRequest) error {
	if conversion.IsNil(bR) {
		return nil
	}

	if conversion.IsNil(request) {
		return nil
	}

	requests, err := bR.GetRestRequests()
	if err != nil {
		return err
	}
	if conversion.IsNil(requests) {
		requests = []RestRequest{}
	}

	requests = append(requests, request)

	return bR.SetRestRequests(requests)
}

// GetBatchRequestID returns the id of the request.
func (bR *BatchRequestModel) GetBatchRequestID() (*string, error) {
	if conversion.IsNil(bR) {
		return nil, nil
	}
	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, batchRequestIDKey)
}

// SetBatchRequestID sets the id of the request.
func (bR *BatchRequestModel) SetBatchRequestID(id *string) error {
	if conversion.IsNil(bR) {
		return nil
	}
	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, batchRequestIDKey, id)
}

// GetRestRequests returns batched requests.
func (bR *BatchRequestModel) GetRestRequests() ([]RestRequest, error) {
	if conversion.IsNil(bR) {
		return nil, nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []RestRequest](backingStore, restRequestsKey)
}

// SetRestRequests sets the batched requests.
func (bR *BatchRequestModel) SetRestRequests(requests []RestRequest) error {
	if conversion.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, restRequestsKey, requests)
}
