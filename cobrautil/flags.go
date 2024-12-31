package cobrautil

import (
	"strings"

	flag "github.com/spf13/pflag"
)

type Flags struct {
	*flag.FlagSet
	Cmd Cmd
}
type FlagOpts struct {
	Default  any
	Required bool
}

func (f *Flags) StringFlag(name, shorthand, usage string) *string {
	return f.FlagSet.StringP(name, shorthand, "", usage)
}
func (f *Flags) StringFlagWithOpts(name, shorthand, usage string, opts FlagOpts) *string {
	def, _ := opts.Default.(string)
	fs := f.FlagSet.StringP(name, shorthand, def, usage)
	f.applyFlagOpts(name, def == "", opts)
	return fs
}
func (f *Flags) IntFlag(name, shorthand, usage string) *int {
	return f.FlagSet.IntP(name, shorthand, 0, usage)
}
func (f *Flags) IntFlagWithOpts(name, shorthand, usage string, opts FlagOpts) *int {
	def, _ := opts.Default.(int)
	fs := f.FlagSet.IntP(name, shorthand, def, usage)
	f.applyFlagOpts(name, def == 0, opts)
	return fs
}
func (f *Flags) BoolFlag(name, shorthand, usage string) *bool {
	return f.FlagSet.BoolP(name, shorthand, false, usage)
}
func (f *Flags) BoolFlagWithOpts(name, shorthand, usage string, opts FlagOpts) *bool {
	def, _ := opts.Default.(bool)
	fs := f.FlagSet.BoolP(name, shorthand, def, usage)
	f.applyFlagOpts(name, def == false, opts)
	return fs
}
func (f *Flags) StringSliceFlag(name, shorthand, usage string) *[]string {
	return f.FlagSet.StringSliceP(name, shorthand, []string{}, usage)
}
func (f *Flags) StringSliceFlagWithOpts(name, shorthand, usage string, opts FlagOpts) *[]string {
	def, _ := opts.Default.([]string)
	fs := f.FlagSet.StringSliceP(name, shorthand, def, usage)
	f.applyFlagOpts(name, strings.Join(def, "") == "", opts)
	return fs
}

func (f *Flags) applyFlagOpts(name string, chkRequired bool, opts FlagOpts) {
	var err error
	if chkRequired && opts.Required {
		err = f.Cmd.Command().MarkFlagRequired(name)
	}
	if err != nil {
		panic(err)
	}
}

func GetFlags(cmd Cmd) *Flags {
	var cobraFlags = cmd.Flags()
	return &Flags{
		FlagSet: cobraFlags,
		Cmd:     cmd,
	}
}
