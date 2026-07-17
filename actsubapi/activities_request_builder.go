package actsubapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	activitiesURLTemplate = "{+baseurl}/api/now/v1/actsub/activities"
)

// ActivitiesRequestBuilder provides operations to manage activities.
type ActivitiesRequestBuilder struct {
	*collectionGetRequestBuilder
}

// NewActivitiesRequestBuilderInternal instantiates a new ActivitiesRequestBuilder.
func NewActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ActivitiesRequestBuilder {
	return &ActivitiesRequestBuilder{
		newCollectionGetRequestBuilder(pathParameters, requestAdapter, activitiesURLTemplate),
	}
}
