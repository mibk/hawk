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
	yys         int
	num         int
	sym         string
	expr        Expr
	exprlist    []Expr
	stmt        Stmt
	stmtlist    []Stmt
	paction     PatternAction
	pactionlist []PatternAction
}

const NUM = 57346
const IDENT = 57347
const EQ = 57348
const NE = 57349
const LE = 57350
const GE = 57351
const LT = 57352
const GT = 57353

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUM",
	"IDENT",
	"EQ",
	"NE",
	"LE",
	"GE",
	"LT",
	"GT",
	"';'",
	"'{'",
	"'}'",
	"'$'",
	"'('",
	"')'",
	"','",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line hawk.y:158

func Compile(src []byte, p *parse.Parser) *Tree {
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
}

const yyNprod = 25
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 51

var yyAct = [...]int{

	4, 21, 36, 37, 23, 32, 11, 31, 1, 22,
	2, 9, 10, 6, 25, 26, 27, 28, 29, 30,
	8, 19, 7, 20, 34, 13, 14, 15, 16, 17,
	18, 33, 8, 22, 35, 9, 10, 3, 38, 13,
	14, 15, 16, 17, 18, 5, 7, 9, 10, 24,
	12,
}
var yyPact = [...]int{

	7, -1000, -6, -1000, 19, -1000, -1000, 43, 31, -1000,
	-12, 7, -1000, 31, 31, 31, 31, 31, 31, -1000,
	-7, -1000, 33, 31, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 31, -15, 33, -1000, -1000, 31, 33,
}
var yyPgo = [...]int{

	0, 0, 13, 31, 1, 23, 45, 37, 10, 8,
}
var yyR1 = [...]int{

	0, 9, 8, 8, 7, 7, 7, 6, 5, 5,
	5, 4, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 3, 3, 3,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 2, 1, 1, 3, 0, 1,
	3, 1, 1, 2, 3, 3, 3, 3, 3, 3,
	1, 4, 0, 1, 3,
}
var yyChk = [...]int{

	-1000, -9, -8, -7, -1, -6, -2, 15, 13, 4,
	5, 12, -6, 6, 7, 8, 9, 10, 11, -2,
	-5, -4, -1, 16, -7, -1, -1, -1, -1, -1,
	-1, 14, 12, -3, -1, -4, 17, 18, -1,
}
var yyDef = [...]int{

	0, -2, 1, 2, 5, 6, 12, 0, 8, 20,
	0, 0, 4, 0, 0, 0, 0, 0, 0, 13,
	0, 9, 11, 22, 3, 14, 15, 16, 17, 18,
	19, 7, 0, 0, 23, 10, 21, 0, 24,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 15, 3, 3, 3,
	16, 17, 3, 3, 18, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 12,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 13, 3, 14,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
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
		//line hawk.y:46
		{
			ast = &Tree{yyDollar[1].pactionlist}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:52
		{
			yyVAL.pactionlist = append(yyVAL.pactionlist, yyDollar[1].paction)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:56
		{
			yyVAL.pactionlist = append(yyDollar[1].pactionlist, yyDollar[3].paction)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:62
		{
			yyVAL.paction = PatternAction{yyDollar[1].expr, BlockStmt{yyDollar[2].stmtlist}}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:66
		{
			yyVAL.paction = PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:70
		{
			yyVAL.paction = PatternAction{Lit(1), BlockStmt{yyDollar[1].stmtlist}}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:76
		{
			yyVAL.stmtlist = yyDollar[2].stmtlist
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:81
		{
			yyVAL.stmtlist = nil
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:85
		{
			yyVAL.stmtlist = append(yyVAL.stmtlist, yyDollar[1].stmt)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:89
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:95
		{
			yyVAL.stmt = ExprStmt{yyDollar[1].expr}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:103
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:107
		{
			yyVAL.expr = Col{parser, yyDollar[2].expr}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:111
		{
			yyVAL.expr = BinaryOp{EQ, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:115
		{
			yyVAL.expr = BinaryOp{NE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:119
		{
			yyVAL.expr = BinaryOp{LE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:123
		{
			yyVAL.expr = BinaryOp{GE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:127
		{
			yyVAL.expr = BinaryOp{LT, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:131
		{
			yyVAL.expr = BinaryOp{GT, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:137
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:141
		{
			yyVAL.expr = CallExpr{parser.Writer, yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:146
		{
			yyVAL.exprlist = nil
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:150
		{
			yyVAL.exprlist = append(yyVAL.exprlist, yyDollar[1].expr)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:154
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
