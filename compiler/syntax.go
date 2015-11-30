package compiler

import "github.com/mibk/hawk/value"

type Tree struct {
	begin    []Stmt
	pActions []Stmt
	end      []Stmt

	vars map[string]value.Value
}

func NewTree() *Tree {
	return &Tree{vars: make(map[string]value.Value)}
}

func (t Tree) Begin() {
	for _, a := range t.begin {
		a.Exec()
	}
}

func (t Tree) End() {
	for _, a := range t.end {
		a.Exec()
	}
}

func (t Tree) AnyPatternActions() bool {
	return len(t.pActions) > 0 || len(t.end) > 0
}

func (t Tree) Exec() {
	for _, a := range t.pActions {
		a.Exec()
	}
}

type BeginAction struct {
	Stmt
}

type EndAction struct {
	Stmt
}

type PatternAction struct {
	pattern Expr
	action  Stmt
}

func (p PatternAction) Exec() Status {
	if p.pattern == nil || p.pattern.Eval().Bool() {
		p.action.Exec()
	}
	return StatusNone
}
