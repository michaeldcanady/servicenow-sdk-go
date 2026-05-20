package caseapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CaseCollectionResponse represents a collection of cases.
type CaseCollectionResponse = newInternal.ServiceNowCollectionResponse[*CaseResultModel]

// CreateCaseCollectionResponseFromDiscriminatorValue is a factory for creating a CaseCollectionResponse.
func CreateCaseCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil
}

// CaseItemResponse represents a single case response.
type CaseItemResponse = newInternal.ServiceNowItemResponse[*CaseResultModel]

// CreateCaseItemResponseFromDiscriminatorValue is a factory for creating a CaseItemResponse.
func CreateCaseItemResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil
}

// ActivitiesResponse represents a single activities response.
type ActivitiesResponse = newInternal.ServiceNowItemResponse[*ActivitiesResultModel]

// CreateActivitiesResponseFromDiscriminatorValue is a factory for creating an ActivitiesResponse.
func CreateActivitiesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*ActivitiesResultModel](CreateActivitiesResultFromDiscriminatorValue), nil
}

// FieldValuesResponse represents a single field values response.
type FieldValuesResponse = newInternal.ServiceNowItemResponse[*FieldValuesResultModel]

// CreateFieldValuesResponseFromDiscriminatorValue is a factory for creating a FieldValuesResponse.
func CreateFieldValuesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*FieldValuesResultModel](CreateFieldValuesResultFromDiscriminatorValue), nil
}
