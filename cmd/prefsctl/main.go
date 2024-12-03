package main

import (
	"os"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	_ "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"
)

func main() {
	var outcome *cobrautil.Outcome
	var err error
	var cfg *cobrautil.Config
	var cli *cobrautil.CLI

	ctx := cobrautil.DefaultContext()
	cfg, err = GetConfig(ctx)
	if err != nil {
		goto end
	}
	cli = cobrautil.NewCLI(cfg, os.Args[1:])
	err = cli.Initialize(ctx)
	if err != nil {
		goto end
	}
	outcome = cli.Execute(ctx, cli.Args)
end:
	if err != nil {
		outcome = cobrautil.NewErrorOutcome(err)
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
