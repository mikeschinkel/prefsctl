package macprefs_test

import (
	"context"
	"testing"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	_ "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	GoFormat = macprefs.GoFormat
)

var (
	GlobalFlags = &macprefs.GlobalFlags
	GetDefaults = macprefs.GetDefaults
)

type (
	GetDefaultsArgs = macprefs.GetDefaultsArgs
	Printer         = macprefs.Printer
	OutputFormat    = macprefs.OutputFormat
)

func TestGetDefaults(t *testing.T) {
	tests := []struct {
		name      string
		ptr       Printer
		args      GetDefaultsArgs
		output    OutputFormat
		errWanted error
	}{
		{
			name:      "",
			args:      GetDefaultsArgs{},
			output:    GoFormat,
			errWanted: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			GlobalFlags.Output = stdlibex.Ptr(string(tt.output))
			GlobalFlags.Quiet = stdlibex.Ptr(true)
			tt.args.Printer = &errutil.BufferPrinter{}
			err := GetDefaults(ctx, tt.args)
			if errutil.ErrorCheckFails(t, "GetDefaults", tt.errWanted, err) {
				return
			}
		})
	}
}
