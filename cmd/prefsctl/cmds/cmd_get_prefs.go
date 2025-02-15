package cmds

import (
	"reflect"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/spf13/cobra"
)

func init() {
	cobrautil.AddInitializer(func(cli *CLI) {
		getCmd.AddCmd(getPrefsCmd)
	})
}

var getPrefsProps = &GetPrefsProps{}

type GetPrefsProps struct {
	BaseProps
	IncludeUnchanged *bool
	Domains          *[]string
	//filename macprefs.FilenamePtr
	//dummy *string
}

var getPrefsCmd = NewCmdFromOpts(CmdOpts{
	Parent: getCmd,
	Command: &cobra.Command{
		Use:   "prefs",
		Short: "Get preference prefs",
	},
	Props:   getPrefsProps,
	RunFunc: runGetPrefsFunc,
	Flags: []*CmdFlag{
		{
			Name:     IncludeUnchangedFlagName,
			Type:     reflect.Bool,
			Descr:    "Include preferences that have been unchanged from their default values",
			Default:  false,
			Required: false,
			AssignFunc: func(value any) {
				// This assigns the pointer, the value has not yet been retrieved from os.Args
				getPrefsProps.IncludeUnchanged = value.(*bool)
			},
		},
		DomainsFlag(func(value any) {
			// This assigns the pointer, the value has not yet been retrieved from os.Args
			getPrefsProps.Domains = value.(*[]string)
		}),
		//{
		//	Name:      macprefs.FilenameFlag,
		//	Type:      reflect.String,
		//	Required:  true,
		//	Shorthand: 'f',
		//	AssignFunc: func(value any) {
		// 		// This assigns the pointer, the value has not yet been retrieved from os.Args
		//		getPrefsProps.filename = macprefs.FilenamePtr(value.(*string))
		//	},
		//},
	},
})

func runGetPrefsFunc(ctx Context, cfg cobrautil.Config, cmd Cmd) cobrautil.CmdResult {
	return macprefs.GetPrefs(ctx, cfg, cmd, macprefs.QueryArgs{
		IncludeUnchanged: *getPrefsProps.IncludeUnchanged,
		Domains:          macprefs.ToDomainNames(*getPrefsProps.Domains),
	}).CobraUtilResult(cmd)
}
