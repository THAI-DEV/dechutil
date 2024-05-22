package dechutil

func CloneMap[K comparable, V any](src map[K]V) map[K]V {
	clone := make(map[K]V, len(src))
	for k, v := range src {
		clone[k] = v
	}

	return clone
}
