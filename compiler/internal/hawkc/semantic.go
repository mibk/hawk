package hawkc

func analyse(prog *Program) {
	for _, fn := range prog.funcs {
		walkStmt(fn.Body, fn.scope)
	}
}

func walkStmt(stmt Stmt, scope Scope) {
	if stmt == nil {
		return
	}
	switch s := stmt.(type) {
	case *BlockStmt:
		for _, s := range s.List {
			walkStmt(s, scope)
		}
	case *AssignStmt:
		s.scope = scope
		walkExpr(s.Left, scope)
		walkExpr(s.Right, scope)
	case *IfStmt:
		walkExpr(s.X, scope)
		walkStmt(s.Body, scope)
		walkStmt(s.Else, scope)
	case *ForStmt:
		walkStmt(s.Init, scope)
		walkExpr(s.Cond, scope)
		walkStmt(s.Post, scope)
		walkStmt(s.Body, scope)
	case *ForeachStmt:
		walkExpr(s.Key, scope)
		walkExpr(s.Val, scope)
		walkStmt(s.Body, scope)
	case *ReturnStmt:
		walkExpr(s.X, scope)
	case *PrintStmt:
		for _, e := range s.Args {
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
		walkExpr(e.Cond, scope)
		walkExpr(e.Yes, scope)
		walkExpr(e.No, scope)
	case *CallExpr:
		for _, e := range e.Args {
			walkExpr(e, scope)
		}
	case *Ident:
		e.scope = scope
	case *FieldExpr:
		walkExpr(e.X, scope)
	case *BinaryExpr:
		walkExpr(e.X, scope)
		walkExpr(e.Y, scope)
	case *UnaryExpr:
		walkExpr(e.X, scope)
	}
}
