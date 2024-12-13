package iterator

import (
	"context"
	"errors"
	"net/url"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type KiotaIterator interface {
	Send(context.Context, string) (HasResult, error)
}

type kiotaIteratorImpl struct {
	reqAdapter      abstractions.RequestAdapter
	constructorFunc serialization.ParsableFactory
	headers         *abstractions.RequestHeaders
	reqOptions      []abstractions.RequestOption
	errorMappings   abstractions.ErrorMappings
}

func NewKiotaIterator(reqAdapter abstractions.RequestAdapter,
	constructorFunc serialization.ParsableFactory,
	headers *abstractions.RequestHeaders,
	errorMappings abstractions.ErrorMappings,
	reqOptions ...abstractions.RequestOption,
) *kiotaIteratorImpl {
	return &kiotaIteratorImpl{
		reqAdapter:      reqAdapter,
		constructorFunc: constructorFunc,
		headers:         headers,
		reqOptions:      reqOptions,
		errorMappings:   errorMappings,
	}
}

func (it *kiotaIteratorImpl) Send(ctx context.Context, rawURI string) (HasResult, error) {
	requestInfo := abstractions.NewRequestInformation()
	requestInfo.Method = abstractions.GET

	uri, err := url.Parse(rawURI)
	if err != nil {
		return nil, err
	}

	requestInfo.SetUri(*uri)
	requestInfo.Headers.AddAll(it.headers)
	requestInfo.AddRequestOptions(it.reqOptions)

	response, err := it.reqAdapter.Send(context.Background(), requestInfo, it.constructorFunc, it.errorMappings)
	if err != nil {
		return nil, err
	}

	typedResponse, ok := response.(HasResult)
	if !ok {
		return nil, errors.New("invalid response")
	}

	return typedResponse, nil
}
