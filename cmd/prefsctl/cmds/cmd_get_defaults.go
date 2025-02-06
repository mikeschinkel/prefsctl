package cmds

import (
	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/spf13/cobra"
)

func init() {
	cobrautil.AddInitializer(func(cli *CLI) {
		getCmd.AddCmd(getDefaultsCmd)
	})
}

var getDefaultsProps = &GetDefaultsProps{}

type GetDefaultsProps struct {
	BaseProps
	Domains *[]string
	//UseCurrent *bool
	//filename macprefs.FilenamePtr
	//dummy *string
}

var getDefaultsCmd = NewCmdFromOpts(CmdOpts{
	Parent: getCmd,
	Command: &cobra.Command{
		Use:   "defaults",
		Short: "Get preference defaults",
	},
	Props:   getDefaultsProps,
	RunFunc: runGetDefaultsFunc,
	Flags: []*CmdFlag{
		DomainsFlag(func(value any) {
			// This assigns the pointer, the value has not yet been retrieved from os.Args
			getDefaultsProps.Domains = value.(*[]string)
		}),
		//{
		//	Name:     UseCurrentFlagName,
		//	Type:     reflect.Bool,
		//	Descr:    "Use current values from macOS for preference values not yet marked valid",
		//	Default:  false,
		//	Required: false,
		//	AssignFunc: func(value any) {
		//		// This assigns the pointer, the value has not yet been retrieved from os.Args
		//		getDefaultsProps.UseCurrent = value.(*bool)
		//	},
		//{
		//	Name:      macprefs.FilenameFlag,
		//	Type:      reflect.String,
		//	Required:  true,
		//	Shorthand: 'f',
		//	AssignFunc: func(value any) {
		// 		// This assigns the pointer, the value has not yet been retrieved from os.Args
		//		getDefaultsProps.filename = macprefs.FilenamePtr(value.(*string))
		//	},
		//},
	},
})

func runGetDefaultsFunc(ctx Context, cfg cobrautil.Config, cmd Cmd) cobrautil.CmdResult {
	//p := cmd.Props.(*GetDefaultsProps)
	return macprefs.GetDefaults(ctx, cfg, cmd, macprefs.QueryArgs{
		//UseCurrent: *getDefaultsProps.UseCurrent,
		Domains: macprefs.ToDomainNames(*getDefaultsProps.Domains),
	}).CobraUtilResult(cmd)
}
