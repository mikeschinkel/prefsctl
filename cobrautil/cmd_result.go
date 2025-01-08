package cobrautil

import (
	"github.com/spf13/cobra"
)

var NoResult = CmdResult{}

type CmdResult struct {
	cmd     Cmd
	Success string
	err     error
}

func NewCmdResult(cmd Cmd, success string, err error) CmdResult {
	return CmdResult{
		cmd:     cmd,
		Success: success,
		err:     err,
	}
}
func NewSuccessResult(cmd Cmd, success string) CmdResult {
	return NewCmdResult(cmd, success, nil)
}
func NewErrorResult(cmd Cmd, err error) CmdResult {
	return NewCmdResult(cmd, "", err)
}

func (n CmdResult) IsErr() bool {
	return n.err != nil
}

func (n CmdResult) SetErr(err error) {
	n.err = err
}

func (n CmdResult) Err() error {
	return n.err
}

func (n CmdResult) Cmd() Cmd {
	return n.cmd
}

func (n CmdResult) Command() *cobra.Command {
	return n.cmd.Command()
}

func (n CmdResult) String() string {
	s := n.err.Error()
	if s == "" {
		s = "Success"
	}
	return s
}
func (n CmdResult) Error() string {
	return n.err.Error()
}
