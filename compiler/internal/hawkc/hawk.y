%{
package hawkc

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mibk/hawk/scan"
)

var (
	scanner *scan.Scanner
	ast    *Program

	defaultAction *BlockStmt
)
%}

%union {
	num       int
	sym       string
	symlist   []string
	decl      Decl
	decllist  []Decl
	expr      Expr
	exprlist  []Expr
	stmt      Stmt
	stmtlist  []Stmt
	blockstmt *BlockStmt
}

%type <decl>      decl paction funcdecl
%type <symlist>   arglist
%type <decllist>  decllist
%type <expr>      expr oexpr uexpr
%type <exprlist>  exprlist
%type <stmt>      pipeline stmt ostmt ifstmt else if_or_block forstmt foreachstmt
%type <blockstmt> blockstmt
%type <stmtlist>  stmtlist

%token <num>  NUM BOOL
%token <sym>  IDENT STRING PRINT
%token        BEGIN END
%token        IF ELSE
%token        FOR IN BREAK CONTINUE
%token        INC DEC
%token        ADDEQ SUBEQ MULEQ DIVEQ MODEQ
%token        FUNC RETURN

%right '?' ':'
%left OROR
%left ANDAND
%left EQ NE LE GE '<' '>'
%left '.'
%left '+' '-'
%left '*' '/' '%'

%%

top:
	decllist ';'
	{
		for _, d := range $1 {
			switch d := d.(type) {
			case *BeginAction:
				ast.begin = append(ast.begin, d)
			case *EndAction:
				ast.end = append(ast.end, d)
			case *PatternAction:
				ast.pActions = append(ast.pActions, d)
			case *FuncDecl:
				ast.funcs[d.name] = d
			default:
				panic(fmt.Sprintf("unexpected type: %T", d))
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
		$$ = &PatternAction{$1, $2}
	}
|	expr
	{
		$$ = &PatternAction{$1, defaultAction}
	}
|	BEGIN blockstmt
	{
		$$ = &BeginAction{$2}
	}
|	END blockstmt
	{
		$$ = &EndAction{$2}
	}

funcdecl:
	FUNC IDENT '(' arglist ')' blockstmt
	{
		$$ = &FuncDecl{&FuncScope{}, $2, $4, $6}
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
		$$ = &BlockStmt{$2}
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
		$$ = $1
	}
|	pipeline '|' STRING
	{
		$$ = &PipeStmt{genDebugInfo(), $1, $3}
	}

stmt:
	expr
	{
		$$ = &ExprStmt{$1}
	}
|	IDENT '=' expr
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &Ident{ast, $1}, $3}
	}
|	IDENT '[' ']' '=' expr
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &IndexExpr{genDebugInfo(), &Ident{ast, $1}, nil}, $5}
	}
|	IDENT '[' expr ']' '=' expr
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &IndexExpr{genDebugInfo(), &Ident{ast, $1}, $3}, $6}
	}
|	IDENT ADDEQ expr
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &Ident{ast, $1}, &BinaryExpr{genDebugInfo(), Add, &Ident{ast, $1}, $3}}
	}
|	IDENT SUBEQ expr
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &Ident{ast, $1}, &BinaryExpr{genDebugInfo(), Sub, &Ident{ast, $1}, $3}}
	}
|	IDENT MULEQ expr
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &Ident{ast, $1}, &BinaryExpr{genDebugInfo(), Mul, &Ident{ast, $1}, $3}}
	}
|	IDENT DIVEQ expr
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &Ident{ast, $1}, &BinaryExpr{genDebugInfo(), Div, &Ident{ast, $1}, $3}}
	}
|	IDENT MODEQ expr
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &Ident{ast, $1}, &BinaryExpr{genDebugInfo(), Mod, &Ident{ast, $1}, $3}}
	}
	// TODO: should rather be 'uexpr INC' and catch unwanted usage during
	//	semantic analysis, but for now...
|	IDENT INC
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &Ident{ast, $1}, &BinaryExpr{genDebugInfo(), Add, &Ident{ast, $1}, Lit(1)}}
	}
|	IDENT DEC
	{
		$$ = &AssignStmt{genDebugInfo(), ast, &Ident{ast, $1}, &BinaryExpr{genDebugInfo(), Sub, &Ident{ast, $1}, Lit(1)}}
	}
|	ifstmt
	{
		$$ = $1
	}
|	forstmt
	{
		$$ = $1
	}
|	foreachstmt
	{
		$$ = $1
	}
|	BREAK
	{
		$$ = &StatusStmt{StatusBreak}
	}
|	CONTINUE
	{
		$$ = &StatusStmt{StatusContinue}
	}
|	RETURN oexpr
	{
		$$ = &ReturnStmt{ast, $2}
	}
|	PRINT exprlist
	{
		$$ = &PrintStmt{genDebugInfo(), $1, $2}
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
		$$ = &IfStmt{genDebugInfo(), $2, $3, $4}
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
		$$ = $1
	}

forstmt:
	FOR ostmt ';' oexpr ';' ostmt blockstmt
	{
		$$ = &ForStmt{genDebugInfo(), $2, $4, $6, $7}
	}
|	FOR oexpr blockstmt
	{
		$$ = &ForStmt{genDebugInfo(), nil, $2, nil, $3}
	}

foreachstmt:
	FOR IDENT IN expr blockstmt
	{
		$$ = &ForeachStmt{genDebugInfo(), &Ident{ast, $2}, nil, $4, $5}
	}
|	FOR IDENT ',' IDENT IN expr blockstmt
	{
		$$ = &ForeachStmt{genDebugInfo(), &Ident{ast, $2}, &Ident{ast, $4}, $6, $7}
	}


expr:
	uexpr
	{
		$$ = $1
	}
|	expr '?' expr ':' expr
	{
		$$ = &TernaryExpr{genDebugInfo(), $1, $3, $5}
	}
|	'$' uexpr
	{
		$$ = &FieldExpr{genDebugInfo(), scanner, $2}
	}
|	expr OROR expr
	{
		$$ = &BinaryExpr{genDebugInfo(), OrOr, $1, $3}
	}
|	expr ANDAND expr
	{
		$$ = &BinaryExpr{genDebugInfo(), AndAnd, $1, $3}
	}
|	expr EQ expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Eq, $1, $3}
	}
|	expr NE expr
	{
		$$ = &BinaryExpr{genDebugInfo(), NotEq, $1, $3}
	}
|	expr LE expr
	{
		$$ = &BinaryExpr{genDebugInfo(), LtEq, $1, $3}
	}
|	expr GE expr
	{
		$$ = &BinaryExpr{genDebugInfo(), GtEq, $1, $3}
	}
|	expr '<' expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Lt, $1, $3}
	}
|	expr '>' expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Gt, $1, $3}
	}
|	expr '+' expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Add, $1, $3}
	}
|	expr '-' expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Sub, $1, $3}
	}
|	expr '*' expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Mul, $1, $3}
	}
|	expr '/' expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Div, $1, $3}
	}
|	expr '%' expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Mod, $1, $3}
	}
|	expr '.' expr
	{
		$$ = &BinaryExpr{genDebugInfo(), Concat, $1, $3}
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
|	BOOL
	{
		$$ = BoolLit($1 != 0)
	}
|	'+' uexpr
	{
		$$ = $2
	}
|	'-' uexpr
	{
		$$ = &UnaryExpr{genDebugInfo(), Minus, $2}
	}
|	'!' uexpr
	{
		$$ = &UnaryExpr{genDebugInfo(), Not, $2}
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
		$$ = &CallExpr{genDebugInfo(), $1, nil}
	}
|	IDENT '(' exprlist ocomma ')'
	{
		$$ = &CallExpr{genDebugInfo(), $1, $3}
	}
|	'[' ']'
	{
		$$ = &ArrayLit{}
	}
|	'[' exprlist ocomma ']'
	{
		$$ = &ArrayLit{$2}
	}
	// TODO: Allow more expr than just IDENT.
|	IDENT '[' expr ']'
	{
		$$ = &IndexExpr{genDebugInfo(), &Ident{ast, $1}, $3}
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

// Compile compiles a Hawk program from src. It is not safe
// for concurrent use.
func Compile(src io.Reader) (*Program, error) {
	scanner = new(scan.Scanner)
	defaultAction = &BlockStmt{[]Stmt{
		&PrintStmt{fun: "print", args: []Expr{&FieldExpr{sc: scanner, num: Lit(0)}}},
	}}
	ast = NewProgram(scanner)
	lexlineno = 1
	nlsemi = false
	l := &yyLex{reader: bufio.NewReader(src)}
	yyParse(l)
	analyse(ast)
	return ast, l.err
}
