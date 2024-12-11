package mapsutils

import (
	"cmp"
	"iter"
	"slices"
)

func KeysSorted[Map ~map[O]V, O cmp.Ordered, V any](m Map) iter.Seq2[O, V] {
	return func(yield func(O, V) bool) {
		sorted := make([]O, len(m))
		i := 0
		for k := range m {
			sorted[i] = k
			i++
		}
		slices.Sort(sorted)
		for _, k := range sorted {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}
