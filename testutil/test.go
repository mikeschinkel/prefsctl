package testutil

import (
	"bytes"
	"net/http/httptest"

	"github.com/mikeschinkel/prefsctl/logging"
)

type Test interface {
	Common() *TestCommon
}

var initializers []InitializerFunc

type InitializerFunc func(ContextForTests)

func AddInitializer(initFunc InitializerFunc) {
	initializers = append(initializers, initFunc)
}

func InitializeTest(c4t ContextForTests, test Test) {
	c4t.Helper()
	logging.TestLogOutput = bytes.Buffer{}
	c4t.Update(ContextForTestArgs{
		Test:     test,
		Recorder: httptest.NewRecorder(),
	})
	for _, initFunc := range initializers {
		initFunc(c4t)
	}
}
