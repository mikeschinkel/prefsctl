package testutil

import (
	"context"
	"net/http/httptest"
	"testing"
)

type ContextForTests interface {
	context.Context
	testing.TB
	Depth() int
	IncrDepth() ContextForTests
	DecrDepth() ContextForTests
	Debugf(format string, args ...any) ContextForTests
	Debug(args ...any) ContextForTests
	SetTest(Test)
	Reset(*testing.T) ContextForTests
	SetTestingT(*testing.T)
	TestingT() *testing.T
	Run(string, func(t *testing.T)) bool
	SetContext(Context)
	Context() Context
	Test() Test
	Warnf(string, ...any) ContextForTests
	Warn(args ...any) ContextForTests
	Common() *TestCommon
	Update(ContextForTestArgs)
	Recorder() *httptest.ResponseRecorder
	ConfigFilepath() string
}

var _ ContextForTests = (*contextForTests)(nil)

type contextForTests struct {
	*testing.T
	context  context.Context
	test     Test
	common   *TestCommon
	recorder *httptest.ResponseRecorder
	depth    int
}

func (c4t *contextForTests) ConfigFilepath() string {
	return "./config.json"
}

func (c4t *contextForTests) Depth() int {
	return c4t.depth
}
func (c4t *contextForTests) IncrDepth() ContextForTests {
	c4t.depth++
	return c4t
}
func (c4t *contextForTests) DecrDepth() ContextForTests {
	c4t.depth--
	return c4t
}

func (c4t *contextForTests) Run(name string, fn func(t *testing.T)) bool {
	return c4t.T.Run(name, fn)
}

func (c4t *contextForTests) Reset(t *testing.T) ContextForTests {
	c4t.SetTestingT(t)
	return c4t
}
func (c4t *contextForTests) SetTestingT(t *testing.T) {
	c4t.T = t
}

func (c4t *contextForTests) Recorder() *httptest.ResponseRecorder {
	return c4t.recorder
}

func (c4t *contextForTests) TestingT() *testing.T {
	return c4t.T
}

func (c4t *contextForTests) Common() *TestCommon {
	return c4t.common
}

func (c4t *contextForTests) Test() Test {
	return c4t.test
}

func (c4t *contextForTests) SetTest(test Test) {
	c4t.test = test
}

func (c4t *contextForTests) Context() Context {
	return c4t.context
}

func (c4t *contextForTests) SetContext(c Context) {
	c4t.context = c
}

func (c4t *contextForTests) Done() <-chan struct{} {
	return c4t.Context().Done()
}

func (c4t *contextForTests) Err() error {
	return c4t.Context().Err()
}

func (c4t *contextForTests) Value(key any) any {
	return c4t.Context().Value(key)
}

type ContextForTestArgs struct {
	T        *testing.T
	Context  context.Context
	Test     Test
	Common   *TestCommon
	Recorder *httptest.ResponseRecorder
}

func NewContextForTests(t *testing.T) ContextForTests {
	return &contextForTests{
		T:        t,
		context:  context.Background(),
		recorder: httptest.NewRecorder(),
	}
}

func (c4t *contextForTests) Update(args ContextForTestArgs) {
	if args.Context != nil {
		c4t.SetContext(args.Context)
	}
	if args.Test != nil {
		c4t.test = args.Test
		c4t.common = c4t.test.Common()
	}
	if args.Common != nil {
		c4t.common = args.Common
	}
	if args.T != nil {
		c4t.T = args.T
		c4t.common.T = c4t.T
	}
	if args.Recorder != nil {
		c4t.recorder = args.Recorder
	}
	if c4t.test.Common().T == nil {
		c4t.test.Common().T = c4t.T
	}
	//if setter, ok := c4t.client.(TestSetter); ok {
	//	setter.SetTest(c4t.test)
	//}
}

func (c4t *contextForTests) Debugf(format string, args ...any) ContextForTests {
	Depth(c4t.depth).Debugf(format, args...)
	return c4t
}
func (c4t *contextForTests) Debug(args ...any) ContextForTests {
	Depth(c4t.depth).Debug(args...)
	return c4t
}

func (c4t *contextForTests) Warnf(format string, args ...any) ContextForTests {
	Depth(c4t.depth).Warnf(format, args...)
	return c4t
}
func (c4t *contextForTests) Warn(args ...any) ContextForTests {
	Depth(c4t.depth).Warn(args...)
	return c4t
}
