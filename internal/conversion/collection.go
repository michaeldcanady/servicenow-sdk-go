package conversion

import "fmt"

// CastCollection converts a slice of one type to another.
func CastCollection[T, R any](collection []T) ([]R, error) {
	var err error
	newCollection := CollectionApply(collection, func(in T) (R, bool) {
		out, ok := any(in).(R)
		if !ok {
			var emptyR R
			err = fmt.Errorf("item not %T", emptyR)
			return emptyR, false
		}
		return out, true
	})
	if err != nil {
		return nil, err
	}

	return newCollection, nil
}

// CollectionApply applies a mutator function to each item in a slice.
func CollectionApply[T, R any](collection []T, mutator func(in T) (R, bool)) []R {
	outputs := make([]R, len(collection))
	for i, item := range collection {
		output, ok := mutator(item)
		if !ok {
			break
		}

		outputs[i] = output
	}
	return outputs
}
