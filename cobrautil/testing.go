//go:build test

package cobrautil

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/stdlibex"
	"github.com/mikeschinkel/prefsctl/testutil"
)

type ContextForTests = testutil.ContextForTests

type BeforeCLITestsFunc func(ContextForTests) error
type BeforeCLITestFunc func(ContextForTests, CLICommandTest) error
type AfterCLITestFunc func(ContextForTests, CLICommandTest, CmdResult) error
type AfterCLITestsFunc func(ContextForTests) error
type CheckTestErrorsFunc func(ContextForTests, CLICommandTest, *error) error

type OnCLITestSuccessFunc func(CmdResult) any
type CLICommandTest struct {
	Name                string   // Test case name
	Args                []string // CLI arguments
	ErrExpected         bool     // Expected error
	ErrMessage          string
	BeforeTestFunc      BeforeCLITestFunc
	CheckTestErrorsFunc CheckTestErrorsFunc
	AfterTestFunc       AfterCLITestFunc
	Config              Config
	OutWriter           io.Writer
	ErrWriter           io.Writer
}

var BeforeTests BeforeCLITestsFunc = func(testutil.ContextForTests) error {
	return nil
}

var AfterTests AfterCLITestsFunc = func(testutil.ContextForTests) error {
	return nil
}

var TestNameLogArg = "test_name"

func RunCLICommandTests(c4t testutil.ContextForTests, tests []CLICommandTest) any {
	return RunCLICommandTestsWithSuccessFunc(c4t, nil, tests)
}

func RunCLICommandTestsWithSuccessFunc(c4t testutil.ContextForTests, successFunc OnCLITestSuccessFunc, tests []CLICommandTest) any {
	var errs errutil.MultiErr
	results := make([]any, 0)
	err := BeforeTests(c4t)
	if err != nil {
		c4t.Errorf("Failed to run pre-tests func: %v", err)
		goto end
	}
	for _, tt := range tests {
		tt.OutWriter = &bytes.Buffer{}
		tt.ErrWriter = &bytes.Buffer{}
		c4t.TestingT().Run(tt.Name, func(t *testing.T) {
			result, err := RunCLICommandTest(c4t, successFunc, tt)
			if err != nil {
				errs.Add(errors.Join(err, fmt.Errorf("%s=%s", TestNameLogArg, tt.Name)))
			}
			results = append(results, result)
		})
	}
	err = AfterTests(c4t)
	if err != nil {
		c4t.Errorf("Failed to run pre-tests func: %v", err)
		goto end
	}
end:
	return results
}

func RunCLICommandTest(c4t testutil.ContextForTests, successFunc OnCLITestSuccessFunc, test CLICommandTest) (result any, err error) {
	var cmdResult CmdResult

	if test.Config == nil {
		test.Config = NewConfig(&ConfigArgs{
			Filepath: c4t.ConfigFilepath(),
		})
	}

	if runBeforeTestFails(c4t, test, &err) {
		goto end
	}
	cmdResult, err = Execute(rootCmd, ExecuteArgs{
		Config:    test.Config,
		CLIArgs:   test.Args,
		OutWriter: test.OutWriter,
		ErrWriter: test.ErrWriter,
	})
	if runTestErrorChecksFail(c4t, test, &err) {
		goto end
	}
	if runAfterTestFails(c4t, test, cmdResult, &err) {
		goto end
	}
	if successFunc != nil {
		result = onCLITestSuccess(test, cmdResult, successFunc)
		goto end
	}
end:
	return result, err
}

func runTestErrorChecksFail(c4t testutil.ContextForTests, test CLICommandTest, err *error) bool {
	var errs errutil.MultiErr
	var got error
	var want string

	c4t.Helper()

	// Check to see if error was expected and if expectations were met
	switch {
	case *err != nil:
		if !test.ErrExpected {
			errs.Add(fmt.Errorf("error DID occur but was NOT expected"))
		}
	default:
		if test.ErrExpected {
			errs.Add(fmt.Errorf("error did NOT occur but WAS expected"))
		}
	}

	// Check is expected error message matches or not
	got = *err
	want = test.ErrMessage
	switch {
	case got == nil:
		if want != "" {
			errs.Add(fmt.Errorf("expected error message: %v, got: nil error", want))
		}
	case got.Error() == "":
		if want != "" {
			errs.Add(fmt.Errorf("expected error message: %v, got: empty error", want))
		}
	case want == "":
		errs.Add(fmt.Errorf("expected NO error message, got: %v", got.Error()))
	}

	if errs.IsError() {
		*err = errs.Err()
		c4t.Errorf(stdlibex.Title((*err).Error()))
	}
	return *err != nil
}

func onCLITestSuccess(test CLICommandTest, result CmdResult, onSuccessFunc OnCLITestSuccessFunc) any {
	if result.IsErr() {
		return nil
	}
	if test.ErrExpected {
		return nil
	}
	return onSuccessFunc(result)
}

func runBeforeTestFails(c4t testutil.ContextForTests, test CLICommandTest, errPtr *error) bool {
	var err error

	if test.BeforeTestFunc == nil {
		goto end
	}
	err = test.BeforeTestFunc(c4t, test)
	if err != nil {
		errPtr = &err
	}
end:
	return err != nil
}

func runAfterTestFails(c4t testutil.ContextForTests, test CLICommandTest, result CmdResult, errPtr *error) bool {
	var err error

	if test.AfterTestFunc == nil {
		goto end
	}
	err = test.AfterTestFunc(c4t, test, result)
	if err != nil {
		errPtr = &err
	}
end:
	return err != nil
}
