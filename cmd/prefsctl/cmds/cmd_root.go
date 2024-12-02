package cmds

import (
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/cliutil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = cliutil.NewCommandFromArgs(cliutil.CommandOpts{
	Command: &cobra.Command{
		Use:   "prefsctl",
		Short: "CLI for managing macOS preferences",
		Long:  "CLI for managing macOS preferences, especially for use with Ansible",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cliutil.SetCalledCmd(cmd)
			if *macprefs.GlobalFlags.Quiet {
				cliutil.SetQuiet(cmd)
			}
		},
		// Silence usage as we present usage after calling Cobra
		SilenceUsage: true,
		// Silence errors as we handle errors after calling Cobra
		SilenceErrors: true,
	},
})

func GetConfig(ctx Context) (cfg *cliutil.Config, err error) {
	var options cliutil.OptionsMap
	fp, err := cliutil.ConfigFilepath()
	if err != nil {
		goto end
	}
	options = cliutil.OptionsMap{}
	cfg = cliutil.NewConfig(&cliutil.ConfigArgs{
		Filepath: fp,
		Options:  options,
	})
	err = cfg.Initialize(ctx)
	if err != nil {
		goto end
	}
end:
	return cfg, err
}

func init() {
	cliutil.SetRootCmd(rootCmd)
	cliutil.AddInitializer(func(cli *cliutil.CLI) {
		macprefs.GlobalFlags.Quiet = rootCmd.PersistentFlags().BoolP(
			cliutil.QuietFlag,
			cliutil.QuietFlagShorthand,
			false,
			"Disable informational messages to stdOut",
		)
		macprefs.GlobalFlags.Output = rootCmd.PersistentFlags().StringP(
			macprefs.OutputFlag,
			macprefs.OutputFlagShorthand,
			string(macprefs.TXTFormat),
			fmt.Sprintf("Specify the format for output; one of: %s",
				strings.Join(sliceconv.ToStrings(macprefs.AllFormats), ","),
			),
		)
	})
}
