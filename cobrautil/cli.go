package cobrautil

import (
	"context"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/mikeschinkel/prefsctl/stdlibex"
	"github.com/spf13/cobra"
)

var rootCmd Cmd
var rootCmdMutex sync.Mutex

var calledCmd *cobra.Command
var calledCmdMutex sync.Mutex

func RootCmd() Cmd {
	return rootCmd
}

//func SetArgsPassed(args []string) {
//	argsPassed = args
//}

type CLI struct {
	Config Config
	Args   []string
}

func NewCLI() *CLI {
	return &CLI{}
}

func (cli *CLI) Execute(ctx Context, args []string) (result CmdResult, err error) {
	rootCmd.SetArgs(args)
	err = rootCmd.Command().Execute()
	result = NewCmdResult(
		NewCmd(cli, calledCmd),
		err,
	)
	return result, err
}

func DefaultContext() Context {
	return context.Background()
}

func (cli *CLI) Initialize(ctx Context, cfg Config, args []string) error {
	var exists bool

	cli.Config = cfg
	cli.Args = args

	file, err := ConfigFilepath()
	if err != nil {
		goto end
	}
	err = stdlibex.EnsureDir(filepath.Dir(file))
	if err != nil {
		goto end
	}
	exists, err = stdlibex.FileExists(file)
	if err != nil {
		goto end
	}
	if !exists {
		err = SaveConfig(ctx, cli.Config, file)
	}
	if err != nil {
		goto end
	}
	RunInitializers(cli)
end:
	return err
}

var regex = regexp.MustCompile("\n([a-z0-9_]+=)")

func (cli *CLI) ShowUsage(r CmdResult) {
	rootCmd.PrintErr("\nERROR: ")
	rootCmd.PrintErrln(regex.ReplaceAllString(r.Err().Error(), "; $1"))
	rootCmd.PrintErrln("")
	rootCmd.PrintErrln(r.Command().UsageString())
	rootCmd.PrintErrln("")
}
