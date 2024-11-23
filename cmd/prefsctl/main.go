package main

import (
	"os"

	"github.com/mikeschinkel/prefsctl/cliutil"
)

func main() {
	var outcome *cliutil.Outcome
	var err error
	var cfg *cliutil.Config
	var cli *cliutil.CLI

	ctx := cliutil.DefaultContext()
	cfg, err = GetConfig(ctx)
	if err != nil {
		goto end
	}
	cli = cliutil.NewCLI(cfg, os.Args[1:])
	err = cli.Initialize(ctx)
	if err != nil {
		goto end
	}
	outcome = cli.Execute(ctx, cli.Args)
end:
	if err != nil {
		outcome = cliutil.NewErrorOutcome(err)
	}
	if outcome.WasError() {
		//goland:noinspection GoDfaNilDereference
		cli.ShowUsage(outcome)
		os.Exit(1)
	}
}

//func main2() {
//	if err := rootCmd.Execute(); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//}
