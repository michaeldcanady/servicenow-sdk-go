package iterator

// PageInfo contains information about an iterator's paging state.
type PageInfo struct {
	// Token is the token used to retrieve the next page of items from the
	// API. You may set Token immediately after creating an iterator to
	// begin iteration at a particular point. If Token is the empty string,
	// the iterator will begin with the first eligible item.
	//
	// The result of setting Token after the first call to Next is undefined.
	//
	// After the underlying API method is called to retrieve a page of items,
	// Token is set to the next-page token in the response.
	Token string

	// MaxSize is the maximum number of items returned by a call to the API.
	// Set MaxSize as a hint to optimize the buffering behavior of the iterator.
	// If zero, the page size is determined by the underlying service.
	//
	// Use Pager to retrieve a page of a specific, exact size.
	MaxSize int

	// The error state of the iterator. Manipulated by PageInfo.next and Pager.
	// This is a latch: it starts as nil, and once set should never change.
	err error

	// If true, no more calls to fetch should be made. Set to true when fetch
	// returns an empty page token. The iterator is Done when this is true AND
	// the buffer is empty.
	atEnd bool

	// Function that fetches a page from the underlying service. It should pass
	// the pageSize and pageToken arguments to the service, fill the buffer
	// with the results from the call, and return the next-page token returned
	// by the service. The function must not remove any existing items from the
	// buffer. If the underlying RPC takes an int32 page size, pageSize should
	// be silently truncated.
	fetch func(pageSize int, pageToken string) (nextPageToken string, err error)

	// Function that returns the number of currently buffered items.
	bufLen func() int

	// Function that returns the buffer, after setting the buffer variable to nil.
	takeBuf func() interface{}

	// Set to true on first call to PageInfo.next or Pager.NextPage. Used to check
	// for calls to both Next and NextPage with the same iterator.
	nextCalled, nextPageCalled bool
}

// newPageInfo creates and returns a PageInfo and a next func. If an iterator can
// support paging, its iterator-creating method should call this. Each time the
// iterator's Next is called, it should call the returned next fn to determine
// whether a next item exists, and if so it should pop an item from the buffer.
//
// The fetch, bufLen and takeBuf arguments provide access to the iterator's
// internal slice of buffered items. They behave as described in PageInfo, above.
//
// The return value is the PageInfo.next method bound to the returned PageInfo value.
// (Returning it avoids exporting PageInfo.next.)
//
// Note: the returned PageInfo and next fn do not remove items from the buffer.
// It is up to the iterator using these to remove items from the buffer:
// typically by performing a pop in its Next. If items are not removed from the
// buffer, memory may grow unbounded.
func newPageInfo(fetch func(int, string) (string, error), bufLen func() int, takeBuf func() interface{}) (pi *PageInfo, next func() error) {
	pi = &PageInfo{
		fetch:   fetch,
		bufLen:  bufLen,
		takeBuf: takeBuf,
	}
	return pi, pi.next
}

// Remaining returns the number of items available before the iterator makes another API call.
func (pi *PageInfo) Remaining() int { return pi.bufLen() }

// next provides support for an iterator's Next function. An iterator's Next
// should return the error returned by next if non-nil; else it can assume
// there is at least one item in its buffer, and it should return that item and
// remove it from the buffer.
func (pi *PageInfo) next() error {
	pi.nextCalled = true
	if pi.err != nil { // Once we get an error, always return it.
		// TODO(jba): fix so users can retry on transient errors? Probably not worth it.
		return pi.err
	}
	if pi.nextPageCalled {
		pi.err = errMixed
		return pi.err
	}
	// Loop until we get some items or reach the end.
	for pi.bufLen() == 0 && !pi.atEnd {
		if err := pi.fill(pi.MaxSize); err != nil {
			pi.err = err
			return pi.err
		}
		if pi.Token == "" {
			pi.atEnd = true
		}
	}
	// Either the buffer is non-empty or pi.atEnd is true (or both).
	if pi.bufLen() == 0 {
		// The buffer is empty and pi.atEnd is true, i.e. the service has no
		// more items.
		pi.err = errDone
	}
	return pi.err
}

// Call the service to fill the buffer, using size and pi.Token. Set pi.Token to the
// next-page token returned by the call.
// If fill returns a non-nil error, the buffer will be empty.
func (pi *PageInfo) fill(size int) error {
	tok, err := pi.fetch(size, pi.Token)
	if err != nil {
		pi.takeBuf() // clear the buffer
		return err
	}
	pi.Token = tok
	return nil
}
