package servicenowsdkgo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
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

// serviceNowServiceClient is the core service used by ServiceNowServiceClient to make requests to Service-Now's APIs
type serviceNowServiceClient struct {
	newInternal.RequestBuilder
}

// registerDefaultSerializers registers default serializers
func registerDefaultSerializers() {
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
}

// registerDefaultDeserializers registers default deserializers
func registerDefaultDeserializers() {
	abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return jsonserialization.NewJsonParseNodeFactory() })
	abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return textserialization.NewTextParseNodeFactory() })
	abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return formserialization.NewFormParseNodeFactory() })
}

// newServiceNowServiceClientWithOptions creates new serviceNowServiceClient using provided options.
func newServiceNowServiceClientWithOptions(
	authenticationProvider authentication.AuthenticationProvider,
	opts ...serviceNowServiceClientOption,
) (*serviceNowServiceClient, error) {
	requestAdapter, err := newInternal.NewServiceNowRequestAdapter(authenticationProvider)
	if err != nil {
		return nil, err
	}

	config, err := buildServiceClientConfig(opts...)
	if err != nil {
		return nil, err
	}

	var baseURL = config.rawURI
	if baseURL == "" && config.instance == "" {
		return nil, errors.New("have to use either withURL or WithInstance")
	}

	if config.instance != "" {
		baseURL = fmt.Sprintf("https://%s.%s", config.instance, defaultServiceNowHost)
	}

	var backingStoreFactory = config.backingStoreFactory
	if internal.IsNil(backingStoreFactory) {
		backingStoreFactory = store.BackingStoreFactoryInstance
	}

	return newServiceNowServiceClient(requestAdapter, backingStoreFactory, baseURL)
}

// newServiceNowServiceClient creates a new ServiceNowBaseServiceClient with the given parameters
func newServiceNowServiceClient(
	requestAdapter abstractions.RequestAdapter,
	backingStoreFactory store.BackingStoreFactory,
	baseURL string,
) (*serviceNowServiceClient, error) {
	baseURL = strings.TrimSpace(baseURL)
	if baseURL == "" {
		return nil, errors.New("baseURL is empty")
	}
	if !newInternal.IsNil(backingStoreFactory) {
		requestAdapter.EnableBackingStore(backingStoreFactory)
	}

	requestAdapter.SetBaseUrl(baseURL)
	pathParameters := map[string]string{baseURLParameter: baseURL}

	registerDefaultSerializers()
	registerDefaultDeserializers()
	return &serviceNowServiceClient{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, baseURLVariable, pathParameters),
	}, nil
}
