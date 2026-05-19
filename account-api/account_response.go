package accountapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountCollectionResponse represents a collection of accounts.
type AccountCollectionResponse = newInternal.ServiceNowCollectionResponse[*AccountModel]

// CreateAccountCollectionResponseFromDiscriminatorValue is a factory for creating an AccountCollectionResponse.
func CreateAccountCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*AccountModel](CreateAccountFromDiscriminatorValue), nil
}

// AccountCollectionResponseModel is the implementation of AccountCollectionResponse.
type AccountCollectionResponseModel = newInternal.BaseServiceNowCollectionResponse[*AccountModel]

// AccountItemResponse represents a single account response.
type AccountItemResponse = newInternal.ServiceNowItemResponse[*AccountModel]

// CreateAccountItemResponseFromDiscriminatorValue is a factory for creating an AccountItemResponse.
func CreateAccountItemResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*AccountModel](CreateAccountFromDiscriminatorValue), nil
}

// AccountItemResponseModel is the implementation of AccountItemResponse.
type AccountItemResponseModel = newInternal.BaseServiceNowItemResponse[*AccountModel]
