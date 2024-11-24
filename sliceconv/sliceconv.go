package sliceconv

import (
	"fmt"

	"github.com/mikeschinkel/prefsctl/errutil"
)

// ToAny converts a slice of whatever type to a slice of any
//
//goland:noinspection GoUnusedExportedFunction
func ToAny[ST []T, T any](slice ST) []any {
	items := make([]any, len(slice))
	for i, item := range slice {
		items[i] = item
	}
	return items
}

// ToStrings converts a slice of whatever type to a slice of string
func ToStrings[ST []T, T any](slice ST) []string {
	items := make([]string, len(slice))
	for i, item := range slice {
		items[i] = fmt.Sprintf("%v", item)
	}
	return items
}

// ToPtrs converts a slice of any type to a slice of pointers to that type
func ToPtrs[E any](from []E) []*E {
	items := make([]*E, len(from))
	for i, item := range from {
		items[i] = &item
	}
	return items
}

// Func converts a slice of one type to a slice of another type
func Func[F, T any](from []F, fn func(F) T) []T {
	items := make([]T, len(from))
	for i, item := range from {
		items[i] = fn(item)
	}
	return items
}

// ToMapFunc converts a slice of a map using a func
//
//goland:noinspection GoUnusedExportedFunction
func ToMapFunc[K comparable, V, T any](slice []T, fn func(T) (bool, K, V, error)) (map[K]V, error) {
	var errs errutil.MultiErr
	items := make(map[K]V)
	for _, item := range slice {
		include, k, v, err := fn(item)
		if !include {
			continue
		}
		if err != nil {
			errs.Add(err)
		}
		items[k] = v
	}
	return items, errs.Err()
}

// ToStringsFunc converts a slice of a slice of strings using a func
func ToStringsFunc[T any](slice []T, fn func(T) (bool, string, error)) ([]string, error) {
	var errs errutil.MultiErr
	items := make([]string, len(slice))
	for _, item := range slice {
		include, s, err := fn(item)
		if !include {
			continue
		}
		if err != nil {
			errs.Add(err)
		}
		items = append(items, s)
	}
	return items, errs.Err()
}
