package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CaseCollectionResponse represents a collection of cases.
type CaseCollectionResponse = internal.ServiceNowCollectionResponse[*CaseResultModel]

// CreateCaseCollectionResponseFromDiscriminatorValue is a factory for creating a CaseCollectionResponse.
func CreateCaseCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil
}

// CaseItemResponse represents a single case response.
type CaseItemResponse = internal.ServiceNowItemResponse[*CaseResultModel]

// CreateCaseItemResponseFromDiscriminatorValue is a factory for creating a CaseItemResponse.
func CreateCaseItemResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil
}

// ActivitiesResponse represents a single activities response.
type ActivitiesResponse = internal.ServiceNowItemResponse[*ActivitiesResultModel]

// CreateActivitiesResponseFromDiscriminatorValue is a factory for creating an ActivitiesResponse.
func CreateActivitiesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*ActivitiesResultModel](CreateActivitiesResultFromDiscriminatorValue), nil
}

// FieldValuesResponse represents a single field values response.
type FieldValuesResponse = internal.ServiceNowItemResponse[*FieldValuesResultModel]

// CreateFieldValuesResponseFromDiscriminatorValue is a factory for creating a FieldValuesResponse.
func CreateFieldValuesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*FieldValuesResultModel](CreateFieldValuesResultFromDiscriminatorValue), nil
}
