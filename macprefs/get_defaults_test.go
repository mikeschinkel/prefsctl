package macprefs_test

import (
	"context"
	"testing"

	"github.com/mikeschinkel/go-diffator"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macosutil/macosutilmock"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/macprefs/macprefstest"
	_ "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	GoFormat = macprefs.GoFormat
)

var (
	GlobalFlags   = &macprefs.GlobalFlags
	GetDefaults   = macprefs.GetDefaults
	MockMacOSUtil = macosutilmock.MockMacOSUtil
)

type (
	QueryArgs        = macprefs.QueryArgs
	OutputFormat     = macprefs.OutputFormat
	PreferenceDomain = macosutil.PreferenceDomain
	Preference       = macosutil.Preference
	Data             = macosutilmock.Data
)

func TestGetDefaults(t *testing.T) {
	tests := []struct {
		name            string
		args            QueryArgs
		output          OutputFormat
		errWanted       error
		mockData        Data
		domains         []PreferenceDomain
		domainsErr      error
		domainPrefs     map[PreferenceDomain][]*Preference
		domainsPrefErrs map[PreferenceDomain]error
		expectedOutput  string
	}{
		{
			name: "",
			args: QueryArgs{
				OmitEmpty: false,
			},
			output:         GoFormat,
			errWanted:      nil,
			expectedOutput: macprefstest.ExpectedOutputForTest(),
			mockData: Data{
				Domains:     macprefstest.DomainsForTest(),
				DomainPrefs: macprefstest.DomainPrefsForTest(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			GlobalFlags.Output = stdlibex.Ptr(string(tt.output))
			GlobalFlags.Quiet = stdlibex.Ptr(true)
			tt.args.Printer = &errutil.BufferPrinter{}
			MockMacOSUtil(tt.mockData)
			err := GetDefaults(ctx, tt.args)
			if errutil.ErrorCheckFails(t, "GetDefaults", tt.errWanted, err) {
				return
			}
			expected := tt.expectedOutput
			received := tt.args.PrinterOutput()
			if expected != received {
				diff := diffator.CompareStrings(expected, received, nil)
				t.Errorf("Body <(expected/received)>:\n\t%s", diff)
				return
			}
		})
	}
}
