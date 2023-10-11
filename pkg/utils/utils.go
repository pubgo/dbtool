package utils

func MapKeys[K comparable, V any](mm map[K]V) []K {
	if mm == nil {
		return nil
	}

	var keys []K
	for k := range mm {
		keys = append(keys, k)
	}
	return keys
}
