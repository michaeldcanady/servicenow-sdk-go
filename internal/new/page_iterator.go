package internal

import (
	"context"
	"errors"
	"net/url"
	"strings"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type PageIterator[T serialization.Parsable] struct {
	currentPage     PageResult[T]
	reqAdapter      abstractions.RequestAdapter
	pauseIndex      int
	constructorFunc serialization.ParsableFactory
	headers         *abstractions.RequestHeaders
	reqOptions      []abstractions.RequestOption
	errorMappings   abstractions.ErrorMappings
}

// NewPageIterator creates an iterator instance
//
// It has three parameters. res is the graph response from the initial request and represents the first page.
// reqAdapter is used for getting the next page and constructorFunc is used for serializing next page's response to the specified type.
func NewPageIterator[T serialization.Parsable](res ServiceNowCollectionResponse[T], reqAdapter abstractions.RequestAdapter, constructorFunc serialization.ParsableFactory) (*PageIterator[T], error) {
	if reqAdapter == nil {
		return nil, errors.New("reqAdapter can't be nil")
	}

	page, err := convertToPage(res)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": CreateServiceNowErrorFromDiscriminatorValue,
	}

	return &PageIterator[T]{
		currentPage:     page,
		reqAdapter:      reqAdapter,
		pauseIndex:      0,
		constructorFunc: constructorFunc,
		headers:         abstractions.NewRequestHeaders(),
		errorMappings:   errorMapping,
	}, nil
}

func (i *PageIterator[T]) Iterate(ctx context.Context, reverse bool, callback func(T) bool) error {
	for {
		keepIterating := i.enumerate(callback)

		if !keepIterating {
			return nil
		}

		var err error
		var page PageResult[T]
		if reverse {
			page, err = i.Previous(ctx)
		} else {
			page, err = i.Next(ctx)
		}

		if err != nil {
			return err
		}

		if len(page.value) == 0 {
			return nil
		}

		i.pauseIndex = 0
	}
}

func (i *PageIterator[T]) Previous(ctx context.Context) (PageResult[T], error) {
	if i.currentPage.prevLink == nil || strings.TrimSpace(*i.currentPage.prevLink) == "" {
		return PageResult[T]{}, nil
	}

	collection, err := i.fetchPage(ctx, i.currentPage.prevLink)
	if err != nil {
		return PageResult[T]{}, err
	}

	i.currentPage, err = convertToPage(collection)

	return i.currentPage, err
}

func (i *PageIterator[T]) Next(ctx context.Context) (PageResult[T], error) {
	if i.currentPage.nextLink == nil || strings.TrimSpace(*i.currentPage.nextLink) == "" {
		return PageResult[T]{}, nil
	}

	collection, err := i.fetchPage(ctx, i.currentPage.nextLink)
	if err != nil {
		return PageResult[T]{}, err
	}

	i.currentPage, err = convertToPage(collection)

	return i.currentPage, err
}

func (i *PageIterator[T]) First(ctx context.Context) (PageResult[T], error) {
	if i.currentPage.firstLink == nil || strings.TrimSpace(*i.currentPage.firstLink) == "" {
		return PageResult[T]{}, nil
	}

	collection, err := i.fetchPage(ctx, i.currentPage.firstLink)
	if err != nil {
		return PageResult[T]{}, err
	}

	i.currentPage, err = convertToPage(collection)

	return i.currentPage, err
}

func (i *PageIterator[T]) Last(ctx context.Context) (PageResult[T], error) {
	if i.currentPage.lastLink == nil || strings.TrimSpace(*i.currentPage.lastLink) == "" {
		return PageResult[T]{}, nil
	}

	collection, err := i.fetchPage(ctx, i.currentPage.lastLink)
	if err != nil {
		return PageResult[T]{}, err
	}

	i.currentPage, err = convertToPage(collection)

	return i.currentPage, err
}

// SetHeaders sets the headers for the next page request.
func (i *PageIterator[T]) SetHeaders(headers *abstractions.RequestHeaders) {
	i.headers = headers
}

// AddRequestOptions adds the request options for the next page request.
func (i *PageIterator[T]) AddRequestOptions(options ...abstractions.RequestOption) {
	i.reqOptions = append(i.reqOptions, options...)
}

func (i *PageIterator[T]) fetchPage(ctx context.Context, pageLink *string) (ServiceNowCollectionResponse[T], error) {
	var response ServiceNowCollectionResponse[T]
	var rawResponse serialization.Parsable
	var err error

	if pageLink == nil || strings.TrimSpace(*pageLink) == "" {
		return response, nil
	}

	link, err := url.Parse(*pageLink)
	if err != nil {
		return response, errors.New("parsing nextLink url failed")
	}

	requestInfo := abstractions.NewRequestInformation()
	requestInfo.Method = abstractions.GET
	requestInfo.SetUri(*link)
	requestInfo.Headers.AddAll(i.headers)
	requestInfo.AddRequestOptions(i.reqOptions)

	rawResponse, err = i.reqAdapter.Send(ctx, requestInfo, i.constructorFunc, i.errorMappings)
	if err != nil {
		return response, err
	}

	response, ok := rawResponse.(ServiceNowCollectionResponse[T])
	if !ok {
		return response, errors.New("response is of wrong type")
	}

	return response, nil
}

func (pI *PageIterator[T]) enumerate(callback func(item T) bool) bool {
	keepIterating := true

	// the current page has no items to enumerate
	pageItems := pI.currentPage.value
	if pageItems == nil {
		return false
	}

	// start/continue enumerating page items from  pauseIndex.
	// this makes it possible to resume iteration from where we paused iteration.
	for i := pI.pauseIndex; i < len(pageItems); i++ {
		keepIterating = callback(pageItems[i])

		// Set pauseIndex so that we know where to resume from.
		// Resumes from the next item
		pI.pauseIndex = i + 1

		if !keepIterating {
			break
		}
	}

	return keepIterating
}
