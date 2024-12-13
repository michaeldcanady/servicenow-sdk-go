package iterator

import (
	"errors"
	"fmt"
	"reflect"
)

// Pager supports retrieving iterator items a page at a time.
type Pager struct {
	pageInfo *PageInfo
	pageSize int
}

// NewPager returns a pager that uses iter. Calls to its NextPage method will
// obtain exactly pageSize items, unless fewer remain. The pageToken argument
// indicates where to start the iteration. Pass the empty string to start at
// the beginning, or pass a token retrieved from a call to Pager.NextPage.
//
// If you use an iterator with a Pager, you must not call Next on the iterator.
func NewPager(iter Pageable, pageSize int, pageToken string) *Pager {
	p := &Pager{
		pageInfo: iter.PageInfo(),
		pageSize: pageSize,
	}
	p.pageInfo.Token = pageToken
	if pageSize <= 0 {
		p.pageInfo.err = errors.New("iterator: page size must be positive")
	}
	return p
}

// NextPage retrieves a sequence of items from the iterator and appends them
// to slicep, which must be a pointer to a slice of the iterator's item type.
// Exactly p.pageSize items will be appended, unless fewer remain.
//
// The first return value is the page token to use for the next page of items.
// If empty, there are no more pages. Aside from checking for the end of the
// iteration, the returned page token is only needed if the iteration is to be
// resumed a later time, in another context (possibly another process).
//
// The second return value is non-nil if an error occurred. It will never be
// the special iterator sentinel value Done. To recognize the end of the
// iteration, compare nextPageToken to the empty string.
//
// It is possible for NextPage to return a single zero-length page along with
// an empty page token when there are no more items in the iteration.
func (p *Pager) NextPage(slicep interface{}) (nextPageToken string, err error) {
	p.pageInfo.nextPageCalled = true
	if p.pageInfo.err != nil {
		return "", p.pageInfo.err
	}
	if p.pageInfo.nextCalled {
		p.pageInfo.err = errMixed
		return "", p.pageInfo.err
	}
	if p.pageInfo.bufLen() > 0 {
		return "", errors.New("must call NextPage with an empty buffer")
	}
	// The buffer must be empty here, so takeBuf is a no-op. We call it just to get
	// the buffer's type.
	wantSliceType := reflect.PointerTo(reflect.ValueOf(p.pageInfo.takeBuf()).Type())
	if slicep == nil {
		return "", errors.New("nil passed to Pager.NextPage")
	}
	vslicep := reflect.ValueOf(slicep)
	if vslicep.Type() != wantSliceType {
		return "", fmt.Errorf("slicep should be of type %s, got %T", wantSliceType, slicep)
	}
	for p.pageInfo.bufLen() < p.pageSize {
		if err := p.pageInfo.fill(p.pageSize - p.pageInfo.bufLen()); err != nil {
			p.pageInfo.err = err
			return "", p.pageInfo.err
		}
		if p.pageInfo.Token == "" {
			break
		}
	}
	e := vslicep.Elem()
	e.Set(reflect.AppendSlice(e, reflect.ValueOf(p.pageInfo.takeBuf())))
	return p.pageInfo.Token, nil
}
