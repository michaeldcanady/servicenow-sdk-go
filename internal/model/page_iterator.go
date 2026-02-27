package model

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ErrNoMoreItems is returned when the iterator has reached the end of the collection.
var (
	ErrNoMoreItems = errors.New("no more items")
)

// PageIterator represents an iterator for paginated collections.
// T is the type of the items in the collection, which must implement [serialization.Parsable].
type PageIterator[T serialization.Parsable] struct {
	currentPage     PageResult[T]
	originalPage    PageResult[T]
	reqAdapter      abstractions.RequestAdapter
	pauseIndex      int
	constructorFunc serialization.ParsableFactory
	headers         *abstractions.RequestHeaders
	reqOptions      []abstractions.RequestOption
	errorMappings   abstractions.ErrorMappings
}

// NewPageIterator creates a new PageIterator instance.
//
// res is the initial response containing the first page of results.
// reqAdapter is the RequestAdapter used to fetch subsequent pages.
// constructorFunc is the factory function for creating new instances of T.
// options allows for additional configuration of the iterator.
func NewPageIterator[T serialization.Parsable](
	res ServiceNowCollectionResponse[T],
	reqAdapter abstractions.RequestAdapter,
	constructorFunc serialization.ParsableFactory,
	options ...kiota.Option[*PageIterator[T]],
) (*PageIterator[T], error) {
	if reqAdapter == nil {
		return nil, errors.New("reqAdapter can't be nil")
	}

	page, err := convertToPage(res)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internalErrors.CreateServiceNowErrorFromDiscriminatorValue,
	}

	iterator := &PageIterator[T]{
		currentPage:     page,
		originalPage:    page,
		reqAdapter:      reqAdapter,
		pauseIndex:      0,
		constructorFunc: constructorFunc,
		headers:         abstractions.NewRequestHeaders(),
		errorMappings:   errorMapping,
	}

	if err := kiota.ApplyOptions(iterator, options...); err != nil {
		return nil, err
	}

	return iterator, nil
}

// Reset returns the iterator to the initial state (first page, first item).
func (i *PageIterator[T]) Reset() {
	i.currentPage = i.originalPage
	i.pauseIndex = 0
}

// ResetPage restarts the iteration of the current page.
func (i *PageIterator[T]) ResetPage() {
	i.pauseIndex = 0
}

// WithHeaders sets the headers for the next page request.
func WithHeaders[T serialization.Parsable](headers *abstractions.RequestHeaders) kiota.Option[*PageIterator[T]] {
	return func(i *PageIterator[T]) error {
		i.headers = headers
		return nil
	}
}

// WithRequestOptions adds the request options for the next page request.
func WithRequestOptions[T serialization.Parsable](options ...abstractions.RequestOption) kiota.Option[*PageIterator[T]] {
	return func(i *PageIterator[T]) error {
		i.reqOptions = append(i.reqOptions, options...)
		return nil
	}
}

// Iterate traverses the pages and invokes the callback for each item.
//
// reverse determines the direction of page traversal.
// callback should return true to continue iteration, or false to stop.
func (i *PageIterator[T]) Iterate(ctx context.Context, reverse bool, callback func(T) bool) error {
	if reverse && i.pauseIndex == 0 {
		i.pauseIndex = len(i.currentPage.value)
	}

	for {
		var item T
		var err error

		if reverse {
			item, err = i.PreviousItem(ctx)
		} else {
			item, err = i.NextItem(ctx)
		}

		if errors.Is(err, ErrNoMoreItems) {
			return nil
		}
		if err != nil {
			return err
		}

		if !callback(item) {
			return nil
		}
	}
}

// HasNext returns true if there are more items to iterate over in the forward direction.
func (i *PageIterator[T]) HasNext() bool {
	return (i.pauseIndex < len(i.currentPage.value)) ||
		(i.currentPage.nextLink != nil && strings.TrimSpace(*i.currentPage.nextLink) != "")
}

// HasPrevious returns true if there are more items to iterate over in the reverse direction.
func (i *PageIterator[T]) HasPrevious() bool {
	return (i.pauseIndex > 0) ||
		(i.currentPage.prevLink != nil && strings.TrimSpace(*i.currentPage.prevLink) != "")
}

// NextItem returns the next item in the collection, fetching the next page if necessary.
//
// Returns [ErrNoMoreItems] if the end of the collection is reached.
func (i *PageIterator[T]) NextItem(ctx context.Context) (T, error) {
	for {
		if i.pauseIndex < 0 {
			i.pauseIndex = 0
		}

		if i.pauseIndex < len(i.currentPage.value) {
			item := i.currentPage.value[i.pauseIndex]
			i.pauseIndex++
			return item, nil
		}

		if !i.HasNext() {
			var zero T
			return zero, ErrNoMoreItems
		}

		_, err := i.Next(ctx)
		if err != nil {
			var zero T
			return zero, err
		}
	}
}

// PreviousItem returns the previous item in the collection, fetching the previous page if necessary.
//
// Returns [ErrNoMoreItems] if the beginning of the collection is reached.
func (i *PageIterator[T]) PreviousItem(ctx context.Context) (T, error) {
	for {
		if i.pauseIndex > len(i.currentPage.value) {
			i.pauseIndex = len(i.currentPage.value)
		}

		if i.pauseIndex > 0 {
			i.pauseIndex--
			item := i.currentPage.value[i.pauseIndex]
			return item, nil
		}

		if !i.HasPrevious() {
			var zero T
			return zero, ErrNoMoreItems
		}

		_, err := i.Previous(ctx)
		if err != nil {
			var zero T
			return zero, err
		}
	}
}

// Previous fetches and returns the previous page of results.
func (i *PageIterator[T]) Previous(ctx context.Context) (PageResult[T], error) {
	if i.currentPage.prevLink == nil || strings.TrimSpace(*i.currentPage.prevLink) == "" {
		return PageResult[T]{}, nil
	}

	collection, err := i.fetchPage(ctx, i.currentPage.prevLink)
	if err != nil {
		return PageResult[T]{}, err
	}

	page, err := convertToPage(collection)
	if err != nil {
		return PageResult[T]{}, err
	}

	i.updatePage(page, len(page.value))

	return i.currentPage, nil
}

// Next fetches and returns the next page of results.
func (i *PageIterator[T]) Next(ctx context.Context) (PageResult[T], error) {
	if i.currentPage.nextLink == nil || strings.TrimSpace(*i.currentPage.nextLink) == "" {
		return PageResult[T]{}, nil
	}

	collection, err := i.fetchPage(ctx, i.currentPage.nextLink)
	if err != nil {
		return PageResult[T]{}, err
	}

	page, err := convertToPage(collection)
	if err != nil {
		return PageResult[T]{}, err
	}

	i.updatePage(page, 0)

	return i.currentPage, nil
}

// First fetches and returns the first page of results.
func (i *PageIterator[T]) First(ctx context.Context) (PageResult[T], error) {
	if i.currentPage.firstLink == nil || strings.TrimSpace(*i.currentPage.firstLink) == "" {
		return PageResult[T]{}, nil
	}

	collection, err := i.fetchPage(ctx, i.currentPage.firstLink)
	if err != nil {
		return PageResult[T]{}, err
	}

	page, err := convertToPage(collection)
	if err != nil {
		return PageResult[T]{}, err
	}

	i.updatePage(page, 0)

	return i.currentPage, nil
}

// Last fetches and returns the last page of results.
func (i *PageIterator[T]) Last(ctx context.Context) (PageResult[T], error) {
	if i.currentPage.lastLink == nil || strings.TrimSpace(*i.currentPage.lastLink) == "" {
		return PageResult[T]{}, nil
	}

	collection, err := i.fetchPage(ctx, i.currentPage.lastLink)
	if err != nil {
		return PageResult[T]{}, err
	}

	page, err := convertToPage(collection)
	if err != nil {
		return PageResult[T]{}, err
	}

	i.updatePage(page, len(page.value))

	return i.currentPage, nil
}

func (i *PageIterator[T]) updatePage(page PageResult[T], pauseIndex int) {
	i.currentPage = page
	i.pauseIndex = pauseIndex
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
