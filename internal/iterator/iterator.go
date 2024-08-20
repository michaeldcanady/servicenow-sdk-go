package iterator

import (
	"context"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

type iterator[T any] struct {
	info            *PageInfo
	client          intCore.Client2
	constructorFunc intCore.Parsable[T]
	nextFunc        func() error
	pager           *pager
	items           []T
}

func Newiterator[T any](resp intCore.CollectionResponse[T], client intCore.Client2, constructorFunc intCore.Parsable[T]) Iterable[T] {
	iter := &iterator[T]{
		info:            nil,
		client:          client,
		constructorFunc: constructorFunc,
		pager:           nil,
	}

	info, _ := newPageInfo(iter.fetch, iter.len, iter.takeBuf)
	iter.info = info
	iter.nextFunc = info.next

	return iter
}

func (it *iterator[T]) fetch(_ int, pageToken string) (string, error) {
	reqInfo := intCore.NewRequestInformation(intCore.WithMethod(intCore.GET), intCore.WithURL(pageToken))
	//mapping := intCore.ErrorMapping

	resp, err := it.client.SendWithContext(context.Background(), reqInfo, nil)
	if err != nil {
		return "", err
	}
	collResponse, err := it.constructorFunc(resp)
	if err != nil {
		return "", err
	}

	it.items = collResponse.Results()

	//TOOD: make return next page link
	return "", nil
}

func (it *iterator[T]) len() int {
	return len(it.items)
}

func (it *iterator[T]) takeBuf() interface{} {
	it.items = make([]T, 0)

	return it.items
}

func (it *iterator[T]) Next() (T, error) {
	var empty T

	if err := it.nextFunc(); err != nil {
		return empty, err
	}

	item := it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *iterator[T]) PageInfo() *PageInfo {
	return it.info
}
