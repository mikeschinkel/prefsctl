package testutil

import (
	"testing"
)

var _ Test = (*TestCommon)(nil)

type TestCommon struct {
	Name   string
	Reason string
	Status int
	Output string
	LogOut []string
	T      *testing.T
}

func NewTestCommon(c4t ContextForTests) *TestCommon {
	return &TestCommon{
		T:      c4t.TestingT(),
		LogOut: []string{},
		Output: "",
		Status: 0,
	}
}
func (c *TestCommon) Common() *TestCommon {
	return c
}
