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
const OROR = 57364
const ANDAND = 57365
const EQ = 57366
const NE = 57367
const LE = 57368
const GE = 57369

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
	"'{'",
	"'}'",
	"'='",
	"'$'",
	"'!'",
	"'('",
	"')'",
	"','",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line hawk.y:359

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

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 5,
	38, 52,
	-2, 5,
	-1, 53,
	38, 51,
	-2, 26,
	-1, 86,
	38, 52,
	-2, 12,
}

const yyNprod = 66
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 277

var yyAct = [...]int{

	45, 5, 18, 47, 83, 4, 89, 88, 41, 34,
	35, 19, 10, 16, 11, 112, 40, 99, 5, 85,
	52, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 72, 17, 70, 105, 44,
	12, 13, 71, 29, 30, 31, 32, 33, 19, 9,
	14, 15, 71, 82, 86, 31, 32, 33, 20, 84,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 3, 92, 93, 94, 95, 96, 97,
	68, 108, 102, 1, 2, 98, 43, 100, 101, 81,
	103, 42, 91, 10, 46, 11, 51, 8, 48, 52,
	107, 53, 49, 50, 109, 106, 104, 36, 111, 110,
	37, 38, 39, 69, 0, 0, 114, 113, 0, 0,
	0, 12, 13, 0, 0, 0, 0, 0, 90, 0,
	9, 14, 15, 10, 46, 11, 51, 0, 0, 52,
	0, 53, 49, 50, 20, 0, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, 32, 33, 0,
	19, 12, 13, 10, 16, 11, 0, 6, 7, 0,
	9, 14, 15, 0, 0, 0, 0, 0, 79, 80,
	74, 75, 76, 77, 78, 0, 0, 0, 0, 0,
	0, 12, 13, 0, 0, 0, 0, 0, 0, 0,
	9, 14, 15, 73, 0, 0, 41, 20, 87, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 20, 0, 21, 22, 23, 24, 25, 26,
	27, 28, 29, 30, 31, 32, 33, 10, 16, 11,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 0, 0, 12, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 15,
}
var yyPact = [...]int{

	159, -1000, -1, -1000, -27, 200, -27, -27, -1000, 233,
	-1000, -1000, 233, 233, 233, 8, -35, 159, -1000, 129,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, -1000, -1000, -1000, -1000, -1000, -1000,
	36, 8, -1000, -2, -1000, 200, 163, -1000, -1000, -1000,
	-1000, 8, 8, 129, 185, 215, 226, 11, 11, 11,
	11, 11, 11, 21, 21, -1000, -1000, -1000, -1000, -37,
	-39, 200, 89, 8, 8, 8, 8, 8, 8, -1000,
	-1000, -39, 122, -20, -27, -1000, 200, 8, -1000, 8,
	-1000, -1000, 200, 200, 200, 200, 200, 200, 27, 8,
	-1000, 200, 76, 200, -1000, 10, -22, 200, -1000, -1000,
	-1000, -1000, 129, -27, -1000,
}
var yyPgo = [...]int{

	0, 0, 5, 97, 37, 113, 19, 4, 73, 3,
	106, 104, 98, 86, 2, 84, 83, 82,
}
var yyR1 = [...]int{

	0, 16, 15, 15, 8, 8, 8, 8, 14, 13,
	13, 13, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 7, 7, 9, 10,
	10, 11, 11, 12, 12, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	17, 3, 4, 4, 5, 5,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 2, 1, 2, 2, 4, 0,
	1, 3, 1, 3, 3, 3, 3, 3, 3, 2,
	2, 1, 1, 1, 1, 2, 0, 1, 4, 0,
	2, 1, 1, 7, 3, 1, 5, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 0, 1, 1, 1, 2, 2, 2, 3, 1,
	0, 6, 1, 3, 0, 1,
}
var yyChk = [...]int{

	-1000, -16, -15, -8, -2, -1, 8, 9, -3, 41,
	4, 6, 32, 33, 42, 43, 5, 37, -14, 38,
	22, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, -14, -14, -3, -3, -3, -3,
	-1, 43, -8, -13, -6, -1, 5, -9, -12, 13,
	14, 7, 10, 12, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, 44, -5,
	-4, -1, 37, 40, 17, 18, 19, 20, 21, 15,
	16, -4, -1, -7, -2, -6, -1, 23, 44, 45,
	39, -6, -1, -1, -1, -1, -1, -1, -14, 37,
	-14, -1, -17, -1, -10, 11, -2, -1, 5, -11,
	-9, -14, 37, -7, -14,
}
var yyDef = [...]int{

	51, -2, 1, 2, 0, -2, 0, 0, 35, 0,
	53, 54, 0, 0, 0, 0, 59, 51, 4, 9,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 6, 7, 37, 55, 56, 57,
	0, 64, 3, 0, 10, 12, 59, 21, 22, 23,
	24, 0, 0, -2, 0, 38, 39, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50, 58, 0,
	65, 62, 0, 0, 0, 0, 0, 0, 0, 19,
	20, 25, 0, 0, 0, 27, -2, 0, 60, 0,
	8, 11, 13, 14, 15, 16, 17, 18, 29, 51,
	34, 36, 0, 63, 28, 0, 0, 52, 61, 30,
	31, 32, 26, 0, 33,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 42, 3, 3, 41, 36, 3, 3,
	43, 44, 34, 32, 45, 33, 3, 35, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 23, 37,
	30, 40, 31, 22, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 38, 3, 39,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	24, 25, 26, 27, 28, 29,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:56
		{
			for i := 0; i < len(yyDollar[1].stmtlist); {
				pa := yyDollar[1].stmtlist[i]
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
				yyDollar[1].stmtlist = append(yyDollar[1].stmtlist[:i], yyDollar[1].stmtlist[i+1:]...)
			}
			ast.pActions = yyDollar[1].stmtlist
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:77
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:81
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:87
		{
			yyVAL.stmt = PatternAction{yyDollar[1].expr, BlockStmt{yyDollar[2].stmtlist}}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:91
		{
			yyVAL.stmt = PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:95
		{
			yyVAL.stmt = BeginAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:99
		{
			yyVAL.stmt = EndAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:105
		{
			yyVAL.stmtlist = yyDollar[2].stmtlist
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:110
		{
			yyVAL.stmtlist = nil
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:114
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:118
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:124
		{
			yyVAL.stmt = ExprStmt{yyDollar[1].expr}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:128
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, yyDollar[3].expr}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:132
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Add, Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:136
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Sub, Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:140
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Mul, Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:144
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Div, Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:148
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Mod, Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:154
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Add, Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:158
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Sub, Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:162
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:166
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:170
		{
			yyVAL.stmt = StatusStmt{StatusBreak}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:174
		{
			yyVAL.stmt = StatusStmt{StatusContinue}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:178
		{
			yyVAL.stmt = CallStmt{parser.Writer, yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:183
		{
			yyVAL.stmt = nil
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:187
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:193
		{
			yyVAL.stmt = IfStmt{yyDollar[2].expr, BlockStmt{yyDollar[3].stmtlist}, yyDollar[4].stmt}
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:198
		{
			yyVAL.stmt = nil
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:202
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:208
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:212
		{
			yyVAL.stmt = BlockStmt{yyDollar[1].stmtlist}
		}
	case 33:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:218
		{
			yyVAL.stmt = ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, BlockStmt{yyDollar[7].stmtlist}}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:222
		{
			yyVAL.stmt = ForStmt{nil, yyDollar[2].expr, nil, BlockStmt{yyDollar[3].stmtlist}}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:229
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:233
		{
			yyVAL.expr = TernaryExpr{yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:237
		{
			yyVAL.expr = FieldExpr{parser, yyDollar[2].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:241
		{
			yyVAL.expr = BinaryExpr{OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:245
		{
			yyVAL.expr = BinaryExpr{AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:249
		{
			yyVAL.expr = BinaryExpr{Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:253
		{
			yyVAL.expr = BinaryExpr{NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:257
		{
			yyVAL.expr = BinaryExpr{LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:261
		{
			yyVAL.expr = BinaryExpr{GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:265
		{
			yyVAL.expr = BinaryExpr{Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:269
		{
			yyVAL.expr = BinaryExpr{Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:273
		{
			yyVAL.expr = BinaryExpr{Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:277
		{
			yyVAL.expr = BinaryExpr{Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:281
		{
			yyVAL.expr = BinaryExpr{Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:285
		{
			yyVAL.expr = BinaryExpr{Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:289
		{
			yyVAL.expr = BinaryExpr{Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:294
		{
			yyVAL.expr = nil
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:298
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:305
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:309
		{
			yyVAL.expr = StringLit(yyDollar[1].sym)
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:313
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:317
		{
			yyVAL.expr = UnaryExpr{Minus, yyDollar[2].expr}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:321
		{
			yyVAL.expr = UnaryExpr{Not, yyDollar[2].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:325
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:329
		{
			yyVAL.expr = Ident{ast, yyDollar[1].sym}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:333
		{
			yyVAL.expr = CallExpr{parser.Writer, yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 61:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:337
		{
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:342
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:346
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:351
		{
			yyVAL.exprlist = nil
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:355
		{
			yyVAL.exprlist = yyDollar[1].exprlist
		}
	}
	goto yystack /* stack new state and value */
}
