package accountapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountCollectionResponse represents a collection of accounts.
type AccountCollectionResponse = core.ServiceNowCollectionResponse[*AccountModel]

// CreateAccountCollectionResponseFromDiscriminatorValue is a factory for creating an AccountCollectionResponse.
func CreateAccountCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*AccountModel](CreateAccountFromDiscriminatorValue), nil
}

// AccountCollectionResponseModel is the implementation of AccountCollectionResponse.
type AccountCollectionResponseModel = core.BaseServiceNowCollectionResponse[*AccountModel]

// AccountItemResponse represents a single account response.
type AccountItemResponse = core.ServiceNowItemResponse[*AccountModel]

// CreateAccountItemResponseFromDiscriminatorValue is a factory for creating an AccountItemResponse.
func CreateAccountItemResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*AccountModel](CreateAccountFromDiscriminatorValue), nil
}

// AccountItemResponseModel is the implementation of AccountItemResponse.
type AccountItemResponseModel = core.BaseServiceNowItemResponse[*AccountModel]
