package collections

type Array []interface{}

func Any(elements []bool) bool {
	for _, e := range elements {
		if e {
			return true
		}
	}

	return false
}

func All(elements []bool) bool {
	for _, e := range elements {
		if !e {
			return false
		}
	}

	return true
}

func Map[K, V interface{}](elements []K, mapper func(K) V) []V {
	results := []V{}
	for _, e := range elements {
		results = append(results, mapper(e))
	}
	return results
}

func Reduce[K, V interface{}](init V, elements []K, reducer func(V, K) V) V {
	result := init
	for _, e := range elements {
		result = reducer(result, e)
	}
	return result
}

func Filter[K interface{}](elements []K, filter func(K) bool) []K {
	results := []K{}
	for _, e := range elements {
		if filter(e) {
			results = append(results, e)
		}
	}

	return results
}
