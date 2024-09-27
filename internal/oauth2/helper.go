package oauth2

type FilterFunc[T any] func(T) bool

func Filter[T any](slice []any, filterFunc FilterFunc[T]) []T {
	var result []T
	for _, item := range slice {
		elem, ok := item.(T)
		if !ok {
			continue
		}
		if filterFunc(elem) {
			result = append(result, elem)
		}
	}
	return result
}
