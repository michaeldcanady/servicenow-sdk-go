// Copyright 2016 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package iterator provides support for standard Google API iterators.
// See https://github.com/GoogleCloudPlatform/gcloud-golang/wiki/Iterator-Guidelines.
package iterator

import (
	"context"
	"errors"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// errDone is returned by an iterator's Next method when the iteration is
// complete; when there are no more items to return.
var errDone = errors.New("no more items in iterator")

// We don't support mixed calls to Next and NextPage because they play
// with the paging state in incompatible ways.
var errMixed = errors.New("iterator: Next and NextPage called on same iterator")

type ForwardIterator[T serialization.Parsable] struct {
	Iterator[T]
	KiotaIterator
	// ctx used as a way to pass context indirectly
	ctx context.Context
}

func NewForwardIterator[T serialization.Parsable](reqAdapter abstractions.RequestAdapter, constructorFunc serialization.ParsableFactory) (*ForwardIterator[T], error) {
	it := &ForwardIterator[T]{
		KiotaIterator: NewKiotaIterator(reqAdapter, constructorFunc, nil, nil),
	}

	bIT := BaseIterator[T]{}

	bIT.pageInfo, bIT.nextFunc = newPageInfo(
		it.fetch,
		bIT.bufLen,
		bIT.takeBuf,
	)

	it.Iterator = &bIT

	return it, nil
}

func (it *ForwardIterator[T]) fetch(_ int, pageToken string) (string, error) {
	response, err := it.Send(it.ctx, pageToken)
	if err != nil {
		return "", err
	}

	if pageToken == "" {
		return "", errDone
	}

	items, err := response.GetResult()
	if err != nil {
		return "", err
	}

	buf := make([]T, 0, len(items))
	for _, item := range items {
		typedItem, ok := item.(T)
		if !ok {
			return "", errors.New("item is not T")
		}
		buf = append(buf, typedItem)
	}

	it.setBuf(buf)

	nextResponse, ok := response.(HasNextLink)
	if !ok {
		return "", nil
	}

	nextLink, err := nextResponse.GetNextLink()
	if err != nil {
		return "", err
	}

	if nextLink == nil || *nextLink == "" {
		return "", nil
	}

	return *nextLink, nil
}

func (it *ForwardIterator[T]) Next(ctx context.Context) (T, error) {
	it.ctx = ctx
	defer func() {
		it.ctx = nil
	}()

	return it.Iterator.Next()
}

func (it *ForwardIterator[T]) Iterate(ctx context.Context, callback func(item T) bool) error {
	for {
		item, err := it.Next(ctx)
		if err == errDone || !callback(item) {
			return nil
		}
	}
}
