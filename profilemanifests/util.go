package profilemanifests

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}

func noop(x ...any) {

}

func parseMajorVersion(v string) (int, error) {
	parts := strings.Split(v, ".")
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid version format")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	return major, nil
}
