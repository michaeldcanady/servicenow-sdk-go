package iterator

type HasNextLink interface {
	GetNextLink() (*string, error)
}
