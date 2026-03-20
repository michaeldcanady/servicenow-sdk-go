package utils

// Mutator represents a function that transforms a value.
type Mutator[T, S any] func(input T) (S, error)

func NoOpMutator[T any](input T) (T, error) {
	return input, nil
}
