package prefsyaml

import (
	"fmt"
	"log/slog"

	"github.com/hashicorp/go-version"
	"github.com/mikeschinkel/prefsctl/logargs"
)

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

// SupportsOSVersion tests to see if the key is supported by the current macOS
// version, assuming that MacOSMin is not empty, otherwise assume supported.
func SupportsOSVersion(ver, minVer, maxVer OSVersion) (supports bool) {
	var err error
	var v, vMin, vMax *version.Version

	// If we can't tell then assume supported
	supports = true

	if minVer == "" && maxVer == "" {
		goto end
	}
	v, err = version.NewVersion(string(ver))
	if err != nil {
		slog.Warn(err.Error())
		goto end
	}
	vMin, err = version.NewVersion(string(minVer))
	if err != nil {
		slog.Info("min version not a valid version", logargs.Version, minVer)
		goto end
	}
	if minVer != "" && maxVer == "" {
		// 0 for v==vMin, 1 for v>vMin
		supports = v.Compare(vMin) != -1
		goto end
	}
	vMax, err = version.NewVersion(string(maxVer))
	if err != nil {
		slog.Info("max version not a valid version", logargs.Version, maxVer)
		goto end
	}
	if minVer == "" {
		// -1 for v<vMax
		supports = v.Compare(vMax) == -1
		goto end
	}
	// 0 for v==vMin, 1 for v>vMin
	supports = v.Compare(vMin) != -1 &&
		// -1 for v<vMax
		v.Compare(vMax) == -1
end:
	return supports
}
