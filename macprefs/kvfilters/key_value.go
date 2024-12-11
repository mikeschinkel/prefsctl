package kvfilters

import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
)

type KeyValue interface {
	fmt.Stringer
	Key() Code
	Value() string
	Labels() *Labels
	ErrorInfo() error
	Valid() bool
}

func (g *group) KeyValues() (kvs []KeyValue) {
	return g.keyValues
}

var _ KeyValue = (*keyValue)(nil)

type keyValue struct {
	key     Code
	value   string
	labels  *Labels
	invalid bool
}

func (kv *keyValue) Labels() *Labels {
	return kv.labels
}

func (kv *keyValue) Valid() bool {
	return !kv.invalid
}

func (kv *keyValue) String() string {
	return fmt.Sprintf("%s=%s", kv.key, kv.value)
}

func (kv *keyValue) Key() Code {
	return kv.key
}

func (kv *keyValue) Value() string {
	return kv.value
}
func (kv *keyValue) ErrorInfo() error {
	return errors.Join(
		fmt.Errorf("%s=%s", logargs.KeyLogArg, kv.key),
		fmt.Errorf("%s=%s", logargs.ValueLogArg, kv.value),
	)
}
