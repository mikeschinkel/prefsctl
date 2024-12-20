package macosutil_test

import (
	"strings"
	"testing"

	"github.com/mikeschinkel/prefsctl/macosutil"
	"github.com/mikeschinkel/prefsctl/sliceconv"
)

var (
	GetVersionNumber      = macosutil.GetVersionNumber
	ValidateVersionNumber = macosutil.ValidateVersionNumber
	ValidateVersionName   = macosutil.ValidateVersionName
	ValidVersionNumbers   = macosutil.ValidVersionNumbers
	GetVersionCode        = macosutil.GetVersionCode
	VersionCode           = macosutil.VersionCode
	VersionName           = macosutil.VersionName
	GetVersionName        = macosutil.GetVersionName
)

func Test_macOSUtils_GetVersionNumber(t *testing.T) {
	gotV, err := GetVersionNumber()
	if err != nil {
		t.Error(err)
		return
	}
	if !ValidateVersionNumber(gotV) {
		t.Errorf("GetVersionNumber(); got invalid version number '%s'; expected one of these (ignoring patch, and minor if major is >10):\n%s",
			gotV,
			strings.Join(sliceconv.ToStrings(ValidVersionNumbers()), ", "),
		)
	}
}

func Test_macOSUtils_VersionCode(t *testing.T) {
	gotCode, err := VersionCode()
	if err != nil {
		t.Errorf("VersionCode() expected no error; got %v", err)
		return
	}
	gotName := GetVersionName(gotCode)
	if gotName == "" {
		t.Errorf("GetVersionName() expected a  macOS version name for code '%s'; got zero string", gotCode)
		return
	}
	if !ValidateVersionName(gotName) {
		t.Errorf("ValidateVersionName() expected a valid macOS version name; got '%s' instead", gotName)
		return
	}
}

func Test_macOSUtils_VersionName(t *testing.T) {
	gotName, err := VersionName()
	if err != nil {
		t.Errorf("VersionName() expected no error; got %v", err)
		return
	}
	if !ValidateVersionName(gotName) {
		t.Errorf("ValidateVersionName() expected a valid macOS version name; got '%s' instead", gotName)
		return
	}
}

func Test_macOSUtils_ValidVersionNumbers(t *testing.T) {
	nums := ValidVersionNumbers()
	if len(nums) == 0 {
		t.Error("ValidVersionNumbers() expected many version numbers, got none")
		return
	}
	if GetVersionCode(nums[0]) != macosutil.Cheetah {
		t.Errorf("ValidVersionNumbers() expected first version number '%s' to be for macOS %s", nums[0], macosutil.Cheetah)
		return
	}
	num := nums[len(nums)-1]
	if GetVersionCode(num) != macosutil.Sequoia {
		t.Errorf("ValidVersionNumbers() expected last version number '%s' to be for macOS %s", num, macosutil.Sequoia)
		return
	}
}
