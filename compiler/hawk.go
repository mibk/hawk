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
	ast    *Root

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
const OROR = 57359
const ANDAND = 57360
const EQ = 57361
const NE = 57362
const LE = 57363
const GE = 57364

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

//line hawk.y:333

func Compile(r io.Reader, p *parse.Parser) (*Root, error) {
	ast = NewRoot()
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
	31, 46,
	-2, 5,
	-1, 52,
	31, 45,
	-2, 21,
	-1, 79,
	31, 46,
	-2, 12,
}

const yyNprod = 60
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 234

var yyAct = [...]int{

	44, 5, 18, 46, 76, 4, 68, 81, 80, 33,
	34, 40, 19, 98, 78, 51, 39, 86, 5, 72,
	73, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 43, 91, 19, 71, 70, 17,
	40, 69, 28, 29, 30, 31, 32, 30, 31, 32,
	94, 69, 75, 79, 88, 3, 1, 74, 77, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 84, 41, 2, 42, 8, 47, 85, 66,
	87, 95, 89, 90, 67, 83, 35, 93, 0, 36,
	37, 38, 92, 0, 97, 96, 0, 0, 10, 45,
	11, 50, 100, 99, 51, 0, 52, 48, 49, 0,
	0, 0, 0, 0, 0, 10, 45, 11, 50, 12,
	13, 51, 0, 52, 48, 49, 82, 0, 9, 14,
	15, 10, 16, 11, 0, 0, 12, 13, 10, 16,
	11, 0, 6, 7, 0, 9, 14, 15, 0, 0,
	0, 0, 12, 13, 0, 0, 0, 0, 0, 12,
	13, 9, 14, 15, 0, 10, 16, 11, 9, 14,
	15, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 0, 19, 12, 13, 0, 0,
	0, 0, 0, 0, 0, 0, 14, 15, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32,
}
var yyPact = [...]int{

	134, -1000, 9, -1000, -19, 181, -19, -19, -1000, 161,
	-1000, -1000, 161, 161, 161, 127, -25, 134, -1000, 111,
	127, 127, 127, 127, 127, 127, 127, 127, 127, 127,
	127, 127, 127, -1000, -1000, -1000, -1000, -1000, -1000, 42,
	127, -1000, 8, -1000, 181, 4, -1000, -1000, -1000, -1000,
	127, 127, 111, 193, 204, 17, 17, 17, 17, 17,
	17, 20, 20, -1000, -1000, -1000, -1000, -29, -31, 181,
	94, 127, -1000, -1000, -31, 154, -13, -19, -1000, 181,
	-1000, 127, -1000, -1000, 181, 24, 127, -1000, 45, 181,
	-1000, 5, -17, 181, -1000, -1000, -1000, -1000, 111, -19,
	-1000,
}
var yyPgo = [...]int{

	0, 0, 5, 76, 6, 84, 14, 4, 55, 3,
	83, 81, 77, 75, 2, 74, 56, 54,
}
var yyR1 = [...]int{

	0, 16, 15, 15, 8, 8, 8, 8, 14, 13,
	13, 13, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 7, 7, 9, 10, 10, 11, 11, 12, 12,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 17, 3, 4, 4, 5, 5,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 2, 1, 2, 2, 4, 0,
	1, 3, 1, 3, 2, 2, 1, 1, 1, 1,
	2, 0, 1, 4, 0, 2, 1, 1, 7, 3,
	1, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 0, 1, 1, 1, 2,
	2, 2, 3, 1, 0, 6, 1, 3, 0, 1,
}
var yyChk = [...]int{

	-1000, -16, -15, -8, -2, -1, 8, 9, -3, 34,
	4, 6, 25, 26, 35, 36, 5, 30, -14, 31,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 29, -14, -14, -3, -3, -3, -3, -1,
	36, -8, -13, -6, -1, 5, -9, -12, 13, 14,
	7, 10, 12, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, 37, -5, -4, -1,
	30, 33, 15, 16, -4, -1, -7, -2, -6, -1,
	37, 38, 32, -6, -1, -14, 30, -14, -17, -1,
	-10, 11, -2, -1, 5, -11, -9, -14, 30, -7,
	-14,
}
var yyDef = [...]int{

	45, -2, 1, 2, 0, -2, 0, 0, 30, 0,
	47, 48, 0, 0, 0, 0, 53, 45, 4, 9,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 6, 7, 31, 49, 50, 51, 0,
	58, 3, 0, 10, 12, 53, 16, 17, 18, 19,
	0, 0, -2, 32, 33, 34, 35, 36, 37, 38,
	39, 40, 41, 42, 43, 44, 52, 0, 59, 56,
	0, 0, 14, 15, 20, 0, 0, 0, 22, -2,
	54, 0, 8, 11, 13, 24, 45, 29, 0, 57,
	23, 0, 0, 46, 55, 25, 26, 27, 21, 0,
	28,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 3, 3, 34, 29, 3, 3,
	36, 37, 27, 25, 38, 26, 3, 28, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 30,
	23, 33, 24, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 31, 3, 32,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22,
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
		//line hawk.y:54
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
		//line hawk.y:75
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:79
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:85
		{
			yyVAL.stmt = PatternAction{yyDollar[1].expr, BlockStmt{yyDollar[2].stmtlist}}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:89
		{
			yyVAL.stmt = PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:93
		{
			yyVAL.stmt = BeginAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:97
		{
			yyVAL.stmt = EndAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:103
		{
			yyVAL.stmtlist = yyDollar[2].stmtlist
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:108
		{
			yyVAL.stmtlist = nil
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:112
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:116
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:122
		{
			yyVAL.stmt = ExprStmt{yyDollar[1].expr}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:126
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, yyDollar[3].expr}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:132
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Add, Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:136
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{Sub, Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:140
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:144
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:148
		{
			yyVAL.stmt = StatusStmt{StatusBreak}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:152
		{
			yyVAL.stmt = StatusStmt{StatusContinue}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:156
		{
			yyVAL.stmt = CallStmt{parser.Writer, yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:161
		{
			yyVAL.stmt = nil
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:165
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:171
		{
			yyVAL.stmt = IfStmt{yyDollar[2].expr, BlockStmt{yyDollar[3].stmtlist}, yyDollar[4].stmt}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:176
		{
			yyVAL.stmt = nil
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:180
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:186
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:190
		{
			yyVAL.stmt = BlockStmt{yyDollar[1].stmtlist}
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:196
		{
			yyVAL.stmt = ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, BlockStmt{yyDollar[7].stmtlist}}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:200
		{
			yyVAL.stmt = ForStmt{nil, yyDollar[2].expr, nil, BlockStmt{yyDollar[3].stmtlist}}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:207
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:211
		{
			yyVAL.expr = FieldExpr{parser, yyDollar[2].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:215
		{
			yyVAL.expr = BinaryExpr{OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:219
		{
			yyVAL.expr = BinaryExpr{AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:223
		{
			yyVAL.expr = BinaryExpr{Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:227
		{
			yyVAL.expr = BinaryExpr{NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:231
		{
			yyVAL.expr = BinaryExpr{LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:235
		{
			yyVAL.expr = BinaryExpr{GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:239
		{
			yyVAL.expr = BinaryExpr{Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:243
		{
			yyVAL.expr = BinaryExpr{Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:247
		{
			yyVAL.expr = BinaryExpr{Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:251
		{
			yyVAL.expr = BinaryExpr{Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:255
		{
			yyVAL.expr = BinaryExpr{Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:259
		{
			yyVAL.expr = BinaryExpr{Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:263
		{
			yyVAL.expr = BinaryExpr{Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:268
		{
			yyVAL.expr = nil
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:272
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:279
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:283
		{
			yyVAL.expr = StringLit(yyDollar[1].sym)
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:287
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:291
		{
			yyVAL.expr = UnaryExpr{Minus, yyDollar[2].expr}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:295
		{
			yyVAL.expr = UnaryExpr{Not, yyDollar[2].expr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:299
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:303
		{
			yyVAL.expr = Ident{ast, yyDollar[1].sym}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:307
		{
			yyVAL.expr = CallExpr{parser.Writer, yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:311
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:316
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:320
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:325
		{
			yyVAL.exprlist = nil
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:329
		{
			yyVAL.exprlist = yyDollar[1].exprlist
		}
	}
	goto yystack /* stack new state and value */
}
