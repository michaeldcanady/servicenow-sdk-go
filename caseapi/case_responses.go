package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CaseCollectionResponse represents a collection of cases.
type CaseCollectionResponse = core.ServiceNowCollectionResponse[*CaseResultModel]

// CreateCaseCollectionResponseFromDiscriminatorValue is a factory for creating a CaseCollectionResponse.
func CreateCaseCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil
}

// CaseItemResponse represents a single case response.
type CaseItemResponse = core.ServiceNowItemResponse[*CaseResultModel]

// CreateCaseItemResponseFromDiscriminatorValue is a factory for creating a CaseItemResponse.
func CreateCaseItemResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue), nil
}

// ActivitiesResponse represents a single activities response.
type ActivitiesResponse = core.ServiceNowItemResponse[*ActivitiesResultModel]

// CreateActivitiesResponseFromDiscriminatorValue is a factory for creating an ActivitiesResponse.
func CreateActivitiesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ActivitiesResultModel](CreateActivitiesResultFromDiscriminatorValue), nil
}

// FieldValuesResponse represents a single field values response.
type FieldValuesResponse = core.ServiceNowItemResponse[*FieldValuesResultModel]

// CreateFieldValuesResponseFromDiscriminatorValue is a factory for creating a FieldValuesResponse.
func CreateFieldValuesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*FieldValuesResultModel](CreateFieldValuesResultFromDiscriminatorValue), nil
}
