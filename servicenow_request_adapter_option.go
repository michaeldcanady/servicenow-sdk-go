package servicenowsdkgo

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

var (
	// WithClient option to use a custom http.Client for the ServiceNowRequestAdapter
	WithClient = internal.WithClient
	// WithSerializationFactory option to add additional serialization.SerializationFactories to the ServiceNowRequestAdapter
	WithSerializationFactory = internal.WithSerializationFactory
	// WithParseNodeFactory option to add additional serialization.ParseNodeFactories to the ServiceNowRequestAdapter
	WithParseNodeFactory = internal.WithParseNodeFactory
)
