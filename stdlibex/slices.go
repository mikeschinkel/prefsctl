package stdlibex

import (
	"cmp"
	"slices"
)

//goland:noinspection GoUnusedExportedFunction

//goland:noinspection GoUnusedExportedFunction

func RemoveElements[S []T, T any](slice S, startPos, count int) S {
	s := slices.Clone(slice)
	return append(s[:startPos], s[startPos+count:]...)
}

func DeleteElement[T any](s []T, pos int) []T {
	s = slices.Clone(s)
	lastPos := len(s) - 1
	switch {
	case pos == lastPos:
		s = s[:pos]
	case pos < lastPos:
		s = append(s[:pos], s[pos+1:]...)
	}
	return s
}

func SpliceElements[S []T, T any](slice S, startPos, count int, ele ...T) S {
	s := slices.Clone(slice)
	if startPos >= 0 {
		s = append(s[:startPos], ele...)
	}
	if startPos+count <= len(slice) {
		s = append(s, slices.Clone(slice[startPos+count:])...)
	}
	return s
}

//// Element returns an element of a slice by index,
//// or a default value it index does not exist
//func Element[S []T, T any](s S, i int, d T) (e T) {
//	e = d
//	if i < len(s) {
//		e = s[i]
//	}
//	return e
//}

//goland:noinspection GoUnusedExportedFunction
func RemoveElementFunc[T any](slice []T, isElementFn func(T) bool) []T {
	i := 0
	slice = slices.Clone(slice)
	for i < len(slice) {
		if !isElementFn(slice[i]) {
			i++
			continue
		}
		switch {
		case i == 0:
			slice = slice[1:]
		case i == len(slice)-1:
			slice = slice[:len(slice)-1]
		default:
			slice = append(slice[:i], slice[i+1:]...)
		}
		break
	}
	return slice
}

// ContainsSameElements Compares two 'value' arrays — 'value' meaning
// containing no elements that are references such as pointers, slices, maps or
// channels — and return true if they have all the same elements and no more,
// even if the elements are in different positions in the slice.
//
//goland:noinspection GoUnusedExportedFunction
func ContainsSameElements[E1, E2 any, C comparable](s1 []E1, fn1 func(E1) C, s2 []E2, fn2 func(E2) C) (same bool) {
	var m map[C]struct{}
	if len(s1) != len(s2) {
		goto end
	}
	m = make(map[C]struct{}, len(s1))
	for _, e1 := range slices.Clone(s1) {
		m[fn1(e1)] = struct{}{}
	}
	for _, e2 := range slices.Clone(s2) {
		if _, ok := m[fn2(e2)]; !ok {
			goto end
		}
	}
	same = true
end:
	return same
}

// GetMissingElements Compares two 'value' arrays — 'value' meaning
// containing no elements that are references such as pointers, slices, maps or
// channels — and returns the elements missing in s2 that exist in s1.
//
//goland:noinspection GoUnusedExportedFunction
func GetMissingElements[E1, E2 any, C comparable](s1 []E1, fn1 func(E1) C, s2 []E2, fn2 func(E2) C) (missing []C) {
	var m map[C]struct{}
	if len(s1) != len(s2) {
		goto end
	}
	m = make(map[C]struct{}, len(s1))
	for _, e1 := range s1 {
		m[fn1(e1)] = struct{}{}
	}
	missing = make([]C, 0)
	for _, e2 := range s2 {
		key := fn2(e2)
		if _, ok := m[key]; ok {
			delete(m, key)
			goto end
		}
		missing = append(missing, key)
	}
end:
	return missing
}

//goland:noinspection GoUnusedExportedFunction
func GetElementFunc[E any](slice []E, isElementFn func(E) bool) (ele E, found bool) {
	for _, e := range slice {
		if isElementFn(e) {
			ele = e
			found = true
			break
		}
	}
	return ele, found
}

// DiffSlicesFunc takes two slices and collects any elements FOUND in the 1st
// slice that are NOT IN the 2nd slice, e.g.
//
//	DiffSlicesFunc([]int{0,2,4,6,8},[]int{1,2,3,4,5,6,7},fn) => []int{0,8}
//
// Note the example uses scalars but this is typically to be used for slices of objects
func DiffSlicesFunc[S ~[]E, E any](s1, s2 S, idFunc func(E) string) (diff S) {
	diff = make(S, 0, max(len(s1), len(s2)))
	m1 := make(map[string]struct{})
	for _, e2 := range s2 {
		m1[idFunc(e2)] = struct{}{}
	}
	for _, e1 := range s1 {
		if _, ok := m1[idFunc(e1)]; !ok {
			diff = append(diff, e1)
		}
	}
	return diff
}

// DiffSlices takes two slices and collects any elements FOUND in the 1st
// slice that are NOT IN the 2nd slice, e.g.
//
//	DiffSlices([]int{0,2,4,6,8},[]int{1,2,3,4,5,6,7},fn) => []int{0,8}
func DiffSlices[S ~[]E, E comparable](s1, s2 S) (diff S) {
	diff = make(S, 0, max(len(s1), len(s2)))
	m1 := make(map[E]struct{})
	for _, e2 := range s2 {
		m1[e2] = struct{}{}
	}
	for _, e1 := range s1 {
		if _, ok := m1[e1]; !ok {
			diff = append(diff, e1)
		}
	}
	return diff
}

// MergeSlicesFunc takes two slices of the same element type and merges them, removing any duplicates
//
//	MergeSlicesFunc([]int{0,2,4,6,8},[]int{1,2,3,4,5,6,7},fn) => []int{0,1,2,3,4,5,6,7,8}
//
// Note the example uses scalars but this is typically to be used for slices of objects
func MergeSlicesFunc[S ~[]E, E any](s1, s2 S, cmpFunc func(E, E) int) S {
	return DedupSlicesFunc(slices.Concat(s1, s2), cmpFunc)
}

// DedupSlicesFunc removes duplicate slices based on the return value of cmpFunc
//
//	DedupSlicesFunc([]int{0,2,4,2,3,7,3,2,8},fn) => []int{0,2,4,3,7,8}
//
// Note the example uses scalars but this is typically to be used for slices of objects
func DedupSlicesFunc[S ~[]E, E any](s S, cmpFunc func(E, E) int) S {
	slices.SortFunc(s, cmpFunc)
	return slices.CompactFunc(s, func(e1, e2 E) bool {
		return cmpFunc(e1, e2) == 0
	})
}

// MergeSlices takes two slices of the same element type and merges them, removing any duplicates
//
//	MergeSlices([]int{0,2,4,6,8},[]int{1,2,3,4,5,6,7},fn) => []int{0,1,2,3,4,5,6,7,8}
func MergeSlices[S ~[]E, E cmp.Ordered](s1, s2 S) S {
	s := slices.Concat(s1, s2)
	slices.Sort(s)
	return slices.Compact(s)
}

// FilterSlices returns the slice with all non-zero elements removed.
//
//	FilterSlices([]string{"A", "", "CNAME"}) => []string{"A", "CNAME"}
func FilterSlices[E comparable](s []E) []E {
	var zero E
	filtered := make([]E, 0)
	for _, e := range s {
		if e == zero {
			continue
		}
		filtered = append(filtered, e)
	}
	return filtered
}
