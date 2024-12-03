package cobrautil

import (
	"io"

	"github.com/spf13/cobra"
)

func SetQuiet(cmd *cobra.Command) {
	cmd.SetOut(io.Discard)
}
func SetCalledCmd(cmd *cobra.Command) {
	calledCmdMutex.Lock()
	calledCmd = cmd
	calledCmdMutex.Unlock()
}
func SetRootCmd(cmd *Command) {
	rootCmdMutex.Lock()
	rootCmd = cmd
	rootCmdMutex.Unlock()
}
