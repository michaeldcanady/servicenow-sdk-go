package accountapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountCollectionResponse represents a collection of accounts.
type AccountCollectionResponse = internal.ServiceNowCollectionResponse[*AccountModel]

// CreateAccountCollectionResponseFromDiscriminatorValue is a factory for creating an AccountCollectionResponse.
func CreateAccountCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*AccountModel](CreateAccountFromDiscriminatorValue), nil
}

// AccountCollectionResponseModel is the implementation of AccountCollectionResponse.
type AccountCollectionResponseModel = internal.BaseServiceNowCollectionResponse[*AccountModel]

// AccountItemResponse represents a single account response.
type AccountItemResponse = internal.ServiceNowItemResponse[*AccountModel]

// CreateAccountItemResponseFromDiscriminatorValue is a factory for creating an AccountItemResponse.
func CreateAccountItemResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*AccountModel](CreateAccountFromDiscriminatorValue), nil
}

// AccountItemResponseModel is the implementation of AccountItemResponse.
type AccountItemResponseModel = internal.BaseServiceNowItemResponse[*AccountModel]
