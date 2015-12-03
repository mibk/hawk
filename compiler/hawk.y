%{
package compiler

import (
	"bufio"
	"bytes"
	"io"

	"github.com/mibk/hawk/parse"
)

var (
	parser *parse.Parser
	ast    *Program

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
%type <exprlist> exprlist  oexprlist
%type <stmt>     stmt ostmt paction ifstmt else if_or_block forstmt
%type <stmtlist> stmtlist blockstmt pactionlist

%token <num> NUM
%token <sym> IDENT STRING PRINT
%token       BEGIN END
%token       IF ELSE
%token       FOR BREAK CONTINUE
%token       INC DEC

%right '?' ':'
%left OROR
%left ANDAND
%left EQ NE LE GE '<' '>'
%left '+' '-'
%left '*' '/' '%'

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
	// TODO: should rather be 'uexpr INC' and catch unwanted usage during
	//	semantic analysis, but for now...
|	IDENT INC
	{
		$$ = AssignStmt{ast, $1, BinaryExpr{Add, Ident{ast, $1}, Lit(1)}}
	}
|	IDENT DEC
	{
		$$ = AssignStmt{ast, $1, BinaryExpr{Sub, Ident{ast, $1}, Lit(1)}}
	}
|	ifstmt
	{
		$$ = $1
	}
|	forstmt
	{
		$$ = $1
	}
|	BREAK
	{
		$$ = StatusStmt{StatusBreak}
	}
|	CONTINUE
	{
		$$ = StatusStmt{StatusContinue}
	}
|	PRINT exprlist
	{
		$$ = CallStmt{parser.Writer, $1, $2}
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
|	expr '?' expr ':' expr
	{
		$$ = TernaryExpr{$1, $3, $5}
	}
|	'$' uexpr
	{
		$$ = FieldExpr{parser, $2}
	}
|	expr OROR expr
	{
		$$ = BinaryExpr{OrOr, $1, $3}
	}
|	expr ANDAND expr
	{
		$$ = BinaryExpr{AndAnd, $1, $3}
	}
|	expr EQ expr
	{
		$$ = BinaryExpr{Eq, $1, $3}
	}
|	expr NE expr
	{
		$$ = BinaryExpr{NotEq, $1, $3}
	}
|	expr LE expr
	{
		$$ = BinaryExpr{LtEq, $1, $3}
	}
|	expr GE expr
	{
		$$ = BinaryExpr{GtEq, $1, $3}
	}
|	expr '<' expr
	{
		$$ = BinaryExpr{Lt, $1, $3}
	}
|	expr '>' expr
	{
		$$ = BinaryExpr{Gt, $1, $3}
	}
|	expr '+' expr
	{
		$$ = BinaryExpr{Add, $1, $3}
	}
|	expr '-' expr
	{
		$$ = BinaryExpr{Sub, $1, $3}
	}
|	expr '*' expr
	{
		$$ = BinaryExpr{Mul, $1, $3}
	}
|	expr '/' expr
	{
		$$ = BinaryExpr{Div, $1, $3}
	}
|	expr '%' expr
	{
		$$ = BinaryExpr{Mod, $1, $3}
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
		$$ = UnaryExpr{Minus, $2}
	}
|	'!' uexpr
	{
		$$ = UnaryExpr{Not, $2}
	}
|	'(' expr ')'
	{
		$$ = $2
	}
|	IDENT
	{
		$$ = Ident{ast, $1}
	}
|	IDENT '(' oexprlist ')'
	{
		$$ = CallExpr{parser.Writer, $1, $3}
	}
	IDENT
	{
	}

exprlist:
	expr
	{
		$$ = []Expr{$1}
	}
|	exprlist ',' expr
	{
		$$ = append($1, $3)
	}

oexprlist:
	{
		$$ = nil
	}
|	exprlist
	{
		$$ = $1
	}

%%

func Compile(r io.Reader, p *parse.Parser) (*Program, error) {
	ast = NewProgram(p)
	parser = p
	defaultAction = BlockStmt{[]Stmt{
		ExprStmt{CallExpr{parser.Writer, "print", []Expr{FieldExpr{parser, Lit(0)}}}},
	}}
	l := &yyLex{reader: bufio.NewReader(r), buf: new(bytes.Buffer)}
	yyParse(l)
	return ast, l.err
}
