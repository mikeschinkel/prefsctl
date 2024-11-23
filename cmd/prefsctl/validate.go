package main

import (
	"github.com/mikeschinkel/prefsctl/cliutil"
	"github.com/spf13/cobra"
)

func validateAndAssignArgs(numArgs int, assignFunc func([]string) error) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) (err error) {
		err = cobra.ExactArgs(numArgs)(cmd, args)
		if err != nil {
			cliutil.SetCalledCmd(cmd) // Since rootCmd.PersistentPreRun will never get called
			goto end
		}
		err = assignFunc(args)
	end:
		return err
	}
}
