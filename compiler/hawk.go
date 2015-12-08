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
	"'|'",
	"'='",
	"'$'",
	"'!'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line hawk.y:408

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

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	43, 62,
	-2, 7,
	-1, 20,
	43, 61,
	-2, 1,
	-1, 60,
	43, 61,
	-2, 36,
	-1, 98,
	43, 62,
	-2, 21,
}

const yyNprod = 78
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 304

var yyAct = [...]int{

	51, 7, 50, 53, 49, 95, 6, 78, 48, 21,
	82, 37, 38, 104, 22, 118, 119, 113, 103, 44,
	120, 7, 59, 75, 61, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 72, 73, 74, 13, 19,
	14, 45, 8, 9, 130, 115, 79, 89, 90, 84,
	85, 86, 87, 88, 81, 22, 10, 20, 92, 79,
	94, 98, 3, 123, 91, 97, 93, 96, 15, 16,
	106, 126, 45, 101, 18, 34, 35, 36, 83, 39,
	102, 12, 17, 46, 107, 108, 109, 110, 111, 112,
	105, 80, 1, 47, 13, 19, 14, 114, 54, 116,
	117, 127, 122, 2, 121, 100, 5, 4, 0, 13,
	52, 14, 58, 0, 121, 59, 92, 60, 55, 56,
	0, 125, 124, 0, 15, 16, 129, 128, 57, 0,
	18, 77, 0, 0, 132, 97, 131, 12, 17, 15,
	16, 13, 52, 14, 58, 18, 11, 59, 22, 60,
	55, 56, 12, 17, 32, 33, 34, 35, 36, 40,
	57, 0, 41, 42, 43, 0, 0, 0, 0, 0,
	0, 15, 16, 13, 19, 14, 0, 18, 0, 0,
	0, 0, 0, 0, 12, 17, 23, 0, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 0, 0, 15, 16, 22, 13, 19, 14, 18,
	0, 0, 0, 0, 0, 0, 12, 17, 23, 0,
	24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 0, 0, 76, 15, 16, 0, 0,
	0, 0, 18, 0, 0, 0, 0, 0, 0, 0,
	17, 23, 99, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 23, 0, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 26, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36,
}
var yyPact = [...]int{

	34, -1000, 18, -1000, -1000, -1000, -29, 242, -29, -29,
	74, -1000, 202, -1000, -1000, 202, 202, 202, 169, 1,
	34, -1000, 105, 169, 169, 169, 169, 169, 169, 169,
	169, 169, 169, 169, 169, 169, 169, -1000, -1000, -17,
	-1000, -1000, -1000, -1000, 194, 90, -1000, 15, -35, -1000,
	-1000, 242, 32, -1000, -1000, -1000, -1000, 169, 169, 169,
	137, 227, 254, 265, 120, 120, 120, 120, 120, 120,
	39, 39, -1000, -1000, -1000, 68, -1000, -1000, -24, 242,
	-31, 105, 64, 169, 169, 169, 169, 169, 169, -1000,
	-1000, -1000, 242, -25, 162, 6, -29, -1000, 242, 169,
	-26, -1000, -21, 169, -1000, -35, -1000, 242, 242, 242,
	242, 242, 242, 169, 52, 169, -1000, 242, -29, 66,
	-1000, 242, -1000, 12, 5, -1000, -1000, -1000, -1000, -1000,
	137, -29, -1000,
}
var yyPgo = [...]int{

	0, 62, 107, 106, 105, 103, 0, 6, 146, 7,
	8, 4, 5, 3, 102, 101, 98, 93, 2, 92,
	91, 80,
}
var yyR1 = [...]int{

	0, 19, 5, 5, 1, 1, 2, 2, 2, 2,
	3, 4, 4, 4, 18, 17, 17, 17, 10, 10,
	10, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 12, 12, 13, 14,
	14, 15, 15, 16, 16, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 7, 7, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 9, 9, 20, 20, 21, 21,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 1, 2, 1, 2, 2,
	6, 0, 1, 3, 4, 0, 1, 3, 1, 1,
	3, 1, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 1, 1, 2, 2, 0, 1, 4, 0,
	2, 1, 1, 7, 3, 1, 5, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 0, 1, 1, 1, 2, 2, 2, 3, 1,
	3, 5, 1, 3, 0, 1, 0, 1,
}
var yyChk = [...]int{

	-1000, -19, -5, -1, -2, -3, -7, -6, 8, 9,
	22, -8, 47, 4, 6, 34, 35, 48, 40, 5,
	39, -18, 43, 24, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, -18, -18, 5,
	-8, -8, -8, -8, -6, 40, -1, -17, -10, -11,
	-18, -6, 5, -13, -16, 13, 14, 23, 7, 10,
	12, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, 40, 41, 41, -9, -6,
	-20, 39, 45, 46, 17, 18, 19, 20, 21, 15,
	16, -7, -6, -9, -6, -12, -7, -11, -6, 25,
	-4, 5, -21, 42, 44, -10, 6, -6, -6, -6,
	-6, -6, -6, 42, -18, 39, -18, -6, 41, 42,
	41, -6, -14, 11, -7, -18, 5, -15, -13, -18,
	39, -12, -18,
}
var yyDef = [...]int{

	61, -2, 0, 2, 4, 5, 0, -2, 0, 0,
	0, 45, 0, 63, 64, 0, 0, 0, 0, 69,
	-2, 6, 15, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 8, 9, 0,
	47, 65, 66, 67, 0, 0, 3, 74, 16, 18,
	19, 21, 69, 30, 31, 32, 33, 61, 0, 0,
	-2, 0, 48, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 11, 68, 70, 76, 72,
	0, 75, 0, 0, 0, 0, 0, 0, 0, 28,
	29, 34, 62, 35, 0, 0, 0, 37, -2, 0,
	0, 12, 0, 77, 14, 17, 20, 22, 23, 24,
	25, 26, 27, 0, 39, 61, 44, 46, 0, 0,
	71, 73, 38, 0, 0, 10, 13, 40, 41, 42,
	36, 0, 43,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 48, 3, 3, 47, 38, 3, 3,
	40, 41, 36, 34, 42, 35, 3, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 25, 39,
	32, 46, 33, 24, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 43, 45, 44,
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
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:160
		{
			yyVAL.stmt = BlockStmt{yyDollar[1].stmtlist}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:164
		{
			yyVAL.stmt = PipeStmt{yyDollar[1].stmt, yyDollar[3].sym}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:170
		{
			yyVAL.stmt = ExprStmt{yyDollar[1].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:174
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:178
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:182
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:186
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Mul, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:190
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Div, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:194
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Mod, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:200
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:204
		{
			yyVAL.stmt = &AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:208
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:212
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:216
		{
			yyVAL.stmt = StatusStmt{StatusBreak}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:220
		{
			yyVAL.stmt = StatusStmt{StatusContinue}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:224
		{
			yyVAL.stmt = ReturnStmt{ast, yyDollar[2].expr}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:228
		{
			yyVAL.stmt = CallStmt{yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:233
		{
			yyVAL.stmt = nil
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:237
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:243
		{
			yyVAL.stmt = IfStmt{yyDollar[2].expr, BlockStmt{yyDollar[3].stmtlist}, yyDollar[4].stmt}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:248
		{
			yyVAL.stmt = nil
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:252
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:258
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:262
		{
			yyVAL.stmt = BlockStmt{yyDollar[1].stmtlist}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:268
		{
			yyVAL.stmt = ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, BlockStmt{yyDollar[7].stmtlist}}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:272
		{
			yyVAL.stmt = ForStmt{nil, yyDollar[2].expr, nil, BlockStmt{yyDollar[3].stmtlist}}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:279
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:283
		{
			yyVAL.expr = TernaryExpr{yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:287
		{
			yyVAL.expr = FieldExpr{parser, yyDollar[2].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:291
		{
			yyVAL.expr = BinaryExpr{OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:295
		{
			yyVAL.expr = BinaryExpr{AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:299
		{
			yyVAL.expr = BinaryExpr{Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:303
		{
			yyVAL.expr = BinaryExpr{NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:307
		{
			yyVAL.expr = BinaryExpr{LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:311
		{
			yyVAL.expr = BinaryExpr{GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:315
		{
			yyVAL.expr = BinaryExpr{Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:319
		{
			yyVAL.expr = BinaryExpr{Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:323
		{
			yyVAL.expr = BinaryExpr{Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:327
		{
			yyVAL.expr = BinaryExpr{Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:331
		{
			yyVAL.expr = BinaryExpr{Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:335
		{
			yyVAL.expr = BinaryExpr{Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:339
		{
			yyVAL.expr = BinaryExpr{Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:344
		{
			yyVAL.expr = nil
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:348
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:355
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:359
		{
			yyVAL.expr = StringLit(yyDollar[1].sym)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:363
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:367
		{
			yyVAL.expr = UnaryExpr{Minus, yyDollar[2].expr}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:371
		{
			yyVAL.expr = UnaryExpr{Not, yyDollar[2].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:375
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:379
		{
			yyVAL.expr = &Ident{ast, yyDollar[1].sym}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:383
		{
			yyVAL.expr = CallExpr{yyDollar[1].sym, nil}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:387
		{
			yyVAL.expr = CallExpr{yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:393
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:397
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
