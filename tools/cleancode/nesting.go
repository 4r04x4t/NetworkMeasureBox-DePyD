package main

import (
	"go/ast"
	"go/token"
)

// tooDeep returns the position of every construct that opens nesting level maxNesting+1 inside a
// function body. The body is level 0; if, for, range, switch, type switch, select, a bare block
// and a function literal each open one level; an else-if stays on the level of its if; case and
// default clauses add none. Constructs inside one that is already too deep are not reported again.
func tooDeep(body *ast.BlockStmt) []token.Pos {
	w := &nestWalker{elseIf: map[*ast.IfStmt]bool{}, bare: map[*ast.BlockStmt]bool{}}
	w.walk(body, 0)
	return w.found
}

type nestWalker struct {
	elseIf map[*ast.IfStmt]bool
	bare   map[*ast.BlockStmt]bool
	found  []token.Pos
}

// walk visits n at the given level and descends into its children with their levels.
func (w *nestWalker) walk(n ast.Node, level int) {
	w.markChildren(n)
	ast.Inspect(n, func(child ast.Node) bool {
		if child == nil || child == n {
			return child != nil
		}
		w.markChildren(child)
		if !w.opensLevel(child) {
			return true
		}
		if level+1 > maxNesting {
			w.found = append(w.found, child.Pos())
			return false
		}
		w.walk(child, level+1)
		return false
	})
}

// markChildren records the else-if statements and bare blocks below n before they are visited.
func (w *nestWalker) markChildren(n ast.Node) {
	switch s := n.(type) {
	case *ast.IfStmt:
		if next, ok := s.Else.(*ast.IfStmt); ok {
			w.elseIf[next] = true
		}
	case *ast.BlockStmt:
		w.markBare(s.List)
	case *ast.CaseClause:
		w.markBare(s.Body)
	case *ast.CommClause:
		w.markBare(s.Body)
	case *ast.LabeledStmt:
		if b, ok := s.Stmt.(*ast.BlockStmt); ok {
			w.bare[b] = true
		}
	}
}

func (w *nestWalker) markBare(list []ast.Stmt) {
	for _, st := range list {
		if b, ok := st.(*ast.BlockStmt); ok {
			w.bare[b] = true
		}
	}
}

func (w *nestWalker) opensLevel(n ast.Node) bool {
	switch s := n.(type) {
	case *ast.IfStmt:
		return !w.elseIf[s]
	case *ast.ForStmt, *ast.RangeStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt, *ast.SelectStmt, *ast.FuncLit:
		return true
	case *ast.BlockStmt:
		return w.bare[s]
	}
	return false
}
