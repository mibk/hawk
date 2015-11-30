package compiler

type Status int

const (
	StatusNone Status = iota
	StatusBreak
	StatusContinue
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
		switch stmt.Exec() {
		case StatusBreak:
			return StatusBreak
		case StatusContinue:
			return StatusNone
		}
	}
	return StatusNone
}

type AssignStmt struct {
	tree *Tree
	name string
	expr Expr
}

func (a AssignStmt) Exec() Status {
	a.tree.vars[a.name] = a.expr.Eval()
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
		if f.body.Exec() == StatusBreak {
			break
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

type CallStmt CallExpr

func (c CallStmt) Exec() Status {
	CallExpr(c).Eval()
	return StatusNone
}
