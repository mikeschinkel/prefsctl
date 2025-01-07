package testutil

import (
	"context"
)

type Context = context.Context

type TestSetter interface {
	SetTest(Test)
}
