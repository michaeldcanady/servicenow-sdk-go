package servicenowsdkgo

import (
	"errors"
	"maps"
	"strings"
	"sync"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	formserialization "github.com/microsoft/kiota-serialization-form-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	multipartserialization "github.com/microsoft/kiota-serialization-multipart-go"
	textserialization "github.com/microsoft/kiota-serialization-text-go"
)

const (
	// defaultServiceNowHost the default Service-Now url
	defaultServiceNowHost = "service-now.com"
	// baseURLParameter the path parameter for the base url
	baseURLParameter = "baseurl"
	// baseURLVariable the url template variable for the base url
	baseURLVariable = "{+baseurl}"
)

var (
	registerSerializersOnce   sync.Once
	registerDeserializersOnce sync.Once
)

// ServiceNowServiceClient is the core service used by ServiceNowServiceClient to make requests to Service-Now's APIs
type ServiceNowServiceClient struct {
	kiota.RequestBuilder
}

// registerDefaultSerializers registers default serializers
func registerDefaultSerializers() {
	registerSerializersOnce.Do(func() {
		abstractions.RegisterDefaultSerializer(func() serialization.SerializationWriterFactory {
			return jsonserialization.NewJsonSerializationWriterFactory()
		})
		abstractions.RegisterDefaultSerializer(func() serialization.SerializationWriterFactory {
			return textserialization.NewTextSerializationWriterFactory()
		})
		abstractions.RegisterDefaultSerializer(func() serialization.SerializationWriterFactory {
			return formserialization.NewFormSerializationWriterFactory()
		})
		abstractions.RegisterDefaultSerializer(func() serialization.SerializationWriterFactory {
			return multipartserialization.NewMultipartSerializationWriterFactory()
		})
	})
}

// registerDefaultDeserializers registers default deserializers
func registerDefaultDeserializers() {
	registerDeserializersOnce.Do(func() {
		abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return jsonserialization.NewJsonParseNodeFactory() })
		abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return textserialization.NewTextParseNodeFactory() })
		abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return formserialization.NewFormParseNodeFactory() })
	})
}

// NewServiceNowServiceClientWithOptions creates new serviceNowServiceClient using provided options.
func NewServiceNowServiceClientWithOptions(
	authenticationProvider authentication.AuthenticationProvider,
	opts ...ServiceNowServiceClientOption,
) (*ServiceNowServiceClient, error) {
	config, err := buildServiceClientConfig(opts...)
	if err != nil {
		return nil, err
	}

	requestAdapter, err := config.getRequestAdapter()
	if err != nil {
		return nil, err
	}

	var backingStoreFactory = config.backingStoreFactory
	if utils.IsNil(backingStoreFactory) {
		backingStoreFactory = store.BackingStoreFactoryInstance
	}

	return NewServiceNowServiceClient(requestAdapter, backingStoreFactory, config.rawURI)
}

// NewServiceNowServiceClient creates a new ServiceNowBaseServiceClient with the given parameters
func NewServiceNowServiceClient(
	requestAdapter abstractions.RequestAdapter,
	backingStoreFactory store.BackingStoreFactory,
	baseURL string,
) (*ServiceNowServiceClient, error) {
	baseURL = strings.TrimSpace(baseURL)
	if baseURL == "" {
		return nil, errors.New("baseURL is empty")
	}
	if !utils.IsNil(backingStoreFactory) {
		requestAdapter.EnableBackingStore(backingStoreFactory)
	}

	pathParameters := map[string]string{baseURLParameter: baseURL}

	registerDefaultSerializers()
	registerDefaultDeserializers()
	return &ServiceNowServiceClient{
		RequestBuilder: kiota.NewBaseRequestBuilder(requestAdapter, baseURLVariable, pathParameters),
	}, nil
}

// Now returns a NowRequestBuilder associated with the client.
func (c *ServiceNowServiceClient) Now() *NowRequestBuilder {
	return NewNowRequestBuilderInternal(maps.Clone(c.GetPathParameters()), c.GetRequestAdapter())
}
