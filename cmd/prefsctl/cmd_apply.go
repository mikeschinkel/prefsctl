package main

import (
	"reflect"

	"github.com/mikeschinkel/prefsctl/cliutil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/spf13/cobra"
)

var applyProps = &ApplyProps{}

type ApplyProps struct {
	cliutil.BaseProps
	filename macprefs.FilenamePtr
}

// applyCmd represents the `ip deallocate` command
var applyCmd = cliutil.NewCommandFromArgs(cliutil.CommandOpts{
	Parent: rootCmd,
	Command: &cobra.Command{
		Use:   "apply -f <filename>",
		Short: "Apply a MacPrefs manifest",
		Long:  "Apply a MacPrefs manifest to set macOS preferences",
	},
	Props: applyProps,
	Flags: []*cliutil.CommandFlag{
		{
			Name:      macprefs.FilenameFlag,
			Type:      reflect.String,
			Required:  true,
			Shorthand: 'f',
			AssignFunc: func(value any) {
				applyProps.filename = macprefs.FilenamePtr(value.(*string))
			},
		},
	},
	RunFunc: func(ctx Context, props cliutil.Props) (cliutil.Result, error) {
		return macprefs.Apply(ctx, macprefs.ApplyArgs{
			Filename: macprefs.Filename(*props.(*ApplyProps).filename),
		})
	},
	SuccessMsg: "Successfully applied manifest:\n\n%s",
})
