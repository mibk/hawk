//line hawk.y:2
package compiler

import __yyfmt__ "fmt"

//line hawk.y:2
import (
	"github.com/mibk/hawk/parse"
)

var (
	parser *parse.Parser
	ast    *Tree

	defaultAction BlockStmt
)

//line hawk.y:17
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
const BEGIN = 57348
const END = 57349
const IF = 57350
const ELSE = 57351
const FOR = 57352
const EQ = 57353
const NE = 57354
const LE = 57355
const GE = 57356
const LT = 57357
const GT = 57358

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUM",
	"IDENT",
	"BEGIN",
	"END",
	"IF",
	"ELSE",
	"FOR",
	"EQ",
	"NE",
	"LE",
	"GE",
	"LT",
	"GT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"';'",
	"'{'",
	"'}'",
	"'='",
	"'$'",
	"'('",
	"')'",
	"','",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line hawk.y:279

func Compile(src []byte, p *parse.Parser) *Tree {
	ast = NewTree()
	parser = p
	defaultAction = BlockStmt{[]Stmt{
		ExprStmt{CallExpr{parser.Writer, "print", []Expr{Col{parser, Lit(0)}}}},
	}}
	yyParse(&yyLex{src: src})
	return ast
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 38,
	22, 38,
	-2, 17,
	-1, 60,
	22, 39,
	-2, 13,
}

const yyNprod = 49
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 153

var yyAct = [...]int{

	33, 4, 5, 35, 57, 58, 55, 17, 42, 28,
	29, 70, 71, 42, 10, 41, 37, 4, 59, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 32,
	10, 24, 25, 26, 27, 73, 80, 68, 56, 60,
	54, 11, 15, 63, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 12, 13, 66, 11, 15, 67,
	61, 69, 9, 14, 26, 27, 1, 16, 2, 75,
	12, 13, 76, 65, 74, 3, 79, 78, 31, 14,
	8, 36, 11, 34, 82, 81, 37, 77, 38, 72,
	30, 62, 43, 39, 40, 12, 13, 11, 15, 6,
	7, 64, 0, 9, 14, 0, 0, 0, 11, 34,
	12, 13, 37, 0, 38, 10, 0, 0, 9, 14,
	0, 12, 13, 0, 0, 0, 0, 0, 0, 9,
	14, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 0, 10, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27,
}
var yyPact = [...]int{

	93, -1000, 46, -1000, 120, -1000, -8, -8, -1000, 53,
	104, -1000, 53, 53, 37, -13, 93, -1000, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, -1000, -1000,
	-1000, 19, -1000, 132, -18, -1000, -1000, 37, 104, -1000,
	-1000, 33, 37, -1000, 14, 14, 14, 14, 14, 14,
	45, 45, -1000, -1000, 78, 37, 120, 16, -8, -1000,
	132, -1000, -16, 132, -1000, -1000, 132, 26, 37, -1000,
	-1000, 37, -1000, 8, 15, 132, 132, -1000, -1000, -1000,
	104, -8, -1000,
}
var yyPgo = [...]int{

	0, 0, 5, 80, 91, 18, 4, 75, 3, 89,
	87, 81, 78, 2, 68, 66,
}
var yyR1 = [...]int{

	0, 15, 14, 14, 7, 7, 7, 7, 7, 13,
	12, 12, 12, 5, 5, 5, 5, 6, 6, 8,
	9, 9, 10, 10, 11, 11, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	3, 3, 3, 3, 3, 3, 4, 4, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 2, 1, 1, 2, 2, 4,
	0, 1, 3, 1, 3, 1, 1, 0, 1, 4,
	0, 2, 1, 1, 7, 3, 1, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 0, 1,
	1, 2, 2, 3, 1, 4, 0, 1, 3,
}
var yyChk = [...]int{

	-1000, -15, -14, -7, -1, -13, 6, 7, -3, 25,
	22, 4, 17, 18, 26, 5, 21, -13, 11, 12,
	13, 14, 15, 16, 17, 18, 19, 20, -13, -13,
	-3, -12, -5, -1, 5, -8, -11, 8, 10, -3,
	-3, -1, 26, -7, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, 21, 24, -1, -6, -2, -5,
	-1, 27, -4, -1, 23, -5, -1, -13, 21, -13,
	27, 28, -9, 9, -2, -1, -1, -10, -8, -13,
	21, -6, -13,
}
var yyDef = [...]int{

	0, -2, 1, 2, 5, 6, 0, 0, 26, 0,
	10, 40, 0, 0, 0, 44, 0, 4, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 7, 8,
	27, 0, 11, 13, 44, 15, 16, 0, -2, 41,
	42, 0, 46, 3, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 0, 0, 0, 0, 0, 18,
	-2, 43, 0, 47, 9, 12, 14, 20, 38, 25,
	45, 0, 19, 0, 0, 39, 48, 21, 22, 23,
	17, 0, 24,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 25, 3, 3, 3,
	26, 27, 19, 17, 28, 18, 3, 20, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 21,
	3, 24, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 22, 3, 23,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
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
		//line hawk.y:47
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
		//line hawk.y:68
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:72
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:78
		{
			yyVAL.stmt = PatternAction{yyDollar[1].expr, BlockStmt{yyDollar[2].stmtlist}}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:82
		{
			yyVAL.stmt = PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:86
		{
			yyVAL.stmt = PatternAction{Lit(1), BlockStmt{yyDollar[1].stmtlist}}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:90
		{
			yyVAL.stmt = BeginAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:94
		{
			yyVAL.stmt = EndAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:100
		{
			yyVAL.stmtlist = yyDollar[2].stmtlist
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:105
		{
			yyVAL.stmtlist = nil
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:109
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:113
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:119
		{
			yyVAL.stmt = ExprStmt{yyDollar[1].expr}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:123
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:127
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:131
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:136
		{
			yyVAL.stmt = nil
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:140
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:146
		{
			yyVAL.stmt = IfStmt{yyDollar[2].expr, BlockStmt{yyDollar[3].stmtlist}, yyDollar[4].stmt}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:151
		{
			yyVAL.stmt = nil
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:155
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:161
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:165
		{
			yyVAL.stmt = BlockStmt{yyDollar[1].stmtlist}
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:171
		{
			yyVAL.stmt = ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, BlockStmt{yyDollar[7].stmtlist}}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:175
		{
			yyVAL.stmt = ForStmt{nil, yyDollar[2].expr, nil, BlockStmt{yyDollar[3].stmtlist}}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:182
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:186
		{
			yyVAL.expr = Col{parser, yyDollar[2].expr}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:190
		{
			yyVAL.expr = BinaryExpr{EQ, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:194
		{
			yyVAL.expr = BinaryExpr{NE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:198
		{
			yyVAL.expr = BinaryExpr{LE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:202
		{
			yyVAL.expr = BinaryExpr{GE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:206
		{
			yyVAL.expr = BinaryExpr{LT, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:210
		{
			yyVAL.expr = BinaryExpr{GT, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:214
		{
			yyVAL.expr = BinaryExpr{ADD, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:218
		{
			yyVAL.expr = BinaryExpr{SUB, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:222
		{
			yyVAL.expr = BinaryExpr{MUL, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:226
		{
			yyVAL.expr = BinaryExpr{DIV, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:231
		{
			yyVAL.expr = nil
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:235
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:242
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:246
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:250
		{
			yyVAL.expr = UnaryExpr{SUB, yyDollar[2].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:254
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:258
		{
			yyVAL.expr = Ident{ast, yyDollar[1].sym}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:262
		{
			yyVAL.expr = CallExpr{parser.Writer, yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:267
		{
			yyVAL.exprlist = nil
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:271
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:275
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
