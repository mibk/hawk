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

%type <expr>     expr oexpr uexpr
%type <exprlist> exprlist
%type <stmt>     stmt ostmt paction ifstmt else if_or_block forstmt
%type <stmtlist> stmtlist blockstmt pactionlist

%token <num> NUM
%token <sym> IDENT STRING
%token       BEGIN END
%token       IF ELSE
%token       FOR

%left EQ NE LE GE LT GT
%left '+' '-'
%left '*' '/'

%start top

%%

top:
	pactionlist
	{
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
	oexpr blockstmt
	{
		$$ = PatternAction{$1, BlockStmt{$2}}
	}
|	expr
	{
		$$ = PatternAction{$1, defaultAction}
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
|	IDENT '=' expr
	{
		$$ = AssignStmt{ast, $1, $3}
	}
|	ifstmt
	{
		$$ = $1
	}
|	forstmt
	{
		$$ = $1
	}

ostmt:
	{
		$$ = nil
	}
|	stmt
	{
		$$ = $1
	}

ifstmt:
	IF expr blockstmt else
	{
		$$ = IfStmt{$2, BlockStmt{$3}, $4}
	}

else:
	{
		$$ = nil
	}
|	ELSE if_or_block
	{
		$$ = $2
	}

if_or_block:
	ifstmt
	{
		$$ = $1
	}
|	blockstmt
	{
		$$ = BlockStmt{$1}
	}

forstmt:
	FOR ostmt ';' oexpr ';' ostmt blockstmt
	{
		$$ = ForStmt{$2, $4, $6, BlockStmt{$7}}
	}
|	FOR oexpr blockstmt
	{
		$$ = ForStmt{nil, $2, nil, BlockStmt{$3}}
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
		$$ = BinaryExpr{EQ, $1, $3}
	}
|	expr NE expr
	{
		$$ = BinaryExpr{NE, $1, $3}
	}
|	expr LE expr
	{
		$$ = BinaryExpr{LE, $1, $3}
	}
|	expr GE expr
	{
		$$ = BinaryExpr{GE, $1, $3}
	}
|	expr LT expr
	{
		$$ = BinaryExpr{LT, $1, $3}
	}
|	expr GT expr
	{
		$$ = BinaryExpr{GT, $1, $3}
	}
|	expr '+' expr
	{
		$$ = BinaryExpr{ADD, $1, $3}
	}
|	expr '-' expr
	{
		$$ = BinaryExpr{SUB, $1, $3}
	}
|	expr '*' expr
	{
		$$ = BinaryExpr{MUL, $1, $3}
	}
|	expr '/' expr
	{
		$$ = BinaryExpr{DIV, $1, $3}
	}

oexpr:
	{
		$$ = nil
	}
|	expr
	{
		$$ = $1
	}


uexpr:
	NUM
	{
		$$ = Lit($1)
	}
|	STRING
	{
		$$ = StringLit($1)
	}
|	'+' uexpr
	{
		$$ = $2
	}
|	'-' uexpr
	{
		$$ = UnaryExpr{SUB, $2}
	}
|	'(' expr ')'
	{
		$$ = $2
	}
|	IDENT
	{
		$$ = Ident{ast, $1}
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
	ast = NewTree()
	parser = p
	defaultAction = BlockStmt{[]Stmt{
		ExprStmt{CallExpr{parser.Writer, "print", []Expr{Col{parser, Lit(0)}}}},
	}}
	yyParse(&yyLex{src: src})
	return ast
}
