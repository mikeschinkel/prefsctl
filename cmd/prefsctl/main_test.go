package main_test

import (
	"os"
	"testing"

	"github.com/mikeschinkel/prefsctl/cmd/prefsctl/cmds"
	"github.com/mikeschinkel/prefsctl/cobrautil"
)

// TestMain is the entry point for testing, allowing you to run setup and
// teardown code.
func TestMain(m *testing.M) {
	tm := cobrautil.NewTestMain()
	tm.OnSetupFunc = func(tm *cobrautil.TestMain) error {
		return nil
	}
	tm.OnTeardownFunc = func(tm *cobrautil.TestMain) error {
		return nil
	}
	tm.RootCmd = cmds.RootCmd
	exitCode := tm.Run(m)
	os.Exit(exitCode)
}

type cmd []string

func TestGetDefaults(t *testing.T) {
	c4t := cobrautil.GetCurrentContextForTests().Reset(t)

	cobrautil.RunCLICommandTests(c4t, []cobrautil.CLICommandTest{
		{
			Name:          "Get Defaults Go code",
			Args:          cmd{"get", "defaults", "-o", "go"},
			SuccessResult: "TODO",
		},
	})
}
