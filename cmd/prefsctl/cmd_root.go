package main

import (
	"fmt"
	"strings"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/sliceconv"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = cobrautil.NewCommandFromArgs(cobrautil.CommandOpts{
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

func GetConfig(ctx Context) (cfg *cobrautil.Config, err error) {
	var options cobrautil.OptionsMap
	fp, err := cobrautil.ConfigFilepath()
	if err != nil {
		goto end
	}
	options = cobrautil.OptionsMap{}
	cfg = cobrautil.NewConfig(&cobrautil.ConfigArgs{
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
	cobrautil.SetRootCmd(rootCmd)
	cobrautil.AddInitializer(func(cli *cobrautil.CLI) {
		macprefs.GlobalFlags.Quiet = rootCmd.PersistentFlags().BoolP(
			cobrautil.QuietFlag,
			cobrautil.QuietFlagShorthand,
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
