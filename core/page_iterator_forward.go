package core

// ForwardPageIterator[T, C] represents iterator for forward iteration of pages
type ForwardPageIterator[T any, C CollectionResponse[T]] struct {
	PageIterator[T, C]
	pauseIndex int
}

// NewForwardPageIterator[T, C] Creates a new ForwardPageIterator object.
func NewForwardPageIterator[T any, C CollectionResponse[T]](iterator PageIterator[T, C]) *ForwardPageIterator[T, C] {
	return &ForwardPageIterator[T, C]{
		PageIterator: iterator,
		pauseIndex:   0,
	}
}

// enumerate iterates through the items on the current page and invokes the callback.
func (pI *ForwardPageIterator[T, C]) enumerate(callback func(item *T) bool) bool {
	keepIterating := true

	pageItems := pI.Current().Result
	if pageItems == nil {
		return false
	}

	for i := pI.pauseIndex; i < len(pageItems); i++ {
		keepIterating = callback(pageItems[i])

		if !keepIterating {
			break
		}

		pI.pauseIndex = i + 1
	}

	return keepIterating
}

// Iterate iterates through pages and invokes the provided callback for each page item.
func (pI *ForwardPageIterator[T, C]) Iterate(callback func(pageItem *T) bool) error {
	if callback == nil {
		return ErrNilCallback
	}

	for {
		keepIterating := pI.enumerate(callback)

		if !keepIterating {
			// Callback returned false, stop iterating through pages.
			return nil
		}

		if pI.Current().NextPageLink == "" {
			// NextPageLink is empty, stop iterating through pages.
			return nil
		}

		nextPage, err := pI.Next()
		if err != nil {
			return err
		}

		pI.SetCurret(nextPage)
		pI.pauseIndex = 0
	}
}

// next fetches the next page of results.
func (pI *ForwardPageIterator[T, C]) Next() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.Current().NextPageLink)
}
