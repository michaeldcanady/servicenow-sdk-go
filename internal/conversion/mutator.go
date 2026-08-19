package conversion

type Mutator[T, S any] func(input T) (S, error)
