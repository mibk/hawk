%{
package compiler

import (
	"github.com/mibk/hawk/parse"
)

var (
	parser *parse.Parser
	ast    *Tree
)

%}

%union {
	num  int
	sym  string
	expr Expr
}

%type	<expr>	expr uexpr

%token	<num>	NUM

%left		EQ NE LE GE LT GT

%start top

%%

top:
	expr
	{
		ast = &Tree{$1}
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

%%

func Compile(src []byte, p *parse.Parser) *Tree {
	parser = p
	yyParse(&yyLex{src: src})
	return ast
}
