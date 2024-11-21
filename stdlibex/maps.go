package stdlibex

// MapKeys returns the keys of the map m as a slice.
// The keys will be an indeterminate order.
//
//goland:noinspection GoUnusedExportedFunction
func MapKeys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// MapValues returns the values of the map m as a slice.
// The keys will be an indeterminate order.
//
//goland:noinspection GoUnusedExportedFunction
func MapValues[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// MakeKeysMap takes a slice and returns a map[comparable]struct{} where each
// element from the slice becomes a map key.
func MakeKeysMap[K comparable](s []K) map[K]struct{} {
	m := make(map[K]struct{}, len(s))
	for _, k := range s {
		m[k] = struct{}{}
	}
	return m
}

// MapDefault takes a map and its comparable key and returns the key's associated
// value if it exists but if not returns the default value passed as a 3rd
// parameter. element from the slice becomes a map key.
//
//goland:noinspection GoUnusedExportedFunction
func MapDefault[K comparable, T any](m map[K]T, key K, def T) T {
	v, ok := m[key]
	if !ok {
		v = def
	}
	return v
}
