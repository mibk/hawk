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
	symlist  []string
	decl     Decl
	decllist []Decl
	expr     Expr
	exprlist []Expr
	stmt     Stmt
	stmtlist []Stmt
}

%type <decl>     decl paction funcdecl
%type <symlist>  arglist
%type <decllist> decllist
%type <expr>     expr oexpr uexpr
%type <exprlist> exprlist
%type <stmt>     pipeline stmt ostmt ifstmt else if_or_block forstmt
%type <stmtlist> stmtlist blockstmt

%token <num> NUM
%token <sym> IDENT STRING PRINT
%token       BEGIN END
%token       IF ELSE
%token       FOR BREAK CONTINUE
%token       INC DEC
%token       ADDEQ SUBEQ MULEQ DIVEQ MODEQ
%token       FUNC RETURN

%right '?' ':'
%left OROR
%left ANDAND
%left EQ NE LE GE '<' '>'
%left '+' '-'
%left '*' '/' '%'

%%

top:
	decllist ';'
	{
		for _, d := range $1 {
			switch d := d.(type) {
			case BeginAction:
				ast.begin = append(ast.begin, d)
			case EndAction:
				ast.end = append(ast.end, d)
			case PatternAction:
				ast.pActions = append(ast.pActions, d)
			case FuncDecl:
				ast.funcs[d.name] = d
			default:
				panic("unreachable")
			}
		}
	}

decllist:
	decl
	{
		$$ = []Decl{$1}
	}
|	decllist ';' decl
	{
		$$ = append($1, $3)
	}

decl:
	paction
	{
		$$ = $1
	}
|	funcdecl
	{
		$$ = $1
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

funcdecl:
	FUNC IDENT '(' arglist ')' blockstmt
	{
		$$ = FuncDecl{new(FuncScope), $2, $4, BlockStmt{$6}}
	}

arglist:
	{
		$$ = nil
	}
|	IDENT
	{
		$$ = []string{$1}
	}
|	arglist ',' IDENT
	{
		$$ = append($1, $3)
	}

blockstmt:
	'{' stmtlist osemi '}'
	{
		$$ = $2
	}

stmtlist:
	{
		$$ = nil
	}
|	pipeline
	{
		$$ = []Stmt{$1}
	}
|	stmtlist ';' pipeline
	{
		$$ = append($1, $3)
	}

pipeline:
	stmt
	{
		$$ = $1
	}
|	blockstmt
	{
		$$ = BlockStmt{$1}
	}
|	pipeline '|' STRING
	{
		$$ = PipeStmt{$1, $3}
	}

stmt:
	expr
	{
		$$ = ExprStmt{$1}
	}
|	IDENT '=' expr
	{
		$$ = &AssignStmt{ast, $1, $3}
	}
|	IDENT ADDEQ expr
	{
		$$ = &AssignStmt{ast, $1, BinaryExpr{Add, &Ident{ast, $1}, $3}}
	}
|	IDENT SUBEQ expr
	{
		$$ = &AssignStmt{ast, $1, BinaryExpr{Sub, &Ident{ast, $1}, $3}}
	}
|	IDENT MULEQ expr
	{
		$$ = &AssignStmt{ast, $1, BinaryExpr{Mul, &Ident{ast, $1}, $3}}
	}
|	IDENT DIVEQ expr
	{
		$$ = &AssignStmt{ast, $1, BinaryExpr{Div, &Ident{ast, $1}, $3}}
	}
|	IDENT MODEQ expr
	{
		$$ = &AssignStmt{ast, $1, BinaryExpr{Mod, &Ident{ast, $1}, $3}}
	}
	// TODO: should rather be 'uexpr INC' and catch unwanted usage during
	//	semantic analysis, but for now...
|	IDENT INC
	{
		$$ = &AssignStmt{ast, $1, BinaryExpr{Add, &Ident{ast, $1}, Lit(1)}}
	}
|	IDENT DEC
	{
		$$ = &AssignStmt{ast, $1, BinaryExpr{Sub, &Ident{ast, $1}, Lit(1)}}
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
|	RETURN oexpr
	{
		$$ = ReturnStmt{ast, $2}
	}
|	PRINT exprlist
	{
		$$ = CallStmt{$1, $2}
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
		$$ = &Ident{ast, $1}
	}
|	IDENT '(' ')'
	{
		$$ = CallExpr{$1, nil}
	}
|	IDENT '(' exprlist ocomma ')'
	{
		$$ = CallExpr{$1, $3}
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


osemi:
|	';'

ocomma:
|	','

%%

func Compile(r io.Reader, p *parse.Parser) (*Program, error) {
	ast = NewProgram(p)
	parser = p
	defaultAction = BlockStmt{[]Stmt{
		ExprStmt{CallExpr{"print", []Expr{FieldExpr{parser, Lit(0)}}}},
	}}
	l := &yyLex{reader: bufio.NewReader(r), buf: new(bytes.Buffer)}
	yyParse(l)
	analyse(ast)
	return ast, l.err
}
