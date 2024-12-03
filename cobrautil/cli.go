package cobrautil

import (
	"context"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/mikeschinkel/prefsctl/stdlibex"
	"github.com/spf13/cobra"
)

var rootCmd *Command
var rootCmdMutex sync.Mutex

var calledCmd *cobra.Command
var calledCmdMutex sync.Mutex

func RootCmd() *Command {
	return rootCmd
}

//func SetArgsPassed(args []string) {
//	argsPassed = args
//}

type CLI struct {
	Config *Config
	Args   []string
}

func NewCLI(cfg *Config, args []string) *CLI {
	return &CLI{
		Config: cfg,
		Args:   args,
	}
}

func (cli *CLI) Execute(_ Context, args []string) *Outcome {
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	return NewCmdOutcome(err, calledCmd)
}

func DefaultContext() Context {
	return context.Background()
}

func (cli *CLI) Initialize(ctx Context) error {
	var exists bool
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

func (cli *CLI) ShowUsage(r *Outcome) {
	rootCmd.PrintErr("\nERROR: ")
	rootCmd.PrintErrln(regex.ReplaceAllString(r.Err.Error(), "; $1"))
	rootCmd.PrintErrln("")
	rootCmd.PrintErrln(r.Command.UsageString())
	rootCmd.PrintErrln("")
}
