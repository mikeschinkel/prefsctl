package cobrautil

import (
	"io"
	"os"
)

type ExecuteArgs struct {
	Config        Config
	CLIArgs       []string
	NoUsageOnErr  bool
	OutWriter     io.Writer
	ErrWriter     io.Writer
	ConfigOptions OptionsMap
}

func Execute(rootCmd Cmd, args ExecuteArgs) (result CmdResult, err error) {
	var cfg Config

	SetRootCmd(rootCmd)
	cc := rootCmd.Command()
	if args.OutWriter == nil {
		args.OutWriter = os.Stdout
	}
	if args.ErrWriter == nil {
		args.OutWriter = os.Stderr
	}
	cc.SetOut(args.OutWriter)
	cc.SetErr(args.ErrWriter)
	cli := NewCLI()
	ctx := DefaultContext()
	if args.Config == nil {
		cfg = NewConfig(ConfigArgs{Options: args.ConfigOptions})
		err = cfg.Initialize(ctx)
		if err != nil {
			result = NewErrorResult(nil, err)
			goto end
		}
	}
	if args.CLIArgs == nil {
		args.CLIArgs = os.Args[1:]
	}
	err = cli.Initialize(ctx, cfg, args.CLIArgs)
	if err != nil {
		result = NewErrorResult(nil, err)
		goto end
	}
	result, err = cli.Execute(ctx, cli.Args)
end:
	if err != nil && !args.NoUsageOnErr {
		cli.ShowUsage(result)
	}
	return result, err
}
