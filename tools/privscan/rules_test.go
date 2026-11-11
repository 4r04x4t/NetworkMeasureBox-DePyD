package main

import (
	"strings"
	"testing"
)

// Positive values are assembled at run time from parts, so that no value the scanner must find
// is ever a literal in this source (the scanner runs over its own package).
func j(parts ...string) string { return strings.Join(parts, "") }

func scanText(t *testing.T, sc *Scanner, text string) []Hit {
	t.Helper()
	hits, err := sc.Scan("f", strings.NewReader(text))
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	return hits
}

func categories(hits []Hit) string {
	var out []string
	for _, h := range hits {
		out = append(out, h.Category)
	}
	return strings.Join(out, ",")
}

type ruleCase struct {
	name string
	text string
	want string // categories of the hits, comma-separated; empty for none
}

func runCases(t *testing.T, cases []ruleCase) {
	t.Helper()
	sc := NewScanner(nil, nil)
	for _, c := range cases {
		if got := categories(scanText(t, sc, c.text)); got != c.want {
			t.Errorf("%s: %q gave %q, want %q", c.name, c.text, got, c.want)
		}
	}
}

func TestEmail(t *testing.T) {
	runCases(t, []ruleCase{
		{"address", j("to: jane.roe", "@", "mailhost", ".net"), catEmail},
		{"subdomain of a real domain", j("x", "@", "mail.provider", ".com"), catEmail},
		{"reserved domain", "user@example.com and a@b.example.org", ""},
		{"reserved names", "a@host.example b@x.invalid c@y.test d@z.localhost", ""},
		{"module query", "golang.org/x/tools@v0.52.0 and tool@v1.2.3-rc.beta", ""},
		{"no-reply", j("noreply", "@", "vendor", ".com"), ""},
		{"no top-level name", "pkg@latest user@localhost", ""},
	})
}

func TestIPv4(t *testing.T) {
	runCases(t, []ruleCase{
		{"public", j("server ", "11.22", ".33.44", " up"), catIPv4},
		{"documentation", "192.0.2.1 198.51.100.7 203.0.113.255", ""},
		{"private and special", "10.1.2.3 172.16.0.1 192.168.1.1 127.0.0.1 169.254.1.1 0.0.0.0 255.255.255.255 224.0.0.1", ""},
		{"longer dotted number", j("1.2", ".3.4.5 ", "v11.22", ".33.44.55"), ""},
		{"not an address", "999.1.1.1 1.2.3 01.2.3.4", ""},
		{"end of sentence", j("at ", "11.22", ".33.44."), catIPv4},
	})
}

// F00-R20, question Q3: carrier-grade NAT and tailnet addresses are a hit.
func TestCGNATIsHit(t *testing.T) {
	runCases(t, []ruleCase{{"cgnat", j("100.", "100.1.2"), catIPv4}})
}

func TestIPv6(t *testing.T) {
	runCases(t, []ruleCase{
		{"global", j("addr 2a00", ":1450::", "200e"), catIPv6},
		{"full form", j("2a00:1450:4001", ":82a:0:0:0:200e"), catIPv6},
		{"documentation", "2001:db8::1 3fff::1 [2001:db8:0:1::2]:80", ""},
		{"special", ":: ::1 fd12:3456::1 ff02::1 64:ff9b::c000:201", ""},
		{"mapped documentation", "::ffff:192.0.2.1", ""},
		{"mapped public", j("::ffff:", "11.22", ".33.44"), "ipv4,ipv6"},
		{"identifier", "std::vector and Foo::Bar", ""},
		{"time", "12:34:56", ""},
	})
}

func TestEUI64LinkLocalFromRealMACIsHit(t *testing.T) {
	runCases(t, []ruleCase{
		{"universal", j("fe80::", "0200:5eff:", "fe00:5301"), catIPv6},
		{"locally administered", "fe80::ff:fe00:1", ""},
		{"not EUI-64", "fe80::1 fe80::1234:5678:9abc:def0", ""},
	})
}

func TestMAC(t *testing.T) {
	runCases(t, []ruleCase{
		{"universal colon", j("00:00:5e", ":00:53:01"), catMAC},
		{"universal dash", j("00-00-5E", "-00-53-01"), catMAC},
		{"universal dot", j("0000.5e", "00.5301"), catMAC},
		{"locally administered", "02:00:00:00:00:01 06-00-00-00-00-02", ""},
		{"multicast and zero", "01:00:5e:00:00:01 ff:ff:ff:ff:ff:ff 00:00:00:00:00:00", ""},
		{"inside a longer token", j("2001:db8:0:00:00", ":5e:00:53:01"), ""},
	})
}

func TestPath(t *testing.T) {
	runCases(t, []ruleCase{
		{"linux home", j("/ho", "me/", "jdoe", "/src"), catPath},
		{"mac home", j("/Us", "ers/", "jdoe/x"), catPath},
		{"windows", j(`C:\`, `Users\`, `Jane Roe\Downloads`), catPath},
		{"wsl mount", j("/mnt/c/", "Users/", "jdoe/"), catPath},
		{"allowed", `/home/nmb/ /home/<user>/ /home/$USER/ C:\Users\Public\ ~/src C:\Users\%USERNAME%`, ""},
	})
}

func TestKey(t *testing.T) {
	alnum36 := strings.Repeat("aB3dE6gH9", 4)
	runCases(t, []ruleCase{
		{"private key block", j("-----BEGIN ", "OPENSSH PRIVATE", " KEY-----"), catKey},
		{"github token", j("gh", "p_", alnum36), catKey},
		{"aws", j("AK", "IA", "ABCDEFGH23456789"), catKey},
		{"telegram", j("1234567890", ":", alnum36[:35]), catKey},
		{"api key", j("s", "k-", "proj-", alnum36[:24]), catKey},
		{"wireguard", j("PrivateKey = ", strings.Repeat("aB3+", 10), "aB3="), catKey},
		{"quoted assignment", j(`password = "`, "Zq8#Lw2!Rt5x", `"`), catKey},
		{"config assignment", j("TELEGRAM_TOKEN=", "Zq8kLw2mRt5xV7"), catKey},
		{"placeholders", `token = "<your token>" password: "${PASSWORD}" secret=__KEEP_CURRENT__ api_key="example-key-123456"`, ""},
		{"low entropy", `password = "aaaaaaaaaaaaaaaa1"`, ""},
		{"code", "token := strings.TrimSpace(line) // secret handling", ""},
		{"go.sum hash", "golang.org/x/mod v0.42.0 h1:abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQ=", ""},
	})
}

func TestEntropy(t *testing.T) {
	if e := entropy([]byte("aaaa")); e != 0 {
		t.Errorf("entropy of a repeated byte = %v", e)
	}
	if e := entropy([]byte("abcd")); e != 2 {
		t.Errorf("entropy of four distinct bytes = %v", e)
	}
}
