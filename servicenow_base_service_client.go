package servicenowsdkgo

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	formserialization "github.com/microsoft/kiota-serialization-form-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	multipartserialization "github.com/microsoft/kiota-serialization-multipart-go"
	textserialization "github.com/microsoft/kiota-serialization-text-go"
)

type ServiceNowBaseServiceClient struct {
	abstractions.BaseRequestBuilder
}

func (rB *ServiceNowBaseServiceClient) Now() NowRequestBuilder2 {
	return *NewNowRequestBuilder2Internal(rB.BaseRequestBuilder.PathParameters, rB.BaseRequestBuilder.RequestAdapter)
}

func NewServiceNowBaseServiceClient(
	requestAdapter abstractions.RequestAdapter,
	backingStoreFactory store.BackingStoreFactory,
	baseURL string,
) *ServiceNowBaseServiceClient {
	m := &ServiceNowBaseServiceClient{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
	}

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

	m.BaseRequestBuilder.RequestAdapter.SetBaseUrl(baseURL)

	abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return jsonserialization.NewJsonParseNodeFactory() })
	abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return textserialization.NewTextParseNodeFactory() })
	abstractions.RegisterDefaultDeserializer(func() serialization.ParseNodeFactory { return formserialization.NewFormParseNodeFactory() })

	m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
	m.BaseRequestBuilder.RequestAdapter.EnableBackingStore(backingStoreFactory)
	return m
}
