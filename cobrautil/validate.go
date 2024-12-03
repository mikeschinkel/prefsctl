package cobrautil

import (
	"github.com/spf13/cobra"
)

func ValidateAndAssignArgs(numArgs int, assignFunc func([]string) error) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) (err error) {
		err = cobra.ExactArgs(numArgs)(cmd, args)
		if err != nil {
			SetCalledCmd(cmd) // Since rootCmd.PersistentPreRun will never get called
			goto end
		}
		err = assignFunc(args)
	end:
		return err
	}
}
