package stdlibex

import (
	"net"
	"regexp"
)

// Regular expression for validating a well-formed host name.
var hostnameRegex = regexp.MustCompile(`^(?i)[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?(?:\.[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?)*\.[a-z]{2,}$`)

func IsValidHostName(domain Hostname) bool {
	return hostnameRegex.MatchString(string(domain))
}

func IsValidIPv4Address(ip IPv4Address) bool {
	obj := net.ParseIP(string(ip))
	return obj != nil && obj.To4() != nil
}
