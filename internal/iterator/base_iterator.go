package iterator

type BaseIterator[T interface{}] struct {
	pageInfo  *PageInfo
	nextFunc  func() error
	max       int
	items     []T
	prevItems []T
}

type Iterator[T interface{}] interface {
	Next() (T, error)
	PageInfo() *PageInfo
	setBuf(buf []T)
}

func NewBaseIterator[T interface{}](fetch func(int, string) (string, error)) (*BaseIterator[T], error) {
	it := &BaseIterator[T]{}

	it.pageInfo, it.nextFunc = newPageInfo(
		fetch,
		it.bufLen,
		it.takeBuf,
	)

	return it, nil
}

// PageInfo returns a PageInfo, which supports pagination.
func (it *BaseIterator[T]) PageInfo() *PageInfo {
	return it.pageInfo
}

// takeBuf Function that returns the buffer, after setting the buffer variable to nil.
func (it *BaseIterator[T]) takeBuf() interface{} {
	buf := it.items

	it.items = nil

	return buf
}

func (it *BaseIterator[T]) bufLen() int {
	return len(it.items)
}

func (it *BaseIterator[T]) setBuf(buf []T) {
	it.items = buf
}

func (it *BaseIterator[T]) Next() (T, error) {
	var item T

	if err := it.nextFunc(); err != nil {
		return item, err
	}

	item = it.items[0]
	it.prevItems = append(it.prevItems, item)
	it.items = it.items[1:]
	return item, nil
}
