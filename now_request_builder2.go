package servicenowsdkgo

import (
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	nowDefaultURLTemplate   = "{+baseurl}/api/now"
	nowVersionedURLTemplate = "{+baseurl}/api/now/{version}"
)

type NowRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
}

// NewNowRequestBuilder2Internal ...
func NewNowRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder2 {
	m := &NowRequestBuilder2{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, nowDefaultURLTemplate, pathParameters),
	}
	return m
}

// NewNowRequestBuilder2 ...
func NewNowRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *NowRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL
	return NewNowRequestBuilder2Internal(urlParams, requestAdapter)
}

func (rB *NowRequestBuilder2) Table(table string) tableapi.TableRequestBuilder2 {
	return *tableapi.NewDefaultTableRequestBuilder2Internal(rB.BaseRequestBuilder.PathParameters, rB.BaseRequestBuilder.RequestAdapter)
}
