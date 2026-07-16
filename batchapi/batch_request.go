package batchapi

import (
	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
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
	core.BackedModel
}

// NewBatchRequestModel creates a new BatchRequest.
func NewBatchRequestModel() *BatchRequestModel {
	request := &BatchRequestModel{
		core.NewBaseModel(),
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
	return store.DefaultBackedModelAccessorFunc[*BatchRequestModel, *string](bR, batchRequestIDKey)
}

// SetBatchRequestID sets the id of the request.
func (bR *BatchRequestModel) SetBatchRequestID(id *string) error {
	return store.DefaultBackedModelMutatorFunc(bR, batchRequestIDKey, id)
}

// GetRestRequests returns batched requests.
func (bR *BatchRequestModel) GetRestRequests() ([]RestRequest, error) {
	return store.DefaultBackedModelAccessorFunc[*BatchRequestModel, []RestRequest](bR, restRequestsKey)
}

// SetRestRequests sets the batched requests.
func (bR *BatchRequestModel) SetRestRequests(requests []RestRequest) error {
	return store.DefaultBackedModelMutatorFunc(bR, restRequestsKey, requests)
}
