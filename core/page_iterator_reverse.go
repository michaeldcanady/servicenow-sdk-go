package core

// ForwardPageIterator[T, C] represents iterator for backward iteration of pages
type ReversePageIterator[T any, C CollectionResponse[T]] struct {
	PageIterator[T, C]
	pauseIndex int
}

// NewReversePageIterator[T, C] Creates a new ReversePageIterator object.
func NewReversePageIterator[T any, C CollectionResponse[T]](iterator PageIterator[T, C]) *ReversePageIterator[T, C] {
	return &ReversePageIterator[T, C]{
		PageIterator: iterator,
		pauseIndex:   0,
	}
}

// Iterate iterates through pages and invokes the provided callback for each page item.
func (pI *ReversePageIterator[T, C]) Reverse(callback func(pageItem *T) bool) error {
	if callback == nil {
		return ErrNilCallback
	}

	for {
		keepIterating := pI.enumerate(callback)

		if !keepIterating {
			// Callback returned false, stop iterating through pages.
			return nil
		}

		if pI.Current().PreviousPageLink == "" {
			// PreviousPageLink is empty, stop iterating through pages.
			return nil
		}

		previousPage, err := pI.Previous()
		if err != nil {
			return err
		}

		pI.SetCurret(previousPage)
		pI.pauseIndex = 0
	}
}

// enumerate iterates through the items on the current page and invokes the callback.
func (pI *ReversePageIterator[T, C]) enumerate(callback func(item *T) bool) bool {
	keepIterating := true

	pageItems := pI.Current().Result
	if pageItems == nil {
		return false
	}

	for i := len(pageItems) - 1; i >= 0; i-- {
		keepIterating = callback(pageItems[i])

		if !keepIterating {
			break
		}

		pI.pauseIndex = i - 1
	}

	return keepIterating
}

// previous fetches the previous page of results.
func (pI *ReversePageIterator[T, C]) Previous() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.Current().PreviousPageLink)
}

// Last fetches the last page of results.
func (pI *ReversePageIterator[T, C]) Last() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.Current().LastPageLink)
}
