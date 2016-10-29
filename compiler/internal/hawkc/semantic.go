package hawkc

import "fmt"

type Analyser struct {
	scope Scope
}

func analyse(prog *Program) {
	a := &Analyser{prog}
	for _, p := range prog.Begins {
		a.walkPactions(p)
	}
	for _, p := range prog.Pactions {
		a.walkPactions(p)
	}
	for _, p := range prog.Ends {
		a.walkPactions(p)
	}
	for _, fn := range prog.funcs {
		a.scope = fn.scope
		a.walkStmt(fn.Body)
	}
}

func (a *Analyser) walkPactions(pa Stmt) {
	if pa == nil {
		return
	}
	switch pa := pa.(type) {
	case *BeginAction:
		a.walkStmt(pa.Stmt)
	case *EndAction:
		a.walkStmt(pa.Stmt)
	case *PatternAction:
		a.walkExpr(pa.X)
		a.walkStmt(pa.Body)
	default:
		panic(fmt.Sprintf("unknown pattern-action: %T", pa))
	}
}

func (a *Analyser) walkStmt(s Stmt) {
	if s == nil {
		return
	}
	switch s := s.(type) {
	case *ExprStmt:
		a.walkExpr(s.X)
	case *BlockStmt:
		for _, s := range s.List {
			a.walkStmt(s)
		}
	case *PipeStmt:
		a.walkStmt(s.Stmt)
	case *AssignStmt:
		s.scope = a.scope
		a.walkExpr(s.Left)
		a.walkExpr(s.Right)
	case *IfStmt:
		a.walkExpr(s.X)
		a.walkStmt(s.Body)
		a.walkStmt(s.Else)
	case *ForStmt:
		a.walkStmt(s.Init)
		a.walkExpr(s.Cond)
		a.walkStmt(s.Post)
		a.walkStmt(s.Body)
	case *ForeachStmt:
		a.walkExpr(s.Key)
		if s.Val != nil {
			a.walkExpr(s.Val)
		}
		a.walkExpr(s.X)
		a.walkStmt(s.Body)
	case *StatusStmt:
	case *ReturnStmt:
		a.walkExpr(s.X)
	case *PrintStmt:
		for _, e := range s.Args {
			a.walkExpr(e)
		}
	}
}

func (a *Analyser) walkExpr(e Expr) {
	if e == nil {
		return
	}
	switch e := e.(type) {
	case *TernaryExpr:
		a.walkExpr(e.Cond)
		a.walkExpr(e.Yes)
		a.walkExpr(e.No)
	case *CallExpr:
		for _, e := range e.Args {
			a.walkExpr(e)
		}
	case *Ident:
		e.scope = a.scope
	case *FieldExpr:
		a.walkExpr(e.X)
	case *IndexExpr:
		a.walkExpr(e.Index)
		a.walkExpr(e.X)
	case *BinaryExpr:
		a.walkExpr(e.X)
		a.walkExpr(e.Y)
	case *UnaryExpr:
		a.walkExpr(e.X)
	case *MatchExpr:
		a.walkExpr(e.X)
		a.walkExpr(e.Y)
	case *ArrayLit:
		for _, e := range e.Elems {
			a.walkExpr(e)
		}
	}
}
