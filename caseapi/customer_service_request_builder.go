package caseapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const customerServiceURLTemplate = "{+baseurl}/api/sn_customerservice"

// CustomerServiceRequestBuilder provides operations to manage Customer Service APIs.
type CustomerServiceRequestBuilder struct {
	internal.RequestBuilder
}

// NewCustomerServiceRequestBuilderInternal instantiates a new CustomerServiceRequestBuilder.
func NewCustomerServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CustomerServiceRequestBuilder {
	return &CustomerServiceRequestBuilder{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, customerServiceURLTemplate, pathParameters),
	}
}

// Case returns a CaseRequestBuilder.
func (rB *CustomerServiceRequestBuilder) Case() *CaseRequestBuilder {
	return NewCaseRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
