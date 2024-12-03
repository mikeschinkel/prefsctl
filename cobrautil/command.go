package cobrautil

import (
	"context"
	"reflect"

	"github.com/mikeschinkel/prefsctl/stdlibex"
	"github.com/spf13/cobra"
)

type Command struct {
	*cobra.Command
	CLI   *CLI
	Props Props
}

func NewCommand(cli *CLI, command *cobra.Command) *Command {
	return &Command{
		Command: command,
		CLI:     cli,
	}
}

func (c *Command) AddCommand(cmd *Command) {
	if c.Command != cmd.Command {
		c.Command.AddCommand(cmd.Command)
	}
}

type RunFunc = func(Context, *Command, Props) (Result, error)
type CommandOpts struct {
	Parent     *Command
	Command    *cobra.Command
	RunFunc    RunFunc
	Flags      []*CommandFlag
	Props      Props
	SuccessMsg string
}

func NewCommandFromArgs(opts CommandOpts) (cmd *Command) {
	cmd = &Command{
		Command: opts.Command,
	}
	AddInitializer(func(cli *CLI) {
		cmd.CLI = cli
		rootCmd.AddCommand(cmd)
		cmd.Command.RunE = func(cc *cobra.Command, args []string) (err error) {
			var result Result

			ctx := context.Background()
			cfg := cli.Config

			props, err := ProcessProps(cfg, opts.Props)
			if err != nil {
				goto end
			}
			result, err = opts.RunFunc(ctx, cmd, props)
			if err != nil {
				goto end
			}
		end:
			SetResult(result)
			if err == nil && len(opts.SuccessMsg) != 0 {
				//goland:noinspection GoDfaNilDereference
				cc.Printf(opts.SuccessMsg+"\n",
					stdlibex.IndentLines(2, result.String()),
				)
			}
			return err
		}
		cmd.Props = opts.Props
		for _, flag := range opts.Flags {
			flag.Command = cmd
			switch flag.Type {
			case reflect.String:
				flag.AssignFunc(cmd.Flags().StringP(
					flag.Name,
					string(flag.Shorthand),
					flag.DefaultString(),
					flag.Usage,
				))
			case reflect.Int:
				flag.AssignFunc(cmd.Flags().IntP(
					flag.Name,
					string(flag.Shorthand),
					flag.DefaultInt(),
					flag.Usage,
				))
			case reflect.Bool:
				flag.AssignFunc(cmd.Flags().BoolP(
					flag.Name,
					string(flag.Shorthand),
					flag.DefaultBool(),
					flag.Usage,
				))
			default:
				panicf("Flag Type '%s' not implemented", flag.Type)
			}
			if flag.Required {
				// Ignoring the error as the only error thrown is "flag not found" and since we
				// are using a const that should never happy. No really, it should never happen.
				_ = cmd.MarkFlagRequired(flag.Name)
			}
		}
	})
	return cmd
}
