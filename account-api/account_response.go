package accountapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountCollectionResponse represents a collection of accounts.
type AccountCollectionResponse = newInternal.ServiceNowCollectionResponse[*Account]

// CreateAccountCollectionResponseFromDiscriminatorValue is a factory for creating an AccountCollectionResponse.
func CreateAccountCollectionResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*Account](CreateAccountFromDiscriminatorValue), nil
}

// AccountCollectionResponseModel is the implementation of AccountCollectionResponse.
type AccountCollectionResponseModel = newInternal.BaseServiceNowCollectionResponse[*Account]

// AccountItemResponse represents a single account response.
type AccountItemResponse = newInternal.ServiceNowItemResponse[*Account]

// CreateAccountItemResponseFromDiscriminatorValue is a factory for creating an AccountItemResponse.
func CreateAccountItemResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*Account](CreateAccountFromDiscriminatorValue), nil
}

// AccountItemResponseModel is the implementation of AccountItemResponse.
type AccountItemResponseModel = newInternal.BaseServiceNowItemResponse[*Account]
