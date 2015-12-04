package compiler

import (
	"bufio"
	"io"
	"strings"

	"github.com/mibk/hawk/parse"
	"github.com/mibk/hawk/value"
)

type Program struct {
	parser   *parse.Parser
	begin    []Stmt
	pActions []Stmt
	end      []Stmt

	vars map[string]*value.Value
}

func NewProgram(p *parse.Parser) *Program {
	return &Program{parser: p, vars: make(map[string]*value.Value)}
}

func (p Program) Run(in io.Reader) {
	p.Begin()
	if p.anyPatternActions() {
		in := bufio.NewReader(in)
		for {
			line, err := in.ReadBytes('\n')
			if err != nil {
				break
			}
			p.parser.SetFields(strings.Fields(string(line)))
			p.Exec()
		}
		p.End()
	}
}

func (p Program) Begin() {
	for _, a := range p.begin {
		a.Exec()
	}
}

func (p Program) End() {
	for _, a := range p.end {
		a.Exec()
	}
}

func (p Program) anyPatternActions() bool {
	return len(p.pActions) > 0 || len(p.end) > 0
}

func (p Program) Exec() {
	for _, a := range p.pActions {
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
