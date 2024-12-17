//go:build test

package cobrautil

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/mikeschinkel/prefsctl/testutil"
)

const (
	OnErrExitCode = 999
)

var testMain *TestMain

type TestMain struct {
	contextForTests testutil.ContextForTests
	OnSetupFunc     func(*TestMain) error
	OnTeardownFunc  func(*TestMain) error
	RootCmd         Cmd
}

func NewTestMain() *TestMain {
	return &TestMain{}
}

func GetCurrentContextForTests() testutil.ContextForTests {
	return testMain.contextForTests
}

// Run runs setup then tests and then teardown.
func (tm *TestMain) Run(m *testing.M) int {
	var exitCode int
	err := tm.Setup()
	if err != nil {
		tm.Fatalf(err.Error())
	}
	exitCode = m.Run()
	err = tm.Teardown()
	if err != nil {
		tm.Errorf(err.Error())
	}
	return exitCode
}

func (tm *TestMain) Initialize() {
	// Create a ContextForTests but one that has a mostly invalid *testing.T because
	// it was only designed to be instantiated inside of tests vs. in a TestMain.
	tm.contextForTests = testutil.NewContextForTests(&testing.T{})

	disableTestInfo := strings.ToLower(os.Getenv(testutil.DisableDebugInfoVar)) != ""
	if !disableTestInfo {
		testutil.SetDebugOutput(true)
		testutil.Infof("Debug output for testing is ENABLED. To disable it, set %s to 'yes'", testutil.DisableDebugInfoVar)
	}
}

// Setup prepares all child tests of TestMain — those in this `test` package — to
// have a consistent, repeatable environment. It establishes output to be
// discarded (TODO: explain why) and calls TestMain.SetupFunc() if not nil.
func (tm *TestMain) Setup() (err error) {

	// Get a reference to this object so other objects can access it
	testMain = tm

	// Disable testutil.Debugf() output when DISABLE_TEST_INFO is set to "yes"
	// (or anything, actually)
	tm.Initialize()

	c4t := tm.contextForTests

	if tm.RootCmd == nil {
		c4t.Fatalf("RootCmd not set in TestMain")
	}
	SetRootCmd(tm.RootCmd)

	c4t.Debug("Beginning global setup")

	if tm.OnSetupFunc != nil {
		err = tm.OnSetupFunc(tm)
		if err != nil {
			c4t.Debugf("Global setup failed: %#v", err)
			goto end
		}
	}

	c4t.Debug("Global setup complete")
end:
	return err
}

// Teardown allows tearing down of any test scaffolding by calling
// TestMain.TeardownFunc() if not nil.
func (tm *TestMain) Teardown() (err error) {
	c4t := tm.contextForTests
	c4t.Debug("Beginning global teardown")
	if tm.OnTeardownFunc != nil {
		err = tm.OnTeardownFunc(tm)
		if err != nil {
			c4t.Debugf("Global teardown failed: %#v", err)
			goto end
		}
	}
	c4t.Debug("Global teardown complete")
end:
	return err
}

func (tm *TestMain) Errorf(format string, args ...any) {
	fmt.Printf("+++ ERROR: "+format+"\n", args...)
}

func (tm *TestMain) Fatalf(format string, args ...any) {
	fmt.Printf("+++ FATAL: "+format+"\n", args...)
	os.Exit(1)
}
