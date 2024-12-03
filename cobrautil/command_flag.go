package cobrautil

import (
	"reflect"
)

type CommandFlag struct {
	Command    *Command
	Name       string
	Type       reflect.Kind
	Required   bool
	Default    any
	Shorthand  byte
	Usage      string
	AssignFunc func(any)
}

func (f *CommandFlag) panicFunc(typ string) {
	panicf("ERROR: Default value for command '%s' flag '%s' is not a %s",
		f.Command.Use,
		f.Name,
		typ,
	)
}

func (f *CommandFlag) DefaultString() string {
	var d string
	var ok bool
	if f.Default == nil {
		goto end
	}
	d, ok = f.Default.(string)
	if !ok {
		f.panicFunc("string")
	}
end:
	return d
}

func (f *CommandFlag) DefaultInt() int {
	var d int
	var ok bool
	if f.Default == nil {
		goto end
	}
	d, ok = f.Default.(int)
	if !ok {
		f.panicFunc("int")
	}
end:
	return d
}

func (f *CommandFlag) DefaultBool() bool {
	var d bool
	var ok bool
	if f.Default == nil {
		goto end
	}
	d, ok = f.Default.(bool)
	if !ok {
		f.panicFunc("bool")
	}
end:
	return d
}
