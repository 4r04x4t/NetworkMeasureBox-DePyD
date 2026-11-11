package main

import (
	"bytes"
	"math"
	"regexp"
	"strings"
)

// Categories of findings; the allow list names them.
const (
	catEmail   = "email"
	catIPv4    = "ipv4"
	catIPv6    = "ipv6"
	catMAC     = "mac"
	catKey     = "key"
	catPath    = "path"
	catPrivate = "private"
	catName    = "name" // author or committer name of a commit (history mode)
)

// Limits of the secret-assignment rule.
const (
	minSecretLen     = 12
	minSecretEntropy = 3.5 // bits per character
)

// span is a finding inside one line: byte offsets [start, end).
type span struct{ start, end int }

// rule finds the hits of one category in one line; built-in allowances are applied by find.
type rule struct {
	category string
	find     func(line []byte) []span
}

// newRules compiles every rule. The expressions are RE2, so matching is linear in the input.
func newRules() []rule {
	return []rule{
		{category: catEmail, find: emailFinder()},
		{category: catIPv4, find: ipv4Finder()},
		{category: catIPv6, find: ipv6Finder()},
		{category: catMAC, find: macFinder()},
		{category: catKey, find: keyFinder()},
		{category: catPath, find: pathFinder()},
	}
}

func emailFinder() func([]byte) []span {
	re := regexp.MustCompile(`[A-Za-z0-9._%+-]+@((?:[A-Za-z0-9-]+\.)+[A-Za-z]{2,63})`)
	// A Go module query names a semantic version (https://go.dev/ref/mod#versions), whole.
	version := regexp.MustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+(?:-[0-9A-Za-z.-]+)?$`)
	return func(line []byte) []span { return findEmails(re, version, line) }
}

func findEmails(re, version *regexp.Regexp, line []byte) []span {
	var out []span
	for _, m := range re.FindAllSubmatchIndex(line, -1) {
		domain := line[m[2]:m[3]]
		if !version.Match(domain) && !emailAllowed(domain) {
			out = append(out, span{m[0], m[1]})
		}
	}
	return out
}

// emailAllowed: reserved domains and top-level names (RFC 2606, RFC 6761). Any other address,
// no-reply ones included, is a hit unless the reviewed allow list names it exactly.
func emailAllowed(domain []byte) bool {
	d := string(bytes.ToLower(domain))
	for _, reserved := range []string{"example.com", "example.net", "example.org"} {
		if d == reserved || hasSuffix(d, "."+reserved) {
			return true
		}
	}
	for _, tld := range []string{".example", ".invalid", ".test", ".localhost"} {
		if hasSuffix(d, tld) {
			return true
		}
	}
	return false
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// macFinder finds MAC addresses in the colon, dash and dot forms that are universally
// administered unicast addresses; locally administered, multicast and all-zero ones are allowed.
func macFinder() func([]byte) []span {
	forms := []macForm{
		{regexp.MustCompile(`[0-9A-Fa-f]{2}(?::[0-9A-Fa-f]{2}){5}`), ':'},
		{regexp.MustCompile(`[0-9A-Fa-f]{2}(?:-[0-9A-Fa-f]{2}){5}`), '-'},
		{regexp.MustCompile(`[0-9A-Fa-f]{4}\.[0-9A-Fa-f]{4}\.[0-9A-Fa-f]{4}`), '.'},
	}
	return func(line []byte) []span { return findMACs(forms, line) }
}

// macForm is one way of writing a MAC address and its separator.
type macForm struct {
	re  *regexp.Regexp
	sep byte
}

func findMACs(forms []macForm, line []byte) []span {
	var out []span
	for _, f := range forms {
		for _, m := range f.re.FindAllIndex(line, -1) {
			if bounded(line, m[0], m[1], f.sep) && macIsPersonal(line[m[0]:m[1]]) {
				out = append(out, span{m[0], m[1]})
			}
		}
	}
	return out
}

// bounded reports whether line[start:end] is not part of a longer hex token joined by sep.
func bounded(line []byte, start, end int, sep byte) bool {
	if start > 0 && (isHex(line[start-1]) || line[start-1] == sep) {
		return false
	}
	return end >= len(line) || !(isHex(line[end]) || line[end] == sep)
}

func isHex(c byte) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}

func hexVal(c byte) byte {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	default:
		return c - 'A' + 10
	}
}

// macIsPersonal: the first octet decides; bit 0 is multicast, bit 1 locally administered.
func macIsPersonal(mac []byte) bool {
	first := hexVal(mac[0])<<4 | hexVal(mac[1])
	if first&0x03 != 0 {
		return false
	}
	for _, c := range mac {
		if isHex(c) && c != '0' {
			return true
		}
	}
	return false
}

// keyFinder finds private keys, public keys, API tokens and secret assignments.
func keyFinder() func([]byte) []span {
	whole := []keyPattern{
		{re: regexp.MustCompile(`-----BEGIN (?:[A-Z0-9]+ )*PRIVATE KEY-----`)},
		{re: regexp.MustCompile(`(?:ssh-ed25519|ssh-rsa|ecdsa-sha2-nistp[0-9]+) AAAA[0-9A-Za-z+/]{40,}={0,2}`)},
		{re: regexp.MustCompile(`AGE-SECRET-KEY-1[0-9A-Z]{50,}`)},
		{re: regexp.MustCompile(`\bgh[pousr]_[A-Za-z0-9]{36}\b`)},
		{re: regexp.MustCompile(`\bgithub_pat_[A-Za-z0-9_]{22,}`)},
		{re: regexp.MustCompile(`\bglpat-[A-Za-z0-9_-]{20,}`)},
		{re: regexp.MustCompile(`\bxox[abprs]-[A-Za-z0-9-]{10,}`)},
		{re: regexp.MustCompile(`\bAKIA[0-9A-Z]{16}\b`)},
		{re: regexp.MustCompile(`\bAIza[0-9A-Za-z_-]{35}`)},
		// Telegram bot token: bounded in code, because \b misses a token that ends in '-' or
		// follows "bot" in an API address.
		{re: regexp.MustCompile(`[0-9]{8,10}:[A-Za-z0-9_-]{35}`), notBefore: digits, notAfter: idChars},
		{re: regexp.MustCompile(`\btskey-[A-Za-z0-9-]{16,}`)},
		{re: regexp.MustCompile(`(?:^|[^A-Za-z0-9_-])(sk-[A-Za-z0-9_-]*[0-9][A-Za-z0-9_-]*)`), min: 20},
		{re: regexp.MustCompile(`\beyJ[A-Za-z0-9_-]{10,}\.eyJ[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}`)},
		{re: regexp.MustCompile(`(?:PrivateKey|PresharedKey|PublicKey)\s*=\s*[A-Za-z0-9+/]{43}=`)},
	}
	assign := []*regexp.Regexp{
		regexp.MustCompile(`(?i)(?:password|passwd|secret|token|api_?key)[a-z0-9_]*["']?\s*[:=]+\s*["']([^"'\s]+)["']`),
		regexp.MustCompile(`(?i)^\s*(?:export\s+)?[a-z0-9_]*(?:password|passwd|secret|token|api_?key)[a-z0-9_]*\s*[:=]\s*([^\s"'#]+)\s*$`),
	}
	return func(line []byte) []span {
		return append(findKeys(whole, line), findAssignments(assign, line)...)
	}
}

// keyPattern is one key or token form; min is the shortest hit, and the span of group 1 counts
// when the expression has one. A match touched by a byte of notBefore or notAfter is part of a
// longer token and no hit.
type keyPattern struct {
	re                  *regexp.Regexp
	min                 int
	notBefore, notAfter string
}

// Byte sets for keyPattern boundaries.
const (
	digits  = "0123456789"
	idChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"
)

func findKeys(patterns []keyPattern, line []byte) []span {
	var out []span
	for _, p := range patterns {
		for _, m := range p.re.FindAllSubmatchIndex(line, -1) {
			s := span{m[0], m[1]}
			if len(m) > 2 && m[2] >= 0 {
				s = span{m[2], m[3]}
			}
			if s.end-s.start >= p.min && !isPlaceholder(line[s.start:s.end]) && p.bounded(line, s) {
				out = append(out, s)
			}
		}
	}
	return out
}

// bounded reports whether s is not touched by a byte of the pattern's boundary sets.
func (p keyPattern) bounded(line []byte, s span) bool {
	if s.start > 0 && strings.IndexByte(p.notBefore, line[s.start-1]) >= 0 {
		return false
	}
	return s.end >= len(line) || strings.IndexByte(p.notAfter, line[s.end]) < 0
}

func findAssignments(res []*regexp.Regexp, line []byte) []span {
	var out []span
	for _, re := range res {
		for _, m := range re.FindAllSubmatchIndex(line, -1) {
			v := line[m[2]:m[3]]
			if len(v) >= minSecretLen && !isPlaceholder(v) && entropy(v) >= minSecretEntropy {
				out = append(out, span{m[2], m[3]})
			}
		}
	}
	return out
}

// isPlaceholder: values that only stand for a secret.
func isPlaceholder(v []byte) bool {
	lower := bytes.ToLower(v)
	for _, marker := range []string{"example", "<", "${", "__keep_current__", "xxxxxxxx"} {
		if bytes.Contains(lower, []byte(marker)) {
			return true
		}
	}
	return false
}

// entropy returns the Shannon entropy of v in bits per byte.
func entropy(v []byte) float64 {
	var counts [256]int
	for _, c := range v {
		counts[c]++
	}
	var h float64
	for _, n := range counts {
		if n > 0 {
			p := float64(n) / float64(len(v))
			h -= p * math.Log2(p)
		}
	}
	return h
}

// pathFinder finds developer home folders; the service user nmb, shared profiles and
// placeholders such as <user>, $USER or %USERNAME% are allowed (they do not match the name set).
func pathFinder() func([]byte) []span {
	res := []*regexp.Regexp{
		regexp.MustCompile(`/home/([A-Za-z0-9._-]+)`),
		regexp.MustCompile(`/Users/([A-Za-z0-9._-]+)`),
		regexp.MustCompile(`(?i)\b[a-z]:\\+Users\\+([A-Za-z0-9._ -]*[A-Za-z0-9._-])`),
	}
	return func(line []byte) []span { return findPaths(res, line) }
}

func findPaths(res []*regexp.Regexp, line []byte) []span {
	var out []span
	for _, re := range res {
		for _, m := range re.FindAllSubmatchIndex(line, -1) {
			if !pathNameAllowed(string(line[m[2]:m[3]])) {
				out = append(out, span{m[0], m[1]})
			}
		}
	}
	return out
}

func pathNameAllowed(name string) bool {
	switch name {
	case "nmb", "Public", "Default", "Shared", "All Users":
		return true
	}
	return false
}
