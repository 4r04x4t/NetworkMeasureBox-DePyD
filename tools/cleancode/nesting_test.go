package main

import (
	"strings"
	"testing"
)

// nested returns a function body with depth nested if statements.
func nested(depth int) string {
	return "package p\n\nfunc f(x int) {\n" + strings.Repeat("if x > 0 {\n", depth) + "x--\n" +
		strings.Repeat("}\n", depth) + "}\n"
}

func nestingViolations(t *testing.T, body string) []Finding {
	t.Helper()
	res := check(t, "package p\n\nfunc f(x int, c chan int) {\n"+body+"\n}\n")
	var out []Finding
	for _, v := range res.Violations {
		if strings.HasPrefix(v.Rule, "nesting") {
			out = append(out, v)
		}
	}
	return out
}

func TestNestingThreePasses(t *testing.T) {
	if res := check(t, nested(3)); len(res.Violations) != 0 {
		t.Errorf("violations: %v", res.Violations)
	}
}

func TestNestingFourFails(t *testing.T) {
	res := check(t, nested(4))
	if len(res.Violations) != 1 || res.Violations[0].Line != 7 || !strings.HasPrefix(res.Violations[0].Rule, "nesting f") {
		t.Errorf("violations: %v", res.Violations)
	}
	if res := check(t, nested(6)); len(res.Violations) != 1 {
		t.Errorf("deeper constructs reported again: %v", res.Violations)
	}
}

func TestNestingHasNoException(t *testing.T) {
	res := check(t, strings.Replace(nested(4), "func f", "//nmb:allow-long reason\nfunc f", 1))
	if len(res.Violations) != 1 || len(res.Exceptions) != 0 {
		t.Errorf("result: %+v", res)
	}
}

func TestElseIfChainStaysOnLevel(t *testing.T) {
	body := "for {\nfor {\nif x == 1 {\n} else if x == 2 {\n} else if x == 3 {\nx++\n} else {\nx--\n}\n}\n}"
	if v := nestingViolations(t, body); len(v) != 0 {
		t.Errorf("violations: %v", v)
	}
	if v := nestingViolations(t, "for {\nfor {\nif x == 1 {\n} else {\nif x == 2 {\n}\n}\n}\n}"); len(v) != 1 {
		t.Errorf("if inside else must add a level: %v", v)
	}
}

func TestCaseDoesNotAddLevel(t *testing.T) {
	body := "switch x {\ncase 1:\nfor {\nselect {\ncase <-c:\n}\n}\ndefault:\n}"
	if v := nestingViolations(t, body); len(v) != 0 {
		t.Errorf("violations: %v", v)
	}
	body = "switch x {\ncase 1:\nfor {\nselect {\ncase <-c:\nif x > 0 {\n}\n}\n}\n}"
	if v := nestingViolations(t, body); len(v) != 1 {
		t.Errorf("fourth level inside cases: %v", v)
	}
}

func TestFuncLitAddsLevel(t *testing.T) {
	if v := nestingViolations(t, "for {\nif x > 0 {\ngo func() {\n}()\n}\n}"); len(v) != 0 {
		t.Errorf("three levels: %v", v)
	}
	if v := nestingViolations(t, "for {\nif x > 0 {\ngo func() {\nif x > 1 {\n}\n}()\n}\n}"); len(v) != 1 {
		t.Errorf("four levels through a literal: %v", v)
	}
}

func TestBareBlockAndRangeAndTypeSwitchAddLevel(t *testing.T) {
	cases := []string{
		"{\n{\n{\n{\n}\n}\n}\n}",
		"for range c {\nfor range c {\nfor range c {\nfor range c {\n}\n}\n}\n}",
		"var y any\nswitch y.(type) {\ndefault:\nswitch y.(type) {\ndefault:\nfor {\nif x > 0 {\n}\n}\n}\n}",
		"L:\n{\n{\n{\n{\n}\n}\n}\n}\ngoto L",
	}
	for _, body := range cases {
		if v := nestingViolations(t, body); len(v) != 1 {
			t.Errorf("%q: violations %v", body, v)
		}
	}
}
