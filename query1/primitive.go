package query

type Primitive interface {
	Numeric | ~string
}
