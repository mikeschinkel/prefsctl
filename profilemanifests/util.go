package profilemanifests

import (
	"io"

	"github.com/hashicorp/go-version"
	"github.com/mikeschinkel/prefsctl/logargs"
)

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}

// SupportsOSVersion tests to see if the key is supported by the current macOS
// version, assuming that MacOSMin is not empty, otherwise assume supported.
func supportsOSVersion(ver, minVer, maxVer string) (supports bool) {
	var err error
	var v, vMin, vMax *version.Version

	// If we can't tell then assume supported
	supports = true

	if minVer == "" && maxVer == "" {
		goto end
	}
	v, err = version.NewVersion(ver)
	if err != nil {
		slog.Warn(err.Error())
		goto end
	}
	vMin, err = version.NewVersion(minVer)
	if err != nil {
		slog.Info("min version not a valid version", logargs.Version, minVer)
		goto end
	}
	if minVer != "" && maxVer == "" {
		// 0 for v==vMin, 1 for v>vMin
		supports = v.Compare(vMin) != -1
		goto end
	}
	vMax, err = version.NewVersion(maxVer)
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
