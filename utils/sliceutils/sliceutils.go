package sliceutils

func Contains[K comparable](slice []K, value K) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func RemoveDuplicates[K comparable](slice []K) []K {
	noDuplicates := make([]K, 0)

	for _, v := range slice {
		if !Contains(noDuplicates, v) {
			noDuplicates = append(noDuplicates, v)
		}
	}

	return noDuplicates
}

func Map[K any, V any](slice []K, mapFunction func(K) V) []V {
	var mapped []V
	for _, i := range slice {
		mapped = append(mapped, mapFunction(i))
	}

	return mapped
}

func Concatenate[K any](slices ...[]K) []K {
	var slice []K

	for _, s := range slices {
		slice = append(slice, s...)
	}

	return slice
}

func Filter[K any](slice []K, filterFunction func(K) bool) []K {
	var s []K

	for _, v := range slice {
		if filterFunction(v) {
			s = append(s, v)
		}
	}

	return s
}
