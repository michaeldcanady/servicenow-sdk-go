package servicenowsdkgo

import (
	"errors"
	"maps"
	"strings"
	"sync"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
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
	newInternal.RequestBuilder
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

// NewServiceNowServiceClient creates a new ServiceNowServiceClient using the provided options.
func NewServiceNowServiceClient(opts ...ServiceNowServiceClientOption) (*ServiceNowServiceClient, error) {
	config, err := buildServiceClientConfig(opts...)
	if err != nil {
		return nil, err
	}

	requestAdapter, err := config.getRequestAdapter()
	if err != nil {
		return nil, err
	}

	baseURL := config.getBaseURL()
	if strings.TrimSpace(baseURL) == "" {
		return nil, errors.New("baseURL cannot be empty")
	}

	pathParameters := map[string]string{baseURLParameter: baseURL}

	registerDefaultSerializers()
	registerDefaultDeserializers()

	return &ServiceNowServiceClient{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, baseURLVariable, pathParameters),
	}, nil
}

func (rB *ServiceNowServiceClient) Now() *NowRequestBuilder2 {
	return NewServiceNowRequestBuilder3Internal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
