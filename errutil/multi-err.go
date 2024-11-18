package errutil

import (
	"errors"
)

type MultiErr struct {
	errs []error
}

func NewMultiErr() *MultiErr {
	return &MultiErr{
		errs: make([]error, 0),
	}
}

func (e *MultiErr) IsError() bool {
	return len(e.errs) > 0
}
func (e *MultiErr) Add(errs ...error) {
	if len(errs) == 0 {
		return
	}
	if errs[0] == nil {
		return
	}
	e.errs = append(e.errs, errs...)
}
func (e *MultiErr) Err() (errs error) {
	if len(e.errs) == 0 {
		return nil
	}
	return errors.Join(e.errs...)
}