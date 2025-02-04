package sliceconv

// Minimize removes any unused capacity from a slice
func Minimize[T any](s []T) []T {
	minimized := make([]T, len(s))
	copy(minimized, s)
	return minimized
}
