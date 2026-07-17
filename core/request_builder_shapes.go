package core

import (
	"context"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// requestConfigurationShape describes the field layout of abstractions.RequestConfiguration[Q]
// (Headers, Options, QueryParameters *Q) as a ~-constraint, rather than naming
// abstractions.RequestConfiguration[Q] directly - Go doesn't allow ~ on an instantiated
// named generic type from another package, only on the type's underlying literal shape.
//
// This only matches request-configuration types built on the "typed QueryParameters"
// convention (a type alias of abstractions.RequestConfiguration[Q], as actsubapi and
// appserviceapi use). It intentionally does NOT match the other request-configuration
// convention used elsewhere in this codebase (e.g. documentsapi's
// VersionStateRequestBuilderGetRequestConfiguration), whose bespoke structs omit
// QueryParameters entirely (Headers/Options, or Headers/Options/Data) - those are a
// genuinely different shape, not just a differently-named one.
type requestConfigurationShape[Q any] interface {
	~struct {
		Headers         *abstractions.RequestHeaders
		Options         []abstractions.RequestOption
		QueryParameters *Q
	}
}

// C is the request-configuration type used by a given verb (e.g.
// ActivitiesRequestBuilderGetRequestConfiguration). Every real builder in this codebase
// takes *C, never C by value, so these interfaces take *C too.

type CollectionGetRequestBuilder[T serialization.Parsable, Q any, C requestConfigurationShape[Q]] interface {
	Get(ctx context.Context, requestConfiguration *C) (*BaseServiceNowCollectionResponse[T], error)
	ToGetRequestInformation(ctx context.Context, requestConfiguration *C) (*abstractions.RequestInformation, error)
}

// Deliberately no CollectionPostRequestBuilder: no Post method anywhere in this codebase
// returns a collection response - every real Post returns either a single item
// (matched by ItemPostRequestBuilder below) or a bespoke response type outside this
// taxonomy (e.g. appserviceapi's CreateServiceResponse).

type ItemGetRequestBuilder[T serialization.Parsable, Q any, C requestConfigurationShape[Q]] interface {
	Get(ctx context.Context, requestConfiguration *C) (*BaseServiceNowItemResponse[T], error)
	ToGetRequestInformation(ctx context.Context, requestConfiguration *C) (*abstractions.RequestInformation, error)
}

type ItemPostRequestBuilder[T serialization.Parsable, Q any, C requestConfigurationShape[Q]] interface {
	Post(ctx context.Context, body T, requestConfiguration *C) (*BaseServiceNowItemResponse[T], error)
	ToPostRequestInformation(ctx context.Context, body T, requestConfiguration *C) (*abstractions.RequestInformation, error)
}

type ItemPatchRequestBuilder[T serialization.Parsable, Q any, C requestConfigurationShape[Q]] interface {
	Patch(ctx context.Context, body T, requestConfiguration *C) (*BaseServiceNowItemResponse[T], error)
	ToPatchRequestInformation(ctx context.Context, body T, requestConfiguration *C) (*abstractions.RequestInformation, error)
}

type ItemDeleteRequestBuilder[Q any, C requestConfigurationShape[Q]] interface {
	Delete(ctx context.Context, requestConfiguration *C) error
	ToDeleteRequestInformation(ctx context.Context, requestConfiguration *C) (*abstractions.RequestInformation, error)
}
