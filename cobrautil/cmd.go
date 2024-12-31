package cobrautil

import (
	"context"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

type Printer interface {
	Print(i ...any)
	Println(i ...any)
	Printf(format string, i ...any)
	PrintErr(i ...any)
	PrintErrln(i ...any)
	PrintErrf(format string, i ...any)
}

func PersistentBoolFlag(cmd Cmd, args CmdFlagArgs) *bool {
	return cmd.Command().PersistentFlags().BoolP(
		args.Name,
		string(args.Shorthand),
		args.Default.(bool),
		args.Usage,
	)
}

func PersistentStringFlag(cmd Cmd, args CmdFlagArgs) *string {
	return cmd.Command().PersistentFlags().StringP(
		args.Name,
		string(args.Shorthand),
		args.Default.(string),
		args.Usage,
	)
}

type Cmd interface {
	HasSubCmds() bool
	SubCmd(string) Cmd
	Props() Props
	Flags() *flag.FlagSet
	SetProps(Props)
	SetArgs([]string)
	Execute(Context) error
	Command() *cobra.Command
	AddCmd(Cmd)
	Printer
}

var _ Cmd = (*cmd)(nil)

type cobraCommand = cobra.Command

type cmd struct {
	*cobraCommand
	Parent  Cmd
	CLI     *CLI
	props   Props
	SubCmds map[string]Cmd
	RunFunc RunFunc
}

func (c *cmd) Command() *cobra.Command {
	cc := c.cobraCommand
	if cc == nil {
		cc = rootCmd.Command()
	}
	return cc
}

func (c *cmd) Props() Props {
	return c.props
}
func (c *cmd) Execute(ctx Context) error {
	return c.RunFunc(ctx, c)
}
func (c *cmd) SetProps(props Props) {
	c.props = props
}

func (c *cmd) SetArgs(args []string) {
	c.cobraCommand.SetArgs(args)
}

func NewCmd(cli *CLI, command *cobra.Command) Cmd {
	return &cmd{
		cobraCommand: command,
		CLI:          cli,
		SubCmds:      make(map[string]Cmd),
	}
}

func (c *cmd) SubCmd(usage string) Cmd {
	subCmd, _ := c.SubCmds[usage]
	return subCmd
}

func (c *cmd) HasSubCmds() bool {
	return len(c.SubCmds) != 0
}

func CalledCmd(c Cmd, args []string) (subCmd Cmd) {
	subCmd = c
	if !c.HasSubCmds() {
		goto end
	}
	for _, arg := range args {
		sc := subCmd.SubCmd(arg)
		if sc == nil {
			goto end
		}
		subCmd = sc
	}
end:
	return subCmd
}

func (c *cmd) AddCmd(cmd Cmd) {
	cc := c.Command()
	command := cmd.Command()
	if cc != command {
		cc.AddCommand(command)
		if c.SubCmds == nil {
			c.SubCmds = make(map[string]Cmd)
		}
		c.SubCmds[command.Use] = cmd
	}
}

type RunFunc = func(Context, Cmd) error

type CmdOpts struct {
	Parent     Cmd
	Command    *cobra.Command
	Flags      []*CmdFlag
	Props      Props
	RunFunc    RunFunc
	SuccessMsg string
}

func NewCmdFromOpts(opts CmdOpts) Cmd {
	newCmd := &cmd{
		cobraCommand: opts.Command,
		SubCmds:      make(map[string]Cmd),
	}
	AddInitializer(func(cli *CLI) {
		newCmd.CLI = cli
		newCmd.Parent = opts.Parent
		newCmd.props = opts.Props
		newCmd.RunFunc = opts.RunFunc
		if newCmd.RunFunc == nil {
			newCmd.RunFunc = func(ctx Context, cmd Cmd) (err error) {
				calledCmd := CalledCmd(cmd, cli.Args)
				if cmd != calledCmd {
					err = calledCmd.Execute(ctx)
					goto end
				}
				println("TODO: Call Command Help Here")
			end:
				return err
			}
		}
		for _, f := range opts.Flags {
			f.Cmd = newCmd
			shorthand := string(f.Shorthand)
			if f.Shorthand == 0 {
				shorthand = ""
			}
			switch f.Type {
			case reflect.String:
				f.AssignFunc(newCmd.Flags().StringP(
					f.Name,
					shorthand,
					f.DefaultString(),
					f.Usage,
				))
			case reflect.Int:
				f.AssignFunc(newCmd.Flags().IntP(
					f.Name,
					shorthand,
					f.DefaultInt(),
					f.Usage,
				))
			case reflect.Bool:
				f.AssignFunc(newCmd.Flags().BoolP(
					f.Name,
					shorthand,
					f.DefaultBool(),
					f.Usage,
				))
			case reflect.Slice:
				switch f.Subtype {
				case reflect.String:
					f.AssignFunc(newCmd.Flags().StringSliceP(
						f.Name,
						shorthand,
						f.DefaultSliceString(),
						f.Usage,
					))
				default:
					panicf("Flag Type 'Slice' Subtype '%s' not implemented", f.Subtype)
				}
			default:
				panicf("Flag Type '%s' not implemented", f.Type)
			}
			if f.Required {
				// Ignoring the error as the only error thrown is "flag not found" and since we
				// are using a const that should never happy. No really, it should never happen.
				_ = newCmd.MarkFlagRequired(f.Name)
			}
		}
		newCmd.Command().RunE = func(cc *cobra.Command, args []string) (err error) {
			var props Props
			cfg := cli.Config
			subCmd := CalledCmd(newCmd, args)
			if subCmd == nil {
				panicf("Error finding command for: %s", strings.Join(args, " "))
			}
			props, err = ProcessProps(cfg, subCmd.Props())
			if err != nil {
				goto end
			}
			subCmd.SetProps(props)
			err = subCmd.Execute(context.Background())
			if err != nil {
				goto end
			}
		end:
			if err == nil && len(opts.SuccessMsg) != 0 {
				//goland:noinspection GoDfaNilDereference
				cc.Printf(opts.SuccessMsg)
				//cc.Printf(opts.SuccessMsg+"\n",
				//	stdlibex.IndentLines(2, result.String()),
				//)
			}
			return err
		}
	})
	return newCmd
}
