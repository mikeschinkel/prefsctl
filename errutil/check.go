package errutil

import (
	"errors"
	"testing"
)

// ErrorCheckFails returns true if the "got" error passed is either nil or does
// not match the wanted error — want indicated by a non-nil errWanted — or true
// if error not wanted and the "got" error is not nil. This function is designed
// to be used in an if statement to return from a test without lots of
// boilerplate required.
func ErrorCheckFails(t *testing.T, testName string, errWanted, errGot error) (failed bool) {
	t.Helper()
	switch {
	case errWanted == nil && errGot != nil:
		t.Errorf("%s: did not want error, got '%v'", testName, CleanErrString(errGot))
		failed = true
	case errWanted != nil:
		switch {
		case errGot == nil:
			t.Errorf("%s: wanted error, got nil", testName)
		case !errors.Is(errGot, errWanted):
			t.Errorf("%s: wanted error but got wrong error:\n\tExpected: %s\n\tReceived: %s",
				testName,
				CleanErrString(errWanted),
				CleanErrString(errGot),
			)
		}
		failed = true
	}
	return failed
}
