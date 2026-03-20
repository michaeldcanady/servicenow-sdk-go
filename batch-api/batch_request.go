package batchapi

import (
	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
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
	newInternal.Model
}

// NewBatchRequestModel creates a new BatchRequest.
func NewBatchRequestModel() *BatchRequestModel {
	request := &BatchRequestModel{
		newInternal.NewBaseModel(),
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
	if utils.IsNil(bR) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(batchRequestIDKey)(func() (*string, error) {
			id, err := bR.GetBatchRequestID()
			if err != nil {
				return nil, err
			}

			// ensure request has an id BEFORE serializing
			if utils.IsNil(id) || *id == "" {
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
	if utils.IsNil(bR) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		batchRequestIDKey: internalSerialization.DeserializeStringFunc()(bR.SetBatchRequestID),
		restRequestsKey:   internalSerialization.DeserializeCollectionOfObjectValuesFunc[RestRequest](CreateRestRequestFromDiscriminatorValue)(bR.SetRestRequests),
	}
}

// AddRequest add request to the slice of requests.
func (bR *BatchRequestModel) AddRequest(request RestRequest) error {
	if utils.IsNil(bR) {
		return nil
	}

	if internal.IsNil(request) {
		return nil
	}

	requests, err := bR.GetRestRequests()
	if err != nil {
		return err
	}
	if utils.IsNil(requests) {
		requests = []RestRequest{}
	}

	requests = append(requests, request)

	return bR.SetRestRequests(requests)
}

// GetBatchRequestID returns the id of the request.
func (bR *BatchRequestModel) GetBatchRequestID() (*string, error) {
	if utils.IsNil(bR) {
		return nil, nil
	}
	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, batchRequestIDKey)
}

// SetBatchRequestID sets the id of the request.
func (bR *BatchRequestModel) SetBatchRequestID(id *string) error {
	if utils.IsNil(bR) {
		return nil
	}
	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, batchRequestIDKey, id)
}

// GetRestRequests returns batched requests.
func (bR *BatchRequestModel) GetRestRequests() ([]RestRequest, error) {
	if utils.IsNil(bR) {
		return nil, nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []RestRequest](backingStore, restRequestsKey)
}

// SetRestRequests sets the batched requests.
func (bR *BatchRequestModel) SetRestRequests(requests []RestRequest) error {
	if utils.IsNil(bR) {
		return nil
	}

	backingStore := bR.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, restRequestsKey, requests)
}
