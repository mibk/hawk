package compiler

import "github.com/mibk/hawk/value"

type Root struct {
	begin    []Stmt
	pActions []Stmt
	end      []Stmt

	vars map[string]value.Value
}

func NewRoot() *Root {
	return &Root{vars: make(map[string]value.Value)}
}

func (r Root) Begin() {
	for _, a := range r.begin {
		a.Exec()
	}
}

func (r Root) End() {
	for _, a := range r.end {
		a.Exec()
	}
}

func (r Root) AnyPatternActions() bool {
	return len(r.pActions) > 0 || len(r.end) > 0
}

func (r Root) Exec() {
	for _, a := range r.pActions {
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
