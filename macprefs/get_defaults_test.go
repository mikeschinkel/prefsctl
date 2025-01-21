package macprefs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mikeschinkel/go-diffator"
	"github.com/mikeschinkel/prefsctl/appinfo"
	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/macosutil/macosutilmock"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/mikeschinkel/prefsctl/macprefs/macprefstest"
	_ "github.com/mikeschinkel/prefsctl/prefdefaults"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	YAMLFormat = macprefs.YAMLFormat
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

func Test_GetDefaults(t *testing.T) {
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
			output:         YAMLFormat,
			errWanted:      nil,
			expectedOutput: macprefstest.ExpectedOutputForTest(t),
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
			ptr := &errutil.BufferPrinter{}
			MockMacOSUtil(tt.mockData)
			cfg := cobrautil.NewConfig(cobrautil.ConfigArgs{
				AppName: fmt.Sprintf("%s-test", appinfo.Name),
			})
			err := cfg.Initialize(ctx)
			if err != nil {
				t.Fatalf(err.Error())
			}
			result := GetDefaults(ctx, cfg, ptr, tt.args)
			if errutil.ErrorCheckFails(t, "GetDefaults", tt.errWanted, result.Err) {
				return
			}
			expected := tt.expectedOutput
			received := ptr.String()
			//_ = os.WriteFile(macprefstest.GetDefaultsTestOutputFile, []byte(received), os.ModePerm)
			//t.Logf("\n\n%s\n\n", received)
			if expected != received {
				diff := diffator.CompareStrings(expected, received, nil)
				t.Errorf("Body <(expected/received)>:\n\t%s", diff)
				return
			}
		})
	}
}
