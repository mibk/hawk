//line hawk.y:2
package compiler

import __yyfmt__ "fmt"

//line hawk.y:2
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

//line hawk.y:21
type yySymType struct {
	yys      int
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

const NUM = 57346
const IDENT = 57347
const STRING = 57348
const PRINT = 57349
const BEGIN = 57350
const END = 57351
const IF = 57352
const ELSE = 57353
const FOR = 57354
const BREAK = 57355
const CONTINUE = 57356
const INC = 57357
const DEC = 57358
const ADDEQ = 57359
const SUBEQ = 57360
const MULEQ = 57361
const DIVEQ = 57362
const MODEQ = 57363
const FUNC = 57364
const RETURN = 57365
const OROR = 57366
const ANDAND = 57367
const EQ = 57368
const NE = 57369
const LE = 57370
const GE = 57371

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUM",
	"IDENT",
	"STRING",
	"PRINT",
	"BEGIN",
	"END",
	"IF",
	"ELSE",
	"FOR",
	"BREAK",
	"CONTINUE",
	"INC",
	"DEC",
	"ADDEQ",
	"SUBEQ",
	"MULEQ",
	"DIVEQ",
	"MODEQ",
	"FUNC",
	"RETURN",
	"'?'",
	"':'",
	"OROR",
	"ANDAND",
	"EQ",
	"NE",
	"LE",
	"GE",
	"'<'",
	"'>'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"';'",
	"'('",
	"')'",
	"','",
	"'{'",
	"'}'",
	"'='",
	"'$'",
	"'!'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line hawk.y:394

func Compile(r io.Reader, p *parse.Parser) (*Program, error) {
	ast = NewProgram(p)
	parser = p
	defaultAction = BlockStmt{[]Stmt{
		ExprStmt{CallExpr{parser.Writer, "print", []Expr{FieldExpr{parser, Lit(0)}}}},
	}}
	l := &yyLex{reader: bufio.NewReader(r), buf: new(bytes.Buffer)}
	yyParse(l)
	analyse(ast)
	return ast, l.err
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	43, 59,
	-2, 7,
	-1, 20,
	43, 58,
	-2, 1,
	-1, 58,
	43, 58,
	-2, 33,
	-1, 95,
	43, 59,
	-2, 18,
}

const yyNprod = 75
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 280

var yyAct = [...]int{

	49, 7, 21, 51, 92, 101, 6, 76, 57, 22,
	109, 37, 38, 114, 115, 94, 100, 116, 73, 44,
	45, 7, 126, 111, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 72, 48, 79,
	20, 22, 34, 35, 36, 3, 77, 86, 87, 81,
	82, 83, 84, 85, 119, 122, 89, 77, 91, 95,
	98, 39, 88, 99, 90, 93, 46, 32, 33, 34,
	35, 36, 45, 13, 19, 14, 78, 80, 1, 47,
	52, 103, 104, 105, 106, 107, 108, 123, 118, 2,
	13, 19, 14, 97, 110, 102, 112, 113, 5, 4,
	0, 117, 0, 15, 16, 13, 50, 14, 56, 18,
	117, 57, 89, 58, 53, 54, 17, 121, 120, 0,
	15, 16, 125, 124, 55, 0, 18, 75, 0, 0,
	128, 127, 12, 17, 0, 15, 16, 0, 0, 0,
	0, 18, 13, 19, 14, 0, 0, 12, 17, 23,
	0, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 11, 13, 19, 14, 22, 8,
	9, 0, 15, 16, 0, 0, 0, 40, 18, 0,
	41, 42, 43, 10, 12, 17, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 15, 16, 0, 0, 0,
	0, 18, 0, 0, 0, 0, 0, 12, 17, 23,
	0, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 0, 0, 74, 23, 96, 24,
	25, 26, 27, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 23, 0, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 25, 26, 27,
	28, 29, 30, 31, 32, 33, 34, 35, 36, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
}
var yyPact = [...]int{

	161, -1000, 1, -1000, -1000, -1000, -34, 218, -34, -34,
	56, -1000, 69, -1000, -1000, 69, 69, 69, 138, -20,
	161, -1000, 101, 138, 138, 138, 138, 138, 138, 138,
	138, 138, 138, 138, 138, 138, 138, -1000, -1000, -22,
	-1000, -1000, -1000, -1000, 185, 86, -1000, 0, -1000, 218,
	32, -1000, -1000, -1000, -1000, 138, 138, 138, 101, 203,
	230, 241, 33, 33, 33, 33, 33, 33, 6, 6,
	-1000, -1000, -1000, 55, -1000, -1000, -26, 218, -39, 101,
	138, 138, 138, 138, 138, 138, -1000, -1000, -1000, 218,
	-32, 125, -16, -34, -1000, 218, 138, -28, -1000, -24,
	138, -1000, -1000, 218, 218, 218, 218, 218, 218, 138,
	43, 138, -1000, 218, -34, 50, -1000, 218, -1000, -2,
	-17, -1000, -1000, -1000, -1000, -1000, 101, -34, -1000,
}
var yyPgo = [...]int{

	0, 45, 99, 98, 93, 89, 0, 6, 164, 7,
	15, 4, 3, 88, 87, 80, 79, 2, 78, 76,
	63,
}
var yyR1 = [...]int{

	0, 18, 5, 5, 1, 1, 2, 2, 2, 2,
	3, 4, 4, 4, 17, 16, 16, 16, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 11, 11, 12, 13, 13, 14, 14,
	15, 15, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 7, 7,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 9,
	9, 19, 19, 20, 20,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 1, 2, 1, 2, 2,
	6, 0, 1, 3, 4, 0, 1, 3, 1, 3,
	3, 3, 3, 3, 3, 2, 2, 1, 1, 1,
	1, 2, 2, 0, 1, 4, 0, 2, 1, 1,
	7, 3, 1, 5, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 0, 1,
	1, 1, 2, 2, 2, 3, 1, 3, 5, 1,
	3, 0, 1, 0, 1,
}
var yyChk = [...]int{

	-1000, -18, -5, -1, -2, -3, -7, -6, 8, 9,
	22, -8, 46, 4, 6, 34, 35, 47, 40, 5,
	39, -17, 43, 24, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, -17, -17, 5,
	-8, -8, -8, -8, -6, 40, -1, -16, -10, -6,
	5, -12, -15, 13, 14, 23, 7, 10, 12, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, 40, 41, 41, -9, -6, -19, 39,
	45, 17, 18, 19, 20, 21, 15, 16, -7, -6,
	-9, -6, -11, -7, -10, -6, 25, -4, 5, -20,
	42, 44, -10, -6, -6, -6, -6, -6, -6, 42,
	-17, 39, -17, -6, 41, 42, 41, -6, -13, 11,
	-7, -17, 5, -14, -12, -17, 39, -11, -17,
}
var yyDef = [...]int{

	58, -2, 0, 2, 4, 5, 0, -2, 0, 0,
	0, 42, 0, 60, 61, 0, 0, 0, 0, 66,
	-2, 6, 15, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 8, 9, 0,
	44, 62, 63, 64, 0, 0, 3, 71, 16, 18,
	66, 27, 28, 29, 30, 58, 0, 0, -2, 0,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 11, 65, 67, 73, 69, 0, 72,
	0, 0, 0, 0, 0, 0, 25, 26, 31, 59,
	32, 0, 0, 0, 34, -2, 0, 0, 12, 0,
	74, 14, 17, 19, 20, 21, 22, 23, 24, 0,
	36, 58, 41, 43, 0, 0, 68, 70, 35, 0,
	0, 10, 13, 37, 38, 39, 33, 0, 40,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 47, 3, 3, 46, 38, 3, 3,
	40, 41, 36, 34, 42, 35, 3, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 25, 39,
	32, 45, 33, 24, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 43, 3, 44,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 26, 27, 28, 29, 30, 31,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:61
		{
			for _, d := range yyDollar[1].decllist {
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
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:80
		{
			yyVAL.decllist = []Decl{yyDollar[1].decl}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:84
		{
			yyVAL.decllist = append(yyDollar[1].decllist, yyDollar[3].decl)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:90
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:94
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:100
		{
			yyVAL.decl = PatternAction{yyDollar[1].expr, BlockStmt{yyDollar[2].stmtlist}}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:104
		{
			yyVAL.decl = PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:108
		{
			yyVAL.decl = BeginAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:112
		{
			yyVAL.decl = EndAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:118
		{
			yyVAL.decl = FuncDecl{new(FuncScope), yyDollar[2].sym, yyDollar[4].symlist, BlockStmt{yyDollar[6].stmtlist}}
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:123
		{
			yyVAL.symlist = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:127
		{
			yyVAL.symlist = []string{yyDollar[1].sym}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:131
		{
			yyVAL.symlist = append(yyDollar[1].symlist, yyDollar[3].sym)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:137
		{
			yyVAL.stmtlist = yyDollar[2].stmtlist
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:142
		{
			yyVAL.stmtlist = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:146
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:150
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:156
		{
			yyVAL.stmt = ExprStmt{yyDollar[1].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:160
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:164
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:168
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:172
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Mul, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:176
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Div, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:180
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Mod, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:186
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:190
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:194
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:198
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:202
		{
			yyVAL.stmt = StatusStmt{StatusBreak}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:206
		{
			yyVAL.stmt = StatusStmt{StatusContinue}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:210
		{
			yyVAL.stmt = ReturnStmt{ast, yyDollar[2].expr}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:214
		{
			yyVAL.stmt = CallStmt{parser.Writer, yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:219
		{
			yyVAL.stmt = nil
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:223
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:229
		{
			yyVAL.stmt = IfStmt{yyDollar[2].expr, BlockStmt{yyDollar[3].stmtlist}, yyDollar[4].stmt}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:234
		{
			yyVAL.stmt = nil
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:238
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:244
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:248
		{
			yyVAL.stmt = BlockStmt{yyDollar[1].stmtlist}
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:254
		{
			yyVAL.stmt = ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, BlockStmt{yyDollar[7].stmtlist}}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:258
		{
			yyVAL.stmt = ForStmt{nil, yyDollar[2].expr, nil, BlockStmt{yyDollar[3].stmtlist}}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:265
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:269
		{
			yyVAL.expr = TernaryExpr{yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:273
		{
			yyVAL.expr = FieldExpr{parser, yyDollar[2].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:277
		{
			yyVAL.expr = BinaryExpr{OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:281
		{
			yyVAL.expr = BinaryExpr{AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:285
		{
			yyVAL.expr = BinaryExpr{Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:289
		{
			yyVAL.expr = BinaryExpr{NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:293
		{
			yyVAL.expr = BinaryExpr{LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:297
		{
			yyVAL.expr = BinaryExpr{GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:301
		{
			yyVAL.expr = BinaryExpr{Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:305
		{
			yyVAL.expr = BinaryExpr{Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:309
		{
			yyVAL.expr = BinaryExpr{Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:313
		{
			yyVAL.expr = BinaryExpr{Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:317
		{
			yyVAL.expr = BinaryExpr{Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:321
		{
			yyVAL.expr = BinaryExpr{Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:325
		{
			yyVAL.expr = BinaryExpr{Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:330
		{
			yyVAL.expr = nil
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:334
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:341
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:345
		{
			yyVAL.expr = StringLit(yyDollar[1].sym)
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:349
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:353
		{
			yyVAL.expr = UnaryExpr{Minus, yyDollar[2].expr}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:357
		{
			yyVAL.expr = UnaryExpr{Not, yyDollar[2].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:361
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:365
		{
			yyVAL.expr = &Ident{ast, yyDollar[1].sym}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:369
		{
			yyVAL.expr = CallExpr{parser.Writer, yyDollar[1].sym, nil}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:373
		{
			yyVAL.expr = CallExpr{parser.Writer, yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:379
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:383
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
