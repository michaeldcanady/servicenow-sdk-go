package iterator

// Pageable is implemented by iterators that support paging.
type Pageable interface {
	// PageInfo returns paging information associated with the iterator.
	PageInfo() *PageInfo
}
