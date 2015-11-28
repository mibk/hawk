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
const STRING = 57348
const PRINT = 57349
const BEGIN = 57350
const END = 57351
const IF = 57352
const ELSE = 57353
const FOR = 57354
const INC = 57355
const DEC = 57356
const LOR = 57357
const LAND = 57358
const EQ = 57359
const NE = 57360
const LE = 57361
const GE = 57362
const LT = 57363
const GT = 57364

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
	"INC",
	"DEC",
	"LOR",
	"LAND",
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
	"'%'",
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

//line hawk.y:317

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
	-1, 5,
	29, 44,
	-2, 5,
	-1, 48,
	29, 43,
	-2, 19,
	-1, 75,
	29, 44,
	-2, 12,
}

const yyNprod = 57
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 235

var yyAct = [...]int{

	42, 5, 17, 44, 72, 4, 64, 77, 76, 32,
	33, 38, 18, 94, 74, 37, 47, 5, 68, 69,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 41, 87, 18, 67, 82, 38, 65,
	27, 28, 29, 30, 31, 84, 66, 65, 71, 75,
	29, 30, 31, 70, 73, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 80, 18,
	16, 90, 3, 1, 81, 2, 83, 40, 85, 45,
	91, 79, 8, 89, 86, 63, 0, 0, 88, 39,
	93, 92, 34, 0, 0, 35, 36, 0, 96, 95,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 0, 10, 43, 11, 46, 0, 62,
	47, 0, 48, 0, 0, 0, 0, 0, 0, 10,
	43, 11, 46, 12, 13, 47, 0, 48, 0, 0,
	78, 0, 9, 14, 10, 15, 11, 0, 12, 13,
	10, 15, 11, 0, 6, 7, 0, 9, 14, 0,
	0, 0, 0, 12, 13, 0, 0, 0, 0, 12,
	13, 0, 9, 14, 0, 0, 0, 0, 9, 14,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 29, 30, 31, 10, 15, 11, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 30, 31, 0,
	0, 0, 0, 0, 12, 13, 0, 0, 0, 0,
	0, 0, 0, 0, 14,
}
var yyPact = [...]int{

	146, -1000, 42, -1000, -17, 165, -17, -17, -1000, 201,
	-1000, -1000, 201, 201, 140, -22, 146, -1000, 125, 140,
	140, 140, 140, 140, 140, 140, 140, 140, 140, 140,
	140, 140, -1000, -1000, -1000, -1000, -1000, 85, 140, -1000,
	18, -1000, 165, 5, -1000, -1000, 140, 140, 125, 177,
	191, 17, 17, 17, 17, 17, 17, 25, 25, -1000,
	-1000, -1000, -1000, -26, -28, 165, 110, 140, -1000, -1000,
	-28, 40, 9, -17, -1000, 165, -1000, 140, -1000, -1000,
	165, 23, 140, -1000, 66, 165, -1000, 6, -15, 165,
	-1000, -1000, -1000, -1000, 125, -17, -1000,
}
var yyPgo = [...]int{

	0, 0, 5, 82, 6, 85, 14, 4, 72, 3,
	84, 80, 79, 77, 2, 75, 73, 45,
}
var yyR1 = [...]int{

	0, 16, 15, 15, 8, 8, 8, 8, 14, 13,
	13, 13, 6, 6, 6, 6, 6, 6, 6, 7,
	7, 9, 10, 10, 11, 11, 12, 12, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 17, 3, 4, 4, 5, 5,
}
var yyR2 = [...]int{

	0, 1, 1, 3, 2, 1, 2, 2, 4, 0,
	1, 3, 1, 3, 2, 2, 1, 1, 2, 0,
	1, 4, 0, 2, 1, 1, 7, 3, 1, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 0, 1, 1, 1, 2, 2, 3,
	1, 0, 6, 1, 3, 0, 1,
}
var yyChk = [...]int{

	-1000, -16, -15, -8, -2, -1, 8, 9, -3, 32,
	4, 6, 23, 24, 33, 5, 28, -14, 29, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, -14, -14, -3, -3, -3, -1, 33, -8,
	-13, -6, -1, 5, -9, -12, 7, 10, 12, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, 34, -5, -4, -1, 28, 31, 13, 14,
	-4, -1, -7, -2, -6, -1, 34, 35, 30, -6,
	-1, -14, 28, -14, -17, -1, -10, 11, -2, -1,
	5, -11, -9, -14, 28, -7, -14,
}
var yyDef = [...]int{

	43, -2, 1, 2, 0, -2, 0, 0, 28, 0,
	45, 46, 0, 0, 0, 50, 43, 4, 9, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 6, 7, 29, 47, 48, 0, 55, 3,
	0, 10, 12, 50, 16, 17, 0, 0, -2, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 49, 0, 56, 53, 0, 0, 14, 15,
	18, 0, 0, 0, 20, -2, 51, 0, 8, 11,
	13, 22, 43, 27, 0, 54, 21, 0, 0, 44,
	52, 23, 24, 25, 19, 0, 26,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 32, 27, 3, 3,
	33, 34, 25, 23, 35, 24, 3, 26, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 28,
	3, 31, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 29, 3, 30,
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
		//line hawk.y:50
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
		//line hawk.y:71
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:75
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:81
		{
			yyVAL.stmt = PatternAction{yyDollar[1].expr, BlockStmt{yyDollar[2].stmtlist}}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:85
		{
			yyVAL.stmt = PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:89
		{
			yyVAL.stmt = BeginAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:93
		{
			yyVAL.stmt = EndAction{BlockStmt{yyDollar[2].stmtlist}}
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:99
		{
			yyVAL.stmtlist = yyDollar[2].stmtlist
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:104
		{
			yyVAL.stmtlist = nil
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:108
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:112
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:118
		{
			yyVAL.stmt = ExprStmt{yyDollar[1].expr}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:122
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, yyDollar[3].expr}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:128
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{ADD, Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:132
		{
			yyVAL.stmt = AssignStmt{ast, yyDollar[1].sym, BinaryExpr{SUB, Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:136
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:140
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:144
		{
			yyVAL.stmt = CallStmt{parser.Writer, yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:149
		{
			yyVAL.stmt = nil
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:153
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:159
		{
			yyVAL.stmt = IfStmt{yyDollar[2].expr, BlockStmt{yyDollar[3].stmtlist}, yyDollar[4].stmt}
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:164
		{
			yyVAL.stmt = nil
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:168
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:174
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:178
		{
			yyVAL.stmt = BlockStmt{yyDollar[1].stmtlist}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:184
		{
			yyVAL.stmt = ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, BlockStmt{yyDollar[7].stmtlist}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:188
		{
			yyVAL.stmt = ForStmt{nil, yyDollar[2].expr, nil, BlockStmt{yyDollar[3].stmtlist}}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:195
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:199
		{
			yyVAL.expr = Col{parser, yyDollar[2].expr}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:203
		{
			yyVAL.expr = BinaryExpr{LOR, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:207
		{
			yyVAL.expr = BinaryExpr{LAND, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:211
		{
			yyVAL.expr = BinaryExpr{EQ, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:215
		{
			yyVAL.expr = BinaryExpr{NE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:219
		{
			yyVAL.expr = BinaryExpr{LE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:223
		{
			yyVAL.expr = BinaryExpr{GE, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:227
		{
			yyVAL.expr = BinaryExpr{LT, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:231
		{
			yyVAL.expr = BinaryExpr{GT, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:235
		{
			yyVAL.expr = BinaryExpr{ADD, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:239
		{
			yyVAL.expr = BinaryExpr{SUB, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:243
		{
			yyVAL.expr = BinaryExpr{MUL, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:247
		{
			yyVAL.expr = BinaryExpr{DIV, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:251
		{
			yyVAL.expr = BinaryExpr{MOD, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:256
		{
			yyVAL.expr = nil
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:260
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:267
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:271
		{
			yyVAL.expr = StringLit(yyDollar[1].sym)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:275
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:279
		{
			yyVAL.expr = UnaryExpr{SUB, yyDollar[2].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:283
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:287
		{
			yyVAL.expr = Ident{ast, yyDollar[1].sym}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:291
		{
			yyVAL.expr = CallExpr{parser.Writer, yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 52:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:295
		{
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:300
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:304
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:309
		{
			yyVAL.exprlist = nil
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:313
		{
			yyVAL.exprlist = yyDollar[1].exprlist
		}
	}
	goto yystack /* stack new state and value */
}
