package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// Limits of the clean-code rules (roadmap Phase 0, "Clean code").
const (
	maxFuncLines = 50
	maxFileLines = 500
	maxNesting   = 3

	allowLongFunc = "//nmb:allow-long"
	allowLongFile = "nmb:allow-long-file"
	headerLines   = 5 // a text file's directive must be in its first lines
)

// Finding is one result: a violation, or an exception granted by a directive with a reason.
type Finding struct {
	Path   string
	Line   int
	Rule   string
	Value  int
	Limit  int
	Reason string // set for exceptions only
}

func (f Finding) String() string {
	s := fmt.Sprintf("%s:%d: %s: %d (limit %d)", f.Path, f.Line, f.Rule, f.Value, f.Limit)
	if f.Reason != "" {
		s += ": allowed: " + f.Reason
	}
	return s
}

// Result holds the violations and the granted exceptions of one file.
type Result struct {
	Violations []Finding
	Exceptions []Finding
}

func (r *Result) add(f Finding, reason string, directive bool) {
	switch {
	case !directive:
		r.Violations = append(r.Violations, f)
	case reason == "":
		f.Rule += " (directive without a reason)"
		r.Violations = append(r.Violations, f)
	default:
		f.Reason = reason
		r.Exceptions = append(r.Exceptions, f)
	}
}

// CheckGo applies every rule to one Go source file. A file that cannot be parsed is an error.
func CheckGo(path string, src []byte) (Result, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return Result{}, fmt.Errorf("parsing %s: %w", path, err)
	}
	var res Result
	checkFileLength(&res, path, src, fileDirective(file, "//"))
	for _, decl := range file.Decls {
		for _, fn := range functionsOf(decl) {
			checkFunction(&res, fset, path, fn)
		}
	}
	return res, nil
}

// CheckText applies the file-length rule to a file that is not Go (shell scripts).
func CheckText(path string, src []byte) Result {
	var res Result
	var d directive
	lines := strings.SplitN(string(src), "\n", headerLines+1)
	for _, line := range lines[:min(headerLines, len(lines))] {
		if r, ok := directiveReason(strings.TrimSpace(line), "# "+allowLongFile); ok {
			d = directive{reason: r, found: true}
		}
	}
	checkFileLength(&res, path, src, d)
	return res
}

type directive struct {
	reason string
	found  bool
}

func checkFileLength(res *Result, path string, src []byte, d directive) {
	lines := bytes.Count(src, []byte("\n"))
	if len(src) > 0 && src[len(src)-1] != '\n' {
		lines++
	}
	if lines > maxFileLines {
		res.add(Finding{Path: path, Line: 1, Rule: "file length", Value: lines, Limit: maxFileLines}, d.reason, d.found)
	}
}

// fileDirective finds "//nmb:allow-long-file <reason>" in a comment before the package clause.
func fileDirective(file *ast.File, prefix string) directive {
	for _, group := range file.Comments {
		if group.Pos() > file.Package {
			break
		}
		for _, c := range group.List {
			if r, ok := directiveReason(c.Text, prefix+allowLongFile); ok {
				return directive{reason: r, found: true}
			}
		}
	}
	return directive{}
}

// directiveReason reports whether text is the directive and returns its reason (maybe empty).
func directiveReason(text, name string) (string, bool) {
	if text != name && !strings.HasPrefix(text, name+" ") {
		return "", false
	}
	return strings.TrimSpace(strings.TrimPrefix(text, name)), true
}

// function is a body to check, with the doc comment that may carry an allow-long directive.
type function struct {
	name string
	doc  *ast.CommentGroup
	body *ast.BlockStmt
}

// functionsOf returns the function declarations and the package-level function literals of decl.
func functionsOf(decl ast.Decl) []function {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		if d.Body == nil {
			return nil
		}
		return []function{{name: d.Name.Name, doc: d.Doc, body: d.Body}}
	case *ast.GenDecl:
		var out []function
		for _, spec := range d.Specs {
			out = append(out, literalsOf(spec, d.Doc)...)
		}
		return out
	}
	return nil
}

func literalsOf(spec ast.Spec, declDoc *ast.CommentGroup) []function {
	vs, ok := spec.(*ast.ValueSpec)
	if !ok {
		return nil
	}
	doc := vs.Doc
	if doc == nil {
		doc = declDoc
	}
	var out []function
	for i, v := range vs.Values {
		if lit, ok := v.(*ast.FuncLit); ok && i < len(vs.Names) {
			out = append(out, function{name: vs.Names[i].Name, doc: doc, body: lit.Body})
		}
	}
	return out
}

func checkFunction(res *Result, fset *token.FileSet, path string, fn function) {
	start := fset.Position(fn.body.Lbrace).Line
	end := fset.Position(fn.body.Rbrace).Line
	if lines := end - start - 1; lines > maxFuncLines {
		d := funcDirective(fn.doc)
		f := Finding{Path: path, Line: start, Rule: "function length " + fn.name, Value: lines, Limit: maxFuncLines}
		res.add(f, d.reason, d.found)
	}
	for _, pos := range tooDeep(fn.body) {
		res.add(Finding{Path: path, Line: fset.Position(pos).Line, Rule: "nesting " + fn.name,
			Value: maxNesting + 1, Limit: maxNesting}, "", false)
	}
}

func funcDirective(doc *ast.CommentGroup) directive {
	if doc == nil {
		return directive{}
	}
	for _, c := range doc.List {
		if r, ok := directiveReason(c.Text, allowLongFunc); ok {
			return directive{reason: r, found: true}
		}
	}
	return directive{}
}
