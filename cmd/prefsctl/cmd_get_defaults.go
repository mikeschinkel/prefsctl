package main

import (
	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getDefaultsCmd)
}

var getDefaultsProps = &GetDefaultsProps{}

type GetDefaultsProps struct {
	cobrautil.BaseProps
	//filename macprefs.FilenamePtr
	//dummy *string
}

var getDefaultsCmd = cobrautil.NewCommandFromArgs(cobrautil.CommandOpts{
	Parent: getCmd,
	Command: &cobra.Command{
		Use:   "defaults",
		Short: "Get preference defaults",
	},
	Props: getDefaultsProps,
	Flags: []*cobrautil.CommandFlag{
		//{
		//	Name:      macprefs.FilenameFlag,
		//	Type:      reflect.String,
		//	Required:  true,
		//	Shorthand: 'f',
		//	AssignFunc: func(value any) {
		//		getDefaultsProps.filename = macprefs.FilenamePtr(value.(*string))
		//	},
		//},
		//{
		//	Name:      "Dummy",
		//	Type:      reflect.String,
		//	Required:  true,
		//	Shorthand: 'd',
		//	AssignFunc: func(value any) {
		//		getDefaultsProps.dummy = value.(*string)
		//	},
		//},
	},
	RunFunc: func(ctx Context, cmd *cobrautil.Command, props cobrautil.Props) (cobrautil.Result, error) {
		//p := props.(*GetDefaultsProps)
		return macprefs.GetDefaults(ctx, cmd, macprefs.GetDefaultsArgs{
			//Dummy: *p.dummy,
		})
	},
})
