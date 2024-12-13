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

type ReverseIterator[T serialization.Parsable] struct {
	Iterator[T]
	KiotaIterator
	// ctx used as a way to pass context indirectly
	ctx context.Context
}

func NewReverseIterator[T serialization.Parsable](reqAdapter abstractions.RequestAdapter, constructorFunc serialization.ParsableFactory) (*ReverseIterator[T], error) {
	it := &ReverseIterator[T]{
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

type HasPreviousLink interface {
	GetPreviousLink() (*string, error)
}

func (it *ReverseIterator[T]) fetch(_ int, pageToken string) (string, error) {
	response, err := it.Send(context.Background(), pageToken)
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

	nextResponse, ok := response.(HasPreviousLink)
	if !ok {
		return "", nil
	}

	previousLink, err := nextResponse.GetPreviousLink()
	if err != nil {
		return "", err
	}

	if previousLink == nil || *previousLink == "" {
		return "", nil
	}

	return *previousLink, nil
}

func (it *ReverseIterator[T]) Next(ctx context.Context) (T, error) {
	it.ctx = ctx
	defer func() {
		it.ctx = nil
	}()

	return it.Iterator.Next()
}

func (it *ReverseIterator[T]) Iterate(ctx context.Context, callback func(item T) bool) error {
	for {
		item, err := it.Next(ctx)
		if err == errDone || !callback(item) {
			return nil
		}
	}
}
