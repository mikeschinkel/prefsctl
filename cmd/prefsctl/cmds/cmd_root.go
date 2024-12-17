package cmds

import (
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = NewCmdFromOpts(CmdOpts{
	Command: &cobra.Command{
		Use:   "prefsctl",
		Short: "CLI for managing macOS preferences",
		Long:  "CLI for managing macOS preferences, especially for use with Ansible",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cobrautil.SetCalledCmd(cmd)
			if *macprefs.GlobalFlags.Quiet {
				cobrautil.SetQuiet(cmd)
			}
		},
		// Silence usage as we present usage after calling Cobra
		SilenceUsage: true,
		// Silence errors as we handle errors after calling Cobra
		SilenceErrors: true,
	},
})

func init() {
	cobrautil.AddInitializer(func(cli *CLI) {

		macprefs.GlobalFlags.Quiet = PersistentBoolFlag(RootCmd, CmdFlagArgs{
			Name:      cobrautil.QuietFlagName,
			Shorthand: cobrautil.QuietFlagShorthand,
			Default:   false,
			Usage:     "Disable informational messages to stdOut",
		})

		macprefs.GlobalFlags.Output = PersistentStringFlag(RootCmd, CmdFlagArgs{
			Name:      cobrautil.OutputFlagName,
			Shorthand: cobrautil.OutputFlagShorthand,
			Default:   string(macprefs.TXTFormat),
			Usage: fmt.Sprintf("Specify the format for output; one of: %s",
				strings.Join(sliceconv.ToStrings(macprefs.AllFormats), ","),
			),
		})
	})
}
