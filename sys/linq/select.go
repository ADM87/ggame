package linq

func Select[T any, TResult any](source []T, selector func(T) TResult) []TResult {
	result := make([]TResult, len(source))
	for i, item := range source {
		result[i] = selector(item)
	}
	return result
}

func SelectMany[T any, TResult any](source []T, selector func(T) []TResult) []TResult {
	var result []TResult
	for _, item := range source {
		result = append(result, selector(item)...)
	}
	return result
}
