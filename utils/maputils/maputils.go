package maputils

func GetValueOrDefault[K comparable, V any](_map map[K]V, key K, _default V) V {
	v, contains := _map[key]
	if contains {
		return v
	} else {
		return _default
	}
}
