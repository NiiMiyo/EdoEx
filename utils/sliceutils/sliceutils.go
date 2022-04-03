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
