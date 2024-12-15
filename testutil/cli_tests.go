package testutil

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/stdlibex"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type BeforeCLITestsFunc func(ContextForTests) error
type BeforeCLITestFunc func(ContextForTests, CLICommandTest) error
type AfterCLITestFunc func(ContextForTests, CLICommandTest, *cobrautil.Outcome) error
type AfterCLITestsFunc func(ContextForTests) error
type CheckTestErrorsFunc func(ContextForTests, CLICommandTest, *error) error

type OnCLITestSuccessFunc func(*cobrautil.Outcome) any
type CLICommandTest struct {
	Name                string   // Test case name
	Args                []string // CLI arguments
	ErrExpected         bool     // Expected error
	ErrMessage          string
	BeforeTestFunc      BeforeCLITestFunc
	CheckTestErrorsFunc CheckTestErrorsFunc
	AfterTestFunc       AfterCLITestFunc
}

var BeforeTests BeforeCLITestsFunc = func(ContextForTests) error {
	return nil
}
var AfterTests AfterCLITestsFunc = func(ContextForTests) error {
	return nil
}
var TestNameLogArg = "test_name"

func RunCLICommandTests(c4t ContextForTests, tests []CLICommandTest) any {
	return RunCLICommandTestsWithSuccessFunc(c4t, nil, tests)
}
func RunCLICommandTestsWithSuccessFunc(c4t ContextForTests, successFunc OnCLITestSuccessFunc, tests []CLICommandTest) any {
	var errs errutil.MultiErr
	results := make([]any, 0)
	err := BeforeTests(c4t)
	if err != nil {
		c4t.Error("Failed to run pre-tests func: %v", err)
		goto end
	}
	for _, tt := range tests {
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
		c4t.Error("Failed to run pre-tests func: %v", err)
		goto end
	}
end:
	return results
}

func RunCLICommandTest(c4t ContextForTests, successFunc OnCLITestSuccessFunc, test CLICommandTest) (output any, err error) {
	var cli *cobrautil.CLI
	var outcome *cobrautil.Outcome

	ctx := cobrautil.DefaultContext()
	cfg := cobrautil.NewConfig(&cobrautil.ConfigArgs{
		Filepath: stdlibex.Must(cobrautil.ConfigFilepath()),
	})
	err = cfg.Initialize(ctx)
	if err != nil {
		c4t.Fatalf("Failed to initialize CLI config: %v", err)
	}
	if runBeforeTestFails(c4t, test, &err) {
		goto end
	}
	cli = cobrautil.NewCLI(cfg, test.Args)
	err = cli.Initialize(ctx)
	outcome, err = cli.Execute(ctx, cli.Args)
	if runTestErrorChecksFail(c4t, test, &err) {
		goto end
	}
	if runAfterTestFails(c4t, test, outcome, &err) {
		goto end
	}
	if successFunc != nil {
		output = onCLITestSuccess(test, outcome, successFunc)
		goto end
	}
end:
	return output, err
}

func runTestErrorChecksFail(c4t ContextForTests, test CLICommandTest, err *error) bool {
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
		titler := cases.Title(language.English)
		*err = errs.Err()
		c4t.Errorf(titler.String((*err).Error()))
	}
	return *err != nil
}

func onCLITestSuccess(test CLICommandTest, outcome *cobrautil.Outcome, onSuccessFunc OnCLITestSuccessFunc) any {
	if outcome.WasError() {
		return nil
	}
	if test.ErrExpected {
		return nil
	}
	return onSuccessFunc(outcome)
}

func runBeforeTestFails(c4t ContextForTests, test CLICommandTest, errPtr *error) bool {
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

func runAfterTestFails(c4t ContextForTests, test CLICommandTest, outcome *cobrautil.Outcome, errPtr *error) bool {
	var err error

	if test.AfterTestFunc == nil {
		goto end
	}
	err = test.AfterTestFunc(c4t, test, outcome)
	if err != nil {
		errPtr = &err
	}
end:
	return err != nil
}
