package main

import (
	"io"

	"github.com/mikeschinkel/prefsctl/cliutil"
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
			if *GlobalFlags.Quiet {
				SetQuiet(cmd)
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

var GlobalFlags struct {
	Quiet *bool
}

func SetQuiet(cmd *cobra.Command) {
	cmd.SetOut(io.Discard)
}

func init() {
	cliutil.SetRootCmd(rootCmd)
	GlobalFlags.Quiet = rootCmd.PersistentFlags().BoolP(
		"quiet",
		"q",
		false,
		"Disable informational messages to stdOut",
	)
	cliutil.AddInitializer(func(cli *cliutil.CLI) {
		// Add if needed
	})
}
