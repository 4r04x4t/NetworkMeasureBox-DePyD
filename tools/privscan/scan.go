package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// maxFileSize is the largest input scanned; a bigger one is an error, never a pass.
const maxFileSize = 256 << 20

// maskedPrefix is how many bytes of a hit are shown; the rest is never printed.
const maskedPrefix = 2

// Hit is one finding. Its String form is masked so that logs never repeat a secret.
type Hit struct {
	Path     string
	Line     int
	Col      int
	Category string
	value    []byte
	private  int // line number in the private list, for category private
}

func (h Hit) String() string {
	return fmt.Sprintf("%s:%d:%d: %s: %s", h.Path, h.Line, h.Col, h.Category, h.masked())
}

func (h Hit) masked() string {
	if h.Category == catPrivate {
		return "private value #" + strconv.Itoa(h.private)
	}
	n := min(maskedPrefix, len(h.value))
	prefix := h.value[:n]
	for _, c := range prefix {
		if c < 0x21 || c > 0x7e {
			prefix = []byte("??")[:n]
			break
		}
	}
	return fmt.Sprintf("%s…(%d)", prefix, len(h.value))
}

// Scanner applies the rules, the allow list and the private list to its inputs.
type Scanner struct {
	rules   []rule
	allow   map[string]map[string]bool // category -> exact value
	private []privateValue
}

// NewScanner returns a scanner with every rule, the given allow list and private values.
func NewScanner(allow map[string]map[string]bool, private []privateValue) *Scanner {
	return &Scanner{rules: newRules(), allow: allow, private: private}
}

// Scan reads r line by line (no line-length limit) and returns its hits.
func (s *Scanner) Scan(path string, r io.Reader) ([]Hit, error) {
	br := bufio.NewReader(io.LimitReader(r, maxFileSize+1))
	var hits []Hit
	read := 0
	for lineNo := 1; ; lineNo++ {
		line, err := br.ReadBytes('\n')
		read += len(line)
		if read > maxFileSize {
			return nil, fmt.Errorf("%s: larger than %d bytes", path, maxFileSize)
		}
		hits = append(hits, s.scanLine(path, lineNo, line)...)
		if errors.Is(err, io.EOF) {
			return hits, nil
		}
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", path, err)
		}
	}
}

func (s *Scanner) scanLine(path string, lineNo int, line []byte) []Hit {
	var hits []Hit
	for _, r := range s.rules {
		for _, sp := range r.find(line) {
			v := line[sp.start:sp.end]
			if s.allow[r.category][string(v)] {
				continue
			}
			hits = append(hits, Hit{Path: path, Line: lineNo, Col: sp.start + 1, Category: r.category, value: v})
		}
	}
	if len(s.private) == 0 {
		return hits
	}
	lower := bytes.ToLower(line)
	for _, p := range s.private {
		if i := bytes.Index(lower, p.lower); i >= 0 {
			hits = append(hits, Hit{Path: path, Line: lineNo, Col: i + 1, Category: catPrivate, private: p.line})
		}
	}
	return hits
}
