package hawkc

func analyse(prog *Program) {
	for _, fn := range prog.funcs {
		walkStmt(fn.body, fn.scope)
	}
}

func walkStmt(stmt Stmt, scope Scope) {
	if stmt == nil {
		return
	}
	switch s := stmt.(type) {
	case *BlockStmt:
		for _, s := range s.stmts {
			walkStmt(s, scope)
		}
	case *AssignStmt:
		s.scope = scope
		walkExpr(s.left, scope)
		walkExpr(s.right, scope)
	case *IfStmt:
		walkExpr(s.expr, scope)
		walkStmt(s.stmt, scope)
		walkStmt(s.elseStmt, scope)
	case *ForStmt:
		walkStmt(s.init, scope)
		walkExpr(s.cond, scope)
		walkStmt(s.step, scope)
		walkStmt(s.body, scope)
	case *ForeachStmt:
		walkExpr(s.key, scope)
		walkExpr(s.val, scope)
		walkStmt(s.body, scope)
	case *ReturnStmt:
		walkExpr(s.expr, scope)
	case *PrintStmt:
		for _, e := range s.args {
			walkExpr(e, scope)
		}
	}
}

func walkExpr(expr Expr, scope Scope) {
	if expr == nil {
		return
	}
	switch e := expr.(type) {
	case *TernaryExpr:
		walkExpr(e.cond, scope)
		walkExpr(e.yes, scope)
		walkExpr(e.no, scope)
	case *CallExpr:
		for _, e := range e.args {
			walkExpr(e, scope)
		}
	case *Ident:
		e.scope = scope
	case *FieldExpr:
		walkExpr(e.num, scope)
	case *BinaryExpr:
		walkExpr(e.left, scope)
		walkExpr(e.right, scope)
	case *UnaryExpr:
		walkExpr(e.expr, scope)
	}
}
