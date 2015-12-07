package compiler

type Status int

const (
	StatusNone Status = iota
	StatusBreak
	StatusContinue
	StatusReturn
)

type Stmt interface {
	Exec() Status
}

type ExprStmt struct {
	expr Expr
}

func (e ExprStmt) Exec() Status {
	e.expr.Eval()
	return StatusNone
}

type BlockStmt struct {
	stmts []Stmt
}

func (b BlockStmt) Exec() Status {
	for _, stmt := range b.stmts {
		switch s := stmt.Exec(); s {
		case StatusBreak, StatusReturn:
			return s
		case StatusContinue:
			return StatusNone
		}
	}
	return StatusNone
}

type AssignStmt struct {
	scope Scope
	name  string
	expr  Expr
}

func (a AssignStmt) Exec() Status {
	a.scope.SetVar(a.name, a.expr.Eval())
	return StatusNone
}

type IfStmt struct {
	expr     Expr
	stmt     Stmt
	elseStmt Stmt
}

func (i IfStmt) Exec() Status {
	if i.expr.Eval().Bool() {
		return i.stmt.Exec()
	} else if i.elseStmt != nil {
		return i.elseStmt.Exec()
	}
	return StatusNone
}

type ForStmt struct {
	init Stmt
	cond Expr
	step Stmt
	body Stmt
}

func (f ForStmt) Exec() Status {
	if f.init != nil {
		f.init.Exec()
	}
	for f.cond == nil || f.cond.Eval().Bool() {
		switch f.body.Exec() {
		case StatusBreak:
			break
		case StatusReturn:
			return StatusReturn
		}
		if f.step != nil {
			f.step.Exec()
		}
	}
	return StatusNone
}

type StatusStmt struct {
	status Status
}

func (s StatusStmt) Exec() Status {
	return s.status
}

type ReturnStmt struct {
	tree *Program
	expr Expr
}

func (r ReturnStmt) Exec() Status {
	if r.expr != nil {
		r.tree.retval = r.expr.Eval()
	}
	return StatusReturn
}

type CallStmt CallExpr

func (c CallStmt) Exec() Status {
	CallExpr(c).Eval()
	return StatusNone
}
