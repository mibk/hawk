%{
package compiler

import (
	"github.com/mibk/hawk/parse"
)

var (
	parser *parse.Parser
	ast    *Tree

	defaultAction BlockStmt
)

%}

%union {
	num      int
	sym      string
	expr     Expr
	exprlist []Expr
	stmt     Stmt
	stmtlist []Stmt
}

%type <expr>     expr uexpr
%type <exprlist> exprlist
%type <stmt>     stmt paction
%type <stmtlist> stmtlist blockstmt pactionlist

%token <num> NUM
%token <sym> IDENT
%token       BEGIN END

%left EQ NE LE GE LT GT

%start top

%%

top:
	pactionlist
	{
		ast = &Tree{}
		for i := 0; i < len($1); {
			pa := $1[i]
			switch pa.(type) {
			case BeginAction:
				ast.begin = append(ast.begin, pa)
				goto del
			case EndAction:
				ast.end = append(ast.end, pa)
				goto del
			}
			i++
			continue
		del:
			$1 = append($1[:i], $1[i+1:]...)
		}
		ast.pActions = $1
	}

pactionlist:
	paction
	{
		$$ = []Stmt{$1}
	}
|	pactionlist ';' paction
	{
		$$ = append($1, $3)
	}

paction:
	expr blockstmt
	{
		$$ = PatternAction{$1, BlockStmt{$2}}
	}
|	expr
	{
		$$ = PatternAction{$1, defaultAction}
	}
|	blockstmt
	{
		$$ = PatternAction{Lit(1), BlockStmt{$1}}
	}
|	BEGIN blockstmt
	{
		$$ = BeginAction{BlockStmt{$2}}
	}
|	END blockstmt
	{
		$$ = EndAction{BlockStmt{$2}}
	}

blockstmt:
	'{' stmtlist ';' '}'
	{
		$$ = $2
	}

stmtlist:
	{
		$$ = nil
	}
|	stmt
	{
		$$ = []Stmt{$1}
	}
|	stmtlist ';' stmt
	{
		$$ = append($1, $3)
	}

stmt:
	expr
	{
		$$ = ExprStmt{$1}
	}



expr:
	uexpr
	{
		$$ = $1
	}
|	'$' uexpr
	{
		$$ = Col{parser, $2}
	}
|	expr EQ expr
	{
		$$ = BinaryOp{EQ, $1, $3}
	}
|	expr NE expr
	{
		$$ = BinaryOp{NE, $1, $3}
	}
|	expr LE expr
	{
		$$ = BinaryOp{LE, $1, $3}
	}
|	expr GE expr
	{
		$$ = BinaryOp{GE, $1, $3}
	}
|	expr LT expr
	{
		$$ = BinaryOp{LT, $1, $3}
	}
|	expr GT expr
	{
		$$ = BinaryOp{GT, $1, $3}
	}

uexpr:
	NUM
	{
		$$ = Lit($1)
	}
|	IDENT '(' exprlist ')'
	{
		$$ = CallExpr{parser.Writer, $1, $3}
	}

exprlist:
	{
		$$ = nil
	}
|	expr
	{
		$$ = []Expr{$1}
	}
|	exprlist ',' expr
	{
		$$ = append($1, $3)
	}

%%

func Compile(src []byte, p *parse.Parser) *Tree {
	parser = p
	defaultAction = BlockStmt{[]Stmt{
		ExprStmt{CallExpr{parser.Writer, "print", []Expr{Col{parser, Lit(0)}}}},
	}}
	yyParse(&yyLex{src: src})
	return ast
}
