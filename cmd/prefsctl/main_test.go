package main

import (
	"os"
	"testing"

	"github.com/mikeschinkel/prefsctl/testutil"
)

// TestMain is the entry point for testing, allowing you to run setup and
// teardown code.
func TestMain(m *testing.M) {
	tm := testutil.NewTestMain()
	tm.OnSetupFunc = func(tm *testutil.TestMain) error {
		return nil
	}
	tm.OnTeardownFunc = func(tm *testutil.TestMain) error {
		return nil
	}
	exitCode := tm.Run(m)
	os.Exit(exitCode)
}

type cmd []string

func TestGetDefaults(t *testing.T) {
	c4t := testutil.GetCurrentContextForTests().Reset(t)
	testutil.RunCLICommandTestsWithSuccessFunc(c4t, nil, []testutil.CLICommandTest{
		{
			Name:        "Get Defaults Go code",
			Args:        cmd{"get", "defaults", "-o", "go"},
			ErrExpected: false,
		},
	})
}
