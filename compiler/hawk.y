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
	paction	 PatternAction
}

%type	<expr>	expr uexpr
%type	<exprlist>	exprlist
%type	<stmt>	stmt
%type	<stmtlist>	stmtlist blockstmt
%type	<paction>	paction

%token	<num>	NUM
%token	<sym>	IDENT

%left		EQ NE LE GE LT GT

%start	top

%%

top:
	paction
	{
		ast = &Tree{$1}
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

blockstmt:
	'{' stmtlist '}'
	{
		$$ = $2
	}

stmtlist:
	{
		$$ = nil
	}
|	stmt
	{
		$$ = append($$, $1)
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
		$$ = append($$, $1)
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
