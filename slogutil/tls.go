package slogutil

import (
	"crypto/tls"
	"fmt"

	"github.com/mikeschinkel/prefsctl/logargs"
)

// TLSLogArgs returns a human-readable representation of the TLS state, omitting zero-value fields.
func TLSLogArgs(s *tls.ConnectionState) (args []any) {
	if s == nil {
		goto end
	}

	if s.Version != 0 {
		args = append(args, logargs.TLSVersion, _tls(s.Version).String())
	}

	if !s.HandshakeComplete {
		args = append(args, logargs.TLSHandshake, "pending")
	} else {
		args = append(args, logargs.TLSHandshake, "complete")
	}

	if len(s.PeerCertificates) > 0 {
		args = append(args, logargs.PeerCertCount, len(s.PeerCertificates))
	}

	if len(s.VerifiedChains) > 0 {
		args = append(args, logargs.VerifiedChainCount, len(s.VerifiedChains))
	}

	if s.CipherSuite != 0 {
		args = append(args, logargs.CipherSuite, _cipher(s.CipherSuite).String())
	}

	if s.NegotiatedProtocol != "" {
		args = append(args, logargs.NegotiatedProtocol, s.NegotiatedProtocol)
	}

	if s.ServerName != "" {
		args = append(args, logargs.ServerName, s.ServerName)
	}

end:
	return args
}

type _tls uint16

// String converts a TLS version number to a human-readable string.
func (t _tls) String() string {
	switch uint16(t) {
	case tls.VersionTLS10:
		return "1.0"
	case tls.VersionTLS11:
		return "1.1"
	case tls.VersionTLS12:
		return "1.2"
	case tls.VersionTLS13:
		return "1.3"
	default:
		return "???"
	}
}

type _cipher uint16

// String converts a cipher suite number to a human-readable string.
func (c _cipher) String() string {
	switch uint16(c) {
	case tls.TLS_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_RSA_WITH_AES_128_CBC_SHA"
	case tls.TLS_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_RSA_WITH_AES_256_CBC_SHA"
	case tls.TLS_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_RSA_WITH_AES_128_GCM_SHA256"
	case tls.TLS_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_RSA_WITH_AES_256_GCM_SHA384"
	case tls.TLS_AES_128_GCM_SHA256:
		return "TLS_AES_128_GCM_SHA256"
	case tls.TLS_AES_256_GCM_SHA384:
		return "TLS_AES_256_GCM_SHA384"
	case tls.TLS_CHACHA20_POLY1305_SHA256:
		return "TLS_CHACHA20_POLY1305_SHA256"
	default:
		return fmt.Sprintf("Unknown (%d)", c)
	}
}
