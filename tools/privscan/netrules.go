package main

import (
	"bytes"
	"net/netip"
)

// allowedIPv4 lists the ranges that are no personal data: documentation (RFC 5737), private
// (RFC 1918), loopback, link-local, "this network", IETF protocol assignments, benchmarking,
// multicast and reserved including the broadcast address. 100.64.0.0/10 (RFC 6598, carrier NAT
// and tailnets) is deliberately not here: its addresses are a hit.
func allowedIPv4() []netip.Prefix {
	return prefixes("192.0.2.0/24", "198.51.100.0/24", "203.0.113.0/24", "10.0.0.0/8", "172.16.0.0/12",
		"192.168.0.0/16", "127.0.0.0/8", "169.254.0.0/16", "0.0.0.0/8", "192.0.0.0/24", "198.18.0.0/15",
		"224.0.0.0/4", "240.0.0.0/4")
}

// allowedIPv6 lists documentation (RFC 3849, RFC 9637), unspecified, loopback, unique local,
// multicast and NAT64 ranges. Link-local is decided by linkLocalIsPersonal.
func allowedIPv6() []netip.Prefix {
	return prefixes("2001:db8::/32", "3fff::/20", "::/128", "::1/128", "fc00::/7", "ff00::/8", "64:ff9b::/96")
}

func prefixes(list ...string) []netip.Prefix {
	out := make([]netip.Prefix, 0, len(list))
	for _, p := range list {
		out = append(out, netip.MustParsePrefix(p))
	}
	return out
}

func inAny(addr netip.Addr, list []netip.Prefix) bool {
	for _, p := range list {
		if p.Contains(addr) {
			return true
		}
	}
	return false
}

// ipv4Finder finds dotted quads that are not part of a longer dotted number.
func ipv4Finder() func([]byte) []span {
	allowed := allowedIPv4()
	return func(line []byte) []span { return findIPv4(line, allowed) }
}

func findIPv4(line []byte, allowed []netip.Prefix) []span {
	var out []span
	for _, r := range runs(line, func(c byte) bool { return c >= '0' && c <= '9' || c == '.' }) {
		s, ok := quadIn(line, r)
		if !ok {
			continue
		}
		addr, ok := parseQuad(line[s.start:s.end])
		if !ok || inAny(addr, allowed) {
			continue
		}
		out = append(out, s)
	}
	return out
}

// quadIn returns the dotted-quad candidate of the digit-and-dot run r: the run without its
// trailing dots. A run that starts with a dot is part of a longer dotted name (libx.so.1.2.3.4)
// and no candidate, unless it stands between two letter labels with one dot on each side, as in
// the nip.io-style name app.<a.b.c.d>.nip.io: an address embedded there is no version number.
// A quad followed by ".<digit>" is part of a longer dotted number and no candidate; such an
// address at the end of a sentence without a space is a known limit (brief 4.5).
func quadIn(line []byte, r span) (span, bool) {
	tok := bytes.TrimRight(line[r.start:r.end], ".")
	s := span{r.start, r.start + len(tok)}
	if len(tok) > 0 && tok[0] == '.' {
		if !betweenLabels(line, r) {
			return s, false
		}
		s.start++
	}
	return s, bytes.Count(line[s.start:s.end], []byte(".")) == 3
}

// betweenLabels reports whether run r is ".<digits and dots>." with a letter before and after.
func betweenLabels(line []byte, r span) bool {
	if r.end-r.start < 3 || r.start == 0 || r.end >= len(line) {
		return false
	}
	return line[r.start+1] != '.' && line[r.end-1] == '.' && line[r.end-2] != '.' &&
		isLetter(line[r.start-1]) && isLetter(line[r.end])
}

func isLetter(c byte) bool { return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' }

// parseQuad parses tok, four dot-separated runs of decimal digits, as an IPv4 address. Leading
// zeros are padding, read in decimal (011 is 11, never octal 9): 192.168.001.001 is a private
// address and a zero-padded public address hits. A part above 255 or an empty part is none.
func parseQuad(tok []byte) (netip.Addr, bool) {
	parts := bytes.Split(tok, []byte("."))
	if len(parts) != 4 {
		return netip.Addr{}, false
	}
	var b [4]byte
	for i, p := range parts {
		v, ok := quadPart(p)
		if !ok {
			return netip.Addr{}, false
		}
		b[i] = v
	}
	return netip.AddrFrom4(b), true
}

// quadPart reads one part of a dotted quad: one or more decimal digits with a value of 0..255.
func quadPart(p []byte) (byte, bool) {
	if len(p) == 0 {
		return 0, false
	}
	v := 0
	for _, c := range bytes.TrimLeft(p, "0") {
		if c < '0' || c > '9' {
			return 0, false
		}
		if v = v*10 + int(c-'0'); v > 255 {
			return 0, false
		}
	}
	return byte(v), true
}

// ipv6Finder finds tokens of hex digits, colons and dots with "::" or seven colons that parse as
// an IPv6 address and are not part of an identifier (as in std::string).
func ipv6Finder() func([]byte) []span {
	allowed, allowed4 := allowedIPv6(), allowedIPv4()
	return func(line []byte) []span { return findIPv6(line, allowed, allowed4) }
}

func findIPv6(line []byte, allowed, allowed4 []netip.Prefix) []span {
	if !bytes.Contains(line, []byte(":")) {
		return nil
	}
	var out []span
	for _, r := range runs(line, func(c byte) bool { return isHex(c) || c == ':' || c == '.' }) {
		tok := bytes.TrimRight(line[r.start:r.end], ".")
		if !bytes.Contains(tok, []byte("::")) && bytes.Count(tok, []byte(":")) < 7 {
			continue
		}
		if identChar(line, r.start-1) || identChar(line, r.end) {
			continue
		}
		addr, err := netip.ParseAddr(string(tok))
		if err != nil || !addr.Is6() {
			continue
		}
		if ipv6Allowed(addr, allowed, allowed4) {
			continue
		}
		out = append(out, span{r.start, r.start + len(tok)})
	}
	return out
}

func ipv6Allowed(addr netip.Addr, allowed, allowed4 []netip.Prefix) bool {
	if addr.Is4In6() {
		return inAny(addr.Unmap(), allowed4)
	}
	if addr.IsLinkLocalUnicast() {
		return !linkLocalIsPersonal(addr)
	}
	return inAny(addr, allowed)
}

// linkLocalIsPersonal: an EUI-64 interface id (ff:fe in the middle) made from a universally
// administered unicast MAC carries that MAC. In modified EUI-64 the universal/local bit is
// inverted, so a set bit 1 of byte 8 means the MAC was universally administered.
func linkLocalIsPersonal(addr netip.Addr) bool {
	b := addr.As16()
	return b[11] == 0xff && b[12] == 0xfe && b[8]&0x02 != 0 && b[8]&0x01 == 0
}

// identChar reports whether line[i] exists and continues an identifier.
func identChar(line []byte, i int) bool {
	if i < 0 || i >= len(line) {
		return false
	}
	c := line[i]
	return c == '_' || c >= 'g' && c <= 'z' || c >= 'G' && c <= 'Z'
}

// runs returns the maximal runs of bytes for which in is true.
func runs(line []byte, in func(byte) bool) []span {
	var out []span
	start := -1
	for i := 0; i <= len(line); i++ {
		inside := i < len(line) && in(line[i])
		switch {
		case inside && start < 0:
			start = i
		case !inside && start >= 0:
			out = append(out, span{start, i})
			start = -1
		}
	}
	return out
}
